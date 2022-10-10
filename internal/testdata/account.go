package testdata

import (
	"time"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
)

// NewAddress creates a new test address
func NewAddress() *ixapi.Address {
	return &ixapi.Address{
		Country:             "Germany",
		Locality:            "Berlin",
		Region:              NewOptString("Region"),
		PostalCode:          "10117",
		StreetAddress:       "Straßenweg 9",
		PostOfficeBoxNumber: NewOptString("PO 123"),
	}
}

// NewBillingInformation creates new test billing info
func NewBillingInformation() *ixapi.BillingInformation {
	return &ixapi.BillingInformation{
		Name:      "Billing LLC",
		Address:   NewAddress(),
		VatNumber: NewOptString("NL1235890"),
	}
}

// NewStatus creates a new status
func NewStatus() []*ixapi.Status {
	return []*ixapi.Status{
		&ixapi.Status{
			Severity:  6,
			Tag:       "info",
			Message:   "message",
			Timestamp: time.Now().UTC(),
		},
	}
}

// NewAccount creates a new test account
func NewAccount() *ixapi.Account {
	return &ixapi.Account{
		ID:                       "23",
		Name:                     "account name",
		State:                    NewOptString("production"),
		Status:                   NewStatus(),
		ManagingAccount:          NewOptString("MACCT-12345"),
		LegalName:                NewOptString("legal name"),
		BillingInformation:       NewBillingInformation(),
		ExternalRef:              NewOptString("ext ref"),
		Address:                  NewAddress(),
		MetroAreaNetworkPresence: []string{"FRA", "AMS", "LON"},
	}
}
