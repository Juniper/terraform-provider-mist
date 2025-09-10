
  name = "comprehensive_ap_profile"
  aeroscout {
    enabled = true
    host = "aeroscout.example.com"
    locate_connected = true
    port = 1144
  }
  airista {
    enabled = true
    host = "airista.example.com"
    port = 8080
  }
  ble_config {
    beacon_enabled = true
    beacon_rate = 5
    beacon_rate_mode = "custom"
    beam_disabled = [1, 3, 5]
    custom_ble_packet_enabled = true
    custom_ble_packet_frame = "0201061aff4c000215"
    custom_ble_packet_freq_msec = 1000
    ibeacon_enabled = true
    ibeacon_adv_power = -4
    ibeacon_freq_msec = 1000
    ibeacon_major = 123
    ibeacon_minor = 456
    ibeacon_uuid = "550e8400-e29b-41d4-a716-446655440000"
    power = 5
    power_mode = "custom"
  }
  disable_eth1 = false
  disable_eth2 = false
  disable_eth3 = false
  disable_module = false
  esl_config {
    enabled = true
    type = "imagotag"
    host = "esl.example.com"
    port = 443
    channel = 37
    verify_cert = true
    vlan_id = 1
  }
  ip_config {
    type = "static"
    ip = "192.168.1.10"
    netmask = "255.255.255.0"
    gateway = "192.168.1.1"
    dns = ["8.8.8.8", "8.8.4.4"]
    dns_suffix = ["example.com", "local"]
    mtu = 1500
    vlan_id = 100
  }
  lacp_config {
    enabled = true
  }
  led {
    enabled = true
    brightness = 200
  }
  mesh {
    enabled = true
    role = "base"
    group = 1
    bands = ["24", "5"]
  }
  ntp_servers = ["pool.ntp.org", "time.google.com"]
  poe_passthrough = true
  port_config = {
    "eth0" = {
      disabled = false
      forwarding = "all"
      enable_mac_auth = true
      mac_auth_preferred = false
      mac_auth_protocol = "eap-md5"
      port_vlan_id = 100
      dynamic_vlan = {
        enabled = true
        default_vlan_id = 1
        type = "standard"
        vlans = {
          "staff" = "10"
          "guest" = "20"
        }
      }
      mist_nac = {
        enabled = true
        network = "corp"
        acct_interim_interval = 3600
        auth_servers_retries = 3
        auth_servers_timeout = 30
        coa_enabled = true
        coa_port = 3799
      }
    }
  }

