package provider

import (
	"testing"

	"github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/schemas"
)

func TestProvider(t *testing.T) {
	p := New("test")
	if p == nil {
		t.Fatal("could not create provider instance")
	}
}

func TestSchemaVersions(t *testing.T) {
	if ixapi.SchemaVersion != schemas.SchemaVersion {
		// In this case: Regenerate client / schema files
		t.Fatal("client schema version must match terraform schema version")
	}
}
