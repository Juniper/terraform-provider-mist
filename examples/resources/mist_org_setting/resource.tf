resource "mist_org_setting" "terraform_test" {
  org_id              = mist_org.terraform_test.id
  ap_updown_threshold = 10
  cradlepoint = {
    cp_api_id   = "cp_api_id_test"
    cp_api_key  = "secret"
    ecm_api_id  = "ecm_api_id_test"
    ecm_api_key = "secret"
  }
  device_updown_threshold  = 10
  disable_pcap             = false
  disable_remote_shell     = true
  gateway_updown_threshold = 10
  mxedge_mgmt = {
    mist_password = "root_secret_password"
    root_password = "root_secret_password"
    oob_ip_type   = "dhcp"
    oob_ip_type6  = "disabled"
  }
  password_policy = {
    enabled                  = true
    freshness                = 180
    min_length               = 12
    requires_special_char    = true
    requires_two_factor_auth = false
  }
  security = {
    disable_local_ssh = true
  }
  switch_updown_threshold = 10
  synthetic_test = {
    disabled = false
    vlans = [{
      vlan_ids         = ["10", "30"]
      custom_test_urls = ["http://www.abc.com/", "https://10.3.5.1:8080/about"]
      }, {
      vlan_ids = ["20"]
      disabled = true
    }]
  }
  ui_idle_timeout = 120
}
