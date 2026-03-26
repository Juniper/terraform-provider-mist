terraform {
  required_providers {
    mist = {
      source = "juniper/mist"
    }
  }
}

provider "mist" {
  host      = "api.mist.com"
  apitoken = "EtsUU0JcxXpyFTZ6x4u4khaFXotuc3eW5UVhFZWkmOJi4DhUvZZaOmQGbB4Nu3GmBGoBgCKBC2PvPwGZSVet8iit0BwHeb3m"
}

# Resource for importing existing mxedge
# Import with: terraform import mist_org_mxedge.existing_mxedge "<org_id>.<mxedge_id>"
resource "mist_org_mxedge" "existing_mxedge" {
  org_id = "b3b6ad7a-de9e-438c-8cdf-3ddcc1e124e2"
  name   = "edgey_1"
  model = "ME-X1"
  site_id = "971ecb5c-e694-4753-a867-42bf18d60e92"
}


# Resource for importing existing mxedge
# Import with: terraform import mist_org_mxedge.existing_mxedge "<org_id>.<mxedge_id>"
resource "mist_org_mxedge" "new_mxedge" {
  org_id = "b3b6ad7a-de9e-438c-8cdf-3ddcc1e124e2"
  name   = "edgey_2"
  model = "ME-X1"
  site_id = "971ecb5c-e694-4753-a867-42bf18d60e92"
}