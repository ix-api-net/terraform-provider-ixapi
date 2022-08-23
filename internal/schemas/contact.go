package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
)

// FlattenContact makes a flat contact
func FlattenContact(
	contact *ixapi.Contact,
) map[string]interface{} {
	res := map[string]interface{}{}
	res["managing_account"] = contact.ManagingAccount
	res["consuming_account"] = contact.ConsumingAccount
	if contact.ExternalRef != nil {
		res["external_ref"] = *contact.ExternalRef
	}
	if contact.Name != nil {
		res["name"] = *contact.Name
	}
	if contact.Telephone != nil {
		res["telephone"] = *contact.Telephone
	}
	if contact.Email != nil {
		res["email"] = *contact.Email
	}
	res["id"] = contact.ID
	return res
}

// ContactRequestFromResourceData makes a new structured request
func ContactRequestFromResourceData(r *schema.ResourceData) *ixapi.ContactRequest {
	res := ResourceData{r}
	req := &ixapi.ContactRequest{
		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		ExternalRef:      res.GetStringOpt("external_ref"),
		Name:             res.GetStringOpt("name"),
		Telephone:        res.GetStringOpt("telephone"),
		Email:            res.GetStringOpt("email"),
	}
	return req
}

// ContactPatchFromResourceData creates a contact update
func ContactPatchFromResourceData(r *schema.ResourceData) *ixapi.ContactPatch {
	res := ResourceData{r}
	patch := &ixapi.ContactPatch{}
	if res.HasChange("managing_account") {
		patch.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("consuming_account") {
		patch.ConsumingAccount = res.GetStringOpt("consuming_account")
	}
	if res.HasChange("name") {
		patch.Name = res.GetStringOpt("name")
	}
	if res.HasChange("telephone") {
		patch.Telephone = res.GetStringOpt("telephone")
	}
	if res.HasChange("email") {
		patch.Email = res.GetStringOpt("email")
	}
	return patch
}

// ContactSetResourceData assigns state
func ContactSetResourceData(contact *ixapi.Contact, res *schema.ResourceData) {
	res.Set("id", contact.ID)
	res.Set("managing_account", contact.ManagingAccount)
	res.Set("consuming_account", contact.ConsumingAccount)
	if contact.ExternalRef != nil {
		res.Set("external_ref", *contact.ExternalRef)
	}
	if contact.Email != nil {
		res.Set("email", *contact.Email)
	}
	if contact.Telephone != nil {
		res.Set("telephone", *contact.Telephone)
	}
	if contact.Name != nil {
		res.Set("name", *contact.Name)
	}
}
