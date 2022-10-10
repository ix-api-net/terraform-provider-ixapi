package resources

import (
	"context"
	"testing"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/testdata"
)

func TestContactCreate(t *testing.T) {
	resource := NewContactResource()
	res := resource.Data(nil)

	// Set request
	res.Set("roles", []string{"noc"})
	res.Set("email", "email@addr")

	contact := testdata.NewContact()
	roles := []*ixapi.Role{testdata.NewRoleNOC()}
	assignment := testdata.NewRoleAssignment()
	assignments := []*ixapi.RoleAssignment{assignment}
	reqs := 0
	api := ixapi.NewTestClient(map[string]any{
		"/roles": roles,
		"/contacts": ixapi.TestResponseFunc(func(body []byte) (any, error) {
			ixapi.AssertBodyContains(t, body, `email@addr`)
			return contact, nil
		}),
		"/contacts/2342": contact,
		"/role-assignments": ixapi.TestResponseFunc(func(body []byte) (any, error) {
			if reqs == 0 {
				// CREATE responds with a single result
				return assignment, nil
			}
			reqs++
			return assignments, nil
		}),
		"/role-assignments/1111": assignment,
	})

	diags := contactCreate(context.Background(), res, api)
	if diags != nil {
		t.Fatal(diags)
	}

	if res.Get("name").(string) != "contact name" {
		t.Error("unexpected contact resource", res)
	}

	if res.Id() != "2342" {
		t.Error("unexpected ID assigned to contact:", res.Id())
	}
}
