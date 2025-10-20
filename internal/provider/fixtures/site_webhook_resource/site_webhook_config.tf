
name = "test-comprehensive-webhook"
url = "https://webhook.example.com/endpoint"
type = "oauth2"
enabled = true
topics = [
  "alarms",
  "device-events"
]
headers = {
  "Authorization" = "Bearer comprehensive-token"
  "X-Webhook-Source" = "mist-cloud"
  "X-Event-Version" = "v2"
  "Content-Type" = "application/json"
  "X-API-Key" = "api-key-12345"
}
oauth2_grant_type = "client_credentials"
oauth2_client_id = "comprehensive-client-id"
oauth2_client_secret = "comprehensive-client-secret"
oauth2_token_url = "https://oauth2.example.com/token"
oauth2_scopes = [
  "webhook:write",
  "events:read",
  "webhooks:manage",
  "admin:read"
]
single_event_per_message = true
verify_cert = false
