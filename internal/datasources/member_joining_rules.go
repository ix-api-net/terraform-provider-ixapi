package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/crud"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

func memberJoiningRuleDataSourceSchema() map[string]*schema.Schema {
	return schemas.IntoDataSourceSchema(
		schemas.Combine(
			schemas.AllowMemberJoiningRuleSchema(),
			schemas.DenyMemberJoiningRuleSchema(),
			map[string]*schema.Schema{
				"type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Type of member joining rule: allow or deny",
				},
			}))
}

// NewMemberJoiningRulesDataSource creates a data source for
// querying member joining rules
func NewMemberJoiningRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query member joining rules",
		ReadContext: crud.Read(memberJoiningRulesRead),
		Schema: map[string]*schema.Schema{
			"type": schemas.DataSourceQuery(
				"Filter member joining rules by type: allow or deny"),
			"network_service": schemas.DataSourceQuery(
				"Select member joining rules for this network service"),
			"member_joining_rules": schemas.IntoDataSourceResultsSchema(
				memberJoiningRuleDataSourceSchema()),
		},
	}
}

func memberJoiningRulesQuery(
	res *schema.ResourceData,
) *ixapi.MemberJoiningRulesListQuery {
	qry := &ixapi.MemberJoiningRulesListQuery{}
	ns, hasNS := res.GetOk("network_service")
	if hasNS {
		qry.NetworkService = ns.(string)
	}
	return qry
}

// Fetch member joining rules
func memberJoiningRulesRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	qry := memberJoiningRulesQuery(res)
	results, err := api.MemberJoiningRulesList(ctx, qry)
	if err != nil {
		return err
	}

	ruleType, hasRuleType := res.GetOk("type")

	filtered := make([]ixapi.MemberJoiningRule, 0, len(results))
	for _, rule := range results {
		if hasRuleType && ruleType.(string) != rule.PolymorphicType() {
			continue
		}
		filtered = append(filtered, rule)
	}
	flat, err := schemas.FlattenModels(filtered)
	if err != nil {
		return err
	}

	if err := res.Set("member_joining_rules", flat); err != nil {
		return err
	}
	res.SetId(schemas.Timestamp())
	return nil
}

// NewMemberJoiningRuleDataSource creates a datasource for querying
// a member joining rule
func NewMemberJoiningRuleDataSource() *schema.Resource {
	s := memberJoiningRuleDataSourceSchema()
	s["id"].Optional = false
	s["id"].Computed = false
	s["id"].Required = true
	return &schema.Resource{
		Description: "Use this data source to refernce a joining rule by ID",
		ReadContext: crud.Read(memberJoiningRuleRead),
		Schema:      s,
	}
}

// Fetch a single member joining rule by ID
func memberJoiningRuleRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Get("id").(string)
	rule, err := api.MemberJoiningRulesRead(ctx, id)
	if err != nil {
		return err
	}
	if err := schemas.SetResourceData(rule, res); err != nil {
		return err
	}
	res.SetId(id)
	return nil
}
