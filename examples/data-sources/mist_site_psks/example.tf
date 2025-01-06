data "mist_site_psks" "psks_vip" {
  site_id = "15fca2ac-b1a6-47cc-9953-cc6906281550"
  
  // Filtering options
  name = "psk_one"
  role   = "vip"
  ssid = "psk_ssid"
}
