# result is a device object
resource "mist_device_ap" "test_ap" {
  device_id = provider::mist::search_inventory_by_claimcode(resource.mist_org_inventory.inventory.devices, "CPKL2EXXXXXXXXX").id
  site_id   = provider::mist::search_inventory_by_claimcode(resource.mist_org_inventory.inventory.devices, "CPKL2EXXXXXXXXX").site_id
  name      = "test_ap"
}
