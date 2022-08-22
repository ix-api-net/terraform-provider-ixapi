package ixapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Errors

// APIError is a generic api error
type APIError struct {
	ProblemResponse
}

// NotFoundError indicates that a resource was not found
type NotFoundError struct {
	ProblemResponse
}

// AuthenticationError indicates that the authentication
// was not successful.
type AuthenticationError struct {
	ProblemResponse
}

// ValidationError indicates that the validation of user data
// failed. The Properties attribute should contain
// a list of property names and reasons.
type ValidationError struct {
	ProblemResponse
	Properties []ValidationErrorProp `json:"properties"`
}

// ValidationErrorProp A failed validation
type ValidationErrorProp struct {
	// Name is a name
	Name string `json:"name,omitempty"`

	// Reason is a reason
	Reason json.RawMessage `json:"reason,omitempty"`
}

func collectReasons(res map[string]interface{}) string {
	reasons := "{ "
	for k, v := range res {
		reason := v.(string)
		reasons += fmt.Sprintf("%s: %s", k, reason)
	}
	reasons += " }"
	return reasons
}

// Error implements the error interface
func (e ValidationError) Error() string {
	props := ""
	plen := len(e.Properties) - 1
	for i, prop := range e.Properties {
		props += fmt.Sprintf("%s: %s", prop.Name, prop.Reason)
		if i < plen {
			props += ", "
		}
	}
	return fmt.Sprintf("%s %s",
		e.Title, props)
}

// PermissionError indicates that insufficient rights were
// given, when trying to access a resource.
type PermissionError struct {
	ProblemResponse
}

// Authentication Strategies

// AuthenticationProvider is an interface for authentication
type AuthenticationProvider interface {
	authenticate(ctx context.Context, client *Client) error
}

// AuthAPIKeySecret will use the legacy authentication endpoint
type AuthAPIKeySecret struct {
	Key    string
	Secret string
}

// Private authenticate the client using key secret credentials
func (s *AuthAPIKeySecret) authenticate(
	ctx context.Context,
	c *Client,
) error {
	auth, err := c.AuthTokenCreate(ctx, &AuthTokenRequest{
		APIKey:    s.Key,
		APISecret: s.Secret,
	})
	if err != nil {
		return err
	}
	c.header.Set("Authorization", "Bearer "+auth.AccessToken)
	return nil
}

// Client is an IX-API http client
type Client struct {
	http.Client

	APIURL string
	header http.Header
}

// NewClient creates a new client instance
func NewClient(server string) *Client {
	return &Client{
		APIURL: server,
		header: http.Header{},
	}
}

// Private resourceURL concatinates the api base with the resource
func (c *Client) resourceURL(res string, params ...string) string {
	base := c.APIURL
	if strings.HasSuffix(base, "/") {
		base = base[:len(base)-1]
	}
	p := base + res
	if len(params) > 0 {
		p = strings.ReplaceAll(p, "{id}", params[0])
	}
	return p
}

// Authenticate using a authentication provider
func (c *Client) Authenticate(
	ctx context.Context,
	auth AuthenticationProvider,
) error {
	return auth.authenticate(ctx, c)
}
