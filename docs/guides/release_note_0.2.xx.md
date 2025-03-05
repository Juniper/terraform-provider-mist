---
subcategory: "Release Notes"
page_title: "v0.2.xx"
description: |-
    Release Notes for v0.2.xx
---

# Release Notes for v0.2.xx

## Release Notes for v0.2.23
**release data**:

This release is adding new attributes based on the Mist Cloud push from February 2025.

### New Cloud support
Add support for Mist Cloud Global 05 (manage.gc4.mist.com / api.gc4.mist.com)

### New Datasource
* `mist_const_fingerprints`: The Fingerprint information can be used within `matching` and `not_matching` attributes of the NAC Rule resource (`mist_org_nacrule`)

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

#### mist_device_ap_stats
* remove attribute (removed from API) `mist_device_ap_stats.device_ap_stats[*].use_auto_placement`

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



## Release Notes for v0.2.22
**release date** : February 28th, 2025

### Fixes
* [Issue 89](https://github.com/Juniper/terraform-provider-mist/issues/89): fix typo in `snmp_config.v3_config.usm.users.authentication_type` allowed values in the following resources: `mist_device_switch`, `mist_org_networktemplate`, `mist_site_networktemplate`
* [Issue 90](https://github.com/Juniper/terraform-provider-mist/issues/90): adding a new validator to not accept an empty list in the `snmp_config.client_list` attribute in the following resources: `mist_device_switch`, `mist_org_networktemplate`, `mist_site_networktemplate`



## Release Notes for v0.2.21
**release date** : February 27th, 2025

### New Resource
* `mist_org_avprofile`: Allows configuring the Antivirus profiles for SRX Gateways. The Antivirus profiles can be used within the following resources:  
  * `mist_org_servicepolicy.antivirus`
  * `mist_org_gatewaytemplate.service_policies.antivirus`
  * `mist_org_deviceprofile_gateway.service_policies.antivirus`
  * `mist_device_gateway.service_policies.antivirus`

### New Data Sources
* `mist_upgrade_device`: Allows triggering a firmware upgrade on a single device. Site and Org upgrades will be available at a later time.

### New Functions
* `search_vc_by_member_claimcode`, `search_vc_by_member_mac` and `search_vc_by_member_serial` have been added to easily retrieve the Virtual Chassis' 
`id` and `site_id` based on the claim code, MAC address or Serial Number of ones of the Virtual Chassis member.

### Changes
* add `mist_device_ap.flow_control` attribute
* add the `antivirus` and `sll_proxy` attributes to `mist_org_servicepolicy`, `mist_org_gatewaytemplate.service_policies`, `mist_org_deviceprofile_gateway.service_policies` and `mist_device_gateway.service_policies`
* [Issue 85](https://github.com/Juniper/terraform-provider-mist/issues/85): adding a new validator to `switch_matching.rules.name` to only allow the rule name `default` within the last rule of the list. This is applied to the resources `mist_org_networktemplate` and `mist_site_networktemplate`
* update the `mist_org_inventory` resource to support the new Virtual Chassis "virtual MAC" that will be deployed later this year by Mist.   
This new "virtual MAC" will be assigned to each new Virtual Chassis, and will be used to generate the `device_id` instead of the MAC Address of the primary member. 
This change will allow replacing the Virtual Chassis primary member without any impact on the Virtual Chassis `device_id`.  
In addition, new "validations" have been added to the `mist_org_inventory.inventory` resource to validate all the Virtual Chassis members are assigned to the same site. 

### Fixes
* [Issue 88](https://github.com/Juniper/terraform-provider-mist/issues/88): fix issues with the `remote_syslog` attribute in the following resources: `mist_device_switch`, `mist_org_networktemplate`, `mist_site_networktemplate`
* [Issue 87](https://github.com/Juniper/terraform-provider-mist/issues/87): fix issues with the `snmp_config` attribute when `v3_config` is not defined in the following resources: `mist_device_switch`, `mist_org_networktemplate`, `mist_site_networktemplate`
* [Issue 86](https://github.com/Juniper/terraform-provider-mist/issues/86): fix issues with the `switch_mgmt.tacacs` attribute in the following resources: `mist_org_networktemplate`, `mist_site_networktemplate`
* [Issue 84](https://github.com/Juniper/terraform-provider-mist/issues/84): fix issues with the `snmp_config.v3_config` attribute in the following resources: `mist_device_switch`, `mist_org_networktemplate`, `mist_site_networktemplate`
* [Issue 83](https://github.com/Juniper/terraform-provider-mist/issues/83): fix an issue with the `mist_org_nactags` datasource
* [Issue 81](https://github.com/Juniper/terraform-provider-mist/issues/81): fix typo in `mist_org_networktemplate` property `bypass_auth_when_server_down_for_unkonwn_client`


## Release Notes for v0.2.20
**release date** : February 14th, 2025

### Fixes
* Fix issue when `mist_org_sso_role.privileges.views` has empty value


## Release Notes for v0.2.19
**release date** : February 14th, 2025

### New Data Sources
* `mist_device_versions`: Allows retrieving the list of available firmware for a specific type of hardware
### New Resources
* `mist_upgrade_device`: Allows triggering a firmware upgrade on a single device. Site and Org upgrades will be available at a later time.

### Changes
* change the `mist_org_sso_role.privileges.views` from `string` to `list` to match the Mist API Structure. This change requires updating the configuration and the state to match the new format if this attribute has already been configured.

### Fixes
* [Issue 70](https://github.com/Juniper/terraform-provider-mist/issues/70): In some conditions, the list of MAC addresses returned by Mist when assigning device profiles to a device may be different from the configured order, which generates an error in the provider. The type of the list in the provider has been changed from `List` to `Set` to avoid this issue. This change doesn't impact the existing configurations or deployments.
* [Issue 74](https://github.com/Juniper/terraform-provider-mist/issues/74): The procedure to unset the `alarmtemplate_id` at the org level was not removing the reference, generating an error in the provider. This has been fixed and the Alarm Template can now be unreferenced as expected.
* [Issue 76](https://github.com/Juniper/terraform-provider-mist/issues/76): When multiple sites where added to the same sitegroup during the same run, the concurrent processing of the provider was causing an unexpected result in the Mist side in some conditions (the `site_ids` attribute was not updated as expected in the sitegroup object). To avoid this situation, the `mist_site` resource has been updated to force the provider to create/update/delete them sequentially, avoiding this "race condition" issue.
* [Issue 77](https://github.com/Juniper/terraform-provider-mist/issues/77): The `mist_org_webhooks` and `mist_site_webhooks` datasource were missing some attributes, generating an error in the provider. The missing attributes have been added.
* There was an issue when importing the `mist_org_inventory` resource, resulting in an empty inventory. This has been fixed, and the import function will import/generate the inventory with all the devices in the Mist Org Inventory.

## Release Notes for v0.2.18
**release date** : January 10th, 2025

### New Resources
* `mist_org_nac_endpoints`

### New Data Sources
* `mist_org_alarmtemplates`
* `mist_org_evpn_topologies`
* `mist_org_nac_endpoints`
* `mist_org_sso_roles`
* `mist_org_wlans`
* `mist_site_wlans`

### Changes
* Add `mist_org_sso_role.privilege.views` attribute
* Add `mist_site_setting.juniper_srx` attribute

### Fixes
* In some conditions, the assigned/unassigned devices process may fail on the Mist side, and the list of MAC addresses returned may not match what is expected. This case was not processed by the provider, resulting a "This is a bug in the provider, which should be reported in the provider's own issue tracker." error message.  
The provider is now returning a specific warning message when it was not able to unassign a device, and a specific error when it was not able to assign a device (issue [#70](https://github.com/Juniper/terraform-provider-mist/issues/70))
* `vpn_access.static_nat.wan_name` and `vpn_access.destination_nat.wan_name` was removed in a previous version, but some parts of the code were not updated. This issue was impacting the `mist_org_network`, `mist_device_gateway`, `mist_org_gatewaytemplate` and `mist_org_deviceprofile_gateway` resources (issue [#71](https://github.com/Juniper/terraform-provider-mist/issues/71))

## Release Notes for v0.2.17
**release date** : January 2nd, 2025

### Fixes
* fix `mist_org_psk.vlan_id` and `mist_site_psk.vlan_id` issue when converting the value from the Go SDK to Terraform.

## Release Notes for v0.2.16
**release date** : January 2nd, 2025

### Fixes
* fix `mist_org_wlan.dynamic_vlan.default_vlan_ids` and `mist_site_wlan.dynamic_vlan.default_vlan_ids` issue when converting the value from the Go SDK to Terraform.


## Release Notes for v0.2.15
**release date** : December 27th, 2024

!> Breaking changes. See below

### Breaking Changes
* The Following attributes have been changed from int64 to string to allow "mist variable" support:
  * `mist_org_network.internet_access.destination_nat.port`
  * `mist_org_network.vpn_access.destination_nat.port`
  * `mist_device_gateway.networks.internet_access.destination_nat.port`
  * `mist_device_gateway.networks.vpn_access.destination_nat.port`
  * `mist_deviceprofile_gateway.networks.internet_access.destination_nat.port`
  * `mist_deviceprofile_gateway.networks.vpn_access.destination_nat.port`
  * `mist_org_gatewaytemplate.networks.internet_access.destination_nat.port`
  * `mist_org_gatewaytemplate.networks.vpn_access.destination_nat.port`
* The following attributes have been changed from `optional` to `required`:
  * `mist_org_network.internet_access.static_nat.internal_ip`
  * `mist_org_network.internet_access.static_nat.name`
  * `mist_org_network.vpn_access.static_nat.internal_ip`
  * `mist_org_network.vpn_access.static_nat.name`
  * `mist_device_gateway.tunnel_configs.primary.hosts`
  * `mist_device_gateway.tunnel_configs.primary.wan_names`
  * `mist_device_gateway.tunnel_configs.secondary.hosts`
  * `mist_device_gateway.tunnel_configs.secondary.wan_names`
  * `mist_device_gateway.networks.internet_access.static_nat.internal_ip`
  * `mist_device_gateway.networks.internet_access.static_nat.name`
  * `mist_device_gateway.networks.vpn_access.static_nat.internal_ip`
  * `mist_device_gateway.networks.vpn_access.static_nat.name`
  * `mist_deviceprofile_gateway.tunnel_configs.primary.hosts`
  * `mist_deviceprofile_gateway.tunnel_configs.primary.wan_names`
  * `mist_deviceprofile_gateway.tunnel_configs.secondary.hosts`
  * `mist_deviceprofile_gateway.tunnel_configs.secondary.wan_names`
  * `mist_deviceprofile_gateway.networks.internet_access.static_nat.internal_ip`
  * `mist_deviceprofile_gateway.networks.internet_access.static_nat.name`
  * `mist_deviceprofile_gateway.networks.vpn_access.static_nat.internal_ip`
  * `mist_deviceprofile_gateway.networks.vpn_access.static_nat.name`
  * `mist_org_gatewaytemplate.tunnel_configs.primary.hosts`
  * `mist_org_gatewaytemplate.tunnel_configs.primary.wan_names`
  * `mist_org_gatewaytemplate.tunnel_configs.secondary.hosts`
  * `mist_org_gatewaytemplate.tunnel_configs.secondary.wan_names`
  * `mist_org_gatewaytemplate.networks.internet_access.static_nat.internal_ip`
  * `mist_org_gatewaytemplate.networks.internet_access.static_nat.name`
  * `mist_org_gatewaytemplate.networks.vpn_access.static_nat.internal_ip`
  * `mist_org_gatewaytemplate.networks.vpn_access.static_nat.name`


### Improvements
* add the `api_debug` flag to the provider properties to enable the logging of the SDK Requests and Responses
* improve the `mist_device_gateway`, `mist_deviceprofile_gateway` and `mist_org_gatewaytemplate` documentation
* improve validators for the following resources:
  * `mist_device_gateway`
  * `mist_device_switch`
  * `mist_deviceprofile_gateway`
  * `mist_org_gatewaytemplate` 
  * `mist_org_networktemplate` 
  * `mist_org_wlan` 
  * `mist_org_wxtag` 
  * `mist_site_networktemplate` 
  * `mist_site_wlan` 
  * `mist_site_wxtag` 
  * `mist_org_service` 

### Fixes
* fix the transformation of the `VlanIdWithVariable` SDK property. In some conditions, the value sent to the provider was an HEX string instead of the VLAN ID (or variable) value.

### Other Changes
Changes to the `mist_org_network` resource based on the OpenAPI changes:
* add the `internet_access.destination_nat.wan_name` attribute
* add the `internet_access.static_nat.wan_name` attribute
* add the `multicast` attribute

Remove the following optional attributes from the `mist_org_sso` resource (these attributes are used with NAC IDP, not Org SSO):
* `scim_enabled`
* `scim_secret_token`

Add the following optional attributes to the `mist_org_nacidp` resource:
* `oauth_ping_identity_region`
* `scim_enabled`
* `scim_secret_token`

Add the following optional attributes to the `mist_device_switch` resource:
* `notes`
* `local_port_config.note`

Changes to the `mist_device_gateway`, `mist_deviceprofile_gateway` and `mist_org_gatewaytemplate` resource based on the OpenAPI changes:
* attributes added:
  * `networks.internet_access.destination_nat.wan_name`
  * `networks.internet_access.static_nat.wan_name`
  * `port_config.wan_networks`
  * `routing_policies.action.aggregate`
  * `tunnel_configs.auto_provision.primary.probe_ips`
  * `tunnel_configs.auto_provision.secondary.probe_ips`
  * `tunnel_configs.auto_provision.provider`
  * `tunnel_configs.auto_provision.region`
* attributes removed:
  * `tunnel_configs.auto_provision.primary.num_hosts` (this setting is configured in the `tunnel_provider_options` object)
  * `tunnel_configs.auto_provision.secondary.num_hosts` (this setting is configured in the `tunnel_provider_options` object)
* attributes updated:
  * `tunnel_provider_options.jse.name` renamed to `tunnel_provider_options.jse.org_name` 
  * rework the whole `tunnel_provider_options.zscaler` object to match the Mist API structure (see the resource documentation for more details)




## Release Notes for v0.2.14
**release date** : December 20th, 2024


### Fixes
* Fixing issue when `mist_org_wlan.ap_ids` or `mist_org_wlan.ap_ids` is present but has an `null` value







## Release Notes for v0.2.13
**release date** : December 20th, 2024

### Improvements
* `mist_device_gateway_cluster`: Improve the creation resource behavior when one of both of the cluster nodes already belong to a cluster.  The provider will no more raise an error when the existing cluster in the Mist Cloud is matching the planned cluster (same primary node, same secondary node).
* `mist_org_inventory`: improve the deprecation message.

### Fixes
* [Issue 65](https://github.com/Juniper/terraform-provider-mist/issues/65): Fixing the `port_config.wan_source_nat` attribute in the `mist_device_gateway`, `mist_org_deviceprofile_gateway` and `mist_org_gatewaytemplate` resources 
* [Issue 66](https://github.com/Juniper/terraform-provider-mist/issues/66): Fixing `mist_org_wlan` resource following the v0.2.12 changes 
* [Issue 68](https://github.com/Juniper/terraform-provider-mist/issues/68): Fixing `tunnel_configs.auto_provision` attribute in `gateway` resources 
* Fixing issue when removing `rftemplate_id` from the `mist_site` resource







## Release Notes for v0.2.12
**release date** : December 13th, 2024

!> This release may generate multiple changes to the `org_wlan_resource` and `site_wlan_resource` resources during the first configuration sync. This is due to the new default values defined, and will not impact to actual SSID configuration deployed on the Access Points

### Changes
#### Documentation
* improve `org_wlan_resource` and `site_wlan_resource` resources documentation

#### WLAN resources default values
Changes applied to `org_wlan_resource` and `site_wlan_resource` to reduce configuration drift when saving the WLAN from the Mist UI. 
These changes try to mimic the Mist UI default values; however, some of them are changing based on other parameter values which make it currently impossible to eliminate the configuration drift.

List of the default value changes:

| Attribute                                  | Previous Default                                                    | New Default                                                                                                                                     |
|--------------------------------------------|---------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| `acct_servers`                             | not set                                                             | `types.ListValueMust(AcctServersValue{}.Type(ctx), []attr.Value{})`                                                                             |
| `airwatch.api_key`                         | not set                                                             | `""`                                                                                                                                            |
| `airwatch.console_url`                     | not set                                                             | `""`                                                                                                                                            |
| `airwatch.password`                        | not set                                                             | `""`                                                                                                                                            |
| `airwatch.username`                        | not set                                                             | `""`                                                                                                                                            |
| `airwatch`                                 | not set                                                             | `types.ObjectValueMust(AirwatchValue{}.AttributeTypes(ctx), ... )`                                                                              |
| `ap_ids`                                   | `types.ListValueMust(AcctServersValue{}.Type(ctx), []attr.Value{})` | `types.ListNull(types.StringType)`                                                                                                              |
| `app_limit`                                | not set                                                             | `types.ObjectValueMust(AppLimitValue{}.AttributeTypes(ctx), ... )`                                                                              |
| `app_limit`                                | not set                                                             | `types.ObjectValueMust(AppQosValue{}.AttributeTypes(ctx), ... )`                                                                                |
| `auth.anticlog_threshold`                  | `16`                                                                | removed                                                                                                                                         | 
| `auth.keys`                                | not set                                                             | `types.ListValueMust(types.StringType, []attr.Value{types.StringValue(""),types.StringValue(""),types.StringValue(""),types.StringValue(""),})` | 
| `auth.owe`                                 | `"disabled"`                                                        | removed                                                                                                                                         | 
| `auth.wep_as_secondary_auth`               | `false`                                                             | removed                                                                                                                                         | 
| `auth_servers`                             | not set                                                             | `types.ListValueMust(AcctServersValue{}.Type(ctx), []attr.Value{})`                                                                             |
| `auth_servers_nas_id`                      | not set                                                             | `""`                                                                                                                                            |
| `auth_servers_nas_ip`                      | not set                                                             | `""`                                                                                                                                            |
| `bonjour`                                  | not set                                                             | `types.ObjectValueMust(BonjourValue{}.AttributeTypes(ctx), ... )`                                                                               |
| `cisco_cwa`                                | not set                                                             | `types.ObjectValueMust(CiscoCwaValue{}.AttributeTypes(ctx), ... )`                                                                              |
| `client_limit_down`                        | not set                                                             | `1000`                                                                                                                                          |
| `client_limit_up`                          | not set                                                             | `512`                                                                                                                                           |
| `coa_servers`                              | not set                                                             | `types.ListValueMust(CoaServersValue{}.Type(ctx), []attr.Value{})`                                                                              |
| `disable_when_gateway_unreachable`         | `false`                                                             | removed                                                                                                                                         |
| `disable_when_mxtunnel_down`               | `false`                                                             | removed                                                                                                                                         |
| `dns_server_rewrite`                       | not set                                                             | `types.ObjectValueMust(DnsServerRewriteValue{}.AttributeTypes(ctx), ... )`                                                                      |
| `hotspot20`                                | not set                                                             | `types.ObjectValueMust(Hotspot20Value{}.AttributeTypes(ctx), ... )`                                                                             |
| `mist_nac`                                 | not set                                                             | `types.ObjectValueMust(MistNacValue{}.AttributeTypes(ctx), ... )`                                                                               |
| `mxtunnel_ids`                             | `types.ListValueMust(types.StringType, []attr.Value{})`             | `types.ListNull(types.StringType)`                                                                                                              |
| `mxtunnel_name`                            | `types.ListValueMust(types.StringType, []attr.Value{})`             | `types.ListNull(types.StringType)`                                                                                                              |
| `portal.allow_wlan_id_roam`                | `false`                                                             | removed                                                                                                                                         |
| `portal.amazon_email_domains`              | `types.ListValueMust(types.StringType, []attr.Value{})`             | `types.ListNull(types.StringType)`                                                                                                              |
| `portal.broadnet_sid`                      | `"MIST"`                                                            | removed                                                                                                                                         |
| `portal.broadnet_user_id`                  | `""`                                                                | removed                                                                                                                                         |
| `portal.clickatell_api_key`                | `""`                                                                | removed                                                                                                                                         |
| `portal.cross_site`                        | `false`                                                             | removed                                                                                                                                         |
| `portal.facebook_email_domains`            | `types.ListValueMust(types.StringType, []attr.Value{})`             | `types.ListNull(types.StringType)`                                                                                                              |
| `portal.google_email_domains`              | `types.ListValueMust(types.StringType, []attr.Value{})`             | `types.ListNull(types.StringType)`                                                                                                              |
| `portal.gupshup_password`                  | `""`                                                                | removed                                                                                                                                         |
| `portal.gupshup_userid`                    | `""`                                                                | removed                                                                                                                                         |
| `portal.microsoft_email_domains`           | `types.ListValueMust(types.StringType, []attr.Value{})`             | `types.ListNull(types.StringType)`                                                                                                              |
| `portal.puzzel_password`                   | `""`                                                                | removed                                                                                                                                         |
| `portal.puzzel_service_id`                 | `""`                                                                | removed                                                                                                                                         |
| `portal.puzzel_username`                   | `""`                                                                | removed                                                                                                                                         |
| `portal.sponsor_auto_approve`              | `false`                                                             | removed                                                                                                                                         |
| `portal.telstra_client_id`                 | `""`                                                                | removed                                                                                                                                         |
| `portal.telstra_client_secret`             | `""`                                                                | removed                                                                                                                                         |
| `portal.twilio_auth_token`                 | `""`                                                                | removed                                                                                                                                         |
| `portal.twilio_phone_number`               | `""`                                                                | removed                                                                                                                                         |
| `portal.twilio_sid`                        | `""`                                                                | removed                                                                                                                                         |
| `portal`                                   | not set                                                             | `types.ObjectValueMust(PortalValue{}.AttributeTypes(ctx), ... )`                                                                                |
| `qos`                                      | not set                                                             | `types.ObjectValueMust(ObjectValueMust{}.AttributeTypes(ctx), ... )`                                                                            |
| `radsec`                                   | not set                                                             | `types.ObjectValueMust(RadsecValue{}.AttributeTypes(ctx), ... )`                                                                                |
| `rateset`                                  | not set                                                             | `types.MapValueMust(RatesetValue{}.AttributeTypes(ctx), ... )`                                                                                  |
| `reconnect_clients_when_roaming_mxcluster` | `false`                                                             | removed                                                                                                                                         |
| `schedule`                                 | not set                                                             | `types.ObjectValueMust(ScheduleValue{}.AttributeTypes(ctx), ... )`                                                                              |
| `vlan_ids`                                 | not set                                                             | `types.ListValueMust(types.StringType, []attr.Value{})`                                                                                         |
| `wlan_limit_down`                          | not set                                                             | `20000`                                                                                                                                         |
| `wlan_limit_up`                            | not set                                                             | `10000`                                                                                                                                         |

#### WLAN resources validators
Validators applied to the WLAN resources attributes have been updated
to simplify the resource configuration and improve the configuration validity.

List of the validator changes:

| Attribute                       | Previous Default                                                                                                 | New Default                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|---------------------------------|------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `acct_servers`                  | `listvalidator.SizeAtLeast(1)`                                                                                   | removed                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| `airwatch.api_key`              | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                                     |
| `airwatch.console_url`          | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                                     |
| `airwatch.password`             | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                                     |
| `airwatch.username`             | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                                     |
| `auth.key_idx`                  | `mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("wep"))`     | `int64validator.Between(1, 4)`                                                                                                                                                                                                                                                                                                                                                                                                                    |
| `auth.keys`                     | `mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("wep"))`     | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("wep"))`                                                                                                                                                                                                                                                                                                                                     |
| `auth_servers`                  | `listvalidator.SizeAtLeast(1)`                                                                                   | removed                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| `dynamic_vlan`                  | `mistvalidator.CannotBeTrueWhenValueIs(path.MatchRoot("vlan_enabled"), types.BoolValue(true))`                   | `mistvalidator.CanOnlyTrueWhenValueIs(path.MatchRoot("vlan_enabled"), types.BoolValue(true)), boolvalidator.Any( mistvalidator.CanOnlyTrueWhenValueIs(path.MatchRoot("auth").AtName("enable_mac_auth"), types.BoolValue(true)), mistvalidator.CanOnlyTrueWhenValueIs(path.MatchRoot("auth").AtName("type"), types.StringValue("eap")), mistvalidator.CanOnlyTrueWhenValueIs(path.MatchRoot("auth").AtName("type"), types.StringValue("eap192")))` |
| `portal.azure_client_id`        | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("azure_enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                               |
| `portal.azure_client_secret`    | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("azure_enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                               |
| `portal.azure_tenant_id`        | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("azure_enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                               |
| `portal.broadnet_password`      | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("broadnet"))`                                                                                                                                                                                                                                                                                                                        |
| `portal.broadnet_sid`           | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("broadnet"))`                                                                                                                                                                                                                                                                                                                        |
| `portal.broadnet_user_id`       | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("broadnet"))`                                                                                                                                                                                                                                                                                                                        |
| `portal.clickatell_api_key`     | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("clickatell"))`                                                                                                                                                                                                                                                                                                                      |
| `portal.external_portal_url`    | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("auth"), types.StringValue("external"))`                                                                                                                                                                                                                                                                                                                                |
| `portal.facebook_client_id`     | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("facebook_enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                            |
| `portal.facebook_client_secret` | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("facebook_enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                            |
| `portal.forward_url`            | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("forward"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                                     |
| `portal.gupshup_password`       | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("gupshup"))`                                                                                                                                                                                                                                                                                                                         |
| `portal.gupshup_userid`         | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("gupshup"))`                                                                                                                                                                                                                                                                                                                         |
| `portal.password`               | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("passphrase_enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                          |
| `portal.puzzel_password`        | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("puzzel"))`                                                                                                                                                                                                                                                                                                                          |
| `portal.puzzel_service_id`      | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("puzzel"))`                                                                                                                                                                                                                                                                                                                          |
| `portal.puzzel_username`        | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("puzzel"))`                                                                                                                                                                                                                                                                                                                          |
| `portal.sms_provider`           | `stringvalidator.OneOf("", "broadnet", "clickatell", "gupshup", "manual", "puzzel", "telstra", "twilio")`        | `stringvalidator.OneOf("", "broadnet", "clickatell", "gupshup", "manual", "puzzel", "telstra", "twilio"), mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_enabled"), types.BoolValue(true))`                                                                                                                                                                                                                        |
| `portal.sso_idp_cert`           | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("auth"), types.StringValue("sso"))`                                                                                                                                                                                                                                                                                                                                     |
| `portal.sso_idp_sso_url`        | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("auth"), types.StringValue("sso"))`                                                                                                                                                                                                                                                                                                                                     |
| `portal.sso_issuer`             | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("auth"), types.StringValue("sso"))`                                                                                                                                                                                                                                                                                                                                     |
| `portal.telstra_client_id`      | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("telstra"))`                                                                                                                                                                                                                                                                                                                         |
| `portal.telstra_client_secret`  | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("telstra"))`                                                                                                                                                                                                                                                                                                                         |
| `portal.twilio_auth_token`      | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("telstra"))`                                                                                                                                                                                                                                                                                                                         |
| `portal.twilio_phone_number`    | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("twilio"))`                                                                                                                                                                                                                                                                                                                          |
| `portal.twilio_sid`             | not set                                                                                                          | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("twilio"))`                                                                                                                                                                                                                                                                                                                          |
| `vlan_ids`                      | `listvalidator.ValueStringsAre(stringvalidator.Any(mistvalidator.ParseInt(1, 4094), mistvalidator.ParseVar())),` | `listvalidator.ValueStringsAre(stringvalidator.Any(mistvalidator.ParseInt(1, 4094), mistvalidator.ParseVar())),mistvalidator.RequiredWhenValueIs(path.MatchRoot("vlan_pooling"), types.BoolValue(true))`                                                                                                                                                                                                                                          |
| `vlan_pooling`                  | not set                                                                                                          | `mistvalidator.CanOnlyTrueWhenValueIs(path.MatchRoot("vlan_enabled"), types.BoolValue(true))`                                                                                                                                                                                                                                                                                                                                                     |

#### Remove Attributes
| Attribute              | Reason                                                                                                                                         |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| `dynamic_psk.vlan_ids` | OpenAPI Specification issue. This attribute is not supported by the Mist API                                                                   |
| `portal_template_url`  | Read Only attribute returned by the Mist API. The returned URL has limited lifetime so it doesn't make sense to store it in the resource state |
| `thumbnail`            | Read Only attribute returned by the Mist API. The returned URL has limited lifetime so it doesn't make sense to store it in the resource state |

### Fixes
* [Issue 63](https://github.com/Juniper/terraform-provider-mist/issues/63): Adding `Optional` type to `alarmtemplate.rules.delivery`to fix synchronization issue