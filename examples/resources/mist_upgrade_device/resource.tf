resource "mist_upgrade_device" "ap_upgrade" {
  site_id = mist_site.terraform_test.id
  device_id = mist_device_switch.ap_one.id
  target_version = "0.14.29543"
}

resource "mist_upgrade_device" "switch_upgrade" {
  site_id = mist_site.terraform_test.id
  device_id = mist_device_switch.switch_one.id
  target_version = "24.2R1-S1.10"
  reboot = true
  sync_upgrade_timeout = 3600
}