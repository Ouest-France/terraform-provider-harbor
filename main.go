package main

import (
	"github.com/Ouest-France/terraform-provider-harbor/harbor"

	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return harbor.Provider()
		},
	})
}
