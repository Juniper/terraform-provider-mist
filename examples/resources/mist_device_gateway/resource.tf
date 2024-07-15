resource "mist_device_gateway" "gateway_one" {
  name        = "gateway_one"
  device_id = mist_device_gateway_cluster.cluster_one.device_id
  site_id   = mist_device_gateway_cluster.cluster_one.site_id
  oob_ip_config = {
    type = "dhcp"
  }
  dns_servers = ["8.8.8.8"]
  additional_config_cmds = [
    "annotate system \" -- custom-main -- Template level --\"",
    "delete apply-groups custom-main",
    "delete groups custom-main",
    "set groups custom-main",
    "set groups custom-main system services ssh root-login allow",
    "set apply-groups custom-main",
  ]
}