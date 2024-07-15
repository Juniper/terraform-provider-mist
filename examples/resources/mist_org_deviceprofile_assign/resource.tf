
resource "mist_org_deviceprofile_assign" "deviceprofile_assign" {
  org_id = mist_org.terraform_test.id
  deviceprofile_id = mist_org_deviceprofile_gateway.hub_one.id
  macs = [
    "4c9614000000",
    "4c9614000001"
  ]
}