
  device_id = "00000000-0000-0000-1000-5c5b35000032"
  name      = "test-switch-comprehensive"
  additional_config_cmds = ["set system host-name switch1", "set system domain-name example.com"]
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
  port_config = {
    "eth0" = "up"
    "eth1" = "down"
    "eth2" = "up"
  }
  image_url = "https://example.com/switch.png"
  notes     = "Comprehensive switch config"
  model     = "QFX5100"
  serial    = "SW-1234567890"
  x         = 100.5
  y         = 200.75

