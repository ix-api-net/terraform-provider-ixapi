package testdata

import (
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

// NewMemberJoiningRuleAllow creates a member joining rule
func NewMemberJoiningRuleAllow() *ixapi.AllowMemberJoiningRule {
	return &ixapi.AllowMemberJoiningRule{
		Type: ixapi.AllowMemberJoiningRuleType,

		ManagingAccount:  "managing:123",
		ConsumingAccount: "consuming:123",

		CapacityMax:    NewOptInt(2300),
		NetworkService: "ns:23",
	}
}

// NewMemberJoiningRuleDeny creates a member joining rule
func NewMemberJoiningRuleDeny() *ixapi.DenyMemberJoiningRule {
	return &ixapi.DenyMemberJoiningRule{
		Type: ixapi.DenyMemberJoiningRuleType,

		ManagingAccount:  "managing:123",
		ConsumingAccount: "consuming:123",

		NetworkService: "ns:23",
	}
}
