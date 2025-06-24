
terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/Juniper/mist"
    }
  }
}

provider "mist" {
  host = "api.mistsys.com"
  apitoken = "pe7stWeewIecNLcRRGnF9T98AQ8JTZOGCwja9Iu8xQS9KhjG9bKCCMnaWNsXd9qpfhQGTRh17ZUg6qcPDmOFxgXKYDVoE28h"
}
