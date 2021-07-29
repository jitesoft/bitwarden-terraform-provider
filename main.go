package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/jitesoft/bitwarden-terraform-provider/pkg"
)

func main () {
	options := plugin.ServeOpts{
		ProviderFunc: pkg.Provider,
	}

	plugin.Serve(&options)
}
