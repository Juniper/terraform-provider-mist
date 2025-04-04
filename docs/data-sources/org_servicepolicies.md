---
page_title: "mist_org_servicepolicies Data Source - terraform-provider-mist"
subcategory: "WAN Assurance"
description: |-
  This data source provides the list of WAN Assurance Service Policies (Application Policies).
  The Service Policies can be used in the service_policies object by referencing the Service Policy ID as the servicepolicy_id in:
  the Gateway configuration (mist_device_gateway.service_policies)the Gateway Templates (mist_org_gatewaytemplate.service_policies)the HUB Profiles (mist_org_deviceprofile_gateway.service_policies)
  They can be used to manage common policies between multiples configurations
---

# mist_org_servicepolicies (Data Source)

This data source provides the list of WAN Assurance Service Policies (Application Policies).

The Service Policies can be used in the `service_policies` object by referencing the Service Policy ID as the `servicepolicy_id` in:
* the Gateway configuration (`mist_device_gateway.service_policies`)
* the Gateway Templates (`mist_org_gatewaytemplate.service_policies`)
* the HUB Profiles (`mist_org_deviceprofile_gateway.service_policies`)
They can be used to manage common policies between multiples configurations


## Example Usage

```terraform
data "mist_org_servicepolicies" "servicepolicies" {
  org_id  = "15fca2ac-b1a6-47cc-9953-cc6906281550"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String)

### Read-Only

- `org_servicepolicies` (Attributes Set) (see [below for nested schema](#nestedatt--org_servicepolicies))

<a id="nestedatt--org_servicepolicies"></a>
### Nested Schema for `org_servicepolicies`

Read-Only:

- `aamw` (Attributes) For SRX Only (see [below for nested schema](#nestedatt--org_servicepolicies--aamw))
- `action` (String) enum: `allow`, `deny`
- `antivirus` (Attributes) For SRX-only (see [below for nested schema](#nestedatt--org_servicepolicies--antivirus))
- `appqoe` (Attributes) For SRX Only (see [below for nested schema](#nestedatt--org_servicepolicies--appqoe))
- `created_time` (Number) When the object has been created, in epoch
- `ewf` (Attributes List) (see [below for nested schema](#nestedatt--org_servicepolicies--ewf))
- `id` (String) Unique ID of the object instance in the Mist Organization
- `idp` (Attributes) (see [below for nested schema](#nestedatt--org_servicepolicies--idp))
- `local_routing` (Boolean) access within the same VRF
- `modified_time` (Number) When the object has been modified for the last time, in epoch
- `name` (String)
- `org_id` (String)
- `path_preference` (String) By default, we derive all paths available and use them, optionally, you can customize by using `path_preference`
- `services` (List of String)
- `ssl_proxy` (Attributes) For SRX-only (see [below for nested schema](#nestedatt--org_servicepolicies--ssl_proxy))
- `tenants` (List of String)

<a id="nestedatt--org_servicepolicies--aamw"></a>
### Nested Schema for `org_servicepolicies.aamw`

Read-Only:

- `aamwprofile_id` (String) org-level Advanced Advance Anti Malware Profile (SkyAtp) Profile can be used, this takes precedence over 'profile'
- `enabled` (Boolean)
- `profile` (String) enum: `docsonly`, `executables`, `standard`


<a id="nestedatt--org_servicepolicies--antivirus"></a>
### Nested Schema for `org_servicepolicies.antivirus`

Read-Only:

- `avprofile_id` (String) org-level AV Profile can be used, this takes precedence over 'profile'
- `enabled` (Boolean)
- `profile` (String) Default / noftp / httponly / or keys from av_profiles


<a id="nestedatt--org_servicepolicies--appqoe"></a>
### Nested Schema for `org_servicepolicies.appqoe`

Read-Only:

- `enabled` (Boolean)


<a id="nestedatt--org_servicepolicies--ewf"></a>
### Nested Schema for `org_servicepolicies.ewf`

Read-Only:

- `alert_only` (Boolean)
- `block_message` (String)
- `enabled` (Boolean)
- `profile` (String) enum: `critical`, `standard`, `strict`


<a id="nestedatt--org_servicepolicies--idp"></a>
### Nested Schema for `org_servicepolicies.idp`

Read-Only:

- `alert_only` (Boolean)
- `enabled` (Boolean)
- `idpprofile_id` (String) org_level IDP Profile can be used, this takes precedence over `profile`
- `profile` (String) enum: `Custom`, `strict` (default), `standard` or keys from idp_profiles


<a id="nestedatt--org_servicepolicies--ssl_proxy"></a>
### Nested Schema for `org_servicepolicies.ssl_proxy`

Read-Only:

- `ciphers_category` (String) enum: `medium`, `strong`, `weak`
- `enabled` (Boolean)