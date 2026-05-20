
# Resource for importing existing mxcluster
# Import with: terraform import mist_org_mxcluster.existing_mxcluster "<org_id>.<mxcluster_id>"
resource "mist_org_mxcluster" "existing_mxcluster" {
  org_id = <org_id>
  name   = "edgey_cluster"
}


resource "mist_org_mxedge" "mxedge" {
  org_id       = mist_org_mxcluster.existing_mxcluster.org_id
  claim_code = <claim_code>
  name         = "me1"
}