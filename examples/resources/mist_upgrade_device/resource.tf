resource "mist_upgrade_device" "upgrade_one" {
  site_id = mist_site.terraform_test.id
  device_id = mist_device_switch.switch_one.id
  upgrade_to_version = "24.2R1-S1.10"
  reboot = true
  sync_upgrade_timeout = 3600
}