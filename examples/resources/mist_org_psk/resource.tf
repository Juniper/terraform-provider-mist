resource "mist_org_psk" "psk_one" {
  org_id    = mist_org.terraform_test.id
  name       = "JNP-FR-PAR"
  passphrase = "secretone"
  ssid       = mist_org_wlan.wlan_one.ssid
  usage      = "multi"
}