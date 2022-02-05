package nslookup

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func commonDataSource() schema.Resource {
	var transportProtocols = []string{"udp", "tcp"}

	return schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"wait": {
				Description: "Wait N secnds for name to propagate",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
			},
			"resolver": {
				Description: "Use custom resolver",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr": {
							Description:      "IPv4 address",
							Type:             schema.TypeString,
							ValidateDiagFunc: validation.ToDiagFunc(validation.IsIPv4Address),
							Required:         true,
						},
						"port": {
							Description:      "TCP/UDP port",
							Type:             schema.TypeInt,
							ValidateDiagFunc: validation.ToDiagFunc(validation.IsPortNumber),
							Optional:         true,
							Default:          53,
						},
						"proto": {
							Description:      fmt.Sprintf("Transport protocol (%s)", strings.Join(transportProtocols, ", ")),
							Type:             schema.TypeString,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(transportProtocols, true)),
							Optional:         true,
							Default:          "udp",
						},
						"timeout": {
							Description: "Connection timeout in seconds",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     60,
						},
					},
				},
			},
			"data": {},
		},
	}
}

func DataSourceLookupIP() *schema.Resource {
	s := commonDataSource()
	s.Schema["data"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	s.ReadContext = readLookupIPContext
	return &s
}

func readLookupIPContext(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	l := NewLookupFromResourceData(ctx, d)
	name := strings.ToLower(d.Get("name").(string))
	records, err := l.Address(ctx, name)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(records) < 1 {
		return diag.Errorf("'%s' returned empty record set", name)
	}
	d.SetId(fmt.Sprintf("host.%s", name))
	return diag.FromErr(d.Set("data", records))
}

func DataSourceLookupTXT() *schema.Resource {
	s := DataSourceLookupIP()
	s.ReadContext = readLookupTXTContext
	return s
}

func readLookupTXTContext(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	l := NewLookupFromResourceData(ctx, d)
	name := strings.ToLower(d.Get("name").(string))
	records, err := l.Text(ctx, name)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(records) < 1 {
		return diag.Errorf("'%s' returned empty TXT record set", name)
	}
	d.SetId(fmt.Sprintf("txt.%s", name))
	return diag.FromErr(d.Set("data", records))
}

func DataSourceLookupPTR() *schema.Resource {
	s := DataSourceLookupIP()
	s.ReadContext = readLookupPTRContext
	return s
}

func readLookupPTRContext(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	l := NewLookupFromResourceData(ctx, d)
	name := strings.ToLower(d.Get("name").(string))
	records, err := l.Reverse(ctx, name)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(records) < 1 {
		return diag.Errorf("'%s' returned empty PTR record set", name)
	}
	d.SetId(fmt.Sprintf("ptr.%s", name))
	return diag.FromErr(d.Set("data", records))
}

func DataSourceLookupNS() *schema.Resource {
	s := DataSourceLookupIP()
	s.ReadContext = readLookupNSContext
	return s
}

func readLookupNSContext(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	l := NewLookupFromResourceData(ctx, d)
	name := strings.ToLower(d.Get("name").(string))
	records, err := l.Nameserver(ctx, name)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(records) < 1 {
		return diag.Errorf("'%s' returned empty NS record set", name)
	}
	d.SetId(fmt.Sprintf("ns.%s", name))
	return diag.FromErr(d.Set("data", records))
}

func DataSourceLookupMX() *schema.Resource {
	s := commonDataSource()
	s.Schema["data"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"host": {
					Type:     schema.TypeString,
					Required: true,
				},
				"priority": {
					Type:     schema.TypeInt,
					Optional: true,
				},
			},
		},
	}
	s.ReadContext = readLookupMXContext
	return &s
}

func readLookupMXContext(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	l := NewLookupFromResourceData(ctx, d)
	name := strings.ToLower(d.Get("name").(string))
	records, err := l.MailExchange(ctx, name)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(records) < 1 {
		return diag.Errorf("'%s' returned empty MX record set", name)
	}
	data := make([]map[string]interface{}, 0)
	for _, r := range records {
		data = append(data, map[string]interface{}{
			"host":     r.Host,
			"priority": r.Priority,
		})
	}
	d.SetId(fmt.Sprintf("mx.%s", name))
	return diag.FromErr(d.Set("data", data))
}

func DataSourceLookupSRV() *schema.Resource {
	protos := []string{"tcp", "udp"}
	s := commonDataSource()

	s.Schema["proto"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ValidateDiagFunc: validation.ToDiagFunc(
			validation.StringInSlice(protos, true)),
	}
	s.Schema["service"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	s.Schema["data"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"target": {
					Type:     schema.TypeString,
					Required: true,
				},
				"priority": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"weight": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"port": {
					Type:     schema.TypeInt,
					Optional: true,
				},
			},
		},
	}
	s.ReadContext = readLookupSRVContext
	return &s
}

func readLookupSRVContext(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	l := NewLookupFromResourceData(ctx, d)
	name := strings.ToLower(d.Get("name").(string))
	proto := strings.ToLower(d.Get("proto").(string))
	service := strings.ToLower(d.Get("service").(string))
	records, err := l.Service(ctx, service, proto, name)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(records) < 1 {
		return diag.Errorf("'%s' returned empty SRV record set", name)
	}
	data := make([]map[string]interface{}, 0)
	for _, r := range records {
		data = append(data, map[string]interface{}{
			"target":   r.Host,
			"priority": r.Priority,
			"weight":   r.Weight,
			"port":     r.Port,
		})
	}
	d.SetId(fmt.Sprintf("srv.%s.%s.%s", service, proto, name))
	return diag.FromErr(d.Set("data", data))
}
