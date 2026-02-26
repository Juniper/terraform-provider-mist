
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

    destination_nat = {
      "203.0.113.10:8080" = {
        internal_ip = "192.168.100.10"
        name        = "web-server"
        port        = "80"
        wan_name    = "wan"
      }
    }

    static_nat = {
      "203.0.113.20" = {
        internal_ip = "192.168.100.20"
        name        = "mail-server"
        wan_name    = "wan"
      }
    }
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

  tenants = {
    "tenant1" = {
      addresses = ["192.168.100.30", "192.168.100.31"]
    }
  }

  vpn_access = {
    "vpn1" = {
      advertised_subnet              = "10.10.0.0/16"
      allow_ping                     = true
      nat_pool                       = "10.10.1.0/24"
      no_readvertise_to_lan_bgp      = true
      no_readvertise_to_lan_ospf     = true
      no_readvertise_to_overlay      = true
      routed                         = true
      summarized_subnet              = "10.10.0.0/16"
      summarized_subnet_to_lan_bgp   = "10.10.0.0/16"
      summarized_subnet_to_lan_ospf  = "10.10.0.0/16"
      other_vrfs                     = ["vrf2", "vrf3"]

      source_nat = {
        external_ip = "203.0.113.1"
      }

      destination_nat = {
        "203.0.113.30:443" = {
          internal_ip = "10.10.1.10"
          name        = "vpn-web-server"
          port        = "443"
        }
      }

      static_nat = {
        "203.0.113.40" = {
          internal_ip = "10.10.1.20"
          name        = "vpn-mail-server"
        }
      }
    }
  }
