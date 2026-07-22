# Virtual PNI with a Foreign Port Example

This example demonstrates a Virtual PNI between your port and a port owned by a foreign partner over a point-to-point virtual circuit.

## What This Example Creates

1. **Discoverable partner connection** - the partner marks a connection `discoverable`, publishing a p2p product offering under their service-provider name.
2. **Network service** - your side creates the `ixapi_network_service_p2p_vc` against that offering, with the partner as the joining member.
3. **Two configs** - your side and the partner each attach their port with an `ixapi_network_service_config_p2p_vc` on the same network service.

Both sides are shown with two provider configurations for completeness; in practice each side is usually managed by its own account and Terraform configuration.

## Requirements

- terraform-provider-ixapi **>= 1.0.3**
- The DE-CIX Cloud Router extension enabled on the provider (`extension_de_cix_cloud_router_enabled = true`), on both sides.
