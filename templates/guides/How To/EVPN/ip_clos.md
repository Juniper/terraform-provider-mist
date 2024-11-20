---
page_title: "Campus Fabric IP Clos"
description: |-
  Process to follow to create and manage a Campus Fabric IP Clos EVPN Topology with the Mist Provider.
---

# Create and manage a Campus Fabric IP Clos EVPN Topology with the Mist Provider.

## 1. Build the Topology
Use this `mist_org_evpn_topology` resource to create the EVPN Topology. 

```terraform
resource "mist_site_evpn_topology" "evpn_one" {
  site_id = mist_site.site_evpn.id
  name    = "evpn_one"
  evpn_options = {
    routed_at = "core"
    overlay = {
      as = 65000
    },
    core_as_border        = true
    auto_loopback_subnet  = "172.16.192.0/24"
    auto_loopback_subnet6 = "fd33:ab00:2::/64"
    per_vlan_vga_v4_mac   = false
    underlay = {
      as_base  = 65001
      use_ipv6 = false
      subnet   = "10.255.240.0/20"
    }
    auto_router_id_subnet = "172.16.254.0/23"
  }
  switches = [
    {
      mac  = "020004000001"
      role = "core"
    },
    {
      mac  = "02000400002"
      role = "core"
    },
    {
      mac  = "02000400003"
      role = "distribution"
    },
    {
      mac  = "02000400004"
      role = "distribution"
    },
    {
      mac  = "02000400005"
      role = "access"
    },
    {
      mac  = "02000400006"
      role = "access"
    }
  ]
}

```

## 2. Update the device port configs
Update the `mist_device_switch` resources to configure the `port_usages` and assign the corresponding profile to the ports you have used (or plan to use) to inter-connect the switches by assigning them the `evpn_uplink` or `evpn_downlink` profiles:

### Core Switches
  * Links to the Distribution layer configured with `usage`=`evpn_downlink`

```terraform
resource "mist_device_switch" "switch_core_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "020004000001").id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "020004000001").id.site_id
  name      = "core-01"
  port_config = {
    "ge-0/0/0,ge-0/0/1" = {
      usage = "evpn_downlink"
    }
  }
}
```


### Distribution Switches
  * Links to the Core layer configured with `usage`=`evpn_uplink`
  * Links to the Access layer configured with `usage`=`evpn_downlink`

```terraform
resource "mist_device_switch" "switch_distri_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "020004000003").id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "020004000003").id.site_id
  name      = "distri-01"
  port_config = {
    "ge-0/0/0,ge-0/0/1" = {
      usage = "evpn_uplink"
    },
    "ge-0/0/2,ge-0/0/3" = {
      usage = "evpn_downlink"
    }
  }
}
```


### Access Switches
  * Links to the Distribution layer configured with `usage`=`evpn_uplink`

```terraform
resource "mist_device_switch" "switch_access_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "020004000005").id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "020004000005").id.site_id
  name      = "access-01"
  port_config = {
    "ge-0/0/0,ge-0/0/1" = {
      usage = "evpn_uplink"
    }
  }
}
```
