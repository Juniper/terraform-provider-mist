resource "mist_site_webhook" "webhook_one" {
  site_id     = mist_site.terraform_site.id
  name        = "webhook_one"
  type        = "http-post"
  url         = "https://myserver.com:4321/"
  verify_cert = false
  enabled     = true
  topics = [
    "device-events",
    "alarms",
    "audits",
    "client-join",
    "client-info",
    "client-sessions",
    "device-updowns",
    "mxedge-events",
    "nac-events",
    "nac-accounting"
  ]
}
