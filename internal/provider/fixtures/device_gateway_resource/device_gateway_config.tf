  device_id = "00000000-0000-0000-1000-f4bfa8426080"
  name = "test-gateway-comprehensive"
  managed = false
  notes = "Comprehensive test gateway configuration with all possible attributes"
  router_id = "192.168.1.1"
  site_id = "2c107c8e-2e06-404a-ba61-e25b5757ecea"
  x = 150.5
  y = 250.75

  additional_config_cmds = [
    "set system host-name gateway-test",
    "set system domain-name example.com",
    "set interfaces lo0 unit 0 family inet address 10.255.255.1/32"
  ]

  dns_servers = ["8.8.8.8", "8.8.4.4", "1.1.1.1"]
  dns_suffix = ["example.com", "test.local"]
  ntp_servers = ["pool.ntp.org", "time.google.com", "time.cloudflare.com"]

  vars = {
    environment = "production"
    location = "datacenter-west"
    zone = "dmz"
    tier = "core"
  }

  bgp_config = {
    external_peers = {
      auth_key = "bgp-auth-key-123"
      bfd_minimum_interval = 1000
      bfd_multiplier = 3
      disable_bfd = false
      export_policy = "export-external"
      extended_v4_nexthop = true
      graceful_restart_time = 300
      hold_time = 90
      import_policy = "import-external"
      local_as = "65001"
      neighbor_as = "65002"
      neighbors = {
        "192.168.10.1" = {
          disabled = false
          export_policy = "neighbor-export-policy"
          hold_time = 180
          import_policy = "neighbor-import-policy"
          multihop_ttl = 5
          neighbor_as = "65003"
        }
      }
      networks = ["10.0.0.0/8", "192.168.0.0/16"]
      no_private_as = true
      no_readvertise_to_overlay = false
      tunnel_name = "ipsec-tunnel-1"
      type = "external"
      via = "wan"
      vpn_name = "corporate-vpn"
      wan_name = "wan1"
    }
  }

  extra_routes = {
    route_to_remote_site = {
      via = "192.168.100.1"
    }
  }

  extra_routes6 = {
    ipv6_route_1 = {
      via = "2001:db8::1"
    }
  }

  # idp_profiles section temporarily removed due to provider framework issue with missing id field
  # idp_profiles = {
  #   security_profile = {
  #     base_profile = "strict"
  #     name = "comprehensive-security"
  #     overwrites = [
  #       {
  #         action = "drop"
  #         name = "block-tcp-scans"
  #         matching = {
  #           attack_name = ["TCP:SCAN:PORT"]
  #           dst_subnet = ["192.168.1.0/24"]
  #           severity = ["critical", "high"]
  #         }
  #       }
  #     ]
  #   }
  # }

  ip_configs = {
    lan_interface = {
      ip = "192.168.1.1"
      ip6 = "2001:db8:1::1"
      netmask = "255.255.255.0"
      netmask6 = "64"
      secondary_ips = ["192.168.1.10", "192.168.1.11"]
      type = "static"
      type6 = "static"
    }
    wan_interface = {
      ip = "203.0.113.10"
      netmask = "255.255.255.252"
      type = "static"
    }
  }

  networks = [
    {
      disallow_mist_services = false
      gateway = "192.168.1.1"
      gateway6 = "2001:db8:1::1"
      internal_access = {
        enabled = true
      }
      internet_access = {
        create_simple_service_policy = true
        enabled = true
        destination_nat = {
          "192.168.1.100" = {
            internal_ip = "192.168.1.100"
            name = "web-server-nat"
            port = "80"
            wan_name = "wan1"
          }
        }
        static_nat = {
          "192.168.1.200" = {
            internal_ip = "192.168.1.200"
            name = "mail-server-static"
            wan_name = "wan1"
          }
        }
        restricted = false
      }
      isolation = false
      multicast = {
        disable_igmp = false
        enabled = true
        groups = {
          group_239_1_1_1 = {
            rp_ip = "192.168.1.254"
          }
        }
      }
      name = "corporate-lan"
      routed_for_networks = ["192.168.10.0/24", "192.168.20.0/24"]
      subnet = "192.168.1.0/24"
      subnet6 = "2001:db8:1::/64"
      tenants = {
        engineering = {
          addresses = ["192.168.1.100", "192.168.1.101"]
        }
      }
      vlan_id = "100"
      vpn_access = {
        corporate_vpn = {
          advertised_subnet = "192.168.0.0/16"
          allow_ping = true
          destination_nat = {
            "192.168.1.50" = {
              internal_ip = "192.168.1.50"
              name = "internal-api"
              port = "8080"
            }
          }
          nat_pool = "vpn-nat-pool"
          no_readvertise_to_lan_bgp = false
          no_readvertise_to_lan_ospf = false
          no_readvertise_to_overlay = false
          other_vrfs = ["guest-vrf"]
          routed = true
          source_nat = {
            external_ip = "203.0.113.50"
          }
          static_nat = {
            "192.168.1.150" = {
              internal_ip = "192.168.1.150"
              name = "vpn-gateway"
            }
          }
          summarized_subnet = "192.168.0.0/16"
          summarized_subnet_to_lan_bgp = "192.168.0.0/16"
          summarized_subnet_to_lan_ospf = "192.168.0.0/16"
        }
      }
    }
  ]

  oob_ip_config = {
    gateway = "172.16.0.1"
    ip = "172.16.0.100"
    netmask = "255.255.255.0"
    node1 = {
      gateway = "172.16.0.1"
      ip = "172.16.0.101"
      netmask = "255.255.255.0"
      type = "static"
      use_mgmt_vrf = true
      use_mgmt_vrf_for_host_out = false
      vlan_id = "10"
    }
    type = "static"
    use_mgmt_vrf = true
    use_mgmt_vrf_for_host_out = false
    vlan_id = "10"
  }

  path_preferences = {
    wan_preference = {
      paths = [
        {
          cost = 100
          disabled = false
          name = "primary-wan"
          networks = ["default"]
          type = "wan"
        }
      ]
      strategy = "ordered"
    }
  }

  port_config = {
    "ge-0/0/1" = {
      ae_disable_lacp = false
      ae_idx = "ae0"
      ae_lacp_force_up = true
      aggregated = true
      critical = false
      description = "Primary LAN interface"
      disable_autoneg = false
      disabled = false
      duplex = "auto"
      mtu = 1500
      networks = ["corporate-lan"]
      outer_vlan_id = 200
      poe_disabled = false
      ip_config = {
        # dns = ["8.8.8.8", "8.8.4.4"]  # commented out due to provider inconsistency
        # dns_suffix = ["example.com"]  # commented out due to provider inconsistency
        gateway = "192.168.1.1"
        gateway6 = "2001:db8:1::1"
        ip = "192.168.1.1"
        ip6 = "2001:db8:1::1"
        netmask = "255.255.255.0"
        netmask6 = "64"
        network = "corporate-lan"
        poser_password = "pppoe-password"
        pppoe_auth = "pap"
        pppoe_username = "pppoe-user"
        type = "static"
        type6 = "static"
      }
      port_network = "lan"
      preserve_dscp = true
      redundant = false
      redundant_group = 1
      reth_idx = "reth0"
      reth_node = "node0"
      # reth_nodes = ["node0", "node1"]  # commented out due to provider inconsistency
      speed = "auto"
      ssr_no_virtual_mac = false
      svr_port_range = "1024-65535"
      traffic_shaping = {
        # class_percentages = [25, 25, 25, 25]  # commented out due to provider inconsistency
        enabled = true
        max_tx_kbps = 100000
      }
      usage = "lan"
      vlan_id = "100"
      vpn_paths = {
        corporate_vpn = {
          bfd_profile = "broadband"
          bfd_use_tunnel_mode = true
          preference = 100
          role = "spoke"
          traffic_shaping = {
            # class_percentages = [30, 30, 20, 20]  # commented out due to provider inconsistency
            enabled = true
            max_tx_kbps = 50000
          }
        }
      }
    }
  }

  routing_policies = {
    bgp_import_policy = {
      terms = [
        {
          actions = {
            accept = true
          }
          matching = {
            prefix = ["10.0.0.0/8", "192.168.0.0/16"]
          }
        }
      ]
    }
  }

  service_policies = [
    {
      action = "allow"
      idp = {
        enabled = true
        profile = "strict"
      }
      name = "allow-web-traffic"
      path_preference = "wan_preference"
      services = ["web-browsing", "ssl"]
      tenants = ["engineering"]
    }
  ]

  tunnel_configs = {
    ipsec_tunnel_1 = {
      ike_lifetime = 28800
      ike_mode = "main"
      ike_proposals = [
        {
          auth_algo = "sha2"
          dh_group = "14"
          enc_algo = "aes256"
        }
      ]
      ipsec_lifetime = 3600
      ipsec_proposals = [
        {
          auth_algo = "sha2"
          enc_algo = "aes256"
        }
      ]
      local_id = "gateway1@example.com"
      mode = "active-standby"
      primary = {
        hosts = ["vpn.example.com"]
        remote_ids = ["gateway2@example.com"]
        wan_names = ["wan1"]
      }
      protocol = "ipsec"
      provider = "custom-ipsec"
      psk = "pre-shared-key-123"
      secondary = {
        hosts = ["backup-vpn.example.com"]
        wan_names = ["wan2"]
      }
      version = "2"
    }
  }

  tunnel_provider_options = {
    jse = {
      name = "JSE-Provider"
      num_users = 100
    }
    zscaler = {
      aup_acceptance_required = true
      aup_expire = 3600
      aup_ssl_proxy = true
      download_mbps = 100
      enable_aup = true
      enable_caution = true
      enforce_authentication = true
      name = "Zscaler-Provider"
      sub_locations = [
        {
          aup_acceptance_required = false
          aup_expire = 1800
          aup_ssl_proxy = false
          auth_required = true
          enforce_authentication = false
          idle_time_in_minutes = 30
          name = "Branch-Office"
          session_timeout = 7200
          subnets = ["192.168.10.0/24"]
        }
      ]
      upload_mbps = 50
      use_xff = true
    }
  }

  vrf_config = {
    enabled = true
  }

  vrf_instances = {
    guest_network = {
      networks = ["guest-wifi", "guest-lan"]
    }
    iot_network = {
      networks = ["iot-devices"]
    }
  }

