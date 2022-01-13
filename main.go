package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/sethvargo/terraform-provider-filesystem/filesystem"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: filesystem.Provider,
	})
}
