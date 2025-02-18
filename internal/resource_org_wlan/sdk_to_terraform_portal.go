package resource_org_wlan

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func portalSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanPortal) PortalValue {
	var allowWlanIdRoam basetypes.BoolValue
	var amazonClientId basetypes.StringValue
	var amazonClientSecret basetypes.StringValue
	var amazonEmailDomains = misttransform.ListOfStringSdkToTerraformEmpty()
	var amazonEnabled basetypes.BoolValue
	var amazonExpire basetypes.Int64Value
	var auth basetypes.StringValue
	var azureClientId basetypes.StringValue
	var azureClientSecret basetypes.StringValue
	var azureEnabled basetypes.BoolValue
	var azureExpire basetypes.Int64Value
	var azureTenantId basetypes.StringValue
	var broadnetPassword basetypes.StringValue
	var broadnetSid basetypes.StringValue
	var broadnetUserId basetypes.StringValue
	var bypassWhenCloudDown basetypes.BoolValue
	var clickatellApiKey basetypes.StringValue
	var crossSite basetypes.BoolValue
	var emailEnabled basetypes.BoolValue
	var enabled basetypes.BoolValue
	var expire basetypes.Int64Value
	var externalPortalUrl = types.StringValue("")
	var facebookClientId basetypes.StringValue
	var facebookClientSecret basetypes.StringValue
	var facebookEmailDomains = misttransform.ListOfStringSdkToTerraformEmpty()
	var facebookEnabled basetypes.BoolValue
	var facebookExpire basetypes.Int64Value
	var forward basetypes.BoolValue
	var forwardUrl basetypes.StringValue
	var googleClientId basetypes.StringValue
	var googleClientSecret basetypes.StringValue
	var googleEmailDomains = misttransform.ListOfStringSdkToTerraformEmpty()
	var googleEnabled basetypes.BoolValue
	var googleExpire basetypes.Int64Value
	var gupshupPassword basetypes.StringValue
	var gupshupUserid basetypes.StringValue
	var microsoftClientId basetypes.StringValue
	var microsoftClientSecret basetypes.StringValue
	var microsoftEmailDomains = misttransform.ListOfStringSdkToTerraformEmpty()
	var microsoftEnabled basetypes.BoolValue
	var microsoftExpire basetypes.Int64Value
	var passphraseEnabled basetypes.BoolValue
	var passphraseExpire basetypes.Int64Value
	var password basetypes.StringValue
	var predefinedSponsorsEnabled basetypes.BoolValue
	var predefinedSponsorsHideEmail = basetypes.NewBoolValue(false)
	var privacy basetypes.BoolValue
	var puzzelPassword basetypes.StringValue
	var puzzelServiceId basetypes.StringValue
	var puzzelUsername basetypes.StringValue
	var smsEnabled basetypes.BoolValue
	var smsExpire basetypes.Int64Value
	var smsMessageFormat basetypes.StringValue
	var smsProvider basetypes.StringValue
	var sponsorAutoApprove basetypes.BoolValue
	var sponsorEmailDomains = misttransform.ListOfStringSdkToTerraformEmpty()
	var sponsorEnabled basetypes.BoolValue
	var sponsorExpire basetypes.Int64Value
	var sponsorLinkValidityDuration basetypes.StringValue
	var sponsorNotifyAll basetypes.BoolValue
	var sponsorStatusNotify basetypes.BoolValue
	var sponsors = types.MapNull(types.StringType)
	var ssoDefaultRole basetypes.StringValue
	var ssoForcedRole basetypes.StringValue
	var ssoIdpCert basetypes.StringValue
	var ssoIdpSignAlgo basetypes.StringValue
	var ssoIdpSsoUrl basetypes.StringValue
	var ssoIssuer basetypes.StringValue
	var ssoNameidFormat basetypes.StringValue
	var telstraClientId basetypes.StringValue
	var telstraClientSecret basetypes.StringValue
	var twilioAuthToken basetypes.StringValue
	var twilioPhoneNumber basetypes.StringValue
	var twilioSid basetypes.StringValue

	if d != nil && d.AllowWlanIdRoam != nil {
		allowWlanIdRoam = types.BoolValue(*d.AllowWlanIdRoam)
	}
	if d != nil && d.AmazonClientId.Value() != nil {
		amazonClientId = types.StringValue(*d.AmazonClientId.Value())
	}
	if d != nil && d.AmazonClientSecret.Value() != nil {
		amazonClientSecret = types.StringValue(*d.AmazonClientSecret.Value())
	}
	if d != nil && d.AmazonEmailDomains != nil {
		amazonEmailDomains = misttransform.ListOfStringSdkToTerraform(d.AmazonEmailDomains)
	}
	if d != nil && d.AmazonEnabled != nil {
		amazonEnabled = types.BoolValue(*d.AmazonEnabled)
	}
	if d != nil && d.AmazonExpire.Value() != nil {
		amazonExpire = types.Int64Value(int64(*d.AmazonExpire.Value()))
	}
	if d != nil && d.Auth != nil {
		auth = types.StringValue(string(*d.Auth))
	}
	if d != nil && d.AzureClientId.Value() != nil {
		azureClientId = types.StringValue(*d.AzureClientId.Value())
	}
	if d != nil && d.AzureClientSecret.Value() != nil {
		azureClientSecret = types.StringValue(*d.AzureClientSecret.Value())
	}
	if d != nil && d.AzureEnabled != nil {
		azureEnabled = types.BoolValue(*d.AzureEnabled)
	}
	if d != nil && d.AzureExpire.Value() != nil {
		azureExpire = types.Int64Value(int64(*d.AzureExpire.Value()))
	}
	if d != nil && d.AzureTenantId.Value() != nil {
		azureTenantId = types.StringValue(*d.AzureTenantId.Value())
	}
	if d != nil && d.BroadnetPassword != nil {
		broadnetPassword = types.StringValue(*d.BroadnetPassword)
	}
	if d != nil && d.BroadnetSid != nil {
		broadnetSid = types.StringValue(*d.BroadnetSid)
	}
	if d != nil && d.BroadnetUserId != nil {
		broadnetUserId = types.StringValue(*d.BroadnetUserId)
	}
	if d != nil && d.BypassWhenCloudDown != nil {
		bypassWhenCloudDown = types.BoolValue(*d.BypassWhenCloudDown)
	}
	if d != nil && d.ClickatellApiKey != nil {
		clickatellApiKey = types.StringValue(*d.ClickatellApiKey)
	}
	if d != nil && d.CrossSite != nil {
		crossSite = types.BoolValue(*d.CrossSite)
	}
	if d != nil && d.EmailEnabled != nil {
		emailEnabled = types.BoolValue(*d.EmailEnabled)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Expire != nil {
		expire = types.Int64Value(int64(*d.Expire))
	}
	if d != nil && d.ExternalPortalUrl != nil {
		externalPortalUrl = types.StringValue(*d.ExternalPortalUrl)
	}
	if d != nil && d.FacebookClientId.Value() != nil {
		facebookClientId = types.StringValue(*d.FacebookClientId.Value())
	}
	if d != nil && d.FacebookClientSecret.Value() != nil {
		facebookClientSecret = types.StringValue(*d.FacebookClientSecret.Value())
	}
	if d != nil && d.FacebookEmailDomains != nil {
		facebookEmailDomains = misttransform.ListOfStringSdkToTerraform(d.FacebookEmailDomains)
	}
	if d != nil && d.FacebookEnabled != nil {
		facebookEnabled = types.BoolValue(*d.FacebookEnabled)
	}
	if d != nil && d.FacebookExpire.Value() != nil {
		facebookExpire = types.Int64Value(int64(*d.FacebookExpire.Value()))
	}
	if d != nil && d.Forward != nil {
		forward = types.BoolValue(*d.Forward)
	}
	if d != nil && d.ForwardUrl.Value() != nil {
		forwardUrl = types.StringValue(*d.ForwardUrl.Value())
	}
	if d != nil && d.GoogleClientId.Value() != nil {
		googleClientId = types.StringValue(*d.GoogleClientId.Value())
	}
	if d != nil && d.GoogleClientSecret.Value() != nil {
		googleClientSecret = types.StringValue(*d.GoogleClientSecret.Value())
	}
	if d != nil && d.GoogleEmailDomains != nil {
		googleEmailDomains = misttransform.ListOfStringSdkToTerraform(d.GoogleEmailDomains)
	}
	if d != nil && d.GoogleEnabled != nil {
		googleEnabled = types.BoolValue(*d.GoogleEnabled)
	}
	if d != nil && d.GoogleExpire.Value() != nil {
		googleExpire = types.Int64Value(int64(*d.GoogleExpire.Value()))
	}
	if d != nil && d.GupshupPassword != nil {
		gupshupPassword = types.StringValue(*d.GupshupPassword)
	}
	if d != nil && d.GupshupUserid != nil {
		gupshupUserid = types.StringValue(*d.GupshupUserid)
	}
	if d != nil && d.MicrosoftClientId.Value() != nil {
		microsoftClientId = types.StringValue(*d.MicrosoftClientId.Value())
	}
	if d != nil && d.MicrosoftClientSecret.Value() != nil {
		microsoftClientSecret = types.StringValue(*d.MicrosoftClientSecret.Value())
	}
	if d != nil && d.MicrosoftEmailDomains != nil {
		microsoftEmailDomains = misttransform.ListOfStringSdkToTerraform(d.MicrosoftEmailDomains)
	}
	if d != nil && d.MicrosoftEnabled != nil {
		microsoftEnabled = types.BoolValue(*d.MicrosoftEnabled)
	}
	if d != nil && d.MicrosoftExpire.Value() != nil {
		microsoftExpire = types.Int64Value(int64(*d.MicrosoftExpire.Value()))
	}
	if d != nil && d.PassphraseEnabled != nil {
		passphraseEnabled = types.BoolValue(*d.PassphraseEnabled)
	}
	if d != nil && d.PassphraseExpire.Value() != nil {
		passphraseExpire = types.Int64Value(int64(*d.PassphraseExpire.Value()))
	}
	if d != nil && d.Password.Value() != nil {
		password = types.StringValue(*d.Password.Value())
	}
	if d != nil && d.PredefinedSponsorsEnabled != nil {
		predefinedSponsorsEnabled = types.BoolValue(*d.PredefinedSponsorsEnabled)
	}
	if d != nil && d.PredefinedSponsorsHideEmail != nil {
		predefinedSponsorsHideEmail = types.BoolValue(*d.PredefinedSponsorsHideEmail)
	}
	if d != nil && d.Privacy != nil {
		privacy = types.BoolValue(*d.Privacy)
	}
	if d != nil && d.PuzzelPassword != nil {
		puzzelPassword = types.StringValue(*d.PuzzelPassword)
	}
	if d != nil && d.PuzzelServiceId != nil {
		puzzelServiceId = types.StringValue(*d.PuzzelServiceId)
	}
	if d != nil && d.PuzzelUsername != nil {
		puzzelUsername = types.StringValue(*d.PuzzelUsername)
	}
	if d != nil && d.SmsEnabled != nil {
		smsEnabled = types.BoolValue(*d.SmsEnabled)
	}
	if d != nil && d.SmsExpire.Value() != nil {
		smsExpire = types.Int64Value(int64(*d.SmsExpire.Value()))
	}
	if d != nil && d.SmsMessageFormat != nil {
		smsMessageFormat = types.StringValue(*d.SmsMessageFormat)
	}
	if d != nil && d.SmsProvider != nil {
		smsProvider = types.StringValue(string(*d.SmsProvider))
	}
	if d != nil && d.SponsorAutoApprove != nil {
		sponsorAutoApprove = types.BoolValue(*d.SponsorAutoApprove)
	}
	if d != nil && d.SponsorEmailDomains != nil {
		sponsorEmailDomains = misttransform.ListOfStringSdkToTerraform(d.SponsorEmailDomains)
	}
	if d != nil && d.SponsorEnabled != nil {
		sponsorEnabled = types.BoolValue(*d.SponsorEnabled)
	}
	if d != nil && d.SponsorExpire.Value() != nil {
		sponsorExpire = types.Int64Value(int64(*d.SponsorExpire.Value()))
	}
	if d != nil && d.SponsorLinkValidityDuration != nil {
		sponsorLinkValidityDuration = types.StringValue(*d.SponsorLinkValidityDuration)
	}
	if d != nil && d.SponsorNotifyAll != nil {
		sponsorNotifyAll = types.BoolValue(*d.SponsorNotifyAll)
	}
	if d != nil && d.SponsorStatusNotify != nil {
		sponsorStatusNotify = types.BoolValue(*d.SponsorStatusNotify)
	}
	if d != nil && d.Sponsors != nil {
		sponsorsAttr := make(map[string]attr.Value)
		if s, ok := d.Sponsors.AsMapOfString(); ok {
			for k, v := range *s {
				sponsorsAttr[k] = types.StringValue(v)
			}
		}
		sponsors = types.MapValueMust(types.StringType, sponsorsAttr)
	}
	if d != nil && d.SsoDefaultRole != nil {
		ssoDefaultRole = types.StringValue(*d.SsoDefaultRole)
	}
	if d != nil && d.SsoForcedRole != nil {
		ssoForcedRole = types.StringValue(*d.SsoForcedRole)
	}
	if d != nil && d.SsoIdpCert != nil {
		ssoIdpCert = types.StringValue(*d.SsoIdpCert)
	}
	if d != nil && d.SsoIdpSignAlgo != nil {
		ssoIdpSignAlgo = types.StringValue(string(*d.SsoIdpSignAlgo))
	}
	if d != nil && d.SsoIdpSsoUrl != nil {
		ssoIdpSsoUrl = types.StringValue(*d.SsoIdpSsoUrl)
	}
	if d != nil && d.SsoIssuer != nil {
		ssoIssuer = types.StringValue(*d.SsoIssuer)
	}
	if d != nil && d.SsoNameidFormat != nil {
		ssoNameidFormat = types.StringValue(string(*d.SsoNameidFormat))
	}
	if d != nil && d.TelstraClientId != nil {
		telstraClientId = types.StringValue(*d.TelstraClientId)
	}
	if d != nil && d.TelstraClientSecret != nil {
		telstraClientSecret = types.StringValue(*d.TelstraClientSecret)
	}
	if d != nil && d.TwilioAuthToken.Value() != nil {
		twilioAuthToken = types.StringValue(*d.TwilioAuthToken.Value())
	}
	if d != nil && d.TwilioPhoneNumber.Value() != nil {
		twilioPhoneNumber = types.StringValue(*d.TwilioPhoneNumber.Value())
	}
	if d != nil && d.TwilioSid.Value() != nil {
		twilioSid = types.StringValue(*d.TwilioSid.Value())
	}

	dataMapValue := map[string]attr.Value{
		"allow_wlan_id_roam":             allowWlanIdRoam,
		"amazon_client_id":               amazonClientId,
		"amazon_client_secret":           amazonClientSecret,
		"amazon_email_domains":           amazonEmailDomains,
		"amazon_enabled":                 amazonEnabled,
		"amazon_expire":                  amazonExpire,
		"auth":                           auth,
		"azure_client_id":                azureClientId,
		"azure_client_secret":            azureClientSecret,
		"azure_enabled":                  azureEnabled,
		"azure_expire":                   azureExpire,
		"azure_tenant_id":                azureTenantId,
		"broadnet_password":              broadnetPassword,
		"broadnet_sid":                   broadnetSid,
		"broadnet_user_id":               broadnetUserId,
		"bypass_when_cloud_down":         bypassWhenCloudDown,
		"clickatell_api_key":             clickatellApiKey,
		"cross_site":                     crossSite,
		"email_enabled":                  emailEnabled,
		"enabled":                        enabled,
		"expire":                         expire,
		"external_portal_url":            externalPortalUrl,
		"facebook_client_id":             facebookClientId,
		"facebook_client_secret":         facebookClientSecret,
		"facebook_email_domains":         facebookEmailDomains,
		"facebook_enabled":               facebookEnabled,
		"facebook_expire":                facebookExpire,
		"forward":                        forward,
		"forward_url":                    forwardUrl,
		"google_client_id":               googleClientId,
		"google_client_secret":           googleClientSecret,
		"google_email_domains":           googleEmailDomains,
		"google_enabled":                 googleEnabled,
		"google_expire":                  googleExpire,
		"gupshup_password":               gupshupPassword,
		"gupshup_userid":                 gupshupUserid,
		"microsoft_client_id":            microsoftClientId,
		"microsoft_client_secret":        microsoftClientSecret,
		"microsoft_email_domains":        microsoftEmailDomains,
		"microsoft_enabled":              microsoftEnabled,
		"microsoft_expire":               microsoftExpire,
		"passphrase_enabled":             passphraseEnabled,
		"passphrase_expire":              passphraseExpire,
		"password":                       password,
		"predefined_sponsors_enabled":    predefinedSponsorsEnabled,
		"predefined_sponsors_hide_email": predefinedSponsorsHideEmail,
		"privacy":                        privacy,
		"puzzel_password":                puzzelPassword,
		"puzzel_service_id":              puzzelServiceId,
		"puzzel_username":                puzzelUsername,
		"sms_enabled":                    smsEnabled,
		"sms_expire":                     smsExpire,
		"sms_message_format":             smsMessageFormat,
		"sms_provider":                   smsProvider,
		"sponsor_auto_approve":           sponsorAutoApprove,
		"sponsor_email_domains":          sponsorEmailDomains,
		"sponsor_enabled":                sponsorEnabled,
		"sponsor_expire":                 sponsorExpire,
		"sponsor_link_validity_duration": sponsorLinkValidityDuration,
		"sponsor_notify_all":             sponsorNotifyAll,
		"sponsor_status_notify":          sponsorStatusNotify,
		"sponsors":                       sponsors,
		"sso_default_role":               ssoDefaultRole,
		"sso_forced_role":                ssoForcedRole,
		"sso_idp_cert":                   ssoIdpCert,
		"sso_idp_sign_algo":              ssoIdpSignAlgo,
		"sso_idp_sso_url":                ssoIdpSsoUrl,
		"sso_issuer":                     ssoIssuer,
		"sso_nameid_format":              ssoNameidFormat,
		"telstra_client_id":              telstraClientId,
		"telstra_client_secret":          telstraClientSecret,
		"twilio_auth_token":              twilioAuthToken,
		"twilio_phone_number":            twilioPhoneNumber,
		"twilio_sid":                     twilioSid,
	}
	data, e := NewPortalValue(PortalValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
