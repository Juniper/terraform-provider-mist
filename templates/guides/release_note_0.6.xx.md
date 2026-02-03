---
subcategory: "Release Notes"
page_title: "v0.6.xx"
description: |-
    Release Notes for v0.6.xx
---

# Release Notes for v0.6.xx

## Release Notes for v0.6.4
**Release Date**: Febuary 3rd, 2026

### Fixes

### General changes

#### Attributes added

- **`mist_device_switch` resource**:
    - default_port_usage

- **`mist_org_networktemplate` resource**:
    - switch_matching.rules[*].default_port_usage

- **`mist_site_networktemplate` resource**:
    - switch_matching.rules[*].default_port_usage

#### Attributes removed

- **`mist_site_setting` resource**:
    - default_port_usage

## Release Notes for v0.6.3
**Release Date**: December 19th, 2025

### Fixes
* **Routing Policies**: Fixed normalization issues in gateway routing policies for `mist_device_gateway`, `mist_org_deviceprofile_gateway`, and `mist_org_gatewaytemplate` resources to prevent plan/apply mismatches.

### General changes

#### Attributes added

- **`mist_device_switch` resource**:
    - routing_policies

- **`mist_org_networktemplate` resource**:
    - routing_policies

- **`mist_site_networktemplate` resource**:
    - routing_policies

## Release Notes for v0.6.2
**Release Date**: December 18th, 2025

### Breaking Changes

- **`mist_device_ap` resource**:
    - `port_config.vlan_ids`: Changed from `list(number)` to `string` (comma-separated list)

- **`mist_org_deviceprofile_ap` resource**:
    - `port_config.vlan_ids`: Changed from `list(number)` to `string` (comma-separated list)

### General changes

#### Attributes added

- **`mist_device_switch` resource**:
    - bgp_config
    - port_config.networks

- **`mist_org_networktemplate` resource**:
    - bgp_config
    - switch_matching.rules.port_config.networks

- **`mist_site_networktemplate` resource**:
    - switch_matching.rules.port_config.networks

## Release Notes for v0.6.1
**Release Date**: December 3rd, 2025

### Fixes
* **[Issue 152](https://github.com/Juniper/terraform-provider-mist/issues/152):** Added `orgsites` option for sso scope role.
* **[Issue 160](https://github.com/Juniper/terraform-provider-mist/issues/160):** Added `lldp_system_description` option for `src` under switch `port_usages.rules`.
* **[Issue 161](https://github.com/Juniper/terraform-provider-mist/issues/161):** Handle empty `networks` attribute for `dhcp_snooping`.

### General changes

#### Attributes updated
- **`mist_device_ap` and `mist_deviceprofile_ap` resource**:
    - Updated `radio_config.ant_mode` to `radio_config.antenna_select`.

#### Attributes added
- **`mist_device_ap` and `mist_deviceprofile_ap` resource**:
    - radio_config.band_5.antenna_beam_pattern
    - radio_config.band_5_on_24_radio.antenna_beam_pattern
    - radio_config.band_6.antenna_beam_pattern
    - radio_config.rrm_managed

- **`mist_device_gateway`, `mist_deviceprofile_gateway` and `mist_org_gatewaytemplate` resources**:
    - service_policies.skyatp
    - service_policies.syslog

- **`mist_device_switch`, `mist_org_networktemplate` and `mist_site_networktemplate` resource**:
    - port_usage.bypass_auth_when_server_down_for_voip
    - port_usage.poe_priority

- **`mist_org_setting` resource**:
    - marvis.auto_operations.ap_insufficient_capacity
    - marvis.auto_operations.ap_loop
    - marvis.auto_operations.ap_non_compliant
    - marvis.auto_operations.gateway_non_compliant
    - marvis.auto_operations.switch_misconfigured_port
    - marvis.auto_operations.switch_port_stuck
    - ssr.proxy.disabled

- **`mist_site_setting` resource**:
    - marvis.auto_operations.ap_insufficient_capacity
    - marvis.auto_operations.ap_loop
    - marvis.auto_operations.ap_non_compliant
    - marvis.auto_operations.gateway_non_compliant
    - marvis.auto_operations.switch_misconfigured_port
    - marvis.auto_operations.switch_port_stuck
    - proxy.disabled
    - ssr.proxy.disabled
   
## Release Notes for v0.6.0
**Release Date**: October 22nd, 2025

### New Features
* Adding support for APAC 02 Mist Cloud region.


### General changes

#### Attributes added
- **`mist_device_ap` and `mist_deviceprofile_ap` resource**:
    - `radio_config.ant_mode`

- **`mist_device_gateway`, `mist_deviceprofile_gateway` and `mist_org_gatewaytemplate` resources**:
    - dhcp_config.ip6
    - port_config.wan_ext_ip6
    - port_config.wan_source_nat.nat6_pool
    - url_filtering_deny_msg

- **`mist_device_gateway` resource**:
    - gateway_mgmt.config_revert_timer

- ** `mist_org_networktemplate` and `mist_site_networktemplate` resources**:
    - port_usage.stp_disable
    - port_usage.stp_required
    - switch_matching.rules.stp_config

- **`mist_device_switch` resource**:
    - dhcpd_config.fixed_binding.ip6
    - port_usage.stp_disable
    - port_usage.stp_required

- **`mist_org_setting` resource**:
    - juniper_srx.auto_upgrade
    - ssr.proxy
    - ssr.auto_upgrade
    - ui_not_tracking
    - vpn_options.enable_ipv6

- **`mist_site_setting` resource**:
    - juniper_srx.auto_upgrade
    - ssr.proxy
    - ssr.auto_upgrade
    