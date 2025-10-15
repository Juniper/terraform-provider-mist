  analytic = {
    enabled = true
  }
  
  ap_updown_threshold = 5
  
  auto_upgrade = {
    custom_versions = {
      "AP41" = "0.14.26348"
      "AP43" = "0.14.26348"
    }
    day_of_week  = "sun"
    enabled      = true
    time_of_day  = "02:00"
    version      = "stable"
  }
  
  bgp_neighbor_updown_threshold = 10
  blacklist_url                 = "https://blacklist.example.com"
  
  ble_config = {
    beacon_enabled              = true
    beacon_rate_mode           = "custom"
    beacon_rate                 = 3
    beam_disabled              = [1, 3, 6]
    custom_ble_packet_enabled  = true
    custom_ble_packet_frame    = "0x12345678"
    custom_ble_packet_freq_msec = 300
    ibeacon_adv_power          = -16
    ibeacon_beams              = "2-7"
    ibeacon_enabled            = true
    ibeacon_freq_msec          = 1000
    ibeacon_major              = 13
    ibeacon_minor              = 138
    ibeacon_uuid               = "f3f17139-704a-f03a-2786-0400279e37c3"
    power                      = 6
    power_mode                 = "default"
  }
  
  config_auto_revert = false
  
  config_push_policy = {
    no_push = false
    push_window = {
      enabled = true
      hours = {
        fri = "02:00-08:00,13:00-17:00"
        mon = "02:00-08:00,13:00-17:00"
        sat = "02:00-08:00"
        sun = "02:00-08:00"
        thu = "02:00-08:00,13:00-17:00"
        tue = "02:00-08:00,13:00-17:00"
        wed = "02:00-08:00,13:00-17:00"
      }
    }
  }
  
  critical_url_monitoring = {
    enabled = true
    monitored_urls = [
      {
        url = "https://www.google.com"
      },
      {
        url = "https://www.youtube.com"
      }
    ]
  }
  
  default_port_usage   = "wan"
  device_updown_threshold = 180
  enable_unii_4          = false
  
  engagement = {
    dwell_tag_names = {
      bounce   = "Bounce"
      engaged  = "Engaged"
      passerby = "Passerby"
      stationed = "Stationed"
    }
    dwell_tags = {
      bounce   = "1-300"
      engaged  = "18000-43200"
      passerby = "0-0"
      stationed = "43200-"
    }
    hours = {
      fri = "09:00-17:00"
      mon = "09:00-17:00"
      sat = "09:00-12:00"
      sun = "closed"
      thu = "09:00-17:00"
      tue = "09:00-17:00"
      wed = "09:00-17:00"
    }
  }
  
  gateway_updown_threshold = 20
  
  juniper_srx = {
    antivirus_enabled       = true
    apply_to                = ["WAN-SRX-1"]
    ike_allowed_remote_ips = ["192.168.1.0/24"]
    idp_fingerprint_update  = true
    idp_enabled            = true
    intrusion_detection     = true
    utm_idp_features       = ["application-firewall", "application-identification"]
    auto_upgrade = {
      enabled = true
      snapshot = false
      custom_versions = {
        "WAN-SRX-1" = "21.4R3.15"
        "WAN-SRX-2" = "21.4R3.15"
      }
    }
  }
  
  led = {
    brightness = 255
    enabled    = true
  }
  
  marvis = {
    vble_enabled = true
  }
  
  occupancy = {
    assets_enabled    = true
    clients_enabled   = true
    min_duration      = 3000
    sdkclients_enabled = false
    unconnected_clients_enabled = true
  }
  
  persist_config_on_device = false
  
  proxy = {
    url = "http://proxy.example.com:8080"
  }
  
  remove_existing_configs = false
  report_gatt             = false
  
  rogue = {
    enabled              = true
    honeypot_enabled     = false
    min_duration         = 10
    min_rssi            = -80
    whitelisted_bssids  = ["aa:bb:cc:dd:ee:ff"]
    whitelisted_ssids   = ["Test-SSID"]
  }
  
  rtsa = {
    app_waking          = false
    disable_dead_reckoning = false
    disable_pressure_sensor = false
    enabled             = true
    track_asset         = true
  }
  
  simple_alert = {
    ap_offline_notification     = true
    ap_offline_timeout          = 5
    switch_offline_notification = true
    switch_offline_timeout      = 5
  }
  
  skyatp = {
    enabled = true
    send_ip_mac_mapping = true
  }
  
  sle_thresholds = {
    ap_health = {
      cpu_util    = 80
      mem_util    = 85
      temperature = 68
    }
    capacity = {
      max_capacity = 500
    }
    coverage = {
      coverage = -70
    }
    roaming = {
      failure_reasons = [
        {
          reason    = "auth-or-assoc-failure"
          threshold = 10
        }
      ]
    }
    successful_connect = {
      auth_time_threshold   = 10000
      dhcp_time_threshold   = 15000
      dns_time_threshold    = 8000
      failure_reasons = [
        {
          reason    = "auth-or-assoc-failure"
          threshold = 10
        }
      ]
    }
    throughput = {
      download = 10.0
      upload   = 5.0
    }
    time_to_connect = {
      failure_reasons = [
        {
          reason    = "auth-or-assoc-failure"
          threshold = 10
        }
      ]
      threshold = 10000
    }
  }
  
  srx_app = {
    enabled = true
  }
  
  ssr = {
    proxy = {
      url = "http://ssr-proxy.example.com:8080"
    }
    auto_upgrade = {
      enabled = true
      channel = "stable"
      custom_versions = {
        "router-1" = "6.2.3"
        "router-2" = "6.2.3"
      }
    }
  }
  
  ssh_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC..."]
  
  switch_updown_threshold = 15
  
  synthetic_test = {
    disabled = false
    aggressiveness = "auto"
    custom_probes = {
      "probe1" = {
        type = "icmp"
        host = "8.8.8.8"
        threshold = 100
        aggressiveness = "auto"
      }
      "probe2" = {
        type = "tcp"
        host = "google.com"
        port = 80
        threshold = 200
        aggressiveness = "med"
      }
    }
    vlans = [
      {
        vlan_ids = ["100", "200"]
        disabled = false
        custom_test_urls = ["https://example.com", "https://google.com"]
      }
    ]
    wan_speedtest = {
      enabled = true
      time_of_day = "02:00"
    }
  }
  
  track_anonymous_devices = true
  
  uplink_port_config = {
    dot1x_enabled      = false
    keep_wlans_up_if_down = false
  }
  
  vars = {
    "custom_var1" = "value1"
    "custom_var2" = "value2"
  }
  
  vna = {
    enabled = true
  }
  
  vpn_path_updown_threshold = 30
  vpn_peer_updown_threshold = 25
  
  wan_vna = {
    enabled = true
  }
  
  watched_station_url = "https://watched.example.com"
  whitelist_url       = "https://whitelist.example.com"
  
  wids = {
    repeated_auth_failures = {
      duration  = 60
      threshold = 10
    }
  }
  
  wifi = {
    ap_affinity_threshold   = 12
    beacon_protection       = false
    beacon_protection_enabled = false
    band_steering_enabled   = true
    cisco_enabled          = false
    disable_11k            = false
    disable_radios_when_power_constrained = false
    enable_arp_spoof_check = false
    enable_shared_radio_scanning = true
    enabled                = true
    locate_connected       = true
    locate_unconnected     = false
    mesh_allow_dfs         = false
    mesh_enable_crm        = false
    mesh_enabled           = false
    mesh_psk               = "meshpassword123"
    mesh_ssid              = "MeshNetwork"
    proxy_arp              = "default"
  }
  
  wired_vna = {
    enabled = true
  }
  
  zone_occupancy_alert = {
    email_notifiers = ["admin@example.com"]
    enabled         = true
    threshold       = 5
  }
