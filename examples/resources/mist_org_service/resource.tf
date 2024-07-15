resource "mist_org_service" "service_one" {
  org_id = mist_org.terraform_test.id
  name   = "service_one"
  addresses = [
    "10.3.0.0/24",
    "10.4.0.0/24"
  ]
  type = "custom"
  specs = [
    {
      protocol   = "tcp"
      port_range = "443"
    }
  ]
}
