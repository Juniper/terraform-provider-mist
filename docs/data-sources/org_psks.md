---
page_title: "mist_org_psks Data Source - terraform-provider-mist"
subcategory: "Wi-Fi Assurance"
description: |-
  This data source provides the list Org Psks.
  A multi PSK (Pre-Shared Key) is a feature that allows the use of multiple PSKs for securing network connections.It provides a simple and comprehensive way to onboard client devices without relying on client mac addresses.Each psk has its own key name, which can be used for user-level accountability, key rotation, and visibility in the management platform. It supports the creation, rotation, and auto-expiration of psks, and allows vlan assignment and role assignment for dynamic per-user policies.Multi PSKs create virtual broadcast domains and can be used for end-user onboarding via authenticated sso login.
---

# mist_org_psks (Data Source)

This data source provides the list Org Psks.

A multi PSK (Pre-Shared Key) is a feature that allows the use of multiple PSKs for securing network connections.  
It provides a simple and comprehensive way to onboard client devices without relying on client mac addresses.  
Each psk has its own key name, which can be used for user-level accountability, key rotation, and visibility in the management platform. It supports the creation, rotation, and auto-expiration of psks, and allows vlan assignment and role assignment for dynamic per-user policies.  
Multi PSKs create virtual broadcast domains and can be used for end-user onboarding via authenticated sso login.


## Example Usage

```terraform
data "mist_org_psks" "psks_vip" {
  org_id = "15fca2ac-b1a6-47cc-9953-cc6906281550"

  // Filtering options
  name = "psk_one"
  role = "vip"
  ssid = "psk_ssid"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String)

### Optional

- `name` (String)
- `role` (String)
- `ssid` (String)

### Read-Only

- `org_psks` (Attributes Set) (see [below for nested schema](#nestedatt--org_psks))

<a id="nestedatt--org_psks"></a>
### Nested Schema for `org_psks`

Read-Only:

- `admin_sso_id` (String) sso id for psk created from psk portal
- `created_time` (Number) When the object has been created, in epoch
- `email` (String) email to send psk expiring notifications to
- `expire_time` (Number) Expire time for this PSK key (epoch time in seconds). Default `null` (as no expiration)
- `expiry_notification_time` (Number) Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire
- `id` (String) Unique ID of the object instance in the Mist Organization
- `mac` (String) If `usage`==`single`, the mac that this PSK ties to, empty if `auto-binding`
- `macs` (List of String) If `usage`==`macs`, this list contains N number of client mac addresses or mac patterns(1122*) or both. This list is capped at 5000
- `max_usage` (Number) For Org PSK Only. Max concurrent users for this PSK key. Default is 0 (unlimited)
- `modified_time` (Number) When the object has been modified for the last time, in epoch
- `name` (String)
- `note` (String)
- `notify_expiry` (Boolean) If set to true, reminder notification will be sent when psk is about to expire
- `notify_on_create_or_edit` (Boolean) If set to true, notification will be sent when psk is created or edited
- `old_passphrase` (String, Sensitive) previous passphrase of the PSK if it has been rotated
- `org_id` (String)
- `passphrase` (String, Sensitive) passphrase of the PSK (8-63 character or 64 in hex)
- `role` (String)
- `ssid` (String) SSID this PSK should be applicable to
- `usage` (String) enum: `macs`, `multi`, `single`
- `vlan_id` (String)