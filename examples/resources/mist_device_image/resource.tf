resource "mist_device_image" "device_image_one" {
  device_id = provider::mist::search_inventory_by_claimcode(resource.mist_org_inventory.inventory, "CPKL2EXXXXXXXXX").id
  site_id   = provider::mist::search_inventory_by_claimcode(resource.mist_org_inventory.inventory, "CPKL2EXXXXXXXXX").site_id
  file    = "/Users/johndoe/Documents/image.jpg"
  image_number = 1
}