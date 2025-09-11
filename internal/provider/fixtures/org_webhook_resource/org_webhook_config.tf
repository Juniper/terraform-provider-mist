  name   = "Basic HTTP Post Webhook"
  url    = "https://example.com/webhook"
  type   = "http-post"
  topics = ["alarms", "device-events"]
␞
  name       = "HTTP Post with Secret"
  url        = "https://api.example.com/webhooks/mist"
  type       = "http-post"
  secret     = "my-webhook-secret-key"
  topics     = ["client-join", "client-sessions"]
  verify_cert = true
  enabled    = true
␞
  name   = "HTTP Post with Single Event"
  url    = "https://webhook.service.com/endpoint"
  type   = "http-post"
  topics = ["audits", "device-updowns"]
  single_event_per_message = true
␞
  name                = "OAuth2 Client Credentials"
  url                 = "https://oauth.example.com/webhook"
  type                = "oauth2"
  oauth2_grant_type   = "client_credentials"
  oauth2_client_id    = "my-client-id"
  oauth2_client_secret = "my-client-secret"
  oauth2_token_url    = "https://oauth.example.com/token"
  oauth2_scopes       = ["webhook.write", "events.read"]
  topics              = ["guest-authorizations", "nac-events"]
␞
  name              = "OAuth2 Password Grant"
  url               = "https://secure.webhook.com/api"
  type              = "oauth2"
  oauth2_grant_type = "password"
  oauth2_username   = "webhook-user"
  oauth2_password   = "secure-password"
  oauth2_token_url  = "https://secure.webhook.com/oauth/token"
  topics            = ["mxedge-events", "nac-accounting"]
  enabled           = true
␞
  name        = "Splunk Webhook"
  url         = "https://splunk.company.com:8088/services/collector"
  type        = "splunk"
  splunk_token = "B5A79AAD-D822-46CC-80D1-819F80D7BFB0"
  topics      = ["alarms", "audits", "device-events"]
  verify_cert = false
␞
  name   = "All Topics Webhook"
  url    = "https://comprehensive.webhook.com/all-events"
  type   = "http-post"
  topics = [
    "alarms",
    "audits",
    "client-info",
    "client-join",
    "client-sessions",
    "device-events",
    "device-updowns",
    "guest-authorizations",
    "mxedge-events",
    "nac-accounting",
    "nac-events"
  ]
  single_event_per_message = false
  verify_cert = true
  enabled = true
␞
  name    = "Disabled Webhook"
  url     = "https://disabled.webhook.com/endpoint"
  topics  = ["alarms"]
  enabled = false