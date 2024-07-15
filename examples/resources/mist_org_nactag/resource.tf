resource "mist_org_nactag" "tag_one" {
  name   = "tag_one"
  type   = "match"
  match  = "client_mac"
  org_id = mist_org.terraform_test.id
  values = [
    "5c5b35*"
  ]
}
