
  name                    = "Example SAML SSO"
  idp_cert                = "-----BEGIN CERTIFICATE-----\nMIIBkTCB+wIJALJ8UUKmgH1GMA0GCSqGSIb3DQEBCwUAMBQxEjAQBgNVBAMMCXRl\nc3QtY2VydDAeFw0yNDEyMjcwMDAwMDBaFw0yNTEyMjcwMDAwMDBaMBQxEjAQBgNV\nBAMMCXRlc3QtY2VydDBcMA0GCSqGSIb3DQEBAQUAA0sAMEgCQQCxUC6+OeSgM1Fh\nOdKqA5C1XQfFdKK0C8JxUQKHjOKE8Q1j8I+FHFOdKGY5TKZrIvOLMbOeXJGF7Wl5\nxD0dVhZdAgMBAAEwDQYJKoZIhvcNAQELBQADQQA3F8+8MzE5E5GHj5E5TQ==\n-----END CERTIFICATE-----"
  idp_sign_algo           = "sha256"
  idp_sso_url             = "https://idp.example.com/sso/saml"
  issuer                  = "https://idp.example.com/issuer"
  custom_logout_url       = "https://idp.example.com/logout"
  default_role            = "viewer"
  ignore_unmatched_roles  = false
  nameid_format           = "email"
  role_attr_extraction    = "cn"
  role_attr_from          = "Role"

