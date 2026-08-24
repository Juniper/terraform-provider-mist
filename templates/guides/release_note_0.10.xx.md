---
subcategory: "Release Notes"
page_title: "v0.10.xx"
description: |-
    Release Notes for v0.10.xx
---

# Release Notes for v0.10.xx

## Release Notes for v0.10.0
**Release Date**: August 2026

### General Changes

#### Attributes added

- **`mist_device_ap` resource**:
  - `enable_unii_4` — Enable UNII-4 (5.9 GHz) radio band support on the AP
  - `mqtt_config.default_topic` — Default MQTT topic used when no AssetFilter-specific topic is matched
  - `uwb_config` — UWB (Ultra-Wideband) configuration block (`enabled`, `host`, `port`, `slot`, `type`)

- **`mist_org_deviceprofile_ap` resource**:
  - `mqtt_config.default_topic` — Default MQTT topic used when no AssetFilter-specific topic is matched
  - `uwb_config` — UWB configuration block (`enabled`, `host`, `port`, `slot`, `type`)

- **`mist_org_rftemplate` resource**:
  - `enable_unii_4` — Enable UNII-4 (5.9 GHz) radio band support

- **`mist_site_setting` resource**:
  - `enable_unii_4` — Enable UNII-4 (5.9 GHz) radio band support
  - `gateway_mgmt.disable_idp_pcap` — Disable IDP packet capture on the gateway
  - `uwb_config` — UWB configuration block (`enabled`, `host`, `port`, `slot`, `type`)

- **`mist_device_gateway`, `mist_org_gatewaytemplate`, and `mist_org_deviceprofile_gateway` resources**:
  - `gateway_mgmt.disable_idp_pcap` — Disable IDP packet capture on the gateway
  - `networks.<key>.multicast` — Per-network multicast configuration block (`disable_igmp`, `enabled`, `groups`)
  - `port_config.<key>.wan_probe_override.hostnames` — List of hostnames used for WAN probe override health checks
  - `port_config.<key>.wan_probe_override.http` — HTTP-based WAN probe override configuration (`accepted_status_codes`, `urls`)
  - `routing_policies.<key>.terms[].action.next_policy` — When `true`, evaluation continues at the next routing policy instead of terminating
  - `routing_policies.<key>.terms[].action.next_term` — When `true`, evaluation continues at the next term within the current routing policy
  - `tunnel_configs.<key>.node0.internal_ip6s` — List of internal IPv6 addresses for tunnel node 0
  - `tunnel_configs.<key>.node0.probe_hostnames` — List of probe hostnames for tunnel node 0
  - `tunnel_configs.<key>.node0.probe_http` — HTTP probe configuration for tunnel node 0 (`accepted_status_codes`, `urls`)
  - `tunnel_configs.<key>.node0.probe_ip6s` — List of probe IPv6 addresses for tunnel node 0
  - `tunnel_configs.<key>.node1.internal_ip6s` — List of internal IPv6 addresses for tunnel node 1
  - `tunnel_configs.<key>.node1.probe_hostnames` — List of probe hostnames for tunnel node 1
  - `tunnel_configs.<key>.node1.probe_http` — HTTP probe configuration for tunnel node 1 (`accepted_status_codes`, `urls`)
  - `tunnel_configs.<key>.node1.probe_ip6s` — List of probe IPv6 addresses for tunnel node 1

- **`mist_org_network` resource**:
  - `multicast` — Multicast configuration block (`disable_igmp`, `enabled`, `groups.<key>.rp_ip`)

- **`mist_device_switch`, `mist_org_deviceprofile_switch`, `mist_org_networktemplate`, and `mist_site_networktemplate` resources**:
  - `acl_policies[].disabled` — Disable an individual ACL policy entry without removing it
  - `networks.<key>.multicast` — Per-network multicast configuration block (`enabled`, `igmp_version`)
  - `snmpv3_config.notify_filter[].categories` — List of SNMP notification category names to include in the filter entry
  - `vrf_instances.<key>.multicast_config` — VRF-level multicast configuration (`anycast_rp`, `rp_ip`, `sbd_subnet`, `sbd_vlan_id`)

- **`mist_org_networktemplate` resource**:
  - `multicast_config` — Top-level multicast configuration block for the network template (`anycast_rp`, `rp_ip`, `sbd_subnet`, `sbd_vlan_id`)

- **`mist_site_networktemplate` resource**:
  - `uwb_config` — UWB configuration block (`enabled`, `host`, `port`, `slot`, `type`)

- **`mist_org_webhook` and `mist_site_webhook` resources**:
  - `default_action` — Default action applied to webhook events that do not match any rule (`allow` or `drop`)
  - `rules` — Ordered list of webhook routing rules; each rule specifies an `action`, a `matching` expression, and the `topic` it applies to

- **`mist_org_nac_portal` resource**:
  - `enable_location` — Enable location-awareness in the NAC portal

- **`mist_org_psk` and `mist_site_psk` resources**:
  - `usermac_labels` — List of user MAC labels to associate with the PSK for policy matching

#### Attributes updated

- **`mist_org_psk` and `mist_site_psk` resources**:
  - `usage`: added `usermac_labels` as a valid enum value

- **`mist_org_wlan` and `mist_site_wlan` resources**:
  - `dynamic_psk.local_vlan_ids` — New attribute; list of local VLAN IDs available for Dynamic PSK assignment
