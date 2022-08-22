package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Support markdown syntax in description
	schema.DescriptionKind = schema.StringMarkdown

	// Add defaults to exported descriptions
	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		return strings.TrimSpace(desc)
	}
}

// configure initializes the API client and loads the credentials
func configure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(
		ctx context.Context,
		res *schema.ResourceData,
	) (any, diag.Diagnostics) {
		// Get API credentials and create client

		// ...
		return nil, nil
	}
}

// New creates a new provider function
func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{},
			ResourcesMap:   map[string]*schema.Resource{},
		}

		p.ConfigureContextFunc = configure(p)

		return p
	}
}
