  org_id = "{org_id}"
  name   = "Test Alarm Template"

  delivery {
    enabled             = true
    to_org_admins      = true
    to_site_admins     = false
    additional_emails  = ["admin@example.com","alerts@example.com"]
  }

  rules = {
    "ap_bad_cable" = {
      enabled = true
      delivery = {
        enabled            = true
        to_org_admins     = true
        to_site_admins    = false
        additional_emails = ["network@example.com"]
      }
    }
    "ap_config_failed" = {
      enabled = true
      delivery = {
        enabled            = true
        to_org_admins     = false
        to_site_admins    = true
        additional_emails = []
      }
    }
    "ap_disconnected" = {
      enabled = false
    }
    "switch_bad_cable" = {
      enabled = true
      delivery = {
        enabled            = true
        to_org_admins     = true
        to_site_admins    = true
        additional_emails = [
          "switch-admin@example.com",
          "infrastructure@example.com"
        ]
      }
    }
    "gateway_down" = {
      enabled = true
      delivery = {
        enabled            = false
        to_org_admins     = false
        to_site_admins    = false
        additional_emails = []
      }
    }
  }

