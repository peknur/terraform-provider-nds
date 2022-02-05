package nslookup_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/peknur/terraform-provider-nds/internal/acctest"
)

func TestAccIP(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "nds_nslookup_ip" "test" {
					name = "example.debsu.fi"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nds_nslookup_ip.test", "data.0", "127.0.0.1"),
					resource.TestCheckResourceAttr("data.nds_nslookup_ip.test", "data.#", "1"),
				),
			},
		}},
	)
}

func TestAccNS(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "nds_nslookup_ns" "test" {
					name = "debsu.fi"
				}
				`,
				Check: resource.TestCheckResourceAttr("data.nds_nslookup_ns.test", "data.#", "2"),
			},
		}},
	)
}

func TestAccTXT(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "nds_nslookup_txt" "test" {
					name = "example.debsu.fi"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nds_nslookup_txt.test", "data.0", "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s."),
					resource.TestCheckResourceAttr("data.nds_nslookup_txt.test", "data.#", "1"),
				),
			},
		}},
	)
}

func TestAccMX(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "nds_nslookup_mx" "test" {
					name = "example.debsu.fi"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nds_nslookup_mx.test", "data.0.priority", "1"),
					resource.TestCheckResourceAttr("data.nds_nslookup_mx.test", "data.0.host", "localhost.localdomain."),
					resource.TestCheckResourceAttr("data.nds_nslookup_mx.test", "data.#", "1"),
				),
			},
		}},
	)
}

func TestAccSRV(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "nds_nslookup_srv" "test" {
					name = "example.debsu.fi"
					proto = "tcp"
					service = "xmpp-client"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nds_nslookup_srv.test", "data.0.weight", "0"),
					resource.TestCheckResourceAttr("data.nds_nslookup_srv.test", "data.0.target", "example.com."),
					resource.TestCheckResourceAttr("data.nds_nslookup_srv.test", "data.0.priority", "5"),
					resource.TestCheckResourceAttr("data.nds_nslookup_srv.test", "data.0.port", "5222"),
					resource.TestCheckResourceAttr("data.nds_nslookup_srv.test", "data.#", "1"),
				),
			},
		}},
	)
}

func TestAccCustomResolver(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "nds_nslookup_ip" "my_ip" {
					name = "myip.opendns.com"
					resolver {
						addr = "208.67.222.222" # resolver1.opendns.com
					}
				}
				`,
				Check: resource.TestCheckResourceAttrSet("data.nds_nslookup_ip.my_ip", "data.0"),
			},
		}},
	)
}

func TestAccPTR(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "nds_nslookup_ptr" "test" {
					name = "1.1.1.1"
				}
				`,
				Check: resource.TestCheckResourceAttr("data.nds_nslookup_ptr.test", "data.0", "one.one.one.one."),
			},
		}},
	)
}
