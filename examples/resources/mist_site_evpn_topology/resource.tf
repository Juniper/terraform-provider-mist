resource "mist_site_site_evpn_topology" "evpn_one" {
  site_id = mist_site.terraform_test.id
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
