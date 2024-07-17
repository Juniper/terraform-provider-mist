resource "mist_org_servicepolicy" "servicepolicy_one" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "guest"
  ]
  services = [
    "guest-internet"
  ]
  action    = "allow"
  idp = {
    enabled    = true
    profile    = "standard"
    alert_only = true
  }
  name = "Guest-IDP"
}