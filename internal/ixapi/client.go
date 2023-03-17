package ixapi

import (
	"context"
	"net/http"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
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
	c.SetBearer(auth.AccessToken)
	return nil
}

// OAuth2ClientCredentials will use OAuth2 for authentication
type OAuth2ClientCredentials struct {
	Key      string
	Secret   string
	TokenURL string
	Scopes   []string
}

// Implement AuthenticationProvider interface
func (flow *OAuth2ClientCredentials) authenticate(
	ctx context.Context,
	c *Client,
) error {
	config := &clientcredentials.Config{
		ClientID:     flow.Key,
		ClientSecret: flow.Secret,
		TokenURL:     flow.TokenURL,
		Scopes:       flow.Scopes,
	}
	token, err := config.Token(ctx)
	if err != nil {
		return err
	}
	c.SetBearer(token.AccessToken)
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

// SetBearer allows setting the bearer token in the client.
// This can be used to implement custom authentication.
func (c *Client) SetBearer(token string) {
	c.header.Set("Authorization", "Bearer "+token)
}

// Authenticate using a authentication provider
func (c *Client) Authenticate(
	ctx context.Context,
	auth AuthenticationProvider,
) error {
	return auth.authenticate(ctx, c)
}
