package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/pikacloud/terraform-provider-pikacloud/pikacloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pikacloud.Provider,
	})
}
