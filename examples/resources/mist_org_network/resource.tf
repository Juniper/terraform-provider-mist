resource "mist_org_network" "network_one" {
  org_id                 = mist_org.terraform_test.id
  name                   = "network_one"
  subnet                 = "10.4.0.0/24"
  disallow_mist_services = false
}