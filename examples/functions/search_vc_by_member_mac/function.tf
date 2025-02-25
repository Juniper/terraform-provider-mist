# result is a device object
resource "mist_device_switch" "virtual_chassis_one" {
  device_id = provider::mist::search_vc_by_member_mac(resource.mist_org_inventory.inventory, "c0ffee000000").id
  site_id   = provider::mist::search_vc_by_member_mac(resource.mist_org_inventory.inventory, "c0ffee000000").site_id
  name      = "virtual_chassis_one"
}
