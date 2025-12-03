  org_id = "{org_id}"
  name   = "Test API Token"
  privileges = [{
      role  = "admin"
      scope = "org"
    }]
  src_ips = [
    "192.168.1.100",
    "10.0.0.0/24",
    "203.0.113.1"
  ]
␞
  org_id = "{org_id}"
  name   = "Site Admin Token"
  privileges = [
    {
      role    = "admin"
      scope   = "site"
      site_id = "mist_site.TestSite.id"
    }
  ]
  src_ips = [
    "10.1.1.50",
    "172.16.0.0/16"
  ]
␞
  org_id = "{org_id}"
  name   = "Multi-Privilege Token"
  privileges = [
    {
      role  = "read"
      scope = "org"
    },
    {
      role    = "write"
      scope   = "site"
      site_id = "mist_site.TestSite.id"
    }
  ]
  src_ips = [
    "0.0.0.0/0"
  ]
␞
  org_id = "{org_id}"
  name   = "Installer Token"

  privileges = [
    {
      role    = "installer"
      scope   = "site"
      site_id = "mist_site.TestSite.id"
    }
  ]

  src_ips = [
    "192.168.100.10",
    "192.168.100.11"
  ]
␞
  org_id = "{org_id}"
  name   = "Read-Only Token"

  privileges = [{
    role  = "read"
    scope = "org"
  }]
␞
  org_id = "{org_id}"
  name   = "Org Sites Token"

  privileges = [{
    role  = "admin"
    scope = "orgsites"
  }]

  src_ips = ["192.168.0.0/16"]

