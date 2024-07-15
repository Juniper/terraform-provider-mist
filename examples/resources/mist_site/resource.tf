resource "mist_site" "juniper_france" {
  org_id       = mist_org.terraform_test.id
  name         = "JNP-FR-PAR"
  country_code = "FR"
  timezone     = "Europe/Paris"
  address      = "41 Rue de Villiers, 92100 Neuilly sur Seine, France"
  notes        = "Created with Terraform, Updated with Terraform"
  latlng = {
    lat = 48.899268
    lng = 2.214447
  }
  sitegroup_ids      = [mist_org_sitegroup.test_group.id, mist_org_sitegroup.test_group2.id]
  networktemplate_id = mist_org_networktemplate.switch_template.id
  rftemplate_id      = mist_org_rftemplate.test_rf.id
  gatewaytemplate_id = mist_org_gatewaytemplate.test-api.id
}