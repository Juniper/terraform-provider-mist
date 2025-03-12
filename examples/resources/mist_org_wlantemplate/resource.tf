
resource "mist_org_wlantemplate" "wlantemplate_one" {
  name   = "wlantemplate_one"
  org_id = mist_org.terraform_test.id
  applies = {
    site_ids = [
      mist_site.terraform_site.id
    ]
  }
}