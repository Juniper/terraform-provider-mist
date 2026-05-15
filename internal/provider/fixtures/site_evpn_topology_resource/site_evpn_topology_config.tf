name = "test_site_evpn_topology"

switches "000000000001" {
  role = "distribution"
  pod  = 1
}

switches "000000000002" {
  role = "access"
  pod  = 1
}

evpn_options {
  auto_loopback_subnet   = "100.101.0.0/16"
  auto_loopback_subnet6  = "fd01::/48"
  auto_router_id_subnet  = "100.100.0.0/24"
  auto_router_id_subnet6 = "fd02::/48"
  core_as_border         = false
  enable_inband_mgmt     = true
  enable_inband_ztp      = false
  per_vlan_vga_v4_mac    = false
  per_vlan_vga_v6_mac    = false
  routed_at              = "core"

  overlay {
    as = 65000
  }

  underlay {
    as_base          = 65100
    routed_id_prefix = "/26"
    subnet           = "10.255.240.0/20"
    use_ipv6         = false
  }
}

pod_names = {
  "1" = "pod1"
}
