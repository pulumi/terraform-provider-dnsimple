package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/dnsimple/terraform-provider-dnsimple/dnsimple"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dnsimple.Provider})
}
