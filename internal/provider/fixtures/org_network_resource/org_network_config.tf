
  name                  = "test-org-network"
  subnet                = "192.168.100.0/24"
  subnet6               = "fd00:1000::/64"
  gateway               = "192.168.100.1"
  gateway6              = "fd00:1000::1"
  vlan_id               = "100"
  disallow_mist_services = false
  isolation             = false
  routed_for_networks   = ["192.168.200.0/24", "192.168.201.0/24"]

  internal_access = {
    enabled = true
  }

  internet_access = {
    enabled                      = true
    create_simple_service_policy = false
    restricted                   = false
  }

  multicast = {
    enabled      = true
    disable_igmp = false

    groups = {
      "225.1.0.3/32" = {
        rp_ip = "192.168.100.50"
      }
      "225.1.0.4/32" = {
        rp_ip = "192.168.100.51"
      }
    }
  }


  

