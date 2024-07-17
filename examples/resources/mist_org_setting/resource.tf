resource "mist_org_vpn" "vpn_one" {
  org_id = mist_org.terraform_test.id
  name   = "vpn_one"
  paths = {
    "AWS_Hub_Profile1-WAN1" : {
      bfd_profile = "broadband"
    },
    "AWS_Hub_Profile1-WAN2" : {},
  }
}