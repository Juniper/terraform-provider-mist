terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  host     = "api.mist.com"
  apitoken = "xxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
