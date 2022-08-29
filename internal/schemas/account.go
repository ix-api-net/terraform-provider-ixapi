package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
)

// AddressFromResource decodes embedded address infos
func AddressFromResource(res Resource) *ixapi.Address {
	if res == nil {
		return nil
	}
	return &ixapi.Address{
		Country:             res.GetString("country"),
		Locality:            res.GetString("locality"),
		Region:              res.GetStringOpt("region"),
		PostalCode:          res.GetString("postal_code"),
		StreetAddress:       res.GetString("street_address"),
		PostOfficeBoxNumber: res.GetStringOpt("post_office_box_number"),
	}
}

// FlattenAddress makes a flat address
func FlattenAddress(addr *ixapi.Address) []interface{} {
	if addr == nil {
		return []interface{}{}
	}

	a := map[string]interface{}{}
	a["country"] = addr.Country
	a["locality"] = addr.Locality
	if addr.Region != nil {
		a["region"] = *addr.Region
	}
	a["postal_code"] = addr.PostalCode
	a["street_address"] = addr.StreetAddress
	if addr.PostOfficeBoxNumber != nil {
		a["post_office_box_number"] = *addr.PostOfficeBoxNumber
	}

	return []interface{}{a}
}

// BillingInformationFromResource decodes billing information
// from an embedded resource
func BillingInformationFromResource(res Resource) *ixapi.BillingInformation {
	if res == nil {
		return nil
	}

	address := AddressFromResource(res.GetResource("address"))
	return &ixapi.BillingInformation{
		Name:      res.GetString("name"),
		VatNumber: res.GetStringOpt("vat_number"),
		Address:   address,
	}
}

// FlattenBillingInformation flattens the billing information
func FlattenBillingInformation(billing *ixapi.BillingInformation) []interface{} {
	if billing == nil {
		return []interface{}{}
	}

	b := map[string]interface{}{}
	b["name"] = billing.Name
	b["address"] = FlattenAddress(billing.Address)
	if billing.VatNumber != nil {
		b["vat_number"] = *billing.VatNumber
	}
	return []interface{}{b}
}

// AccountRequestFromResourceData builds an account create request
func AccountRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.AccountRequest, diag.Diagnostics) {
	res := ResourceData{r}

	billing := BillingInformationFromResource(
		res.GetResource("billing_information"))

	address := AddressFromResource(
		res.GetResource("address"))

	req := &ixapi.AccountRequest{
		ManagingAccount:    res.GetStringOpt("managing_account"),
		Name:               res.GetString("name"),
		LegalName:          res.GetStringOpt("legal_name"),
		BillingInformation: billing,
		ExternalRef:        res.GetStringOpt("external_ref"),
		Discoverable:       res.GetBoolOpt("discoverable"),
		Address:            address,
	}

	return req, nil
}

// AccountPatchFromResourceData makes a new Patch payload
func AccountPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.AccountPatch, diag.Diagnostics) {
	res := ResourceData{r}
	req := &ixapi.AccountPatch{}
	if res.HasChange("managing_account") {
		req.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("name") {
		req.Name = res.GetStringOpt("name")
	}
	if res.HasChange("legal_name") {
		req.LegalName = res.GetStringOpt("legal_name")
	}
	if res.HasChange("billing_information") {
		req.BillingInformation = BillingInformationFromResource(
			res.GetResource("billing_information"))
	}
	if res.HasChange("external_ref") {
		req.ExternalRef = res.GetStringOpt("external_ref")
	}
	if res.HasChange("discoverable") {
		req.Discoverable = res.GetBoolOpt("discoverable")
	}
	if res.HasChange("address") {
		req.Address = AddressFromResource(
			res.GetResource("address"))
	}
	return req, nil
}

// AccountSetResourceData sets the resource data for an account
func AccountSetResourceData(acc *ixapi.Account, res ResourceSetter) error {
	if acc.ManagingAccount != nil {
		res.Set("managing_account", *acc.ManagingAccount)
	}
	if err := res.Set("name", acc.Name); err != nil {
		return err
	}

	if acc.LegalName != nil {
		res.Set("legal_name", *acc.LegalName)
	}
	if acc.ExternalRef != nil {
		res.Set("external_ref", *acc.ExternalRef)
	}
	if acc.Discoverable != nil {
		res.Set("discoverable", *acc.Discoverable)
	}
	if acc.MetroAreaNetworkPresence != nil {
		res.Set("metro_area_network_presence", acc.MetroAreaNetworkPresence)
	}

	res.Set("billing_information", FlattenBillingInformation(acc.BillingInformation))
	res.Set("address", FlattenAddress(acc.Address))
	res.Set("state", acc.State)
	res.Set("id", acc.ID)
	return nil
}

// FlattenAccount makes a flat account
func FlattenAccount(acc *ixapi.Account) map[string]interface{} {
	res := NewFlatResource()
	AccountSetResourceData(acc, res)
	return res.Flatten()
}
