---
subcategory: "Guides"
page_title: "Getting Started"
description: |-
    Getting Started
---

# Getting Started


This Guide provides a quick example to show how to configure and use the Juniper-Mist provider to deploy a Mist Organization and Site with a working Wired and Wireless network.

The complete example is available at the end of this page.


## Create the terraform Configuration
Before starting, create a new empty Folder we will use to store the configuration and the Terraform files. 

In this folder, create an empty file `main.tf`. This file will be used to define the Terraform configuration with the HCL Language. This configuration will store all the objects and their definition we want to deploy during this guide.

~> It is recommended to split the Terraform configuration into multiple files, but to simplify the example, all the configuration will be in the same file `main.tf`. 

### Configure the Provider
The first step is to configure the provider. To perform this task, it is required to already have an account on any supported Mist Cloud (see [Supported Mist Clouds](https://registry.terraform.io/providers/Juniper/mist/latest/docs)) with a [User API Token](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/topics/task/create-token-for-rest-api.html#task_wdg_4kw_dcc).

Then, in the `main.tf` we need to configure the provider:

```terraform
terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  // replace the host and apitoken values with Mist Cloud host you are using the API Token you created
  host     = "api.mist.com"
  apitoken = "xxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

Additional configuration parameters are described in the [Provider Schema](https://registry.terraform.io/providers/Juniper/mist/latest/docs#schema), and their corresponding environment variables are listed in the [Provider Environment Variables](https://registry.terraform.io/providers/Juniper/mist/latest/docs#environment-variables)


### Create the Mist Organization
In this example, all the objects will be deployed by Terraform, but it is also possible to reuse existing objects like an existing Mist Organization.


To create the Mist Organization, we need to define a [`mist_org` resource](https://registry.terraform.io/providers/Juniper/mist/latest/docs/resources/org) with. This example is only showing the minimum configuration required to create the Mist Organization, but additional parameters are listed in the [Mist Organization Schema](https://registry.terraform.io/providers/Juniper/mist/latest/docs/resources/org#schema).


```terraform
resource "mist_org" "org_one" {
  name             = "My Provider Organization"
}
```

### Create the Site Group
We will then create a Site Group. This Site Group will be used to assign the WLAN Template to the Site later in this guide.
Even if there is multiple ways to assign a WLAN Template to a Site, this approach allows us to create the objects in a streamlined way.


To create the Site Group, we will use the [`mist_org_sitegroup` resource](https://registry.terraform.io/providers/Juniper/mist/latest/docs/resources/org_sitegroup) and configure the Site Group Name and the Organization ID where this object will be created. This `org_id` will be a reference to the ID of the Mist Organization created with the `mist_org` resource above.

```terraform
resource "mist_org_sitegroup" "sitegroup_one" {
  org_id = mist_org.org_one.id
  name   = "My Provider Site Group"
}
```

### Create the WLAN Template and the WLAN
Next, we will use the [`mist_org_wlantemplate` resource](https://registry.terraform.io/providers/Juniper/mist/latest/docs/resources/org_wlantemplate) to create the WLAN Template, and assign it to the Site Group above.

```terraform
resource "mist_org_wlantemplate" "wlantemplate_one" {
  name   = "My Provider WLAN Template"
  org_id = mist_org.org_one.id
  applies = {
    sitegroup_ids = [
      mist_org_sitegroup.sitegroup_one.id
    ]
  }
}
```

Once we have the WLAN Template configuration, we can define the Org WLAN and assign it to the WLAN Template. For this purpose, we will use the [`mist_org_wlan` resource](https://registry.terraform.io/providers/Juniper/mist/latest/docs/resources/org_wlan):

```terraform
resource "mist_org_wlan" "wlan_one" {
  ssid              = "My Provider WLAN"
  org_id      = mist_org.org_one.id
  template_id = mist_org_wlantemplate.wlantemplate_one.id
  auth = {
    type = "psk"
    psk  = "secretpsk"
  }
}
```


### Create the Network Template
Then, we will create the Network (switch) Template. This is just an example of the minimum required configuration to use the [`mist_org_networktemplate` resource](https://registry.terraform.io/providers/Juniper/mist/latest/docs/resources/org_networktemplate):

```terraform
resource "mist_org_networktemplate" "networktemplate_one" {
  name   = "MyProviderNetworkTemplate"
  org_id = mist_org.org_one.id
  networks = {
    guest_vlan = {
      vlan_id = 10
    }
    user_vlan = {
      vlan_id = 11
    }
  }
  port_usages = {
    guest = {
      mode         = "access"
      port_network = "guest_vlan"
    }
    user = {
      mode      = "access"
      port_auth = "dot1x"
    }
    trunk = {
      all_networks = true
      enable_qos   = true
      mode         = "trunk"
      port_network = "guest_vlan"
    }
  }
  radius_config = {
    acct_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
    auth_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
  }
  switch_matching = {
    enable = true
    rules = [
      {
        name       = "switch_rule_one"
        match_name = "abc"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "guest"
          }
        }
      }
    ]
  }
}
```

### Create the Site
Now all the configuration is done, we can create the Mist Site with the [`mist_site` resource](https://registry.terraform.io/providers/Juniper/mist/latest/docs/resources/site) and put all together.


```terraform
resource "mist_site" "site_one" {
  org_id       = mist_org.org_one.id
  name         = "Site One"
  address      = "41 rue de Villiers, 92100 Neuilly sur Seine, France"
  country_code = "FR"
  sitegroup_ids = [
    mist_org_sitegroup.sitegroup_one.id
  ]
  networktemplate_id = mist_org_networktemplate.networktemplate_one.id
}
```

### Optional - Claim devices and assign them to the Site
To finish, we can claim devices to the Mist Organization and assign them to the Mist Site. To claim a device, it is required to have its Claim Code (the Claim Code can be found on the sticker with a QRCode directly on the hardware). 

~> For devices that are not "Cloud-Ready" (without Claim Code / QRcode), it is possible to manually adopt them with the corresponding configuration commands and then manage them with the provider. However, the adoption process cannot be done with by the provider.

To manage the Mist Organization inventory, we will use the [`mist_org_inventory` resource](https://registry.terraform.io/providers/Juniper/mist/latest/docs/resources/org_inventory):

```terraform
resource "mist_org_inventory" "inventory" {
  inventory = {
    // to claim a device with its Claim Code
    // replace "XXXXXXXXXXXXXXX" with the actual Claim Code
    "XXXXXXXXXXXXXXX" = {
      site_id = mist_site.terraform_site.id
      unclaim_when_destroyed = true
    }
    // to manage a device already claimed or adopted
    // replace "XXXXXXXXXXXXXXX" with the actual Claim Code
    "020004000000" = {
      site_id = mist_site.terraform_site.id
      unclaim_when_destroyed = false
    }
  }
}
```

## Use the Terraform CLI to deploy the configuration
When you [installed Terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli), this installed the Terraform CLI that will be used in this chapter.

### Initialize the Workspace
Before being able to deploy the created configuration with Terraform, it is required to initialize the Terraform Workspace. The `terraform init` command initializes a working directory containing Terraform configuration files. This is the first command that should be run after writing a new Terraform configuration or cloning an existing one from version control. It is safe to run this command multiple times.

Open a terminal and go to the directory where your `main.tf` file is located, then run `terraform init`

```shell
$ terraform init
Initializing the backend...
Initializing provider plugins...
- Finding latest version of juniper/mist...
- Installing juniper/mist v0.2.7...
- Installed juniper/mist v0.2.7 (signed by a HashiCorp partner, key ID 1211DC34850D21DE)
Partner and community providers are signed by their developers.
If you'd like to know more about provider signing, you can read about it here:
https://www.terraform.io/docs/cli/plugins/signing.html
Terraform has created a lock file .terraform.lock.hcl to record the provider
selections it made above. Include this file in your version control repository
so that Terraform can guarantee to make the same selections by default when
you run "terraform init" in the future.

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```

### Validate the configuration
To validate the planned configuration is valid, we can now use the command `terraform validate`. This will allow Terraform to parse the whole configuration and report any possible issue (missing attribute, wrong values, ...):

```shell
$ terraform validate
Success! The configuration is valid, but there were some validation warnings as shown above.
```

### Deploy the configuration
There is other Terraform commands that can be used before deploying the changes, but this is out of the scope of this guide. We will now directly deploy the configuration with the `terraform apply` command.

When using this command, Terraform will display the list of changes that will be applied and ask for a validation:
```shell
$ terraform apply

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

... 

Plan: 6 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

mist_org.org_one: Creating...
mist_org.org_one: Creation complete after 1s [id=e2e8566f-099c-4866-a44d-5c535656201a]
mist_org_sitegroup.sitegroup_one: Creating...
mist_org_networktemplate.networktemplate_one: Creating...
mist_org_sitegroup.sitegroup_one: Creation complete after 1s [id=01298c49-e81d-46fb-a5f3-ebd1a5ace2c9]
mist_org_wlantemplate.wlantemplate_one: Creating...
mist_site.site_one: Creating...
mist_org_networktemplate.networktemplate_one: Creation complete after 1s [id=5d176886-ff53-4f76-9dd3-49b5e78bba58]
mist_org_wlantemplate.wlantemplate_one: Creation complete after 0s [id=a5d8648e-9952-4180-8ffe-1cb28c773fbf]
mist_site.site_one: Creation complete after 0s [id=e2cc354f-0da1-4941-bd73-329918983958]
mist_org_wlan.wlan_one: Creating...
mist_org_wlan.wlan_one: Creation complete after 0s [id=50b5b3bb-11f5-4eec-89fb-ed7b761d943b]

Apply complete! Resources: 6 added, 0 changed, 0 destroyed.
```

The Organization is now created and configured. You can go to the Mist UI to access it and check the configuration.

You can also apply changes to the HCL configuration in the `main.tf` file, and use the `terraform apply` command to list the changes planned by Terraform and apply them to you Organization.

### Destroy the configuration
Congratulation, you reached the end of this guide. You can continue to discover the provider possibilities with the organization we just created, or you can ask Terraform to destroy it.

!> When validating the `terraform destroy` command, Terraform will delete all the objects deployed with this HCL configuration, the Organization itself included. The configuration done outside of Terraform will not be listed in the "destroy" command summary, but will be deleted when the Organization will be destroyed. Be sure to double check everything before using this command!

```shell
$ terraform destroy
mist_org.org_one: Refreshing state... [id=e2e8566f-099c-4866-a44d-5c535656201a]
mist_org_sitegroup.sitegroup_one: Refreshing state... [id=01298c49-e81d-46fb-a5f3-ebd1a5ace2c9]
mist_org_networktemplate.networktemplate_one: Refreshing state... [id=5d176886-ff53-4f76-9dd3-49b5e78bba58]
mist_org_wlantemplate.wlantemplate_one: Refreshing state... [id=a5d8648e-9952-4180-8ffe-1cb28c773fbf]
mist_site.site_one: Refreshing state... [id=e2cc354f-0da1-4941-bd73-329918983958]
mist_org_wlan.wlan_one: Refreshing state... [id=50b5b3bb-11f5-4eec-89fb-ed7b761d943b]

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  - destroy

Terraform will perform the following actions:

...

Plan: 0 to add, 0 to change, 6 to destroy.

Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes

mist_site.site_one: Destroying... [id=e2cc354f-0da1-4941-bd73-329918983958]
mist_org_wlan.wlan_one: Destroying... [id=50b5b3bb-11f5-4eec-89fb-ed7b761d943b]
mist_site.site_one: Destruction complete after 0s
mist_org_wlan.wlan_one: Destruction complete after 0s
mist_org_networktemplate.networktemplate_one: Destroying... [id=5d176886-ff53-4f76-9dd3-49b5e78bba58]
mist_org_wlantemplate.wlantemplate_one: Destroying... [id=a5d8648e-9952-4180-8ffe-1cb28c773fbf]
mist_org_wlantemplate.wlantemplate_one: Destruction complete after 0s
mist_org_networktemplate.networktemplate_one: Destruction complete after 0s
mist_org_sitegroup.sitegroup_one: Destroying... [id=01298c49-e81d-46fb-a5f3-ebd1a5ace2c9]
mist_org_sitegroup.sitegroup_one: Destruction complete after 0s
mist_org.org_one: Destroying... [id=e2e8566f-099c-4866-a44d-5c535656201a]
mist_org.org_one: Destruction complete after 1s

Destroy complete! Resources: 6 destroyed.
```


## Full HCL configuration
```terraform
// Provider configuration
terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  // replace the host and apitoken values with Mist Cloud host you are using the API Token you created
  host     = "api.mist.com"
  apitoken = "xxxxxxxxxxxxxxxxxxxxxxxxxxx"
}

// Create the Mist Organization
resource "mist_org" "org_one" {
  name = "My Provider Organization"
}

// Create the Site Group
resource "mist_org_sitegroup" "sitegroup_one" {
  org_id = mist_org.org_one.id
  name   = "My Provider Site Group"
}

// Create the WLAN Template and the WLAN
resource "mist_org_wlantemplate" "wlantemplate_one" {
  name   = "My Provider WLAN Template"
  org_id = mist_org.org_one.id
  applies = {
    sitegroup_ids = [
      mist_org_sitegroup.sitegroup_one.id
    ]
  }
}

resource "mist_org_wlan" "wlan_one" {
  ssid        = "My Provider WLAN"
  org_id      = mist_org.org_one.id
  template_id = mist_org_wlantemplate.wlantemplate_one.id
  auth = {
    type = "psk"
    psk  = "secretpsk"
  }
}

resource "mist_org_networktemplate" "networktemplate_one" {
  name   = "MyProviderNetworkTemplate"
  org_id = mist_org.org_one.id
  networks = {
    guest_vlan = {
      vlan_id = 10
    }
    user_vlan = {
      vlan_id = 11
    }
  }
  port_usages = {
    guest = {
      mode         = "access"
      port_network = "guest_vlan"
    }
    user = {
      mode      = "access"
      port_auth = "dot1x"
    }
    trunk = {
      all_networks = true
      enable_qos   = true
      mode         = "trunk"
      port_network = "guest_vlan"
    }
  }
  radius_config = {
    acct_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
    auth_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
  }
  switch_matching = {
    enable = true
    rules = [
      {
        name       = "switch_rule_one"
        match_name = "abc"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "guest"
          }
        }
      }
    ]
  }
}

resource "mist_site" "site_one" {
  org_id       = mist_org.org_one.id
  name         = "Site One"
  address      = "41 rue de Villiers, 92100 Neuilly sur Seine, France"
  country_code = "FR"
  sitegroup_ids = [
    mist_org_sitegroup.sitegroup_one.id
  ]
  networktemplate_id = mist_org_networktemplate.networktemplate_one.id
}

```
