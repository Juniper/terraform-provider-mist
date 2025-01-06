data "mist_device_switch_stats" "switch_stats" {
  org_id = "15fca2ac-b1a6-47cc-9953-cc6906281550"

  // Filtering options
  mac         = "485a0d000000"
  site_id     = "4a422ae5-7ca0-4599-87a3-8e49aa63685f"
  status      = "connected"
  evpn_unused = true
  evpntopo_id = "92984e2f-94db-4cd8-9763-9cf83fbd079e"

  // Stats time range - option #1
  // cannot be used with the `start`/`end` attribute
  // when using the `duration` attribute, `end`==`now` 
  duration = "1d"

  // Stats time range - option #2
  // cannot be used with the `duration` attribute
  start = 1736031600
  end   = 1736175934
}
