resource "mist_org_rftemplate" "rftempalte_one" {
  name         = "rftempalte_one"
  org_id       = mist_org.terraform_test.id
  band_24_usage = "auto"
  band_5 = {
    ant_gain = 2
    power    = 8
    channels = [
      60,
      104,
      132
    ]
    bandwidth = 20
  }
  band_6 = {
    ant_gain = 2
    power    = 8
  }
  band_24 = {
    ant_gain          = 1
    allow_rrm_disable = true
    power_min         = 18
    power_max         = 18
    bandwidth         = 20
  }
  ant_gain_5   = 2
  ant_gain_6   = 2
  ant_gain_24  = 1
  country_code = "FR"
}
