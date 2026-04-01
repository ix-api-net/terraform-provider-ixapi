package ixapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// cloudRouterURL builds an absolute URL for a DE-CIX extension API path.
// It uses the host base of the configured API URL because the extension API is mounted
// at the server root (e.g. /api/v3/decix-vrf-v1/...) rather than under the versioned IX-API prefix.
// An optional id substitutes {id} in the path; an optional query string is appended if non-empty.
func (c *Client) cloudRouterURL(path, query string, id ...string) string {
	if len(id) > 0 {
		path = strings.ReplaceAll(path, "{id}", id[0])
	}
	u := hostBase(c.APIURL) + path
	if query != "" {
		u += "?" + query
	}
	return u
}

// CloudRoutersListQuery holds query parameters for listing Cloud Router resources.
type CloudRoutersListQuery struct {
	ManagingAccount  string `json:"managing_account,omitempty"`
	ConsumingAccount string `json:"consuming_account,omitempty"`
	ExternalRef      string `json:"external_ref,omitempty"`
}

// RawQuery encodes the query parameters as a URL query string.
func (q *CloudRoutersListQuery) RawQuery() string {
	qry := url.Values{}
	if q.ManagingAccount != "" {
		qry.Add("managing_account", q.ManagingAccount)
	}
	if q.ConsumingAccount != "" {
		qry.Add("consuming_account", q.ConsumingAccount)
	}
	if q.ExternalRef != "" {
		qry.Add("external_ref", q.ExternalRef)
	}
	return qry.Encode()
}

// CloudRoutersList fetches all Cloud Routers, optionally filtered by query parameters.
func (c *Client) CloudRoutersList(
	ctx context.Context,
	qry ...*CloudRoutersListQuery,
) ([]*CloudRouter, error) {
	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/vrfs", params)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := []*CloudRouter{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// CloudRoutersCreate creates a new Cloud Router and returns the created resource.
func (c *Client) CloudRoutersCreate(
	ctx context.Context,
	req *CloudRouterRequest,
) (*CloudRouter, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/vrfs", "")
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

	if ret.StatusCode <= http.StatusAccepted {
		res := &CloudRouter{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// CloudRoutersRead fetches a single Cloud Router by ID.
func (c *Client) CloudRoutersRead(
	ctx context.Context,
	id string,
) (*CloudRouter, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/vrfs/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &CloudRouter{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// CloudRoutersDestroy deletes a Cloud Router by ID.
func (c *Client) CloudRoutersDestroy(
	ctx context.Context,
	id string,
) error {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/vrfs/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return err
	}

	if ret.StatusCode <= http.StatusAccepted {
		return nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return err
		}
		res.Status = ret.StatusCode
		return res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return err
		}
		res.Status = ret.StatusCode
		return res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return err
		}
		res.Status = ret.StatusCode
		return res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return err
		}
		res.Status = ret.StatusCode
		return res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return err
	}
	res.Status = ret.StatusCode
	return res
}

// CloudRouterNetworkServiceConfigsListQuery holds query parameters for listing Cloud Router network service configs.
type CloudRouterNetworkServiceConfigsListQuery struct {
	Type        string `json:"type,omitempty"`
	BGPPassword string `json:"bgp_password,omitempty"`
	BFD         *bool  `json:"bfd,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	Offset      int    `json:"offset,omitempty"`
}

// RawQuery encodes the query parameters as a URL query string.
func (q *CloudRouterNetworkServiceConfigsListQuery) RawQuery() string {
	qry := url.Values{}
	if q.Type != "" {
		qry.Add("type", q.Type)
	}
	if q.BGPPassword != "" {
		qry.Add("bgp_password", q.BGPPassword)
	}
	if q.BFD != nil {
		if *q.BFD {
			qry.Add("bfd", "1")
		} else {
			qry.Add("bfd", "0")
		}
	}
	if q.Limit > 0 {
		qry.Add("limit", fmt.Sprintf("%d", q.Limit))
	}
	if q.Offset > 0 {
		qry.Add("offset", fmt.Sprintf("%d", q.Offset))
	}
	return qry.Encode()
}

// CloudRouterNetworkServiceConfigsList fetches all Cloud Router network service configs, optionally filtered by query parameters.
func (c *Client) CloudRouterNetworkServiceConfigsList(
	ctx context.Context,
	qry ...*CloudRouterNetworkServiceConfigsListQuery,
) ([]*CloudRouterNetworkServiceConfig, error) {
	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/network-service-configs", params)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := []*CloudRouterNetworkServiceConfig{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		for _, item := range res {
			vlanConfig, err := decodeVLANConfig(item.VLANConfigRaw)
			if err != nil {
				return nil, err
			}
			item.VLANConfig = vlanConfig
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// CloudRouterNetworkServiceConfigsCreate creates a new Cloud Router network service config and returns the created resource.
func (c *Client) CloudRouterNetworkServiceConfigsCreate(
	ctx context.Context,
	req *CloudRouterNetworkServiceConfigRequest,
) (*CloudRouterNetworkServiceConfig, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/network-service-configs", "")
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

	if ret.StatusCode <= http.StatusAccepted {
		res := &CloudRouterNetworkServiceConfig{}
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

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// CloudRouterNetworkServiceConfigsRead fetches a single Cloud Router network service config by ID.
func (c *Client) CloudRouterNetworkServiceConfigsRead(
	ctx context.Context,
	id string,
) (*CloudRouterNetworkServiceConfig, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/network-service-configs/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &CloudRouterNetworkServiceConfig{}
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

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// CloudRouterNetworkServiceConfigsPatch partially updates a Cloud Router network service config by ID.
func (c *Client) CloudRouterNetworkServiceConfigsPatch(
	ctx context.Context,
	id string,
	patch *CloudRouterNetworkServiceConfigPatch,
) (*CloudRouterNetworkServiceConfig, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/network-service-configs/{id}", "", id)
	data, err := json.Marshal(patch)
	if err != nil {
		return nil, err
	}

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &CloudRouterNetworkServiceConfig{}
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

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// CloudRouterNetworkServiceConfigsDestroy deletes a Cloud Router network service config by ID.
func (c *Client) CloudRouterNetworkServiceConfigsDestroy(
	ctx context.Context,
	id string,
) error {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/network-service-configs/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	for k, v := range c.header {
		hreq.Header.Set(k, v[0])
	}
	ret, err := c.Do(hreq)
	if err != nil {
		return err
	}
	defer ret.Body.Close()
	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return err
	}

	if ret.StatusCode <= http.StatusAccepted {
		return nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return err
		}
		res.Status = ret.StatusCode
		return res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return err
		}
		res.Status = ret.StatusCode
		return res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return err
		}
		res.Status = ret.StatusCode
		return res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return err
		}
		res.Status = ret.StatusCode
		return res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return err
	}
	res.Status = ret.StatusCode
	return res
}

// CloudRouterProductOfferingsListQuery holds query parameters for listing Cloud Router product offerings.
type CloudRouterProductOfferingsListQuery struct {
	Limit                    int    `json:"limit,omitempty"`
	Offset                   int    `json:"offset,omitempty"`
	ID                       string `json:"id,omitempty"`
	Bandwidth                int    `json:"bandwidth,omitempty"`
	Name                     string `json:"name,omitempty"`
	ServiceMetroArea         string `json:"service_metro_area,omitempty"`
	ServiceMetroAreaNetwork  string `json:"service_metro_area_network,omitempty"`
	ContractPeriod           string `json:"contract_period,omitempty"`
}

// RawQuery encodes the query parameters as a URL query string.
func (q *CloudRouterProductOfferingsListQuery) RawQuery() string {
	qry := url.Values{}
	if q.Limit > 0 {
		qry.Add("limit", fmt.Sprintf("%d", q.Limit))
	}
	if q.Offset > 0 {
		qry.Add("offset", fmt.Sprintf("%d", q.Offset))
	}
	if q.ID != "" {
		qry.Add("id", q.ID)
	}
	if q.Bandwidth > 0 {
		qry.Add("bandwidth", fmt.Sprintf("%d", q.Bandwidth))
	}
	if q.Name != "" {
		qry.Add("name", q.Name)
	}
	if q.ServiceMetroArea != "" {
		qry.Add("service_metro_area", q.ServiceMetroArea)
	}
	if q.ServiceMetroAreaNetwork != "" {
		qry.Add("service_metro_area_network", q.ServiceMetroAreaNetwork)
	}
	if q.ContractPeriod != "" {
		qry.Add("contract_period", q.ContractPeriod)
	}
	return qry.Encode()
}

// CloudRouterProductOfferingsList fetches all Cloud Router product offerings, optionally filtered by query parameters.
func (c *Client) CloudRouterProductOfferingsList(
	ctx context.Context,
	qry ...*CloudRouterProductOfferingsListQuery,
) ([]*CloudRouterProductOffering, error) {
	params := ""
	if len(qry) > 0 && qry[0] != nil {
		params = qry[0].RawQuery()
	}

	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/product-offerings", params)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := []*CloudRouterProductOffering{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// CloudRouterProductOfferingsRead fetches a single Cloud Router product offering by ID.
func (c *Client) CloudRouterProductOfferingsRead(
	ctx context.Context,
	id string,
) (*CloudRouterProductOffering, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/product-offerings/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &CloudRouterProductOffering{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// CloudRouterNetworkServiceConfigGetBGPState fetches the BGP session state for a given network service config ID.
func (c *Client) CloudRouterNetworkServiceConfigGetBGPState(
	ctx context.Context,
	id string,
) (*BGPStateResponse, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/network-service-configs/{id}/bgp-state", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &BGPStateResponse{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// CloudRouterNetworkServiceConfigGetBFDState fetches the BFD session state for a given network service config ID.
func (c *Client) CloudRouterNetworkServiceConfigGetBFDState(
	ctx context.Context,
	id string,
) (*BFDStateResponse, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/network-service-configs/{id}/bfd-state", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &BFDStateResponse{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// Prefix Lists

// PrefixListsList fetches all prefix lists, optionally filtered by managing account.
func (c *Client) PrefixListsList(
	ctx context.Context,
	managingAccount string,
) ([]*PrefixList, error) {
	params := ""
	if managingAccount != "" {
		params = "managing_account=" + url.QueryEscape(managingAccount)
	}

	u := c.cloudRouterURL("/api/v3/decix-vrf-v1/prefix-lists", params)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := []*PrefixList{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PrefixListsCreate creates a new prefix list and returns the created resource.
func (c *Client) PrefixListsCreate(
	ctx context.Context,
	req *PrefixListRequest,
) (*PrefixList, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/prefix-lists", "")
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

	if ret.StatusCode <= http.StatusAccepted {
		res := &PrefixList{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PrefixListsRead fetches a single prefix list by ID.
func (c *Client) PrefixListsRead(
	ctx context.Context,
	id string,
) (*PrefixList, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/prefix-lists/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &PrefixList{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PrefixListsUpdate replaces a prefix list by ID and returns the updated resource.
func (c *Client) PrefixListsUpdate(
	ctx context.Context,
	id string,
	req *PrefixListRequest,
) (*PrefixList, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/prefix-lists/{id}", "", id)
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

	if ret.StatusCode <= http.StatusAccepted {
		res := &PrefixList{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PrefixListsDelete deletes a prefix list by ID and returns the deleted resource.
func (c *Client) PrefixListsDelete(
	ctx context.Context,
	id string,
) (*PrefixList, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/prefix-lists/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &PrefixList{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// Policies

// PoliciesList fetches all routing policies, optionally filtered by managing account.
func (c *Client) PoliciesList(
	ctx context.Context,
	managingAccount string,
) ([]*Policy, error) {
	params := ""
	if managingAccount != "" {
		params = "managing_account=" + url.QueryEscape(managingAccount)
	}

	u := c.cloudRouterURL("/api/v3/decix-vrf-v1/policies", params)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := []*Policy{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PoliciesCreate creates a new routing policy and returns the created resource.
func (c *Client) PoliciesCreate(
	ctx context.Context,
	req *PolicyRequest,
) (*Policy, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/policies", "")
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

	if ret.StatusCode <= http.StatusAccepted {
		res := &Policy{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PoliciesRead fetches a single routing policy by ID.
func (c *Client) PoliciesRead(
	ctx context.Context,
	id string,
) (*Policy, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/policies/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &Policy{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PoliciesUpdate replaces a routing policy by ID and returns the updated resource.
func (c *Client) PoliciesUpdate(
	ctx context.Context,
	id string,
	req *PolicyRequest,
) (*Policy, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/policies/{id}", "", id)
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

	if ret.StatusCode <= http.StatusAccepted {
		res := &Policy{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// PoliciesDelete deletes a routing policy by ID and returns the deleted resource.
func (c *Client) PoliciesDelete(
	ctx context.Context,
	id string,
) (*Policy, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/policies/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &Policy{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// VRF ARP Table

// ArpTableList fetches ARP table entries, optionally filtered by VRF ID and network service config ID.
func (c *Client) ArpTableList(
	ctx context.Context,
	vrfID string,
	nscID string,
) ([]*ArpEntry, error) {
	params := url.Values{}
	if vrfID != "" {
		params.Set("vrf", vrfID)
	}
	if nscID != "" {
		params.Set("network_service_config", nscID)
	}
	u := c.cloudRouterURL("/api/v3/decix-vrf-v1/arp-table", params.Encode())
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := []*ArpEntry{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// VRF Routes

// VrfRoutesList fetches routes from a VRF routing table, optionally filtered by VRF ID.
func (c *Client) VrfRoutesList(
	ctx context.Context,
	vrfID string,
) ([]*VrfRoute, error) {
	params := url.Values{}
	if vrfID != "" {
		params.Set("vrf", vrfID)
	}
	u := c.cloudRouterURL("/api/v3/decix-vrf-v1/routes", params.Encode())
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := []*VrfRoute{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// Static Routes

// StaticRoutesList fetches static routes, optionally filtered by VRF ID and network service config ID.
func (c *Client) StaticRoutesList(
	ctx context.Context,
	vrfID string,
	nscID string,
) ([]*StaticRoute, error) {
	params := url.Values{}
	if vrfID != "" {
		params.Set("vrf", vrfID)
	}
	if nscID != "" {
		params.Set("network_service_config", nscID)
	}
	u := c.cloudRouterURL("/api/v3/decix-vrf-v1/static-routes", params.Encode())
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := []*StaticRoute{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// StaticRoutesCreate creates a new static route and returns the created resource.
func (c *Client) StaticRoutesCreate(
	ctx context.Context,
	req *StaticRouteRequest,
) (*StaticRoute, error) {
	u := c.cloudRouterURL("/api/v3/decix-vrf-v1/static-routes", "")
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPost, u, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &StaticRoute{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// StaticRoutesRead fetches a single static route by ID.
func (c *Client) StaticRoutesRead(
	ctx context.Context,
	id string,
) (*StaticRoute, error) {
	u := c.cloudRouterURL("/api/v3/decix-vrf-v1/static-routes/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &StaticRoute{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// StaticRoutesUpdate replaces a static route by ID and returns the updated resource.
func (c *Client) StaticRoutesUpdate(
	ctx context.Context,
	id string,
	req *StaticRouteRequest,
) (*StaticRoute, error) {
	u := c.cloudRouterURL("/api/v3/decix-vrf-v1/static-routes/{id}", "", id)
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	hreq, err := http.NewRequestWithContext(
		ctx, http.MethodPut, u, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &StaticRoute{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusBadRequest {
		res := &ValidationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// StaticRoutesDelete deletes a static route by ID and returns the deleted resource.
func (c *Client) StaticRoutesDelete(
	ctx context.Context,
	id string,
) (*StaticRoute, error) {
	u := c.cloudRouterURL("/api/v3/decix-vrf-v1/static-routes/{id}", "", id)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := &StaticRoute{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusForbidden {
		res := &PermissionError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NSC Routes

// NetworkServiceConfigReceivedRoutesList fetches BGP routes received from peers on a given network service config.
func (c *Client) NetworkServiceConfigReceivedRoutesList(
	ctx context.Context,
	nscID string,
) ([]*BGPRoute, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/network-service-configs/{id}/received-routes", "", nscID)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := []*BGPRoute{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}

// NetworkServiceConfigAdvertisedRoutesList fetches BGP routes advertised to peers on a given network service config.
func (c *Client) NetworkServiceConfigAdvertisedRoutesList(
	ctx context.Context,
	nscID string,
) ([]*BGPRoute, error) {
	url := c.cloudRouterURL("/api/v3/decix-vrf-v1/network-service-configs/{id}/advertised-routes", "", nscID)
	hreq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	if ret.StatusCode <= http.StatusAccepted {
		res := []*BGPRoute{}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if ret.StatusCode == http.StatusNotFound {
		res := &NotFoundError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}
	if ret.StatusCode == http.StatusUnauthorized {
		res := &AuthenticationError{}
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}
		res.Status = ret.StatusCode
		return nil, res
	}

	res := &APIError{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	res.Status = ret.StatusCode
	return nil, res
}
