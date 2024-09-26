package resource_org_wlan

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func portalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan PortalValue) *models.WlanPortal {

	sponsors := make(map[string]string)
	for k, v := range plan.Sponsors.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		sponsors[k] = v_plan.ValueString()
	}

	data := models.WlanPortal{}
	data.AmazonClientId = models.NewOptional(plan.AmazonClientId.ValueStringPointer())
	data.AmazonClientId = models.NewOptional(plan.AmazonClientId.ValueStringPointer())
	data.AmazonClientSecret = models.NewOptional(plan.AmazonClientSecret.ValueStringPointer())
	data.AmazonEmailDomains = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AmazonEmailDomains)
	data.AmazonEnabled = plan.AmazonEnabled.ValueBoolPointer()
	data.AmazonExpire = models.NewOptional(models.ToPointer(plan.AmazonExpire.ValueFloat64()))
	data.Auth = models.ToPointer(models.WlanPortalAuthEnum(string(plan.Auth.ValueString())))
	data.AzureClientId = models.NewOptional(plan.AzureClientId.ValueStringPointer())
	data.AzureClientSecret = models.NewOptional(plan.AzureClientSecret.ValueStringPointer())
	data.AzureEnabled = plan.AzureEnabled.ValueBoolPointer()
	data.AzureExpire = models.NewOptional(models.ToPointer(plan.AzureExpire.ValueFloat64()))
	data.AzureTenantId = models.NewOptional(plan.AzureTenantId.ValueStringPointer())
	data.BroadnetPassword = plan.BroadnetPassword.ValueStringPointer()
	data.BroadnetSid = plan.BroadnetSid.ValueStringPointer()
	data.BroadnetUserId = plan.BroadnetUserId.ValueStringPointer()
	data.BypassWhenCloudDown = plan.BypassWhenCloudDown.ValueBoolPointer()
	data.ClickatellApiKey = plan.ClickatellApiKey.ValueStringPointer()
	data.CrossSite = plan.CrossSite.ValueBoolPointer()
	data.EmailEnabled = plan.EmailEnabled.ValueBoolPointer()
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.Expire = models.ToPointer(plan.Expire.ValueFloat64())
	data.ExternalPortalUrl = plan.ExternalPortalUrl.ValueStringPointer()
	data.FacebookClientId = models.NewOptional(plan.FacebookClientId.ValueStringPointer())
	data.FacebookClientSecret = models.NewOptional(plan.FacebookClientSecret.ValueStringPointer())
	data.FacebookEmailDomains = mist_transform.ListOfStringTerraformToSdk(ctx, plan.FacebookEmailDomains)
	data.FacebookEnabled = plan.FacebookEnabled.ValueBoolPointer()
	data.FacebookExpire = models.NewOptional(models.ToPointer(plan.FacebookExpire.ValueFloat64()))
	data.Forward = plan.Forward.ValueBoolPointer()
	data.ForwardUrl = models.NewOptional(plan.ForwardUrl.ValueStringPointer())
	data.GoogleClientId = models.NewOptional(plan.GoogleClientId.ValueStringPointer())
	data.GoogleClientSecret = models.NewOptional(plan.GoogleClientSecret.ValueStringPointer())
	data.GoogleEmailDomains = mist_transform.ListOfStringTerraformToSdk(ctx, plan.GoogleEmailDomains)
	data.GoogleEnabled = plan.GoogleEnabled.ValueBoolPointer()
	data.GoogleExpire = models.NewOptional(models.ToPointer(plan.GoogleExpire.ValueFloat64()))
	data.GupshupPassword = plan.GupshupPassword.ValueStringPointer()
	data.GupshupUserid = plan.GupshupUserid.ValueStringPointer()
	data.MicrosoftClientId = models.NewOptional(plan.MicrosoftClientId.ValueStringPointer())
	data.MicrosoftClientSecret = models.NewOptional(plan.MicrosoftClientSecret.ValueStringPointer())
	data.MicrosoftEmailDomains = mist_transform.ListOfStringTerraformToSdk(ctx, plan.MicrosoftEmailDomains)
	data.MicrosoftEnabled = plan.MicrosoftEnabled.ValueBoolPointer()
	data.MicrosoftExpire = models.NewOptional(models.ToPointer(plan.MicrosoftExpire.ValueFloat64()))
	data.PassphraseEnabled = plan.PassphraseEnabled.ValueBoolPointer()
	data.PassphraseExpire = models.NewOptional(models.ToPointer(plan.PassphraseExpire.ValueFloat64()))
	data.Password = models.NewOptional(plan.Password.ValueStringPointer())
	data.PredefinedSponsorsEnabled = plan.PredefinedSponsorsEnabled.ValueBoolPointer()
	data.Privacy = plan.Privacy.ValueBoolPointer()
	data.PuzzelPassword = plan.PuzzelPassword.ValueStringPointer()
	data.PuzzelServiceId = plan.PuzzelServiceId.ValueStringPointer()
	data.PuzzelUsername = plan.PuzzelUsername.ValueStringPointer()
	data.SmsEnabled = plan.SmsEnabled.ValueBoolPointer()
	data.SmsExpire = models.NewOptional(models.ToPointer(plan.SmsExpire.ValueFloat64()))
	data.SmsMessageFormat = plan.SmsMessageFormat.ValueStringPointer()
	data.SmsProvider = models.ToPointer(models.WlanPortalSmsProviderEnum(string(plan.SmsProvider.ValueString())))
	data.SponsorAutoApprove = plan.SponsorAutoApprove.ValueBoolPointer()
	data.SponsorEmailDomains = mist_transform.ListOfStringTerraformToSdk(ctx, plan.SponsorEmailDomains)
	data.SponsorEnabled = plan.SponsorEnabled.ValueBoolPointer()
	data.SponsorExpire = models.NewOptional(models.ToPointer(plan.SponsorExpire.ValueFloat64()))
	data.SponsorLinkValidityDuration = plan.SponsorLinkValidityDuration.ValueStringPointer()
	data.SponsorNotifyAll = plan.SponsorNotifyAll.ValueBoolPointer()
	data.SponsorStatusNotify = plan.SponsorStatusNotify.ValueBoolPointer()
	data.Sponsors = models.ToPointer(models.WlanPortalSponsorsContainer.FromMapOfString(sponsors))
	data.SsoDefaultRole = plan.SsoDefaultRole.ValueStringPointer()
	data.SsoForcedRole = plan.SsoForcedRole.ValueStringPointer()
	data.SsoIdpCert = plan.SsoIdpCert.ValueStringPointer()
	data.SsoIdpSignAlgo = (*models.WlanPortalIdpSignAlgoEnum)(plan.SsoIdpSignAlgo.ValueStringPointer())
	data.SsoIdpSsoUrl = plan.SsoIdpSsoUrl.ValueStringPointer()
	data.SsoIssuer = plan.SsoIssuer.ValueStringPointer()
	data.SsoNameidFormat = models.ToPointer(models.WlanPortalSsoNameidFormatEnum(string(plan.SsoNameidFormat.ValueString())))
	data.TelstraClientId = plan.TelstraClientId.ValueStringPointer()
	data.TelstraClientSecret = plan.TelstraClientSecret.ValueStringPointer()
	data.TwilioAuthToken = models.NewOptional(plan.TwilioAuthToken.ValueStringPointer())
	data.TwilioPhoneNumber = models.NewOptional(plan.TwilioPhoneNumber.ValueStringPointer())
	data.TwilioSid = models.NewOptional(plan.TwilioSid.ValueStringPointer())

	return &data
}
