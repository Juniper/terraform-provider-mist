resource "mist_org_avprofile" "avprofile_one" {
  org_id          = mist_org.terraform_test.id
  fallback_action = "block"
  max_filesize    = 5000
  mime_whitelist = [
    "image/png"
  ]
  name = "avprofile_one"
  protocols = [
    "ftp",
    "http",
    "imap",
    "pop3",
    "smtp"
  ]
  url_whitelist = ["www.google.fr"]
}
