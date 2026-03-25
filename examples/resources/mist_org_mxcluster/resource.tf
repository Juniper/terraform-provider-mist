



resource "mist_org_mxcluster" "new_mxcluster" {
  name    = "test-mxcluster"
  org_id  = mist_org.terraform_test.id
  site_id = mist_site.terraform_test.id
  
  # Mist DAS (Dynamic Authorization Service) configuration
  mist_das = {
    enabled = true
    coa_servers = [
      {
        disable_event_timestamp_check  = false
        enabled                        = true
        host                          = "10.10.10.10"
        port                          = 3799
        require_message_authenticator = true
        secret                        = "coa-secret-123"
      },
      {
        enabled = true
        host    = "10.10.10.11"
        port    = 3799
        secret  = "coa-secret-456"
      }
    ]
  }

  # Mist NAC (Network Access Control) configuration
  mist_nac = {
    enabled         = true
    acct_server_port = 1813
    auth_server_port = 1812
    secret          = "nac-shared-secret"
  }

  # MxEdge management configuration
  mxedge_mgmt = {
    config_auto_revert = true
    fips_enabled       = false
    mist_password      = "mist-password-123"
    root_password      = "root-password-456"
    oob_ip_type        = "static"
    oob_ip_type6       = "dhcp"
  }

  # Proxy configuration
  proxy = {
    disabled = false
    url      = "http://proxy.example.com:8080"
  }

  # RadSec (RADIUS over TLS) configuration
  radsec = {
    enabled          = true
    match_ssid       = true
    nas_ip_source    = "tunnel"
    server_selection = "ordered"
    src_ip_source    = "tunnel"
    proxy_hosts      = ["radsec1.example.com", "radsec2.example.com"]
    
    acct_servers = [
      {
        host   = "acct1.example.com"
        port   = 2083
        secret = "acct-secret-123"
        ssids  = ["Corporate", "Guest"]
      },
      {
        host   = "acct2.example.com"
        port   = 2083
        secret = "acct-secret-456"
        ssids  = ["Corporate"]
      }
    ]
    
    auth_servers = [
      {
        host                   = "auth1.example.com"
        port                   = 2083
        secret                 = "auth-secret-123"
        inband_status_check    = true
        inband_status_interval = 60
        keywrap_enabled        = true
        keywrap_format         = "hex"
        keywrap_kek            = "keywrap-kek-value"
        keywrap_mack           = "keywrap-mack-value"
        retry                  = 3
        timeout                = 30
        ssids                  = ["Corporate", "Guest"]
      },
      {
        host    = "auth2.example.com"
        port    = 2083
        secret  = "auth-secret-456"
        retry   = 3
        timeout = 30
        ssids   = ["Corporate"]
      }
    ]
  }

  # Tunterm AP subnets - list of subnets where APs can establish Mist Tunnels from
  tunterm_ap_subnets = [
    "192.168.10.0/24",
    "192.168.20.0/24",
    "10.100.0.0/16"
  ]

  # Tunterm DHCP relay configuration per VLAN
  tunterm_dhcpd_config = {
    "100" = {
      enabled = true
      type    = "relay"
      servers = ["10.100.0.5", "10.100.0.6"]
    }
    "200" = {
      enabled = true
      type    = "relay"
      servers = ["10.200.0.5"]
    }
  }

  # Tunterm extra routes - additional routes for Mist Tunneled VLANs
  tunterm_extra_routes = {
    "172.16.0.0/16" = {
      via = "10.0.0.10"
    }
    "192.168.100.0/24" = {
      via = "10.0.0.20"
    }
  }

  # Tunterm hosts - hostnames or IPs for Mist Tunnel peers
  tunterm_hosts = [
    "mxedge1.example.com",
    "mxedge2.example.com",
    "10.10.10.100"
  ]

  # Tunterm hosts order - list of indices for tunterm_hosts
  tunterm_hosts_order = [0, 1, 2]

  # Tunterm hosts selection strategy
  tunterm_hosts_selection = "ordered"

  # Tunterm monitoring configuration - list of monitoring groups (each group is a list)
  tunterm_monitoring = [
    [  # First monitoring group
      {
        host        = "10.0.0.1"
        port        = 443
        protocol    = "https"
        src_vlan_id = 100
        timeout     = 10
      },
      {
        host        = "10.0.0.2"
        port        = 80
        protocol    = "http"
        src_vlan_id = 100
        timeout     = 5
      }
    ],
    [  # Second monitoring group
      {
        host        = "8.8.8.8"
        port        = 443
        protocol    = "https"
        src_vlan_id = 200
        timeout     = 10
      }
    ]
  ]

  # Tunterm monitoring disabled flag
  tunterm_monitoring_disabled = false
}
