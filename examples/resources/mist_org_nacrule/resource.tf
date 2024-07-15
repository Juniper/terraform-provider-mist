resource "mist_org_nacrule" "nacrule_one" {
  name   = "rule_one"
  action = "allow"
  org_id = mist_org.terraform_test.id
  matching = {
    port_types = [
      "wired"
    ],
    auth_type = "mab",
    nactags = [
      "c055c60b-351a-4311-8ee5-9b7be5e5f902"
    ]
  }
  apply_tags = [
    "61c11327-5e1b-40ed-bbbf-5e95642c4f59",
    "3f292454-ac5f-4a36-9aff-d0518d90b47a"
  ]
  enabled = true
  order   = 9
}
