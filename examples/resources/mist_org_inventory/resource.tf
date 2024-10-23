resource "mist_org_inventory" "inventory" {
  org_id = mist_org.terraform_test.id
  inventory = {
    # Device Claim Code
    "CPKL2EXXXXXXXXX" = {}
    "G87JHBFXXXXXXXX" = {
      site_id = mist_site.terraform_site.id
      unclaim_when_destroyed = true
    }
    # MAC Address
    "2c2131000000" = {
      site_id                = mist_site.terraform_site.id
      unclaim_when_destroyed = true
    }
    "2c2131000001" = {
      unclaim_when_destroyed = false
    }    
  }
}

# deprecated
resource "mist_org_inventory" "inventory" {
  org_id = mist_org.terraform_test.id
  devices = [
    # Device Claim Code
    {
      claim_code = "CPKL2EXXXXXXXXX"
      site_id = mist_site.terraform_site.id
      unclaim_when_destroyed = true
    },
    # MAC Address
    {
      mac = "2c2131000000"
      site_id                = mist_site.terraform_site.id
      unclaim_when_destroyed = true
    }   
  ]
}