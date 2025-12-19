
  device_id = "00000000-0000-0000-1000-5c5b35000032"
  name      = "test-switch-comprehensive"
  additional_config_cmds = ["set system host-name switch1", "set system domain-name example.com"]
  bgp_config = {
    "bgp_config_1" = {
      type                 = "internal"
      networks             = ["lan", "wan"]
      bfd_minimum_interval = 150
      local_as             = "65100"
      hold_time            = 60
      auth_key             = "bgpkey1"
      export_policy        = "export_policy_1"
      import_policy        = "import_policy_1"
      neighbors = {
        "10.1.0.1" = {
          neighbor_as   = "65100"
          hold_time     = 90
          import_policy = "import_policy_1"
          export_policy = "export_policy_1"
        }
      }
    }
    "bgp_config_2" = {
      type                 = "external"
      networks             = ["wan"]
      bfd_minimum_interval = 250
      local_as             = "65200"
      hold_time            = 150
      auth_key             = "bgpkey2"
      export_policy        = "export_policy_2"
      import_policy        = "import_policy_2"
      neighbors = {
        "192.168.100.1" = {
          neighbor_as   = "65300"
          hold_time     = 210
          import_policy = "import_policy_2"
          export_policy = "export_policy_2"
          multihop_ttl  = 10
        }
      }
    }
  }
  acl_policies = [
    {
      name = "policy1"
      src_tags = ["src1", "src2"]
      actions = [
        {
          action = "allow"
          dst_tag = "dst1"
        },
        {
          action = "deny"
          dst_tag = "dst2"
        }
      ]
    },
    {
      name = "policy2"
      src_tags = ["src3"]
      actions = [
        {
          action = "allow"
          dst_tag = "dst3"
        }
      ]
    }
  ]
  acl_tags = {
    "tag1" = {
      ether_types = ["ipv4", "ipv6"]
      gbp_tag     = 1001
      macs        = ["00:11:22:33:44:55", "66:77:88:99:AA:BB"]
      network     = "net1"
      port_usage  = "access"
      radius_group = "group1"
      specs = [
        {
          port_range = "1-10"
          protocol   = "tcp"
        },
        {
          port_range = "11-20"
          protocol   = "udp"
        }
      ]
      subnets = ["192.168.1.0/24", "10.0.0.0/8"]
      type    = "mac"
    },
    "tag2" = {
      ether_types = ["ipv4"]
      gbp_tag     = 2002
      macs        = ["AA:BB:CC:DD:EE:FF"]
      network     = "net2"
      port_usage  = "trunk"
      radius_group = "group2"
      specs = [
        {
          port_range = "21-30"
          protocol   = "icmp"
        }
      ]
      subnets = ["172.16.0.0/16"]
      type    = "network"
    }
  }
  dhcp_snooping = {
    all_networks           = false
    enable_arp_spoof_check = true
    enable_ip_source_guard = true
    enabled                = true
    networks               = ["net1", "net2", "net3"]
  }
  dhcpd_config = {
    enabled = true
    config = {
      "lan" = {
        dns_servers  = ["8.8.8.8", "8.8.4.4"]
        dns_suffix   = ["example.com"]
        gateway      = "192.168.1.1"
        ip_start     = "192.168.1.100"
        ip_end       = "192.168.1.200"
        ip_start6    = "2001:db8::100"
        ip_end6      = "2001:db8::200"
        lease_time   = 86400
        type         = "local"
        type6        = "local"
        server_id_override = false
        servers      = ["192.168.1.1"]
        servers6     = ["2001:db8::1"]
        fixed_bindings = {
          "client1" = {
            ip   = "192.168.1.50"
            ip6  = "2001:db8::50"
            name = "test-client"
          }
          "client2" = {
            ip   = "192.168.1.51"
            name = "another-client"
          }
        }
        options = {
          "option1" = {
            type  = "ip"
            value = "192.168.1.1"
          }
        }
        vendor_encapsulated = {
          "vendor1" = {
            type  = "string"
            value = "test-value"
          }
        }
      }
    }
  }
  port_config = {
    "ge-0/0/0" = {
      usage         = "inet"
      description   = "Internet port"
      networks      = ["wan", "internet"]
      mtu           = 9000
      speed         = "1g"
      duplex        = "full"
    }
    "ge-0/0/1" = {
      usage         = "inet"
      description   = "LAN port"
      networks      = ["lan"]
      poe_disabled  = true
      speed         = "10g"
    }
  }
  port_usages = {
    "access_port" = {
      all_networks     = false
      allow_dhcpd      = true
      bypass_auth_when_server_down_for_voip = true
      description      = "Access port usage"
      disabled         = false
      enable_qos       = true
      mac_limit        = "5"
      mode             = "access"
      networks         = ["lan"]
      poe_priority     = "high"
      port_network     = "lan"
      speed            = "auto"
      stp_disable      = false
      stp_required     = true
      rules = [
        {
          src   = "mac"
          equals = "aa:bb:cc:dd:ee:ff"
          usage = "access"
        }
      ]
      storm_control = {
        no_broadcast          = false
        no_multicast          = false
        no_registered_multicast = false
        no_unknown_unicast    = false
        percentage            = 80
      }
    }
    "trunk_port" = {
      all_networks     = true
      allow_dhcpd      = false
      bypass_auth_when_server_down_for_voip = false
      description      = "Trunk port usage"
      disabled         = false
      enable_qos       = false
      mac_limit        = "10"
      mode             = "trunk"
      networks         = ["lan", "wan"]
      poe_priority     = "medium"
      port_network     = "lan"
      speed            = "1g"
      stp_disable      = true
      stp_required     = false
      rules = [
        {
          src   = "vlan"
          equals_any = ["100", "200"]
          usage = "trunk"
        }
      ]
    }
  }

  routing_policies = {
    test_import = {
      terms = [
        {
          matching = {
            prefix = [
              "10.1.0.0/24"
            ],
            as_path = [
              "234"
            ],
            protocol = [
              "direct"
            ],
            community = [
              "my_com"
            ]
          },
          actions = {
            accept = true,
            community = [
              "test"
            ],
            prepend_as_path = [
              "1234"
            ],
            local_preference = "2345"
          },
          name = "test_direct"
        },
        {
          matching = {
            prefix = [
              "10.0.0.0/8"
            ],
            as_path = [
              "65000"
            ],
            protocol = [
              "bgp"
            ],
            community = [
              "test"
            ]
          },
          actions = {
            accept = true,
            community = [
              "my_community"
            ],
            prepend_as_path = [
              "432"
            ],
            local_preference = "532"
          },
          name = "test_bgp"
        },
        {
          matching = {
            prefix = [
              "10.2.0.0/24"
            ],
            as_path = [
              "45332"
            ]
          },
          actions = {
            accept = true
          },
          name = "test_none"
        },
        {
          matching = {
            prefix = [
              "10.3.0.0/24"
            ],
            as_path = [
              "2314"
            ],
            protocol = [
              "evpn"
            ]
          },
          actions = {
              accept = true
          },
          name = "test_evpn"
        },
        {
          matching = {
            prefix = [
              "10.5.0.0/25"
            ],
            protocol = [
              "ospf"
            ]
          },
          actions = {
            accept = true
          },
          name = "test_ospf"
        },
        {
          matching = {
            protocol = [
              "static"
            ]
          },
          actions = {
            accept = true
          },
          name = "test_static"
        }
      ]
    }
  }
  image_url = "https://example.com/switch.png"
  notes     = "Comprehensive switch config"
  model     = "QFX5100"
  serial    = "SW-1234567890"
  x         = 100.5
  y         = 200.75

