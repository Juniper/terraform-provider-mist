---
page_title: "mist_site_wlan_portal_template Resource - terraform-provider-mist"
subcategory: "Wi-Fi Assurance"
description: |-
  This resource is used customize the WLAN Guest Portal.
  The WLAN Portal Template can be used to define:
  Guest Authentication methods and parameters (access duration, ...)Default values of the text fields and labels on the portalValues of the text fields and labels based on the User Agent (locales property)
  Notes:
  There is no feedback from the API, so there is no possibility to validate the changes. The resource states is directly generated based on the resource plan.* There is no option to delete or revert the changes. Deleting the resource will just remove it from the states. Once removed, it is possible to create a new one. It will replace the previous template
---

# mist_site_wlan_portal_template (Resource)

This resource is used customize the WLAN Guest Portal.
The WLAN Portal Template can be used to define:
* Guest Authentication methods and parameters (access duration, ...)
* Default values of the text fields and labels on the portal
* Values of the text fields and labels based on the User Agent (`locales` property)

**Notes:**
* There is no feedback from the API, so there is no possibility to validate the changes. The resource states is directly generated based on the resource plan.* There is no option to delete or revert the changes. Deleting the resource will just remove it from the states. Once removed, it is possible to create a new one. It will replace the previous template


## Example Usage

```terraform
resource "mist_site_wlan_portal_template" "wlan_one" {
  site_id = mist_site.terraform_test.id
  wlan_id = mist_site_wlan.wlan_one.id
  portal_template = {
    sms_message_format    = "Code {{code}} expires in {{duration}} minutes."
    sms_validity_duration = "10"
    page_title            = "Welcome To My Demo Portal"
    locales = {
      "fr-FR" = {
        page_title = "Bienvenue sur mon portail de démo"
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `portal_template` (Attributes) Portal template wlan settings (see [below for nested schema](#nestedatt--portal_template))
- `site_id` (String)
- `wlan_id` (String) Site WLAN ID

<a id="nestedatt--portal_template"></a>
### Nested Schema for `portal_template`

Optional:

- `access_code_alternate_email` (String)
- `alignment` (String) defines alignment on portal. enum: `center`, `left`, `right`
- `auth_button_amazon` (String) Label for Amazon auth button
- `auth_button_azure` (String) Label for Azure auth button
- `auth_button_email` (String) Label for Email auth button
- `auth_button_facebook` (String) Label for Facebook auth button
- `auth_button_google` (String) Label for Google auth button
- `auth_button_microsoft` (String) Label for Microsoft auth button
- `auth_button_passphrase` (String) Label for passphrase auth button
- `auth_button_sms` (String) Label for SMS auth button
- `auth_button_sponsor` (String) Label for Sponsor auth button
- `auth_label` (String)
- `back_link` (String) Label of the link to go back to /logon
- `color` (String) Portal main color
- `color_dark` (String)
- `color_light` (String)
- `company` (Boolean) Whether company field is required
- `company_error` (String) Error message when company not provided
- `company_label` (String) Label of company field
- `email` (Boolean) Whether email field is required
- `email_access_domain_error` (String) Error message when a user has valid social login but doesn't match specified email domains.
- `email_cancel` (String) Label for cancel confirmation code submission using email auth
- `email_code_cancel` (String)
- `email_code_error` (String)
- `email_code_field_label` (String)
- `email_code_message` (String)
- `email_code_submit` (String)
- `email_code_title` (String)
- `email_error` (String) Error message when email not provided
- `email_field_label` (String)
- `email_label` (String) Label of email field
- `email_message` (String)
- `email_submit` (String) Label for confirmation code submit button using email auth
- `email_title` (String) Title for the Email registration
- `field1` (Boolean) Whether to ask field1
- `field1error` (String) Error message when field1 not provided
- `field1label` (String) Label of field1
- `field1required` (Boolean) Whether field1 is required field
- `field2` (Boolean) Whether to ask field2
- `field2error` (String) Error message when field2 not provided
- `field2label` (String) Label of field2
- `field2required` (Boolean) Whether field2 is required field
- `field3` (Boolean) Whether to ask field3
- `field3error` (String) Error message when field3 not provided
- `field3label` (String) Label of field3
- `field3required` (Boolean) Whether field3 is required field
- `field4` (Boolean) Whether to ask field4
- `field4error` (String) Error message when field4 not provided
- `field4label` (String) Label of field4
- `field4required` (Boolean) Whether field4 is required field
- `locales` (Attributes Map) Can be used to localize the portal based on the User Agent. Allowed property key values are:
  `ar`, `ca-ES`, `cs-CZ`, `da-DK`, `de-DE`, `el-GR`, `en-GB`, `en-US`, `es-ES`, `fi-FI`, `fr-FR`, 
  `he-IL`, `hi-IN`, `hr-HR`, `hu-HU`, `id-ID`, `it-IT`, `ja-J^`, `ko-KT`, `ms-MY`, `nb-NO`, `nl-NL`, 
  `pl-PL`, `pt-BR`, `pt-PT`, `ro-RO`, `ru-RU`, `sk-SK`, `sv-SE`, `th-TH`, `tr-TR`, `uk-UA`, `vi-VN`, 
  `zh-Hans`, `zh-Hant` (see [below for nested schema](#nestedatt--portal_template--locales))
- `logo` (String) path to the background image file. File must be a `png` image`
- `marketing_policy_link` (String) label of the link to go to /marketing_policy
- `marketing_policy_opt_in` (Boolean) Whether marketing policy optin is enabled
- `marketing_policy_opt_in_label` (String) label for marketing optin
- `marketing_policy_opt_in_text` (String) marketing policy text
- `message` (String)
- `multi_auth` (Boolean)
- `name` (Boolean) Whether name field is required
- `name_error` (String) Error message when name not provided
- `name_label` (String) Label of name field
- `opt_out_default` (Boolean) Default value for the `Do not store` checkbox
- `optout` (Boolean) Whether to display Do Not Store My Personal Information
- `optout_label` (String) Label for Do Not Store My Personal Information
- `page_title` (String)
- `passphrase_cancel` (String) Label for the Passphrase cancel button
- `passphrase_error` (String) Error message when invalid passphrase is provided
- `passphrase_label` (String) Passphrase
- `passphrase_message` (String)
- `passphrase_submit` (String) Label for the Passphrase submit button
- `passphrase_title` (String) Title for passphrase details page
- `powered_by` (Boolean) Whether to show \"Powered by Mist\"
- `privacy` (Boolean) Whether to require the Privacy Term acceptance
- `privacy_policy_accept_label` (String) Prefix of the label of the link to go to Privacy Policy
- `privacy_policy_error` (String) Error message when Privacy Policy not accepted
- `privacy_policy_link` (String) Label of the link to go to Privacy Policy
- `privacy_policy_text` (String) Text of the Privacy Policy
- `required_field_label` (String) Label to denote required field
- `responsive_layout` (Boolean)
- `sign_in_label` (String) Label of the button to signin
- `sms_carrier_default` (String)
- `sms_carrier_error` (String)
- `sms_carrier_field_label` (String) Label for mobile carrier drop-down list
- `sms_code_cancel` (String) Label for cancel confirmation code submission
- `sms_code_error` (String) Error message when confirmation code is invalid
- `sms_code_field_label` (String)
- `sms_code_message` (String)
- `sms_code_submit` (String) Label for confirmation code submit button
- `sms_code_title` (String)
- `sms_country_field_label` (String)
- `sms_country_format` (String)
- `sms_have_access_code` (String) Label for checkbox to specify that the user has access code
- `sms_is_twilio` (Boolean)
- `sms_message_format` (String) Format of access code sms message. {{code}} and {{duration}} are placeholders and should be retained as is.
- `sms_number_cancel` (String) Label for canceling mobile details for SMS auth
- `sms_number_error` (String)
- `sms_number_field_label` (String) Label for field to provide mobile number
- `sms_number_format` (String)
- `sms_number_message` (String)
- `sms_number_submit` (String) Label for submit button for code generation
- `sms_number_title` (String) Title for phone number details
- `sms_username_format` (String)
- `sms_validity_duration` (Number) How long confirmation code should be considered valid (in minutes)
- `sponsor_back_link` (String)
- `sponsor_cancel` (String)
- `sponsor_email` (String) Label for Sponsor Email
- `sponsor_email_error` (String)
- `sponsor_email_template` (String) HTML template to replace/override default sponsor email template 
Sponsor Email Template supports following template variables:
  * `approve_url`: Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized
  * `deny_url`: Renders URL to reject the request
  * `guest_email`: Renders Email ID of the guest
  * `guest_name`: Renders Name of the guest
  * `field1`: Renders value of the Custom Field 1
  * `field2`: Renders value of the Custom Field 2
  * `sponsor_link_validity_duration`: Renders validity time of the request (i.e. Approve/Deny URL)
  * `auth_expire_minutes`: Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes)
- `sponsor_info_approved` (String)
- `sponsor_info_denied` (String)
- `sponsor_info_pending` (String)
- `sponsor_name` (String) Label for Sponsor Name
- `sponsor_name_error` (String)
- `sponsor_note_pending` (String)
- `sponsor_request_access` (String) Submit button label request Wifi Access and notify sponsor about guest request
- `sponsor_status_approved` (String) Text to display if sponsor approves request
- `sponsor_status_denied` (String) Text to display when sponsor denies request
- `sponsor_status_pending` (String) Text to display if request is still pending
- `sponsor_submit` (String) Submit button label to notify sponsor about guest request
- `sponsors_error` (String)
- `sponsors_field_label` (String)
- `tos` (Boolean)
- `tos_accept_label` (String) Prefix of the label of the link to go to tos
- `tos_error` (String) Error message when tos not accepted
- `tos_link` (String) Label of the link to go to tos
- `tos_text` (String) Text of the Terms of Service

<a id="nestedatt--portal_template--locales"></a>
### Nested Schema for `portal_template.locales`

Optional:

- `auth_button_amazon` (String) Label for Amazon auth button
- `auth_button_azure` (String) Label for Azure auth button
- `auth_button_email` (String) Label for Email auth button
- `auth_button_facebook` (String) Label for Facebook auth button
- `auth_button_google` (String) Label for Google auth button
- `auth_button_microsoft` (String) Label for Microsoft auth button
- `auth_button_passphrase` (String) Label for passphrase auth button
- `auth_button_sms` (String) Label for SMS auth button
- `auth_button_sponsor` (String) Label for Sponsor auth button
- `auth_label` (String)
- `back_link` (String) Label of the link to go back to /logon
- `company_error` (String) Error message when company not provided
- `company_label` (String) Label of company field
- `email_access_domain_error` (String) Error message when a user has valid social login but doesn't match specified email domains.
- `email_cancel` (String) Label for cancel confirmation code submission using email auth
- `email_code_cancel` (String)
- `email_code_error` (String)
- `email_code_field_label` (String)
- `email_code_message` (String)
- `email_code_submit` (String)
- `email_code_title` (String)
- `email_error` (String) Error message when email not provided
- `email_field_label` (String)
- `email_label` (String) Label of email field
- `email_message` (String)
- `email_submit` (String) Label for confirmation code submit button using email auth
- `email_title` (String) Title for the Email registration
- `field1error` (String) Error message when field1 not provided
- `field1label` (String) Label of field1
- `field2error` (String) Error message when field2 not provided
- `field2label` (String) Label of field2
- `field3error` (String) Error message when field3 not provided
- `field3label` (String) Label of field3
- `field4error` (String) Error message when field4 not provided
- `field4label` (String) Label of field4
- `marketing_policy_link` (String) label of the link to go to /marketing_policy
- `marketing_policy_opt_in` (Boolean) Whether marketing policy optin is enabled
- `marketing_policy_opt_in_label` (String) label for marketing optin
- `marketing_policy_opt_in_text` (String) marketing policy text
- `message` (String)
- `name_error` (String) Error message when name not provided
- `name_label` (String) Label of name field
- `optout_label` (String) Label for Do Not Store My Personal Information
- `page_title` (String)
- `passphrase_cancel` (String) Label for the Passphrase cancel button
- `passphrase_error` (String) Error message when invalid passphrase is provided
- `passphrase_label` (String) Passphrase
- `passphrase_message` (String)
- `passphrase_submit` (String) Label for the Passphrase submit button
- `passphrase_title` (String) Title for passphrase details page
- `privacy_policy_accept_label` (String) Prefix of the label of the link to go to Privacy Policy
- `privacy_policy_error` (String) Error message when Privacy Policy not accepted
- `privacy_policy_link` (String) Label of the link to go to Privacy Policy
- `privacy_policy_text` (String) Text of the Privacy Policy
- `required_field_label` (String) Label to denote required field
- `sign_in_label` (String) Label of the button to signin
- `sms_carrier_default` (String)
- `sms_carrier_error` (String)
- `sms_carrier_field_label` (String) Label for mobile carrier drop-down list
- `sms_code_cancel` (String) Label for cancel confirmation code submission
- `sms_code_error` (String) Error message when confirmation code is invalid
- `sms_code_field_label` (String)
- `sms_code_message` (String)
- `sms_code_submit` (String) Label for confirmation code submit button
- `sms_code_title` (String)
- `sms_country_field_label` (String)
- `sms_country_format` (String)
- `sms_have_access_code` (String) Label for checkbox to specify that the user has access code
- `sms_message_format` (String) Format of access code sms message. {{code}} and {{duration}} are placeholders and should be retained as is.
- `sms_number_cancel` (String) Label for canceling mobile details for SMS auth
- `sms_number_error` (String)
- `sms_number_field_label` (String) Label for field to provide mobile number
- `sms_number_format` (String)
- `sms_number_message` (String)
- `sms_number_submit` (String) Label for submit button for code generation
- `sms_number_title` (String) Title for phone number details
- `sms_username_format` (String)
- `sponsor_back_link` (String)
- `sponsor_cancel` (String)
- `sponsor_email` (String) Label for Sponsor Email
- `sponsor_email_error` (String)
- `sponsor_info_approved` (String)
- `sponsor_info_denied` (String)
- `sponsor_info_pending` (String)
- `sponsor_name` (String) Label for Sponsor Name
- `sponsor_name_error` (String)
- `sponsor_note_pending` (String)
- `sponsor_request_access` (String) Submit button label request Wifi Access and notify sponsor about guest request
- `sponsor_status_approved` (String) Text to display if sponsor approves request
- `sponsor_status_denied` (String) Text to display when sponsor denies request
- `sponsor_status_pending` (String) Text to display if request is still pending
- `sponsor_submit` (String) Submit button label to notify sponsor about guest request
- `sponsors_error` (String)
- `sponsors_field_label` (String)
- `tos_accept_label` (String) Prefix of the label of the link to go to tos
- `tos_error` (String) Error message when tos not accepted
- `tos_link` (String) Label of the link to go to tos
- `tos_text` (String) Text of the Terms of Service


