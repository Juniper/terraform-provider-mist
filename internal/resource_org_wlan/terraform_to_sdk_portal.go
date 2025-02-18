package resource_org_wlan

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func portalTerraformToSdk(plan PortalValue) *models.WlanPortal {

	sponsors := make(map[string]string)
	for k, v := range plan.Sponsors.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(basetypes.StringValue)
		sponsors[k] = vPlan.ValueString()
	}

	data := models.WlanPortal{}
	if plan.AllowWlanIdRoam.ValueBoolPointer() != nil {
		data.AllowWlanIdRoam = plan.AllowWlanIdRoam.ValueBoolPointer()
	}
	if plan.AmazonClientId.ValueStringPointer() != nil {
		data.AmazonClientId = models.NewOptional(plan.AmazonClientId.ValueStringPointer())
	}
	if plan.AmazonClientId.ValueStringPointer() != nil {
		data.AmazonClientId = models.NewOptional(plan.AmazonClientId.ValueStringPointer())
	}
	if plan.AmazonClientSecret.ValueStringPointer() != nil {
		data.AmazonClientSecret = models.NewOptional(plan.AmazonClientSecret.ValueStringPointer())
	}
	if !plan.AmazonEmailDomains.IsNull() && !plan.AmazonEmailDomains.IsUnknown() {
		data.AmazonEmailDomains = misttransform.ListOfStringTerraformToSdk(plan.AmazonEmailDomains)
	}
	if plan.AmazonEnabled.ValueBoolPointer() != nil {
		data.AmazonEnabled = plan.AmazonEnabled.ValueBoolPointer()
	}
	if plan.AmazonExpire.ValueInt64Pointer() != nil {
		data.AmazonExpire = models.NewOptional(models.ToPointer(int(plan.AmazonExpire.ValueInt64())))
	}
	if plan.Auth.ValueStringPointer() != nil {
		data.Auth = models.ToPointer(models.WlanPortalAuthEnum(plan.Auth.ValueString()))
	}
	if plan.AzureClientId.ValueStringPointer() != nil {
		data.AzureClientId = models.NewOptional(plan.AzureClientId.ValueStringPointer())
	}
	if plan.AzureClientSecret.ValueStringPointer() != nil {
		data.AzureClientSecret = models.NewOptional(plan.AzureClientSecret.ValueStringPointer())
	}
	if plan.AzureEnabled.ValueBoolPointer() != nil {
		data.AzureEnabled = plan.AzureEnabled.ValueBoolPointer()
	}
	if plan.AzureExpire.ValueInt64Pointer() != nil {
		data.AzureExpire = models.NewOptional(models.ToPointer(int(plan.AzureExpire.ValueInt64())))
	}
	if plan.AzureTenantId.ValueStringPointer() != nil {
		data.AzureTenantId = models.NewOptional(plan.AzureTenantId.ValueStringPointer())
	}
	if plan.BroadnetPassword.ValueStringPointer() != nil {
		data.BroadnetPassword = plan.BroadnetPassword.ValueStringPointer()
	}
	if plan.BroadnetSid.ValueStringPointer() != nil {
		data.BroadnetSid = plan.BroadnetSid.ValueStringPointer()
	}
	if plan.BroadnetUserId.ValueStringPointer() != nil {
		data.BroadnetUserId = plan.BroadnetUserId.ValueStringPointer()
	}
	if plan.BypassWhenCloudDown.ValueBoolPointer() != nil {
		data.BypassWhenCloudDown = plan.BypassWhenCloudDown.ValueBoolPointer()
	}
	if plan.ClickatellApiKey.ValueStringPointer() != nil {
		data.ClickatellApiKey = plan.ClickatellApiKey.ValueStringPointer()
	}
	if plan.CrossSite.ValueBoolPointer() != nil {
		data.CrossSite = plan.CrossSite.ValueBoolPointer()
	}
	if plan.EmailEnabled.ValueBoolPointer() != nil {
		data.EmailEnabled = plan.EmailEnabled.ValueBoolPointer()
	}
	if plan.Enabled.ValueBoolPointer() != nil {
		data.Enabled = plan.Enabled.ValueBoolPointer()
	}
	if plan.Expire.ValueInt64Pointer() != nil {
		data.Expire = models.ToPointer(int(plan.Expire.ValueInt64()))
	}
	if plan.ExternalPortalUrl.ValueStringPointer() != nil {
		data.ExternalPortalUrl = plan.ExternalPortalUrl.ValueStringPointer()
	}
	if plan.FacebookClientId.ValueStringPointer() != nil {
		data.FacebookClientId = models.NewOptional(plan.FacebookClientId.ValueStringPointer())
	}
	if plan.FacebookClientSecret.ValueStringPointer() != nil {
		data.FacebookClientSecret = models.NewOptional(plan.FacebookClientSecret.ValueStringPointer())
	}
	if !plan.FacebookEmailDomains.IsNull() && !plan.MicrosoftEmailDomains.IsUnknown() {
		data.FacebookEmailDomains = misttransform.ListOfStringTerraformToSdk(plan.FacebookEmailDomains)
	}
	if plan.FacebookEnabled.ValueBoolPointer() != nil {
		data.FacebookEnabled = plan.FacebookEnabled.ValueBoolPointer()
	}
	if plan.FacebookExpire.ValueInt64Pointer() != nil {
		data.FacebookExpire = models.NewOptional(models.ToPointer(int(plan.FacebookExpire.ValueInt64())))
	}
	if plan.Forward.ValueBoolPointer() != nil {
		data.Forward = plan.Forward.ValueBoolPointer()
	}
	if plan.ForwardUrl.ValueStringPointer() != nil {
		data.ForwardUrl = models.NewOptional(plan.ForwardUrl.ValueStringPointer())
	}
	if plan.GoogleClientId.ValueStringPointer() != nil {
		data.GoogleClientId = models.NewOptional(plan.GoogleClientId.ValueStringPointer())
	}
	if plan.GoogleClientSecret.ValueStringPointer() != nil {
		data.GoogleClientSecret = models.NewOptional(plan.GoogleClientSecret.ValueStringPointer())
	}
	if !plan.GoogleEmailDomains.IsNull() && !plan.MicrosoftEmailDomains.IsUnknown() {
		data.GoogleEmailDomains = misttransform.ListOfStringTerraformToSdk(plan.GoogleEmailDomains)
	}
	if plan.GoogleEnabled.ValueBoolPointer() != nil {
		data.GoogleEnabled = plan.GoogleEnabled.ValueBoolPointer()
	}
	if plan.GoogleExpire.ValueInt64Pointer() != nil {
		data.GoogleExpire = models.NewOptional(models.ToPointer(int(plan.GoogleExpire.ValueInt64())))
	}
	if plan.GupshupPassword.ValueStringPointer() != nil {
		data.GupshupPassword = plan.GupshupPassword.ValueStringPointer()
	}
	if plan.GupshupUserid.ValueStringPointer() != nil {
		data.GupshupUserid = plan.GupshupUserid.ValueStringPointer()
	}
	if plan.MicrosoftClientId.ValueStringPointer() != nil {
		data.MicrosoftClientId = models.NewOptional(plan.MicrosoftClientId.ValueStringPointer())
	}
	if plan.MicrosoftClientSecret.ValueStringPointer() != nil {
		data.MicrosoftClientSecret = models.NewOptional(plan.MicrosoftClientSecret.ValueStringPointer())
	}
	if !plan.MicrosoftEmailDomains.IsNull() && !plan.MicrosoftEmailDomains.IsUnknown() {
		data.MicrosoftEmailDomains = misttransform.ListOfStringTerraformToSdk(plan.MicrosoftEmailDomains)
	}
	if plan.MicrosoftEnabled.ValueBoolPointer() != nil {
		data.MicrosoftEnabled = plan.MicrosoftEnabled.ValueBoolPointer()
	}
	if plan.MicrosoftExpire.ValueInt64Pointer() != nil {
		data.MicrosoftExpire = models.NewOptional(models.ToPointer(int(plan.MicrosoftExpire.ValueInt64())))
	}
	if plan.PassphraseEnabled.ValueBoolPointer() != nil {
		data.PassphraseEnabled = plan.PassphraseEnabled.ValueBoolPointer()
	}
	if plan.PassphraseExpire.ValueInt64Pointer() != nil {
		data.PassphraseExpire = models.NewOptional(models.ToPointer(int(plan.PassphraseExpire.ValueInt64())))
	}
	if plan.Password.ValueStringPointer() != nil {
		data.Password = models.NewOptional(plan.Password.ValueStringPointer())
	}
	if plan.PredefinedSponsorsEnabled.ValueBoolPointer() != nil {
		data.PredefinedSponsorsEnabled = plan.PredefinedSponsorsEnabled.ValueBoolPointer()
	}
	if plan.Privacy.ValueBoolPointer() != nil {
		data.Privacy = plan.Privacy.ValueBoolPointer()
	}
	if plan.PuzzelPassword.ValueStringPointer() != nil {
		data.PuzzelPassword = plan.PuzzelPassword.ValueStringPointer()
	}
	if plan.PuzzelServiceId.ValueStringPointer() != nil {
		data.PuzzelServiceId = plan.PuzzelServiceId.ValueStringPointer()
	}
	if plan.PuzzelUsername.ValueStringPointer() != nil {
		data.PuzzelUsername = plan.PuzzelUsername.ValueStringPointer()
	}
	if plan.SmsEnabled.ValueBoolPointer() != nil {
		data.SmsEnabled = plan.SmsEnabled.ValueBoolPointer()
	}
	if plan.SmsExpire.ValueInt64Pointer() != nil {
		data.SmsExpire = models.NewOptional(models.ToPointer(int(plan.SmsExpire.ValueInt64())))
	}
	if plan.SmsMessageFormat.ValueStringPointer() != nil {
		data.SmsMessageFormat = plan.SmsMessageFormat.ValueStringPointer()
	}
	if plan.SmsProvider.ValueStringPointer() != nil {
		data.SmsProvider = models.ToPointer(models.WlanPortalSmsProviderEnum(plan.SmsProvider.ValueString()))
	}
	if plan.SponsorAutoApprove.ValueBoolPointer() != nil {
		data.SponsorAutoApprove = plan.SponsorAutoApprove.ValueBoolPointer()
	}
	if !plan.SponsorEmailDomains.IsNull() && !plan.SponsorEmailDomains.IsUnknown() {
		data.SponsorEmailDomains = misttransform.ListOfStringTerraformToSdk(plan.SponsorEmailDomains)
	}
	if plan.SponsorEnabled.ValueBoolPointer() != nil {
		data.SponsorEnabled = plan.SponsorEnabled.ValueBoolPointer()
	}
	if plan.SponsorExpire.ValueInt64Pointer() != nil {
		data.SponsorExpire = models.NewOptional(models.ToPointer(int(plan.SponsorExpire.ValueInt64())))
	}
	if plan.SponsorLinkValidityDuration.ValueStringPointer() != nil {
		data.SponsorLinkValidityDuration = plan.SponsorLinkValidityDuration.ValueStringPointer()
	}
	if plan.SponsorNotifyAll.ValueBoolPointer() != nil {
		data.SponsorNotifyAll = plan.SponsorNotifyAll.ValueBoolPointer()
	}
	if plan.SponsorStatusNotify.ValueBoolPointer() != nil {
		data.SponsorStatusNotify = plan.SponsorStatusNotify.ValueBoolPointer()
	}
	if !plan.Sponsors.IsNull() && !plan.Sponsors.IsUnknown() {
		data.Sponsors = models.ToPointer(models.WlanPortalSponsorsContainer.FromMapOfString(sponsors))
	}
	if plan.SsoDefaultRole.ValueStringPointer() != nil {
		data.SsoDefaultRole = plan.SsoDefaultRole.ValueStringPointer()
	}
	if plan.SsoForcedRole.ValueStringPointer() != nil {
		data.SsoForcedRole = plan.SsoForcedRole.ValueStringPointer()
	}
	if plan.SsoIdpCert.ValueStringPointer() != nil {
		data.SsoIdpCert = plan.SsoIdpCert.ValueStringPointer()
	}
	if plan.SsoIdpSignAlgo.ValueStringPointer() != nil {
		data.SsoIdpSignAlgo = (*models.WlanPortalIdpSignAlgoEnum)(plan.SsoIdpSignAlgo.ValueStringPointer())
	}
	if plan.SsoIdpSsoUrl.ValueStringPointer() != nil {
		data.SsoIdpSsoUrl = plan.SsoIdpSsoUrl.ValueStringPointer()
	}
	if plan.SsoIssuer.ValueStringPointer() != nil {
		data.SsoIssuer = plan.SsoIssuer.ValueStringPointer()
	}
	if plan.SsoNameidFormat.ValueStringPointer() != nil {
		data.SsoNameidFormat = models.ToPointer(models.WlanPortalSsoNameidFormatEnum(plan.SsoNameidFormat.ValueString()))
	}
	if plan.TelstraClientId.ValueStringPointer() != nil {
		data.TelstraClientId = plan.TelstraClientId.ValueStringPointer()
	}
	if plan.TelstraClientSecret.ValueStringPointer() != nil {
		data.TelstraClientSecret = plan.TelstraClientSecret.ValueStringPointer()
	}
	if plan.TwilioAuthToken.ValueStringPointer() != nil {
		data.TwilioAuthToken = models.NewOptional(plan.TwilioAuthToken.ValueStringPointer())
	}
	if plan.TwilioPhoneNumber.ValueStringPointer() != nil {
		data.TwilioPhoneNumber = models.NewOptional(plan.TwilioPhoneNumber.ValueStringPointer())
	}
	if plan.TwilioSid.ValueStringPointer() != nil {
		data.TwilioSid = models.NewOptional(plan.TwilioSid.ValueStringPointer())
	}

	return &data
}
