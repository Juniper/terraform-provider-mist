
resource "mist_device_ap" "ap_one" {
  name      = "test_ap"
  device_id = mist_org_inventory.inventory.devices[0].id
  site_id   = mist_org_inventory.inventory.devices[0].site_id
}