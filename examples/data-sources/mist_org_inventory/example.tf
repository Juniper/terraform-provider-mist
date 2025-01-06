data "mist_org_inventory" "inventory" {
  org_id = "15fca2ac-b1a6-47cc-9953-cc6906281550"

  // Filtering options
  mac        = "485a0d000000"
  model      = "EX4100-F-12P"
  serial     = "F00000000000"
  unassigned = false
  vc         = true
  vc_mac     = "485a0d000001"
  site_id    = "4a422ae5-7ca0-4599-87a3-8e49aa63685f"
}
