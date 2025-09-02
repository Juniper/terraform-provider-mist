// Comprehensive WLAN - Maximum field coverage without conflicting constraints
  ssid = "Comprehensive_WLAN"
  enabled = true
  hide_ssid = true
  isolation = true
  acct_immediate_update = true
  acct_interim_interval = 600
  allow_ipv6_ndp = true
  allow_mdns = true
  allow_ssdp = true
  ap_ids = ["00000000-0000-0000-1000-5c5b35000010", "00000000-0000-0000-1000-5c5b35000011"]
  apply_to = "site"
  bands = ["24", "5", "6"]
  arp_filter = true
  band_steer = true
  band_steer_force_band5 = true
  block_blacklist_clients = true
  disable_11ax = true
  disable_ht_vht_rates = true
  disable_uapsd = true
  disable_v1_roam_notify = true
  disable_v2_roam_notify = true
  disable_when_gateway_unreachable = true
  disable_when_mxtunnel_down = true
  disable_wmm = true
  dtim = 3
  enable_local_keycaching = true
  enable_wireless_bridging = true
  enable_wireless_bridging_dhcp_tracking = true
  fast_dot1x_timers = true
  hostname_ie = true
  l2_isolation = true
  legacy_overds = true
  limit_bcast = true
  limit_probe_response = true
  max_idletime = 7200
  max_num_clients = 128
  no_static_dns = true
  no_static_ip = true
  reconnect_clients_when_roaming_mxcluster = true
  roam_mode = "11r"
  sle_excluded = true
  use_eapol_v1 = true
  vlan_enabled = true
  vlan_id = "100"
  vlan_pooling = true
  vlan_ids = ["10", "20", "30", "40"]
  
  auth = {
    type = "eap"
    enable_mac_auth = true
    eap_reauth = true
    private_wlan = true
    wep_as_secondary_auth = true
  }
  
  auth_servers = [
    {
      host = "radius1.company.com"
      port = 1812
      secret = "radius-secret"
      require_message_authenticator = true
    },
    {
      host = "radius2.company.com"
      port = 1812
      secret = "radius-secret-2"
      require_message_authenticator = true
    }
  ]
  
  acct_servers = [
    {
      host = "radius1.company.com"
      port = 1813
      secret = "radius-secret"
    },
    {
      host = "radius2.company.com"
      port = 1813
      secret = "radius-secret-2"
    }
  ]
  
  coa_servers = [
    {
      ip = "192.168.1.50"
      port = "3799"
      secret = "coa-secret-1"
      enabled = true
      disable_event_timestamp_check = false
    },
    {
      ip = "192.168.1.51"
      port = "3799"
      secret = "coa-secret-2"
      enabled = true
      disable_event_timestamp_check = true
    }
  ]
  
  portal_allowed_hostnames = ["portal.company.com", "guest.company.com", "wifi.company.com"]
  portal_allowed_subnets   = ["192.168.1.0/24", "10.0.0.0/8", "{{orgsubnetsguest}}"]
  portal_denied_hostnames  = ["blocked.example.com", "malware.test.com"]
  
  airwatch = {
    api_key = "test-api-key"
    console_url = "https://console.awmdm.com"
    enabled = true
    password = "test-password"
    username = "test-user"
  }
  
  # app_limit = {
  #   enabled = true
  #   apps = {
  #     facebook = 1000
  #     youtube = 500
  #   }
  #   wxtag_ids = {
  #     "tag1" = 2000
  #   }
  # }
  
  bonjour = {
    enabled = true
    additional_vlan_ids = ["10", "20"]
    services = {
      airplay = {
        disable_local = false
        radius_groups = ["group1", "group2"]
        scope = "same_site"
      }
      homekit = {
        disable_local = true
        radius_groups = ["group3"]
        scope = "same_ap"
      }
    }
  }
  
  cisco_cwa = {
    enabled = true
    allowed_hostnames = ["captive-portal.company.com", "wifi-login.company.com"]
    allowed_subnets = ["192.168.100.0/24", "10.0.0.0/8"]
    blocked_subnets = ["192.168.99.0/24", "172.16.0.0/12"]
  }
  
  dns_server_rewrite = {
    enabled = true
    radius_groups = {
      "employees" = "8.8.8.8"
      "guests" = "1.1.1.1"
      "contractors" = "208.67.222.222"
    }
  }
  
  dynamic_psk = {
    enabled = true
    source = "radius"
    default_psk = "fallback123"
    default_vlan_id = "100"
    force_lookup = true
  }
  
  mist_nac = {
    enabled = true
  }
  
  radsec = {
    enabled = true
    coa_enabled = true
    idle_timeout = "300"
    server_name = "radsec.company.com"
    use_mxedge = true
    use_site_mxedge = false
    mxcluster_ids = ["00000000-0000-0000-4000-5c5b35000400", "00000000-0000-0000-4000-5c5b35000401"]
    proxy_hosts = ["proxy1.company.com", "proxy2.company.com"]
    servers = [
      {
        host = "192.168.10.100"
        port = 2083
      },
      {
        host = "192.168.10.101"
        port = 2083
      }
    ]
  }
  
  portal = {
    enabled = true
    auth = "password"
    expire = 86400
    password = "GuestPassword123"
    bypass_when_cloud_down = true
    cross_site = true
    email_enabled = true
    forward = false
    privacy = true
    sms_enabled = true
    sponsor_enabled = true
    predefined_sponsors_enabled = true
    predefined_sponsors_hide_email = true
  }
  
  app_qos = {
    enabled = true
    apps = {
      teams = {
        dscp = 34
      }
      zoom = {
        dscp = 34
      }
      webex = {
        dscp = 34
      }
    }
  }
  
  qos = {
    class = "video"
    overwrite = true
  }
  
  rateset = {
    "24" = {
      template = "custom"
      min_rssi = -70
      legacy = ["6b", "9", "12b", "18", "24b", "36", "48", "54"]
    }
    "5" = {
      template = "custom"
      min_rssi = -65
      legacy = ["6b", "9", "12b", "18", "24b", "36", "48", "54"]
      ht = "00ff"
      vht = "03ff"
    }
    "6" = {
      template = "custom"
      min_rssi = -60
      he = "0fff"
      eht = "ffff"
    }
  }
  
  schedule = {
    enabled = true
    hours = {
      mon = "08:00-18:00"
      tue = "08:00-18:00"
      wed = "08:00-18:00"
      thu = "08:00-18:00"
      fri = "08:00-18:00"
      sat = "10:00-16:00"
      sun = "10:00-16:00"
    }
  }
  
  client_limit_down_enabled = true
  client_limit_down = 50
  client_limit_up_enabled = true
  client_limit_up = 25
  wlan_limit_down_enabled = true
  wlan_limit_down = 1000
  wlan_limit_up_enabled = true
  wlan_limit_up = 500
  wxtunnel_id = "00000000-0000-0000-1000-5c5b35000100"
  wxtunnel_remote_id = "remote-tunnel-123"
  mxtunnel_ids = ["00000000-0000-0000-2000-5c5b35000200", "00000000-0000-0000-2000-5c5b35000201"]
  mxtunnel_name = ["primary-mx-tunnel", "backup-mx-tunnel"]
  wxtag_ids = ["00000000-0000-0000-3000-5c5b35000300", "00000000-0000-0000-3000-5c5b35000301", "00000000-0000-0000-3000-5c5b35000302"]
  
  interface = "eth0"
  auth_server_selection = "ordered"
  auth_servers_nas_id = "mist-ap-comprehensive"
  auth_servers_nas_ip = "192.168.1.100"
  auth_servers_retries = 5
  auth_servers_timeout = 10
