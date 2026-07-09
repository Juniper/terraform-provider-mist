---
subcategory: "Release Notes"
page_title: "v0.9.xx"
description: |-
    Release Notes for v0.9.xx
---

# Release Notes for v0.9.xx

## Release Notes for v0.9.1
**Release Date**: July 9th, 2026

### Bug Fixes

#### Resources fixed

- **`mist_org_mxedge` resource**:
  - Fixed `mxcluster_id` state drift: the attribute is now preserved across read operations when the API returns `null` for the cluster assignment, preventing spurious plan diffs.
  - Fixed `for_site` handling: site assignment is now performed via the dedicated `AssignOrgMxEdgeToSite` API call instead of being sent in the Create/Update request body (the API silently ignores `for_site`/`site_id` in those payloads); Read operations now fall back to `ListOrgMxEdges(for_site=any)` when the org-level `GetOrgMxEdge` returns 404 for a site-assigned device
  - Fixed device claiming: when `claim_code` is set, the resource now correctly retrieves the claimed device via the inventory claim response MAC address, applies any additional plan attributes (name, `mxcluster_id`), and handles duplicate-claim responses gracefully

### General Changes

#### Resources added

- **`mist_org_mxtunnel` resource**: New resource to manage Org-level Mist Tunnels, including IPsec settings, auto-preemption scheduling, VLAN assignments, and MxEdge cluster bindings

#### Attributes added

- **`mist_site_setting` resource**:
  - `mxtunnels`
  - `tunterm_monitoring`
  - `tunterm_monitoring_disabled`
  - `tunterm_multicast_config`

---

## Release Notes for v0.9.0
**Release Date**: July 6th, 2026

### Bug Fixes

#### Resources fixed

- **`mist_device_ap`, `mist_org_deviceprofile_ap`, and `mist_org_rftemplate` resources**:
  - Fixed `power` field validation in radio band schemas (`ap_radio_band24`, `ap_radio_band5`, `ap_radio_band6`, `rftemplate_radio_band24`, `rftemplate_radio_band5`, `rftemplate_radio_band6`): corrected minimum value from 3 (2.4 GHz) / 5 (5 GHz, 6 GHz) to `0` to match actual API behaviour; value is in dBm with range 0–25 for static power, or `null`/unset for auto power mode
  - Fixed `rftemplate_radio_band6.power_max` default and maximum from 18 to 17

### General Changes

#### Resources added

- **`mist_org_deviceprofile_switch` resource**: New resource to manage switch device profiles at the organization level

#### Attributes added

- **`mist_org_wlan` and `mist_site_wlan` resources**:
  - `enable_ftm` — Enable FTM (Fine-Time Measurement, 802.11mc); configures the AP as an FTM Responder to allow clients to perform ranging requests (default `false`)
  - `portal.smsglobal_sender` — Sender's number or sender ID for SMSGlobal portal SMS delivery

- **`mist_device_ap` and `mist_org_deviceprofile_ap` resources**:
  - `mqtt_config` — MQTT configuration block (`enabled`, `broker_host`, `broker_port`, `broker_proto`, `username`, `password`, `format`); BLE advertisements matching an AssetFilter `mqtt_topic` are forwarded to the configured broker

- **`mist_device_switch`, `mist_org_networktemplate`, and `mist_site_networktemplate` resources**:
  - `port_config.<key>.ae_lacp_passive` — Set LACP to passive mode on the AE interface (default `false`)

- **`mist_org_networktemplate`, `mist_site_networktemplate`, and `mist_site_setting` resources**:
  - `port_usages.<key>.server_fail_retry_interval` — Interval in seconds to retry 802.1X authentication after a RADIUS server failure (range 120–65535, default 120); only applicable when `port_auth` == `dot1x`

- **`mist_org_mxcluster` resource**:
  - `mist_nacedge` — NAC Edge survivability settings block (`enabled`, `caching_site_ids`, `nac_edge_hosts`, `auth_ttl`, `default_vlan`, `default_dot1x_vlan`); requires `mist_nac` to be enabled on the cluster

- **`mist_org_nac_portal` resource**:
  - `id` — Portal ID exposed directly from the API response
  - `org_id` — Organisation ID exposed directly from the API response

- **`mist_org_setting` resource**:
  - `api_policy.src_ips` — Optional list of allowed source IP addresses/CIDR subnets (max 10) for org API access
  - `marvis.disable_proactive_monitoring` — Disable Marvis proactive monitoring

- **`mist_org_sso` resource**:
  - `openroaming_ssids` — SSIDs for OpenRoaming (replaces the nested `openroaming` object)
  - `openroaming_wba_client_cert` — WBA client certificate for OpenRoaming (replaces the nested `openroaming` object)
  - `openroaming_wba_client_key` — WBA client key for OpenRoaming (replaces the nested `openroaming` object)

- **`mist_site_setting` resource**:
  - `iotproxy_visionline.cacerts` — PEM-encoded CA certificates to verify the Visionline collector's TLS certificate when using a self-signed certificate
  - `mist_nacedge.caching_site_ids` — List of site UUIDs whose auth requests should be cached by NAC Edges assigned to this site

#### Attributes updated

- **`mist_org_evpn_topology` and `mist_site_evpn_topology` resources**:
  - `evpn_topology_switch_role`: added `border` as a valid role value

- **`mist_org_mxedge` resource**:
  - `mxcluster_id` now accepts `null`

- **`mist_org_sso` resource**:
  - `openroaming` nested object is now deprecated; use the new flat fields `openroaming_ssids`, `openroaming_wba_client_cert`, and `openroaming_wba_client_key` instead
