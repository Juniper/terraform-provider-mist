---
page_title: "mist_org_mxcluster Resource - terraform-provider-mist"
subcategory: "Devices"
description: |-
  This resource manages MxCluster (cluster of MxEdge devices) in the Mist Organization.
  A Mist Edge Cluster is a group of one or more Org Mist Edge devices (mist_org_mxedge) providing tunnel termination, edge network services and RADIUS proxy capabilities.
---

# mist_org_mxcluster (Resource)

This resource manages MxCluster (cluster of MxEdge devices) in the Mist Organization.

A Mist Edge Cluster is a group of one or more Org Mist Edge devices (mist_org_mxedge) providing tunnel termination, edge network services and RADIUS proxy capabilities.


## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `org_id` (String)

### Optional

- `mist_das` (Attributes) Configure cloud-assisted dynamic authorization service on this cluster of mist edges (see [below for nested schema](#nestedatt--mist_das))
- `mist_nac` (Attributes) (see [below for nested schema](#nestedatt--mist_nac))
- `mxedge_mgmt` (Attributes) (see [below for nested schema](#nestedatt--mxedge_mgmt))
- `proxy` (Attributes) Proxy Configuration to talk to Mist (see [below for nested schema](#nestedatt--proxy))
- `radsec` (Attributes) MxEdge RadSec Configuration (see [below for nested schema](#nestedatt--radsec))
- `site_id` (String)
- `tunterm_ap_subnets` (List of String) List of subnets where we allow AP to establish Mist Tunnels from
- `tunterm_dhcpd_config` (Attributes Map) DHCP server/relay configuration of Mist Tunneled VLANs. Property key is the VLAN ID (see [below for nested schema](#nestedatt--tunterm_dhcpd_config))
- `tunterm_extra_routes` (Attributes Map) Extra routes for Mist Tunneled VLANs. Property key is a CIDR (see [below for nested schema](#nestedatt--tunterm_extra_routes))
- `tunterm_hosts` (List of String) Hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP)
- `tunterm_hosts_order` (List of Number) List of index of tunterm_hosts
- `tunterm_hosts_selection` (String) Ordering of tunterm_hosts for mxedge within the same mxcluster. enum:
  * `shuffle`: the ordering of tunterm_hosts is randomized by the device''s MAC
  * `shuffle-by-site`: shuffle by site_id+tunnel_id (so when client connects to a specific Tunnel, it will go to the same (order of) mxedge, and we load-balancing between tunnels)
  * `ordered`: order decided by tunterm_hosts_order
- `tunterm_monitoring` (List of List of Object)
- `tunterm_monitoring_disabled` (Boolean)

### Read-Only

- `id` (String) Unique ID of the object instance in the Mist Organization
- `radsec_tls` (Attributes) (see [below for nested schema](#nestedatt--radsec_tls))

<a id="nestedatt--mist_das"></a>
### Nested Schema for `mist_das`

Optional:

- `coa_servers` (Attributes List) Dynamic authorization clients configured to send CoA|DM to mist edges on port 3799 (see [below for nested schema](#nestedatt--mist_das--coa_servers))
- `enabled` (Boolean)

<a id="nestedatt--mist_das--coa_servers"></a>
### Nested Schema for `mist_das.coa_servers`

Optional:

- `disable_event_timestamp_check` (Boolean) Whether to disable Event-Timestamp Check
- `enabled` (Boolean)
- `host` (String) This server configured to send CoA|DM to mist edges
- `port` (Number) Mist edges will allow this host on this port
- `require_message_authenticator` (Boolean) Whether to require Message-Authenticator in requests
- `secret` (String, Sensitive)



<a id="nestedatt--mist_nac"></a>
### Nested Schema for `mist_nac`

Optional:

- `acct_server_port` (Number)
- `auth_server_port` (Number)
- `enabled` (Boolean)
- `secret` (String, Sensitive)

Read-Only:

- `client_ips` (Attributes Map) Property key is the RADIUS Client IP/Subnet. (see [below for nested schema](#nestedatt--mist_nac--client_ips))

<a id="nestedatt--mist_nac--client_ips"></a>
### Nested Schema for `mist_nac.client_ips`



<a id="nestedatt--mxedge_mgmt"></a>
### Nested Schema for `mxedge_mgmt`

Optional:

- `config_auto_revert` (Boolean)
- `fips_enabled` (Boolean)
- `mist_password` (String, Sensitive)
- `oob_ip_type` (String) enum: `dhcp`, `disabled`, `static`
- `oob_ip_type6` (String) enum: `autoconf`, `dhcp`, `disabled`, `static`
- `root_password` (String, Sensitive)


<a id="nestedatt--proxy"></a>
### Nested Schema for `proxy`

Optional:

- `disabled` (Boolean)
- `url` (String)


<a id="nestedatt--radsec"></a>
### Nested Schema for `radsec`

Optional:

- `acct_servers` (Attributes List) List of RADIUS accounting servers, optional, order matters where the first one is treated as primary (see [below for nested schema](#nestedatt--radsec--acct_servers))
- `auth_servers` (Attributes List) List of RADIUS authentication servers, order matters where the first one is treated as primary (see [below for nested schema](#nestedatt--radsec--auth_servers))
- `enabled` (Boolean) Whether to enable service on Mist Edge i.e. RADIUS proxy over TLS
- `match_ssid` (Boolean) Whether to match ssid in request message to select from a subset of RADIUS servers
- `nas_ip_source` (String) SSpecify NAS-IP-ADDRESS, NAS-IPv6-ADDRESS to use with auth_servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`
- `proxy_hosts` (List of String) Hostnames or IPs for Mist AP to use as the TLS Server (i.e. they are reachable from AP) in addition to `tunterm_hosts`
- `server_selection` (String) When ordered, Mist Edge will prefer and go back to the first radius server if possible. enum: `ordered`, `unordered`
- `src_ip_source` (String) Specify IP address to connect to auth_servers and acct_servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`

<a id="nestedatt--radsec--acct_servers"></a>
### Nested Schema for `radsec.acct_servers`

Optional:

- `host` (String) IP / hostname of RADIUS server
- `port` (Number) Acct port of RADIUS server
- `secret` (String, Sensitive) Secret of RADIUS server
- `ssids` (List of String) List of ssids that will use this server if match_ssid is true and match is found


<a id="nestedatt--radsec--auth_servers"></a>
### Nested Schema for `radsec.auth_servers`

Optional:

- `host` (String) IP / hostname of RADIUS server
- `inband_status_check` (Boolean) Whether to enable inband status check
- `inband_status_interval` (Number) Inband status interval, in seconds
- `keywrap_enabled` (Boolean) If used for Mist APs, enable keywrap algorithm. Default is false
- `keywrap_format` (String) if used for Mist APs. enum: `ascii`, `hex`
- `keywrap_kek` (String) If used for Mist APs, encryption key
- `keywrap_mack` (String) If used for Mist APs, Message Authentication Code Key
- `port` (Number) Auth port of RADIUS server
- `retry` (Number) Authentication request retry
- `secret` (String, Sensitive) Secret of RADIUS server
- `ssids` (List of String) List of ssids that will use this server if match_ssid is true and match is found
- `timeout` (Number) Authentication request timeout, in seconds



<a id="nestedatt--tunterm_dhcpd_config"></a>
### Nested Schema for `tunterm_dhcpd_config`

Optional:

- `enabled` (Boolean)
- `servers` (List of String)
- `type` (String) enum: `relay`


<a id="nestedatt--tunterm_extra_routes"></a>
### Nested Schema for `tunterm_extra_routes`

Optional:

- `via` (String)


<a id="nestedatt--radsec_tls"></a>
### Nested Schema for `radsec_tls`

Optional:

- `keypair` (String)


