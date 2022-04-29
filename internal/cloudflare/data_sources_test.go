package cloudflare_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/peknur/terraform-provider-nds/internal/acctest"
)

func TestAccIpRanges(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "nds_cloudflare_ip_ranges" "ipv4" {
					ip_version = 4
				}
				`,
				Check: resource.TestCheckResourceAttr("data.nds_cloudflare_ip_ranges.ipv4", "data.0", "103.21.244.0/22"),
			},
			{
				Config: `
				data "nds_cloudflare_ip_ranges" "ipv4" {
				}
				`,
				Check: resource.TestCheckResourceAttr("data.nds_cloudflare_ip_ranges.ipv4", "data.0", "103.21.244.0/22"),
			},
			{
				Config: `
				data "nds_cloudflare_ip_ranges" "ipv6" {
					ip_version = 6
				}
				`,
				Check: resource.TestCheckResourceAttr("data.nds_cloudflare_ip_ranges.ipv6", "data.0", "2400:cb00::/32"),
			},
		},
	},
	)
}
