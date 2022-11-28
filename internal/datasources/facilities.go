package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewFacilitiesDataSource creates a data source schema
// for fetching a list of filtered facilities
func NewFacilitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the data source to fetch a list of facilities",
		ReadContext: crud.Read(facilitiesRead),
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
			"name": schemas.DataSourceQuery(
				"Filter by name"),
			"facilities": schemas.IntoDataSourceResultsSchema(
				schemas.FacilitySchema(),
			),
		},
	}
}

// Fetch facilities from API
func fetchFacilities(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) ([]*ixapi.Facility, error) {
	// Filters
	metroArea, hasMetroArea := res.GetOk("metro_area")
	metroAreaNetwork, hasMetroAreaNetwork := res.GetOk("metro_area_network")
	addressCountry, hasAddressCountry := res.GetOk("address_country")
	addressLocality, hasAddressLocality := res.GetOk("address_locality")
	postalCode, hasPostalCode := res.GetOk("postal_code")
	organisationName, hasOrganisationName := res.GetOk("organisation_name")
	pdbFacilityID, hasPdbFacilityID := res.GetOk("peeringdb_facility_id")
	id, hasID := res.GetOk("id")
	name, hasName := res.GetOk("name")

	// Query
	qry := &ixapi.FacilitiesListQuery{}
	if hasID {
		qry.ID = []string{id.(string)}
	}
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

	// Fetch from API
	result, err := api.FacilitiesList(ctx, qry)
	if err != nil {
		return nil, err
	}

	// Local additional filters
	filtered := make([]*ixapi.Facility, 0, len(result))
	for _, fac := range result {
		// Just to be sure, in case IXPs are incomplete in their
		// filter implementation, we filter by ID if present.
		if hasID && fac.ID != id.(string) {
			continue
		}
		if hasPdbFacilityID && fac.PeeringdbFacilityID == nil {
			continue
		}
		if hasPdbFacilityID && *fac.PeeringdbFacilityID != pdbFacilityID.(int) {
			continue
		}
		if hasName && strings.ToLower(fac.Name) != strings.ToLower(name.(string)) {
			continue
		}
		filtered = append(filtered, fac)
	}
	return filtered, nil
}

// Operations
func facilitiesRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	result, err := fetchFacilities(ctx, res, api)
	if err != nil {
		return err
	}

	facilities, err := schemas.FlattenModels(result)
	if err != nil {
		return err
	}

	if err := res.Set("facilities", facilities); err != nil {
		return err
	}
	res.SetId(schemas.Timestamp())
	return nil
}

// NewFacilityDataSource creates a data source schema for
// reading a single facility
func NewFacilityDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get a specific facility",
		ReadContext: crud.Read(facilityRead),
		Schema: schemas.IntoDataSourceSchema(
			schemas.FacilitySchema(),
		),
	}
}

// Fetch a single facility
func facilityRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	// Filters
	results, err := fetchFacilities(ctx, res, api)
	if err != nil {
		return err
	}
	if len(results) == 0 {
		return fmt.Errorf("A facility could not be found")
	}
	if len(results) > 1 {
		return fmt.Errorf("Multiple facilities were returned")
	}

	facility := results[0]
	if err := schemas.SetResourceData(facility, res); err != nil {
		return err
	}
	res.SetId(facility.ID)

	return nil
}
