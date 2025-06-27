---
subcategory: "Release Notes"
page_title: "v0.4.xx"
description: |-
    Release Notes for v0.4.xx
---

# Release Notes for v0.4.xx


## Release Notes for v0.4.0
**Release Date**: June 27th, 2025 

### Fixes
* **[Issue 112](https://github.com/Juniper/terraform-provider-mist/issues/112):** Resolved a bug that prevented proper configuration of the `mist_org_wlan.wxtag_ids` and `mist_org_wlan.wxtag_ids` resources attributes.


### General changes

#### Attributes removed
* `mist_org_setting.mxedge_fips_enabled` has been removed to match the API structure (use `mist_org_setting.mxedge_mgmt.fips_enabled` instead)

#### Import re-enabled
* The `import` function for the `mist_org_sso_role` resource has been re-enabled

#### Attributes added
- **`mist_device_gateway`, `mist_org_gateway_template` and `mist_deviceprofile_gateway` resource**
 - `.port_config[]. wan_disable_speedtest` has been added
 - `.tunnel_configs.service_connection` has been added
 - `.tunnel_provider_options.prisma` has been added

- **`mist_device_switch`, `mist_org_switch_template` and `deviceprofile_switch` resource**
- `.acl_tags.ether_types` has been added
- `.acl_tags.port_usage` has been added
- `.port_usage.port_network` has been added
- `.port_usage.community_vlan_id` has been added
- `.remote_syslog.cacerts` has been added
- `.remote_syslog.files.enable_tls` has been added
- `.snmp_config.engine_id_type` has been added
- `.switch_mgmt.protect_re.hit_count` has been added
- `.switch_mgmt.remove_existing_configs` has been added
- `.vrrp_config.groups.accept_data` has been added (mist_device_switch only)
- `.vrrp_config.groups.preempt` has been added (mist_device_switch only)

- **`mist_org_evpn_topology` and `mist_site_evpn_topology` resources**
 - `.evpn_options.enable_inband_ztp` has been added

- **`mist_org_setting` resource**
- `.marvis` has been added
- `.ssr` has been added
- `.switch` has been added
- `.synthetic_test. aggressiveness` has been added
- `.synthetic_test.custom_probes` has been added
- `.synthetic_test.lan_networks` has been added


- **`mist_org_wlan` and `mist_site_wlan` resources**
 - `.portal.smsglobal_api_key` has been added
 - `.portal.smsglobal_api_secret` has been added


- **`mist_site_setting` resource**
 -  `.default_port_usage` has been added
 -  `.gateway_mgmt.protect_re.hit_count` has been added
 -  `.marvis` has been added
 -  `.sle_thresholds` has been added
 -  `.ssr.conductor_token` has been added
 -  `.synthetic_test. aggressiveness` has been added
 -  `.custom_probes.custom_probes` has been added

- **`mist_site_webhook` resource**
 -  `.assetfilter_ids` has been added
