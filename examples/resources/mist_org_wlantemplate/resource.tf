
resource "mist_org_wlantemplate" "wlantempalte_one" {
  name   = "wlantempalte_one"
  org_id = mist_org.terraform_test.id
  applies = {
    site_ids = [
      mist_site.terraform_site.id
    ]
  }
}