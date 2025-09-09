ssid = "TestWLAN-Comprehensive"
enabled = true
hide_ssid = false
band_steer = true
band_steer_force_band5 = false
bands = ["24", "5", "6"]
vlan_enabled = true
vlan_id = "100"
vlan_pooling = true
vlan_ids = ["101", "102", "103"]
isolation = false
l2_isolation = false
limit_bcast = true
allow_mdns = true
allow_ssdp = true
allow_ipv6_ndp = true
max_num_clients = 20
max_idletime = 1800
dtim = 2
fast_dot1x_timers = true
enable_local_keycaching = true
enable_wireless_bridging = true
enable_wireless_bridging_dhcp_tracking = true
arp_filter = true
block_blacklist_clients = true
disable_11ax = false
disable_ht_vht_rates = false
disable_uapsd = false
disable_wmm = false
disable_v1_roam_notify = false
disable_v2_roam_notify = false
disable_when_gateway_unreachable = false
disable_when_mxtunnel_down = false
hostname_ie = true
legacy_overds = false
limit_probe_response = false
no_static_dns = true
no_static_ip = true
reconnect_clients_when_roaming_mxcluster = false
sle_excluded = false
use_eapol_v1 = false
auth = {
  type = "eap"
  eap_reauth = true
  enable_mac_auth = false
  multi_psk_only = false
  private_wlan = false
  wep_as_secondary_auth = false
  anticlog_threshold = 16
  pairwise = ["wpa2-ccmp", "wpa3"]
  key_idx = 1
  keys = ["key1", "key2", "key3", "key4"]
}
auth_servers = [
  {
    host = "radius1.example.com"
    port = "1812"
    secret = "RadiusSecret123"
    keywrap_enabled = true
    keywrap_format = "hex"
    keywrap_kek = "kek123"
    keywrap_mack = "mack456"
    require_message_authenticator = true
  },
  {
    host = "radius2.example.com"
    port = "1812"
    secret = "RadiusSecret456"
    keywrap_enabled = false
    require_message_authenticator = true
  }
]
auth_server_selection = "ordered"
auth_servers_nas_id = "ap-nas-id"
auth_servers_nas_ip = "10.1.1.1"
auth_servers_retries = 3
auth_servers_timeout = 5
acct_immediate_update = true
acct_interim_interval = 600
acct_servers = [
  {
    host = "accounting1.example.com"
    port = "1813"
    secret = "AcctSecret123"
    keywrap_enabled = true
    keywrap_format = "ascii"
    keywrap_kek = "acctkek123"
    keywrap_mack = "acctmack456"
  },
  {
    host = "accounting2.example.com"
    port = "1813"
    secret = "AcctSecret456"
    keywrap_enabled = false
  }
]
interface = "all"
apply_to = "site"
roam_mode = "11r"
wlan_limit_down_enabled = true
wlan_limit_down = 50000
wlan_limit_up_enabled = true
wlan_limit_up = 10000
client_limit_down_enabled = true
client_limit_down = "5000"
client_limit_up_enabled = true
client_limit_up = "1000"
bonjour = {
  enabled = false
}
cisco_cwa = {
  enabled = true
  allowed_hostnames = ["captive.example.com", "auth.example.com"]
  allowed_subnets = ["10.0.0.0/8", "192.168.0.0/16"]
  blocked_subnets = ["10.10.10.0/24", "172.16.1.0/24"]
}
coa_servers = [
  {
    enabled = true
    ip = "10.1.1.10"
    port = "3799"
    secret = "CoASecret123"
    disable_event_timestamp_check = false
  },
  {
    enabled = true
    ip = "10.1.1.11"
    port = "3799"
    secret = "CoASecret456"
    disable_event_timestamp_check = true
  }
]
portal = {
  enabled = true
  auth = "none"
  expire = 86400
  privacy = true
  bypass_when_cloud_down = true
  email_enabled = false
  sms_enabled = false
  sponsor_enabled = false
  guest_portal_access = true
}
portal_allowed_hostnames = ["portal.example.com", "auth.example.com", "captive.example.com"]
portal_denied_hostnames = ["blocked.example.com", "malware.example.com"]
qos = {
  class = "voice"
  overwrite = true
}
app_limit = {
  enabled = true
  apps = {
    "facebook" = 1000
    "youtube" = 5000
    "netflix" = 8000
    "spotify" = 2000
  }
  wxtag_ids = {
    "streaming-apps" = 10000
    "social-media" = 5000
  }
}
app_qos = {
  enabled = true
  apps = {
    "video-streaming" = {
      dscp = "34"
      src_subnet = "10.0.0.0/8"
      dst_subnet = "0.0.0.0/0"
    }
    "voice-calls" = {
      dscp = "46"
      src_subnet = "192.168.0.0/16"
      dst_subnet = "0.0.0.0/0"
    }
  }
  others = [
    {
      dscp = "46"
      protocol = "tcp"
      port_ranges = "80,443"
      src_subnet = "10.0.0.0/8"
      dst_subnet = "0.0.0.0/0"
    },
    {
      dscp = "26"
      protocol = "udp"
      port_ranges = "53"
      src_subnet = "192.168.0.0/16"
      dst_subnet = "8.8.8.8/32"
    }
  ]
}
airwatch = {
  enabled = false
}
radsec = {
  enabled = true
  use_mxedge = true
  use_site_mxedge = true
  servers = [
    {
      host = "10.1.1.20"
      port = 2083
    },
    {
      host = "10.1.1.21"
      port = 2083
    }
  ]
}
rateset = {
  "24" = {
    ht = "6.5"
    legacy = "12"
    min_rssi = -75
    template = "default"
  }
  "5" = {
    ht = "6.5"
    legacy = "12"
    min_rssi = -70
    template = "high_density"
  }
  "6" = {
    ht = "6.5"
    legacy = "12"
    min_rssi = -65
    template = "enterprise"
  }
}
schedule = {
  enabled = true
  hours = {
    mon = "06:00-22:00"
    tue = "06:00-22:00"
    wed = "06:00-22:00"
    thu = "06:00-22:00"
    fri = "06:00-22:00"
    sat = "08:00-20:00"
    sun = "08:00-20:00"
  }
}
hotspot20 = {
  enabled = true
  domain_name = ["example.com", "corp.example.com"]
  nai_realms = ["@example.com", "@corp.example.com"]
  venue_name = "Corporate Headquarters"
  venue_type = "business"
  roaming_ois = ["001122", "334455", "667788"]
}
inject_dhcp_option_82 = {
  enabled = true
  circuit_id = "ap-name"
  remote_id = "ap-mac"
}
mist_nac = {
  enabled = true
}
wxtag_ids = ["corporate-devices", "iot-devices", "guest-devices"]
ap_ids = ["550e8400-e29b-41d4-a716-446655440001", "550e8400-e29b-41d4-a716-446655440002", "550e8400-e29b-41d4-a716-446655440003"]
mxtunnel_ids = ["tunnel-primary", "tunnel-backup"]
mxtunnel_name = ["primary-mx-tunnel", "backup-mx-tunnel"]
wxtunnel_id = "wx-tunnel-main"
wxtunnel_remote_id = "remote-wx-main"
