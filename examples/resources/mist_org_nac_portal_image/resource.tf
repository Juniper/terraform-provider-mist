
# First, create a NAC portal
resource "mist_org_nac_portal" "portal" {
  org_id      = mist_org.terraform_test.id
  name        = "Portal with Background"
  type        = "guest_portal"
  access_type = "wireless"
  ssid        = "Guest-WiFi"
}

# Upload a background image for the NAC portal
resource "mist_org_nac_portal_image" "background" {
  org_id       = mist_org.terraform_test.id
  nacportal_id = mist_org_nac_portal.portal.id
  file         = "${path.module}/images/background.jpg"
}
