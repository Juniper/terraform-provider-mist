---
subcategory: "Release Notes"
page_title: "v0.2.13"
description: |-
    Release Notes for v0.2.13
---

# Release Notes for v0.2.13

**version**      : v0.2.13
**release date** : December 20th, 2024

## Improvements
* `mist_device_gateway_cluster`: Improve the creation resource behavior when one of both of the cluster nodes already belong to a cluster.  The provider will no more raise an error when the existing cluster in the Mist Cloud is matching the planned cluster (same primary node, same secondary node).
* `mist_org_inventory`: improve the deprecation message.

## Fixes
* [Issue 65](https://github.com/Juniper/terraform-provider-mist/issues/65): Fixing the `port_config.wan_source_nat` attribute in the `mist_device_gateway`, `mist_org_deviceprofile_gateway` and `mist_org_gatewaytemplate` resources 
* [Issue 66](https://github.com/Juniper/terraform-provider-mist/issues/66): Fixing `mist_org_wlan` resource following the v0.2.12 changes 
* [Issue 68](https://github.com/Juniper/terraform-provider-mist/issues/68): Fixing `tunnel_configs.auto_provision` attribute in `gateway` resources 
* Fixing issue when removing `rftemplate_id` from the `mist_site` resource