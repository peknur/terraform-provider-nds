package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/peknur/terraform-provider-nds/internal/provider"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return provider.New()
		},
	}

	if debugMode {
		if err := plugin.Debug(context.Background(), "github.com/peknur/terraform-provider-nds", opts); err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
