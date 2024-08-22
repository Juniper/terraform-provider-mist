resource "mist_device_gateway_cluster" "cluster_one" {
  site_id   = mist_site.terraform_site2.id
  nodes = [
    { mac = "4c961000000" },
    { mac = "4c961000001" }
  ]
}