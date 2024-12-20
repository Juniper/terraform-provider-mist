---
subcategory: "Release Notes"
page_title: "v0.2.12"
description: |-
    Release Notes for v0.2.12
---

# Release Notes for v0.2.12

**version**      : v0.2.12
**release date** : December 13th, 2024

!> This release may generate multiple changes to the `org_wlan_resource` and `site_wlan_resource` resources during the first configuration sync. This is due to the new default values defined, and will not impact to actual SSID configuration deployed on the Access Points

## Changes
### Documentation
* improve `org_wlan_resource` and `site_wlan_resource` resources documentation

### WLAN resources default values
Changes applied to `org_wlan_resource` and `site_wlan_resource` to reduce configuration drift when saving the WLAN from the Mist UI. These changes try to mimic the Mist UI default values, however, some of them are changing based on other parameter values which make it currently impossible to completely eliminate the configuration drift.

List of the default value changes:
| Attribute | Previous Default | New Default |
| --------- | ----------- | ---------------- |
| `acct_servers` | not set | `types.ListValueMust(AcctServersValue{}.Type(ctx), []attr.Value{})` |
| `airwatch.api_key` | not set | `""` |
| `airwatch.console_url` | not set | `""` |
| `airwatch.password` | not set | `""` |
| `airwatch.username` | not set | `""` |
| `airwatch` | not set | `types.ObjectValueMust(AirwatchValue{}.AttributeTypes(ctx), ... )` |
| `ap_ids` | `types.ListValueMust(AcctServersValue{}.Type(ctx), []attr.Value{})` | `types.ListNull(types.StringType)` |
| `app_limit` | not set | `types.ObjectValueMust(AppLimitValue{}.AttributeTypes(ctx), ... )` |
| `app_limit` | not set | `types.ObjectValueMust(AppQosValue{}.AttributeTypes(ctx), ... )` |
| `auth.anticlog_threshold` | `16` | removed | 
| `auth.keys` | not set | `types.ListValueMust(types.StringType, []attr.Value{types.StringValue(""),types.StringValue(""),types.StringValue(""),types.StringValue(""),})` | 
| `auth.owe` | `"disabled"` | removed | 
| `auth.wep_as_secondary_auth` | `false` | removed | 
| `auth_servers` | not set | `types.ListValueMust(AcctServersValue{}.Type(ctx), []attr.Value{})` |
| `auth_servers_nas_id` | not set | `""` |
| `auth_servers_nas_ip` | not set | `""` |
| `bonjour` | not set | `types.ObjectValueMust(BonjourValue{}.AttributeTypes(ctx), ... )` |
| `cisco_cwa` | not set | `types.ObjectValueMust(CiscoCwaValue{}.AttributeTypes(ctx), ... )` |
| `client_limit_down` | not set | `1000` |
| `client_limit_up` | not set | `512` |
| `coa_servers` | not set | `types.ListValueMust(CoaServersValue{}.Type(ctx), []attr.Value{})` |
| `disable_when_gateway_unreachable` | `false` | removed |
| `disable_when_mxtunnel_down` | `false` | removed |
| `dns_server_rewrite` | not set | `types.ObjectValueMust(DnsServerRewriteValue{}.AttributeTypes(ctx), ... )` |
| `hotspot20` | not set | `types.ObjectValueMust(Hotspot20Value{}.AttributeTypes(ctx), ... )` |
| `mist_nac` | not set | `types.ObjectValueMust(MistNacValue{}.AttributeTypes(ctx), ... )` |
| `mxtunnel_ids` | `types.ListValueMust(types.StringType, []attr.Value{})` | `types.ListNull(types.StringType)` |
| `mxtunnel_name` | `types.ListValueMust(types.StringType, []attr.Value{})` | `types.ListNull(types.StringType)` |
| `portal.allow_wlan_id_roam` | `false` | removed |
| `portal.amazon_email_domains` | `types.ListValueMust(types.StringType, []attr.Value{})` | `types.ListNull(types.StringType)` |
| `portal.broadnet_sid` | `"MIST"` | removed |
| `portal.broadnet_user_id` | `""` | removed |
| `portal.clickatell_api_key` | `""` | removed |
| `portal.cross_site` | `false` | removed |
| `portal.facebook_email_domains` | `types.ListValueMust(types.StringType, []attr.Value{})` | `types.ListNull(types.StringType)` |
| `portal.google_email_domains` | `types.ListValueMust(types.StringType, []attr.Value{})` | `types.ListNull(types.StringType)` |
| `portal.gupshup_password` | `""` | removed |
| `portal.gupshup_userid` | `""` | removed |
| `portal.microsoft_email_domains` | `types.ListValueMust(types.StringType, []attr.Value{})` | `types.ListNull(types.StringType)` |
| `portal.puzzel_password` | `""` | removed |
| `portal.puzzel_service_id` | `""` | removed |
| `portal.puzzel_username` | `""` | removed |
| `portal.sponsor_auto_approve` | `false` | removed |
| `portal.telstra_client_id` | `""` | removed |
| `portal.telstra_client_secret` | `""` | removed |
| `portal.twilio_auth_token` | `""` | removed |
| `portal.twilio_phone_number` | `""` | removed |
| `portal.twilio_sid` | `""` | removed |
| `portal`| not set | `types.ObjectValueMust(PortalValue{}.AttributeTypes(ctx), ... )` |
| `qos`| not set | `types.ObjectValueMust(ObjectValueMust{}.AttributeTypes(ctx), ... )` |
| `radsec`| not set | `types.ObjectValueMust(RadsecValue{}.AttributeTypes(ctx), ... )` |
| `rateset`| not set | `types.MapValueMust(RatesetValue{}.AttributeTypes(ctx), ... )` |
| `reconnect_clients_when_roaming_mxcluster`| `false`  | removed |
| `schedule`| not set | `types.ObjectValueMust(ScheduleValue{}.AttributeTypes(ctx), ... )` |
| `vlan_ids`| not set | `types.ListValueMust(types.StringType, []attr.Value{})` |
| `wlan_limit_down` | not set | `20000` |
| `wlan_limit_up` | not set | `10000` |

### WLAN resources validators
Validators applied to the WLAN resources attributes have been updates to simplify the resource configuration and improve the configuration validity.

List of the validator changes:
| Attribute | Previous Default | New Default |
| --------- | ----------- | ---------------- |
| `acct_servers`  | `listvalidator.SizeAtLeast(1)` | removed |
| `airwatch.api_key` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("enabled"), types.BoolValue(true))` |
| `airwatch.console_url` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("enabled"), types.BoolValue(true))` |
| `airwatch.password` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("enabled"), types.BoolValue(true))` |
| `airwatch.username` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("enabled"), types.BoolValue(true))` |
| `auth.key_idx`  | `mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("wep"))` | `int64validator.Between(1, 4)` |
| `auth.keys`  | `mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("wep"))` | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("wep"))` |
| `auth_servers`  | `listvalidator.SizeAtLeast(1)` | removed |
| `dynamic_vlan` | `mistvalidator.CannotBeTrueWhenValueIs(path.MatchRoot("vlan_enabled"), types.BoolValue(true))` | `mistvalidator.CanOnlyTrueWhenValueIs(path.MatchRoot("vlan_enabled"), types.BoolValue(true)), boolvalidator.Any( mistvalidator.CanOnlyTrueWhenValueIs(path.MatchRoot("auth").AtName("enable_mac_auth"), types.BoolValue(true)), mistvalidator.CanOnlyTrueWhenValueIs(path.MatchRoot("auth").AtName("type"), types.StringValue("eap")), mistvalidator.CanOnlyTrueWhenValueIs(path.MatchRoot("auth").AtName("type"), types.StringValue("eap192")))` |
| `portal.azure_client_id` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("azure_enabled"), types.BoolValue(true))` |
| `portal.azure_client_secret` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("azure_enabled"), types.BoolValue(true))` |
| `portal.azure_tenant_id` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("azure_enabled"), types.BoolValue(true))` |
| `portal.broadnet_password` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("broadnet"))` |
| `portal.broadnet_sid` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("broadnet"))` |
| `portal.broadnet_user_id` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("broadnet"))` |
| `portal.clickatell_api_key` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("clickatell"))` |
| `portal.external_portal_url` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("auth"), types.StringValue("external"))` |
| `portal.facebook_client_id` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("facebook_enabled"), types.BoolValue(true))` |
| `portal.facebook_client_secret` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("facebook_enabled"), types.BoolValue(true))` |
| `portal.forward_url` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("forward"), types.BoolValue(true))` |
| `portal.gupshup_password` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("gupshup"))` |
| `portal.gupshup_userid` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("gupshup"))` |
| `portal.password` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("passphrase_enabled"), types.BoolValue(true))` |
| `portal.puzzel_password` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("puzzel"))` |
| `portal.puzzel_service_id` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("puzzel"))` |
| `portal.puzzel_username` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("puzzel"))` |
| `portal.sms_provider` | `stringvalidator.OneOf("", "broadnet", "clickatell", "gupshup", "manual", "puzzel", "telstra", "twilio")` | `stringvalidator.OneOf("", "broadnet", "clickatell", "gupshup", "manual", "puzzel", "telstra", "twilio"), mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_enabled"), types.BoolValue(true))` |
| `portal.sso_idp_cert` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("auth"), types.StringValue("sso"))` |
| `portal.sso_idp_sso_url` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("auth"), types.StringValue("sso"))` |
| `portal.sso_issuer` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("auth"), types.StringValue("sso"))` |
| `portal.telstra_client_id` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("telstra"))` |
| `portal.telstra_client_secret` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("telstra"))` |
| `portal.twilio_auth_token` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("telstra"))` |
| `portal.twilio_phone_number` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("twilio"))` |
| `portal.twilio_sid` | not set | `mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("sms_provider"), types.StringValue("twilio"))` |
| `vlan_ids`| `listvalidator.ValueStringsAre(stringvalidator.Any(mistvalidator.ParseInt(1, 4094), mistvalidator.ParseVar())),` | `listvalidator.ValueStringsAre(stringvalidator.Any(mistvalidator.ParseInt(1, 4094), mistvalidator.ParseVar())),mistvalidator.RequiredWhenValueIs(path.MatchRoot("vlan_pooling"), types.BoolValue(true))` |
| `vlan_pooling` | not set | `mistvalidator.CanOnlyTrueWhenValueIs(path.MatchRoot("vlan_enabled"), types.BoolValue(true))` |

### Remove Attributes
| Attribute | Reason |
| --------- | ----------- |
| `dynamic_psk.vlan_ids` | OpenAPI Specification issue. This attribute is not supported by the Mist API |
| `portal_template_url` | Read Only attribute returned by the Mist API. The returned URL has limited lifetime so it doesn't make sense to store it in the resource state |
| `thumbnail` | Read Only attribute returned by the Mist API. The returned URL has limited lifetime so it doesn't make sense to store it in the resource state |

## Fixes
* [Issue 63](https://github.com/Juniper/terraform-provider-mist/issues/63): Adding `Optional` type to `alarmtemplate.rules.delivery`to fix synchronisation issue