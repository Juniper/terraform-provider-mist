---
page_title: "1. EVPN Topology Overview"
subcategory: "How To"
description: |-
  Process to follow to create and manage an EVPN Topology with the Mist Provider.
---

# Create and manage an EVPN Topology with the Mist Provider.

Juniper Networks campus fabrics provide a single, standards-based Ethernet VPN-Virtual Extensible LAN (EVPN-VXLAN) solution that you can deploy on any campus. You can deploy campus fabrics on a two-tier network with a collapsed core or a campus-wide system that involves multiple buildings with separate distribution and core layers.

You can build and manage a campus fabric using the Mist Provider with the `mist_org_evpn_topology` resource. This topic describes the required steps to configure and manage EVPN Topology for the following architectures:

* **Campus Fabric IP Clos**: includes core / distribution / access layers
* **Campus Fabric Core-Distribution**: includes core / distribution layers to form a 2-tier EVPN topology where ESI-Lag is configured distribution to connect to access switches
* **EVPN Multihoming**: in a small/medium campus, EVPN Multihoming for Collapsed Core can be used where core switches are inter-connected to do EVPN

For more information about each EVPN Topology architecture, please refer to the [Juniper Documentation](https://www.juniper.net/documentation/us/en/software/mist/mist-wired/topics/concept/choose-campus-fabric-topology.html)

-> The EVPN Multihoming architecture can only be created at the site level with the `mist_site_evpn_topology` resource. The two others can be created at the site or at the org level with the `mist_site_evpn_topology` or `mist_org_evpn_topology` resources.