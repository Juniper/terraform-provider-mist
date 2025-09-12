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
  disable_11be = true
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
  
  # Authentication configuration for dynamic VLAN support
  auth = {
    type = "eap"
    enable_mac_auth = true
    pairwise = ["wpa2-ccmp", "wpa3-ccmp"]
    anticlog_threshold = 20
    eap_reauth = true
    key_idx = 1
    keys = ["1234567890", "abcdef1234", "0987654321", "fedcba9876"]
    multi_psk_only = true
    private_wlan = true
    psk = "MySecurePassword123!"
    wep_as_secondary_auth = true
  }
  
  # Dynamic VLAN configuration
  dynamic_vlan = {
    enabled = true
    type = "standard"
    default_vlan_ids = ["100", "200-250"]
    local_vlan_ids = ["10", "20"]
    vlans = {
      "100" = ""
      "200" = ""
      "300" = ""
    }
  }
  
  auth_servers = [
    {
      host = "radius1.company.com"
      port = 1812
      secret = "radius-secret"
      require_message_authenticator = true
      keywrap_enabled = true
      keywrap_format = "hex"
      keywrap_kek = "abcdef1234567890"
      keywrap_mack = "0987654321fedcba"
    },
    {
      host = "radius2.company.com"
      port = 1812
      secret = "radius-secret-2"
      require_message_authenticator = false
      keywrap_enabled = false
    }
  ]
  
  acct_servers = [
    {
      host = "radius1.company.com"
      port = 1813
      secret = "radius-secret"
      keywrap_enabled = true
      keywrap_format = "hex"
      keywrap_kek = "1234567890abcdef"
      keywrap_mack = "fedcba0987654321"
      require_message_authenticator = true
    },
    {
      host = "radius2.company.com"
      port = 1813
      secret = "radius-secret-2"
      keywrap_enabled = false
      keywrap_format = "ascii"
      require_message_authenticator = false
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
  
  portal = {
    enabled = true
    auth = "sso"
    
    # Social login providers - enable all with test values
    amazon_enabled = true
    amazon_client_id = "test-amazon-client-id"
    amazon_client_secret = "test-amazon-client-secret"
    amazon_expire = 1440
    amazon_email_domains = ["amazon.com", "aws.com"]
    
    azure_enabled = true
    azure_client_id = "test-azure-client-id"
    azure_client_secret = "test-azure-client-secret"
    azure_expire = 1440
    azure_tenant_id = "test-azure-tenant-id"
    
    facebook_enabled = true
    facebook_client_id = "test-facebook-client-id"
    facebook_client_secret = "test-facebook-client-secret"
    facebook_expire = 1440
    facebook_email_domains = ["facebook.com", "meta.com"]
    
    google_enabled = true
    google_client_id = "test-google-client-id"
    google_client_secret = "test-google-client-secret"
    google_expire = 1440
    google_email_domains = ["google.com", "gmail.com"]
    
    microsoft_enabled = true
    microsoft_client_id = "test-microsoft-client-id"
    microsoft_client_secret = "test-microsoft-client-secret"
    microsoft_expire = 1440
    microsoft_email_domains = ["microsoft.com", "outlook.com"]
    
    # SMS providers - all enabled for maximum coverage
    sms_enabled = true
    sms_expire = 60
    sms_provider = "twilio"
    sms_message_format = "Code: {{code}}"
    
    # Twilio configuration
    twilio_auth_token = "test-twilio-auth-token"
    twilio_phone_number = "+1234567890"
    twilio_sid = "test-twilio-sid"
    
    # Other SMS providers
    clickatell_api_key = "test-clickatell-api-key"
    gupshup_password = "test-gupshup-password"
    gupshup_userid = "test-gupshup-userid"
    puzzel_password = "test-puzzel-password"
    puzzel_service_id = "test-puzzel-service"
    puzzel_username = "test-puzzel-user"
    telstra_client_id = "test-telstra-client-id"
    telstra_client_secret = "test-telstra-client-secret"
    
    # Broadnet configuration
    broadnet_password = "test-broadnet-password"
    broadnet_sid = "test-broadnet-sid"
    broadnet_user_id = "test-broadnet-user"
    
    # Email and passphrase authentication
    email_enabled = true
    passphrase_enabled = true
    passphrase_expire = 720
    
    # Sponsor settings - enable all features
    sponsor_enabled = true
    sponsor_expire = 2880
    sponsor_auto_approve = false
    sponsor_notify_all = true
    sponsor_status_notify = true
    sponsor_link_validity_duration = "7d"
    predefined_sponsors_enabled = true
    predefined_sponsors_hide_email = false
    sponsors = {
      "sponsor1" = "sponsor1@company.com"
      "sponsor2" = "sponsor2@company.com"
      "admin" = "admin@company.com"
    }
    sponsor_email_domains = ["company.com", "partners.com"]
    
    # SSO configuration
    sso_default_role = "user"
    sso_forced_role = "guest"
    sso_idp_cert = "-----BEGIN CERTIFICATE-----\nMIICertificateContent\n-----END CERTIFICATE-----"
    sso_idp_sign_algo = "sha256"
    sso_idp_sso_url = "https://sso.company.com/saml/login"
    sso_issuer = "company-issuer"
    sso_nameid_format = "email"
    
    # Portal behavior settings
    allow_wlan_id_roam = true
    bypass_when_cloud_down = true
    cross_site = true
    forward = true
    forward_url = "https://portal.company.com/welcome"
    external_portal_url = "https://external-portal.company.com"
    privacy = true
    
    # Session settings
    expire = 1440
    password = "portal-admin-password"
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
  
  hotspot20 = {
    enabled = true
    domain_name = ["example.com", "company.com"]
    nai_realms = ["realm1.example.com", "realm2.company.com"]
    operators = ["Example Corp", "Company Ltd"]
    rcoi = ["00-11-22", "33-44-55"]
    venue_name = "Corporate WiFi Network"
  }

  inject_dhcp_option_82 = {
    enabled = true
    circuit_id = "site-{{site_name}}-ap-{{ap_name}}"
  }

  app_limit = {
    enabled = true
    apps = {
      "facebook" = 100
      "youtube" = 200
      "netflix" = 150
    }
    wxtag_ids = {
      "high_priority" = 300
      "low_priority" = 50
    }
  }

  app_qos = {
          apps = {
            teams = {
              dscp       = 34
              dst_subnet = "192.168.10.0/24"
              src_subnet = "10.1.0.0/16"
            }
            webex = {
              dscp       = 34
              dst_subnet = "192.168.11.0/24"
              src_subnet = "10.2.0.0/16"
            }
            zoom = {
              dscp       = 34
              dst_subnet = "192.168.12.0/24"
              src_subnet = "10.3.0.0/16"
            }
          }
          enabled = true
          others = [
            {
              dscp        = 46
              dst_subnet  = "192.168.99.0/24"
              port_ranges = "8000-9000,9500-9600"
              protocol    = "tcp"
              src_subnet  = "10.99.0.0/16"
            }
          ]
        }
        qos = {
          class     = "video"
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
␞
// External Portal Test Case - portal.auth set to external to test portal_api_secret
  ssid = "External_Portal_Test_WLAN"
  enabled = true
  
  portal = {
    enabled = true
    auth = "external"
    external_portal_url = "https://external-portal.example.com"
  }
␞
// OWE Test Case - auth.owe field for open authentication
  ssid   = "OWE_Test_WLAN"
  
  auth = {
    type = "open"
    owe  = "enabled"
  }
