resource "mist_org_wlan_portal_image" "wlan_one" {
  site_id = mist_org.terraform_test.id
  wlan_id = mist_org.wlan_one.id
  file    = "/Users/johndoe/Documents/image.jpg"
}
