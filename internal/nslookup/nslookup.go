package nslookup

import (
	"context"
	"fmt"
	"net"
	"sort"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type MailExchange struct {
	Host     string
	Priority uint16
}

type ServiceTarget struct {
	Host     string
	Priority uint16
	Weight   uint16
	Port     uint16
}

// Lookup is currently thin wrapper around Go's net lookup functions with
// possibility to use custom resolver
type Lookup struct {
	resolver *net.Resolver
}

// Service returns service targest sorted by priority.
func (l *Lookup) Service(ctx context.Context, service, proto, name string) (targets []ServiceTarget, err error) {
	_, r, err := l.resolver.LookupSRV(ctx, service, proto, name)
	if err != nil {
		return targets, err
	}
	for _, s := range r {
		targets = append(targets, ServiceTarget{
			Priority: s.Priority,
			Weight:   s.Weight,
			Port:     s.Port,
			Host:     s.Target,
		})
	}
	sort.Slice(targets, func(i, j int) bool {
		return targets[i].Priority < targets[j].Priority
	})
	return targets, err
}

// MailExchange returns the DNS MX records for the given host name sorted by priority.
func (l *Lookup) MailExchange(ctx context.Context, host string) (hosts []MailExchange, err error) {
	r, err := l.resolver.LookupMX(ctx, host)
	if err != nil {
		return hosts, err
	}
	for _, m := range r {
		hosts = append(hosts, MailExchange{
			Host:     m.Host,
			Priority: m.Pref,
		})
	}
	return hosts, err
}

// Address returns a slice of host's IPv4 and IPv6 addresses.
func (l *Lookup) Address(ctx context.Context, host string) (addrs []string, err error) {
	r, err := l.resolver.LookupIPAddr(ctx, host)
	if err != nil {
		return addrs, err
	}
	for _, ip := range r {
		addrs = append(addrs, ip.String())
	}
	return addrs, err
}

// Reverse performs a reverse lookup for the given address, returning a list of names mapping to that address.
func (l *Lookup) Reverse(ctx context.Context, addr string) (names []string, err error) {
	return l.resolver.LookupAddr(ctx, addr)
}

// Text returns the DNS TXT records for the given domain name.
func (l *Lookup) Text(ctx context.Context, host string) (texts []string, err error) {
	return l.resolver.LookupTXT(ctx, host)
}

// Nameserver returns the DNS NS records for the given domain name.
func (l *Lookup) Nameserver(ctx context.Context, host string) (hosts []string, err error) {
	r, err := l.resolver.LookupNS(ctx, host)
	if err != nil {
		return hosts, err
	}
	for _, ns := range r {
		hosts = append(hosts, ns.Host)
	}
	return hosts, err
}

func NewLookup(ctx context.Context, proto string, addr string, timeout time.Duration) *Lookup {
	return &Lookup{resolver: &net.Resolver{
		PreferGo:     true, // needed for custom Dial to work
		StrictErrors: false,
		Dial: func(ctx context.Context, _, _ string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: timeout,
			}
			return d.DialContext(ctx, proto, addr)
		},
	}}
}

func NewLookupFromResourceData(ctx context.Context, d *schema.ResourceData) *Lookup {
	if _, ok := d.GetOk("resolver"); !ok {
		return &Lookup{resolver: net.DefaultResolver}
	}
	p := d.Get("resolver.0").(map[string]interface{})
	addr := fmt.Sprintf("%s:%d", p["addr"].(string), p["port"].(int))
	return NewLookup(ctx, p["proto"].(string), addr, time.Duration(p["timeout"].(int))*time.Second)
}
