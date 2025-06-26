
terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/Juniper/mist"
    }
  }
}

provider "mist" {
  host = "api.mistsys.com"
  apitoken = ""
}
