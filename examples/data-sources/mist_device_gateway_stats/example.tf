data "mist_device_gateway_stats" "gateway_stats" {
  org_id  = "15fca2ac-b1a6-47cc-9953-cc6906281550"

  // Filtering options
  mac     = "e8a245000000"
  site_id = "4a422ae5-7ca0-4599-87a3-8e49aa63685f"
  status  = "connected"

  // Stats time range - option #1
  // cannot be used with the `start`/`end` attribute
  // when using the `duration` attribute, `end`==`now` 
  duration = "1d"

  // Stats time range - option #2
  // cannot be used with the `duration` attribute
  start = 1736031600
  end   = 1736175934
}