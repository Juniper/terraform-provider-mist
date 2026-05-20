---
subcategory: "Release Notes"
page_title: "v0.8.xx"
description: |-
    Release Notes for v0.8.xx
---

# Release Notes for v0.8.xx

## Release Notes for v0.8.0
**Release Date**: May 20th, 2026

### Bug Fixes

#### Resources fixed

- **`mist_org_nacidp` resource**:
  - Fixed provider inconsistency error ("was `okta.com`, but now `null`") caused by `oauth_provider_domain` being missing from the SDK↔Terraform conversion — the field is now properly read from the API response and written to state

- **`mist_org_sso` resource**:
  - Same fix as `mist_org_nacidp`: `oauth_provider_domain` was missing from both conversion functions and is now handled correctly

- **`mist_org_networktemplate` resource**:
  - Fixed provider error "Missing PortConfigValue Attribute Value" for `ae_lacp_force_up` in `switch_matching.rules.port_config` — field was in the schema but absent from the switch-matching conversion code

- **`mist_org_deviceprofile_gateway` resource**:
  - Fixed `oob_ip_config.vlan_id` not being read from API response

### General Changes

#### Attributes added

- **`mist_device_ap` and `mist_org_deviceprofile_ap` resources**:
  - `zigbee_config` — Zigbee radio configuration (enabled, channel, allow_join, pan_id, extended_pan_id)
  - `mesh.use_wpa3_on_5` — Enable WPA3 on the 5 GHz mesh link

- **`mist_device_switch` resource**:
  - `port_config.<key>.ae_lacp_force_up` — Force up an LACP aggregate interface even when no LACP PDUs are received
  - `port_usages.<key>.poe_keep_state_when_reboot` — Keep PoE port state across switch reboots
  - `port_config_overwrite.<key>.poe_keep_state_when_reboot` — Keep PoE port state across switch reboots (overwrite)

- **`mist_device_gateway` resource**:
  - `port_config.<key>.poe_keep_state_when_reboot` — Keep PoE port state across gateway reboots
  - `bgp_config.<key>.neighbors.<key>.tunnel_via` — VPN tunnel used for this BGP neighbour

- **`mist_org_evpn_topology` and `mist_site_evpn_topology` resources**:
  - `evpn_options.enable_inband_mgmt` — Enable in-band management traffic for the EVPN topology

- **`mist_org_gatewaytemplate` resource**:
  - `gateway_mgmt` — Gateway management settings block (console, probe, auto-signature update, root password, admin accounts, TACACS, etc.)
  - `port_config.<key>.poe_keep_state_when_reboot` — Keep PoE port state across reboots
  - `bgp_config.<key>.neighbors.<key>.tunnel_via` — VPN tunnel used for BGP neighbour

- **`mist_org_networktemplate` and `mist_site_networktemplate` resources**:
  - `switch_matching.rules.<key>.port_config.<key>.ae_lacp_force_up` — Force up an LACP aggregate interface
  - `port_usages.<key>.poe_keep_state_when_reboot` — Keep PoE port state across switch reboots

- **`mist_org_psk` and `mist_site_psk` resources**:
  - `vlan_name` — Named VLAN reference for the PSK

- **`mist_org_setting` resource**:
  - `mist_nac.allow_teap_machine_auth_only` — Restrict TEAP authentication to machine certificates only
  - `mist_nac.mdm.coa_type` — CoA (Change of Authorization) type for MDM-integrated NAC
  - `marvis.self_driving` — Enable Marvis self-driving network automation

- **`mist_org_wlan` and `mist_site_wlan` resources**:
  - `auth.enable_gcmp256` — Enable GCMP-256 cipher suite
  - `auth.enable_beacon_protection` — Enable 802.11w beacon protection (Management Frame Protection for beacons)

- **`mist_site_setting` resource**:
  - `allow_mist` — Allow Mist support access to the site
  - `ap_synthetic_test` — AP synthetic test configuration (additional_vlan_ids)
  - `iotproxy` — IoT proxy configuration including Visionline integration
  - `vars_annotations` — Named variable annotations map for tunnel and network references
