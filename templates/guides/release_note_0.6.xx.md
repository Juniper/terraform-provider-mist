---
subcategory: "Release Notes"
page_title: "v0.6.xx"
description: |-
    Release Notes for v0.6.xx
---

# Release Notes for v0.6.xx


## Release Notes for v0.6.0
**Release Date**: September 21st, 2025

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
    