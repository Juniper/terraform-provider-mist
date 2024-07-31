# terraform-provider-mist


## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21
- [Mist API Go package](https://pkg.go.dev/github.com/tmunzer/mistapi-go) >= 0.2.5

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Using the provider

There are two ways to get and use the provider.

* Downloading & installing it from registry.terraform.io
* Building it from source

### From registry 
To install this provider, copy and paste this code into your Terraform configuration. Then, run terraform init.

```terraform
terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  host     = "api.mist.com"
  apitoken = "xxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

### From build
1. Clone this repository
2. From the `terraform-provider-mist` folder, do 
  * `go mod tidy` to install the depencies
  * `go install .` to install the provider
3. create a `.terraformrc` file in your home folder with (replace `<home_folder_path>` with your actual home folder paht):
```
provider_installation {
  dev_overrides {
    "registry.terraform.io/juniper/mist" = "<home_folder_path>/go/bin/",

  }
  direct {}
}
```
4. Create a terraform configuration file (must be name with the `.tf` extension):
```terraform
terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  host     = "api.mist.com"
  apitoken = "xxxxxxxxxxxxxxxxxxxxxxxxx"
}
```


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
