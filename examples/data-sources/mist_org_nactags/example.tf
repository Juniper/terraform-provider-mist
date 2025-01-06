data "mist_org_nactags" "nactags" {
  org_id  = "15fca2ac-b1a6-47cc-9953-cc6906281550"

  // Filtering options
  type = "match"
  match = "cert_issuer"
}