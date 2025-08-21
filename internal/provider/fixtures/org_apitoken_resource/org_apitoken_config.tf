  org_id = "901c5705-ca11-4bf1-9158-31f7195618ef"
  name   = "Test API Token"
  privileges {
    role  = "admin"
    scope = "org"
  }
  src_ips = [
    "192.168.1.100",
    "10.0.0.0/24",
    "203.0.113.1"
  ]
␞
  org_id = "901c5705-ca11-4bf1-9158-31f7195618ef"
  name   = "Site Admin Token"
  privileges {
    role    = "admin"
    scope   = "site"
    site_id = "2c107c8e-2e06-404a-ba61-e25b5757ecea"
  }
  src_ips = [
    "10.1.1.50",
    "172.16.0.0/16"
  ]
␞
  org_id = "901c5705-ca11-4bf1-9158-31f7195618ef"
  name   = "Multi-Privilege Token"
  privileges {
    role  = "read"
    scope = "org"
  }
  privileges {
    role    = "write"
    scope   = "site"
    site_id = "2c107c8e-2e06-404a-ba61-e25b5757ecea"
  }
  privileges {
    role         = "helpdesk"
    scope        = "sitegroup"
    sitegroup_id = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
  }
  src_ips = [
    "0.0.0.0/0"
  ]
␞
  org_id = "901c5705-ca11-4bf1-9158-31f7195618ef"
  name   = "Installer Token"

  privileges {
    role    = "installer"
    scope   = "site"
    site_id = "2c107c8e-2e06-404a-ba61-e25b5757ecea"
  }

  src_ips = [
    "192.168.100.10",
    "192.168.100.11"
  ]
␞
  org_id = "901c5705-ca11-4bf1-9158-31f7195618ef"
  name   = "Read-Only Token"

  privileges {
    role  = "read"
    scope = "org"
  }

