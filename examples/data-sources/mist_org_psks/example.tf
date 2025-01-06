data "mist_org_psks" "psks_vip" {
  org_id = "15fca2ac-b1a6-47cc-9953-cc6906281550"

  // Filtering options
  name = "psk_one"
  role   = "vip"
  ssid = "psk_ssid"
}
