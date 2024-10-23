# result is a device object
resource "mist_device_ap" "test_ap" {
  device_id = provider::mist::search_inventory_by_serial(resource.mist_org_inventory.inventory, "CPKL2EXXXXXXXXX").id
  site_id   = provider::mist::search_inventory_by_serial(resource.mist_org_inventory.inventory, "CPKL2EXXXXXXXXX").site_id
  name      = "test_ap"
}
