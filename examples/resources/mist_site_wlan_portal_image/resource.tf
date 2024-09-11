resource "mist_site_wlan_portal_image" "wlan_one" {
  site_id = mist_site.terraform_test.id
  wlan_id = mist_site_wlan.wlan_one.id
  file    = "/Users/johndoe/Documents/image.jpg"
}
