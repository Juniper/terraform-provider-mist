

  name    = "test-mxedge"
  model   = "ME-X5"

  # Optional attributes
  note              = "Test MxEdge for comprehensive attribute testing"
  for_site          = false
  mxagent_registered = false
  tunterm_registered = false
  mxcluster_id      = "00000000-0000-0000-0000-000000000000"

  # NTP servers configuration
  ntp_servers = [
    "time1.google.com",
    "time2.google.com"
  ]

  # Services configuration
  services = ["tunterm"]

  # MxEdge management configuration
  mxedge_mgmt = {
    config_auto_revert = true
    fips_enabled       = false
    mist_password      = "secure-password-123"
    root_password      = "root-password-456"
    oob_ip_type        = "static"
    oob_ip_type6       = "dhcp"
  }

  # Out-of-band IP configuration
  oob_ip_config = {
    type     = "static"
    ip       = "192.168.1.100"
    netmask  = "255.255.255.0"
    gateway  = "192.168.1.1"
    dns      = ["8.8.8.8", "8.8.4.4"]
    type6    = "static"
    ip6      = "2001:db8::100"
    netmask6 = "64"
    gateway6 = "2001:db8::1"
    autoconf6 = false
    dhcp6     = false
  }

  # Proxy configuration
  proxy = {
    url = "http://proxy.example.com:8080"
  }

  # Tunterm IP configuration
  tunterm_ip_config = {
    ip       = "10.0.0.1"
    netmask  = "255.255.255.0"
    gateway  = "10.0.0.254"
    ip6      = "fd00::1"
    netmask6 = "64"
    gateway6 = "fd00::254"
  }

  # Tunterm other IP configs per VLAN
  tunterm_other_ip_configs = {
    "100" = {
      ip      = "10.100.0.1"
      netmask = "255.255.255.0"
    }
    "200" = {
      ip      = "10.200.0.1"
      netmask = "255.255.255.0"
    }
  }

  # Tunterm extra routes
  tunterm_extra_routes = {
    "172.16.0.0/16" = {
      via = "10.0.0.10"
    }
    "192.168.100.0/24" = {
      via = "10.0.0.20"
    }
  }

  # Tunterm DHCP relay configuration
  tunterm_dhcpd_config = {
    "global" = {
      enabled = true
      type    = "relay"
      servers = ["10.0.0.5", "10.0.0.6"]
    }
    "100" = {
      enabled = true
      type    = "relay"
      servers = ["10.100.0.5"]
    }
  }

  # Versions
  versions = {
    mxagent = "0.1.234"
    tunterm = "0.2.345"
  }

