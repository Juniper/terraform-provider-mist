
  name = "comprehensive_switch_profile_0"
  additional_config_cmds = [
    "set system host-name switch-profile",
    "set system domain-name example.com"
  ]
  acl_policies = [
    {
      name     = "allow_internal"
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
      type    = "subnet"
      subnets = ["10.0.0.0/8", "172.16.0.0/12"]
    }
    "servers" = {
      type = "mac"
      macs = ["aa:bb:cc:dd:ee:01", "aa:bb:cc:dd:ee:02"]
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
        type       = "server"
        ip_start   = "192.168.1.10"
        ip_end     = "192.168.1.200"
        gateway    = "192.168.1.1"
        dns_servers = ["8.8.8.8", "8.8.4.4"]
        dns_suffix  = ["example.com"]
        lease_time  = 86400
        fixed_bindings = {
          "aabbccddeeff" = {
            ip   = "192.168.1.50"
            name = "fixed-server"
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
      via = "192.168.1.1"
    }
    "10.100.0.0/16" = {
      via        = "10.0.0.1"
      metric     = 100
      preference = 10
    }
  }
  extra_routes6 = {
    "::/0" = {
      via        = "fe80::1"
      no_resolve = false
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
    type    = "static"
    ip      = "192.168.1.10"
    netmask = "255.255.255.0"
    gateway = "192.168.1.1"
    network = "mgmt"
    dns     = ["8.8.8.8"]
  }
  mist_nac = {
    enabled = true
    network = "mgmt"
  }
  networks = {
    "mgmt" = {
      vlan_id = "10"
      subnet  = "192.168.10.0/24"
      gateway = "192.168.10.1"
    }
    "user" = {
      vlan_id   = "20"
      subnet    = "192.168.20.0/24"
      gateway   = "192.168.20.1"
      isolation = true
    }
    "guest" = {
      vlan_id = "30"
      subnet  = "192.168.30.0/24"
      gateway = "192.168.30.1"
    }
  }
  ntp_servers = ["time1.example.com", "time2.example.com"]
  oob_ip_config = {
    type    = "dhcp"
    network = "oob"
  }
  ospf_areas = {
    "0" = {
      type             = "default"
      include_loopback = true
      networks = {
        "mgmt" = {
          interface_type  = "broadcast"
          hello_interval  = 10
          dead_interval   = 40
          auth_type       = "none"
          passive         = false
        }
      }
    }
  }
  other_ip_configs = {
    "vlan10" = {
      type    = "static"
      ip      = "10.0.10.1"
      netmask = "255.255.255.0"
    }
  }
  port_config = {
    "ge-0/0/0-1" = {
      usage       = "uplink"
      aggregated  = true
      ae_idx      = 0
      ae_lacp_force_up = true
      ae_lacp_passive  = true
      description = "Uplink to core"
      speed       = "auto"
    }
    "ge-0/0/2-47" = {
      usage    = "access_port"
      networks = ["user"]
      description = "User access ports"
    }
  }
  port_mirroring = {
    "mirror1" = {
      input_port_ids_ingress = ["ge-0/0/10"]
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
    "employee_port" = {
      mode         = "access"
      port_network = "mgmt"
      port_auth    = "dot1x"
      voip_network = "voip"
    }
  }
  radius_config = {
    auth_server_selection = "ordered"
    auth_servers_retries  = 3
    auth_servers_timeout  = 5
    coa_enabled           = true
    coa_port              = "3799"
    network               = "mgmt"
    auth_servers = [
      {
        host   = "192.168.1.100"
        port   = "1812"
        secret = "radius-secret-1"
      }
    ]
    acct_servers = [
      {
        host   = "192.168.1.100"
        port   = "1813"
        secret = "radius-secret-1"
      }
    ]
  }
  remote_syslog = {
    enabled           = true
    send_to_all_servers = false
    network           = "mgmt"
    servers = [
      {
        host     = "syslog.example.com"
        port     = "514"
        protocol = "udp"
        severity = "any"
        facility = "any"
        contents = [
          {
            facility = "any"
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
