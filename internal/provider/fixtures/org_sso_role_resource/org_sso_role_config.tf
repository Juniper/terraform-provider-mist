
  name   = "Example SSO Role"
  privileges = [{
    role  = "admin"
    scope = "org"
    views = ["reporting", "marketing", "super_observer", "location", "security", "switch_admin", "mxedge_admin", "lobby_admin"]
  },
  {
    role         = "read"
    scope        = "sitegroup"
  },
  {
    role         = "read"
    scope        = "site"
  },   
  {
    role  = "helpdesk"
    scope = "org"
    views = ["switch_admin", "marketing"]
  },
  {
    role  = "write"
    scope = "orgsites"
    views = ["location", "security"]
  },
  {
    role  = "installer"
    scope = "org"
    view  = "security"
  }]

