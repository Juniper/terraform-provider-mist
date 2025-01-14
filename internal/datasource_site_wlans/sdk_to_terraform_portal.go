package datasource_site_wlans

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func portalSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanPortal) basetypes.ObjectValue {
	var allow_wlan_id_roam basetypes.BoolValue
	var amazon_client_id basetypes.StringValue
	var amazon_client_secret basetypes.StringValue
	var amazon_email_domains basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var amazon_enabled basetypes.BoolValue
	var amazon_expire basetypes.Int64Value
	var auth basetypes.StringValue
	var azure_client_id basetypes.StringValue
	var azure_client_secret basetypes.StringValue
	var azure_enabled basetypes.BoolValue
	var azure_expire basetypes.Int64Value
	var azure_tenant_id basetypes.StringValue
	var broadnet_password basetypes.StringValue
	var broadnet_sid basetypes.StringValue
	var broadnet_user_id basetypes.StringValue
	var bypass_when_cloud_down basetypes.BoolValue
	var clickatell_api_key basetypes.StringValue
	var cross_site basetypes.BoolValue
	var email_enabled basetypes.BoolValue
	var enabled basetypes.BoolValue
	var expire basetypes.Int64Value
	var external_portal_url basetypes.StringValue = types.StringValue("")
	var facebook_client_id basetypes.StringValue
	var facebook_client_secret basetypes.StringValue
	var facebook_email_domains basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var facebook_enabled basetypes.BoolValue
	var facebook_expire basetypes.Int64Value
	var forward basetypes.BoolValue
	var forward_url basetypes.StringValue
	var google_client_id basetypes.StringValue
	var google_client_secret basetypes.StringValue
	var google_email_domains basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var google_enabled basetypes.BoolValue
	var google_expire basetypes.Int64Value
	var gupshup_password basetypes.StringValue
	var gupshup_userid basetypes.StringValue
	var microsoft_client_id basetypes.StringValue
	var microsoft_client_secret basetypes.StringValue
	var microsoft_email_domains basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var microsoft_enabled basetypes.BoolValue
	var microsoft_expire basetypes.Int64Value
	var passphrase_enabled basetypes.BoolValue
	var passphrase_expire basetypes.Int64Value
	var password basetypes.StringValue
	var predefined_sponsors_enabled basetypes.BoolValue
	var predefined_sponsors_hide_email basetypes.BoolValue = basetypes.NewBoolValue(false)
	var privacy basetypes.BoolValue
	var puzzel_password basetypes.StringValue
	var puzzel_service_id basetypes.StringValue
	var puzzel_username basetypes.StringValue
	var sms_enabled basetypes.BoolValue
	var sms_expire basetypes.Int64Value
	var sms_message_format basetypes.StringValue
	var sms_provider basetypes.StringValue
	var sponsor_auto_approve basetypes.BoolValue
	var sponsor_email_domains basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var sponsor_enabled basetypes.BoolValue
	var sponsor_expire basetypes.Int64Value
	var sponsor_link_validity_duration basetypes.StringValue
	var sponsor_notify_all basetypes.BoolValue
	var sponsor_status_notify basetypes.BoolValue
	var sponsors basetypes.MapValue = types.MapNull(types.StringType)
	var sso_default_role basetypes.StringValue
	var sso_forced_role basetypes.StringValue
	var sso_idp_cert basetypes.StringValue
	var sso_idp_sign_algo basetypes.StringValue
	var sso_idp_sso_url basetypes.StringValue
	var sso_issuer basetypes.StringValue
	var sso_nameid_format basetypes.StringValue
	var telstra_client_id basetypes.StringValue
	var telstra_client_secret basetypes.StringValue
	var twilio_auth_token basetypes.StringValue
	var twilio_phone_number basetypes.StringValue
	var twilio_sid basetypes.StringValue

	if d != nil && d.AllowWlanIdRoam != nil {
		allow_wlan_id_roam = types.BoolValue(*d.AllowWlanIdRoam)
	}
	if d != nil && d.AmazonClientId.Value() != nil {
		amazon_client_id = types.StringValue(*d.AmazonClientId.Value())
	}
	if d != nil && d.AmazonClientSecret.Value() != nil {
		amazon_client_secret = types.StringValue(*d.AmazonClientSecret.Value())
	}
	if d != nil && d.AmazonEmailDomains != nil {
		amazon_email_domains = mist_transform.ListOfStringSdkToTerraform(ctx, d.AmazonEmailDomains)
	}
	if d != nil && d.AmazonEnabled != nil {
		amazon_enabled = types.BoolValue(*d.AmazonEnabled)
	}
	if d != nil && d.AmazonExpire.Value() != nil {
		amazon_expire = types.Int64Value(int64(*d.AmazonExpire.Value()))
	}
	if d != nil && d.Auth != nil {
		auth = types.StringValue(string(*d.Auth))
	}
	if d != nil && d.AzureClientId.Value() != nil {
		azure_client_id = types.StringValue(*d.AzureClientId.Value())
	}
	if d != nil && d.AzureClientSecret.Value() != nil {
		azure_client_secret = types.StringValue(*d.AzureClientSecret.Value())
	}
	if d != nil && d.AzureEnabled != nil {
		azure_enabled = types.BoolValue(*d.AzureEnabled)
	}
	if d != nil && d.AzureExpire.Value() != nil {
		azure_expire = types.Int64Value(int64(*d.AzureExpire.Value()))
	}
	if d != nil && d.AzureTenantId.Value() != nil {
		azure_tenant_id = types.StringValue(*d.AzureTenantId.Value())
	}
	if d != nil && d.BroadnetPassword != nil {
		broadnet_password = types.StringValue(*d.BroadnetPassword)
	}
	if d != nil && d.BroadnetSid != nil {
		broadnet_sid = types.StringValue(*d.BroadnetSid)
	}
	if d != nil && d.BroadnetUserId != nil {
		broadnet_user_id = types.StringValue(*d.BroadnetUserId)
	}
	if d != nil && d.BypassWhenCloudDown != nil {
		bypass_when_cloud_down = types.BoolValue(*d.BypassWhenCloudDown)
	}
	if d != nil && d.ClickatellApiKey != nil {
		clickatell_api_key = types.StringValue(*d.ClickatellApiKey)
	}
	if d != nil && d.CrossSite != nil {
		cross_site = types.BoolValue(*d.CrossSite)
	}
	if d != nil && d.EmailEnabled != nil {
		email_enabled = types.BoolValue(*d.EmailEnabled)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Expire != nil {
		expire = types.Int64Value(int64(*d.Expire))
	}
	if d != nil && d.ExternalPortalUrl != nil {
		external_portal_url = types.StringValue(*d.ExternalPortalUrl)
	}
	if d != nil && d.FacebookClientId.Value() != nil {
		facebook_client_id = types.StringValue(*d.FacebookClientId.Value())
	}
	if d != nil && d.FacebookClientSecret.Value() != nil {
		facebook_client_secret = types.StringValue(*d.FacebookClientSecret.Value())
	}
	if d != nil && d.FacebookEmailDomains != nil {
		facebook_email_domains = mist_transform.ListOfStringSdkToTerraform(ctx, d.FacebookEmailDomains)
	}
	if d != nil && d.FacebookEnabled != nil {
		facebook_enabled = types.BoolValue(*d.FacebookEnabled)
	}
	if d != nil && d.FacebookExpire.Value() != nil {
		facebook_expire = types.Int64Value(int64(*d.FacebookExpire.Value()))
	}
	if d != nil && d.Forward != nil {
		forward = types.BoolValue(*d.Forward)
	}
	if d != nil && d.ForwardUrl.Value() != nil {
		forward_url = types.StringValue(*d.ForwardUrl.Value())
	}
	if d != nil && d.GoogleClientId.Value() != nil {
		google_client_id = types.StringValue(*d.GoogleClientId.Value())
	}
	if d != nil && d.GoogleClientSecret.Value() != nil {
		google_client_secret = types.StringValue(*d.GoogleClientSecret.Value())
	}
	if d != nil && d.GoogleEmailDomains != nil {
		google_email_domains = mist_transform.ListOfStringSdkToTerraform(ctx, d.GoogleEmailDomains)
	}
	if d != nil && d.GoogleEnabled != nil {
		google_enabled = types.BoolValue(*d.GoogleEnabled)
	}
	if d != nil && d.GoogleExpire.Value() != nil {
		google_expire = types.Int64Value(int64(*d.GoogleExpire.Value()))
	}
	if d != nil && d.GupshupPassword != nil {
		gupshup_password = types.StringValue(*d.GupshupPassword)
	}
	if d != nil && d.GupshupUserid != nil {
		gupshup_userid = types.StringValue(*d.GupshupUserid)
	}
	if d != nil && d.MicrosoftClientId.Value() != nil {
		microsoft_client_id = types.StringValue(*d.MicrosoftClientId.Value())
	}
	if d != nil && d.MicrosoftClientSecret.Value() != nil {
		microsoft_client_secret = types.StringValue(*d.MicrosoftClientSecret.Value())
	}
	if d != nil && d.MicrosoftEmailDomains != nil {
		microsoft_email_domains = mist_transform.ListOfStringSdkToTerraform(ctx, d.MicrosoftEmailDomains)
	}
	if d != nil && d.MicrosoftEnabled != nil {
		microsoft_enabled = types.BoolValue(*d.MicrosoftEnabled)
	}
	if d != nil && d.MicrosoftExpire.Value() != nil {
		microsoft_expire = types.Int64Value(int64(*d.MicrosoftExpire.Value()))
	}
	if d != nil && d.PassphraseEnabled != nil {
		passphrase_enabled = types.BoolValue(*d.PassphraseEnabled)
	}
	if d != nil && d.PassphraseExpire.Value() != nil {
		passphrase_expire = types.Int64Value(int64(*d.PassphraseExpire.Value()))
	}
	if d != nil && d.Password.Value() != nil {
		password = types.StringValue(*d.Password.Value())
	}
	if d != nil && d.PredefinedSponsorsEnabled != nil {
		predefined_sponsors_enabled = types.BoolValue(*d.PredefinedSponsorsEnabled)
	}
	if d != nil && d.PredefinedSponsorsHideEmail != nil {
		predefined_sponsors_hide_email = types.BoolValue(*d.PredefinedSponsorsHideEmail)
	}
	if d != nil && d.Privacy != nil {
		privacy = types.BoolValue(*d.Privacy)
	}
	if d != nil && d.PuzzelPassword != nil {
		puzzel_password = types.StringValue(*d.PuzzelPassword)
	}
	if d != nil && d.PuzzelServiceId != nil {
		puzzel_service_id = types.StringValue(*d.PuzzelServiceId)
	}
	if d != nil && d.PuzzelUsername != nil {
		puzzel_username = types.StringValue(*d.PuzzelUsername)
	}
	if d != nil && d.SmsEnabled != nil {
		sms_enabled = types.BoolValue(*d.SmsEnabled)
	}
	if d != nil && d.SmsExpire.Value() != nil {
		sms_expire = types.Int64Value(int64(*d.SmsExpire.Value()))
	}
	if d != nil && d.SmsMessageFormat != nil {
		sms_message_format = types.StringValue(*d.SmsMessageFormat)
	}
	if d != nil && d.SmsProvider != nil {
		sms_provider = types.StringValue(string(*d.SmsProvider))
	}
	if d != nil && d.SponsorAutoApprove != nil {
		sponsor_auto_approve = types.BoolValue(*d.SponsorAutoApprove)
	}
	if d != nil && d.SponsorEmailDomains != nil {
		sponsor_email_domains = mist_transform.ListOfStringSdkToTerraform(ctx, d.SponsorEmailDomains)
	}
	if d != nil && d.SponsorEnabled != nil {
		sponsor_enabled = types.BoolValue(*d.SponsorEnabled)
	}
	if d != nil && d.SponsorExpire.Value() != nil {
		sponsor_expire = types.Int64Value(int64(*d.SponsorExpire.Value()))
	}
	if d != nil && d.SponsorLinkValidityDuration != nil {
		sponsor_link_validity_duration = types.StringValue(*d.SponsorLinkValidityDuration)
	}
	if d != nil && d.SponsorNotifyAll != nil {
		sponsor_notify_all = types.BoolValue(*d.SponsorNotifyAll)
	}
	if d != nil && d.SponsorStatusNotify != nil {
		sponsor_status_notify = types.BoolValue(*d.SponsorStatusNotify)
	}
	if d != nil && d.Sponsors != nil {
		sponsors_attr := make(map[string]attr.Value)
		if s, ok := d.Sponsors.AsMapOfString(); ok {
			for k, v := range *s {
				sponsors_attr[k] = types.StringValue(string(v))
			}
		}
		sponsors = types.MapValueMust(types.StringType, sponsors_attr)
	}
	if d != nil && d.SsoDefaultRole != nil {
		sso_default_role = types.StringValue(*d.SsoDefaultRole)
	}
	if d != nil && d.SsoForcedRole != nil {
		sso_forced_role = types.StringValue(*d.SsoForcedRole)
	}
	if d != nil && d.SsoIdpCert != nil {
		sso_idp_cert = types.StringValue(*d.SsoIdpCert)
	}
	if d != nil && d.SsoIdpSignAlgo != nil {
		sso_idp_sign_algo = types.StringValue(string(*d.SsoIdpSignAlgo))
	}
	if d != nil && d.SsoIdpSsoUrl != nil {
		sso_idp_sso_url = types.StringValue(*d.SsoIdpSsoUrl)
	}
	if d != nil && d.SsoIssuer != nil {
		sso_issuer = types.StringValue(*d.SsoIssuer)
	}
	if d != nil && d.SsoNameidFormat != nil {
		sso_nameid_format = types.StringValue(string(*d.SsoNameidFormat))
	}
	if d != nil && d.TelstraClientId != nil {
		telstra_client_id = types.StringValue(*d.TelstraClientId)
	}
	if d != nil && d.TelstraClientSecret != nil {
		telstra_client_secret = types.StringValue(*d.TelstraClientSecret)
	}
	if d != nil && d.TwilioAuthToken.Value() != nil {
		twilio_auth_token = types.StringValue(*d.TwilioAuthToken.Value())
	}
	if d != nil && d.TwilioPhoneNumber.Value() != nil {
		twilio_phone_number = types.StringValue(*d.TwilioPhoneNumber.Value())
	}
	if d != nil && d.TwilioSid.Value() != nil {
		twilio_sid = types.StringValue(*d.TwilioSid.Value())
	}

	data_map_attr_type := PortalValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_wlan_id_roam":             allow_wlan_id_roam,
		"amazon_client_id":               amazon_client_id,
		"amazon_client_secret":           amazon_client_secret,
		"amazon_email_domains":           amazon_email_domains,
		"amazon_enabled":                 amazon_enabled,
		"amazon_expire":                  amazon_expire,
		"auth":                           auth,
		"azure_client_id":                azure_client_id,
		"azure_client_secret":            azure_client_secret,
		"azure_enabled":                  azure_enabled,
		"azure_expire":                   azure_expire,
		"azure_tenant_id":                azure_tenant_id,
		"broadnet_password":              broadnet_password,
		"broadnet_sid":                   broadnet_sid,
		"broadnet_user_id":               broadnet_user_id,
		"bypass_when_cloud_down":         bypass_when_cloud_down,
		"clickatell_api_key":             clickatell_api_key,
		"cross_site":                     cross_site,
		"email_enabled":                  email_enabled,
		"enabled":                        enabled,
		"expire":                         expire,
		"external_portal_url":            external_portal_url,
		"facebook_client_id":             facebook_client_id,
		"facebook_client_secret":         facebook_client_secret,
		"facebook_email_domains":         facebook_email_domains,
		"facebook_enabled":               facebook_enabled,
		"facebook_expire":                facebook_expire,
		"forward":                        forward,
		"forward_url":                    forward_url,
		"google_client_id":               google_client_id,
		"google_client_secret":           google_client_secret,
		"google_email_domains":           google_email_domains,
		"google_enabled":                 google_enabled,
		"google_expire":                  google_expire,
		"gupshup_password":               gupshup_password,
		"gupshup_userid":                 gupshup_userid,
		"microsoft_client_id":            microsoft_client_id,
		"microsoft_client_secret":        microsoft_client_secret,
		"microsoft_email_domains":        microsoft_email_domains,
		"microsoft_enabled":              microsoft_enabled,
		"microsoft_expire":               microsoft_expire,
		"passphrase_enabled":             passphrase_enabled,
		"passphrase_expire":              passphrase_expire,
		"password":                       password,
		"predefined_sponsors_enabled":    predefined_sponsors_enabled,
		"predefined_sponsors_hide_email": predefined_sponsors_hide_email,
		"privacy":                        privacy,
		"puzzel_password":                puzzel_password,
		"puzzel_service_id":              puzzel_service_id,
		"puzzel_username":                puzzel_username,
		"sms_enabled":                    sms_enabled,
		"sms_expire":                     sms_expire,
		"sms_message_format":             sms_message_format,
		"sms_provider":                   sms_provider,
		"sponsor_auto_approve":           sponsor_auto_approve,
		"sponsor_email_domains":          sponsor_email_domains,
		"sponsor_enabled":                sponsor_enabled,
		"sponsor_expire":                 sponsor_expire,
		"sponsor_link_validity_duration": sponsor_link_validity_duration,
		"sponsor_notify_all":             sponsor_notify_all,
		"sponsor_status_notify":          sponsor_status_notify,
		"sponsors":                       sponsors,
		"sso_default_role":               sso_default_role,
		"sso_forced_role":                sso_forced_role,
		"sso_idp_cert":                   sso_idp_cert,
		"sso_idp_sign_algo":              sso_idp_sign_algo,
		"sso_idp_sso_url":                sso_idp_sso_url,
		"sso_issuer":                     sso_issuer,
		"sso_nameid_format":              sso_nameid_format,
		"telstra_client_id":              telstra_client_id,
		"telstra_client_secret":          telstra_client_secret,
		"twilio_auth_token":              twilio_auth_token,
		"twilio_phone_number":            twilio_phone_number,
		"twilio_sid":                     twilio_sid,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
