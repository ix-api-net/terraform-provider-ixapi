
**This repository is work in progress and not an official release.**

# IX-API Terraform Provider

This Terraform provider is using the [ix-api](https://ix-api.net)
for configuring and provisisioning IXP services.


## Requirements
 * Terraform >= 1.0
 * Go >= 1.17


## Configure the provider

Simple provider configuration example using
the `legacy` authentication strategy.

```hcl
provider "ixapi" {
    api = "http://localhost:8000/api/v2"
    api_key = "my_api_key"
    # api_secret = "" # Use $IX_API_SECRET environment variable
}
```

### Environment Variables

You can also use the environment variables:

 * `$IX_API_AUTH`: Choose the authentication strategy. 
   Defaults to `legacy`. Can be set to `oauth2`.
 * `$IX_API_HOST`: The IX-API endpoint in the format: `https://<server>/api/v2`
 * `$IX_API_KEY`: The key provided by the exchange.
 * `$IX_API_SECRET`: Also provided by the exchange.
 * `$IX_API_OAUTH2_TOKEN_URL`: The OAuth2 token endpoint.
 * `$IX_API_OAUTH2_SCOPES`: A comma-separated list of OAuth2 scopes. (Optional)

### OAuth2

In order to use OAuth2 to retrieve an access token, you
need to provide the `oauth2_token_url` in addition to
the `api_key` and `api_secret`. Key and secret will be used
as `client_id` and `client_secret`.

The `auth` strategy must be set to `oauth2`.

```hcl
provider "ixapi" {
    auth = "oauth2"
    api = "http://localhost:8000/api/v2"
    api_key = "my_api_key"
    api_secret = "..."
    oauth2_token_url = "http://localhost:8000/auth/oauth2/token"
    oauth2_scopes = "ix-api"  # Optional
}
```


## Using The Provider

The following examples illustrate basic usage.

```hcl
# Querying: Show all facilities in the metro area FRA
data "ixapi_metro_area" "fra" {
  iata_code = "FRA"  # Resolve metro area by IATA code
}

data "ixapi_facilities" "fra" {
  metro_area = data.ixapi_metro_area.fra.id
}

output "facilities" {
  value = data.ixapi_facilities.fra.facilities
}

data "ixapi_account" "reseller" {
  external_ref = "demo_reseller"
}
```

### Using Resources

Create an account and add a contact.

```hcl
resource "ixapi_account" "hajnet" {
  managing_account = data.ixapi_account.reseller.id
  name = "Blåhaj Networks Inc."
  address {
    country = "DE"
    locality = "Berlin"
    postal_code = "11111"
    street_address = "Straßenweg 11"
  }
}

locals {
  reseller_id = data.ixapi_account.reseller.id
  customer_id = resource.ixapi_account.hajnet.id
}

resource "ixapi_contact" "hajnet_support" {
  managing_account = local.reseller_id 
  consuming_account = local.hajnet_id
  roles = ["noc", "implementation" ]
  email = "mail@example.com" 
  telephone = "+0 42 1234567890"
}

```



## Development

For development, you need to add the development build
of the terraform provider. You can do so, by adding
the following snippet to you `~/.terraformrc`:

```hcl
provider_installation {
    dev_overrides {
        "ix-api.net/ix-api/ixapi" = "/<full_path_to>/go/src/github.com/ix-api-net/terraform-provider-ixapi/bin"
    }

    direct {}
}
```

And then in the terraform file use:

```hcl
terraform {
    required_providers {
        ixapi = {
            source = "ix-api.net/ix-api/ixapi"
        }
    }
}
```

Use a `ix-api-sandbox-v2` as local API server.




