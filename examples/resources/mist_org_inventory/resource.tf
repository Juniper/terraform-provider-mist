resource "mist_org_inventory" "inventory" {
  org_id = mist_org.terraform_test.id
  inventory = {
    # Device Claim Code (16 characters)
    # Used to claim and manage "Cloud Ready" devices
    "CPKL2EXXXXXXXXX" = {}
    "G87JHBFXXXXXXXX" = {
      site_id = mist_site.terraform_site.id
      unclaim_when_destroyed = true
    }
    # MAC Address (12 characters)
    # Used to manage devices manually claimed or adopted devices
    "2c2131000000" = {
      site_id                = mist_site.terraform_site.id
      unclaim_when_destroyed = true
    }
    "2c2131000001" = {
      unclaim_when_destroyed = false
    }    
  }
}

