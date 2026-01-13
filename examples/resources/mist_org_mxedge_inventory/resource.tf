

resource "mist_org_mxedge_inventory" "mxedges" {
  org_id = "b3b6ad7a-de9e-438c-8cdf-3ddcc1e124e2"
  mxedges = {
    "DKF4DSY7TZMDH9P" = {}
    "HGP4DSY7TZMD4F6" = {
      site_id = "971ecb5c-e694-4753-a867-42bf18d60e92"
    }
    "00000000-0000-0000-1000-d420b0f003ef" = {}   
    "00000000-0000-0000-1000-e320b0f003gh" = {
     site_id                = "971ecb5c-e694-4753-a867-42bf18d60e92"
    } 
  }
}

