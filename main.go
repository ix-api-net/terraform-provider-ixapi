package main

import (
	"flag"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Run the docs generation tool:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
var (
	version string = "dev"
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
		ProviderAddr: "ix-api.net/ix-api/ix-api", // for now
		ProviderFunc: provider.New(version),
	}

	plugin.Serve(opts)
}
