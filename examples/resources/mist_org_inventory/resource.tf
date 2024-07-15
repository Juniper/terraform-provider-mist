resource "mist_org_inventory" "inventory_one" {
  org_id = mist_org.terraform_test.id
  devices = [
    {
      claim_code = "<device_claim_code>"
      site_id    = mist_site.terraform_site.id
    },
    {
      claim_code = "<device_claim_code>"
    }
  ]
}