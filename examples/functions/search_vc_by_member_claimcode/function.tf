# result is a device object
resource "mist_device_switch" "virtual_chassis_one" {
  device_id = provider::mist::search_vc_by_member_claimcode(resource.mist_org_inventory.inventory, "CPKL2EXXXXXXXXX").id
  site_id   = provider::mist::search_vc_by_member_claimcode(resource.mist_org_inventory.inventory, "CPKL2EXXXXXXXXX").site_id
  name      = "virtual_chassis_one"
}
