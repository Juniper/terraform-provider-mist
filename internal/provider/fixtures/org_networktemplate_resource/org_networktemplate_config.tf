
  name   = "test-networktemplate"

  acl_policies = [
    {
      name     = "test-policy"
      src_tags = ["tag1", "tag2"]
      actions = [
        {
          action  = "allow"
          dst_tag = "dst-tag1"
        }
      ]
    }
  ]

  acl_tags = {
    "tag1" = {
      type         = "network"
      network      = "lan"
      gbp_tag      = 100
      ether_types  = ["ipv4", "ipv6"]
      macs         = ["aa:bb:cc:dd:ee:ff"]
      subnets      = ["192.168.1.0/24"]
      port_usage   = "access"
      radius_group = "guest"
      specs = [
        {
          protocol   = "tcp"
          port_range = "80"
        }
      ]
    }
  }

  additional_config_cmds = [
    "set system host-name test-switch",
    "set protocols ospf area 0.0.0.0 interface all"
  ]

  dhcp_snooping = {
    enabled                 = true
    all_networks           = false
    enable_arp_spoof_check = true
    enable_ip_source_guard = true
    networks               = ["lan", "guest"]
  }

  dns_servers = ["8.8.8.8", "8.8.4.4"]
  dns_suffix  = ["example.com", "test.local"]

  extra_routes = {
    "10.0.0.0/8" = {
      via        = "192.168.1.1"
      metric     = 100
      preference = 10
      discard    = false
      no_resolve = false
      next_qualified = {
        "192.168.1.2" = {
          metric     = 200
          preference = 20
        }
      }
    }
  }

  extra_routes6 = {
    "2001:db8::/32" = {
      via        = "2001:db8::1"
      metric     = 100
      preference = 10
      discard    = false
      no_resolve = false
      next_qualified = {
        "2001:db8::2" = {
          metric     = 200
          preference = 20
        }
      }
    }
  }

  mist_nac = {
    enabled = true
    network = "nac"
  }

  networks = {
    "lan" = {
      vlan_id         = "100"
      subnet          = "192.168.1.0/24"
      gateway         = "192.168.1.1"
      subnet6         = "2001:db8:1::/64"
      gateway6        = "2001:db8:1::1"
      isolation       = false
      isolation_vlan_id = "200"
    }
    "guest" = {
      vlan_id         = "200"
      subnet          = "192.168.2.0/24"
      gateway         = "192.168.2.1"
      isolation       = true
      isolation_vlan_id = "201"
    }
    "management" = {
      vlan_id         = "10"
      subnet          = "192.168.10.0/24"
      gateway         = "192.168.10.1"
      isolation       = false
    }
    "voip" = {
      vlan_id         = "300"
      subnet          = "10.10.10.0/24"
      gateway         = "10.10.10.1"
      subnet6         = "2001:db8:300::/64"
      gateway6        = "2001:db8:300::1"
      isolation       = false
    }
  }

  ntp_servers = ["pool.ntp.org", "time.google.com"]

  ospf_areas = {
    "0.0.0.0" = {
      type             = "default"
      include_loopback = true
      networks = {
        "lan" = {
          auth_type                = "md5"
          auth_password           = "secret12"
          interface_type          = "broadcast"
          hello_interval          = 10
          dead_interval           = 40
          metric                  = 10
          passive                 = false
          bfd_minimum_interval    = 200
          no_readvertise_to_overlay = false
          export_policy           = "policy1"
          import_policy           = "policy2"
          auth_keys = {
            "1" = "key1value"
            "2" = "key2value"
          }
        }
        "guest" = {
          auth_type                = "password"
          auth_password           = "guestpwd"
          interface_type          = "p2p"
          hello_interval          = 5
          dead_interval           = 20
          metric                  = 20
          passive                 = true
          bfd_minimum_interval    = 100
          no_readvertise_to_overlay = true
          export_policy           = "guest_export"
          import_policy           = "guest_import"
        }
      }
    }
    "0.0.0.1" = {
      type             = "stub"
      include_loopback = false
      networks = {
        "management" = {
          auth_type                = "none"
          interface_type          = "broadcast"
          hello_interval          = 30
          dead_interval           = 120
          metric                  = 5
          passive                 = false
          bfd_minimum_interval    = 300
        }
      }
    }
  }

  port_mirroring = {
    "mirror1" = {
      input_networks_ingress = ["lan"]
      input_port_ids_ingress = ["ge-0/0/1", "ge-0/0/2"]
      input_port_ids_egress  = ["ge-0/0/3"]
      output_port_id         = "ge-0/0/10"
    }
    "mirror2" = {
      input_networks_ingress = ["guest"]
      input_port_ids_ingress = ["ge-0/0/4"]
      input_port_ids_egress  = ["ge-0/0/5"]
      output_port_id         = "ge-0/0/11"
    }
  }

  port_usages = {
    "access" = {
      mode                    = "access"
      port_network           = "lan"
      description            = "Access port usage"
      disabled               = false
      enable_qos             = true
      mac_limit              = 10
      persist_mac            = true
      poe_disabled           = false
      speed                  = "auto"
      duplex                 = "auto"
      disable_autoneg        = false
      mtu                    = 1500
      port_auth              = "dot1x"
      enable_mac_auth        = true
      mac_auth_only          = false
      mac_auth_preferred     = true
      mac_auth_protocol      = "eap-md5"
      reauth_interval        = 3600
      allow_dhcpd            = false
      allow_multiple_supplicants = false
      bypass_auth_when_server_down = true
      bypass_auth_when_server_down_for_unkown_client = false
      inter_switch_link      = false
      stp_edge               = true
      stp_no_root_port       = false
      stp_p2p                = true
      use_vstp               = false
      all_networks           = false
      dynamic_vlan_networks  = ["dynamic1", "dynamic2"]
      guest_network          = "guest"
      server_fail_network    = "quarantine"
      server_reject_network  = "reject"
      voip_network           = "voip"
      reset_default_when     = "link_down"
      community_vlan_id      = 300
      ui_evpntopo_id         = "topo1"
      storm_control = {
        no_broadcast           = false
        no_multicast          = false
        no_registered_multicast = false
        no_unknown_unicast    = false
        percentage            = 80
        disable_port          = false
      }
    }
    "trunk" = {
      mode         = "trunk"
      all_networks = true
      description  = "Trunk port usage"
      networks     = ["lan", "guest", "management"]
      native_vlan_id = "1"
      voip_network = "voip"
    }
    "dot1x_trunk" = {
      mode                    = "trunk"
      description            = "802.1X authenticated trunk port"
      port_auth              = "dot1x"
      mac_limit              = 5
      dynamic_vlan_networks  = ["dynamic_vlan1"]
      networks               = ["lan", "guest"]
      enable_qos             = true
      reauth_interval        = 7200
      allow_multiple_supplicants = true
      guest_network          = "guest"
      server_fail_network    = "quarantine"
      server_reject_network  = "reject"
    }
    "dynamic_access" = {
      mode                    = "dynamic"
      port_auth              = "dot1x"
      enable_mac_auth        = true
      all_networks           = false
      dynamic_vlan_networks  = ["dynamic_vlan1"]
      rules = [
        {
          equals = "admin"
          equals_any = ["user1", "user2"]
          expression = "{user.role == admin}"
          usage      = "trunk"
        }
      ]
    }
  }

  radius_config = {
    acct_immediate_update   = true
    acct_interim_interval   = 600
    auth_servers_retries    = 3
    auth_servers_timeout    = 5
    auth_server_selection   = "ordered"
    coa_enabled            = true
    coa_port               = "3799"
    fast_dot1x_timers      = true
    network                = "management"
    source_ip              = "192.168.1.10"
    acct_servers = [
      {
        host             = "radius1.example.com"
        secret           = "secret123"
        port             = "1813"
        keywrap_enabled  = false
        keywrap_format   = "hex"
        keywrap_kek      = "kekvalue"
        keywrap_mack     = "mackvalue"
      }
    ]
    auth_servers = [
      {
        host                         = "radius1.example.com"
        secret                       = "secret123"
        port                         = "1812"
        require_message_authenticator = true
        keywrap_enabled             = false
        keywrap_format              = "hex"
        keywrap_kek                 = "kekvalue"
        keywrap_mack                = "mackvalue"
      }
    ]
  }

  remote_syslog = {
    enabled            = true
    time_format        = "millisecond"
    send_to_all_servers = false
    network            = "management"
    cacerts            = ["cert1", "cert2"]
    archive = {
      files = "10"
      size  = "1024m"
    }
    console = {
      contents = [
        {
          facility = "kernel"
          severity = "info"
        }
      ]
    }
    files = [
      {
        file             = "/var/log/messages"
        enable_tls       = false
        explicit_priority = false
        structured_data  = false
        match            = ".*error.*"
        archive = {
          files = "5"
          size  = "100m"
        }
        contents = [
          {
            facility = "daemon"
            severity = "warning"
          }
        ]
      }
    ]
    servers = [
      {
        host             = "syslog.example.com"
        port             = "514"
        protocol         = "udp"
        facility         = "daemon"
        severity         = "info"
        tag              = "switch"
        explicit_priority = false
        structured_data  = false
        match            = ".*"
        routing_instance = "default"
        server_name      = "syslog1"
        source_address   = "192.168.1.10"
        contents = [
          {
            facility = "kernel"
            severity = "error"
          }
        ]
      }
    ]
    users = [
      {
        user  = "admin"
        match = ".*critical.*"
        contents = [
          {
            facility = "security"
            severity = "alert"
          }
        ]
      }
    ]
  }

  remove_existing_configs = true

  snmp_config = {
    enabled      = true
    name         = "test-switch"
    description  = "Test network template switch"
    location     = "Data Center 1"
    contact      = "admin@example.com"
    engine_id    = "0x80001234"
    engine_id_type = "use_mac_address"
    network      = "management"
    client_list = [
      {
        client_list_name = "readonly_clients"
        clients          = ["192.168.1.0/24", "10.0.0.0/8"]
      }
    ]
    trap_groups = [
      {
        group_name = "critical_traps"
        version    = "v2"
        targets    = ["192.168.1.100", "192.168.1.101"]
        categories = ["chassis", "routing", "sonet"]
      }
    ]
    v2c_config = [
      {
        community_name   = "public"
        authorization    = "read-only"
        client_list_name = "readonly_clients"
        view             = "default"
      }
    ]
    v3_config = {
      notify = [
        {
          name = "trap_notify"
          tag  = "trap_targets"
          type = "trap"
        }
      ]
      notify_filter = [
        {
          profile_name = "filter1"
          contents = [
            {
              oid     = "1.3.6.1.2.1.1"
              include = true
            }
          ]
        }
      ]
      target_address = [
        {
          address              = "192.168.1.100"
          address_mask         = "255.255.255.255"
          target_address_name  = "trap_target1"
          port                 = "162"
          tag_list             = "trap_targets"
          target_parameters    = "trap_params"
        }
      ]
      target_parameters = [
        {
          message_processing_model = "v3"
          name                    = "trap_params"
          notify_filter          = "filter1"
          security_level         = "privacy"
          security_model         = "usm"
          security_name          = "admin"
        }
      ]
      usm = [
        {
          engine_type      = "local_engine"
          remote_engine_id = "0x80001234567890"
          users = [
            {
              name                    = "admin"
              authentication_type     = "authentication-md5"
              authentication_password = "auth_pass123"
              encryption_type         = "privacy-des"
              encryption_password     = "priv_pass123"
            }
          ]
        }
      ]
      vacm = {
        access = [
          {
            group_name = "admin_group"
            prefix_list = [
              {
                context_prefix  = "default"
                notify_view     = "all"
                read_view       = "all"
                security_level  = "privacy"
                security_model  = "usm"
                type           = "context_prefix"
                write_view     = "all"
              }
            ]
          }
        ]
        security_to_group = {
          security_model = "usm"
          content = [
            {
              group         = "admin_group"
              security_name = "admin"
            }
          ]
        }
      }
    }
    views = [
      {
        view_name = "all"
        oid       = "1"
        include   = true
      }
    ]
  }

  switch_matching = {
    enable = true
    rules = [
      {
        name         = "access_switch_rule"
        match_name   = "access-sw"
        match_model  = "EX3400"
        match_role   = "access"
        match_type   = "name"
        match_value  = "access"
        match_name_offset = 10
        additional_config_cmds = ["set vlans guest vlan-id 200"]
        ip_config = {
          network = "management"
          type    = "dhcp"
        }
        oob_ip_config = {
          type                     = "dhcp"
          use_mgmt_vrf            = true
          use_mgmt_vrf_for_host_out = false
        }
        port_config = {
          "ge-0/0/0" = {
            usage = "access"
            speed = "1g"
          }
          "ge-0/0/1" = {
            usage = "trunk"
            speed = "10g"
          }
        }
        port_mirroring = {
          "local_mirror" = {
            input_port_ids_ingress = ["ge-0/0/2"]
            output_port_id         = "ge-0/0/10"
          }
        }
      }
    ]
  }

  switch_mgmt = {
    ap_affinity_threshold   = 12
    cli_banner             = "Welcome to the test switch"
    cli_idle_timeout       = 60
    config_revert_timer    = 10
    dhcp_option_fqdn       = true
    disable_oob_down_alarm = false
    fips_enabled           = false
    mxedge_proxy_host      = "proxy.example.com"
    mxedge_proxy_port      = "8080"
    remove_existing_configs = true
    root_password          = "secret123"
    use_mxedge_proxy       = true
    local_accounts = {
      "testuser" = {
        password = "testpass123"
        role     = "admin"
      }
      "readonly" = {
        password = "readonly123"
        role     = "read"
      }
    }
    protect_re = {
      enabled         = true
      hit_count       = true
      allowed_services = ["ssh", "icmp"]
      trusted_hosts   = ["192.168.1.0/24", "10.0.0.0/8"]
      custom = [
        {
          port_range = "8080-8090"
          protocol   = "tcp"
          subnets    = ["192.168.1.0/24"]
        }
      ]
    }
    tacacs = {
      enabled      = true
      default_role = "read"
      network      = "management"
      acct_servers = [
        {
          host    = "tacacs1.example.com"
          port    = "49"
          secret  = "secret123"
          timeout = 10
        }
      ]
      tacplus_servers = [
        {
          host    = "tacacs1.example.com"
          port    = "49"
          secret  = "secret123"
          timeout = 5
        }
      ]
    }
  }

  vrf_config = {
    enabled = true
  }

  vrf_instances = {
    "management" = {
      networks = ["mgmt"]
      extra_routes = {
        "0.0.0.0/0" = {
          via = "192.168.100.1"
        }
      }
    }
    "guest" = {
      networks = ["guest", "public"]
      extra_routes = {
        "10.0.0.0/8" = {
          via = "192.168.2.254"
        }
        "172.16.0.0/12" = {
          via = "192.168.2.1"
        }
      }
    }
    "voice" = {
      networks = ["voip"]
      extra_routes = {
        "0.0.0.0/0" = {
          via = "10.10.10.254"
        }
      }
    }
  }

