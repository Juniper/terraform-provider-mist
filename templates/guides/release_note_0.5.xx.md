---
subcategory: "Release Notes"
page_title: "v0.5.xx"
description: |-
    Release Notes for v0.5.xx
---

# Release Notes for v0.5.xx


## Release Notes for v0.5.1
**Release Date**: August 22st, 2025 

### Fixes
* **[Issue 129](https://github.com/Juniper/terraform-provider-mist/issues/129):** Remove default values for the new `.mist_nac` attributes in `mist_org_wlan` and `mist_site_wlan` resources
* Remove the default values that were added to the `.port_usages` attributes in the `mist_device_switch`, `mist_org_networktemplate` and `mist_site_networktemplate` in v0.5.0


## Release Notes for v0.5.0
**Release Date**: August 21st, 2025 

### Fixes
* **[Issue 128](https://github.com/Juniper/terraform-provider-mist/issues/128):** Resolved a bug that prevented proper configuration of the `app_limit.wxtag_ids` attribute in `mist_org_wlan` resource


### General changes

#### Attributes added
- **`mist_device_ap` and `mist_deviceprofile_ap` resource**
 - `.airista` has been added
 - `.airista.enabled` has been added
 - `.airista.host` has been added
 - `.airista.port` has been added
 - `.port_config.mist_nac.acct_interim_interval` has been added
 - `.port_config.mist_nac.auth_servers_retries` has been added
 - `.port_config.mist_nac.auth_servers_timeout` has been added
 - `.port_config.mist_nac.coa_enabled` has been added
 - `.port_config.mist_nac.coa_port` has been added
 - `.port_config.mist_nac.fast_dot1x_timers` has been added
 - `.port_config.mist_nac.network` has been added
 - `.port_config.mist_nac.source_ip` has been added


- **`mist_device_gateway`, `mist_org_gateway_template` and `mist_deviceprofile_gateway` resource**
 - `.dhcp_config.ip_end6` has been renamed to `.dhcp_config.ip6_end`
 - `.dhcp_config.ip_start6` has been renamed to `.dhcp_config.ip6_start`
 - `.dhcp_config.servers6` has been renamed to `.dhcp_config.serversv6`


- **`mist_device_switch` resource**
 - `.local_port_config.storm_control.disable_port` has been added
 - `.ospf_config.export_policy` has been added
 - `.ospf_config.import_policy` has been added
 - `.port_config_overwrite` has been added
 - `.port_config_overwrite.description` has been added
 - `.port_config_overwrite.disabled` has been added
 - `.port_config_overwrite.duplex` has been added
 - `.port_config_overwrite.mac_limit` has been added
 - `.port_config_overwrite.poe_disabled` has been added
 - `.port_config_overwrite.port_network` has been added
 - `.port_config_overwrite.speed` has been added


- **`mist_device_switch`, `mist_org_switch_template` and `deviceprofile_switch` resource**
 - `.port_mirroring.output_ip_address` has been added
 - `.port_usage.storm_control.disable_port` has been added
 - `.remote_syslog.servers.server_name` has been added


- **`mist_org_nacrule` resources**
 - `.guest_auth_state` has been added


- **`mist_org_nactag` resources**
 - `.nacportal_id` has been added


- **`mist_org_wlan` and `mist_site_wlan` resources**
 - `.mist_nac.acct_interim_interval` has been added
 - `.mist_nac.auth_servers_retries` has been added
 - `.mist_nac.auth_servers_timeout` has been added
 - `.mist_nac.coa_enabled` has been added
 - `.mist_nac.coa_port` has been added
 - `.mist_nac.fast_dot1x_timers` has been added
 - `.mist_nac.network` has been added
 - `.mist_nac.source_ip` has been added


- **`mist_site_setting` resource**
 - `.auto_upgrade_esl` has been added
 - `.auto_upgrade_esl.allow_downgrade` has been added
 - `.auto_upgrade_esl.custom_versions` has been added
 - `.auto_upgrade_esl.day_of_the_week` has been added
 - `.auto_upgrade_esl.enabled` has been added
 - `.auto_upgrade_esl.time_of_day` has been added
 - `.auto_upgrade_esl.version` has been added
 - `.bgp_neighbor_updown_threshold` has been added
 - `.gateway_mgmt.probe_hostsv6` has been added
 - `.vpn_path_updown_threshold` has been added
 - `.vpn_peer_updown_threshold` has been added


### Resources default values changed

Changes have been applied to resources to reduce configuration drift when importing resources or saving changes from the Mist UI. These updates aim to align Terraform resource states with the Mist UI default values. However, some default values are dynamic and depend on other parameter values, making it currently impossible to completely eliminate configuration drift in certain scenarios.

~> **Warning** Some default values have been removed from the Terraform Provider resource schemas.  
These changes may lead to configuration drift if the affected attributes are not explicitly defined in your HCL configuration.  
Attributes without explicit definitions will default to `null`, but this will not alter the actual configuration in the Mist Cloud (the Mist Cloud will use the default value). To avoid discrepancies, ensure that all required attributes are explicitly set in your configuration.

*  **`mist_device_switch`, `mist_org_networktemplate` and `mist_site_networktemplate` resources**
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.port_usages.all_networks` | StaticBool(false) | N/A |
| `.port_usages.allow_multiple_supplicants` | StaticBool(false) | N/A |
| `.port_usages.bypass_auth_when_server_down` | StaticBool(false) | N/A |
| `.port_usages.bypass_auth_when_server_down_for_unknown_client` | StaticBool(false) | N/A |
| `.port_usages.disabled` | StaticBool(false) | N/A |
| `.port_usages.enable_mac_auth` | StaticBool(false) | N/A |
| `.port_usages.enable_qos` | StaticBool(false) | N/A |
| `.port_usages.inter_switch_link` | StaticBool(false) | N/A |
| `.port_usages.inter_switch_link` | StaticBool(false) | N/A |
| `.port_usages.mac_auth_protocol` | StaticString("eap-md5") | N/A |
| `.port_usages.mac_limit` | StaticInt64(0) | N/A |
| `.port_usages.persist_mac` | StaticBool(false) | N/A |
| `.port_usages.storm_control.no_broadcast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.no_multicast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.no_registered_multicast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.no_unknown_unicast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.percentage` | StaticInt64(80) | N/A |
| `.port_usages.stp_edge` | StaticBool(false) | N/A |
| `.port_usages.stp_no_root_port` | StaticBool(false) | N/A |
| `.port_usages.stp_p2p` | StaticBool(false) | N/A |
| `.port_usages.use_vstp` | StaticBool(false) | N/A |