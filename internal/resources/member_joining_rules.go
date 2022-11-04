package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewMemberJoiningRuleAllow creates a member joining rule
// of type 'allow' resource
func NewMemberJoiningRuleAllow() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_member_joining_rule_allow` to create an explicit grant to join a network service for a `consuming_account`.",

		CreateContext: crud.Create(mjrAllowCreate),
		ReadContext:   crud.Read(mjrAllowRead),
		UpdateContext: crud.Update(mjrAllowUpdate),
		DeleteContext: crud.Delete(mjrAllowDelete),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: schemas.AllowMemberJoiningRuleSchema(),
	}
}

func mjrAllowRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.AllowMemberJoiningRuleRequest, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.AllowMemberJoiningRuleRequest{
		Type: ixapi.AllowMemberJoiningRuleType,

		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),

		ExternalRef: res.GetStringOpt("external_ref"),

		CapacityMin: res.GetIntOpt("capacity_min"),
		CapacityMax: res.GetIntOpt("capacity_max"),

		NetworkService: res.GetString("network_service"),
	}
	return req, nil
}

func mjrAllowPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.AllowMemberJoiningRulePatch, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.AllowMemberJoiningRulePatch{
		Type: ixapi.AllowMemberJoiningRuleType,
	}

	if res.HasChange("managing_account") {
		req.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("consuming_account") {
		req.ManagingAccount = res.GetStringOpt("consuming_account")
	}

	if res.HasChange("external_ref") {
		req.ExternalRef = res.GetStringOpt("external_ref")
	}

	if res.HasChange("capacity_min") {
		req.CapacityMin = res.GetIntOpt("capacity_min")
	}
	if res.HasChange("capacity_max") {
		req.CapacityMax = res.GetIntOpt("capacity_max")
	}

	if res.HasChange("network_service") {
		return nil, fmt.Errorf(
			"the network service can not be changed - please create a new joining rule")
	}

	return req, nil
}

// Create
func mjrAllowCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := mjrAllowRequestFromResourceData(res)
	if err != nil {
		return err
	}

	rule, err := api.MemberJoiningRulesCreate(ctx, req)
	if err != nil {
		return err
	}
	allow, ok := rule.(*ixapi.AllowMemberJoiningRule)
	if !ok {
		return ErrUnexpectedPolymorphic(rule, ixapi.AllowMemberJoiningRuleType)
	}
	res.SetId(allow.ID)
	return mjrAllowRead(ctx, res, api)
}

// Read
func mjrAllowRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	rule, err := api.MemberJoiningRulesRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("") // Resource is gone
		return nil
	}
	if err != nil {
		return err
	}

	allow, ok := rule.(*ixapi.AllowMemberJoiningRule)
	if !ok {
		return ErrUnexpectedPolymorphic(rule, ixapi.AllowMemberJoiningRuleType)
	}
	if err := schemas.SetResourceData(allow, res); err != nil {
		return err
	}

	return nil
}

// Update
func mjrAllowUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := mjrAllowPatchFromResourceData(res)
	if err != nil {
		return err
	}

	_, err = api.MemberJoiningRulesPatch(ctx, id, patch)
	if err != nil {
		return err
	}

	return mjrAllowRead(ctx, res, api)
}

// Delete
func mjrAllowDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()

	_, err := api.MemberJoiningRulesDestroy(ctx, id)
	if err != nil {
		return err
	}

	return mjrAllowRead(ctx, res, api)
}

// NewMemberJoiningRuleDeny creates a member joining rule
// of type 'allow' resource
func NewMemberJoiningRuleDeny() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_member_joining_rule_deny` to prevent a `consuming_account` to join a network service.",

		CreateContext: crud.Create(mjrDenyCreate),
		ReadContext:   crud.Read(mjrDenyRead),
		UpdateContext: crud.Update(mjrDenyUpdate),
		DeleteContext: crud.Delete(mjrDenyDelete),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: schemas.DenyMemberJoiningRuleSchema(),
	}
}

func mjrDenyRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.DenyMemberJoiningRuleRequest, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.DenyMemberJoiningRuleRequest{
		Type: ixapi.DenyMemberJoiningRuleType,

		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),

		ExternalRef: res.GetStringOpt("external_ref"),

		NetworkService: res.GetString("network_service"),
	}
	return req, nil
}

func mjrDenyPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.DenyMemberJoiningRulePatch, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.DenyMemberJoiningRulePatch{
		Type: ixapi.DenyMemberJoiningRuleType,
	}

	if res.HasChange("managing_account") {
		req.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("consuming_account") {
		req.ManagingAccount = res.GetStringOpt("consuming_account")
	}

	if res.HasChange("external_ref") {
		req.ExternalRef = res.GetStringOpt("external_ref")
	}

	if res.HasChange("network_service") {
		return nil, fmt.Errorf(
			"the network service can not be changed - please create a new joining rule")
	}

	return req, nil
}

// Create
func mjrDenyCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := mjrDenyRequestFromResourceData(res)
	if err != nil {
		return err
	}

	rule, err := api.MemberJoiningRulesCreate(ctx, req)
	if err != nil {
		return err
	}
	allow, ok := rule.(*ixapi.DenyMemberJoiningRule)
	if !ok {
		return ErrUnexpectedPolymorphic(rule, ixapi.DenyMemberJoiningRuleType)
	}
	res.SetId(allow.ID)
	return mjrDenyRead(ctx, res, api)
}

// Read
func mjrDenyRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	rule, err := api.MemberJoiningRulesRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("") // Resource is gone
		return nil
	}
	if err != nil {
		return err
	}

	allow, ok := rule.(*ixapi.DenyMemberJoiningRule)
	if !ok {
		return ErrUnexpectedPolymorphic(rule, ixapi.DenyMemberJoiningRuleType)
	}
	if err := schemas.SetResourceData(allow, res); err != nil {
		return err
	}

	return nil
}

// Update
func mjrDenyUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := mjrDenyPatchFromResourceData(res)
	if err != nil {
		return err
	}

	_, err = api.MemberJoiningRulesPatch(ctx, id, patch)
	if err != nil {
		return err
	}

	return mjrDenyRead(ctx, res, api)
}

// Delete
func mjrDenyDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()

	_, err := api.MemberJoiningRulesDestroy(ctx, id)
	if err != nil {
		return err
	}

	return mjrDenyRead(ctx, res, api)
}
