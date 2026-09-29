---
subcategory: "Release Notes"
page_title: "v0.11.xx"
description: |-
    Release Notes for v0.11.xx
---

# Release Notes for v0.11.xx

## Release Notes for v0.11.0
**Release Date**: September 2026

### Breaking Changes

- **`mist_org_setting` and `mist_site_setting` resources**: `synthetic_test.wan_speedtest.enabled` has been removed and replaced by `synthetic_test.wan_speedtest.disabled`, matching a rename made in the underlying Mist API. The new attribute has inverted polarity. Any configuration setting `enabled` will fail to plan after upgrading and must be updated manually:

  ```hcl
  # Before
  synthetic_test {
    wan_speedtest {
      enabled = true
    }
  }

  # After
  synthetic_test {
    wan_speedtest {
      disabled = false
    }
  }
  ```

  There is no automatic state migration for this attribute; update your configuration before running `terraform plan` on the upgraded provider.

- **`mist_site_evpn_topology` and `mist_org_evpn_topology` resources**: `switches.<mac>` no longer exposes `deviceprofile_id`, `downlink_ips`, `downlinks`, `esilaglinks`, `evpn_id`, `mac`, `model`, `router_id`, `site_id`, `suggested_downlinks`, `suggested_esilaglinks`, `suggested_uplinks`, and `uplinks`. These were pure computed values reported by the API; only `pod`, `pods`, and `role` remain. Remove any reference to the dropped attributes (for example in `output` blocks) before upgrading.

### General Changes

#### Attributes added

- **`mist_device_gateway`, `mist_org_gatewaytemplate` resources**:
  - `mnha_config.enabled` — Whether MNHA (Multi-Node HA) mode is enabled for this gateway (SRX only)

- **`mist_device_gateway`, `mist_org_gatewaytemplate`, `mist_org_deviceprofile_gateway`, and `mist_org_network` resources**:
  - `zone_id` (top-level on `mist_org_network`, per-network on the gateway resources) — SecurityZone the network belongs to; when set, the zone name is used as the security zone name on the SRX, and multiple networks can share the same zone

- **`mist_device_switch`, `mist_org_deviceprofile_switch`, `mist_org_networktemplate`, and `mist_site_networktemplate` resources**:
  - `vrf_instances.<key>.multicast_config.peg_enabled` — Enables the PIM EVPN Gateway on `is_l3_border` devices; required for external sources or receivers in EVPN topologies
  - `vrf_instances.<key>.multicast_config.rp_mac` — Device MAC address of a fabric RP in EVPN topologies
  - `vrf_instances.<key>.multicast_config.sbd_wan_rpf` — Builds an eBGP mesh between PEG borders over SBD IRBs so WAN-learned routes can satisfy the PIM RPF check during a border WAN-uplink failure
  - `port_usages.<key>.no_local_port_config` — Whether this port usage can be overridden in local port configuration

- **`mist_org_networktemplate` resource**:
  - `multicast_config.peg_enabled`, `multicast_config.rp_mac`, `multicast_config.sbd_wan_rpf` — Same as the `vrf_instances` fields above, applied to the top-level (master VRF) multicast configuration

- **`mist_site_setting` resource**:
  - `mist_nac_user_role_source` — Source of the Mist NAC user role sent to Juniper SRX gateways (`idp_role`, `radius_group`, or `none`)

- **`mist_org_setting` resource**:
  - `cacerts_configs` — Preferred per-issuer CA certificate configuration list (`cert`, `crl_enabled`, `crl_url`, `name`, `ocsp_enabled`, `ocsp_url`); when non-empty, takes precedence over `cacerts`
  - `enforce_src_ips_for_tokens` — When `true`, org API tokens without their own `src_ips` also respect the org policy `src_ips`
  - `enable_eap_md5_for_mab` — Enable EAP-MD5 for MAB (not FIPS compliant; use only for legacy device support)

#### Fixed

- **`mist_device_gateway`, `mist_org_deviceprofile_gateway`, and `mist_org_gatewaytemplate` resources**: `port_config.<key>.reth_nodes` was removed from the schema in a previous release, but stale references to it remained in the hand-written conversion code, causing errors when configuring these resources. The stale references have been removed.
