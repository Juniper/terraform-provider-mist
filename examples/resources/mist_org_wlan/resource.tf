resource "mist_org_wlan" "wlan_one" {
  ssid              = "wlan_one"
  org_id      = mist_org.terraform_test.id
  template_id = mist_org_wlantemplate.test101.id
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
  interface   = "all"
}