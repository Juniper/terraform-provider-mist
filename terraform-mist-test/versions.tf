terraform {
  required_providers {
    mist = {
      source  = "terraform.local/local/mist"
      version = "0.0.1"
    }
  }
}

provider "mist" {}