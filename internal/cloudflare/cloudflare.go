package cloudflare

import (
	"bufio"
	"context"
	"net"
	"net/http"
	"sort"
)

const (
	ipv4URL string = "https://www.cloudflare.com/ips-v4"
	ipv6URL string = "https://www.cloudflare.com/ips-v6"
)

func IPv4Ranges(ctx context.Context) ([]string, error) {
	return ipRanges(ctx, ipv4URL)
}

func IPv6Ranges(ctx context.Context) ([]string, error) {
	return ipRanges(ctx, ipv6URL)
}

func ipRanges(ctx context.Context, url string) ([]string, error) {
	ips := make([]string, 0)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return ips, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ips, err
	}
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		_, net, err := net.ParseCIDR(scanner.Text())
		if err != nil {
			return ips, err
		}
		ips = append(ips, net.String())
	}
	sort.Strings(ips)
	return ips, nil
}
