package main

import (
	"github.com/jackap/terraform-provider-tfsystem/tfsystem"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: tfsystem.Provider})
}
