---
page_title: "mist_org_vpn Resource - terraform-provider-mist"
subcategory: "WAN Assurance"
description: |-
  This resource manages the Org VPN.
---

# mist_org_vpn (Resource)

This resource manages the Org VPN.


## Example Usage

```terraform
resource "mist_org_setting" "terraform_test" {
  org_id = mist_org.terraform_test.id
  password_policy = {
    enabled                  = true
    min_length               = 8
    requires_special_char    = true
    requires_two_factor_auth = true
  }
  mist_nac = {
    eu_only = true
  }
  synthetic_test = {
    disabled = false
    vlans = [
      {
        vlan_ids = [
          "8",
          "999"
        ],
        disabled = true
      }
    ]
  }
  api_policy = {
    no_reveal = false
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `paths` (Attributes Map) For `type`==`hub_spoke`, Property key is the VPN name. For `type`==`mesh`, Property key is the Interface name (see [below for nested schema](#nestedatt--paths))

### Optional

- `org_id` (String)
- `path_selection` (Attributes) Only if `type`==`hub_spoke` (see [below for nested schema](#nestedatt--path_selection))
- `type` (String) enum: `hub_spoke`, `mesh`

### Read-Only

- `id` (String) Unique ID of the object instance in the Mist Organization

<a id="nestedatt--paths"></a>
### Nested Schema for `paths`

Optional:

- `bfd_profile` (String) enum: `broadband`, `lte`
- `bfd_use_tunnel_mode` (Boolean) If `type`==`mesh` and for SSR only, whether to use tunnel mode
- `ip` (String) If different from the wan port
- `peer_paths` (Attributes Map) If `type`==`mesh`, Property key is the Peer Interface name (see [below for nested schema](#nestedatt--paths--peer_paths))
- `pod` (Number)
- `traffic_shaping` (Attributes) (see [below for nested schema](#nestedatt--paths--traffic_shaping))

<a id="nestedatt--paths--peer_paths"></a>
### Nested Schema for `paths.peer_paths`

Optional:

- `preference` (Number)


<a id="nestedatt--paths--traffic_shaping"></a>
### Nested Schema for `paths.traffic_shaping`

Optional:

- `class_percentage` (List of Number) percentages for different class of traffic: high / medium / low / best-effort adding up to 100
- `enabled` (Boolean)
- `max_tx_kbps` (Number)



<a id="nestedatt--path_selection"></a>
### Nested Schema for `path_selection`

Optional:

- `strategy` (String) enum: `disabled`, `simple`, `manual`



## Import
Using `terraform import`, import `mist_org_vpn` with:
```shell
# Org VPN can be imported by specifying the org_id and the vpn_id
terraform import mist_org_vpn.vpn_one 17b46405-3a6d-4715-8bb4-6bb6d06f316a.d3c42998-9012-4859-9743-6b9bee475309
```