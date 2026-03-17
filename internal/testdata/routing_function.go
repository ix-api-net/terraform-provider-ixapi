package testdata

import "github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"

// NewRoutingFunction makes a testing routing function
func NewRoutingFunction() *ixapi.RoutingFunction {
	return &ixapi.RoutingFunction{
		ID:               "rf:42",
		ManagingAccount:  "managing:123",
		ConsumingAccount: "consuming:123",
		BillingAccount:   "billing:123",
		ProductOffering:  "product:123",
		ASN:              65000,
		ExternalRef:      NewOptString("ext-ref-rf"),
		Capacity:         NewOptInt(1000),
		State:            "production",
	}
}
