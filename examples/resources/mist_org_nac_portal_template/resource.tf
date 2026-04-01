

# Example 1: Basic NAC Portal Template with centered alignment
resource "mist_org_nac_portal_template" "centered_template" {
  org_id     = mist_org.terraform_test.id
  alignment  = "center"
  color      = "#1074bc"
  powered_by = false
}

# Example 2: NAC Portal Template with logo and left alignment
resource "mist_org_nac_portal_template" "logo_template" {
  org_id     = mist_org.terraform_test.id
  alignment  = "left"
  color      = "#ff6600"
  logo       = "logo.png"
  powered_by = true
}
