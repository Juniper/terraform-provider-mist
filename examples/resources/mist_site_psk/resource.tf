resource "mist_site_psk" "psk_one" {
  site_id    = mist_site.terraform_site.id
  name       = "JNP-FR-PAR"
  passphrase = "secretone"
  ssid       = mist_site_wlan.wlan_one.ssid
  usage      = "multi"
}
