---
page_title: "EVPN Multihoming"
description: |-
  Process to follow to create and manage an EVPN Multihoming EVPN Topology with the Mist Provider.
---

# Create and manage an EVPN Multihoming EVPN Topology with the Mist Provider.

## 1. Build the Topology
Use this `mist_org_evpn_topology` resource to create the EVPN Topology.

```terraform
resource "mist_org_site_evpn_topology" "evpn_one" {
  name = "evpn_one"
  evpn_options = {
    routed_at = "edge"
    overlay = {
      as = 65000
    }
    core_as_border      = true
    per_vlan_vga_v4_mac = false
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
      role = "collapsed-core"
    },
    {
      mac  = "02000400002"
      role = "collapsed-core"
    },
    {
      mac  = "02000400003"
      role = "esilag-access"
    },
    {
      mac  = "02000400004"
      role = "esilag-access"
    }
  ]
}
```

## 2. Update the device port configs
Create or Update the `mist_site_networktemplate` resource to add a `port_usage` that will be used for the ESI-LAG link between the EVPN topology and the rest of the network.

-> The example below is using the port profile name "x-esilag". This name can be customized and must be used when assigning the ESI-LAG profile the the switch ports in the step 3.

```terraform
resource "mist_site_networktemplate" "networktemplate_one" {
  site_id = mist_site.terraform_test.id
  port_usages = {
    x-esilag = {
      mode         = "trunk"
      disabled     = false
      port_network = null
      voip_network = null
      stp_edge     = false
      all_networks = false
      networks = [
        "user"
      ]
      port_auth     = null
      speed         = "auto"
      duplex        = "auto"
      mac_limit     = "0"
      poe_disabled  = true
      enable_qos    = false
      storm_control = {}
      mtu           = "9200"
    }
  }
}
```

## 3. Update the device port configs
Update the `mist_device_switch` resources to configure the `port_usages` and assign the corresponding profile to the ports you have used (or plan to use) to inter-connect the switches by assigning them the `evpn_uplink` or `evpn_downlink` profiles:

### Core Switches
  * Links to other Cores with 1 x `evpn_downlink` and 1 x `evpn_uplink`
  * Links to each Distribution switch with `usage`=`<esilag name>`, `aggregated` and `esilag` to `true` and `ae_idx` configured

```terraform
resource "mist_device_switch" "switch_core_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "020004000001").id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "020004000001").id.site_id
  name      = "core-01"
  port_config = {
    "ge-0/0/0" = {
      usage = "evpn_uplink"
    },
    "ge-0/0/1" = {
      usage = "evpn_downlink"
    },
    "ge-0/0/2" = {
      usage      = "x-esilag"
      aggregated = true
      esilag     = true
      ae_idx     = 0
    },
    "ge-0/0/3" = {
      usage      = "x-esilag"
      aggregated = true
      esilag     = true
      ae_idx     = 1
    }
  }
}
```

### Distribution Switches
  * Links to the Core layer configured with `usage`=`<esilag name>`, `aggregated` and `esilag` to `true` and `ae_idx` configured
```terraform
resource "mist_device_switch" "switch_distri_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "020004000003").id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "020004000001").id.site_id
  name      = "distri-01"
  port_config = {
    "ge-0/0/0,ge-0/0/1" = {
      usage      = "x-esilag"
      aggregated = true
      esilag     = true
      ae_idx     = 0
    }
  }
}
```
