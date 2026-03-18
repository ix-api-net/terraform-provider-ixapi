package ixapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// This file is generated from the IX-API schema.
// DO NOT EDIT.

// Errors
var (
	// ErrInvalidPolymorphicType is raised when a polymorphic type could
	// not be resolved.
	ErrInvalidPolymorphicType = errors.New("unknown polymorphic type")
)

// AuthTokenCreate **Operation:** `auth_token_create`
//
//	**Summary:** Authenticate an API user identified by `api_key` and
//
// `api_secret`.
func (c *Client) AuthTokenCreate(
	ctx context.Context,
	req *AuthTokenRequest,

) (*AuthToken, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/auth/token" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &AuthToken{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// AuthTokenRefresh **Operation:** `auth_token_refresh`
//
//	**Summary:** Reauthenticate the API user, issue a new `access_token`
//
// and `refresh_token` pair by providing the `refresh_token`
// in the request body.
func (c *Client) AuthTokenRefresh(
	ctx context.Context,
	req *RefreshTokenRequest,

) (*AuthToken, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/auth/refresh" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &AuthToken{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// FacilitiesListQuery has all query parameters for FacilitiesList
type FacilitiesListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// CapabilitySpeed is a capability_speed
	CapabilitySpeed int `json:"capability_speed,omitempty"`

	// CapabilitySpeedLt is a capability_speed_lt
	CapabilitySpeedLt int `json:"capability_speed_lt,omitempty"`

	// CapabilitySpeedLte is a capability_speed_lte
	CapabilitySpeedLte int `json:"capability_speed_lte,omitempty"`

	// CapabilitySpeedGt is a capability_speed_gt
	CapabilitySpeedGt int `json:"capability_speed_gt,omitempty"`

	// CapabilitySpeedGte is a capability_speed_gte
	CapabilitySpeedGte int `json:"capability_speed_gte,omitempty"`

	// CapabilityMediaType is a capability_media_type
	CapabilityMediaType string `json:"capability_media_type,omitempty"`

	// OrganisationName is a organisation_name
	OrganisationName string `json:"organisation_name,omitempty"`

	// MetroArea is a metro_area
	MetroArea string `json:"metro_area,omitempty"`

	// MetroAreaNetwork is a metro_area_network
	MetroAreaNetwork string `json:"metro_area_network,omitempty"`

	// AddressCountry is a address_country
	AddressCountry string `json:"address_country,omitempty"`

	// AddressLocality is a address_locality
	AddressLocality string `json:"address_locality,omitempty"`

	// PostalCode is a postal_code
	PostalCode string `json:"postal_code,omitempty"`
}

// RawQuery creates a query string for FacilitiesListQuery
func (f *FacilitiesListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(f.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", f.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", f.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = f.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = fmt.Sprintf("%v", f.CapabilitySpeed)
	if val != "0" {
		qry.Add("capability_speed", val)
	}
	val = fmt.Sprintf("%v", f.CapabilitySpeedLt)
	if val != "0" {
		qry.Add("capability_speed_lt", val)
	}
	val = fmt.Sprintf("%v", f.CapabilitySpeedLte)
	if val != "0" {
		qry.Add("capability_speed_lte", val)
	}
	val = fmt.Sprintf("%v", f.CapabilitySpeedGt)
	if val != "0" {
		qry.Add("capability_speed_gt", val)
	}
	val = fmt.Sprintf("%v", f.CapabilitySpeedGte)
	if val != "0" {
		qry.Add("capability_speed_gte", val)
	}
	val = f.CapabilityMediaType
	if val != "" {
		qry.Add("capability_media_type", val)
	}
	val = f.OrganisationName
	if val != "" {
		qry.Add("organisation_name", val)
	}
	val = f.MetroArea
	if val != "" {
		qry.Add("metro_area", val)
	}
	val = f.MetroAreaNetwork
	if val != "" {
		qry.Add("metro_area_network", val)
	}
	val = f.AddressCountry
	if val != "" {
		qry.Add("address_country", val)
	}
	val = f.AddressLocality
	if val != "" {
		qry.Add("address_locality", val)
	}
	val = f.PostalCode
	if val != "" {
		qry.Add("postal_code", val)
	}
	return qry.Encode()
}

// FacilitiesList **Operation:** `facilities_list`
//
//	**Summary:** Get a (filtered) list of `facilities`.
func (c *Client) FacilitiesList(
	ctx context.Context,

	qry ...*FacilitiesListQuery,
) ([]*Facility, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/facilities" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*Facility{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// FacilitiesRead **Operation:** `facilities_read`
//
//	**Summary:** Retrieve a facility by id
func (c *Client) FacilitiesRead(
	ctx context.Context,
	id string,

) (*Facility, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/facilities/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Facility{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// DevicesListQuery has all query parameters for DevicesList
type DevicesListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// CapabilitySpeed is a capability_speed
	CapabilitySpeed int `json:"capability_speed,omitempty"`

	// CapabilitySpeedLt is a capability_speed_lt
	CapabilitySpeedLt int `json:"capability_speed_lt,omitempty"`

	// CapabilitySpeedLte is a capability_speed_lte
	CapabilitySpeedLte int `json:"capability_speed_lte,omitempty"`

	// CapabilitySpeedGt is a capability_speed_gt
	CapabilitySpeedGt int `json:"capability_speed_gt,omitempty"`

	// CapabilitySpeedGte is a capability_speed_gte
	CapabilitySpeedGte int `json:"capability_speed_gte,omitempty"`

	// Name is a name
	Name string `json:"name,omitempty"`

	// CapabilityMediaType is a capability_media_type
	CapabilityMediaType string `json:"capability_media_type,omitempty"`

	// Facility is a facility
	Facility string `json:"facility,omitempty"`

	// Pop is a pop
	Pop string `json:"pop,omitempty"`

	// MetroAreaNetwork is a metro_area_network
	MetroAreaNetwork string `json:"metro_area_network,omitempty"`
}

// RawQuery creates a query string for DevicesListQuery
func (d *DevicesListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(d.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", d.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", d.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = d.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = fmt.Sprintf("%v", d.CapabilitySpeed)
	if val != "0" {
		qry.Add("capability_speed", val)
	}
	val = fmt.Sprintf("%v", d.CapabilitySpeedLt)
	if val != "0" {
		qry.Add("capability_speed_lt", val)
	}
	val = fmt.Sprintf("%v", d.CapabilitySpeedLte)
	if val != "0" {
		qry.Add("capability_speed_lte", val)
	}
	val = fmt.Sprintf("%v", d.CapabilitySpeedGt)
	if val != "0" {
		qry.Add("capability_speed_gt", val)
	}
	val = fmt.Sprintf("%v", d.CapabilitySpeedGte)
	if val != "0" {
		qry.Add("capability_speed_gte", val)
	}
	val = d.Name
	if val != "" {
		qry.Add("name", val)
	}
	val = d.CapabilityMediaType
	if val != "" {
		qry.Add("capability_media_type", val)
	}
	val = d.Facility
	if val != "" {
		qry.Add("facility", val)
	}
	val = d.Pop
	if val != "" {
		qry.Add("pop", val)
	}
	val = d.MetroAreaNetwork
	if val != "" {
		qry.Add("metro_area_network", val)
	}
	return qry.Encode()
}

// DevicesList **Operation:** `devices_list`
//
//	**Summary:** List available devices
func (c *Client) DevicesList(
	ctx context.Context,

	qry ...*DevicesListQuery,
) ([]*Device, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/devices" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*Device{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// DevicesRead **Operation:** `devices_read`
//
//	**Summary:** Get a specific device identified by id
func (c *Client) DevicesRead(
	ctx context.Context,
	id string,

) (*Device, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/devices/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Device{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PopsListQuery has all query parameters for PopsList
type PopsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// Facility is a facility
	Facility string `json:"facility,omitempty"`

	// MetroAreaNetwork is a metro_area_network
	MetroAreaNetwork string `json:"metro_area_network,omitempty"`

	// CapabilityMediaType is a capability_media_type
	CapabilityMediaType string `json:"capability_media_type,omitempty"`

	// CapabilitySpeed is a capability_speed
	CapabilitySpeed int `json:"capability_speed,omitempty"`

	// CapabilitySpeedLt is a capability_speed_lt
	CapabilitySpeedLt int `json:"capability_speed_lt,omitempty"`

	// CapabilitySpeedLte is a capability_speed_lte
	CapabilitySpeedLte int `json:"capability_speed_lte,omitempty"`

	// CapabilitySpeedGt is a capability_speed_gt
	CapabilitySpeedGt int `json:"capability_speed_gt,omitempty"`

	// CapabilitySpeedGte is a capability_speed_gte
	CapabilitySpeedGte int `json:"capability_speed_gte,omitempty"`
}

// RawQuery creates a query string for PopsListQuery
func (p *PopsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(p.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", p.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", p.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = p.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = p.Facility
	if val != "" {
		qry.Add("facility", val)
	}
	val = p.MetroAreaNetwork
	if val != "" {
		qry.Add("metro_area_network", val)
	}
	val = p.CapabilityMediaType
	if val != "" {
		qry.Add("capability_media_type", val)
	}
	val = fmt.Sprintf("%v", p.CapabilitySpeed)
	if val != "0" {
		qry.Add("capability_speed", val)
	}
	val = fmt.Sprintf("%v", p.CapabilitySpeedLt)
	if val != "0" {
		qry.Add("capability_speed_lt", val)
	}
	val = fmt.Sprintf("%v", p.CapabilitySpeedLte)
	if val != "0" {
		qry.Add("capability_speed_lte", val)
	}
	val = fmt.Sprintf("%v", p.CapabilitySpeedGt)
	if val != "0" {
		qry.Add("capability_speed_gt", val)
	}
	val = fmt.Sprintf("%v", p.CapabilitySpeedGte)
	if val != "0" {
		qry.Add("capability_speed_gte", val)
	}
	return qry.Encode()
}

// PopsList **Operation:** `pops_list`
//
//	**Summary:** List all PoPs
func (c *Client) PopsList(
	ctx context.Context,

	qry ...*PopsListQuery,
) ([]*PointOfPresence, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/pops" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*PointOfPresence{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PopsRead **Operation:** `pops_read`
//
//	**Summary:** Get a single point of presence
func (c *Client) PopsRead(
	ctx context.Context,
	id string,

) (*PointOfPresence, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/pops/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &PointOfPresence{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MetroAreaNetworksListQuery has all query parameters for MetroAreaNetworksList
type MetroAreaNetworksListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// Name is a name
	Name string `json:"name,omitempty"`

	// MetroArea is a metro_area
	MetroArea string `json:"metro_area,omitempty"`

	// ServiceProvider is a service_provider
	ServiceProvider string `json:"service_provider,omitempty"`
}

// RawQuery creates a query string for MetroAreaNetworksListQuery
func (m *MetroAreaNetworksListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(m.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", m.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", m.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = m.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = m.Name
	if val != "" {
		qry.Add("name", val)
	}
	val = m.MetroArea
	if val != "" {
		qry.Add("metro_area", val)
	}
	val = m.ServiceProvider
	if val != "" {
		qry.Add("service_provider", val)
	}
	return qry.Encode()
}

// MetroAreaNetworksList **Operation:** `metro_area_networks_list`
//
//	**Summary:** List all MetroAreaNetworks
func (c *Client) MetroAreaNetworksList(
	ctx context.Context,

	qry ...*MetroAreaNetworksListQuery,
) ([]*MetroAreaNetwork, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/metro-area-networks" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*MetroAreaNetwork{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MetroAreaNetworksRead **Operation:** `metro_area_networks_read`
//
//	**Summary:** Retrieve a MetroAreaNetwork
func (c *Client) MetroAreaNetworksRead(
	ctx context.Context,
	id string,

) (*MetroAreaNetwork, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/metro-area-networks/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &MetroAreaNetwork{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MetroAreasListQuery has all query parameters for MetroAreasList
type MetroAreasListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`
}

// RawQuery creates a query string for MetroAreasListQuery
func (m *MetroAreasListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(m.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", m.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", m.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = m.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	return qry.Encode()
}

// MetroAreasList **Operation:** `metro_areas_list`
//
//	**Summary:** List all MetroAreas
func (c *Client) MetroAreasList(
	ctx context.Context,

	qry ...*MetroAreasListQuery,
) ([]*MetroArea, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/metro-areas" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*MetroArea{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MetroAreasRead **Operation:** `metro_areas_read`
//
//	**Summary:** Get a single MetroArea
func (c *Client) MetroAreasRead(
	ctx context.Context,
	id string,

) (*MetroArea, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/metro-areas/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &MetroArea{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ProductOfferingsListQuery has all query parameters for ProductOfferingsList
type ProductOfferingsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// Type is a type
	Type string `json:"type"`

	// ResourceType is a resource_type
	ResourceType string `json:"resource_type,omitempty"`

	// Name is a name
	Name string `json:"name,omitempty"`

	// HandoverMetroArea is a handover_metro_area
	HandoverMetroArea string `json:"handover_metro_area,omitempty"`

	// HandoverMetroAreaNetwork is a handover_metro_area_network
	HandoverMetroAreaNetwork string `json:"handover_metro_area_network,omitempty"`

	// HandoverPop is a handover_pop
	HandoverPop string `json:"handover_pop,omitempty"`

	// ServiceMetroArea is a service_metro_area
	ServiceMetroArea string `json:"service_metro_area,omitempty"`

	// ServiceMetroAreaNetwork is a service_metro_area_network
	ServiceMetroAreaNetwork string `json:"service_metro_area_network,omitempty"`

	// ServiceProvider is a service_provider
	ServiceProvider string `json:"service_provider,omitempty"`

	// DowngradeAllowed is a downgrade_allowed
	DowngradeAllowed string `json:"downgrade_allowed,omitempty"`

	// UpgradeAllowed is a upgrade_allowed
	UpgradeAllowed string `json:"upgrade_allowed,omitempty"`

	// Bandwidth is a bandwidth
	Bandwidth int `json:"bandwidth,omitempty"`

	// PhysicalPortSpeed is a physical_port_speed
	PhysicalPortSpeed int `json:"physical_port_speed,omitempty"`

	// ServiceProviderRegion is a service_provider_region
	ServiceProviderRegion string `json:"service_provider_region,omitempty"`

	// ServiceProviderPop is a service_provider_pop
	ServiceProviderPop string `json:"service_provider_pop,omitempty"`

	// DeliveryMethod is a delivery_method
	DeliveryMethod string `json:"delivery_method,omitempty"`

	// CloudKey is a cloud_key
	CloudKey string `json:"cloud_key,omitempty"`

	// ContractInitialPeriod is a contract_initial_period
	ContractInitialPeriod string `json:"contract_initial_period,omitempty"`

	// ContractInitialNoticePeriod is a contract_initial_notice_period
	ContractInitialNoticePeriod string `json:"contract_initial_notice_period,omitempty"`

	// ContractRenewalPeriod is a contract_renewal_period
	ContractRenewalPeriod string `json:"contract_renewal_period,omitempty"`

	// ContractRenewalNoticePeriod is a contract_renewal_notice_period
	ContractRenewalNoticePeriod string `json:"contract_renewal_notice_period,omitempty"`

	// Fields is a fields
	Fields string `json:"fields,omitempty"`
}

// RawQuery creates a query string for ProductOfferingsListQuery
func (p *ProductOfferingsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(p.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", p.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", p.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = p.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = p.Type
	if val != "" {
		qry.Add("type", val)
	}
	val = p.ResourceType
	if val != "" {
		qry.Add("resource_type", val)
	}
	val = p.Name
	if val != "" {
		qry.Add("name", val)
	}
	val = p.HandoverMetroArea
	if val != "" {
		qry.Add("handover_metro_area", val)
	}
	val = p.HandoverMetroAreaNetwork
	if val != "" {
		qry.Add("handover_metro_area_network", val)
	}
	val = p.HandoverPop
	if val != "" {
		qry.Add("handover_pop", val)
	}
	val = p.ServiceMetroArea
	if val != "" {
		qry.Add("service_metro_area", val)
	}
	val = p.ServiceMetroAreaNetwork
	if val != "" {
		qry.Add("service_metro_area_network", val)
	}
	val = p.ServiceProvider
	if val != "" {
		qry.Add("service_provider", val)
	}
	val = p.DowngradeAllowed
	if val != "" {
		qry.Add("downgrade_allowed", val)
	}
	val = p.UpgradeAllowed
	if val != "" {
		qry.Add("upgrade_allowed", val)
	}
	val = fmt.Sprintf("%v", p.Bandwidth)
	if val != "0" {
		qry.Add("bandwidth", val)
	}
	val = fmt.Sprintf("%v", p.PhysicalPortSpeed)
	if val != "0" {
		qry.Add("physical_port_speed", val)
	}
	val = p.ServiceProviderRegion
	if val != "" {
		qry.Add("service_provider_region", val)
	}
	val = p.ServiceProviderPop
	if val != "" {
		qry.Add("service_provider_pop", val)
	}
	val = p.DeliveryMethod
	if val != "" {
		qry.Add("delivery_method", val)
	}
	val = p.CloudKey
	if val != "" {
		qry.Add("cloud_key", val)
	}
	val = p.ContractInitialPeriod
	if val != "" {
		qry.Add("contract_initial_period", val)
	}
	val = p.ContractInitialNoticePeriod
	if val != "" {
		qry.Add("contract_initial_notice_period", val)
	}
	val = p.ContractRenewalPeriod
	if val != "" {
		qry.Add("contract_renewal_period", val)
	}
	val = p.ContractRenewalNoticePeriod
	if val != "" {
		qry.Add("contract_renewal_notice_period", val)
	}
	val = p.Fields
	if val != "" {
		qry.Add("fields", val)
	}
	return qry.Encode()
}

// ProductOfferingsList **Operation:** `product_offerings_list`
//
//	**Summary:** List all (filtered) products-offerings available on the platform
func (c *Client) ProductOfferingsList(
	ctx context.Context,

	qry ...*ProductOfferingsListQuery,
) ([]ProductOffering, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/product-offerings" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		msgs := []json.RawMessage{}
		if err := json.Unmarshal(body, &msgs); err != nil {
			return nil, err
		}
		res := make([]ProductOffering, 0, len(msgs))

		for _, msg := range msgs {
			tmp := &PolymorphicProductOffering{}
			if err := json.Unmarshal(msg, tmp); err != nil {
				return nil, err
			}
			ptype := tmp.PolymorphicType()
			switch ptype {

			case ConnectionProductOfferingType:
				rval := &ConnectionProductOffering{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case ExchangeLanNetworkProductOfferingType:
				rval := &ExchangeLanNetworkProductOffering{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case P2PNetworkProductOfferingType:
				rval := &P2PNetworkProductOffering{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case MP2MPNetworkProductOfferingType:
				rval := &MP2MPNetworkProductOffering{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case P2MPNetworkProductOfferingType:
				rval := &P2MPNetworkProductOffering{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case CloudNetworkProductOfferingType:
				rval := &CloudNetworkProductOffering{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case RoutingFunctionProductOfferingType:
				rval := &RoutingFunctionProductOffering{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			}
		}

		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ProductOfferingsRead **Operation:** `product_offerings_read`
//
//	**Summary:** Get a single products-offering by id.
func (c *Client) ProductOfferingsRead(
	ctx context.Context,
	id string,

) (ProductOffering, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/product-offerings/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicProductOffering{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ConnectionProductOfferingType:
			res := &ConnectionProductOffering{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case ExchangeLanNetworkProductOfferingType:
			res := &ExchangeLanNetworkProductOffering{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2PNetworkProductOfferingType:
			res := &P2PNetworkProductOffering{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case MP2MPNetworkProductOfferingType:
			res := &MP2MPNetworkProductOffering{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2MPNetworkProductOfferingType:
			res := &P2MPNetworkProductOffering{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case CloudNetworkProductOfferingType:
			res := &CloudNetworkProductOffering{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case RoutingFunctionProductOfferingType:
			res := &RoutingFunctionProductOffering{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// AvailabilityZonesListQuery has all query parameters for AvailabilityZonesList
type AvailabilityZonesListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// Name is a name
	Name string `json:"name,omitempty"`
}

// RawQuery creates a query string for AvailabilityZonesListQuery
func (a *AvailabilityZonesListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(a.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", a.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", a.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = a.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = a.Name
	if val != "" {
		qry.Add("name", val)
	}
	return qry.Encode()
}

// AvailabilityZonesList **Operation:** `availability_zones_list`
//
//	**Summary:** List all availability zones available on the platform.
func (c *Client) AvailabilityZonesList(
	ctx context.Context,

	qry ...*AvailabilityZonesListQuery,
) ([]*AvailabilityZone, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/availability-zones" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*AvailabilityZone{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// AvailabilityZonesRead **Operation:** `availability_zones_read`
//
//	**Summary:** Get a single availability zone by id.
func (c *Client) AvailabilityZonesRead(
	ctx context.Context,
	id string,

) (*AvailabilityZone, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/availability-zones/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &AvailabilityZone{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortsListQuery has all query parameters for PortsList
type PortsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// StateIsNot is a state__is_not
	StateIsNot string `json:"state__is_not,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount is a consuming_account
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`

	// MediaType is a media_type
	MediaType string `json:"media_type,omitempty"`

	// Pop is a pop
	Pop string `json:"pop,omitempty"`

	// Name is a name
	Name string `json:"name,omitempty"`

	// Device is a device
	Device string `json:"device,omitempty"`

	// Speed is a speed
	Speed string `json:"speed,omitempty"`

	// Connection is a connection
	Connection string `json:"connection,omitempty"`
}

// RawQuery creates a query string for PortsListQuery
func (p *PortsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(p.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", p.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", p.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = p.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = p.State
	if val != "" {
		qry.Add("state", val)
	}
	val = p.StateIsNot
	if val != "" {
		qry.Add("state__is_not", val)
	}
	val = p.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = p.ConsumingAccount
	if val != "" {
		qry.Add("consuming_account", val)
	}
	val = p.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	val = p.MediaType
	if val != "" {
		qry.Add("media_type", val)
	}
	val = p.Pop
	if val != "" {
		qry.Add("pop", val)
	}
	val = p.Name
	if val != "" {
		qry.Add("name", val)
	}
	val = p.Device
	if val != "" {
		qry.Add("device", val)
	}
	val = p.Speed
	if val != "" {
		qry.Add("speed", val)
	}
	val = p.Connection
	if val != "" {
		qry.Add("connection", val)
	}
	return qry.Encode()
}

// PortsList **Operation:** `ports_list`
//
//	**Summary:** List all ports.
func (c *Client) PortsList(
	ctx context.Context,

	qry ...*PortsListQuery,
) ([]*Port, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/ports" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*Port{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortsRead **Operation:** `ports_read`
//
//	**Summary:** Retrieve a port.
func (c *Client) PortsRead(
	ctx context.Context,
	id string,

) (*Port, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/ports/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Port{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortReservationsListQuery has all query parameters for PortReservationsList
type PortReservationsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// StateIsNot is a state__is_not
	StateIsNot string `json:"state__is_not,omitempty"`

	// Connection is a connection
	Connection string `json:"connection,omitempty"`

	// Port is a port
	Port string `json:"port,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`
}

// RawQuery creates a query string for PortReservationsListQuery
func (p *PortReservationsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(p.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", p.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", p.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = p.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = p.State
	if val != "" {
		qry.Add("state", val)
	}
	val = p.StateIsNot
	if val != "" {
		qry.Add("state__is_not", val)
	}
	val = p.Connection
	if val != "" {
		qry.Add("connection", val)
	}
	val = p.Port
	if val != "" {
		qry.Add("port", val)
	}
	val = p.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	return qry.Encode()
}

// PortReservationsList **Operation:** `port_reservations_list`
//
//	**Summary:** List all port reservations.
func (c *Client) PortReservationsList(
	ctx context.Context,

	qry ...*PortReservationsListQuery,
) ([]*PortReservation, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/port-reservations" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*PortReservation{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortReservationsCreate **Operation:** `port_reservations_create`
//
//	**Summary:** Create a new `PortReservation`.
//
// Two workflows for allocating ports is supported and
// dependent on the `cross_connect_initiator` property
// of the corresponding `product-offering`:
//
// Individual LOAs can be uploaded and downloaded for
// each PortAllocation using the endpoint
// `/port-reservations/{id}/loa`.
//
// Please refer to the internet exchange's api usage
// guide for implementation specific details.
func (c *Client) PortReservationsCreate(
	ctx context.Context,
	req *PortReservationRequest,

) (*PortReservation, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/port-reservations" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &PortReservation{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortReservationsRead **Operation:** `port_reservations_read`
//
//	**Summary:** Retrieve a `PortReservation`.
func (c *Client) PortReservationsRead(
	ctx context.Context,
	id string,

) (*PortReservation, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/port-reservations/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &PortReservation{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortReservationsUpdate **Operation:** `port_reservations_update`
//
//	**Summary:** **DEPRECATION NOTICE**: This operation will be removed in favor
//
// of using `PATCH` for all updates.
//
// Update a port reservation.
func (c *Client) PortReservationsUpdate(
	ctx context.Context,
	id string, req *PortReservationUpdate,

) (*PortReservation, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/port-reservations/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &PortReservation{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortReservationsPatch **Operation:** `port_reservations_partial_update`
//
//	**Summary:** Update a port reservation.
func (c *Client) PortReservationsPatch(
	ctx context.Context,
	id string, req *PortReservationPatch,

) (*PortReservation, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/port-reservations/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &PortReservation{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortReservationsDestroy **Operation:** `port_reservations_destroy`
//
//	**Summary:** Request decommissioning the port-reservation.
//
// The associated `port` will be deallocated and
// removed from the `connection`.
func (c *Client) PortReservationsDestroy(
	ctx context.Context,
	id string, req *CancellationRequest,

) (*PortReservation, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/port-reservations/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &PortReservation{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortReservationCancellationPolicyReadQuery has all query parameters for PortReservationCancellationPolicyRead
type PortReservationCancellationPolicyReadQuery struct {
	// DecommissionAt is a decommission_at
	DecommissionAt string `json:"decommission_at,omitempty"`
}

// RawQuery creates a query string for PortReservationCancellationPolicyReadQuery
func (p *PortReservationCancellationPolicyReadQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = p.DecommissionAt
	if val != "" {
		qry.Add("decommission_at", val)
	}
	return qry.Encode()
}

// PortReservationCancellationPolicyRead **Operation:** `port_reservation_cancellation_policy_read`
//
//	**Summary:** The cancellation-policy can be queried to answer
//
// the questions:
//
// If I cancel my subscription, *when will it be technically
// decommissioned*?
// If I cancel my subscription, *until what date will I be charged*?
//
// When the query parameter `decommision_at` is not provided
// it will provide the first possible cancellation date
// and charge period if cancelled at above date.
//
// The granularity of the date field is a day, the start and end
// of which are to be interpreted by the IXP (some may use UTC,
// some may use their local time zone).
func (c *Client) PortReservationCancellationPolicyRead(
	ctx context.Context,
	id string,
	qry ...*PortReservationCancellationPolicyReadQuery,
) (*CancellationPolicy, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/port-reservations/{id}/cancellation-policy"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &CancellationPolicy{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortReservationsLoaDownload **Operation:** `port_reservations_loa_download`
//
//	**Summary:** Download the *Letter Of Authorization* associated
//
// with the port-reservation.
//
// In case of a *subscriber initiated cross-connect*,
// it will be provided by the exchange.
func (c *Client) PortReservationsLoaDownload(
	ctx context.Context,
	id string,

) (Response, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/port-reservations/{id}/loa"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {
		return body, nil
	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PortReservationsLoaUpload **Operation:** `port_reservations_loa_upload`
//
//	**Summary:** Upload a *Letter Of Authorization* for this
//
// `PortReservation`.
func (c *Client) PortReservationsLoaUpload(
	ctx context.Context,
	id string, data []byte,

) (Response, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/port-reservations/{id}/loa"+params, id)
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/octet-stream")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {
		return body, nil
	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ConnectionsListQuery has all query parameters for ConnectionsList
type ConnectionsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// StateIsNot is a state__is_not
	StateIsNot string `json:"state__is_not,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount is a consuming_account
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`

	// Mode is a mode
	Mode string `json:"mode,omitempty"`

	// ModeIsNot is a mode__is_not
	ModeIsNot string `json:"mode__is_not,omitempty"`

	// Name is a name
	Name string `json:"name,omitempty"`

	// MetroArea is a metro_area
	MetroArea string `json:"metro_area,omitempty"`

	// MetroAreaNetwork is a metro_area_network
	MetroAreaNetwork string `json:"metro_area_network,omitempty"`

	// Pop is a pop
	Pop string `json:"pop,omitempty"`

	// Facility is a facility
	Facility string `json:"facility,omitempty"`

	// RoleAssignments is a role_assignments
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Contacts is a contacts
	Contacts []string `json:"contacts,omitempty"`

	// SupportedNetworkService is a supported_network_service
	SupportedNetworkService string `json:"supported_network_service,omitempty"`
}

// RawQuery creates a query string for ConnectionsListQuery
func (c *ConnectionsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(c.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", c.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", c.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = c.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = c.State
	if val != "" {
		qry.Add("state", val)
	}
	val = c.StateIsNot
	if val != "" {
		qry.Add("state__is_not", val)
	}
	val = c.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = c.ConsumingAccount
	if val != "" {
		qry.Add("consuming_account", val)
	}
	val = c.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	val = c.Mode
	if val != "" {
		qry.Add("mode", val)
	}
	val = c.ModeIsNot
	if val != "" {
		qry.Add("mode__is_not", val)
	}
	val = c.Name
	if val != "" {
		qry.Add("name", val)
	}
	val = c.MetroArea
	if val != "" {
		qry.Add("metro_area", val)
	}
	val = c.MetroAreaNetwork
	if val != "" {
		qry.Add("metro_area_network", val)
	}
	val = c.Pop
	if val != "" {
		qry.Add("pop", val)
	}
	val = c.Facility
	if val != "" {
		qry.Add("facility", val)
	}
	val = strings.Join(c.RoleAssignments, ",")
	if val != "" {
		qry.Add("role_assignments", val)
	}
	val = strings.Join(c.Contacts, ",")
	if val != "" {
		qry.Add("contacts", val)
	}
	val = c.SupportedNetworkService
	if val != "" {
		qry.Add("supported_network_service", val)
	}
	return qry.Encode()
}

// ConnectionsList **Operation:** `connections_list`
//
//	**Summary:** List all `connection`s.
func (c *Client) ConnectionsList(
	ctx context.Context,

	qry ...*ConnectionsListQuery,
) ([]*Connection, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/connections" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*Connection{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ConnectionsCreate **Operation:** `connections_create`
//
//	**Summary:** Create a new `connection` and request ports
//
// allocation.
//
// Two workflows for allocating ports is supported and
// dependent on the `cross_connect_initiator` property
// of the corresponding `product-offering`:
//
// When the initiator is the `subscriber`, a Letter Of
// Authorization (LOA) can
// be downloaded from the `/connection/<id>/loa`
// resource. In case the `exchange` is the initiator,
// the LOA can be uploaded to this resource.
//
// Creating a connection will also create
// PortReservations. See the `port_quantity` and
// `subscriber_side_demarcs` attributes for details.
//
// Please refer to the internet exchange's api usage
// guide for implementation specific details.
func (c *Client) ConnectionsCreate(
	ctx context.Context,
	req *ConnectionRequest,

) (*Connection, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/connections" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Connection{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ConnectionsRead **Operation:** `connections_read`
//
//	**Summary:** Read a `connection`.
func (c *Client) ConnectionsRead(
	ctx context.Context,
	id string,

) (*Connection, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/connections/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Connection{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ConnectionsUpdate **Operation:** `connections_update`
//
//	**Summary:** **DEPRECATION NOTICE**: This operation will be removed in favor
//
// of using `PATCH` for all updates.
//
// Update a connection.
func (c *Client) ConnectionsUpdate(
	ctx context.Context,
	id string, req *ConnectionUpdate,

) (*Connection, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/connections/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Connection{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ConnectionsPatch **Operation:** `connections_partial_update`
//
//	**Summary:** Update a connection.
func (c *Client) ConnectionsPatch(
	ctx context.Context,
	id string, req *ConnectionPatch,

) (*Connection, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/connections/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Connection{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ConnectionsDestroy **Operation:** `connections_destroy`
//
//	**Summary:** Request decommissioning the connection.
//
// The cancellation policy of the connection applies
// here and is independent from the
// policy of the network-service and network-service-config
// using the connection.
//
// The connection will assume the state
// `decommission_requested`.
//
// Associated `port-reservation` will be also
// marked for decommissining and ports will
// be deallocated.
//
// The decommissioning request will *not* cascade
// to network services and configs.
func (c *Client) ConnectionsDestroy(
	ctx context.Context,
	id string, req *CancellationRequest,

) (*Connection, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/connections/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Connection{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ConnectionsLoaDownload **Operation:** `connections_loa_download`
//
//	**Summary:** Download the *Letter Of Authorization* associated with the `connection`.
//
// In case of a *subscriber initiated cross-connect*,
// it will be provided by the exchange.
func (c *Client) ConnectionsLoaDownload(
	ctx context.Context,
	id string,

) (Response, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/connections/{id}/loa"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {
		return body, nil
	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ConnectionsLoaUpload **Operation:** `connections_loa_upload`
//
//	**Summary:** Upload a *Letter Of Authorization* for this
//
// `connection`.
//
// The LOA is valid for the entire connection and must
// include all ports.
func (c *Client) ConnectionsLoaUpload(
	ctx context.Context,
	id string, data []byte,

) (Response, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/connections/{id}/loa"+params, id)
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/octet-stream")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {
		return body, nil
	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ConnectionsCancellationPolicyReadQuery has all query parameters for ConnectionsCancellationPolicyRead
type ConnectionsCancellationPolicyReadQuery struct {
	// DecommissionAt is a decommission_at
	DecommissionAt string `json:"decommission_at,omitempty"`
}

// RawQuery creates a query string for ConnectionsCancellationPolicyReadQuery
func (c *ConnectionsCancellationPolicyReadQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = c.DecommissionAt
	if val != "" {
		qry.Add("decommission_at", val)
	}
	return qry.Encode()
}

// ConnectionsCancellationPolicyRead **Operation:** `connections_cancellation_policy_read`
//
//	**Summary:** The cancellation-policy can be queried to answer
//
// the questions:
//
// If I cancel my subscription, *when will it be technically
// decommissioned*?
// If I cancel my subscription, *until what date will I be charged*?
//
// When the query parameter `decommision_at` is not provided
// it will provide the first possible cancellation date
// and charge period if cancelled at above date.
//
// The granularity of the date field is a day, the start and end
// of which are to be interpreted by the IXP (some may use UTC,
// some may use their local time zone).
func (c *Client) ConnectionsCancellationPolicyRead(
	ctx context.Context,
	id string,
	qry ...*ConnectionsCancellationPolicyReadQuery,
) (*CancellationPolicy, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/connections/{id}/cancellation-policy"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &CancellationPolicy{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceConfigsListQuery has all query parameters for NetworkServiceConfigsList
type NetworkServiceConfigsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// StateIsNot is a state__is_not
	StateIsNot string `json:"state__is_not,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount is a consuming_account
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`

	// Type is a type
	Type string `json:"type"`

	// InnerVLAN is a inner_vlan
	InnerVLAN int `json:"inner_vlan,omitempty"`

	// OuterVLAN is a outer_vlan
	OuterVLAN int `json:"outer_vlan,omitempty"`

	// Capacity is a capacity
	Capacity int `json:"capacity,omitempty"`

	// NetworkService is a network_service
	NetworkService string `json:"network_service,omitempty"`

	// Connection is a connection
	Connection string `json:"connection,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// RoleAssignments is a role_assignments
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Contacts is a contacts
	Contacts []string `json:"contacts,omitempty"`
}

// RawQuery creates a query string for NetworkServiceConfigsListQuery
func (n *NetworkServiceConfigsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(n.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", n.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", n.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = n.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = n.State
	if val != "" {
		qry.Add("state", val)
	}
	val = n.StateIsNot
	if val != "" {
		qry.Add("state__is_not", val)
	}
	val = n.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = n.ConsumingAccount
	if val != "" {
		qry.Add("consuming_account", val)
	}
	val = n.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	val = n.Type
	if val != "" {
		qry.Add("type", val)
	}
	val = fmt.Sprintf("%v", n.InnerVLAN)
	if val != "0" {
		qry.Add("inner_vlan", val)
	}
	val = fmt.Sprintf("%v", n.OuterVLAN)
	if val != "0" {
		qry.Add("outer_vlan", val)
	}
	val = fmt.Sprintf("%v", n.Capacity)
	if val != "0" {
		qry.Add("capacity", val)
	}
	val = n.NetworkService
	if val != "" {
		qry.Add("network_service", val)
	}
	val = n.Connection
	if val != "" {
		qry.Add("connection", val)
	}
	val = n.ProductOffering
	if val != "" {
		qry.Add("product_offering", val)
	}
	val = strings.Join(n.RoleAssignments, ",")
	if val != "" {
		qry.Add("role_assignments", val)
	}
	val = strings.Join(n.Contacts, ",")
	if val != "" {
		qry.Add("contacts", val)
	}
	return qry.Encode()
}

// NetworkServiceConfigsList **Operation:** `network_service_configs_list`
//
//	**Summary:** Get all `network-service-config`s.
func (c *Client) NetworkServiceConfigsList(
	ctx context.Context,

	qry ...*NetworkServiceConfigsListQuery,
) ([]NetworkServiceConfig, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-service-configs" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		msgs := []json.RawMessage{}
		if err := json.Unmarshal(body, &msgs); err != nil {
			return nil, err
		}
		res := make([]NetworkServiceConfig, 0, len(msgs))

		for _, msg := range msgs {
			tmp := &PolymorphicNetworkServiceConfig{}
			if err := json.Unmarshal(msg, tmp); err != nil {
				return nil, err
			}
			ptype := tmp.PolymorphicType()
			switch ptype {

			case ExchangeLanNetworkServiceConfigType:
				rval := &ExchangeLanNetworkServiceConfig{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				vlanConfig, err := decodeVLANConfig(rval.VLANConfigRaw)
				if err != nil {
					return nil, err
				}
				rval.VLANConfig = vlanConfig

				res = append(res, rval)

			case P2PNetworkServiceConfigType:
				rval := &P2PNetworkServiceConfig{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				vlanConfig, err := decodeVLANConfig(rval.VLANConfigRaw)
				if err != nil {
					return nil, err
				}
				rval.VLANConfig = vlanConfig

				res = append(res, rval)

			case P2MPNetworkServiceConfigType:
				rval := &P2MPNetworkServiceConfig{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				vlanConfig, err := decodeVLANConfig(rval.VLANConfigRaw)
				if err != nil {
					return nil, err
				}
				rval.VLANConfig = vlanConfig

				res = append(res, rval)

			case MP2MPNetworkServiceConfigType:
				rval := &MP2MPNetworkServiceConfig{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				vlanConfig, err := decodeVLANConfig(rval.VLANConfigRaw)
				if err != nil {
					return nil, err
				}
				rval.VLANConfig = vlanConfig

				res = append(res, rval)

			case CloudNetworkServiceConfigType:
				rval := &CloudNetworkServiceConfig{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				vlanConfig, err := decodeVLANConfig(rval.VLANConfigRaw)
				if err != nil {
					return nil, err
				}
				rval.VLANConfig = vlanConfig

				res = append(res, rval)

			}
		}

		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceConfigsCreate **Operation:** `network_service_configs_create`
//
//	**Summary:** Create a `network-service-config`.
func (c *Client) NetworkServiceConfigsCreate(
	ctx context.Context,
	req NetworkServiceConfigRequest,

) (NetworkServiceConfig, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-service-configs" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkServiceConfig{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ExchangeLanNetworkServiceConfigType:
			res := &ExchangeLanNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case P2PNetworkServiceConfigType:
			res := &P2PNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case P2MPNetworkServiceConfigType:
			res := &P2MPNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case MP2MPNetworkServiceConfigType:
			res := &MP2MPNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case CloudNetworkServiceConfigType:
			res := &CloudNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceConfigsRead **Operation:** `network_service_configs_read`
//
//	**Summary:** Get a `network-service-config`
func (c *Client) NetworkServiceConfigsRead(
	ctx context.Context,
	id string,

) (NetworkServiceConfig, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-service-configs/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkServiceConfig{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ExchangeLanNetworkServiceConfigType:
			res := &ExchangeLanNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case P2PNetworkServiceConfigType:
			res := &P2PNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case P2MPNetworkServiceConfigType:
			res := &P2MPNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case MP2MPNetworkServiceConfigType:
			res := &MP2MPNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case CloudNetworkServiceConfigType:
			res := &CloudNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceConfigsUpdate **Operation:** `network_service_configs_update`
//
//	**Summary:** **DEPRECATION NOTICE**: This operation will be removed in favor
//
// of using `PATCH` for all updates.
//
// Update an exisiting `network-service-config`
func (c *Client) NetworkServiceConfigsUpdate(
	ctx context.Context,
	id string, req NetworkServiceConfigUpdate,

) (NetworkServiceConfig, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-service-configs/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkServiceConfig{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ExchangeLanNetworkServiceConfigType:
			res := &ExchangeLanNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case P2PNetworkServiceConfigType:
			res := &P2PNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case P2MPNetworkServiceConfigType:
			res := &P2MPNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case MP2MPNetworkServiceConfigType:
			res := &MP2MPNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case CloudNetworkServiceConfigType:
			res := &CloudNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceConfigsPatch **Operation:** `network_service_configs_partial_update`
//
//	**Summary:** Update an exisiting `network-service-config`.
func (c *Client) NetworkServiceConfigsPatch(
	ctx context.Context,
	id string, req NetworkServiceConfigPatch,

) (NetworkServiceConfig, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-service-configs/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkServiceConfig{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ExchangeLanNetworkServiceConfigType:
			res := &ExchangeLanNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case P2PNetworkServiceConfigType:
			res := &P2PNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case P2MPNetworkServiceConfigType:
			res := &P2MPNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case MP2MPNetworkServiceConfigType:
			res := &MP2MPNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case CloudNetworkServiceConfigType:
			res := &CloudNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceConfigsDestroy **Operation:** `network_service_configs_destroy`
//
//	**Summary:** Request decommissioning the network service configuration.
//
// The network service config will assume the state
// `decommission_requested`.
// This will cascade to related resources like
// `network-feature-configs`.
func (c *Client) NetworkServiceConfigsDestroy(
	ctx context.Context,
	id string, req *CancellationRequest,

) (NetworkServiceConfig, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-service-configs/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkServiceConfig{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ExchangeLanNetworkServiceConfigType:
			res := &ExchangeLanNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case P2PNetworkServiceConfigType:
			res := &P2PNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case P2MPNetworkServiceConfigType:
			res := &P2MPNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case MP2MPNetworkServiceConfigType:
			res := &MP2MPNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		case CloudNetworkServiceConfigType:
			res := &CloudNetworkServiceConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			vlanConfig, err := decodeVLANConfig(res.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			res.VLANConfig = vlanConfig

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceConfigsPeerStatisticsReadQuery has all query parameters for NetworkServiceConfigsPeerStatisticsRead
type NetworkServiceConfigsPeerStatisticsReadQuery struct {
	// Start is a start
	Start time.Time `json:"start,omitempty"`

	// End is a end
	End time.Time `json:"end,omitempty"`

	// ASN is a asn
	ASN int `json:"asn,omitempty"`

	// MacAddress is a mac_address
	MacAddress string `json:"mac_address,omitempty"`

	// IPAddress is a ip_address
	IPAddress string `json:"ip_address,omitempty"`

	// IPVersion is a ip_version
	IPVersion int `json:"ip_version,omitempty"`
}

// RawQuery creates a query string for NetworkServiceConfigsPeerStatisticsReadQuery
func (n *NetworkServiceConfigsPeerStatisticsReadQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = n.Start.Format(time.RFC3339)
	if val != "" {
		qry.Add("start", val)
	}
	val = n.End.Format(time.RFC3339)
	if val != "" {
		qry.Add("end", val)
	}
	val = fmt.Sprintf("%v", n.ASN)
	if val != "0" {
		qry.Add("asn", val)
	}
	val = n.MacAddress
	if val != "" {
		qry.Add("mac_address", val)
	}
	val = n.IPAddress
	if val != "" {
		qry.Add("ip_address", val)
	}
	val = fmt.Sprintf("%v", n.IPVersion)
	if val != "0" {
		qry.Add("ip_version", val)
	}
	return qry.Encode()
}

// NetworkServiceConfigsPeerStatisticsRead **Operation:** `network_service_configs_peer_statistics_read`
//
//	**Summary:** Read the aggregated peer to peer statistics in relation
//	           to the peer represented by the
//	           `NetworkServiceConfig`.
//	           This operation will return a list of aggregated statistics.
//
// A `start` and `end` query parameter can be used to
// retrieve the aggregated traffic for a given window.
// In this case the key of the returned statistics is `custom`.
//
// With a given `start` and `end` window, the resolution for
// the aggregated data is chosen by the implementation.
//
// You need to check the `accuracy` attribute of the aggregate,
// to see if the data can be used for the desired
// usecase. The `accuracy` is the ratio of *total samples* to
// *expected samples*.
//
// If no `start` or `end` parameter is given, a sliding window
// is assumed and key value pairs of resolutions and aggregated
// statistics are returned.
func (c *Client) NetworkServiceConfigsPeerStatisticsRead(
	ctx context.Context,
	id string,
	qry ...*NetworkServiceConfigsPeerStatisticsReadQuery,
) ([]*PeerAggregate, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-service-configs/{id}/peer-statistics"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*PeerAggregate{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceConfigCancellationPolicyReadQuery has all query parameters for NetworkServiceConfigCancellationPolicyRead
type NetworkServiceConfigCancellationPolicyReadQuery struct {
	// DecommissionAt is a decommission_at
	DecommissionAt string `json:"decommission_at,omitempty"`
}

// RawQuery creates a query string for NetworkServiceConfigCancellationPolicyReadQuery
func (n *NetworkServiceConfigCancellationPolicyReadQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = n.DecommissionAt
	if val != "" {
		qry.Add("decommission_at", val)
	}
	return qry.Encode()
}

// NetworkServiceConfigCancellationPolicyRead **Operation:** `network_service_config_cancellation_policy_read`
//
//	**Summary:** The cancellation-policy can be queried to answer
//
// the questions:
//
// If I cancel my subscription, *when will it be technically
// decommissioned*?
// If I cancel my subscription, *until what date will I be charged*?
//
// When the query parameter `decommision_at` is not provided
// it will provide the first possible cancellation date
// and charge period if cancelled at above date.
//
// The granularity of the date field is a day, the start and end
// of which are to be interpreted by the IXP (some may use UTC,
// some may use their local time zone).
func (c *Client) NetworkServiceConfigCancellationPolicyRead(
	ctx context.Context,
	id string,
	qry ...*NetworkServiceConfigCancellationPolicyReadQuery,
) (*CancellationPolicy, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-service-configs/{id}/cancellation-policy"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &CancellationPolicy{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkFeatureConfigsListQuery has all query parameters for NetworkFeatureConfigsList
type NetworkFeatureConfigsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// StateIsNot is a state__is_not
	StateIsNot string `json:"state__is_not,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount is a consuming_account
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`

	// Type is a type
	Type string `json:"type"`

	// ServiceConfig is a service_config
	ServiceConfig string `json:"service_config,omitempty"`

	// NetworkServiceConfig is a network_service_config
	NetworkServiceConfig string `json:"network_service_config,omitempty"`

	// NetworkFeature is a network_feature
	NetworkFeature string `json:"network_feature,omitempty"`

	// RoleAssignments is a role_assignments
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Contacts is a contacts
	Contacts []string `json:"contacts,omitempty"`
}

// RawQuery creates a query string for NetworkFeatureConfigsListQuery
func (n *NetworkFeatureConfigsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(n.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", n.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", n.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = n.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = n.State
	if val != "" {
		qry.Add("state", val)
	}
	val = n.StateIsNot
	if val != "" {
		qry.Add("state__is_not", val)
	}
	val = n.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = n.ConsumingAccount
	if val != "" {
		qry.Add("consuming_account", val)
	}
	val = n.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	val = n.Type
	if val != "" {
		qry.Add("type", val)
	}
	val = n.ServiceConfig
	if val != "" {
		qry.Add("service_config", val)
	}
	val = n.NetworkServiceConfig
	if val != "" {
		qry.Add("network_service_config", val)
	}
	val = n.NetworkFeature
	if val != "" {
		qry.Add("network_feature", val)
	}
	val = strings.Join(n.RoleAssignments, ",")
	if val != "" {
		qry.Add("role_assignments", val)
	}
	val = strings.Join(n.Contacts, ",")
	if val != "" {
		qry.Add("contacts", val)
	}
	return qry.Encode()
}

// NetworkFeatureConfigsList **Operation:** `network_feature_configs_list`
//
//	**Summary:** Get all network feature configs.
func (c *Client) NetworkFeatureConfigsList(
	ctx context.Context,

	qry ...*NetworkFeatureConfigsListQuery,
) ([]NetworkFeatureConfig, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-feature-configs" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		msgs := []json.RawMessage{}
		if err := json.Unmarshal(body, &msgs); err != nil {
			return nil, err
		}
		res := make([]NetworkFeatureConfig, 0, len(msgs))

		for _, msg := range msgs {
			tmp := &PolymorphicNetworkFeatureConfig{}
			if err := json.Unmarshal(msg, tmp); err != nil {
				return nil, err
			}
			ptype := tmp.PolymorphicType()
			switch ptype {

			case RouteServerNetworkFeatureConfigType:
				rval := &RouteServerNetworkFeatureConfig{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			}
		}

		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkFeatureConfigsCreate **Operation:** `network_feature_configs_create`
//
//	**Summary:** Create a configuration for a `NetworkFeature`
//
// defined in the `NetworkFeature`s collection.
func (c *Client) NetworkFeatureConfigsCreate(
	ctx context.Context,
	req NetworkFeatureConfigRequest,

) (NetworkFeatureConfig, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-feature-configs" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkFeatureConfig{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case RouteServerNetworkFeatureConfigType:
			res := &RouteServerNetworkFeatureConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkFeatureConfigsRead **Operation:** `network_feature_configs_read`
//
//	**Summary:** Get a single network feature config.
func (c *Client) NetworkFeatureConfigsRead(
	ctx context.Context,
	id string,

) (NetworkFeatureConfig, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-feature-configs/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkFeatureConfig{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case RouteServerNetworkFeatureConfigType:
			res := &RouteServerNetworkFeatureConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkFeatureConfigsUpdate **Operation:** `network_feature_configs_update`
//
//	**Summary:** **DEPRECATION NOTICE**: This operation will be removed in favor
//
// of using `PATCH` for all updates.
//
// Update a network feature configuration
func (c *Client) NetworkFeatureConfigsUpdate(
	ctx context.Context,
	id string, req NetworkFeatureConfigUpdate,

) (NetworkFeatureConfig, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-feature-configs/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkFeatureConfig{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case RouteServerNetworkFeatureConfigType:
			res := &RouteServerNetworkFeatureConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkFeatureConfigsPatch **Operation:** `network_feature_configs_partial_update`
//
//	**Summary:** Update a network feature configuration.
func (c *Client) NetworkFeatureConfigsPatch(
	ctx context.Context,
	id string, req NetworkFeatureConfigPatch,

) (NetworkFeatureConfig, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-feature-configs/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkFeatureConfig{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case RouteServerNetworkFeatureConfigType:
			res := &RouteServerNetworkFeatureConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkFeatureConfigsDestroy **Operation:** `network_feature_configs_destroy`
//
//	**Summary:** Remove a network feature config.
//
// The network feature config will be marked as
// `decommission_requested`.
// Decommissioning a network feature config will not
// cascade to related services or service configs.
func (c *Client) NetworkFeatureConfigsDestroy(
	ctx context.Context,
	id string,

) (NetworkFeatureConfig, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-feature-configs/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkFeatureConfig{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case RouteServerNetworkFeatureConfigType:
			res := &RouteServerNetworkFeatureConfig{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// AccountRead **Operation:** `account_read`
//
//	**Summary:** Get the currently authenticated `Account`.
func (c *Client) AccountRead(
	ctx context.Context,

) (*Account, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/account" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Account{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// AccountsListQuery has all query parameters for AccountsList
type AccountsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// StateIsNot is a state__is_not
	StateIsNot string `json:"state__is_not,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// Billable is a billable
	Billable int `json:"billable,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`

	// Name is a name
	Name string `json:"name,omitempty"`

	// ASN is a asn
	ASN int `json:"asn,omitempty"`
}

// RawQuery creates a query string for AccountsListQuery
func (a *AccountsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(a.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", a.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", a.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = a.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = a.State
	if val != "" {
		qry.Add("state", val)
	}
	val = a.StateIsNot
	if val != "" {
		qry.Add("state__is_not", val)
	}
	val = a.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = fmt.Sprintf("%v", a.Billable)
	if val != "0" {
		qry.Add("billable", val)
	}
	val = a.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	val = a.Name
	if val != "" {
		qry.Add("name", val)
	}
	val = fmt.Sprintf("%v", a.ASN)
	if val != "0" {
		qry.Add("asn", val)
	}
	return qry.Encode()
}

// AccountsList **Operation:** `accounts_list`
//
//	**Summary:** Retrieve a list of `Account`s.
//
// This includes all accounts the currently authorized account
// is managing and the current account itself.
//
// Also `discoverable` accounts will be included, however
// sensitive properties, like `address` or `external_ref` will
// either not be present or redacted.
func (c *Client) AccountsList(
	ctx context.Context,

	qry ...*AccountsListQuery,
) ([]*Account, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/accounts" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*Account{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// AccountsCreate **Operation:** `accounts_create`
//
//	**Summary:** Create a new account.
func (c *Client) AccountsCreate(
	ctx context.Context,
	req *AccountRequest,

) (*Account, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/accounts" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Account{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// AccountsRead **Operation:** `accounts_read`
//
//	**Summary:** Get a single account.
func (c *Client) AccountsRead(
	ctx context.Context,
	id string,

) (*Account, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/accounts/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Account{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// AccountsUpdate **Operation:** `accounts_update`
//
//	**Summary:** **DEPRECATION NOTICE**: This operation will be removed in favor
//
// of using `PATCH` for all updates.
//
// Update the entire account.
func (c *Client) AccountsUpdate(
	ctx context.Context,
	id string, req *AccountUpdate,

) (*Account, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/accounts/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Account{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// AccountsPatch **Operation:** `accounts_partial_update`
//
//	**Summary:** Update an account.
func (c *Client) AccountsPatch(
	ctx context.Context,
	id string, req *AccountPatch,

) (*Account, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/accounts/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Account{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// AccountsDestroy **Operation:** `accounts_destroy`
//
//	**Summary:** Accounts can be deleted, when all services and configs
//
// are decommissioned or the account is not longer referenced
// e.g. as a `managing_account` or `billing_account`.
//
// Deleting an account will cascade to `contacts` and
// `role-assignments`.
//
// The request will immediately fail, if the above preconditions
// are not met.
func (c *Client) AccountsDestroy(
	ctx context.Context,
	id string,

) (*Account, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/accounts/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Account{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RolesListQuery has all query parameters for RolesList
type RolesListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// Name is a name
	Name string `json:"name,omitempty"`

	// Contact is a contact
	Contact string `json:"contact,omitempty"`
}

// RawQuery creates a query string for RolesListQuery
func (r *RolesListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(r.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", r.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", r.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = r.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = r.Name
	if val != "" {
		qry.Add("name", val)
	}
	val = r.Contact
	if val != "" {
		qry.Add("contact", val)
	}
	return qry.Encode()
}

// RolesList **Operation:** `roles_list`
//
//	**Summary:** List all roles available.
func (c *Client) RolesList(
	ctx context.Context,

	qry ...*RolesListQuery,
) ([]*Role, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/roles" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*Role{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RolesRead **Operation:** `roles_read`
//
//	**Summary:** Get a single `Role`.
func (c *Client) RolesRead(
	ctx context.Context,
	id string,

) (*Role, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/roles/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Role{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ContactsListQuery has all query parameters for ContactsList
type ContactsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount is a consuming_account
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`
}

// RawQuery creates a query string for ContactsListQuery
func (c *ContactsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(c.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", c.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", c.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = c.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = c.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = c.ConsumingAccount
	if val != "" {
		qry.Add("consuming_account", val)
	}
	val = c.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	return qry.Encode()
}

// ContactsList **Operation:** `contacts_list`
//
//	**Summary:** List available contacts managed by the authorized account.
func (c *Client) ContactsList(
	ctx context.Context,

	qry ...*ContactsListQuery,
) ([]*Contact, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/contacts" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*Contact{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ContactsCreate **Operation:** `contacts_create`
//
//	**Summary:** Create a new contact.
func (c *Client) ContactsCreate(
	ctx context.Context,
	req *ContactRequest,

) (*Contact, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/contacts" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Contact{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ContactsRead **Operation:** `contacts_read`
//
//	**Summary:** Get a contact by it's id
func (c *Client) ContactsRead(
	ctx context.Context,
	id string,

) (*Contact, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/contacts/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Contact{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ContactsUpdate **Operation:** `contacts_update`
//
//	**Summary:** **DEPRECATION NOTICE**: This operation will be removed in favor
//
// of using `PATCH` for all updates.
//
// Update a contact.
func (c *Client) ContactsUpdate(
	ctx context.Context,
	id string, req *ContactUpdate,

) (*Contact, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/contacts/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Contact{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ContactsPatch **Operation:** `contacts_partial_update`
//
//	**Summary:** Update a contact.
func (c *Client) ContactsPatch(
	ctx context.Context,
	id string, req *ContactPatch,

) (*Contact, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/contacts/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Contact{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// ContactsDestroy **Operation:** `contacts_destroy`
//
//	**Summary:** Remove a contact.
//
// Please note, that a contact can only be removed if
// it is not longer in use in a network service or config
// through a role assignment.
func (c *Client) ContactsDestroy(
	ctx context.Context,
	id string,

) (*Contact, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/contacts/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &Contact{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RoleAssignmentsListQuery has all query parameters for RoleAssignmentsList
type RoleAssignmentsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// Contact is a contact
	Contact string `json:"contact,omitempty"`

	// Role is a role
	Role string `json:"role,omitempty"`
}

// RawQuery creates a query string for RoleAssignmentsListQuery
func (r *RoleAssignmentsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(r.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", r.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", r.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = r.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = r.Contact
	if val != "" {
		qry.Add("contact", val)
	}
	val = r.Role
	if val != "" {
		qry.Add("role", val)
	}
	return qry.Encode()
}

// RoleAssignmentsList **Operation:** `role_assignments_list`
//
//	**Summary:** List all role assignments for a contact.
func (c *Client) RoleAssignmentsList(
	ctx context.Context,

	qry ...*RoleAssignmentsListQuery,
) ([]*RoleAssignment, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/role-assignments" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*RoleAssignment{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RoleAssignmentsCreate **Operation:** `role_assignments_create`
//
//	**Summary:** Assign a `Role` to a `Contact`.
//
// The contact needs to have all fields filled, which the
// role requires. If this is not the case a `400`
// `UnableToFulfill` will be returned.
func (c *Client) RoleAssignmentsCreate(
	ctx context.Context,
	req *RoleAssignmentRequest,

) (*RoleAssignment, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/role-assignments" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &RoleAssignment{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RoleAssignmentsRead **Operation:** `role_assignments_read`
//
//	**Summary:** Get a role assignment for a contact.
func (c *Client) RoleAssignmentsRead(
	ctx context.Context,
	id string,

) (*RoleAssignment, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/role-assignments/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &RoleAssignment{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RoleAssignmentsDestroy **Operation:** `role_assignments_destroy`
//
//	**Summary:** Remove a role assignment from a contact.
//
// If the contact is still in use with a given role required,
// this will yield an `UnableToFulfill` error.
func (c *Client) RoleAssignmentsDestroy(
	ctx context.Context,
	id string,

) (*RoleAssignment, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/role-assignments/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &RoleAssignment{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// APIHealthRead **Operation:** `api_health_read`
//
//	**Summary:** Get the IX-API service health status.
func (c *Client) APIHealthRead(
	ctx context.Context,

) (*APIHealth, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/health" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &APIHealth{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// APIImplementationRead **Operation:** `api_implementation_read`
//
//	**Summary:** Get the API implementation details.
func (c *Client) APIImplementationRead(
	ctx context.Context,

) (*APIImplementation, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/implementation" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &APIImplementation{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// APIExtensionsListQuery has all query parameters for APIExtensionsList
type APIExtensionsListQuery struct {
	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`
}

// RawQuery creates a query string for APIExtensionsListQuery
func (a *APIExtensionsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = fmt.Sprintf("%v", a.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", a.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = a.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	return qry.Encode()
}

// APIExtensionsList **Operation:** `api_extensions_list`
//
//	**Summary:** List provider extensions to the IX-API.
func (c *Client) APIExtensionsList(
	ctx context.Context,

	qry ...*APIExtensionsListQuery,
) ([]*APIExtension, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/extensions" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*APIExtension{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// IPsListQuery has all query parameters for IPsList
type IPsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount is a consuming_account
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`

	// NetworkService is a network_service
	NetworkService string `json:"network_service,omitempty"`

	// NetworkServiceConfig is a network_service_config
	NetworkServiceConfig string `json:"network_service_config,omitempty"`

	// NetworkFeature is a network_feature
	NetworkFeature string `json:"network_feature,omitempty"`

	// NetworkFeatureConfig is a network_feature_config
	NetworkFeatureConfig string `json:"network_feature_config,omitempty"`

	// Version is a version
	Version int `json:"version,omitempty"`

	// FQDN is a fqdn
	FQDN string `json:"fqdn,omitempty"`

	// PrefixLength is a prefix_length
	PrefixLength int `json:"prefix_length,omitempty"`

	// ValidNotBefore is a valid_not_before
	ValidNotBefore string `json:"valid_not_before,omitempty"`

	// ValidNotAfter is a valid_not_after
	ValidNotAfter string `json:"valid_not_after,omitempty"`
}

// RawQuery creates a query string for IPsListQuery
func (i *IPsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(i.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", i.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", i.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = i.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = i.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = i.ConsumingAccount
	if val != "" {
		qry.Add("consuming_account", val)
	}
	val = i.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	val = i.NetworkService
	if val != "" {
		qry.Add("network_service", val)
	}
	val = i.NetworkServiceConfig
	if val != "" {
		qry.Add("network_service_config", val)
	}
	val = i.NetworkFeature
	if val != "" {
		qry.Add("network_feature", val)
	}
	val = i.NetworkFeatureConfig
	if val != "" {
		qry.Add("network_feature_config", val)
	}
	val = fmt.Sprintf("%v", i.Version)
	if val != "0" {
		qry.Add("version", val)
	}
	val = i.FQDN
	if val != "" {
		qry.Add("fqdn", val)
	}
	val = fmt.Sprintf("%v", i.PrefixLength)
	if val != "0" {
		qry.Add("prefix_length", val)
	}
	val = i.ValidNotBefore
	if val != "" {
		qry.Add("valid_not_before", val)
	}
	val = i.ValidNotAfter
	if val != "" {
		qry.Add("valid_not_after", val)
	}
	return qry.Encode()
}

// IPsList **Operation:** `ips_list`
//
//	**Summary:** List all ip addresses (and prefixes).
func (c *Client) IPsList(
	ctx context.Context,

	qry ...*IPsListQuery,
) ([]*IPAddress, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/ips" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*IPAddress{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// IPsCreate **Operation:** `ips_create`
//
//	**Summary:** Add an ip host address or network prefix.
func (c *Client) IPsCreate(
	ctx context.Context,
	req *IPAddressRequest,

) (*IPAddress, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/ips" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &IPAddress{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// IPsRead **Operation:** `ips_read`
//
//	**Summary:** Get a single ip addresses by it's id.
func (c *Client) IPsRead(
	ctx context.Context,
	id string,

) (*IPAddress, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/ips/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &IPAddress{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// IPsUpdate **Operation:** `ips_update`
//
//	**Summary:** **DEPRECATION NOTICE**: This operation will be removed in favor
//
// of using `PATCH` for all updates.
//
// Update an ip address object.
//
// You can only update
// IP addresses within your current scope. Not all
// addresses you can read you can update.
//
// If the ip address was allocated for you, you might
// not be able to change anything but the `fqdn`.
func (c *Client) IPsUpdate(
	ctx context.Context,
	id string, req *IPAddressUpdate,

) (*IPAddress, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/ips/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &IPAddress{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// IPsPatch **Operation:** `ips_partial_update`
//
//	**Summary:** Update an ip address.
//
// You can only update
// IP addresses within your current scope. Not all
// addresses you can read you can update.
//
// If the ip address was allocated for you, you might
// not be able to change anything but the `fqdn`.
func (c *Client) IPsPatch(
	ctx context.Context,
	id string, req *IPAddressPatch,

) (*IPAddress, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/ips/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &IPAddress{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MacsListQuery has all query parameters for MacsList
type MacsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount is a consuming_account
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`

	// NetworkServiceConfig is a network_service_config
	NetworkServiceConfig string `json:"network_service_config,omitempty"`

	// Address is a address
	Address string `json:"address,omitempty"`

	// ValidNotBefore is a valid_not_before
	ValidNotBefore string `json:"valid_not_before,omitempty"`

	// ValidNotAfter is a valid_not_after
	ValidNotAfter string `json:"valid_not_after,omitempty"`
}

// RawQuery creates a query string for MacsListQuery
func (m *MacsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(m.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", m.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", m.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = m.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = m.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = m.ConsumingAccount
	if val != "" {
		qry.Add("consuming_account", val)
	}
	val = m.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	val = m.NetworkServiceConfig
	if val != "" {
		qry.Add("network_service_config", val)
	}
	val = m.Address
	if val != "" {
		qry.Add("address", val)
	}
	val = m.ValidNotBefore
	if val != "" {
		qry.Add("valid_not_before", val)
	}
	val = m.ValidNotAfter
	if val != "" {
		qry.Add("valid_not_after", val)
	}
	return qry.Encode()
}

// MacsList **Operation:** `macs_list`
//
//	**Summary:** List all mac addresses managed by the authorized customer.
func (c *Client) MacsList(
	ctx context.Context,

	qry ...*MacsListQuery,
) ([]*MacAddress, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/macs" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*MacAddress{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MacsCreate **Operation:** `macs_create`
//
//	**Summary:** Register a mac address.
func (c *Client) MacsCreate(
	ctx context.Context,
	req *MacAddressRequest,

) (*MacAddress, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/macs" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &MacAddress{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MacsRead **Operation:** `macs_read`
//
//	**Summary:** Get a single mac address by it's id.
func (c *Client) MacsRead(
	ctx context.Context,
	id string,

) (*MacAddress, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/macs/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &MacAddress{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MacsDestroy **Operation:** `macs_destroy`
//
//	**Summary:** Remove a mac address.
func (c *Client) MacsDestroy(
	ctx context.Context,
	id string,

) (*MacAddress, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/macs/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &MacAddress{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServicesListQuery has all query parameters for NetworkServicesList
type NetworkServicesListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// StateIsNot is a state__is_not
	StateIsNot string `json:"state__is_not,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount is a consuming_account
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`

	// Type is a type
	Type string `json:"type"`

	// Pop is a pop
	Pop string `json:"pop,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`
}

// RawQuery creates a query string for NetworkServicesListQuery
func (n *NetworkServicesListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(n.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", n.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", n.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = n.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = n.State
	if val != "" {
		qry.Add("state", val)
	}
	val = n.StateIsNot
	if val != "" {
		qry.Add("state__is_not", val)
	}
	val = n.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = n.ConsumingAccount
	if val != "" {
		qry.Add("consuming_account", val)
	}
	val = n.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	val = n.Type
	if val != "" {
		qry.Add("type", val)
	}
	val = n.Pop
	if val != "" {
		qry.Add("pop", val)
	}
	val = n.ProductOffering
	if val != "" {
		qry.Add("product_offering", val)
	}
	return qry.Encode()
}

// NetworkServicesList **Operation:** `network_services_list`
//
//	**Summary:** List available `NetworkService`s.
func (c *Client) NetworkServicesList(
	ctx context.Context,

	qry ...*NetworkServicesListQuery,
) ([]NetworkService, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		msgs := []json.RawMessage{}
		if err := json.Unmarshal(body, &msgs); err != nil {
			return nil, err
		}
		res := make([]NetworkService, 0, len(msgs))

		for _, msg := range msgs {
			tmp := &PolymorphicNetworkService{}
			if err := json.Unmarshal(msg, tmp); err != nil {
				return nil, err
			}
			ptype := tmp.PolymorphicType()
			switch ptype {

			case ExchangeLanNetworkServiceType:
				rval := &ExchangeLanNetworkService{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case P2PNetworkServiceType:
				rval := &P2PNetworkService{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case P2MPNetworkServiceType:
				rval := &P2MPNetworkService{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case MP2MPNetworkServiceType:
				rval := &MP2MPNetworkService{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case CloudNetworkServiceType:
				rval := &CloudNetworkService{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			}
		}

		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServicesCreate **Operation:** `network_services_create`
//
//	**Summary:** Create a new network service
func (c *Client) NetworkServicesCreate(
	ctx context.Context,
	req NetworkServiceRequest,

) (NetworkService, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkService{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ExchangeLanNetworkServiceType:
			res := &ExchangeLanNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2PNetworkServiceType:
			res := &P2PNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2MPNetworkServiceType:
			res := &P2MPNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case MP2MPNetworkServiceType:
			res := &MP2MPNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case CloudNetworkServiceType:
			res := &CloudNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServicesRead **Operation:** `network_services_read`
//
//	**Summary:** Get a specific `network-service` by id.
func (c *Client) NetworkServicesRead(
	ctx context.Context,
	id string,

) (NetworkService, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkService{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ExchangeLanNetworkServiceType:
			res := &ExchangeLanNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2PNetworkServiceType:
			res := &P2PNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2MPNetworkServiceType:
			res := &P2MPNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case MP2MPNetworkServiceType:
			res := &MP2MPNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case CloudNetworkServiceType:
			res := &CloudNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServicesUpdate **Operation:** `network_services_update`
//
//	**Summary:** **DEPRECATION NOTICE**: This operation will be removed in favor
//
// of using `PATCH` for all updates.
//
// Update a network service.
func (c *Client) NetworkServicesUpdate(
	ctx context.Context,
	id string, req NetworkServiceUpdate,

) (NetworkService, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkService{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ExchangeLanNetworkServiceType:
			res := &ExchangeLanNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2PNetworkServiceType:
			res := &P2PNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2MPNetworkServiceType:
			res := &P2MPNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case MP2MPNetworkServiceType:
			res := &MP2MPNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case CloudNetworkServiceType:
			res := &CloudNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServicesPatch **Operation:** `network_services_partial_update`
//
//	**Summary:** Update a network service
func (c *Client) NetworkServicesPatch(
	ctx context.Context,
	id string, req NetworkServicePatch,

) (NetworkService, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkService{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ExchangeLanNetworkServiceType:
			res := &ExchangeLanNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2PNetworkServiceType:
			res := &P2PNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2MPNetworkServiceType:
			res := &P2MPNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case MP2MPNetworkServiceType:
			res := &MP2MPNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case CloudNetworkServiceType:
			res := &CloudNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServicesDestroy **Operation:** `network_services_destroy`
//
//	**Summary:** Request decomissioning of the network service.
//
// The network service will enter the state of
// `decommission_requested`. The request will
// cascade to related network service and feature
// configs.
//
// An *optional request body* can be provided to request
// a specific service termination date.
//
// If no date is given in the request body, it is assumed to
// be the earliest possible date.
//
// Possible values for `decommission_at` can be queried
// through the `network_service_cancellation_policy_read`
// operation.
//
// The response will contain the dates on which the
// changes will be effected.
func (c *Client) NetworkServicesDestroy(
	ctx context.Context,
	id string, req *CancellationRequest,

) (NetworkService, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkService{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case ExchangeLanNetworkServiceType:
			res := &ExchangeLanNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2PNetworkServiceType:
			res := &P2PNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case P2MPNetworkServiceType:
			res := &P2MPNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case MP2MPNetworkServiceType:
			res := &MP2MPNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case CloudNetworkServiceType:
			res := &CloudNetworkService{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServicesStatisticsRttReadQuery has all query parameters for NetworkServicesStatisticsRttRead
type NetworkServicesStatisticsRttReadQuery struct {
	// ASN is a asn
	ASN string `json:"asn,omitempty"`

	// IP is a ip
	IP string `json:"ip,omitempty"`

	// After is a after
	After int `json:"after,omitempty"`
}

// RawQuery creates a query string for NetworkServicesStatisticsRttReadQuery
func (n *NetworkServicesStatisticsRttReadQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = n.ASN
	if val != "" {
		qry.Add("asn", val)
	}
	val = n.IP
	if val != "" {
		qry.Add("ip", val)
	}
	val = fmt.Sprintf("%v", n.After)
	if val != "0" {
		qry.Add("after", val)
	}
	return qry.Encode()
}

// NetworkServicesStatisticsRttRead **Operation:** `network_services_statistics_rtt_read`
//
//	**Summary:** Get **rtt statistics** from neighbors consuming this network
//
// service. You can filter by `asn` and peer `ip`.
//
// If the `NetworkService` does not support RTT statistics,
// a `404` error response will be returned.
//
// ### Receiving Updates
// To poll for updates, you can provide a `serial` to the
// `after` parameter in order to only receive
// results with a serial greater than the one
// provided. All of the above filters can be applied
// here and allow for receiving updates for a specific peer.
//
// If no `serial` is provided the latest RTT statistics are returned.
//
// ### Event Streaming
// This endpoint supports streaming updates using
// [**Server Sent Events**](https://html.spec.whatwg.org/multipage/server-sent-events.html#server-sent-events).
//
// To subscribe to events, negotiate the response content
// using the HTTP header: `Accept: text/event-stream`.
//
// Filteres are applied as above.
func (c *Client) NetworkServicesStatisticsRttRead(
	ctx context.Context,
	id string,
	qry ...*NetworkServicesStatisticsRttReadQuery,
) ([]*PeerRTT, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services/{id}/rtt-statistics"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*PeerRTT{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceChangeRequestRead **Operation:** `network_service_change_request_read`
//
//	**Summary:** Get the change request.
func (c *Client) NetworkServiceChangeRequestRead(
	ctx context.Context,
	id string,

) (*NetworkServiceChangeRequest, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services/{id}/change-request"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &NetworkServiceChangeRequest{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceChangeRequestCreate **Operation:** `network_service_change_request_create`
//
//	**Summary:** Request a change to the network service.
//
// The B-side participant in a peer-to-peer network service
// (`p2p_vc`) can issue a change request, expressing a
// desired change in the capacity.
//
// The change is accepted when the A-side has configured
// the network service and config with the new bandwidth.
// This is done using the `network_service_update`,
// `network_service_partial_update`,
// `network_service_config_update` or
// `network_service_config_partial_update` operations by
// the A-side.
//
// These changes can sometimes require a change of the
// product offering. The product offering may only
// differ in bandwidth.
//
// The network service will change its state from `production`
// into `production_change_pending`.
//
// A change can by rejected (by the A-side) or retracted
// (by the B-side) using the
// `network_service_change_request_destroy` operation.
//
// Only one change request may be issued at a time.
//
// A change request by the A-side is not a valid request
// and will be rejected.
func (c *Client) NetworkServiceChangeRequestCreate(
	ctx context.Context,
	id string, req *NetworkServiceChangeRequest,

) (*NetworkServiceChangeRequest, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services/{id}/change-request"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &NetworkServiceChangeRequest{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceChangeRequestDestroy **Operation:** `network_service_change_request_destroy`
//
//	**Summary:** Retract or reject a change to the network service.
func (c *Client) NetworkServiceChangeRequestDestroy(
	ctx context.Context,
	id string,

) (*NetworkServiceChangeRequest, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services/{id}/change-request"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &NetworkServiceChangeRequest{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceCancellationPolicyReadQuery has all query parameters for NetworkServiceCancellationPolicyRead
type NetworkServiceCancellationPolicyReadQuery struct {
	// DecommissionAt is a decommission_at
	DecommissionAt string `json:"decommission_at,omitempty"`
}

// RawQuery creates a query string for NetworkServiceCancellationPolicyReadQuery
func (n *NetworkServiceCancellationPolicyReadQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = n.DecommissionAt
	if val != "" {
		qry.Add("decommission_at", val)
	}
	return qry.Encode()
}

// NetworkServiceCancellationPolicyRead **Operation:** `network_service_cancellation_policy_read`
//
//	**Summary:** The cancellation-policy can be queried to answer
//
// the questions:
//
// If I cancel my service, *when will it be technically
// decommissioned*?
// If I cancel my service, *until what date will I be charged*?
//
// When the query parameter `decommision_at` is not provided
// it will provide the first possible cancellation date
// and charge period if cancelled at above date.
//
// The granularity of the date field is a day, the start and end
// of which are to be interpreted by the IXP (some may use UTC,
// some may use their local time zone).
func (c *Client) NetworkServiceCancellationPolicyRead(
	ctx context.Context,
	id string,
	qry ...*NetworkServiceCancellationPolicyReadQuery,
) (*CancellationPolicy, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-services/{id}/cancellation-policy"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &CancellationPolicy{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkFeaturesListQuery has all query parameters for NetworkFeaturesList
type NetworkFeaturesListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// Type is a type
	Type string `json:"type"`

	// Required is a required
	Required string `json:"required,omitempty"`

	// NetworkService is a network_service
	NetworkService string `json:"network_service,omitempty"`

	// Name is a name
	Name string `json:"name,omitempty"`
}

// RawQuery creates a query string for NetworkFeaturesListQuery
func (n *NetworkFeaturesListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(n.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", n.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", n.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = n.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = n.Type
	if val != "" {
		qry.Add("type", val)
	}
	val = n.Required
	if val != "" {
		qry.Add("required", val)
	}
	val = n.NetworkService
	if val != "" {
		qry.Add("network_service", val)
	}
	val = n.Name
	if val != "" {
		qry.Add("name", val)
	}
	return qry.Encode()
}

// NetworkFeaturesList **Operation:** `network_features_list`
//
//	**Summary:** List available network features.
func (c *Client) NetworkFeaturesList(
	ctx context.Context,

	qry ...*NetworkFeaturesListQuery,
) ([]NetworkFeature, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-features" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		msgs := []json.RawMessage{}
		if err := json.Unmarshal(body, &msgs); err != nil {
			return nil, err
		}
		res := make([]NetworkFeature, 0, len(msgs))

		for _, msg := range msgs {
			tmp := &PolymorphicNetworkFeature{}
			if err := json.Unmarshal(msg, tmp); err != nil {
				return nil, err
			}
			ptype := tmp.PolymorphicType()
			switch ptype {

			case RouteServerNetworkFeatureType:
				rval := &RouteServerNetworkFeature{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			}
		}

		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkFeaturesRead **Operation:** `network_features_read`
//
//	**Summary:** Get a single network feature by it's id.
func (c *Client) NetworkFeaturesRead(
	ctx context.Context,
	id string,

) (NetworkFeature, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/network-features/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicNetworkFeature{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case RouteServerNetworkFeatureType:
			res := &RouteServerNetworkFeature{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MemberJoiningRulesListQuery has all query parameters for MemberJoiningRulesList
type MemberJoiningRulesListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount is a consuming_account
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`

	// NetworkService is a network_service
	NetworkService string `json:"network_service,omitempty"`
}

// RawQuery creates a query string for MemberJoiningRulesListQuery
func (m *MemberJoiningRulesListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(m.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", m.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", m.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = m.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = m.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = m.ConsumingAccount
	if val != "" {
		qry.Add("consuming_account", val)
	}
	val = m.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	val = m.NetworkService
	if val != "" {
		qry.Add("network_service", val)
	}
	return qry.Encode()
}

// MemberJoiningRulesList **Operation:** `member_joining_rules_list`
//
//	**Summary:** Get a list of joining rules
func (c *Client) MemberJoiningRulesList(
	ctx context.Context,

	qry ...*MemberJoiningRulesListQuery,
) ([]MemberJoiningRule, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/member-joining-rules" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		msgs := []json.RawMessage{}
		if err := json.Unmarshal(body, &msgs); err != nil {
			return nil, err
		}
		res := make([]MemberJoiningRule, 0, len(msgs))

		for _, msg := range msgs {
			tmp := &PolymorphicMemberJoiningRule{}
			if err := json.Unmarshal(msg, tmp); err != nil {
				return nil, err
			}
			ptype := tmp.PolymorphicType()
			switch ptype {

			case AllowMemberJoiningRuleType:
				rval := &AllowMemberJoiningRule{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			case DenyMemberJoiningRuleType:
				rval := &DenyMemberJoiningRule{}
				if err := json.Unmarshal(msg, rval); err != nil {
					return nil, err
				}

				res = append(res, rval)

			}
		}

		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MemberJoiningRulesCreate **Operation:** `member_joining_rules_create`
//
//	**Summary:** Create a member joining rule
func (c *Client) MemberJoiningRulesCreate(
	ctx context.Context,
	req MemberJoiningRuleRequest,

) (MemberJoiningRule, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/member-joining-rules" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicMemberJoiningRule{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case AllowMemberJoiningRuleType:
			res := &AllowMemberJoiningRule{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case DenyMemberJoiningRuleType:
			res := &DenyMemberJoiningRule{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MemberJoiningRulesRead **Operation:** `member_joining_rules_read`
//
//	**Summary:** Get a single rule
func (c *Client) MemberJoiningRulesRead(
	ctx context.Context,
	id string,

) (MemberJoiningRule, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/member-joining-rules/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicMemberJoiningRule{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case AllowMemberJoiningRuleType:
			res := &AllowMemberJoiningRule{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case DenyMemberJoiningRuleType:
			res := &DenyMemberJoiningRule{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MemberJoiningRulesUpdate **Operation:** `member_joining_rules_update`
//
//	**Summary:** **DEPRECATION NOTICE**: This operation will be removed in favor
//
// of using `PATCH` for all updates.
//
// Update a joining rule.
func (c *Client) MemberJoiningRulesUpdate(
	ctx context.Context,
	id string, req MemberJoiningRuleUpdate,

) (MemberJoiningRule, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/member-joining-rules/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicMemberJoiningRule{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case AllowMemberJoiningRuleType:
			res := &AllowMemberJoiningRule{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case DenyMemberJoiningRuleType:
			res := &DenyMemberJoiningRule{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MemberJoiningRulesPatch **Operation:** `member_joining_rules_partial_update`
//
//	**Summary:** Update a joining rule
func (c *Client) MemberJoiningRulesPatch(
	ctx context.Context,
	id string, req MemberJoiningRulePatch,

) (MemberJoiningRule, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/member-joining-rules/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicMemberJoiningRule{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case AllowMemberJoiningRuleType:
			res := &AllowMemberJoiningRule{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case DenyMemberJoiningRuleType:
			res := &DenyMemberJoiningRule{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// MemberJoiningRulesDestroy **Operation:** `member_joining_rules_destroy`
//
//	**Summary:** Delete a joining rule
func (c *Client) MemberJoiningRulesDestroy(
	ctx context.Context,
	id string,

) (MemberJoiningRule, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/member-joining-rules/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		tmp := &PolymorphicMemberJoiningRule{}
		if err := json.Unmarshal(body, tmp); err != nil {
			return nil, err
		}
		ptype := tmp.PolymorphicType()
		switch ptype {

		case AllowMemberJoiningRuleType:
			res := &AllowMemberJoiningRule{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		case DenyMemberJoiningRuleType:
			res := &DenyMemberJoiningRule{}
			if err := json.Unmarshal(body, res); err != nil {
				return nil, err
			}

			return res, nil

		}

		return nil, ErrInvalidPolymorphicType

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RoutingFunctionsListQuery has all query parameters for RoutingFunctionsList
type RoutingFunctionsListQuery struct {
	// ID is a id
	ID []string `json:"id,omitempty"`

	// PageLimit is a page_limit
	PageLimit int `json:"page_limit,omitempty"`

	// PageOffset is a page_offset
	PageOffset int `json:"page_offset,omitempty"`

	// PageToken is a page_token
	PageToken string `json:"page_token,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// StateIsNot is a state__is_not
	StateIsNot string `json:"state__is_not,omitempty"`

	// ManagingAccount is a managing_account
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount is a consuming_account
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef is a external_ref
	ExternalRef string `json:"external_ref,omitempty"`

	// ASN is a asn
	ASN string `json:"asn,omitempty"`
}

// RawQuery creates a query string for RoutingFunctionsListQuery
func (r *RoutingFunctionsListQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = strings.Join(r.ID, ",")
	if val != "" {
		qry.Add("id", val)
	}
	val = fmt.Sprintf("%v", r.PageLimit)
	if val != "0" {
		qry.Add("page_limit", val)
	}
	val = fmt.Sprintf("%v", r.PageOffset)
	if val != "0" {
		qry.Add("page_offset", val)
	}
	val = r.PageToken
	if val != "" {
		qry.Add("page_token", val)
	}
	val = r.State
	if val != "" {
		qry.Add("state", val)
	}
	val = r.StateIsNot
	if val != "" {
		qry.Add("state__is_not", val)
	}
	val = r.ManagingAccount
	if val != "" {
		qry.Add("managing_account", val)
	}
	val = r.ConsumingAccount
	if val != "" {
		qry.Add("consuming_account", val)
	}
	val = r.ExternalRef
	if val != "" {
		qry.Add("external_ref", val)
	}
	val = r.ASN
	if val != "" {
		qry.Add("asn", val)
	}
	return qry.Encode()
}

// RoutingFunctionsList **Operation:** `routing_functions_list`
//
//	**Summary:** Get a list of routing functions
func (c *Client) RoutingFunctionsList(
	ctx context.Context,

	qry ...*RoutingFunctionsListQuery,
) ([]*RoutingFunction, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/routing-functions" + params)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := []*RoutingFunction{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RoutingFunctionsCreate **Operation:** `routing_functions_create`
//
//	**Summary:** Create a routing function
func (c *Client) RoutingFunctionsCreate(
	ctx context.Context,
	req *RoutingFunctionRequest,

) (*RoutingFunction, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/routing-functions" + params)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &RoutingFunction{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RoutingFunctionsRead **Operation:** `routing_functions_read`
//
//	**Summary:** Get a single routing function instance.
func (c *Client) RoutingFunctionsRead(
	ctx context.Context,
	id string,

) (*RoutingFunction, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/routing-functions/{id}"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &RoutingFunction{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RoutingFunctionsPatch **Operation:** `routing_functions_partial_update`
//
//	**Summary:** Update a routing function
func (c *Client) RoutingFunctionsPatch(
	ctx context.Context,
	id string, req *RoutingFunctionPatch,

) (*RoutingFunction, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/routing-functions/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &RoutingFunction{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RoutingFunctionsDestroy **Operation:** `routing_functions_destroy`
//
//	**Summary:** Request decomissioning the routing function.
//
// The cancellation policy of the routing function applies
// here and is independent from the
// policy of the network-service and network-service-config
// using the routing function.
//
// The routing function will assume the state
// `decommission_requested`.
//
// The decommissioning request will *not* cascade
// to network services and configs.
func (c *Client) RoutingFunctionsDestroy(
	ctx context.Context,
	id string, req *CancellationRequest,

) (*RoutingFunction, error) {

	params := ""
	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/routing-functions/{id}"+params, id)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodDelete, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &RoutingFunction{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// RoutingFunctionsCancellationPolicyReadQuery has all query parameters for RoutingFunctionsCancellationPolicyRead
type RoutingFunctionsCancellationPolicyReadQuery struct {
	// DecommissionAt is a decommission_at
	DecommissionAt string `json:"decommission_at,omitempty"`
}

// RawQuery creates a query string for RoutingFunctionsCancellationPolicyReadQuery
func (r *RoutingFunctionsCancellationPolicyReadQuery) RawQuery() string {
	qry := url.Values{}
	val := ""
	val = r.DecommissionAt
	if val != "" {
		qry.Add("decommission_at", val)
	}
	return qry.Encode()
}

// RoutingFunctionsCancellationPolicyRead **Operation:** `routing_functions_cancellation_policy_read`
//
//	**Summary:** The cancellation-policy can be queried to answer
//
// the questions:
//
// If I cancel my subscription, *when will it be technically
// decommissioned*?
// If I cancel my subscription, *until what date will I be charged*?
//
// When the query parameter `decommision_at` is not provided
// it will provide the first possible cancellation date
// and charge period if cancelled at above date.
//
// The granularity of the date field is a day, the start and end
// of which are to be interpreted by the IXP (some may use UTC,
// some may use their local time zone).
func (c *Client) RoutingFunctionsCancellationPolicyRead(
	ctx context.Context,
	id string,
	qry ...*RoutingFunctionsCancellationPolicyReadQuery,
) (*CancellationPolicy, error) {

	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	if params != "" {
		params = "?" + params
	}

	url := c.resourceURL("/routing-functions/{id}/cancellation-policy"+params, id)

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, err
	}

	// Success
	if ret.StatusCode <= http.StatusAccepted {

		res := &CancellationPolicy{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil

	}

	// Decode error 404
	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode // implementations are not reliable
		return nil, res
	}

	// Decode error 403
	if ret.StatusCode == http.StatusForbidden {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 401
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	// Decode error 400
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	// Decode as generic error
	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}
