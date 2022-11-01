
**This repository is work in progress and not an official release.**

# IX-API Terraform Provider

This Terraform provider is using the [ix-api](https://ix-api.net)
for configuring and provisisioning IXP services.


## Requirements
 * Terraform >= 1.0
 * Go >= 1.17


## Configure the provider


```hcl
provider "ixapi" {
    api = "http://localhost:8000/api/v2"
    api_key = "my_api_key"
    # api_secret = "" # Use $IX_API_SECRET environment variable
}
```

You can also use the environment variables:

 * `$IX_API_HOST`
 * `$IX_API_KEY`
 * `$IX_API_SECRET`



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




