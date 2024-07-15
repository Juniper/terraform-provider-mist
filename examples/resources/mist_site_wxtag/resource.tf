resource "mist_site_wxtag" "wtag_one" {
  site_id = mist_site.terraform_test.id
  name   = "wtag_one"
  values = [
    "10.3.0.0/16"
  ]
  op    = "in"
  type  = "match"
  match = "ip_range_subnet"
}
