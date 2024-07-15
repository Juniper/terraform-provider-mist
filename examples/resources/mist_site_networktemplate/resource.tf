
resource "mist_site_setting" "networktemplate_one" {
  site_id      = mist_site.terraform_test.id
  dns_servers = ["8.8.8.8", "1.1.1.1"]
  dns_suffix  = ["mycorp.com"]
  ntp_servers = ["pool.ntp.org"]
  additional_config_cmds = [
    "set system hostnam test",
    "set system services ssh root-login allow"
  ]
  networks = {
    network_one = {
      vlan_id = 10
    }
    network_two = {
      vlan_id = 11

    }
  }
  port_usages = {
    trunk = {
      all_networks = true
      enable_qos   = true
      mode         = "port_usage_one"
      port_network = "network_one"
    }
  }
  radius_config = {
    acct_interim_interval = 60
    coa_enabled           = true
    network               = "network_one"
    acct_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
    auth_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
  }
  switch_matching = {
    enable = true
    rules = [
      {
        name        = "switch_rule_one"
        match_type  = "match_name[0:3]"
        match_value = "abc"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "port_usage_one"
          }
        }
      }
    ]
  }
}
