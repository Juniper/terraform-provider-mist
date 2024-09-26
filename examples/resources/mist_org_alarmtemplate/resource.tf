resource "mist_org_alarmtemplate" "alarmtemplate_one" {
  org_id = mist_org.terraform_test.id
  name   = "alarmtemplate_one"
  delivery = {
    enabled           = true
    to_org_admins     = true
    additional_emails = ["admin@mycorp.net"]
  }
  rules = {
    health_check_failed : {
      enabled : true
    },
    insufficient_capacity : {
      enabled : true
    },
    insufficient_coverage : {
      enabled : true
    },
    infra_arp_failure : {
      enabled : true
    },
    arp_failure : {
      enabled : true
    }
  }
}
