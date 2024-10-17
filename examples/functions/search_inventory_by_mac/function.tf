# result is a device object
resource "mist_device_ap" "test_ap" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory.devices, "c0ffee000000").id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory.devices, "c0ffee000000").site_id
  name      = "test_ap"
}
