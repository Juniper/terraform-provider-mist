resource "mist_site_wxrule" "wxrule_one" {
  site_id     = mist_site.terraform_test.id
  src_wxtags = [
    mist_org_wxtag.wxtag_one.id
  ]
  enabled = true
  action  = "allow"
  dst_deny_wxtags = [
    mist_org_wxtag.wxtag_two.id
  ]
  order = 1
}
