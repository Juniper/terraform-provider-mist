---
page_title: "mist_org_avprofiles Data Source - terraform-provider-mist"
subcategory: "WAN Assurance"
description: |-
  This data source provides the list of WAN Assurance Antivirus Profiles.
  An Antivirus Profile is used to configure the Antivirus feature on SRX devices. It specifies which content the Antivirus should analyse and which content should be ignored.
  The Antivirus profiles can be used within the following resources:
  mist_org_servicepolicy.antivirusmist_org_gatewaytemplate.service_policies.antivirusmist_org_deviceprofile_gateway.service_policies.antivirusmist_device_gateway.service_policies.antivirus
---

# mist_org_avprofiles (Data Source)

This data source provides the list of WAN Assurance Antivirus Profiles.

An Antivirus Profile is used to configure the Antivirus feature on SRX devices. It specifies which content the Antivirus should analyse and which content should be ignored.

The Antivirus profiles can be used within the following resources: 
 * `mist_org_servicepolicy.antivirus` 
 * `mist_org_gatewaytemplate.service_policies.antivirus` 
 * `mist_org_deviceprofile_gateway.service_policies.antivirus` 
 * `mist_device_gateway.service_policies.antivirus`


## Example Usage

```terraform
data "mist_org_avprofiles" "avprofiles" {
  org_id  = "15fca2ac-b1a6-47cc-9953-cc6906281550"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String)

### Read-Only

- `org_avprofiles` (Attributes Set) (see [below for nested schema](#nestedatt--org_avprofiles))

<a id="nestedatt--org_avprofiles"></a>
### Nested Schema for `org_avprofiles`

Read-Only:

- `created_time` (Number) When the object has been created, in epoch
- `fallback_action` (String) enum: `block`, `permit`
- `id` (String) Unique ID of the object instance in the Mist Organization
- `max_filesize` (Number) In KB
- `mime_whitelist` (List of String)
- `modified_time` (Number) When the object has been modified for the last time, in epoch
- `name` (String)
- `org_id` (String)
- `protocols` (List of String) List of protocols to monitor. enum: `ftp`, `http`, `imap`, `pop3`, `smtp`
- `url_whitelist` (List of String)