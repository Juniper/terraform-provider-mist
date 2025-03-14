---
page_title: "mist_org_sso_role Resource - terraform-provider-mist"
subcategory: "Org"
description: |-
  This resource manages Org SSO Roles for Admin Authentication.
  SSO roles refer to the different functions assigned to users within a Single Sign-On (SSO) system.These roles determine the tasks and actions that users can perform within the SSO system. There are typically predefined roles and custom roles in an SSO system.Roles in SSO provide a well-defined separation of responsibility and visibility, allowing for granular-level access control on SSO objects.
---

# mist_org_sso_role (Resource)

This resource manages Org SSO Roles for Admin Authentication.

SSO roles refer to the different functions assigned to users within a Single Sign-On (SSO) system.  
These roles determine the tasks and actions that users can perform within the SSO system. There are typically predefined roles and custom roles in an SSO system.  
Roles in SSO provide a well-defined separation of responsibility and visibility, allowing for granular-level access control on SSO objects.


## Example Usage

```terraform
resource "mist_org_sso_role" "sso_role_one" {
  org_id = mist_org.terraform_test.id
  name   = "admin_sso"
  privileges = [
    {
      scope   = "site"
      role    = "read"
      site_id = mist_site.terraform_site.id
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `org_id` (String)
- `privileges` (Attributes List) (see [below for nested schema](#nestedatt--privileges))

### Read-Only

- `id` (String) Unique ID of the object instance in the Mist Organization

<a id="nestedatt--privileges"></a>
### Nested Schema for `privileges`

Required:

- `role` (String) access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write`
- `scope` (String) enum: `org`, `site`, `sitegroup`

Optional:

- `site_id` (String) Required if `scope`==`site`
- `sitegroup_id` (String) Required if `scope`==`sitegroup`
- `views` (List of String) Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users. Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.  
You can define custom roles by adding the `views` attribute along with `role` when assigning privileges.  
Below are the list of supported UI views. Note that this is UI only feature.  

  | UI View | Required Role | Description |
  | --- | --- | --- |
  | `reporting` | `read` | full access to all analytics tools |
  | `marketing` | `read` | can view analytics and location maps |
  | `super_observer` | `read` | can view all the organization except the subscription page |
  | `location` | `write` | can view and manage location maps, can view analytics |
  | `security` | `write` | can view and manage site labels, policies and security |
  | `switch_admin` | `helpdesk` | can view and manage Switch ports, can view wired clients |
  | `mxedge_admin` | `admin` | can view and manage Mist edges and Mist tunnels |
  | `lobby_admin` | `admin` | full access to Org and Site Pre-shared keys |



## Import
Using `terraform import`, import `mist_org_sso_role` with:
```shell
# Org PSK can be imported by specifying the org_id and the sso_role_id
terraform import mist_org_sso_role.sso_role_one 17b46405-3a6d-4715-8bb4-6bb6d06f316a.d3c42998-9012-4859-9743-6b9bee475309
```