
resource "mist_device_ap" "ap_one" {
  name      = "test_ap"
  device_id = provider::mist::search_inventory_by_claimcode(resource.mist_org_inventory.inventory, "CPKL2EXXXXXXXXX").id
  site_id   = provider::mist::search_inventory_by_claimcode(resource.mist_org_inventory.inventory, "CPKL2EXXXXXXXXX").site_id
}