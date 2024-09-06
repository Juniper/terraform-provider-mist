resource "mist_org_wlan_portal_template" "wlan_one" {
  org_id  = mist_org.terraform_test.id
  wlan_id = mist_org.wlan_one.id
  portal_template = {
    sms_message_format    = "Code {{code}} expires in {{duration}} minutes."
    sms_validity_duration = "10"
    page_title            = "Welcome To My Demo Portal"
    locales = {
      "fr-FR" = {
        page_title = "Bienvenue sur mon portail de démo"
      }
    }
  }
}
