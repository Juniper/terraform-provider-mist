resource "mist_org_idpprofile" "idpprofile_one" {
  org_id       = mist_org.terraform_test.id
  base_profile = "standard"
  overwrites = [
    {
      name = "server_bypass"
      matching = {
        severity    = []
        dst_subnet  = []
        attack_name = ["SSL:OVERFLOW:KEY-ARG-NO-ENTROPY"]
      }
    },
    {
      name = "guest-bypass"
      matching = {
        severity   = []
        dst_subnet = ["8.8.8.8/32"]
        attack_name = ["UDP:ZERO-DATA"
        ]
      }
    }
  ]
  name = "idpprofile_one"
}
