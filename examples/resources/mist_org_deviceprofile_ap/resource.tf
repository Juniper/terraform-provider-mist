
resource "mist_org_deviceprofile_ap" "deviceprofile_ap_one" {
  name   = "deviceprofile_ap_one"
  org_id = mist_org.terraform_test.id
  esl_config = {
    enabled = true
    host    = "1.2.3.4"
    type    = "native"
  }
}
