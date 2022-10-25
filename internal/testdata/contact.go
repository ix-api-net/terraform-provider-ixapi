package testdata

import "github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"

// NewContact builds an IX-API contact with test data
func NewContact() *ixapi.Contact {
	return &ixapi.Contact{
		ID:               "2342",
		ManagingAccount:  "23",
		ConsumingAccount: "42",
		Name:             NewOptString("contact name"),
		Telephone:        NewOptString("+23 42 1235"),
		Email:            NewOptString("support@customer"),
	}
}

// NewRoleNOC creates a new NOC role
func NewRoleNOC() *ixapi.Role {
	return &ixapi.Role{
		ID:             "452",
		Name:           "noc",
		RequiredFields: []string{"email"},
	}
}

// NewRoleAssignment creates a new test role assignment
func NewRoleAssignment() *ixapi.RoleAssignment {
	return &ixapi.RoleAssignment{
		Role:    "452",
		Contact: "2342",
		ID:      "1111",
	}
}
