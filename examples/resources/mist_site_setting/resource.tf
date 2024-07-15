
resource "mist_site_setting" "site_one" {
  site_id = mist_site.terraform_site.id
  ap_updown_threshold     = 5
  device_updown_threshold = 5
  auto_upgrade = {
    enabled     = true
    day_of_week = "tue"
    time_of_day = "02:00"
    version     = "beta"
  }
  config_auto_revert = true
  persist_config_on_device = true
  proxy = {
    url = "http://myproxy:3128"
  }
  rogue = {
    enabled          = true
    honeypot_enabled = true
    min_duration     = 5
  }
}