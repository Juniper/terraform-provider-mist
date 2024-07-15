resource "mist_org_wxrule" "wxrule_one" {
  org_id      = mist_org.terraform_test.id
  template_id = mist_org_wlantemplate.wlantempalte_one.id
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
