additional_config_cmds = [
  "set system hostname test-switch",
  "set system domain-name test.local"
]
auto_upgrade_linecard = true
# default_port_usage = "lan"
dhcp_snooping = {
  enabled = true
  all_networks = false
  enable_arp_spoof_check = true
  enable_ip_source_guard = true
  networks = ["lan", "guest"]
}
disabled_system_defined_port_usages = ["uplink"]
dns_servers = ["8.8.8.8", "8.8.4.4", "1.1.1.1"]
dns_suffix = ["test.local", "example.com"]
extra_routes = {
  "0.0.0.0/0" = {
    via = "192.168.1.1"
    metric = 100
    preference = 1
    no_resolve = false
    discard = false
  }
  "10.0.0.0/8" = {
    via = "192.168.1.1"
    metric = 50
    preference = 2
    no_resolve = true
    discard = false
    next_qualified = {
      "192.168.1.2" = {
        metric = 60
        preference = 3
      }
    }
  }
}
extra_routes6 = {
  "::/0" = {
    via = "2001:db8::1"
    metric = 100
    preference = 1
    no_resolve = false
    discard = false
  }
}
mist_nac = {
  enabled = true
  network = "lan"
}
networks = {
  "lan" = {
    vlan_id = "100"
    subnet = "192.168.1.0/24"
    gateway = "192.168.1.1"
    dns_servers = ["192.168.1.1"]
    dns_suffix = ["lan.test.local"]
    disallow_mist_services = false
    isolation = false
  }
  "guest" = {
    vlan_id = "200"
    subnet = "192.168.200.0/24"
    gateway = "192.168.200.1"
    dns_servers = ["8.8.8.8"]
    disallow_mist_services = true
    isolation = true
  }
}
ntp_servers = ["pool.ntp.org", "time.nist.gov"]
ospf_areas = {
  "0.0.0.0" = {
    include_loopback = true
    type = "default"
    networks = {
      "192.168.1.0/24" = {
        auth_type = "md5"
        auth_keys = {
          "1" = "ospf_key_1"
          "2" = "ospf_key_2"
        }
        bfd_minimum_interval = 100
        dead_interval = 40
        hello_interval = 10
        interface_type = "nbma"
        metric = 1
        no_readvertise_to_overlay = false
        passive = false
      }
    }
  }
}
port_mirroring = {
  "mirror1" = {
    input_networks_ingress = ["lan"]
    input_port_ids_egress = ["ge-0/0/1"]
    input_port_ids_ingress = ["ge-0/0/2"]
    output_network = "lan"
  }
}
port_usages = {
  "uplink_usage" = {
    mode = "trunk"
    port_network = "lan"
    voip_network = "voice"
    stp_disable = false
    stp_edge = true
    stp_no_root_port = false
    stp_p2p = true
    stp_required = true
    speed = "auto"
    duplex = "auto"
    mac_limit = 10
    persist_mac = false
    poe_disabled = false
    enable_qos = true
    storm_control = {
      no_broadcast = false
      no_multicast = false
      no_registered_multicast = false
      no_unknown_unicast = false
      percentage = 80
    }
  }
}
radius_config = {
  acct_interim_interval = 60
  acct_servers = [
    {
      host = "192.168.1.10"
      port = 1813
      secret = "radius_secret"
    }
  ]
  auth_servers = [
    {
      host = "192.168.1.10"
      port = 1812
      secret = "radius_secret"
    }
  ]
  auth_servers_retries = 3
  auth_servers_timeout = 5
  coa_enabled = true
  coa_port = 3799
  network = "lan"
  source_ip = "192.168.1.100"
}
remote_syslog = {
  archive = {
    files = 10
    size = "1m"
  }
  console = {
    contents = [
      {
        facility = "any"
        severity = "info"
      }
    ]
  }
  enabled = true
  files = [
    {
      archive = {
        files = 5
        size = "1m"
      }
      contents = [
        {
          facility = "any"
          severity = "info"
        }
      ]
      file = "/var/log/messages"
      structured_data = true
    }
  ]
  network = "lan"
  send_to_all_servers = true
  servers = [
    {
      facility = "any"
      host = "192.168.1.20"
      port = 514
      protocol = "udp"
      severity = "info"
      structured_data = true
      tag = "switch"
    }
  ]
  time_format = "millisecond"
  users = [
    {
      contents = [
        {
          facility = "any"
          severity = "emergency"
        }
      ]
      user = "*"
    }
  ]
}
snmp_config = {
  client_list = [
    {
      client_list_name = "mgmt_hosts"
      clients = ["192.168.1.50/32", "192.168.1.51/32"]
    }
  ]
  community = "public"
  contact = "admin@test.local"
  description = "Test Switch"
  enabled = true
  engine_id = "test_engine"
  location = "Data Center"
  name = "test-switch"
  network = "lan"
  trap_groups = [
    {
      categories = []
      group_name = "mgmt_traps"
      targets = ["192.168.1.50"]
      version = "v2"
    }
  ]
  v2c_config = [
    {
      client_list_name = "mgmt_hosts"
      community_name = "public"
      view = "default"
    }
  ]
  v3_config = {
    enabled = true
    notify = [
      {
        name = "mgmt_notify"
        tag = "mgmt_tag"
        type = "trap"
      }
    ]
    notify_filter = [
      {
        contents = [
          {
            include = true
            oid = "1.3.6.1.2.1.1"
          }
        ]
        profile_name = "mgmt_filter"
      }
    ]
    target_address = [
      {
        address = "192.168.1.50"
        address_mask = "255.255.255.255"
        port = 162
        tag_list = "mgmt_tag"
        target_address_name = "mgmt_target"
        target_parameters = "mgmt_params"
      }
    ]
    target_parameters = [
      {
        message_processing_model = "v3"
        name = "mgmt_params"
        notify_filter = "mgmt_filter"
        security_level = "privacy"
        security_model = "usm"
        security_name = "mgmt_user"
      }
    ]
    vacm = {
      access = [
        {
          context_match = "exact"
          context_prefix = ""
          group_name = "mgmt_group"
          notify_view = "mgmt_view"
          read_view = "mgmt_view"
          security_level = "privacy"
          security_model = "usm"
          write_view = "mgmt_view"
        }
      ]
      security_to_group = {
        group_name = "mgmt_group"
        security_model = "usm"
        security_name = "mgmt_user"
      }
      view = [
        {
          oid = "1.3.6.1.2.1.1"
          view_name = "mgmt_view"
          view_type = "include"
        }
      ]
    }
  }
}
switch_matching = {
  enable = true
  rules = [
    {
      match_model = "EX4300"
      match_name = "switch-*"
      match_role = "access"
      name = "access_switches"
      port_config = {
        "ge-0/0/0-23" = {
          usage = "lan"
        }
        "ge-0/0/24-47" = {
          usage = "uplink"
        }
      }
      stp_config = {
        bridge_priority = "32768"
      }
    }
  ]
}
switch_mgmt = {
  config_revert_timer = 10
  dhcp_option_fqdn = true
  mxedge_proxy_host = "192.168.1.30"
  mxedge_proxy_port = 2222
  root_password = "admin123"
  remove_existing_configs = false
  tacacs = {
    acct_servers = [
      {
        host = "192.168.1.15"
        port = 49
        secret = "tacacs_secret"
        timeout = 10
      }
    ]
    auth_servers = [
      {
        host = "192.168.1.15"
        port = 49
        secret = "tacacs_secret"
        timeout = 10
      }
    ]
    default_role = "read"
    enabled = true
    network = "lan"
    tacplus_servers = [
      {
        host = "192.168.1.15"
        port = 49
        secret = "tacacs_secret"
        timeout = 10
      }
    ]
  }
  use_mxedge_proxy = false
}
vrf_config = {
  enabled = true
}
vrf_instances = {
  "mgmt_vrf" = {
    networks = ["lan"]
    extra_routes = {
      "10.0.0.0/8" = {
        via = "192.168.1.1"
      }
    }
  }
  "guest_vrf" = {
    networks = ["guest"]
    extra_routes = {
      "0.0.0.0/0" = {
        via = "192.168.200.1"
      }
    }
  }
}
acl_policies = [
  {
    name = "allow_web_traffic"
    src_tags = ["workstations"]
    actions = [
      {
        action = "allow"
        dst_tag = "web_servers"
      }
    ]
  },
  {
    name = "deny_guest_to_lan"
    src_tags = ["guest_devices"]
    actions = [
      {
        action = "deny"
        dst_tag = "lan_devices"
      }
    ]
  }
]
acl_tags = {
  "workstations" = {
    type = "subnet"
    subnets = ["192.168.1.0/24"]
    network = "lan"
  }
  "web_servers" = {
    type = "subnet"
    subnets = ["192.168.10.0/24"]
    specs = [
      {
        protocol = "tcp"
        port_range = "443"
      }
    ]
  }
  "guest_devices" = {
    type = "network"
    network = "guest"
  }
  "lan_devices" = {
    type = "network"
    network = "lan"
  }
  "mac_based_tag" = {
    type = "mac"
    macs = ["aa:bb:cc:dd:ee:ff", "11:22:33:44:55:66"]
    network = "lan"
  }
  "radius_based_tag" = {
    type = "radius_group"
    radius_group = "engineers"
    specs = [
      {
        protocol = "tcp"
        port_range = "22"
      }
    ]
  }
}


