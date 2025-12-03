  org_id = "{org_id}"
  name = "comprehensive_av_profile"
  protocols = ["ftp", "http", "imap", "pop3", "smtp"]
  fallback_action = "block"
  max_filesize = 25000
  mime_whitelist = ["application/pdf", "image/jpeg", "image/png", "text/plain", "application/zip"]
  url_whitelist = ["example.com", "trusted-site.org", "safe-domain.net"]