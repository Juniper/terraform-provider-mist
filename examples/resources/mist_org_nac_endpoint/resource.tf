resource "mist_org_nac_endpoint" "endpoint_one" {
  mac          = "921b638445cd"
  labels       = ["byod", "flr1"]
  vlan         = "30"
  notes        = "mac address refers to Canon printers"
  name         = "endpoint_one"
  radius_group = "VIP"
}
