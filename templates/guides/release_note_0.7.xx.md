---
subcategory: "Release Notes"
page_title: "v0.7.xx"
description: |-
    Release Notes for v0.7.xx
---

# Release Notes for v0.7.xx

## Release Notes for v0.7.0
**Release Date**: March 3rd, 2026


### General Changes

#### Resources added

- **`mist_org_mxedge` resource**: Manage Mist Edge instances in your organization

#### Attributes added

- **`mist_site` resource**:
  - `routertemplate_id`

- **`mist_org_setting` resource**:
  - `allow_mist`

- **`mist_site_setting` resource**:
  - `allow_mist`
  - `gateway_tunnel_updown_threshold`
  - `ap_synthetic_test.additional_vlan_ids`

- **`mist_org_wlan` and `mist_site_wlan` resources**:
  - `disable_message_authenticator_check`

#### Attributes updated

- **Map schema** (affects site resources with maps):
  - Replaced `mapstack_id` and `mapstack_floor` with `group_name` and `group_idx` for better maps grouping (group_idx typically used for floor)

- **Extra routes schema** (affects gateway and network template resources):
  - Updated `extra_route` and `extra_route6` to support ECMP (Equal-Cost Multi-Path) load balancing with shared `next_hop_via` schema

#### Attributes deprecated

- **Various resources**:
  - `managed` - marked as deprecated
  - `disable_auto_config` - marked as deprecated

**Deprecated Attributes:**
The `managed` and `disable_auto_config` attributes are now deprecated in favour of `mist_configured`. While they continue to work in this release, consider planning for their removal in future versions.

