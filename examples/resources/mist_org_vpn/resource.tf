
resource "mist_org_setting" "terraform_test" {
  org_id = mist_org.terraform_test.id
  password_policy = {
    enabled                  = true
    min_length               = 8
    requires_special_char    = true
    requires_two_factor_auth = true
  }
  mist_nac = {
    eu_only = true
  }
  synthetic_test = {
    disabled = false
    vlans = [
      {
        vlan_ids = [
          "8",
          "999"
        ],
        disabled = true
      }
    ]
  }
  api_policy = {
    no_reveal = false
  }
}