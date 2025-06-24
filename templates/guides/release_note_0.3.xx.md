---
subcategory: "Release Notes"
page_title: "v0.3.xx"
description: |-
    Release Notes for v0.3.xx
---

# Release Notes for v0.3.xx


## Release Notes for v0.3.6
**Release Date**: June 24th, 2025

### Fixes
* **[Issue 108](https://github.com/Juniper/terraform-provider-mist/issues/108):** Resolved a bug that prevented proper configuration of the `mist_device_ap.radio_config` and `mist_deviceprofile_ap.radio_config` resources.


### New Attributes
* **`mist_device_ap` and `mist_org_deviceprofile_ap` Resources**  
  - Added `.port_config` nested object
  - Added `.radio_config.full_automatic_rrm`


## Release Notes for v0.3.5
**Release Date**: May 23th, 2025
### Fixes
* **[Issue 105](https://github.com/Juniper/terraform-provider-mist/issues/105):** Resolved a bug that prevented proper configuration of the `uplink_port_config.keep_wlans_up_if_down` attribute in the `mist_site_setting` resource.

## Release Notes for v0.3.4
**Release Date**: May 16th, 2025
### Fixes
* **[Issue 104](https://github.com/Juniper/terraform-provider-mist/issues/104):** Resolved a bug that prevented proper configuration of the `local_port_config` attribute in the `mist_device_switch` resource.


## Release Notes for v0.3.3
**Release Date**: May 14th, 2025
### Fixes
* **[Issue 103](https://github.com/Juniper/terraform-provider-mist/issues/103):** Resolved a bug that prevented proper configuration of the `radius_config.auth_servers` attribute in the `mist_device_switch` resource.


## Release Notes for v0.3.2
**Release Date**: April 14th, 2025
### Fixes
* **[Issue 97](https://github.com/Juniper/terraform-provider-mist/issues/97):** Changed the `mist_org_service.addresses` validator to be less restrictive and match the Mist UI behavior

* **[Issue 98](https://github.com/Juniper/terraform-provider-mist/issues/98):** Corrected a typo in the `.tunnel_configs.networks` validation for the following resources:  
  - `mist_device_gateway`  
  - `mist_org_deviceprofile_gateway`  
  - `mist_org_gatewaytemplate`  


* **[Issue 99](https://github.com/Juniper/terraform-provider-mist/issues/99):** Adding the missing attributes to the `mist_org_networks` datasource

* **[Issue 100](https://github.com/Juniper/terraform-provider-mist/issues/100):** Fix validators configuration causing a Provider error when the attribute `tunnel_provider_options` was configured in the `mist_device_gateway`, `mist_org_deviceprofile_gateway` and `mist_org_gatewaytemplate` datasources

* Fix issue with `mist_org_nacrule.matching` and `mist_org_nacrule.not_matching` default values

* Fixed a typo in the `.tunnel_configs.provider` enumeration for the following resources:  
  - `mist_device_gateway`  
  - `mist_org_deviceprofile_gateway`  
  - `mist_org_gatewaytemplate`  

* Resolved conversion issues affecting the following resources:  
  - `mist_device_gateway`  
  - `mist_org_deviceprofile_gateway`  
  - `mist_org_gatewaytemplate`  

* Resolved issue when using Mist GLOBAL 05 Cloud (manage.gc4.mist.com)

* Resolved an issue with the `mist_site_setting` resource triggered when all the `vars` are removed from the plan

* Fixed a provider crash occurring during resource deletion when the API returned a nil response.


### Breaking Changes

#### Changes in Attribute Types


~> **Impact** Review and update your Terraform configurations to align with the new attribute types. This ensures compatibility with the latest API behavior and prevents potential runtime issues.

- **`mist_org_wlan` and `mist_site_wlan` Resources**  
  - Updated the type of `.client_limit_down` from `int64` to `string`.
  - Updated the type of `.client_limit_up` from `int64` to `string`.
  - Updated the type of `.wlan_limit_down` from `int64` to `string`.
  - Updated the type of `.wlan_limit_up` from `int64` to `string`.
  - Updated the type of `.acct_servers.port` from `int64` to `string`.
  - Updated the type of `.auth_servers.port` from `int64` to `string`.
  - Updated the type of `.coa_servers.port` from `int64` to `string`.
  - Updated the type of `.radsec.idle_timeout` from `int64` to `string`.


- **`mist_device_switch`, `mist_org_networktemplate` and `mist_site_networktemplate` Resources**  
  - Updated the type of `.acct_servers.port` from `int64` to `string`.
  - Updated the type of `.auth_servers.port` from `int64` to `string`.
  - Updated the type of `.coa_servers.port` from `int64` to `string`.


### Resources default values changed

#### Changes to Reduce Configuration Drift

Changes have been applied to resources to reduce configuration drift when importing resources or saving changes from the Mist UI. These updates aim to align Terraform resource states with the Mist UI default values. However, some default values are dynamic and depend on other parameter values, making it currently impossible to completely eliminate configuration drift in certain scenarios.

~> **Warning** Some default values have been removed from the Terraform Provider resource schemas.  
These changes may lead to configuration drift if the affected attributes are not explicitly defined in your HCL configuration.  
Attributes without explicit definitions will default to `null`, but this will not alter the actual configuration in the Mist Cloud (the Mist Cloud will use the default value). To avoid discrepancies, ensure that all required attributes are explicitly set in your configuration.


List of the default value changed:
*  `mist_org_wlan` and `mist_site_wlan`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.auth_servers_retries` | StaticInt64(2) | N/A |
| `.auth_servers_timeout` | StaticInt64(5) | N/A |
| `.band_steer_force_band5` | StaticBool(false) | N/A |
| `.block_blacklist_clients` | StaticBool(false) | N/A |
| `.bonjour.enabled` | StaticBool(false) | N/A |
| `.cisco_cwa.allowed_hostnames` | N/A | [] |
| `.cisco_cwa.allowed_subnets` | N/A | [] |
| `.cisco_cwa.blocked_subnets` | N/A | [] |
| `.client_limit_down_enabled` | StaticBool(false) | N/A |
| `.client_limit_up_enabled` | StaticBool(false) | N/A |
| `.coa_servers` | [] | N/A |
| `.disable_11be` | StaticBool(false) | N/A |
| `.disable_ht_vht_rates` | StaticBool(false) | N/A |
| `.disable_v1_roam_notify` | StaticBool(false) | N/A |
| `.disable_v2_roam_notify` | StaticBool(false) | N/A |
| `.enable_local_keycaching` | StaticBool(false) | N/A |
| `.enable_wireless_bridging` | StaticBool(false) | N/A |
| `.enable_wireless_bridging_dhcp_tracking` | StaticBool(false) | N/A |
| `.fast_dot1x_timers` | StaticBool(false) | N/A |
| `.isolation` | StaticBool(false) | N/A |
| `.l2_isolation` | StaticBool(false) | N/A |
| `.legacy_overds` | StaticBool(false) | N/A |
| `.max_num_clients` | StaticInt64(0) | N/A |
| `.radsec.server_name` | N/A | StaticString("") | 
| `.radsec.use_site_mxedge` | StaticBool(false) | N/A |
| `.roam_mode` | StaticBool(false) | N/A |
| `.schedule.hours.mon` | StaticString(") | N/A |
| `.schedule.hours.tue` | StaticString(") | N/A |
| `.schedule.hours.wed` | StaticString(") | N/A |
| `.schedule.hours.thu` | StaticString(") | N/A |
| `.schedule.hours.fri` | StaticString(") | N/A |
| `.schedule.hours.sat` | StaticString(") | N/A |
| `.schedule.hours.sun` | StaticString(") | N/A |
| `.sle_excluded` | StaticBool(false) | N/A |
| `.wlan_limit_down_enabled` | StaticBool(false) | N/A |
| `.wlan_limit_up_enabled` | StaticBool(false) | N/A |



*  `mist_org_wlantemplate`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.applies.site_ids` | N/A | [] |
| `.applies.sitegroup_ids` | N/A | [] |
| `.exceptions.site_ids` | N/A | [] |
| `.exceptions.sitegroup_ids` | N/A | [] |


* `mist_site_networktemplate`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.switch_mgmt.use_mxedge_proxy` | N/A | Nested Object |


* `mist_site_setting`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.analytic` | N/A | Nested Object |
| `.auto_upgrade` | N/A | Nested Object |
| `.engagement` | N/A | Nested Object |
| `.gateway_mgmt` | N/A | Nested Object |
| `.gateway_mgmt.admin_sshkeys` | N/A | [] |
| `.gateway_mgmt.probe_hosts` | N/A | [] |
| `.led` | N/A | Nested Object |
| `.occupancy` | N/A | Nested Object |
| `.rogue` | N/A | Nested Object |
| `.rogue.allowed_vlan_ids` | N/A | [] |
| `.rtsa` | N/A | Nested Object |
| `.ssr` | N/A | Nested Object |
| `.synthetic_test` | N/A | Nested Object |
| `.uplink_port_config` | N/A | Nested Object |
| `.wids` | N/A | Nested Object |
| `.wifi` | N/A | Nested Object |
| `.zone_occupancy_alert` | N/A | Nested Object |
| `.zone_occupancy_alert.email_notifiers` | N/A | [] |

--- 

## Release Notes for v0.3.1
**Release Date**: March 31th, 2025 
### Breaking Changes

#### Changes in Allowed Attribute Values
- **`mist_org_service` Resource**  
  - Corrected the `.failover_policy` enum values from `[revertable, non_revertable]` to `[revertible, non_revertible]`.  

~> **Impact** Ensure your Terraform configurations are updated to use the corrected enum values to avoid validation errors.

#### Changes in Attribute Types
To improve compatibility with API type variations and support the use of `{{variables}}` in attribute values, the following type changes have been made:

- **`mist_device_switch`, `mist_org_networktemplate`, and `mist_site_networktemplate` Resources**  
  - Updated the type of `.port_usages.mac_limit` from `int64` to `string`.  
  - Updated the type of `.port_usages.mtu` from `int64` to `string`.

- **`mist_device_gateway`, `mist_org_deviceprofile_gateway`, and `mist_org_gatewaytemplate` Resources**  
  - Updated the type of `.port_config.reth_idx` from `int64` to `string`.

~> **Impact** Review and update your Terraform configurations to align with the new attribute types. This ensures compatibility with the latest API behavior and prevents potential runtime issues.


### General changes

#### Attributes removed
* `mist_org_setting.mxedge_fips_enabled` has been removed to match the API structure (use `mist_org_setting.mxedge_mgmt.fips_enabled` instead)

#### Import function disabled
* temporary removing the `import` function from the `mist_org_sso_role` resource. The import function will be added back in a later version

#### Attributes added
- **`mist_device_switch`, `mist_org_networktemplate` and `mist_site_networktemplate` resources**
 - `.radius_config.acct_immediate_update` has been added
 - `.radius_config.auth_server_selection` has been added, default is `StaticString("ordered")`
 - `.radius_config.coa_enabled` has been added, default is `StaticBool(true)`
 - `.radius_config.coa_port` has been added
 - `.radius_config.fast_dot1x_timers` has been added, default is `StaticBool(false)`

- **`mist_org_deviceprofile_ap` resource**
 -  `.aeroscout.port` has been added


### Resources default values changed

#### Reducing Configuration Drift for Resource IDs

To address the configuration drift caused by the behavior where the resource `id` was marked as "known after apply," improvements have been made to ensure that the `id` attribute uses the state value if unknown. This change reduces unnecessary plan differences and aligns the resource state more accurately with the Mist Cloud API. Users should experience fewer discrepancies when managing resources through Terraform.

#### Changes to Reduce Configuration Drift

Changes have been applied to resources to reduce configuration drift when importing resources or saving changes from the Mist UI. These updates aim to align Terraform resource states with the Mist UI default values. However, some default values are dynamic and depend on other parameter values, making it currently impossible to completely eliminate configuration drift in certain scenarios.

~> **Warning** Some default values have been removed from the Terraform Provider resource schemas.  
These changes may lead to configuration drift if the affected attributes are not explicitly defined in your HCL configuration.  
Attributes without explicit definitions will default to `null`, but this will not alter the actual configuration in the Mist Cloud (the Mist Cloud will use the default value). To avoid discrepancies, ensure that all required attributes are explicitly set in your configuration.

List of the default value changed:
*  `mist_org_alarmtemplate`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.rules.delivery.additional_emails` | N/A | [] |

*  `mist_device_ap` and `mist_org_deviceprofile_ap`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.aeroscout.host` | N/A |StaticString("") |
| `.aeroscout.locate_connected` | StaticBool(true) | StaticBool(false) |
| `.ble_config.beacon_enabled` | StaticBool(false) | StaticBool(true) |
| `.ble_config.beacon_rate` | StaticInt64(0) | N/A |
| `.ble_config.beacon_rate_mode` | StaticString("default") | N/A |
| `.ble_config.custom_ble_packet_enabled` | StaticBool(false) | N/A |
| `.ble_config.custom_ble_packet_frame` | StaticString("") | N/A |
| `.ble_config.custom_ble_packet_freq_msec` | StaticInt64(0) | N/A |
| `.ble_config.eddystone_uid_adv_power` | StaticInt64(0) | N/A |
| `.ble_config.eddystone_uid_beams` | StaticString("") | N/A |
| `.ble_config.eddystone_uid_enabled` | StaticBool(false) | N/A |
| `.ble_config.eddystone_uid_freq_msec` | StaticInt64(0) | N/A |
| `.ble_config.eddystone_uid_instance` | StaticString("") | N/A |
| `.ble_config.eddystone_uid_namespace` | StaticString("") | N/A |
| `.ble_config.eddystone_url_adv_power` | StaticInt64(0) | N/A |
| `.ble_config.eddystone_url_beams` | StaticString("") | N/A |
| `.ble_config.eddystone_url_enabled` | StaticBool(false) | N/A |
| `.ble_config.eddystone_url_freq_msec` | StaticInt64(0) | N/A |
| `.ble_config.eddystone_url_url` | StaticString("") | N/A |
| `.ble_config.ibeacon_adv_power` | StaticInt64(0) | N/A |
| `.ble_config.ibeacon_beams` | StaticString("") | N/A |
| `.ble_config.ibeacon_enabled` | StaticBool(false) | N/A |
| `.ble_config.ibeacon_freq_msec` | StaticInt64(0) | N/A |
| `.ble_config.ibeacon_major` | StaticInt64(0) | N/A |
| `.ble_config.ibeacon_minor` | StaticInt64(0) | N/A |
| `.ble_config.ibeacon_uuid` | StaticString("") | N/A |
| `.ble_config.power` | StaticInt64(0) | N/A |
| `.ble_config.power_mode` | StaticString("default") | N/A |
| `.esl_config.host` | N/A | StaticString("") |
| `.esl_config.type` | N/A | StaticString("imagotag") |
| `.ip_config.mtu` | N/A | StaticInt64(0) |
| `.ip_config.type6` | StaticString("disabled") | N/A |
| `.ip_config.vlan_id` | StaticInt64(1) | N/A |
| `.radio_config.allow_rrm_disable` | StaticBool(false) | N/A |
| `.radio_config.antenna_mode` | StaticString("default") | N/A |
| `.radio_config.band_5.bandwidth` | N/A | StaticInt64(40) |
| `.radio_config.band_5_on_24_radio.bandwidth` | N/A | StaticInt64(40) |
| `.radio_config.indoor_use` | StaticBool(false) | N/A |
| `.radio_config.keep_wlans_up_if_down` | StaticBool(false) | N/A |
| `.usb_config.cacert` | N/A | StaticString("") |
| `.usb_config.host` | N/A | StaticString("") |
| `.usb_config.port` | StaticInt64(0) | N/A |
| `.usb_config.vlan_id` | StaticInt64(1) | N/A |

*  `mist_device_switch`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.dns_servers` | N/A | [] |
| `.dns_suffix` | N/A | [] |
| `.extra_routes.discard` | StaticBool(false) | N/A |
| `.extra_routes.no_resolve` | StaticBool(false) | N/A |
| `.extra_routes6.discard` | StaticBool(false) | N/A |
| `.extra_routes6.no_resolve` | StaticBool(false) | N/A |
| `.networks.isolation` | StaticBool(false) | N/A |
| `.ntp_servers` | N/A | [] |
| `.oob_ip_config.use_mgmt_vrf_for_host_out` | StaticBool(false) | N/A |
| `.port_config.aggregated` | StaticBool(false) | N/A |
| `.port_config.disable_autoneg` | StaticBool(false) | N/A |
| `.port_config.duplex` | StaticString("auto") | N/A |
| `.port_config.mtu` | StaticInt64(1514) | N/A |
| `.port_config.poe_disabled` | StaticBool(false) | N/A |
| `.port_config.speed` | StaticString("auto") | N/A |
| `.port_usages.allow_multiple_supplicants` | StaticBool(false) | N/A |
| `.port_usages.bypass_auth_when_server_down` | StaticBool(false) | N/A |
| `.port_usages.bypass_auth_when_server_down_for_unknown_client` | StaticBool(false) | N/A |
| `.port_usages.description` | N/A | StaticString("") |
| `.port_usages.enable_mac_auth` | StaticBool(false) | N/A |
| `.port_usages.inter_isolation_network_link` | StaticBool(false) | N/A |
| `.port_usages.inter_switch_link` | StaticBool(false) | N/A |
| `.port_usages.mac_auth_protocol` | StaticString("eap-md5") | N/A |
| `.port_usages.mac_limit` | N/A |StaticString("0") |
| `.port_usages.reauth_interval` | StaticString("3600") | N/A |
| `.port_usages.reset_default_when` | StaticString("link_down") | N/A |
| `.port_usages.storm_control.no_broadcast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.no_multicast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.no_registered_multicast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.no_unknown_unicast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.percentage` | StaticInt64(80) | N/A |
| `.port_usages.stp_no_root_port` | StaticBool(false) | N/A |
| `.port_usages.stp_p2p` | StaticBool(false) | N/A |
| `.port_usages.use_vstp` | StaticBool(false) | N/A |
| `.radius_config.auth_servers.require_message_authenticator` | StaticBool(false) | N/A |
| `.radius_config.auth_servers.coa_port` | N/A | StaticString("") |
| `.remote_syslog.send_to_all_servers` | StaticBool(false) | N/A |
| `.remote_syslog.servers.port` | StaticInt64(514) | N/A |
| `.snmp_config.network` | StaticString("default") | N/A |
| `.switch_mgmt.ap_affinity_threshold` | StaticInt64(10) | N/A |
| `.switch_mgmt.dhcp_option_fqdn` | StaticBool(false) | N/A |
| `.switch_mgmt.mxedge_proxy_port` | StaticInt64(2200) | N/A |

*  `mist_org_avprofile`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.mime_whitelist` | N/A | [] |
| `.url_whitelist` | N/A | [] |

*  `mist_org_nacrule`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.apply_tags` | N/A | [] |

*  `mist_org_nactag`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.allow_usermac_override` | StaticBool(false) | N/A |
| `.match_all` | StaticBool(false) | N/A |

*  `mist_org_network`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.routed_for_networks` | N/A | [] |
| `.vpn_access.source_nat` | N/A | {} |
| `.vpn_access.static_nat` | N/A | {} |

*  `mist_org_networktemplate` and `mist_site_networktemplate`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.dns_servers` | N/A | [] |
| `.dns_suffix` | N/A | [] |
| `.extra_routes.discard` | StaticBool(false) | N/A |
| `.extra_routes.no_resolve` | StaticBool(false) | N/A |
| `.extra_routes6.discard` | StaticBool(false) | N/A |
| `.extra_routes6.no_resolve` | StaticBool(false) | N/A |
| `.networks.isolation` | StaticBool(false) | N/A |
| `.ntp_servers` | N/A | []] |
| `.port_usages.allow_multiple_supplicants` | StaticBool(false) | N/A |
| `.port_usages.bypass_auth_when_server_down` | StaticBool(false) | N/A |
| `.port_usages.bypass_auth_when_server_down_for_unknown_client` | StaticBool(false) | N/A |
| `.port_usages.description` | N/A | StaticString("") |
| `.port_usages.enable_mac_auth` | StaticBool(false) | N/A |
| `.port_usages.inter_isolation_network_link` | StaticBool(false) | N/A |
| `.port_usages.inter_switch_link` | StaticBool(false) | N/A |
| `.port_usages.mac_auth_protocol` | StaticString("eap-md5") | N/A |
| `.port_usages.mac_limit` | N/A |StaticString("0") |
| `.port_usages.reauth_interval` | StaticString("3600") | N/A |
| `.port_usages.reset_default_when` | StaticString("link_down") | N/A |
| `.port_usages.storm_control.no_broadcast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.no_multicast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.no_registered_multicast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.no_unknown_unicast` | StaticBool(false) | N/A |
| `.port_usages.storm_control.percentage` | StaticInt64(80) | N/A |
| `.port_usages.stp_no_root_port` | StaticBool(false) | N/A |
| `.port_usages.stp_p2p` | StaticBool(false) | N/A |
| `.port_usages.use_vstp` | StaticBool(false) | N/A |
| `.radius_config.auth_servers.require_message_authenticator` | StaticBool(false) | N/A |
| `.radius_config.auth_servers.coa_port` | N/A | StaticString("") |
| `.remote_syslog.send_to_all_servers` | StaticBool(false) | N/A |
| `.remote_syslog.servers.port` | StaticInt64(514) | N/A |
| `.snmp_config.network` | StaticString("default") | N/A |
| `.ssh_keys` | N/A | [] |
| `.switch_matching.rules.oob_ip_config.use_mgmt_vrf_for_host_out` | StaticBool(false) | N/A |
| `.switch_matching.rules.port_config.aggregated` | StaticBool(false) | N/A |
| `.switch_matching.rules.port_config.disable_autoneg` | StaticBool(false) | N/A |
| `.switch_matching.rules.port_config.duplex` | StaticString("auto") | N/A |
| `.switch_matching.rules.port_config.mtu` | StaticInt64(1514) | N/A |
| `.switch_matching.rules.port_config.poe_disabled` | StaticBool(false) | N/A |
| `.switch_matching.rules.port_config.speed` | StaticString("auto") | N/A |
| `.switch_mgmt.ap_affinity_threshold` | StaticInt64(10) | N/A |
| `.switch_mgmt.dhcp_option_fqdn` | StaticBool(false) | N/A |
| `.switch_mgmt.mxedge_proxy_port` | StaticInt64(2200) | N/A |

* `mist_org_psk` and `mist_site_psk`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.max_usage` | StaticInt64(0) | N/A |

* `mist_org_rftemplate`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.band_24.antenna_mode` | StaticString("default") | N/A |
| `.band_24.power` | StaticInt64(0) | N/A |
| `.band_5.antenna_mode` | StaticString("default") | N/A |
| `.band_5.bandwidth` | N/A | StaticInt64(40) |
| `.band_5_on_24_radio.antenna_mode` | StaticString("default") | N/A |
| `.band_5_on_24_radio.bandwidth` | N/A | StaticInt64(40) |
| `.band_5_on_24_radio.power` | StaticInt64(0) | N/A |
| `.band_5.power` | StaticInt64(0) | N/A |
| `.band_6.antenna_mode` | StaticString("default") | N/A |
| `.band_6.power` | StaticInt64(0) | N/A |
| `.model_specific.ant_gain_24` | N/A | StaticInt64(0) | 
| `.model_specific.ant_gain_5` | N/A | StaticInt64(0) | 
| `.model_specific.ant_gain_6` | N/A | StaticInt64(0) | 
| `.model_specific.band_24.antenna_mode` | StaticString("default") | N/A |
| `.model_specific.band_24.power` | StaticInt64(0) | N/A |
| `.model_specific.band_5.antenna_mode` | StaticString("default") | N/A |
| `.model_specific.band_5.bandwidth` | N/A | StaticInt64(40) |
| `.model_specific.band_5.power` | StaticInt64(0) | N/A |
| `.model_specific.band_5_on_24_radio.antenna_mode` | StaticString("default") | N/A |
| `.model_specific.band_5_on_24_radio.bandwidth` | N/A | StaticInt64(40) |
| `.model_specific.band_5_on_24_radio.power` | StaticInt64(0) | N/A |
| `.model_specific.band_6.antenna_mode` | StaticString("default") | N/A |
| `.model_specific.band_6.power` | StaticInt64(0) | N/A |


* `mist_org_service`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.client_limit_down` | StaticInt64(0) | N/A |
| `.client_limit_up` | StaticInt64(0) | N/A |
| `.failover_policy` | StaticString("revertible") | N/A |
| `.service_limit_down` | StaticInt64(0) | N/A |
| `.service_limit_up` | StaticInt64(0) | N/A |
| `.sle_enabled` | StaticBool(false) | N/A |
| `.ssr_relaxed_tcp_state_enforcement` | StaticBool(false) | N/A |
| `.traffic_class` | StaticString("best_effort") | N/A |

* `mist_org_servicepolicy`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.aamw.profile` | StaticString("standard") | N/A |
| `.ewf.profile` | StaticString("strict") | N/A |
| `.idp.profile` | StaticString("strict") | N/A |

*  `mist_org_setting`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.ap_updown_threshold` | StaticInt64(0) | N/A |
| `.cloudshark.apitoken` | N/A |StaticString("") |
| `.cloudshark.url` | N/A |StaticString("") |
| `.device_updown_threshold` | StaticInt64(0) | N/A |
| `.disable_pcap` | StaticBool(false) | N/A |
| `.disable_remote_shell` | StaticBool(false) | N/A |
| `.gateway_updown_threshold` | StaticInt64(0) | N/A |
| `.mist_nac.disable_rsae_algorithms` | StaticBool(false) | N/A |
| `.mist_nac.eap_ssl_security_level` | StaticInt64(0) | N/A |
| `.mist_nac.idp_machine_cert_lookup_field` | StaticString("automatic") | N/A |
| `.mist_nac.idp_user_cert_lookup_field` | StaticString("automatic") | N/A |
| `.mist_nac.use_ip_version` | StaticString("v4") | N/A |
| `.mist_nac.use_ssl_port` | StaticBool(false) | N/A |
| `.mxedge_mgmt.config_auto_revert` | StaticBool(false) | N/A |
| `.mxedge_mgmt.oob_ip_type` | StaticString(dhcp) | N/A |
| `.mxedge_mgmt.oob_ip_type6` | StaticString(autoconf) | N/A |
| `.switch_updown_threshold` | StaticInt64(0) | N/A |

* `mist_org_sso`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.role_attr_from` | StaticString("Role") | N/A |

* `mist_org_vpn` and `mist_site_vpn`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.paths.bfd_profile` | StaticString("broadband") | N/A |
| `.paths.bfd_use_tunnel_mode` | StaticBool(false) | N/A |
| `.paths.pod` | StaticInt64(1) | N/A |
| `.type` | StaticString("hub_spoke") | N/A |

* `mist_org_webhook` and `mist_site_webhook`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.single_event_per_message` | StaticString("in") | N/A |

* `mist_org_wxtag` and `mist_site_wxtag`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.op` | StaticString("in") | N/A |

* `site`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.notes` | N/A | StaticString("") |
| `.sitegroup_ids` | N/A | [] |

* `site_setting`
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.ap_updown_threshold` | StaticInt64(0) | N/A |
| `.device_updown_threshold` | StaticInt64(0) | N/A |
| `.engagement.bounce` | StaticString("Visitor") | StaticString("") |
| `.engagement.engaged` | StaticString("Associates") | StaticString("") |
| `.engagement.passerby` | StaticString("Passerby") | StaticString("") |
| `.engagement.stationed` | StaticString("Assets") | StaticString("") |
| `.engagement.hours.mon` | StaticString("") | N/A |
| `.engagement.hours.tue` | StaticString("") | N/A |
| `.engagement.hours.wed` | StaticString("") | N/A |
| `.engagement.hours.thu` | StaticString("") | N/A |
| `.engagement.hours.fri` | StaticString("") | N/A |
| `.engagement.hours.sat` | StaticString("") | N/A |
| `.engagement.hours.sun` | StaticString("") | N/A |
| `.gateway_mgmt.auto_signature_update.day_od_week` | N/A | StaticString("") |
| `.gateway_mgmt.config_revert_timer` | StaticInt64(10) | N/A |
| `.gateway_mgmt.disable_console` | StaticBool(false) | N/A |
| `.gateway_mgmt.disable_oob` | StaticBool(false) | N/A |
| `.gateway_mgmt.disable_usb` | StaticBool(false) | N/A |
| `.gateway_mgmt.fips_enabled` | StaticBool(false) | N/A |
| `.gateway_mgmt.security_log_source_interface` | N/A | StaticString("") |
| `.gateway_updown_threshold` | StaticInt64(0) | N/A |
| `.occupancy.min_duration` | StaticInt64(3000) | N/A |
| `.remove_existing_configs` | StaticBool(false) | N/A |
| `.report_gatt` | StaticBool(false) | N/A |
| `.rogue.min_rogue_duration` | StaticInt64(10) | N/A |
| `.rogue.min_rogue_rssi` | StaticInt64(-80) | N/A |
| `.rtsa.disable_pressure_sensor` | StaticBool(false) | N/A |
| `.switch_updown_threshold` | StaticInt64(0) | N/A |
| `.synthetic_test.disabled` | StaticBool(false) | N/A |
| `.track_anonymous_devices` | StaticBool(false) | N/A |
| `.track_anonymous_devices.dot1x` | StaticBool(false) | N/A |
| `.wifi.cisco_enabled` | StaticBool(false) | N/A |
| `.wifi.disable_11k` | StaticBool(false) | N/A |
| `.wifi.disable_radios_when_power_constrained` | StaticBool(false) | N/A |
| `.wifi.enable_arp_spoof_check` | StaticBool(false) | N/A |
| `.wifi.enable_shared_radio_scanning` | StaticBool(true) | N/A |
| `.wifi.mesh_enable_crm` | StaticBool(false) | N/A |


--- 

## Release Notes for v0.3.0
**release data**: March 14th, 2025 

This release is adding new attributes based on the Mist Cloud push from February 2025.

### New Cloud support
Add support for Mist Cloud Global 05 (manage.gc4.mist.com / api.gc4.mist.com)

### Breaking Changes
#### mist_org_inventory
!> This version is removing the deprecated `mist_org_inventory.devices` attribute. Please make sure to migrate to the `mist_org_inventory.inventory` attribute before upgrading to this version.

Process to migrate from the `devices` attribute to the `inventory` one:
- update your `mist_org_inventory.devices` list by adding the attribute `unclaim_when_destroyed`=`false` to each device (this is just for security, the migration process will not touch the devices in the Mist Cloud)
- apply the configuration change. This will only update the resource state and mark each device to be kept "as is" during the migration process (the devices won't be unassigned or unclaimed)
- update your `mist_org_inventory` resource to use the `mist_org_inventory.inventory` attribute and remove the `mist_org_inventory.devices` one. If needed, update the rest of your configuration to match the changes.
- apply the configuration change


#### Other breaking changes
The following changes were required to support API type possibilities and/or to add the possibility to support the use of {{variables}} in the attribute values: 
* `mist_org_nactag` resource:
  * change type of `mist_org_nactag.gbp_tag` from `int64` to `string`
* `mist_device_switch` resource:
  * change type of `mist_device_switch.port_usage.reauth_interval` from `int64` to `string`
  * change type of `mist_device_switch.local_port_config.reauth_interval` from `int64` to `string`
* `mist_org_networktemplate` resource:
  * change type of `mist_org_networktemplate.port_usage.reauth_interval` from `int64` to `string`
  * change type of `mist_site_networktemplate.port_usage.reauth_interval` from `int64` to `string`
* `mist_org_wlan` resource:
  * change type of `mist_org_wlan.app_qos.apps.dscp` from `int64` to `string`
  * change type of `mist_site_wlan.app_qos.otherscp` from `int64` to `string`
  * change type of `mist_site_wlan.app_qos.others.dscp` from `int64` to `string`
* `mist_device_gateway` resource:
  * change type of `mist_device_gateway.bgp_config.local_as` from `int64` to `string`
  * change type of `mist_device_gateway.bgp_config.neighbor_as` from `int64` to `string`
  * change type of `mist_device_gateway.bgp_config.local_as` from `int64` to `string`
* `mist_org_deviceprofile_gateway` resource:
  * change type of `mist_org_deviceprofile_gateway.bgp_config.local_as` from `int64` to `string`
  * change type of `mist_org_deviceprofile_gateway.bgp_config.neighbor_as` from `int64` to `string`
  * change type of `mist_org_deviceprofile_gateway.bgp_config.neighbors.neighbor_as` from `int64` to `string`
* `mist_org_gatewaytemplate` resource:
  * change type of `mist_org_gatewaytemplate.bgp_config.local_as` from `int64` to `string`
  * change type of `mist_org_gatewaytemplate.bgp_config.neighbor_as` from `int64` to `string`
  * change type of `mist_org_gatewaytemplate.bgp_config.neighbors.neighbor_as` from `int64` to `string`
* `mist_site` resource:
  * change `mist_site.tzoffset` to read only to comply with Mist API behavior 

### New Datasource
* `mist_const_fingerprints`: The Fingerprint information can be used within `matching` and `not_matching` attributes of the NAC Rule resource (`mist_org_nacrule`)
* `mist_site_evpn_topologies`

### Resource Changes
#### mist_device_ap
* new attribute `mist_device_ap.lacp_config`
* new attribute `mist_device_ap.mesh.bands`

#### mist_device_gateway
* new attribute `mist_device_gateway.bgp_config.no_private_as`
* new attribute `mist_device_gateway.port_config.redundant_group`
* remove attribute (deprecated) `mist_device_gateway.port_config.vpn_paths.link`
* remove attribute (deprecated) `mist_device_gateway.routing_policies.terms.action.aggregate`
* fix typo to barbiturate `mist_device_gateway.routing_policies.export_communities`

#### mist_device_gateway
* new attribute `mist_device_switch.vrf_instances.evpn_auto_loopback_subnet`
* new attribute `mist_device_switch.vrf_instances.evpn_auto_loopback_subnet6`
* new attribute `mist_device_switch.vrf_instances.extra_routes6`

#### mist_org_deviceprofile_ap
* new attribute `mist_org_deviceprofile_ap.lacp_config`
* new attribute `mist_org_deviceprofile_ap.mesh.bands`

#### mist_deviceprofile_gateway
* new attribute `mist_deviceprofile_gateway.bgp_config.no_private_as`
* new attribute `mist_deviceprofile_gateway.port_config.redundant_group`
* remove attribute (deprecated) `mist_deviceprofile_gateway.port_config.vpn_paths.link`
* remove attribute (deprecated) `mist_deviceprofile_gateway.routing_policies.terms.action.aggregate`
* fix typo to barbiturate `mist_deviceprofile_gateway.routing_policies.export_communities`

#### mist_gatewaytemplate
* new attribute `mist_gatewaytemplate.bgp_config.no_private_as`
* new attribute `mist_gatewaytemplate.port_config.redundant_group`
* remove attribute (deprecated) `mist_gatewaytemplate.port_config.vpn_paths.link`
* remove attribute (deprecated) `mist_gatewaytemplate.routing_policies.terms.action.aggregate`
* fix typo to barbiturate `mist_gatewaytemplate.routing_policies.export_communities`

### mist_org_evpn_topology
* new attribute `mist_org_evpn_topology.evpn_options.evpn_options.per_vlan_vga_v6_mac`

#### mist_org_nacrule
* new attribute `mist_org_nacrule.matching.family`
* new attribute `mist_org_nacrule.matching.mfg`
* new attribute `mist_org_nacrule.matching.model`
* new attribute `mist_org_nacrule.matching.os_type`
* new attribute `mist_org_nacrule.not_matching.family`
* new attribute `mist_org_nacrule.not_matching.mfg`
* new attribute `mist_org_nacrule.not_matching.model`
* new attribute `mist_org_nacrule.not_matching.os_type`

#### mist_org_networktemplate
* new attribute `mist_org_networktemplate.vrf_instances.evpn_auto_loopback_subnet`
* new attribute `mist_org_networktemplate.vrf_instances.evpn_auto_loopback_subnet6`
* new attribute `mist_org_networktemplate.vrf_instances.extra_routes6`

#### mist_org_service
* new attribute `mist_org_service.org_services[*].client_limit_down`
* new attribute `mist_org_service.org_services[*].client_limit_up`
* new attribute `mist_org_service.org_services[*].service_limit_down`
* new attribute `mist_org_service.org_services[*].service_limit_up`

#### mist_org_servicepolicy
* new attribute `mist_org_servicepolicy.aamw_profile`
  * new attribute `mist_org_servicepolicy.aamw_profile.aamwprofile_id`
  * new attribute `mist_org_servicepolicy.aamw_profile.enabled`
  * new attribute `mist_org_servicepolicy.aamw_profile.profile`

#### mist_org_setting
* new attribute `org_setting.junos_shell_access`
  * new attribute `org_setting.junos_shell_access.admin`
  * new attribute `org_setting.junos_shell_access.helpdesk`
  * new attribute `org_setting.junos_shell_access.read`
  * new attribute `org_setting.junos_shell_access.write`
* new attribute `org_setting.mxedge_mgmt.config_auto_revert`
* new attribute `org_setting.switch_mgmt.remove_existing_configs`

#### mist_org_vpn
* new attribute `mist_org_vpn.path_selection`
  * new attribute `mist_org_vpn.path_selection.strategy`
* new attribute `mist_org_vpn.path.bfd_use_tunnel_mode`
* new attribute `mist_org_vpn.path.peer_paths`
  * new attribute `mist_org_vpn.path.peer_paths.preference`
* new attribute `mist_org_vpn.traffic_shaping`
  * new attribute `mist_org_vpn.traffic_shaping.class_percentage`
  * new attribute `mist_org_vpn.traffic_shaping.enabled`
  * new attribute `mist_org_vpn.traffic_shaping.max_tx_kbps`
* new attribute `mist_org_vpn.type`

#### mist_org_webhook
* new attribute `mist_org_webhook.single_event_per_message`

#### mist_org_wlan
* new attribute `mist_org_wlan.wlan.disable_11be`
* new attribute `mist_org_wlan.wlan.rateset.eht`
* new attribute `mist_org_wlan.wlan.rateset.he`

#### mist_org_wlan_portal_template
* new attribute `org_wlan_portal_template.marketing_policy_link`
* new attribute `org_wlan_portal_template.marketing_policy_opt_in`
* new attribute `org_wlan_portal_template.marketing_policy_opt_in_label`
* new attribute `org_wlan_portal_template.marketing_policy_opt_in_text`
* new attribute `org_wlan_portal_template.locale[*].marketing_policy_link`
* new attribute `org_wlan_portal_template.locale[*].marketing_policy_opt_in`
* new attribute `org_wlan_portal_template.locale[*].marketing_policy_opt_in_label`
* new attribute `org_wlan_portal_template.locale[*].marketing_policy_opt_in_text`

#### mist_site
* new attribute `mist_site.tz_offset`

### mist_site_evpn_topology
* new attribute `mist_site_evpn_topology.evpn_options.evpn_options.per_vlan_vga_v6_mac`

#### mist_site_networktemplate
* new attribute `mist_site_networktemplate.vrf_instances.evpn_auto_loopback_subnet`
* new attribute `mist_site_networktemplate.vrf_instances.evpn_auto_loopback_subnet6`
* new attribute `mist_site_networktemplate.vrf_instances.extra_routes6`

#### mist_site_setting
* new attribute `mist_site_setting.enable_unii4`
* new attribute `mist_site_setting.gateway_mgmt.disable_usb`
* new attribute `mist_site_setting.gateway_mgmt.fips_enabled`
* new attribute `mist_site_setting.rogue.allowed_vlan_ids`
* new attribute `mist_site_setting.rogue.min_rogue_duration`
* new attribute `mist_site_setting.rogue.min_rogue_rssi`

#### mist_site_webhook
* new attribute `mist_site_webhook.single_event_per_message`

#### mist_site_wlan
* new attribute `mist_site_wlan.wlan.disable_11be`
* new attribute `mist_site_wlan.wlan.rateset.eht`
* new attribute `mist_site_wlan.wlan.rateset.he`

#### mist_site_wlan_portal_template
* new attribute `mist_site_wlan_portal_template.marketing_policy_link`
* new attribute `mist_site_wlan_portal_template.marketing_policy_opt_in`
* new attribute `mist_site_wlan_portal_template.marketing_policy_opt_in_label`
* new attribute `mist_site_wlan_portal_template.marketing_policy_opt_in_text`
* new attribute `mist_site_wlan_portal_template.locale[*].marketing_policy_link`
* new attribute `mist_site_wlan_portal_template.locale[*].marketing_policy_opt_in`
* new attribute `mist_site_wlan_portal_template.locale[*].marketing_policy_opt_in_label`
* new attribute `mist_site_wlan_portal_template.locale[*].marketing_policy_opt_in_text`

### Data Sources Changes

#### mist_const_webhooks
* new attribute `mist_const_webhooks.const_webhooks[*].allows_single_event_per_message`

#### mist_device_ap_stats
* remove attribute (removed from API) `mist_device_ap_stats.device_ap_stats[*].use_auto_placement`
* new attribute `mist_device_ap_stats.device_ap_stats[*].gps`
  * new attribute `mist_device_ap_stats.device_ap_stats[*].gps.accuracy`
  * new attribute `mist_device_ap_stats.device_ap_stats[*].gps.altitude`
  * new attribute `mist_device_ap_stats.device_ap_stats[*].gps.latitude`
  * new attribute `mist_device_ap_stats.device_ap_stats[*].gps.longitude`
  * new attribute `mist_device_ap_stats.device_ap_stats[*].gps.src`
  * new attribute `mist_device_ap_stats.device_ap_stats[*].gps.timestamp`
* new attribute `mist_device_ap_stats.device_ap_stats[*].num_wlans`
* new attribute `mist_device_ap_stats.device_ap_stats[*].port_stats.rx_peak_bps`
* new attribute `mist_device_ap_stats.device_ap_stats[*].port_stats.tx_peak_bps`

#### mist_device_gateway_stats
* new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.evpn_overlay`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.for_overlay`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.local_as`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.neighbor`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.neighbor_as`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.neighbor_mac`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.node`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.rx_pkts`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.rx_routes`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.state`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.timestamp`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.tx_pkts`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.tx_routes`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.up`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.uptime`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].bgp_peers.vrfName`
* new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.active`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.auth_state`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.disabled`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.for_site`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.full_duplex`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.jitter`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.latency`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.loss`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.lte_iccid`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.lte_imei`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.lte_imsi`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.mac_count`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.mac_limit`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.neighbor_mac`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.neighbor_port_desc`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.neighbor_system_name`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.poe_disabled`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.poe_mode`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.poe_on`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.port_id`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.port_mac`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.port_usage`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.power_draw`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.rx_bcast_pkts`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.rx_bps`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.rx_bytes`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.rx_errors`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.rx_mcast_pkts`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.rx_pkts`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.speed`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.stp_role`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.stp_state`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.tx_bcast_pkts`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.tx_bps`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.tx_bytes`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.tx_errors`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.tx_mcast_pkts`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.tx_pkts`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.type`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.unconfigured`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.up`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.xcvr_model`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.xcvr_part_number`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].ports.xcvr_serial`
* new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.auth_algo`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.encrypt_algo`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.ike_version`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.ip`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.last_event`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.last_flapped`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.node`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.peer_host`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.peer_ip`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.priority`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.protocol`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.rx_bytes`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.rx_pkts`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.tunnel_name`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.tx_bytes`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.tx_pkts`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.up`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.uptime`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].tunnels.wan_name`
* new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.is_active`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.last_seen`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.latency`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.mos`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.mtu`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.peer_mac`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.peer_port_id`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.peer_router_name`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.peer_site_id`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.port_id`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.router_name`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.type`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.up`
  * new attribute `mist_device_gateway_stats.device_gateway_stats[*].vpn_peers.uptime`


#### mist_device_switch_stats
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].module_stat[*].cpu_stat`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].module_stat[*].cpu_stat.idle`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].module_stat[*].cpu_stat.interrupt`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].module_stat[*].cpu_stat.load_avg`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].module_stat[*].cpu_stat.system`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].module_stat[*].cpu_stat.user`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].module_stat[*].locating`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].module_stat[*].mac`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].module_stat[*].mactype`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.active`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.auth_state`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.disabled`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.for_site`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.full_duplex`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.jitter`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.last_flapped`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.latency`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.loss`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.lte_iccid`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.lte_imei`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.lte_imsi`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.mac_count`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.mac_limit`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.neighbor_mac`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.neighbor_port_desc`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.neighbor_system_name`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.org_id`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.poe_disabled`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.poe_mode`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.poe_on`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.port_id`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.port_mac`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.port_usage`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.power_draw`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.rx_bcast_pkts`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.rx_bps`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.rx_bytes`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.rx_errors`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.rx_mcast_pkts`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.rx_pkts`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.site_id`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.speed`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.stp_role`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.stp_state`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.tx_bcast_pkts`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.tx_bps`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.tx_bytes`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.tx_errors`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.tx_mcast_pkts`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.tx_pkts`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.type`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.unconfigured`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.up`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.xcvr_model`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.xcvr_part_number`
  * new attribute `mist_device_switch_stats.deviceswitch_stats[*].ports.xcvr_serial`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].vc_setup_info.current_state`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].vc_setup_info.last_update`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].vc_setup_info.request_time`
* new attribute `mist_device_switch_stats.deviceswitch_stats[*].vc_setup_info.request_type`

### mist_org_evpn_topologies
* new attribute `mist_site_evpn_topologies.org_evpn_topologies[*].evpn_options.per_vlan_vga_v6_mac`

### mist_org_nac_metadata
* new attribute `mist_org_nac_metadata.scim_base_url`

#### mist_org_nacrules
* new attribute `mist_org_narcules.org_narcules[*].enabled`
* new attribute `mist_org_narcules.org_narcules[*].order`

#### mist_org_networks
* new attribute `mist_org_networks.org_networks[*].subnet`
* new attribute `mist_org_networks.org_networks[*].subnet6`
* new attribute `mist_org_networks.org_networks[*].vlan_id`

#### mist_org_rftemplates
* new attribute `mist_org_rftemplates.org_rftemplate[*].country_code`

#### mist_org_services
* new attribute `mist_org_services.org_services[*].addresses`
* new attribute `mist_org_services.org_services[*].app_categories`
* new attribute `mist_org_services.org_services[*].app_subcategories`
* new attribute `mist_org_services.org_services[*].apps`
* new attribute `mist_org_services.org_services[*].client_limit_down`
* new attribute `mist_org_services.org_services[*].client_limit_up`
* new attribute `mist_org_services.org_services[*].description`
* new attribute `mist_org_services.org_services[*].dscp`
* new attribute `mist_org_services.org_services[*].failover_policy`
* new attribute `mist_org_services.org_services[*].hostnames`
* new attribute `mist_org_services.org_services[*].max_jitter`
* new attribute `mist_org_services.org_services[*].max_latency`
* new attribute `mist_org_services.org_services[*].max_loss`
* new attribute `mist_org_services.org_services[*].service_limit_down`
* new attribute `mist_org_services.org_services[*].service_limit_up`
* new attribute `mist_org_services.org_services[*].sle_enabled`
* new attribute `mist_org_services.org_services[*].specs`
* new attribute `mist_org_services.org_services[*].ssr_relaxed_tcp_state_enforcement`
* new attribute `mist_org_services.org_services[*].traffic_class`
* new attribute `mist_org_services.org_services[*].traffic_type`
* new attribute `mist_org_services.org_services[*].type`
* new attribute `mist_org_services.org_services[*].urls`

#### mist_org_servicepolicies
* new attribute `mist_org_servicepolicies.org_servicepolicies[*].aamw`
* new attribute `mist_org_servicepolicies.org_servicepolicies[*].antivirus`
* new attribute `mist_org_servicepolicies.org_servicepolicies[*].ssl_proxy`

### mist_org_sso_metadata
* new attribute `mist_org_sso_metadata.scim_base_url`

#### mist_org_vpns
* new attribute `mist_org_webhooks.org_vpns[*].path_selection`
  * new attribute `mist_org_webhooks.org_vpns[*].path_selection.class_percentage`
  * new attribute `mist_org_webhooks.org_vpns[*].path_selection.enabled`
  * new attribute `mist_org_webhooks.org_vpns[*].path_selection.max_tx_kbps`
* new attribute `mist_org_webhooks.org_vpns[*].type`

#### mist_org_webhooks
* new attribute `mist_org_webhooks.org_webhooks[*].single_event_per_message`

#### mist_org_wlans
* new attribute `mist_org_wlans.org_wlans[*].disable_11be`
* new attribute `mist_org_wlans.org_wlans[*].rateset.eht`
* new attribute `mist_org_wlans.org_wlans[*].rateset.he`

#### mist_org_wxtags
* new attribute `mist_org_wxtags.org_wxtags[*].mac`
* new attribute `mist_org_wxtags.org_wxtags[*].match`
* new attribute `mist_org_wxtags.org_wxtags[*].op`
* new attribute `mist_org_wxtags.org_wxtags[*].specs`
* new attribute `mist_org_wxtags.org_wxtags[*].specs.port_range`
* new attribute `mist_org_wxtags.org_wxtags[*].specs.protocol`
* new attribute `mist_org_wxtags.org_wxtags[*].specs.subnets`
* new attribute `mist_org_wxtags.org_wxtags[*].values`
* new attribute `mist_org_wxtags.org_wxtags[*].vlan_id`

#### mist_site_webhooks
* new attribute `mist_site_webhooks.site_webhooks[*].single_event_per_message`

#### mist_site_wlans
* new attribute `mist_site_wlans.site_wlans[*].disable_11be`
* new attribute `mist_site_wlans.site_wlans[*].rateset.eht`
* new attribute `mist_site_wlans.site_wlans[*].rateset.he`
