
  name = "comprehensive_gateway_profile111"
  additional_config_cmds = [
    "set system host-name gateway-profile",
    "set system domain-name example.com"
  ]  
  bgp_config = {
    "peer1" = {
      auth_key = "bgp-auth-key-123"
      bfd_minimum_interval = 1000
      bfd_multiplier = 3
      disable_bfd = false
      export = "export-all"
      export_policy = "export-all"
      extended_v4_nexthop = true
      graceful_restart_time = 120
      hold_time = 90
      import = "import-all"
      import_policy = "import-all"
      local_as = 65001
      neighbor_as = 65002
      type = "external"
      neighbors = {
        "192.168.1.10" = {
          disabled = false
          export_policy = "export-neighbor"
          hold_time = 180
          import_policy = "import-neighbor"
          multihop_ttl = 5
          neighbor_as = 65003
        }
      }
      networks = ["lan", "guest"]
      no_private_as = false
      no_readvertise_to_overlay = false
      tunnel_name = "tunnel1"
      via = "lan"
      vpn_name = "vpn1"
      wan_name = "wan1"
    }
  }
  
  dhcpd_config = {
    enabled = true
    config = {
      "lan" = {
        dns_servers = ["8.8.8.8", "8.8.4.4"]
        dns_suffix = ["example.com"]
        fixed_bindings = {
          "112233445566" = {
            ip = "192.168.1.100"
            name = "server1"
          }
          "112233445567" = {
            ip = "{{server_ip4_var}}"
            name = "server2"
          }
          "112233445568" = {
            ip6 = "{{server_ip6_var}}"
            name = "server3"
          }
        }
        gateway = "192.168.1.1"
        ip_end = "192.168.1.200"
        ip_start = "192.168.1.50"
        lease_time = 86400
        options = {
          "42" = {
            type = "ip"
            value = "192.168.1.1"
          }
        }
        server_id_override = false
        servers = ["192.168.1.5"]
        type = "local"
        vendor_encapsulated = {
          "43:1" = {
            type = "string"
            value = "test-value"
          }
        }
      }
    }
  }
  
  dns_override = false
  dns_servers = ["8.8.8.8", "1.1.1.1"]
  dns_suffix = ["example.com", "test.com"]
  
  extra_routes = {
    "10.0.0.0/8" = {
      via = "192.168.1.254"
    }
  }
  
  extra_routes6 = {
    "2001:db8::/32" = {
      via = "2001:db8::1"
    }
  }
  
  idp_profiles = {
    "security1" = {
      base_profile = "strict"
      name = "Security Profile 1"
      # org_id = "{org_id}"
      overwrites = [
        {
          action = "drop"
          matching = {
            attack_name = ["malware", "virus"]
            dst_subnet = ["192.168.1.0/24"]
            severity = ["high"]
          }
          # name = "Block Malware"
        }
      ]
    }
  }
  
  ip_configs = {
    "management" = {
      ip = "192.168.100.1"
      ip6 = "2001:db8:100::1"
      netmask = "/24"
      netmask6 = "/64"
      secondary_ips = ["192.168.100.2/24"]
      type = "static"
      type6 = "static"
    }
  }
  
  networks = [
    {
      name = "lan"
      gateway = "192.168.1.1"
      gateway6 = "2001:db8::1"
      subnet = "192.168.1.0/24"
      subnet6 = "2001:db8::/64"
      vlan_id = "100"
      disallow_mist_services = false
      internet_access = {
        enabled = true
        destination_nat = {
          "192.168.1.10" = {
            internal_ip = "10.0.0.10"
            name = "web-server-nat"
            port = 80
          }
        }
        static_nat = {
          "192.168.1.20" = {
            internal_ip = "10.0.0.20"
            name = "mail-server-nat"
          }
        }
      }
      internal_access = {
        enabled = true
      }
      isolation = false
      multicast = {
        enabled = true
        groups = {
          "224.0.1.1" = {
            rp_ip = "192.168.1.1"
          }
        }
      }
      tenants = {
        "tenant1" = {
          addresses = ["192.168.1.50/32"]
        }
      }
      vpn_access = {
        "vpn1" = {
          advertised_subnet = "192.168.1.0/24"
          allow_ping = true
          destination_nat = {
            "192.168.1.30" = {
              internal_ip = "10.0.0.30"
              name = "app-server-nat"
              port = 8080
            }
          }
          nat_pool = "192.168.200.0/24"
          no_readvertise_to_lan_bgp = false
          no_readvertise_to_overlay = false
          routed = true
          source_nat = {
            external_ip = "192.168.200.1"
          }
          static_nat = {
            "192.168.1.40" = {
              internal_ip = "10.0.0.40"
              name = "db-server-nat"
            }
          }
          summarized_subnet = "192.168.0.0/16"
          summarized_subnet_to_lan_bgp = "192.168.0.0/16"
        }
      }
    },
    {
      name = "guest"
      gateway = "192.168.2.1"
      subnet = "192.168.2.0/24"
      vlan_id = "200"
      disallow_mist_services = true
      internet_access = {
        enabled = true
      }
      internal_access = {
        enabled = false
      }
      isolation = true
    }
  ]
  
  ntp_override = false
  ntp_servers = ["pool.ntp.org", "time.google.com"]
  
  oob_ip_config = {
    type = "static"
    ip = "192.168.255.1"
    netmask = "/24"
    gateway = "192.168.255.254"
    use_mgmt_vrf = true
    use_mgmt_vrf_for_host_out = true
    vlan_id = "4094"
    node1 = {
      type = "static"
      ip = "192.168.255.2"
      netmask = "/24"
      gateway = "192.168.255.254"
      use_mgmt_vrf = true
      use_mgmt_vrf_for_host_out = true
      vlan_id = "4094"
    }
  }
  
  path_preferences = {
    "preference1" = {
      strategy = "ordered"
      paths = [
        {
          cost = 100
          disabled = false
          gateway_ip = "192.168.1.1"
          name = "primary-path"
          networks = ["lan"]
          target_ips = ["8.8.8.8"]
          type = "local"
        },
        {
          cost = 200
          disabled = false
          name = "vpn1"
          type = "vpn"
          vpn_name = "vpn1"
        }
      ]
    }
  }
  
  port_config = {
    "ge-0/0/0" = {
      usage = "wan"
      name = "uplink"
      critical = true
      description = "Primary WAN uplink"
      disable_autoneg = false
      disabled = false
      duplex = "auto"
      speed = "auto"
      mtu = 1500
      preserve_dscp = true
      wan_type = "broadband"
      wan_ext_ip = "203.0.113.10"
      wan_ext_ip6 = "2001:db8:85a3::8a2e:370:7334"
      wan_source_nat = {
        disabled = false
        nat_pool = "192.168.100.1-192.168.100.50"
        nat6_pool = "2001:db8:100::1-2001:db8:100::50"
      }
      ip_config = {
        # dns = ["8.8.8.8"]
        # dns_suffix = ["example.com"]
        gateway = "192.168.1.1"
        ip = "192.168.1.10"
        netmask = "/24"
        network = "wan"
        type = "static"
      }
    },
    "ge-0/0/1" = {
      usage = "lan"
      description = "LAN port"
      networks = ["lan"]
      disabled = false
      port_network = "lan"
    }
  }
  
  router_id = "192.168.1.1"
  
  routing_policies = {
    "policy1" = {
      terms = [
        {
          actions = {
            accept = true
            add_community = ["65001:100"]
            add_target_vrfs = ["vrf1"]
            community = "65001:200"
            exclude_as_path = ["65002"]
            exclude_community = ["65002:100"]
            local_preference = 100
            med = 50
            prepend_as_path = ["65001", "65001"]
          }
          matching = {
            as_path = ["^65001"]
            community = ["65001:100"]
            interface = ["ge-0/0/0"]
            ip_prefix = ["192.168.1.0/24"]
            ip_prefix_except = ["192.168.1.1/32"]
            neighbor = ["192.168.1.10"]
            prefix_list = ["prefix-list-1"]
            protocol = ["bgp"]
            route_exists = {
              route = "10.0.0.0/8"
              table = "inet.0"
            }
            tag = ["100"]
            vpn_neighbor_mac = ["aa:bb:cc:dd:ee:ff"]
            vpn_path_sla = {
              max_jitter = 10
              max_latency = 100
              max_loss = 1
            }
          }
        }
      ]
    }
  }
  
  service_policies = [
    {
      name = "Policy-14"
      tenants = ["PRD-Core"]
      services = ["any"]
      action = "allow"
      path_preference = "HUB"
      local_routing = false
      idp = {
        enabled = true
        profile = "critical"
        alert_only = false
      }
      appqoe = {
        enabled = true
      }
      antivirus = {
        enabled = true
        profile = "default"
      }
      ewf = [
        {
          enabled = true
          profile = "critical"
          alert_only = false
          block_message = "Access denied"
        }
      ]
      ssl_proxy = {
        enabled = true
        profile = "default"
      }
      skyatp = {
        dns_dga_detection = "strict"
        dns_tunnel_detection = "default"
        http_inspection = "standard"
        iot_device_policy = "enabled"
      }
      syslog = {
        enabled = true
        server_names = ["syslog1.example.com", "syslog2.example.com"]
      }
    }
  ]
  
  tunnel_configs = {
    "tunnel1" = {
      auto_provision = {
        enable = true
        provider = "jse-ipsec"
        latlng = {
          lat = 37.7749
          lng = -122.4194
        }
        primary = {
          probe_ips = ["8.8.8.8", "1.1.1.1"]
          wan_names = ["wan1"]
        }
        secondary = {
          probe_ips = ["8.8.4.4"]
          wan_names = ["wan2"]
        }
      }
      ike_lifetime = 28800
      ike_mode = "main"
      ike_proposals = [
        {
          auth_algo = "sha1"
          dh_group = "14"
          enc_algo = "aes128"
        }
      ]
      ipsec_lifetime = 3600
      ipsec_proposals = [
        {
          auth_algo = "sha1"
          enc_algo = "aes128"
        }
      ]
      local_id = "local@example.com"
      mode = "active-active"
      overlay_subnet = "10.255.0.0/16"
      primary = {
        hosts = ["203.0.113.1"]
        internal_ip = "10.255.1.1"
        node = "node0"
        probe = {
          interval = 5
          threshold = 3
          type = "icmp"
        }
        remote_id = "remote@example.com"
        wan_names = ["wan1"]
      }
      probe = {
        interval = 10
        threshold = 5
        timeout = 2
        type = "icmp"
      }
      protocol = "ipsec"
      provider = "jse-ipsec"
      # psk = "shared-secret-key"
      secondary = {
        hosts = ["203.0.113.2"]
        internal_ip = "10.255.1.2"
        node = "node1"
        probe = {
          interval = 5
          threshold = 3
          type = "icmp"
        }
        remote_id = "remote2@example.com"
        wan_names = ["wan2"]
      }
      version = "2"
    }
  }
  
  tunnel_provider_options = {
    jse = {
      num_users = 100
      org_name = "JSE Organization"
    }
    prisma = {
      service_account_name = "prisma-service-account"
    }
    zscaler = {
      aup_block_internet_until_accepted = true
      aup_timeout_in_days = 4
      aup_force_ssl_inspection = true
      dn_bandwidth = 100
      aup_enabled = false
      caution_enabled = true
      auth_required = true
      sub_locations = [
        {
          aup_block_internet_until_accepted = false
          aup_timeout_in_days = 179
          aup_force_ssl_inspection = false
          dn_bandwidth = 50
          aup_enabled = false
          caution_enabled = false
          auth_required = false
          name = "Sub Location 1"
          up_bandwidth = 25
        }
      ]
      up_bandwidth = 50
      xff_forward_enabled = true
    }
  }
  
  url_filtering_deny_msg = "Access to this website has been blocked by your organization's security policy. Please contact IT support if you believe this is an error."
  
  vrf_config = {
    enabled = true
  }
  
  vrf_instances = {
    "vrf1" = {
      networks = ["lan"]
      extra_routes = {
        "10.2.0.0/16" = {
          via = "192.168.1.254"
        }
      }
    }
  }