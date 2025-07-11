---
subcategory: "Release Notes"
page_title: "v0.4.xx"
description: |-
    Release Notes for v0.4.xx
---

# Release Notes for v0.4.xx


## Release Notes for v0.4.3
**Release Date**: July 11th, 2025 
* **[Issue 118](https://github.com/Juniper/terraform-provider-mist/issues/118):** Fix the `mist_device_switch.ip_config.dns` validators requiring the DNS to be set when the `ip_config.type` is set to `static`. The DNS is not optional in this case, so the validator has been updated to allow it to be empty.


## Release Notes for v0.4.2
**Release Date**: July 1st, 2025 
* **[Issue 117](https://github.com/Juniper/terraform-provider-mist/issues/117):** Resolved a bug that prevented the `mist_org_deviceprofile_gateway.bpg_config.neighbor_as` attribute from being properly configured.


## Release Notes for v0.4.1
**Release Date**: June 27th, 2025 

### Fixes
* Resolved a bug that prevented proper configuration of the `mist_device_gateway.bpg_config.neighbor_as`, `mist_org_deviceprofile_gateway.bpg_config.neighbor_as` and `mist_org_gatewaytemplate.bpg_config.neighbor_as` resources attributes.
* Remove unexpected `mist_site_networktemplate.marvis` and `mist_site_networktemplate.sle_thresholds` attributes.

### Other Changes
* Minor code clean up and typo corrections.

## Release Notes for v0.4.0
**Release Date**: June 27th, 2025 

### Fixes
* **[Issue 112](https://github.com/Juniper/terraform-provider-mist/issues/112):** Resolved a bug that prevented proper configuration of the `mist_org_wlan.wxtag_ids` and `mist_org_wlan.wxtag_ids` resources attributes.


### General changes

#### New validator
* Adding new validator to `mist_org_wlan.schedule.hours` `mist_site_wlan.schedule.hours` to require at least one day of the week to be set.

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
