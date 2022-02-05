package acctest

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/peknur/terraform-provider-nds/internal/provider"
)

const (
	ProviderName = "nds"
)

var ProviderFactories map[string]func() (*schema.Provider, error)

func init() {

	ProviderFactories = map[string]func() (*schema.Provider, error){
		ProviderName: func() (*schema.Provider, error) {
			return provider.New(), nil
		},
	}
}
