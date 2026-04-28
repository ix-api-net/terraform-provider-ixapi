package ixapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
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

	APIURL             string
	header             http.Header
	CloudRouterEnabled bool
}

// NewClient creates a new client instance
func NewClient(server, version string) *Client {
	c := &Client{
		APIURL: server,
		header: http.Header{},
	}
	if version != "" {
		c.header.Set("User-Agent", "terraform-provider-ixapi/"+version)
	}
	return c
}

// hostBase extracts only the scheme and host from a URL, stripping any path component.
// This is needed because DE-CIX extension API paths are absolute from the server root (e.g. /api/v3/decix-vrf-v1/...),
// while the configured APIURL may include a versioned path suffix (e.g. https://host/v2).
func hostBase(apiURL string) string {
	if u, err := url.Parse(apiURL); err == nil && u.Scheme != "" {
		return u.Scheme + "://" + u.Host
	}
	return strings.TrimSuffix(apiURL, "/")
}

// resourceURL concatenates the API base with the resource path, substituting {id} if provided.
func (c *Client) resourceURL(res string, params ...string) string {
	p := strings.TrimSuffix(c.APIURL, "/") + res
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

// RequireCloudRouterExtension returns an error if the DE-CIX Cloud Router extension is not enabled in the provider configuration.
func (c *Client) RequireCloudRouterExtension() error {
	if !c.CloudRouterEnabled {
		return fmt.Errorf(
			"CloudRouter extension is not enabled. " +
				"Add 'extension_de_cix_cloud_router_enabled = true' to your provider configuration to use CloudRouter features",
		)
	}
	return nil
}
