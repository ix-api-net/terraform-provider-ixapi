package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func NewCloudRouterPolicyResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_de_cix_cloud_router_policy` resource to manage a BGP routing policy.",
		CreateContext: crud.Create(policyCreate),
		UpdateContext: crud.Update(policyUpdate),
		ReadContext:   crud.Read(policyRead),
		DeleteContext: crud.Delete(policyDelete),

		Schema: schemas.PolicySchema(),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func policyRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.PolicyRequest, error) {
	res := schemas.ResourceDataFrom(r)

	entriesRaw := r.Get("entries").([]interface{})
	entries := make([]ixapi.PolicyEntry, len(entriesRaw))
	for i, item := range entriesRaw {
		entryMap := item.(map[string]interface{})
		entry := ixapi.PolicyEntry{
			SequenceNumber: entryMap["sequence_number"].(int),
		}

		if matchPrefixList, ok := entryMap["match_prefix_list"].(string); ok && matchPrefixList != "" {
			entry.MatchPrefixList = &matchPrefixList
		}

		actionList := entryMap["action"].([]interface{})
		if len(actionList) > 0 {
			actionMap := actionList[0].(map[string]interface{})
			action := ixapi.PolicyAction{}

			if filter, ok := actionMap["filter"].(string); ok && filter != "" {
				action.Filter = &filter
			}
			if localPref, ok := actionMap["local_preference"].(int); ok {
				action.LocalPreference = &localPref
			}
			if asPrependList, ok := actionMap["as_path_prepend"].([]interface{}); ok && len(asPrependList) > 0 {
				asPrependMap := asPrependList[0].(map[string]interface{})
				asPrepend := &ixapi.ASPathPrepend{
					Count: asPrependMap["count"].(int),
				}
				if asn, ok := asPrependMap["asn"].(int); ok && asn > 0 {
					asPrepend.ASN = &asn
				}
				action.ASPathPrepend = asPrepend
			}

			entry.Action = action
		}

		entries[i] = entry
	}

	req := &ixapi.PolicyRequest{
		Name:             res.GetString("name"),
		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		Entries:          entries,
	}
	return req, nil
}

func policyCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	req, err := policyRequestFromResourceData(res)
	if err != nil {
		return err
	}

	policy, err := api.PoliciesCreate(ctx, req)
	if err != nil {
		return err
	}
	res.SetId(policy.ID)

	return policyRead(ctx, res, api)
}

func policyRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	policy, err := api.PoliciesRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("")
		return nil
	}
	if err != nil {
		return err
	}

	if err := schemas.SetResourceData(policy, res); err != nil {
		return err
	}
	return nil
}

func policyUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	req, err := policyRequestFromResourceData(res)
	if err != nil {
		return err
	}

	_, err = api.PoliciesUpdate(ctx, id, req)
	if err != nil {
		return err
	}
	return policyRead(ctx, res, api)
}

func policyDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	_, err := api.PoliciesDelete(ctx, id)
	if err != nil && !ixapi.IsErrNotFound(err) {
		return err
	}
	return policyRead(ctx, res, api)
}
