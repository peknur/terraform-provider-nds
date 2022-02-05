package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/peknur/terraform-provider-nds/internal/nslookup"
)

func New() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"nds_nslookup_ip":  nslookup.DataSourceLookupIP(),
			"nds_nslookup_mx":  nslookup.DataSourceLookupMX(),
			"nds_nslookup_srv": nslookup.DataSourceLookupSRV(),
			"nds_nslookup_txt": nslookup.DataSourceLookupTXT(),
			"nds_nslookup_ptr": nslookup.DataSourceLookupPTR(),
			"nds_nslookup_ns":  nslookup.DataSourceLookupNS(),
		},
	}
}
