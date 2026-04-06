package testdata

import (
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ptr"
)

// NewAddress creates a new test address
func NewAddress() *ixapi.Address {
	return &ixapi.Address{
		Country:             "Germany",
		Locality:            "Berlin",
		Region:              ptr.Of("Region"),
		PostalCode:          "10117",
		StreetAddress:       "Straßenweg 9",
		PostOfficeBoxNumber: ptr.Of("PO 123"),
	}
}

// NewBillingInformation creates new test billing info
func NewBillingInformation() *ixapi.BillingInformation {
	return &ixapi.BillingInformation{
		Name:      "Billing LLC",
		Address:   NewAddress(),
		VatNumber: ptr.Of("NL1235890"),
	}
}

// NewStatus creates a new status
func NewStatus() []*ixapi.Status {
	return []*ixapi.Status{
		&ixapi.Status{
			Severity:  6,
			Tag:       "info",
			Message:   "message",
			Timestamp: ixapi.APITimestampNowUTC(),
		},
	}
}

// NewAccount creates a new test account
func NewAccount() *ixapi.Account {
	return &ixapi.Account{
		ID:                       "23",
		Name:                     "account name",
		State:                    ptr.Of("production"),
		Status:                   NewStatus(),
		ManagingAccount:          ptr.Of("MACCT-12345"),
		LegalName:                ptr.Of("legal name"),
		BillingInformation:       NewBillingInformation(),
		ExternalRef:              ptr.Of("ext ref"),
		Address:                  NewAddress(),
		MetroAreaNetworkPresence: []string{"FRA", "AMS", "LON"},
	}
}
