package main

import (
	"flag"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Run the docs generation tool:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
var (
	version string = "dev"
	commit  string = "head"
)

// Flags
var (
	debug bool
)

func init() {
	flag.BoolVar(
		&debug,
		"debug",
		false,
		"set to true to run the provider with support for debuggers like delve",
	)
}

func main() {
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debug,
		ProviderAddr: "registry.terraform.io/ix-api-net/terraform-provider-ixapi",
		ProviderFunc: provider.New(version),
	}

	plugin.Serve(opts)
}
