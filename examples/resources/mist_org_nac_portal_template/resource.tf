
# Create a NAC Portal first
resource "mist_org_nac_portal" "guest_portal" {
  org_id      = mist_org.terraform_test.id
  name        = "Guest Portal"
  type        = "guest_portal"
  access_type = "wireless"
  ssid        = "Guest-Network"
}

# Example 1: Basic NAC Portal Template with centered alignment
resource "mist_org_nac_portal_template" "centered_template" {
  org_id       = mist_org.terraform_test.id
  nacportal_id = mist_org_nac_portal.guest_portal.id
  alignment    = "center"
  color        = "#1074bc"
  powered_by   = false
}

# Example 2: NAC Portal Template with logo and left alignment
resource "mist_org_nac_portal_template" "logo_template" {
  org_id       = mist_org.terraform_test.id
  nacportal_id = mist_org_nac_portal.guest_portal.id
  alignment    = "left"
  color        = "#ff6600"
  logo         = "logo.png"
  powered_by   = true
}
