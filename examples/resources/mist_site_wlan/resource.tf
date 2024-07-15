resource "mist_site_wlan" "wlan_one" {
  ssid              = "wlan_one"
  site_id           = mist_site.terraform_test.id
  bands             = ["5", "6"]
  vlan_id           = 143
  wlan_limit_up     = 10000
  wlan_limit_down   = 20000
  client_limit_up   = 512
  client_limit_down = 1000
  auth = {
    type = "psk"
    psk  = "secretpsk"
  }
  interface = "all"
}
