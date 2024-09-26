resource "mist_device_image" "device_image_one" {
  device_id = mist_org_inventory.inventory.devices[1].id
  site_id   = mist_org_inventory.inventory.devices[1].site_id
  file    = "/Users/johndoe/Documents/image.jpg"
  image_number = 1
}