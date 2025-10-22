name                = "comprehensive-rftemplate"
ant_gain_24         = 3
ant_gain_5          = 5
ant_gain_6          = 7
band_24_usage       = "24"
country_code        = "US"
scanning_enabled    = true

band_24 {
  allow_rrm_disable = true
  ant_gain          = 2
  antenna_mode      = "1x1"
  bandwidth         = 20
  channels          = [1, 6, 11]
  disabled          = true
  power             = 15
  power_max         = 18
  power_min         = 5
  preamble          = "short"
}

band_5 {
  allow_rrm_disable = true
  ant_gain          = 4
  antenna_mode      = "2x2"
  bandwidth         = 80
  channels          = [36, 40, 44, 48]
  disabled          = true
  power             = 15
  power_max         = 17
  power_min         = 8
  preamble          = "long"
}

band_5_on_24_radio {
  allow_rrm_disable = true
  ant_gain          = 3
  antenna_mode      = "1x1"
  bandwidth         = 40
  channels          = [1, 6]
  disabled          = true
  power             = 12
  power_max         = 17
  power_min         = 6
  preamble          = "short"
}

band_6 {
  allow_rrm_disable = true
  ant_gain          = 6
  antenna_mode      = "4x4"
  bandwidth         = 160
  channels          = [1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61]
  disabled          = true
  power             = 18
  power_max         = 18
  power_min         = 10
  preamble          = "auto"
  standard_power    = true
}

model_specific = {
  "AP12" = {
    ant_gain_24     = 2
    ant_gain_5      = 4
    ant_gain_6      = 6
    band_24_usage   = "auto"

    band_24 {
      allow_rrm_disable = true
      ant_gain          = 1
      antenna_mode      = "1x1"
      bandwidth         = 20
      channels          = [1, 6, 11]
      disabled          = true
      power             = 14
      power_max         = 18
      power_min         = 4
      preamble          = "short"
    }

    band_5 {
      allow_rrm_disable = true
      ant_gain          = 3
      antenna_mode      = "2x2"
      bandwidth         = 80
      channels          = [36, 40, 44, 48, 52, 56, 60, 64]
      disabled          = true
      power             = 15
      power_max         = 17
      power_min         = 7
      preamble          = "long"
    }

    band_5_on_24_radio {
      allow_rrm_disable = true
      ant_gain          = 2
      antenna_mode      = "1x1"
      bandwidth         = 40
      channels          = [1, 6]
      disabled          = true
      power             = 11
      power_max         = 16
      power_min         = 5
      preamble          = "auto"
    }

    band_6 {
      allow_rrm_disable = true
      ant_gain          = 5
      antenna_mode      = "4x4"
      bandwidth         = 160
      channels          = [1, 5, 9, 13, 17, 21, 25, 29]
      disabled          = true
      power             = 16
      power_max         = 18
      power_min         = 8
      preamble          = "short"
      standard_power    = true
    }
  }

  "AP21" = {
    ant_gain_24     = 3
    ant_gain_5      = 5
    ant_gain_6      = 7
    band_24_usage   = "24"

    band_24 {
      allow_rrm_disable = true
      ant_gain          = 2
      antenna_mode      = "2x2"
      bandwidth         = 40
      channels          = [1, 6, 11, 14]
      disabled          = true
      power             = 16
      power_max         = 18
      power_min         = 6
      preamble          = "long"
    }

    band_5 {
      allow_rrm_disable = true
      ant_gain          = 4
      antenna_mode      = "4x4"
      bandwidth         = 80
      channels          = [36, 40, 44, 48, 149, 153, 157, 161]
      disabled          = true
      power             = 15
      power_max         = 17
      power_min         = 9
      preamble          = "auto"
    }

    band_5_on_24_radio {
      allow_rrm_disable = true
      ant_gain          = 3
      antenna_mode      = "2x2"
      bandwidth         = 80
      channels          = [1, 6, 11]
      disabled          = true
      power             = 13
      power_max         = 17
      power_min         = 7
      preamble          = "short"
    }

    band_6 {
      allow_rrm_disable = true
      ant_gain          = 6
      antenna_mode      = "4x4"
      bandwidth         = 160
      channels          = [1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45]
      disabled          = true
      power             = 16
      power_max         = 18
      power_min         = 12
      preamble          = "long"
      standard_power    = true
    }
  }
}