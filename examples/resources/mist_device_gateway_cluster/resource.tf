resource "mist_device_gateway_cluster" "cluster_one" {
  site_id   = mist_site.terraform_site2.id
  device_id = "00000000-0000-0000-1000-4c96143de700"
  nodes = [
    { mac = "4c961000000" },
    { mac = "4c961000001" }
  ]
}