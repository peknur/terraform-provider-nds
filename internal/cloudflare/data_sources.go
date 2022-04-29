package cloudflare

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIPRanges() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_version": {
				Description: "IP version number (4 or 6)",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     4,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.IntInSlice([]int{4, 6}),
				),
			},
			"data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		ReadContext: readIPRangesContext,
	}
}

func readIPRangesContext(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var ips []string
	var err error
	version := d.Get("ip_version").(int)
	switch version {
	case 6:
		if ips, err = IPv6Ranges(ctx); err != nil {
			return diag.FromErr(err)
		}
	default:
		if ips, err = IPv4Ranges(ctx); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(fmt.Sprintf("ipv%d_range", version))

	return diag.FromErr(d.Set("data", ips))

}
