
resource "mist_org_nac_portal" "guest_portal" {
  org_id      = mist_org.terraform_test.id
  name        = "Guest Portal"
  type        = "guest_portal"
  access_type = "wireless+wired"
  ssid        = "Guest-Network"
  
  # Certificate and authentication settings
  cert_expire_time          = 365
  eap_type                  = "wpa3"
  enable_telemetry          = true
  expiry_notification_time  = 30
  notify_expiry             = true
  tos                       = "By using this network, you agree to our terms of service and privacy policy."

  # Additional CA certificates for authentication
  additional_cacerts = [
    "-----BEGIN CERTIFICATE-----\nMIIC...certificate...content\n-----END CERTIFICATE-----"
  ]

  # Additional NAC server names
  additional_nac_server_name = [
    "nac1.example.com",
    "nac2.example.com"
  ]

  # Portal configuration
  portal = {
    auth                = "multi"
    expire              = 43200
    external_portal_url = "https://portal.example.com/external"
    force_reconnect     = false
    forward             = true
    forward_url         = "https://example.com/portal/welcome"
    max_num_devices     = 10
    privacy             = true
  }

  # SSO configuration
  sso = {
    idp_cert              = "-----BEGIN CERTIFICATE-----\nMIIC...idp...cert\n-----END CERTIFICATE-----"
    idp_sign_algo         = "sha384"
    idp_sso_url           = "https://idp.example.com/saml/sso"
    issuer                = "https://idp.example.com"
    nameid_format         = "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent"
    use_sso_role_for_cert = false
    
    sso_role_matching = [
      {
        match    = "Administrator"
        assigned = "full-access"
      },
      {
        match    = "Manager"
        assigned = "manager-access"
      },
      {
        match    = "Employee"
        assigned = "employee-access"
      },
      {
        match    = "Guest"
        assigned = "guest-access"
      }
    ]
  }
}
