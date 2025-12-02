
  ap_updown_threshold      = 30
  device_updown_threshold  = 45
  gateway_updown_threshold = 60
  switch_updown_threshold  = 75
  ui_idle_timeout          = 120
  disable_pcap             = false
  disable_remote_shell     = false

  cacerts = [
    "-----BEGIN CERTIFICATE-----\nMIIC2jCCAcICCQDXrcsSSdA+CjANBgkqhkiG9w0BAQsFADAhMR8wHQYDVQQDDBZy\nYWRzZWMtZXhhbXBsZS1kb21haW4xMB4XDTIzMDEwMTAwMDAwMFoXDTI0MDEwMTAw\nMDAwMFowITEfMB0GA1UEAwwWcmFkc2VjLWV4YW1wbGUtZG9tYWluMTCCASIwDQYJ\nKoZIhvcNAQEBBQADggEPADCCAQoCggEBAM3Q2NP+example+cert+data+here\n-----END CERTIFICATE-----"
  ]

  api_policy {
    no_reveal = true
  }

  celona {
    api_key    = "celona_api_key_example"
    api_prefix = "celona_api_prefix_example"
  }

  installer {
    allow_all_devices = true
    allow_all_sites   = false
    extra_site_ids    = ["550e8400-e29b-41d4-a716-446655440001", "550e8400-e29b-41d4-a716-446655440002"]
    grace_period      = 7200
  }

  jcloud_ra {
    org_apitoken_name = "jcloud_ra_org_token_name_example"
    org_id            = "jcloud_ra_org_id_example"
  }

  junos_shell_access {
    admin    = "admin"
    helpdesk = "viewer"
    read     = "viewer"
    write    = "admin"
  }

  marvis {
    auto_operations {
      ap_insufficient_capacity                     = true
      ap_loop                                      = true
      ap_non_compliant                            = true
      bounce_port_for_abnormal_poe_client         = true
      disable_port_when_ddos_protocol_violation   = true
      disable_port_when_rogue_dhcp_server_detected = true
      gateway_non_compliant                       = true
      switch_misconfigured_port                   = true
      switch_port_stuck                           = true
    }
  }

  mgmt {
    mxtunnel_ids = ["550e8400-e29b-41d4-a716-446655440003", "550e8400-e29b-41d4-a716-446655440004"]
    use_mxtunnel = true
    use_wxtunnel = false
  }

  mist_nac {
    disable_rsae_algorithms      = false
    eap_ssl_security_level       = 2
    eu_only                      = false
    idp_machine_cert_lookup_field = "cn"
    idp_user_cert_lookup_field    = "email"
    idps = [
      {
        id           = "550e8400-e29b-41d4-a716-446655440006"
        user_realms  = ["example.com", "test.org"]
        exclude_realms = ["excluded.example.com"]
      }
    ]
    # servier_cert
    use_ip_version = "v4"
    use_ssl_port   = false
  }

  mxedge_mgmt {
    config_auto_revert = true
    fips_enabled       = false
    oob_ip_type        = "dhcp"
    oob_ip_type6       = "dhcp"
  }

  optic_port_config = {
    "et-0/0/47" = {
      channelized = true
      speed       = "25g"
    }
    "et-0/0/48-49" = {
      channelized = false
      speed       = "50g"
    }
  }

  password_policy {
    enabled                   = true
    expiry_in_days           = 90
    min_length               = 12
    requires_special_char    = true
    requires_two_factor_auth = true
  }

  security {
    disable_local_ssh       = false
    limit_ssh_access        = true
  }

  switch_mgmt {
    ap_affinity_threshold = 15
  }

  ssr {
    proxy {
      disabled = true
      url      = "http://proxy.example.com:8080"
    }
  }

  synthetic_test {
    aggressiveness = "high"
    custom_probes = {
      "google_dns" = {
        type         = "icmp"
        host         = "8.8.8.8"
        threshold    = 100
        aggressiveness = "high"
      }
      "web_check" = {
        type         = "curl"
        url          = "https://www.google.com"
        threshold    = 500
        aggressiveness = "auto"
      }
      "tcp_check" = {
        type         = "tcp"
        host         = "10.1.1.1"
        port         = 443
        threshold    = 200
        aggressiveness = "med"
      }
    }
    disabled = false
    lan_networks = [
      {
        networks = ["192.168.1.0/24", "10.0.0.0/8"]
        probes   = ["google_dns", "web_check"]
      }
    ]
    vlans = [
      {
        disabled          = false
        custom_test_urls  = ["https://example.com/test"]
        probes           = ["google_dns", "tcp_check"]
        vlan_ids         = ["100", "200"]
      }
    ]
    wan_speedtest {
      enabled     = true
      time_of_day = "02:00"
    }
  }

  vpn_options {
    as_base   = 65000
    st_subnet = "10.224.0.0/12"
  }

  wan_pma {
    enabled = true
  }

  wired_pma {
    enabled = true
  }

  wireless_pma {
    enabled = true
  }
