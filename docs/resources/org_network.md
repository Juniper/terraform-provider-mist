---
page_title: "mist_org_network Resource - terraform-provider-mist"
subcategory: "WAN Assurance"
description: |-
  This resource manages the WAN Assurance Networks.
  The Networks are used in the service_policies from the Gateway configuration, Gateway templates or HUB Profiles
---

# mist_org_network (Resource)

This resource manages the WAN Assurance Networks.

The Networks are used in the `service_policies` from the Gateway configuration, Gateway templates or HUB Profiles


## Example Usage

```terraform
resource "mist_org_network" "network_one" {
  org_id                 = mist_org.terraform_test.id
  name                   = "network_one"
  subnet                 = "10.4.0.0/24"
  disallow_mist_services = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `org_id` (String)
- `subnet` (String)

### Optional

- `disallow_mist_services` (Boolean) Whether to disallow Mist Devices in the network
- `gateway` (String)
- `gateway6` (String)
- `internal_access` (Attributes) (see [below for nested schema](#nestedatt--internal_access))
- `internet_access` (Attributes) Whether this network has direct internet access (see [below for nested schema](#nestedatt--internet_access))
- `isolation` (Boolean) Whether to allow clients in the network to talk to each other
- `multicast` (Attributes) Whether to enable multicast support (only PIM-sparse mode is supported) (see [below for nested schema](#nestedatt--multicast))
- `routed_for_networks` (List of String) For a Network (usually LAN), it can be routable to other networks (e.g. OSPF)
- `subnet6` (String)
- `tenants` (Attributes Map) Property key must be the user/tenant name (i.e. "printer-1") or a Variable (i.e. "{{myvar}}") (see [below for nested schema](#nestedatt--tenants))
- `vlan_id` (String)
- `vpn_access` (Attributes Map) Property key is the VPN name. Whether this network can be accessed from vpn (see [below for nested schema](#nestedatt--vpn_access))

### Read-Only

- `id` (String) Unique ID of the object instance in the Mist Organization

<a id="nestedatt--internal_access"></a>
### Nested Schema for `internal_access`

Optional:

- `enabled` (Boolean)


<a id="nestedatt--internet_access"></a>
### Nested Schema for `internet_access`

Optional:

- `create_simple_service_policy` (Boolean)
- `destination_nat` (Attributes Map) Property key can be an External IP (i.e. "63.16.0.3"), an External IP:Port (i.e. "63.16.0.3:443"), an External Port (i.e. ":443"), an External CIDR (i.e. "63.16.0.0/30"), an External CIDR:Port (i.e. "63.16.0.0/30:443") or a Variable (i.e. "{{myvar}}"). At least one of the `internal_ip` or `port` must be defined (see [below for nested schema](#nestedatt--internet_access--destination_nat))
- `enabled` (Boolean)
- `restricted` (Boolean) By default, all access is allowed, to only allow certain traffic, make `restricted`=`true` and define service_policies
- `static_nat` (Attributes Map) Property key may be an External IP Address (i.e. "63.16.0.3"), a CIDR (i.e. "63.16.0.12/20") or a Variable (i.e. "{{myvar}}") (see [below for nested schema](#nestedatt--internet_access--static_nat))

<a id="nestedatt--internet_access--destination_nat"></a>
### Nested Schema for `internet_access.destination_nat`

Optional:

- `internal_ip` (String) The Destination NAT destination IP Address. Must be an IP (i.e. "192.168.70.30") or a Variable (i.e. "{{myvar}}")
- `name` (String)
- `port` (String) The Destination NAT destination IP Address. Must be a Port (i.e. "443") or a Variable (i.e. "{{myvar}}")
- `wan_name` (String) SRX Only. If not set, we configure the nat policies against all WAN ports for simplicity


<a id="nestedatt--internet_access--static_nat"></a>
### Nested Schema for `internet_access.static_nat`

Required:

- `internal_ip` (String) The Static NAT destination IP Address. Must be an IP Address (i.e. "192.168.70.3") or a Variable (i.e. "{{myvar}}")
- `name` (String)

Optional:

- `wan_name` (String) SRX Only. If not set, we configure the nat policies against all WAN ports for simplicity. Can be a Variable (i.e. "{{myvar}}")



<a id="nestedatt--multicast"></a>
### Nested Schema for `multicast`

Optional:

- `disable_igmp` (Boolean) If the network will only be the source of the multicast traffic, IGMP can be disabled
- `enabled` (Boolean)
- `groups` (Attributes Map) Group address to RP (rendezvous point) mapping. Property Key is the CIDR (example "225.1.0.3/32") (see [below for nested schema](#nestedatt--multicast--groups))

<a id="nestedatt--multicast--groups"></a>
### Nested Schema for `multicast.groups`

Optional:

- `rp_ip` (String) RP (rendezvous point) IP Address



<a id="nestedatt--tenants"></a>
### Nested Schema for `tenants`

Optional:

- `addresses` (List of String)


<a id="nestedatt--vpn_access"></a>
### Nested Schema for `vpn_access`

Optional:

- `advertised_subnet` (String) If `routed`==`true`, whether to advertise an aggregated subnet toward HUB this is useful when there are multiple networks on SPOKE's side
- `allow_ping` (Boolean) Whether to allow ping from vpn into this routed network
- `destination_nat` (Attributes Map) Property key can be an External IP (i.e. "63.16.0.3"), an External IP:Port (i.e. "63.16.0.3:443"), an External Port (i.e. ":443"), an External CIDR (i.e. "63.16.0.0/30"), an External CIDR:Port (i.e. "63.16.0.0/30:443") or a Variable (i.e. "{{myvar}}"). At least one of the `internal_ip` or `port` must be defined (see [below for nested schema](#nestedatt--vpn_access--destination_nat))
- `nat_pool` (String) If `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub, a subnet is required to create and advertise the route to Hub
- `no_readvertise_to_lan_bgp` (Boolean) toward LAN-side BGP peers
- `no_readvertise_to_lan_ospf` (Boolean) toward LAN-side OSPF peers
- `no_readvertise_to_overlay` (Boolean) toward overlay, how HUB should deal with routes it received from Spokes
- `other_vrfs` (List of String) By default, the routes are only readvertised toward the same vrf on spoke. To allow it to be leaked to other vrfs
- `routed` (Boolean) Whether this network is routable
- `source_nat` (Attributes) If `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub (see [below for nested schema](#nestedatt--vpn_access--source_nat))
- `static_nat` (Attributes Map) Property key may be an External IP Address (i.e. "63.16.0.3"), a CIDR (i.e. "63.16.0.12/20") or a Variable (i.e. "{{myvar}}") (see [below for nested schema](#nestedatt--vpn_access--static_nat))
- `summarized_subnet` (String) toward overlay, how HUB should deal with routes it received from Spokes
- `summarized_subnet_to_lan_bgp` (String) toward LAN-side BGP peers
- `summarized_subnet_to_lan_ospf` (String) toward LAN-side OSPF peers

<a id="nestedatt--vpn_access--destination_nat"></a>
### Nested Schema for `vpn_access.destination_nat`

Optional:

- `internal_ip` (String) The Destination NAT destination IP Address. Must be an IP (i.e. "192.168.70.30") or a Variable (i.e. "{{myvar}}")
- `name` (String)
- `port` (String)


<a id="nestedatt--vpn_access--source_nat"></a>
### Nested Schema for `vpn_access.source_nat`

Optional:

- `external_ip` (String)


<a id="nestedatt--vpn_access--static_nat"></a>
### Nested Schema for `vpn_access.static_nat`

Required:

- `internal_ip` (String) The Static NAT destination IP Address. Must be an IP Address (i.e. "192.168.70.3") or a Variable (i.e. "{{myvar}}")
- `name` (String)



## Import
Using `terraform import`, import `mist_org_network` with:
```shell
# Org Network can be imported by specifying the org_id and the network_id
terraform import mist_org_network.network_one 17b46405-3a6d-4715-8bb4-6bb6d06f316a.d3c42998-9012-4859-9743-6b9bee475309
```