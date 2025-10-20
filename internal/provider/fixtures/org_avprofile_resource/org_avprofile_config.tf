  org_id = "901c5705-ca11-4bf1-9158-31f7195618ef"
  name = "comprehensive_av_profile"
  protocols = ["ftp", "http", "imap", "pop3", "smtp"]
  fallback_action = "block"
  max_filesize = 25000
  mime_whitelist = ["application/pdf", "image/jpeg", "image/png", "text/plain", "application/zip"]
  url_whitelist = ["example.com", "trusted-site.org", "safe-domain.net"]