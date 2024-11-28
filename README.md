# terraform-provider-mist

## Requirements
- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.8
- [Go](https://golang.org/doc/install) >= 1.22.7

## Getting Started

### Install Terraform
Instructions for popular operating systems can be found [here](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli).

### Create a Terraform configuration
The terraform configuration must:
* Be named with a .tf file extension.
* Reference this provider by its global address (`registry.terraform.io/Juniper/mist`) or just `Juniper/mist`.
* Include a provider configuration block which tells the provider which Mist Cloud must be used and which credentials to use.

```hcl
terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/Juniper/mist"
    }
  }
}

provider "mist" {
  host = "api.mist.com"
  apitoken = "xxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

### Terraform Init
Run the following at a command prompt while in the same directory as the configuration file to fetch the Mist provider plugin:

```console
terraform init
```

### Credentials
Mist credentials can be supplied in the provider configuration block or through environment variables (recommended):

* API Token:
```console
export MIST_APITOKEN=<apitoken>
```

* Username and Password:
```console
export MIST_USERNAME=<username>
export MIST_PASSWORD=<password>
```

### Start configuring resources
Full documentation for provider, resources and data sources can be found [here](https://registry.terraform.io/providers/Juniper/mist/latest/docs).

## Contributing
See the open issues for a full list of proposed features (and known issues).
