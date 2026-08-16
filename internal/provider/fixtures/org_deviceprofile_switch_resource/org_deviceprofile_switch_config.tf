
  name = "comprehensive_switch_profile_0"
  additional_config_cmds = [
    "set system host-name switch-profile",
    "set system domain-name example.com"
  ]
  acl_policies = [
    {
      name     = "allow_internal"
      disabled = false
      src_tags = ["internal_hosts", "servers"]
      actions = [
        {
          action  = "allow"
          dst_tag = "internal_hosts"
        },
        {
          action  = "deny"
          dst_tag = "guest_network"
        }
      ]
    }
  ]
  acl_tags = {
    "internal_hosts" = {
      type        = "subnet"
      subnets     = ["10.0.0.0/8", "172.16.0.0/12"]
      ether_types = ["0x0800", "0x86DD"]
    }
    "servers" = {
      type = "mac"
      macs = ["aa:bb:cc:dd:ee:01", "aa:bb:cc:dd:ee:02"]
    }
    "radius_tag" = {
      type         = "radius_group"
      radius_group = "employees"
    }
    "gbp_tag" = {
      type    = "dynamic_gbp"
      gbp_tag = 100
    }
    "port_usage_tag" = {
      type       = "port_usage"
      port_usage = "employee_port"
    }
    "network_tag" = {
      type    = "network"
      network = "mgmt"
      specs = [
        {
          port_range = "80"
          protocol   = "tcp"
        }
      ]
    }
  }
  dhcp_snooping = {
    all_networks          = false
    enable_arp_spoof_check = true
    enable_ip_source_guard = true
    enabled               = true
    networks              = ["mgmt", "user"]
  }
  dhcpd_config = {
    enabled = true
    config = {
      "mgmt" = {
        type              = "server"
        ip_start          = "192.168.1.10"
        ip_end            = "192.168.1.200"
        ip_start6         = "2001:db8::10"
        ip_end6           = "2001:db8::200"
        gateway           = "192.168.1.1"
        dns_servers       = ["8.8.8.8", "8.8.4.4"]
        dns_suffix        = ["example.com"]
        lease_time        = 86400
        server_id_override = false
        servers           = ["192.168.1.5"]
        servers6          = ["2001:db8::1"]
        type6             = "server"
        fixed_bindings = {
          "aabbccddeeff" = {
            ip   = "192.168.1.50"
            name = "fixed-server"
          }
          "112233445566" = {
            ip6  = "2001:db8::50"
            name = "fixed-server6"
          }
        }
        options = {
          "42" = {
            type  = "ip"
            value = "192.168.1.1"
          }
        }
        vendor_encapsulated = {
          "43:1" = {
            type  = "string"
            value = "test-value"
          }
        }
      }
    }
  }
  dns_servers = ["8.8.8.8", "8.8.4.4"]
  dns_suffix  = ["example.com", "corp.example.com"]
  evpn_config = {
    enabled = true
    role    = "collapsed-core"
  }
  extra_routes = {
    "0.0.0.0/0" = {
      via        = "192.168.1.1"
      discard    = false
      no_resolve = false
    }
    "10.100.0.0/16" = {
      via        = "10.0.0.1"
      metric     = 100
      preference = 10
      next_qualified = {
        "10.0.0.2" = {
          metric     = 200
          preference = 20
        }
      }
    }
  }
  extra_routes6 = {
    "::/0" = {
      via        = "fe80::1"
      no_resolve = false
      discard    = false
      metric     = 10
      preference = 5
      next_qualified = {
        "fe80::2" = {
          metric     = 100
          preference = 10
        }
      }
    }
  }
  iot_config = {
    "DI-1" = {
      enabled    = true
      name       = "door-sensor-1"
      input_src  = "IN0"
      alarm_class = "major"
    }
  }
  ip_config = {
    type       = "static"
    ip         = "192.168.1.10"
    netmask    = "255.255.255.0"
    gateway    = "192.168.1.1"
    network    = "mgmt"
    dns        = ["8.8.8.8"]
    dns_suffix = ["example.com"]
  }
  mist_nac = {
    enabled = true
    network = "mgmt"
  }
  networks = {
    "mgmt" = {
      vlan_id           = "10"
      subnet            = "192.168.10.0/24"
      subnet6           = "2001:db8:10::/64"
      gateway           = "192.168.10.1"
      gateway6          = "2001:db8:10::1"
      isolation_vlan_id = "11"
    }
    "user" = {
      vlan_id   = "20"
      subnet    = "192.168.20.0/24"
      gateway   = "192.168.20.1"
      isolation = true
      multicast = {
        enabled      = true
        igmp_version = "3"
      }
    }
    "guest" = {
      vlan_id = "30"
      subnet  = "192.168.30.0/24"
      gateway = "192.168.30.1"
    }
  }
  ntp_servers = ["time1.example.com", "time2.example.com"]
  oob_ip_config = {
    type                      = "static"
    ip                        = "10.255.255.1"
    netmask                   = "/24"
    gateway                   = "10.255.255.254"
    network                   = "oob"
    use_mgmt_vrf              = false
    use_mgmt_vrf_for_host_out = false
  }
  ospf_areas = {
    "0" = {
      type             = "default"
      include_loopback = true
      networks = {
        "mgmt" = {
          interface_type            = "broadcast"
          hello_interval            = 10
          dead_interval             = 40
          auth_type                 = "md5"
          auth_keys = {
            "1" = "ospf-md5-key"
          }
          bfd_minimum_interval      = 500
          export_policy             = "export-ospf"
          import_policy             = "import-ospf"
          metric                    = 100
          no_readvertise_to_overlay = false
          passive                   = false
        }
        "user" = {
          interface_type = "p2p"
          auth_type      = "password"
          auth_password  = "auth1234"
          passive        = true
        }
      }
    }
  }
  other_ip_configs = {
    "vlan10" = {
      type         = "static"
      ip           = "10.0.10.1"
      netmask      = "255.255.255.0"
      evpn_anycast = false
      ip6          = "2001:db8:10::1"
      netmask6     = "/64"
      type6        = "static"
    }
  }
  port_config = {
    "ge-0/0/0-1" = {
      usage            = "uplink"
      aggregated       = true
      ae_idx           = 0
      ae_disable_lacp  = false
      ae_lacp_force_up = true
      ae_lacp_passive  = true
      ae_lacp_slow     = false
      critical         = false
      description      = "Uplink to core"
      disable_autoneg  = false
      duplex           = "auto"
      esilag           = false
      mtu              = 9000
      no_local_overwrite = false
      poe_disabled     = false
      speed            = "auto"
    }
    "ge-0/0/2-47" = {
      usage         = "access_port"
      networks      = ["user"]
      description   = "User access ports"
      dynamic_usage = "default"
      port_network  = "user"
    }
  }
  port_mirroring = {
    "mirror1" = {
      input_networks_ingress = ["mgmt"]
      input_port_ids_egress  = ["ge-0/0/11"]
      input_port_ids_ingress = ["ge-0/0/10"]
      output_ip_address      = "192.168.10.100"
      output_network         = "mgmt"
    }
    "mirror2" = {
      input_port_ids_ingress = ["ge-0/0/20"]
      output_port_id         = "ge-0/0/47"
    }
  }
  port_usages = {
    "access_port" = {
      mode              = "access"
      port_network      = "user"
      stp_edge          = true
      stp_disable       = false
      allow_dhcpd       = true
      poe_disabled      = false
      server_fail_retry_interval = 300
    }
    "uplink" = {
      mode          = "trunk"
      all_networks  = true
    }
    "trunk_port" = {
      mode     = "trunk"
      networks = ["user", "guest"]
    }
    "employee_port" = {
      mode         = "access"
      port_network = "mgmt"
      port_auth    = "dot1x"
      voip_network = "voip"
    }
    "comprehensive_port" = {
      mode                                        = "access"
      allow_multiple_supplicants                  = false
      bypass_auth_when_server_down                = false
      bypass_auth_when_server_down_for_unknown_client = false
      bypass_auth_when_server_down_for_voip       = false
      community_vlan_id                           = 100
      description                                 = "Comprehensive test port"
      disable_autoneg                             = false
      disabled                                    = false
      duplex                                      = "auto"
      dynamic_vlan_networks                       = ["user", "guest"]
      enable_mac_auth                             = false
      enable_qos                                  = false
      guest_network                               = "guest"
      inter_isolation_network_link                = false
      inter_switch_link                           = false
      mac_auth_only                               = false
      mac_auth_preferred                          = false
      mac_auth_protocol                           = "eap-md5"
      mac_limit                                   = "10"
      mtu                                         = "1500"
      persist_mac                                 = false
      poe_disabled                                = false
      poe_keep_state_when_reboot                  = false
      poe_priority                                = "high"
      port_auth                                   = "dot1x"
      port_network                                = "user"
      reauth_interval                             = "3600"
      reset_default_when                          = "link_down"
      server_fail_network                         = "guest"
      server_fail_retry_interval                  = 120
      server_reject_network                       = "guest"
      speed                                       = "auto"
      stp_edge                                    = true
      stp_no_root_port                            = false
      stp_p2p                                     = false
      stp_required                                = false
      use_vstp                                    = false
      voip_network                                = "voip"
      storm_control = {
        disable_port            = false
        no_broadcast            = false
        no_multicast            = false
        no_registered_multicast = false
        no_unknown_unicast      = false
        percentage              = 80
      }
    }
    "dynamic_port" = {
      mode = "dynamic"
      rules = [
        {
          src         = "radius_username"
          description = "Admin rule"
          equals      = "admin"
          usage       = "employee_port"
        },
        {
          src        = "radius_usermac"
          equals_any = ["aa:bb:cc:dd:ee:01", "aa:bb:cc:dd:ee:02"]
          usage      = "access_port"
        },
        {
          src        = "lldp_system_name"
          expression = ".*switch.*"
          usage      = "access_port"
        }
      ]
    }
  }
  radius_config = {
    acct_immediate_update = true
    acct_interim_interval = 60
    auth_server_selection = "ordered"
    auth_servers_retries  = 3
    auth_servers_timeout  = 5
    coa_enabled           = true
    coa_port              = "3799"
    fast_dot1x_timers     = false
    network               = "mgmt"
    source_ip             = "192.168.1.10"
    auth_servers = [
      {
        host                        = "192.168.1.100"
        port                        = "1812"
        secret                      = "radius-secret-1"
        require_message_authenticator = false
        keywrap_enabled             = false
        keywrap_format              = "ascii"
        keywrap_kek                 = "1234567890123456"
        keywrap_mack                = "12345678"
      }
    ]
    acct_servers = [
      {
        host           = "192.168.1.100"
        port           = "1813"
        secret         = "radius-secret-1"
        keywrap_enabled = false
        keywrap_format  = "ascii"
        keywrap_kek     = "1234567890123456"
        keywrap_mack    = "12345678"
      }
    ]
  }
  remote_syslog = {
    enabled             = true
    send_to_all_servers = false
    network             = "mgmt"
    time_format         = "millisecond"
    archive = {
      files = "5"
      size  = "10m"
    }
    cacerts = ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0t"]
    console = {
      contents = [
        {
          facility = "any"
          severity = "any"
        }
      ]
    }
    files = [
      {
        file              = "auth.log"
        match             = "auth"
        enable_tls        = false
        explicit_priority = false
        structured_data   = false
        archive = {
          files = "3"
          size  = "5m"
        }
        contents = [
          {
            facility = "authorization"
            severity = "any"
          }
        ]
      }
    ]
    servers = [
      {
        host             = "syslog.example.com"
        port             = "514"
        protocol         = "udp"
        severity         = "any"
        facility         = "any"
        match            = ".*error.*"
        routing_instance = "default"
        server_name      = "primary-syslog"
        source_address   = "192.168.1.10"
        structured_data  = false
        explicit_priority = false
        tag              = "switch-log"
        contents = [
          {
            facility = "any"
            severity = "any"
          }
        ]
      }
    ]
    users = [
      {
        user  = "admin"
        match = "admin.*"
        contents = [
          {
            facility = "authorization"
            severity = "any"
          }
        ]
      }
    ]
  }
  routing_policies = {
    "export-policy" = {
      terms = [
        {
          name = "accept-lan"
          matching = {
            protocol = ["direct", "static"]
          }
          routing_policy_term_actions = {
            accept = true
          }
        }
      ]
    }
  }
  snmp_config = {
    enabled     = true
    name        = "switch-profile"
    location    = "datacenter-1"
    contact     = "admin@example.com"
    description = "Switch Profile SNMP"
    network     = "mgmt"
    v2c_config = [
      {
        community_name  = "public"
        authorization   = "read-only"
      }
    ]
    trap_groups = [
      {
        group_name = "trap-group-1"
        version    = "v2"
        targets    = ["192.168.1.200"]        
        categories = ["chassis", "routing"]      
      }
    ]
  }
  stp_config = {
    bridge_priority = "32768"
  }
  switch_mgmt = {
    ap_affinity_threshold = 12
    cli_banner            = "Authorized access only"
    cli_idle_timeout      = 30
    config_revert_timer   = 10
    dhcp_option_fqdn      = true
    disable_oob_down_alarm = false
    fips_enabled          = false
    protect_re = {
      enabled          = true
      allowed_services = ["icmp", "ssh"]
      trusted_hosts    = ["10.0.0.0/8"]
    }
    tacacs = {
      enabled      = true
      network      = "mgmt"
      default_role = "read"
      tacplus_servers = [
        {
          host    = "tacacs.example.com"
          port    = "49"
          secret  = "tacacs-secret"
          timeout = 10
        }
      ]
    }
  }
  use_router_id_as_source_ip = false
  vrf_config = {
    enabled = true
  }
  vrf_instances = {
    "vrf-red" = {
      networks = ["user"]
      multicast_config = {
        anycast_rp = false
        rp_ip      = "192.168.20.1"
      }
      extra_routes = {
        "10.200.0.0/16" = {
          via = "10.0.0.1"
        }
      }
    }
  }
  vrrp_config = {
    enabled = true
    groups = {
      "1" = {
        priority = 200
        preempt  = true
      }
    }
  }
