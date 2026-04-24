package testhelpers_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/ix-api-net/terraform-provider-ixapi/acceptance/decix/testhelpers"
)

func TestRequireTestEnv_returnsValue(t *testing.T) {
	t.Setenv("TEST_UTILS_VAR", "expected-value")
	got := testhelpers.RequireTestEnv(t, "TEST_UTILS_VAR")
	if got != "expected-value" {
		t.Errorf("got %q, want %q", got, "expected-value")
	}
}

func stateWithResources(names ...string) *terraform.State {
	resources := make(map[string]*terraform.ResourceState, len(names))
	for _, name := range names {
		resources[name] = &terraform.ResourceState{}
	}
	return &terraform.State{
		Modules: []*terraform.ModuleState{
			{Path: []string{"root"}, Resources: resources},
		},
	}
}

func TestNotExists_resourceAbsent(t *testing.T) {
	state := stateWithResources()
	if err := testhelpers.NotExists("ixapi_foo.bar")(state); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestNotExists_resourcePresent(t *testing.T) {
	state := stateWithResources("ixapi_foo.bar")
	if err := testhelpers.NotExists("ixapi_foo.bar")(state); err == nil {
		t.Error("expected error, got nil")
	}
}

func TestNotExists_otherResourcePresent(t *testing.T) {
	state := stateWithResources("ixapi_other.thing")
	if err := testhelpers.NotExists("ixapi_foo.bar")(state); err != nil {
		t.Errorf("expected nil when different resource exists, got %v", err)
	}
}
