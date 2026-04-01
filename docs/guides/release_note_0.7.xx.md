---
subcategory: "Release Notes"
page_title: "v0.7.xx"
description: |-
    Release Notes for v0.7.xx
---

# Release Notes for v0.7.xx

## Release Notes for v0.7.1
**Release Date**: April 1st, 2026

### General Changes

#### Resources added

- **`mist_org_mxcluster` resource**: Manage Mist Edge clusters in your organization
- **`mist_org_nac_portal` resource**: Manage NAC (Network Access Control) portals for guest and employee authentication
- **`mist_org_nac_portal_template` resource**: Manage customizable NAC portal templates with branding and styling
- **`mist_org_nac_portal_image` resource**: Manage image assets for NAC portal customization

#### Attributes updated
- **`mist_org_mxedge` resource**:
  - Changed note attribute to notes

#### Improvements

- **`mist_org_mxedge` resource**:
  - Fixed site assignment and unassignment operations to properly handle all scenarios

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
  - `uses_description_from_port_usage`

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

