package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewFacilitiesDataSource creates a data source schema
// for fetching a list of filtered facilities
func NewFacilitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the data source to fetch a list of facilities",
		ReadContext: facilitiesRead,
		Schema: map[string]*schema.Schema{
			"metro_area": schemas.DataSourceQuery(
				"Filter facilities by metro area ID"),
			"metro_area_network": schemas.DataSourceQuery(
				"Filter by metro are network ID"),
			"address_country": schemas.DataSourceQuery(
				"Filter by country of the facilitie's address"),
			"address_locality": schemas.DataSourceQuery(
				"Filter by locality ('city') of the facilitie's address"),
			"postal_code": schemas.DataSourceQuery(
				"Filter by postal code of the facilitie's address"),
			"organisation_name": schemas.DataSourceQuery(
				"Filter by name of the organisation operating the facility"),
			"peeringdb_facility_id": schemas.DataSourceQueryInt(
				"Filter byu peeringdb id"),
			"facilities": schemas.IntoDataSourceResultsSchema(
				schemas.FacilitySchema(),
			),
		},
	}
}

// Operations

func facilitiesRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Filters
	metroArea, hasMetroArea := res.GetOk("metro_area")
	metroAreaNetwork, hasMetroAreaNetwork := res.GetOk("metro_area_network")
	addressCountry, hasAddressCountry := res.GetOk("address_country")
	addressLocality, hasAddressLocality := res.GetOk("address_locality")
	postalCode, hasPostalCode := res.GetOk("postal_code")
	organisationName, hasOrganisationName := res.GetOk("organisation_name")
	pdbFacilityID, hasPdbFacilityID := res.GetOk("peeringdb_facility_id")

	// Query
	qry := &ixapi.FacilitiesListQuery{}
	if hasMetroArea {
		qry.MetroArea = metroArea.(string)
	}
	if hasMetroAreaNetwork {
		qry.MetroAreaNetwork = metroAreaNetwork.(string)
	}
	if hasAddressCountry {
		qry.AddressCountry = addressCountry.(string)
	}
	if hasAddressLocality {
		qry.AddressLocality = addressLocality.(string)
	}
	if hasPostalCode {
		qry.PostalCode = postalCode.(string)
	}
	if hasOrganisationName {
		qry.OrganisationName = organisationName.(string)
	}

	result, err := api.FacilitiesList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	// Local additional filters
	filtered := make([]*ixapi.Facility, 0, len(result))
	for _, fac := range result {
		if hasPdbFacilityID && fac.PeeringdbFacilityID == nil {
			continue
		}
		if hasPdbFacilityID && *fac.PeeringdbFacilityID != pdbFacilityID.(int) {
			continue
		}
		filtered = append(filtered, fac)
	}

	facilities, err := schemas.FlattenModels(filtered)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("facilities", facilities); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

// NewFacilityDataSource creates a data source schema for
// reading a single facility
func NewFacilityDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get a specific facility",
		ReadContext: facilityRead,
		Schema: schemas.IntoDataSourceSchema(
			schemas.FacilitySchema(),
		),
	}
}

// Fetch a single facility
func facilityRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Filters
	id, hasID := res.GetOk("id")
	name, hasName := res.GetOk("name")
	pdbFacilityID, hasPdbFacilityID := res.GetOk("peeringdb_facility_id")

	// Query
	qry := &ixapi.FacilitiesListQuery{}
	if hasID {
		qry.ID = []string{id.(string)}
	}
	result, err := api.FacilitiesList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	// Local additional filters
	filtered := make([]*ixapi.Facility, 0, len(result))
	for _, fac := range result {
		if hasName && fac.Name != name.(string) {
			continue
		}
		if hasPdbFacilityID && fac.PeeringdbFacilityID == nil {
			continue
		}
		if hasPdbFacilityID && *fac.PeeringdbFacilityID != pdbFacilityID.(int) {
			continue
		}
		filtered = append(filtered, fac)
	}

	if len(filtered) == 0 {
		return diag.Errorf("a facility could not be found")
	}
	if len(filtered) > 1 {
		return diag.Errorf("multiple facilities were returned")
	}

	if err := schemas.SetResourceData(filtered[0], res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(filtered[0].ID)

	return nil
}
