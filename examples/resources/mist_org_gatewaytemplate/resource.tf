
resource "mist_org_gatewaytemplate" "gwtemplate_one" {
  type   = "spoke"
  name   = "gwtemplate_one"
  org_id = mist_org.terraform_test.id
  port_config = {
    "ge-0/0/3" = {
      name       = "FTTH"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type    = "static"
        ip      = "192.168.1.8"
        netmask = "/24"
        gateway = "192.168.1.1"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/5" = {
      usage            = "lan"
      critical         = false
      aggregated       = true
      ae_disable_lacp  = false
      ae_lacp_force_up = true
      ae_idx           = 0
      redundant        = false
      networks = [
        "PRD-Core",
        "PRD-Mgmt",
        "PRD-Lab"
      ]
    }
  }
  ip_configs = {
    "PRD-Core" = {
      type    = "static"
      ip      = "10.3.100.9"
      netmask = "/24"
    },
    "PRD-Mgmt" = {
      type    = "static"
      ip      = "10.3.172.1"
      netmask = "/24"
    },
    "PRD-Lab" = {
      type    = "static"
      ip      = "10.3.171.1"
      netmask = "/24"
    }
  }
  service_policies = [
    {
      name = "Policy-14"
      tenants = [
        "PRD-Core"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled    = true
        profile    = "critical"
        alert_only = false
      }
    }
  ]
}