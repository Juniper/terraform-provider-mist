
resource "mist_org_deviceprofile_switch" "deviceprofile_switch_one" {
  name   = "deviceprofile_switch_one"
  org_id = mist_org.terraform_test.id
}
