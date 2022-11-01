package provider

import (
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func TestProvider(t *testing.T) {
	p := New("dev")
	if err := p().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}

func TestSchemaVersions(t *testing.T) {
	if ixapi.SchemaVersion != schemas.SchemaVersion {
		// In this case: Regenerate client / schema files
		t.Fatal("client schema version must match terraform schema version")
	}
}
