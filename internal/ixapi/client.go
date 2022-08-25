package ixapi

import (
	"context"
	"net/http"
	"strings"
)

// Errors

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
