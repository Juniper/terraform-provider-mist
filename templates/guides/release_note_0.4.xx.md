---
subcategory: "Release Notes"
page_title: "v0.4.xx"
description: |-
    Release Notes for v0.4.xx
---

# Release Notes for v0.4.xx


## Release Notes for v0.4.6
**Release Date**: July 30th, 2025 

### Fixes
* **[Issue 122](https://github.com/Juniper/terraform-provider-mist/issues/122):** Fix the `.bgp_config` attributes validators in the `mist_device_gateway`, `mist_deviceprofile_gateway` and `mist_org_gatewaytemplate` resources to correct the fields required when the `.bgp_config.via` attribute is set to `vpn`.   
The attributes `.bgp_config.hold_time`, `.bgp_config.local_as`, `.bgp_config.neighbors`,  `.bgp_config.no_private_as` and `bgp_config.type` are now only required when the `.bgp_config.via` is set to `lan`, `tunnel` or `wan`.  
In addition, the `.bgp_config.graceful_restart_time` and `.bgp_config.hold_time` default values have been removed from the provider schema, as they are not applicable when the `.bgp_config.via` is set to `vpn`.



## Release Notes for v0.4.5
**Release Date**: July 29th, 2025 

This release doesn't include any new features or fixes. It is only fixing the release notes formatting.


## Release Notes for v0.4.4
**Release Date**: July 29th, 2025 

### Fixes
* **[Issue 120](https://github.com/Juniper/terraform-provider-mist/issues/120):** Replace the `.tunnel_configs.enable` attribute with `.tunnel_configs.enabled` in the `mist_device_gateway`, `mist_deviceprofile_gateway` and `mist_org_gatewaytemplate` resources to match the API attribute name.
* **[Issue 121](https://github.com/Juniper/terraform-provider-mist/issues/121):** Fix the `mist_org_server.urls` validators to allow the `urls` attribute to be set when the `type` is set to `urls`
* **[Issue 122](https://github.com/Juniper/terraform-provider-mist/issues/122):** Remove some of the attributes default value from the `mist_device_gateway`, `mist_deviceprofile_gateway` and `mist_org_gatewaytemplate` resources (see below for details).

#### Attributes added
- **`mist_device_switch` resource**
 - `.ospf_config` has been added
 - `.ospf_config.enabled` has been added
 - `.ospf_config.area` has been added
 - `.ospf_config.area.no_summary` has been added


### Resources default values changed

Changes have been applied to resources to reduce configuration drift when importing resources or saving changes from the Mist UI. These updates aim to align Terraform resource states with the Mist UI default values. However, some default values are dynamic and depend on other parameter values, making it currently impossible to completely eliminate configuration drift in certain scenarios.

~> **Warning** Some default values have been removed from the Terraform Provider resource schemas.  
These changes may lead to configuration drift if the affected attributes are not explicitly defined in your HCL configuration.  
Attributes without explicit definitions will default to `null`, but this will not alter the actual configuration in the Mist Cloud (the Mist Cloud will use the default value). To avoid discrepancies, ensure that all required attributes are explicitly set in your configuration.

*  **`mist_device_gateway`, `mist_deviceprofile_gateway` and `mist_org_gatewaytemplate` resources**
| Attribute | Previous Default | New Default |
|-----------|-----------|-----------|
| `.bgp_config.bfd_minimum_interval` | StaticInt64(350) | N/A |
| `.bgp_config.bfd_multiplier` | StaticInt64(350) | N/A |
| `.bgp_config.neighbors.networks` | ListNull(types.StringType) | N/A |
| `.bgp_config.no_readvertise_to_overlay` | StaticBool(false) | N/A |
| `.bgp_config.via` | StaticString("lan") | N/A |
| `.dhcpd_config.config.dns_suffix` | ListNull(types.StringType) | N/A |
| `.dhcpd_config.config.lease_time` | StaticInt64(86400) | N/A |
| `.dhcpd_config.config.server_id_override` | StaticBool(false) | N/A |
| `.dhcpd_config.config.servers` | ListNull(types.StringType) | N/A |
| `.dhcpd_config.config.servers6` | ListNull(types.StringType) | N/A |
| `.dhcpd_config.config.type` | StaticString("local") | N/A |
| `.dhcpd_config.config.type6` | StaticString("none") | N/A |
| `.dhcpd_config.enabled` | StaticBool(true) | N/A |
| `.idp_profiles.overwrites.action` | StaticString("alert") | N/A |
| `.path_preferences.paths.networks` | ListNull(types.StringType) | N/A |
| `.path_preferences.paths.target_ips` | ListNull(types.StringType) | N/A |
| `.port_config.ae_disable_lacp` | StaticBool(false) | N/A |
| `.port_config.ae_lacp_force_up` | StaticBool(false) | N/A |
| `.port_config.aggregated` | StaticBool(false) | N/A |
| `.port_config.critical` | StaticBool(false) | N/A |
| `.port_config.disable_autoneg` | StaticBool(false) | N/A |
| `.port_config.dsl_type` | StaticString("vdsl") | N/A |
| `.port_config.dsl_vci` | StaticInt64(35) | N/A |
| `.port_config.dsl_vpi` | StaticInt64(0) | N/A |
| `.port_config.duplex` | StaticString("auto") | N/A |
| `.port_config.lte_auth` | StaticString("none") | N/A |
| `.port_config.networks` | ListNull(types.StringType) | N/A |
| `.port_config.poe_disabled` | StaticBool(false) | N/A |
| `.port_config.ip_config.pppoe_auth` | StaticString("none") | N/A |
| `.port_config.ip_config.type` | StaticString("dhcp") | N/A |
| `.port_config.preserve_dscp` | StaticBool(false) | N/A |
| `.port_config.reth_nodes` | ListNull(types.StringType) | N/A |
| `.port_config.speed` | StaticString("auto") | N/A |
| `.port_config.ssr_no_virtual_mac` | StaticBool(false) | N/A |
| `.port_config.svr_port_range` | StaticString("none") | N/A |
| `.port_config.traffic_shaping.enabled` | StaticBool(false) | N/A |
| `.port_config.vpn_paths.bfd_profile` | StaticString("broadband") | N/A |
| `.port_config.vpn_paths.bfd_use_tunnel_mode` | StaticBool(false) | N/A |
| `.port_config.vpn_paths.role` | StaticString("spoke") | N/A |
| `.port_config.vpn_paths.traffic_shaping.enabled` | StaticBool(false) | N/A |
| `.port_config.wan_arp_policer` | StaticString("default") | N/A |
| `.port_config.wan_disable_speedtest` | StaticBool(false) | N/A |
| `.port_config.wan_extra_routes` | types.MapValueMust(WanExtraRoutesValue{}.Type(ctx), map[string]attr.Value{}) | N/A |
| `.port_config.wan_networks` | ListNull(types.StringType) | N/A |
| `.port_config.wan_probe_override.probe_profile` | StaticString("broadband") | N/A |
| `.port_config.wan_source_nat.disabled` | StaticBool(false) | N/A |
| `.port_config.wan_type` | StaticString("broadband") | N/A |
| `.service_policies.antivirus.enabled` | StaticBool(false) | N/A |
| `.service_policies.appqoe.enabled` | StaticBool(false) | N/A |
| `.service_policies.ewf.enabled` | StaticBool(false) | N/A |
| `.service_policies.ewf.profile` | StaticString("strict") | N/A |
| `.service_policies.idp.enabled` | StaticBool(false) | N/A |
| `.service_policies.idp.profile` | StaticString("strict") | N/A |
| `.service_policies.ssl_proxy.ciphers_category` | StaticString("strict") | N/A |
| `.tunnel_configs.ike_mode` | StaticString("main") | N/A |
| `.tunnel_configs.ike_proposals.dh_group` | StaticString("14") | N/A |
| `.tunnel_configs.ike_proposals.enc_algo` | StaticString("aes256") | N/A |
| `.tunnel_configs.ipsec_proposals.dh_group` | StaticString("14") | N/A |
| `.tunnel_configs.ipsec_proposals.enc_algo` | StaticString("aes256") | N/A |
| `.tunnel_configs.networks` | ListNull(types.StringType) | N/A |
| `.tunnel_configs.version` | StaticString("2") | N/A |




## Release Notes for v0.4.3
**Release Date**: July 11th, 2025 

### Fixes
* **[Issue 118](https://github.com/Juniper/terraform-provider-mist/issues/118):** Fix the `mist_device_switch.ip_config.dns` validators requiring the DNS to be set when the `ip_config.type` is set to `static`. The DNS is not optional in this case, so the validator has been updated to allow it to be empty.


## Release Notes for v0.4.2
**Release Date**: July 1st, 2025 

### Fixes
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
