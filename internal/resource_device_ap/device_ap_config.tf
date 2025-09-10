
  site_id      = "2c107c8e-2e06-404a-ba61-e25b5757ecea"
  device_id    = "00000000-0000-0000-1000-d420b041d5bf"
  name         = "test-ap"
  x            = 150.0
  y            = 250.0
  height       = 2.5
  orientation  = 90
  notes        = "Comprehensive test AP configuration with all attributes"
  locked       = false
  disable_eth1 = true
  disable_eth2 = false
  disable_eth3 = false
  disable_module = false
  poe_passthrough = false
  flow_control = false
  // map_id = "845a23bf-bed9-e43c-4c86-6fa474be7ae5"
  ntp_servers = [
    "pool.ntp.org",
    "time.google.com",
    "time.cloudflare.com"
  ]
  vars = {
    "environment" = "test"
    "location"    = "building-a"
    "floor"       = "2"
    "zone"        = "guest"
    "deployment"  = "comprehensive"
  }
  aeroscout = {
    enabled = true
    host = "aeroscout.example.com"
    locate_connected = true
    port = 3001
  }
  airista = {
    enabled = true
    host = "airista.example.com"
    port = 3004
  }
  ble_config = {
    beacon_enabled = true
    beacon_rate = 1
    beacon_rate_mode = "custom"
    beam_disabled = [1, 3, 5]
    custom_ble_packet_enabled = true
    custom_ble_packet_frame = "1234567890abcdef"
    custom_ble_packet_freq_msec = 1000
    ibeacon_enabled = true
    ibeacon_adv_power = -12
    ibeacon_beams = "2-7"
    ibeacon_freq_msec = 1500
    ibeacon_major = 100
    ibeacon_minor = 200
    ibeacon_uuid = "550e8400-e29b-41d4-a716-446655440000"
    power = 5
    power_mode = "custom"
  }
  centrak = {
    enabled = false
  }
  client_bridge = {
    enabled = true
    ssid = "guest-network"
    auth = {
      type = "psk"
      psk = "secretpassword123"
    }
  }
  esl_config = {
    enabled = true
    type = "native"
    host = "esl.example.com"
    port = 443
    channel = 1
    cacert = "-----BEGIN CERTIFICATE-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END CERTIFICATE-----"
  }
  ip_config = {
    type    = "static"
    ip      = "192.168.1.100"
    netmask = "255.255.255.0"
    gateway = "192.168.1.1"
    dns     = ["8.8.8.8", "8.8.4.4"]
    dns_suffix = ["corp.example.com", "example.com"]
    type6 = "dhcp"
    mtu = 1500
    vlan_id = 100
  }
  lacp_config = {
    enabled = true
  }
  led = {
    enabled = true
    brightness = 75
  }
  mesh = {
    enabled = true
    role    = "base"
    bands   = ["5", "6"]
    group   = 1
  }
  port_config = {
    "ge-0/0/0" = {
      disabled = false
      dynamic_vlan = {
        enabled = true
        type = "standard"
        default_vlan_id = 1
        vlans = {
          "1" = "default"
          "100" = "guest"
          "200" = "iot"
        }
      }
      enable_mac_auth = true
      forwarding = "all"
      mac_auth_preferred = false
      mac_auth_protocol = "pap"
      mist_nac = {
        enabled = false
        network = "default"
        auth_servers_timeout = 5
        auth_servers_retries = 3
        acct_interim_interval = 600
        coa_enabled = false
        coa_port = 3799
        fast_dot1x_timers = false
        source_ip = "192.168.1.100"
      }
      mx_tunnel_id = "550e8400-e29b-41d4-a716-446655440005"
      port_auth = "dot1x"
      port_vlan_id = 100
      radius_config = {
        network = "default"
        acct_interim_interval = 600
        auth_servers_timeout = 5
        auth_servers_retries = 3
        coa_enabled = false
        coa_port = 3799
        source_ip = "192.168.1.100"
        acct_servers = []
        auth_servers = [
          {
            host = "radius-auth.example.com"
            port = "1812"
            secret = "MySecureTestSecret123!"
            keywrap_enabled = false
            keywrap_format = "hex"
            keywrap_kek = "0123456789abcdef0123456789abcdef"
            keywrap_mack = "fedcba9876543210fedcba9876543210"
            require_message_authenticator = false
          }
        ]
      }
      radsec = {
        enabled = false
        idle_timeout = "60"
        server_name = "radsec.example.com"
        use_mxedge = false
        use_site_mxedge = false
        coa_enabled = false
        mxcluster_ids = []
        proxy_hosts = []
        servers = []
      }
      vlan_id = 100
      vlan_ids = [100, 200, 300]
      wxtunnel_id = "550e8400-e29b-41d4-a716-446655440006"
    }
  }
  pwr_config = {
    base = 16000
    prefer_usb_over_wifi = false
  }
  radio_config = {
    allow_rrm_disable = true
    ant_gain_24 = 0
    ant_gain_5 = 2
    ant_gain_6 = 3
    antenna_mode = "default"
    band_24_usage = "auto"
    full_automatic_rrm = false
    indoor_use = true
    scanning_enabled = true
    band_24 = {
      power             = 10
      channel           = 6
      channels          = [1, 6, 11]
      bandwidth         = 20
      allow_rrm_disable = false
      disabled          = false
      ant_gain          = 0
      antenna_mode      = "default"
      power_max         = 16
      power_min         = 10
      preamble          = "auto"
    }
    band_5 = {
      power             = 15
      channel           = 36
      channels          = [36, 40, 44, 48]
      bandwidth         = 80
      allow_rrm_disable = false
      disabled          = false
      ant_gain          = 2
      antenna_mode      = "default"
      power_max         = 16
      power_min         = 12
      preamble          = "auto"
    }
    band_6 = {
      power             = 18
      channel           = 37
      channels          = [37, 41, 45, 49]
      bandwidth         = 160
      allow_rrm_disable = false
      disabled          = false
      ant_gain          = 3
      antenna_mode      = "default"
      power_max         = 17
      power_min         = 10
      preamble          = "auto"
      standard_power    = true
    }
  }
  uplink_port_config = {
    dot1x = false
    keep_wlans_up_if_down = true
  }
  usb_config = {
    enabled = true
    type = "imagotag"
    host = "usb.example.com"
    port = 443
    channel = 2
    cacert = "-----BEGIN CERTIFICATE-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END CERTIFICATE-----"
    verify_cert = true
  }
