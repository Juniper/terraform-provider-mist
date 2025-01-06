data "mist_org_nac_endpoints" "nac_endpoints" {
  org_id = "15fca2ac-b1a6-47cc-9953-cc6906281550"

  // Filtering options
  mac = "4a422a000000"
  labels = [
    "label_one"
  ]
}
