
name = "test-nac-portal-all-attributes"
type = "guest_portal"
access_type = "wireless+wired"
ssid = "Test-All-Attributes"
cert_expire_time = 365
eap_type = "wpa3"
enable_telemetry = true
expiry_notification_time = 30
notify_expiry = true
tos = "By using this network, you agree to our terms of service and privacy policy."
additional_cacerts = [
  "-----BEGIN CERTIFICATE-----\nMIIC1...ca1...cert\n-----END CERTIFICATE-----",
  "-----BEGIN CERTIFICATE-----\nMIIC2...ca2...cert\n-----END CERTIFICATE-----",
  "-----BEGIN CERTIFICATE-----\nMIIC3...ca3...cert\n-----END CERTIFICATE-----"
]
additional_nac_server_name = [
  "nac1.example.com",
  "nac2.example.com",
  "nac3.example.com"
]
portal = {
  auth = "multi"
  expire = 43200
  external_portal_url = "https://portal.example.com/external"
  force_reconnect = false
  forward = true
  forward_url = "https://example.com/portal/welcome"
  max_num_devices = 10
  privacy = true
}
sso = {
  idp_cert = "-----BEGIN CERTIFICATE-----\nMIIC...comprehensive...cert\n-----END CERTIFICATE-----"
  idp_sign_algo = "sha384"
  idp_sso_url = "https://idp.example.com/saml/sso"
  issuer = "https://idp.example.com"
  nameid_format = "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent"
  use_sso_role_for_cert = false
  sso_role_matching = [
    {
      match = "Administrator"
      assigned = "full-access"
    },
    {
      match = "Manager"
      assigned = "manager-access"
    },
    {
      match = "Employee"
      assigned = "employee-access"
    },
    {
      match = "Contractor"
      assigned = "contractor-access"
    },
    {
      match = "Guest"
      assigned = "guest-access"
    }
  ]
}
