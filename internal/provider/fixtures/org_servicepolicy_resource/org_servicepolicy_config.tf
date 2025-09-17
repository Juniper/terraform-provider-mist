name   = "test-servicepolicy-fixture"

action         = "allow"
local_routing  = true
path_preference = "hub"

services = ["ssl-proxy", "idp"]
tenants  = ["default", "guest"]

aamw = {
  enabled         = true
  profile         = "standard"
  aamwprofile_id  = "12345678-1234-1234-1234-123456789012"
}

antivirus = {
  enabled       = true
  profile       = "default"
  avprofile_id  = "12345678-1234-1234-1234-123456789013"
}

appqoe = {
  enabled = true
}

ewf = [
  {
    enabled       = true
    alert_only    = true
    profile       = "strict"
    block_message = "Content blocked by security policy"
  }
]

idp = {
  enabled        = true
  alert_only     = true
  profile        = "strict"
  idpprofile_id  = "12345678-1234-1234-1234-123456789014"
}

ssl_proxy = {
  enabled          = true
  ciphers_category = "medium"
}