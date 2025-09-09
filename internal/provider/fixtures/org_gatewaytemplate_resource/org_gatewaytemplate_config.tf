
  name   = "test_gateway_template"
  type   = "standalone"

  additional_config_cmds = [
    "set system host-name test-gateway",
    "set system time-zone UTC"
  ]

  bgp_config = {
    "test_bgp" = {
      via                        = "lan"
      type                       = "external"
      local_as                   = "65001"
      neighbor_as                = "65002"
      auth_key                   = "test-auth-key"
      bfd_minimum_interval       = 1000
      bfd_multiplier             = 3
      disable_bfd                = false
      export                     = "test-export"
      export_policy              = "test-export-policy"
      extended_v4_nexthop        = false
      graceful_restart_time      = 120
      hold_time                  = 90
      import                     = "test-import"
      import_policy              = "test-import-policy"
      no_private_as              = false
      no_readvertise_to_overlay  = false
      tunnel_name                = "test-tunnel"
      wan_name                   = "wan0"
      vpn_name                   = "test-vpn"
      networks = [
        "192.168.1.0/24",
        "192.168.2.0/24"
      ]
      neighbors = {
        "192.168.1.1" = {
          neighbor_as     = "65002"
          disabled        = false
          export_policy   = "neighbor-export"
          hold_time       = 180
          import_policy   = "neighbor-import"
          multihop_ttl    = 2
        }
      }
    }
  }

  dhcpd_config = {
    enabled = true
    config = {
      "lan" = {
        type                = "local"
        ip_start            = "192.168.1.100"
        ip_end              = "192.168.1.200"
        gateway             = "192.168.1.1"
        lease_time          = 86400
        server_id_override  = false
        dns_servers = [
          "8.8.8.8",
          "8.8.4.4"
        ]
        dns_suffix = [
          "example.com"
        ]
        fixed_bindings = {
          "001122334455" = {
            ip   = "192.168.1.50"
            name = "test-device"
          }
        }
        options = {
          "15" = {
            type  = "string"
            value = "example.com"
          }
        }
        vendor_encapsulated = {
          "311:1" = {
            type  = "string"
            value = "test-vendor"
          }
        }
      }
    }
  }

  dns_override = false
  dns_servers = [
    "8.8.8.8",
    "1.1.1.1"
  ]
  dns_suffix = [
    "test.com",
    "example.org"
  ]

  extra_routes = {
    "10.0.0.0/8" = {
      via = "192.168.1.1"
    }
  }

  extra_routes6 = {
    "2001:db8::/32" = {
      via = "2001:db8::1"
    }
  }

  idp_profiles = {
    "test_profile" = {
      base_profile = "standard"
      name         = "Test IDP Profile"
      org_id       = "{{org_id}}"
      overwrites = [
        {
          action = "drop"
          name   = "test-overwrite"
          matching = {
            attack_name = ["test-attack"]
            dst_subnet  = ["192.168.1.0/24"]
            severity    = ["high"]
          }
        }
      ]
    }
  }

  ip_configs = {
    "lan" = {
      type           = "static"
      ip             = "192.168.1.1"
      netmask        = "/24"
      secondary_ips  = ["192.168.1.2/24"]
    }
  }

  networks = [
    {
      name                      = "lan"
      vlan_id                   = "10"
      subnet                    = "192.168.1.0/24"
      gateway                   = "192.168.1.1"
      subnet6                   = "2001:db8:1::/64"
      gateway6                  = "2001:db8:1::1"
      isolation                 = false
      disallow_mist_services    = false
      routed_for_networks       = ["wan"]
      internal_access = {
        enabled = true
      }
      internet_access = {
        enabled                        = true
        restricted                     = false
        create_simple_service_policy   = false
        destination_nat = {
          "192.168.1.100" = {
            internal_ip = "10.0.1.100"
            port        = "80"
            name        = "web-server"
            wan_name    = "wan0"
          }
        }
        static_nat = {
          "203.0.113.10" = {
            internal_ip = "192.168.1.10"
            name        = "mail-server"
            wan_name    = "wan0"
          }
        }
      }
      multicast = {
        enabled       = true
        disable_igmp  = false
        groups = {
          "225.1.0.3/32" = {
            rp_ip = "192.168.1.10"
          }
        }
      }
      tenants = {
        "printer-1" = {
          addresses = ["192.168.1.201"]
        }
      }
      vpn_access = {
        "test-vpn" = {
          routed                               = true
          allow_ping                           = true
          advertised_subnet                    = "192.168.0.0/16"
          nat_pool                             = "10.10.1.0/24"
          no_readvertise_to_lan_bgp            = false
          no_readvertise_to_lan_ospf           = false
          no_readvertise_to_overlay            = false
          other_vrfs                           = ["guest"]
          summarized_subnet                    = "192.168.0.0/16"
          summarized_subnet_to_lan_bgp         = "192.168.0.0/16"
          summarized_subnet_to_lan_ospf        = "192.168.0.0/16"
          source_nat = {
            external_ip = "203.0.113.20"
          }
          destination_nat = {
            "192.168.1.110" = {
              internal_ip = "10.0.1.110"
              port        = "443"
              name        = "secure-server"
            }
          }
          static_nat = {
            "203.0.113.30" = {
              internal_ip = "192.168.1.30"
              name        = "ftp-server"
            }
          }
        }
      }
    }
  ]

  ntp_override = false
  ntp_servers = [
    "pool.ntp.org",
    "time.google.com"
  ]

  oob_ip_config = {
    type                        = "static"
    ip                          = "10.1.1.10"
    netmask                     = "/24"
    gateway                     = "10.1.1.1"
    use_mgmt_vrf                = true
    use_mgmt_vrf_for_host_out   = true
    vlan_id                     = "100"
    node1 = {
      type                        = "static"
      ip                          = "10.1.1.11"
      netmask                     = "/24"
      gateway                     = "10.1.1.1"
      use_mgmt_vrf                = true
      use_mgmt_vrf_for_host_out   = true
      vlan_id                     = "100"
    }
  }

  path_preferences = {
    "internet-path" = {
      strategy = "ordered"
      paths = [
        {
          type           = "wan"
          name           = "wan0"
          cost           = 10
          disabled       = false
        },
        {
          type           = "vpn"
          name           = "test-vpn"
          cost           = 20
          disabled       = false
          internet_access = true
          wan_name       = "wan1"
        },
        {
          type         = "local"
          gateway_ip   = "192.168.1.1"
          cost         = 30
          disabled     = false
          networks     = ["192.168.1.0/24"]
          target_ips   = ["192.168.1.10"]
        }
      ]
    }
  }

  port_config = {
    "ge-0/0/0" = {
      usage                    = "wan"
      name                     = "wan0"
      wan_type                 = "broadband"
      disabled                 = false
      critical                 = true
      description              = "WAN Port 0"
      disable_autoneg          = false
      duplex                   = "auto"
      speed                    = "auto"
      mtu                      = 1500
      aggregated               = false
      ae_disable_lacp          = false
      ae_idx                   = "0"
      ae_lacp_force_up         = false
      outer_vlan_id            = 100
      vlan_id                  = "10"
      poe_disabled             = false
      preserve_dscp            = false
      redundant                = false
      redundant_group          = 1
      reth_idx                 = "0"
      reth_node                = "node0"
      reth_nodes               = ["node0", "node1"]
      ssr_no_virtual_mac       = false
      svr_port_range           = "16384-32767"
      wan_disable_speedtest    = false
      wan_ext_ip               = "203.0.113.1"
      wan_arp_policer          = "recommended"
      ip_config = {
        type            = "static"
        ip              = "203.0.113.1"
        netmask         = "/30"
        gateway         = "203.0.113.2"
        gateway6        = "2001:db8::1"
        ip6             = "2001:db8::2"
        netmask6        = "/64"
        network         = "wan"
        type6           = "static"
        pppoe_auth      = "pap"
        pppoe_username  = "test-user"
        poser_password  = "test-password"
        dns = [
          "8.8.8.8"
        ]
        dns_suffix = [
          "example.com"
        ]
      }
      wan_extra_routes = {
        "10.10.0.0/16" = {
          via = "203.0.113.2"
        }
      }
      wan_extra_routes6 = {
        "2001:db8:2::/64" = {
          via = "2001:db8::1"
        }
      }
      wan_networks = ["guest"]
      wan_probe_override = {
        ips             = ["8.8.8.8"]
        ip6s            = ["2001:4860:4860::8888"]
        probe_profile   = "broadband"
      }
      wan_source_nat = {
        disabled  = false
        nat_pool  = "203.0.113.0/30"
      }
      traffic_shaping = {
        enabled         = true
        max_tx_kbps     = 100000
        class_percentages = [25, 25, 25, 25]
      }
      vpn_paths = {
        "test-vpn" = {
          role                  = "spoke"
          bfd_profile           = "broadband"
          bfd_use_tunnel_mode   = false
          preference            = 100
          traffic_shaping = {
            enabled         = true
            max_tx_kbps     = 50000
            class_percentages = [30, 30, 20, 20]
          }
        }
      }
    }
    "ge-0/0/1" = {
      usage     = "lan"
      disabled  = false
      networks  = ["lan"]
      port_network = "lan"
      description = "LAN Port 1"
    }
  }

  router_id = "1.1.1.1"

  routing_policies = {
    "test-policy" = {
      terms = [
        {
          actions = {
            accept           = true
            add_community    = ["65001:100"]
            community        = ["65001:200"]
            local_preference = "100"
            prepend_as_path  = ["65001"]
            add_target_vrfs  = ["guest"]
            exclude_as_path  = ["65003"]
            exclude_community = ["65001:300"]
            export_communities = ["65001:400"]
          }
          matching = {
            as_path     = ["^65001.*"]
            community   = ["65001:100"]
            network     = ["192.168.0.0/16"]
            prefix      = ["192.168.1.0/24"]
            protocol    = ["bgp", "static"]
            vpn_path    = ["test-vpn"]
            vpn_neighbor_mac = ["00:11:22:33:44:55"]
            route_exists = {
              route    = "0.0.0.0/0"
              vrf_name = "default"
            }
            vpn_path_sla = {
              max_jitter  = 10
              max_latency = 100
              max_loss    = 1
            }
          }
        }
      ]
    }
  }

  service_policies = [
    {
      name         = "internet-access"
      action       = "allow"
      path_preference = "internet-path"
      local_routing = false
      services = [
        "http",
        "https",
        "dns"
      ]
      tenants = [
        "lan"
      ]
      idp = {
        enabled       = true
        alert_only    = false
        profile       = "strict"
        idpprofile_id = "test-idp-profile"
      }
      antivirus = {
        enabled      = true
        profile      = "default"
        avprofile_id = "test-av-profile"
      }
      appqoe = {
        enabled = true
      }
      ssl_proxy = {
        enabled          = true
        ciphers_category = "strong"
      }
      ewf = [
        {
          enabled       = true
          alert_only    = false
          profile       = "strict"
          block_message = "Access Denied"
        }
      ]
    }
  ]

  tunnel_configs = {
    "test-tunnel" = {
      provider        = "custom-ipsec"
      protocol        = "ipsec"
      local_id        = "test-local-id"
      psk             = "test-preshared-key"
      version         = "2"
      mode            = "active-standby"
      ike_lifetime    = 28800
      ike_mode        = "main"
      ipsec_lifetime  = 3600
      local_subnets   = ["192.168.1.0/24"]
      remote_subnets  = ["10.0.0.0/8"]
      networks        = ["lan"]
      ike_proposals = [
        {
          auth_algo = "sha2"
          dh_group  = "14"
          enc_algo  = "aes256"
        }
      ]
      ipsec_proposals = [
        {
          auth_algo = "sha2"
          dh_group  = "14"
          enc_algo  = "aes256"
        }
      ]
      primary = {
        hosts        = ["203.0.113.100"]
        internal_ips = ["10.1.1.1"]
        probe_ips    = ["8.8.8.8"]
        remote_ids   = ["peer1"]
        wan_names    = ["wan0"]
      }
      secondary = {
        hosts        = ["203.0.113.101"]
        internal_ips = ["10.1.1.2"]
        probe_ips    = ["8.8.4.4"]
        remote_ids   = ["peer2"]
        wan_names    = ["wan0"]
      }
      probe = {
        type      = "icmp"
        interval  = 10
        threshold = 3
        timeout   = 5
      }
      auto_provision = {
        enabled  = false
        provider = "jse-ipsec"
        region   = "us-east-1"
        service_connection = "test-service"
        latlng = {
          lat = 40.7128
          lng = -74.0060
        }
        primary = {
          wan_names = ["wan0"]
          probe_ips = ["8.8.8.8"]
        }
        secondary = {
          wan_names = ["wan1"]
          probe_ips = ["8.8.4.4"]
        }
      }
    }
  }

  tunnel_provider_options = {
    jse = {
      num_users = 100
      org_name  = "Test Organization"
    }
    prisma = {
      service_account_name = "test-service-account"
    }
    zscaler = {
      auth_required                           = false
      aup_enabled                             = true
      aup_block_internet_until_accepted       = true
      aup_force_ssl_inspection                = false
      aup_timeout_in_days                     = 30
      caution_enabled                         = true
      dn_bandwidth                            = 100.5
      idle_time_in_minutes                    = 60
      ofw_enabled                             = true
      surrogate_ip                            = false
      surrogate_ip_enforced_for_known_browsers = false
      surrogate_refresh_time_in_minutes       = 30
      up_bandwidth                            = 50.5
      xff_forward_enabled                     = false
      sub_locations = [
        {
          name                                    = "sub-location-1"
          auth_required                           = false
          aup_enabled                             = true
          aup_block_internet_until_accepted       = true
          aup_force_ssl_inspection                = false
          aup_timeout_in_days                     = 15
          caution_enabled                         = true
          dn_bandwidth                            = 75.0
          idle_time_in_minutes                    = 45
          ofw_enabled                             = true
          surrogate_ip                            = false
          surrogate_ip_enforced_for_known_browsers = false
          surrogate_refresh_time_in_minutes       = 20
          up_bandwidth                            = 25.0
        }
      ]
    }
  }

  vrf_config = {
    enabled = true
  }

  vrf_instances = {
    "guest" = {
      networks = ["guest-network"]
    }
  }
