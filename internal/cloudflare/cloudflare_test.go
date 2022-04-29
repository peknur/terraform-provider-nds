package cloudflare

import (
	"context"
	"testing"
)

func TestIPv4Ranges(t *testing.T) {
	got, err := IPv4Ranges(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(got) < 1 {
		t.Fatalf("IPv4Ranges failed empty slice")
	}
	want := "103.21.244.0/22"
	if want != got[0] {
		t.Errorf("IPv4Ranges first item failed want %s got %s", want, got[0])
	}
}

func TestIPv6Ranges(t *testing.T) {
	got, err := IPv6Ranges(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(got) < 1 {
		t.Fatalf("IPv6Ranges failed empty slice")
	}
	want := "2400:cb00::/32"
	if want != got[0] {
		t.Errorf("IPv6Ranges first item failed want %s got %s", want, got[0])
	}
}
