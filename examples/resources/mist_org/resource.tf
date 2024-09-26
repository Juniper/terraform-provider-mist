resource "mist_org" "org_one" {
  name             = "Org One"
  alarmtemplate_id = mist_org_alarmtemplate.alarmtemplate_one.id
}
