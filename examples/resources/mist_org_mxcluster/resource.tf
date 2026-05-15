
terraform {
  required_providers {
    mist = {
      source = "juniper/mist"
    }
  }
}

provider "mist" {
  host = "api.mistsys.com"
  # api_token can be set via MIST_API_TOKEN environment variable
  # or uncomment below:
  # api_token = "yEtsUU0JcxXpyFTZ6x4u4khaFXotuc3eW5UVhFZWkmOJi4DhUvZZaOmQGbB4Nu3GmBGoBgCKBC2PvPwGZSVet8iit0BwHeb3m"
}

# Resource for importing existing mxcluster
# Import with: terraform import mist_org_mxcluster.existing_mxcluster "76bd3d9a-48b8-468f-83c1-6fb2544f7815.ee32b2a5-1c0f-44ae-8539-6cc7394870c7"
resource "mist_org_mxcluster" "existing_mxcluster" {
  org_id = "76bd3d9a-48b8-468f-83c1-6fb2544f7815"
  name   = "edgey_cluster"
}


resource "mist_org_mxedge" "mxedge" {
  org_id       = mist_org_mxcluster.existing_mxcluster.org_id
  claim_code = "EKF4DSY7TZMDH9P"
  name         = "me1"
}