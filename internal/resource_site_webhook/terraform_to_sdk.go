package resource_site_webhook

import (
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(plan *SiteWebhookModel) (models.Webhook, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Webhook{}
	unset := make(map[string]interface{})

	if !plan.AssetfilterIds.IsNull() && !plan.AssetfilterIds.IsUnknown() {
		var items []uuid.UUID
		for _, item := range plan.AssetfilterIds.Elements() {
			var iface interface{} = item
			val := iface.(basetypes.StringValue)
			items = append(items, uuid.MustParse(val.ValueString()))
		}
		data.AssetfilterIds = items
	} else {
		unset["-assetfilter_ids"] = ""
	}

	if !plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
		data.Enabled = plan.Enabled.ValueBoolPointer()
	} else {
		unset["-enabled"] = ""
	}

	if !plan.Headers.IsNull() && !plan.Headers.IsUnknown() {
		items := make(map[string]string)
		for k, v := range plan.Headers.Elements() {
			items[k] = v.String()
		}
		data.Headers = models.NewOptional(models.ToPointer(items))
	} else {
		unset["-headers"] = ""
	}

	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		data.Name = models.NewOptional(models.ToPointer(plan.Name.ValueString()))
	} else {
		unset["-name"] = ""
	}

	if !plan.Oauth2ClientId.IsNull() && !plan.Oauth2ClientId.IsUnknown() {
		data.Oauth2ClientId = plan.Oauth2ClientId.ValueStringPointer()
	} else {
		unset["-oauth2_client_id"] = ""
	}

	if !plan.Oauth2ClientSecret.IsNull() && !plan.Oauth2ClientSecret.IsUnknown() {
		data.Oauth2ClientSecret = plan.Oauth2ClientSecret.ValueStringPointer()
	} else {
		unset["-oauth2_client_secret"] = ""
	}

	if !plan.Oauth2GrantType.IsNull() && !plan.Oauth2GrantType.IsUnknown() {
		data.Oauth2GrantType = models.ToPointer(models.WebhookOauth2GrantTypeEnum(plan.Oauth2GrantType.ValueString()))
	} else {
		unset["-oauth2_grant_type"] = ""
	}

	if !plan.Oauth2Password.IsNull() && !plan.Oauth2Password.IsUnknown() {
		data.Oauth2Password = plan.Oauth2Password.ValueStringPointer()
	} else {
		unset["-oauth2_password"] = ""
	}

	if !plan.Oauth2Scopes.IsNull() && !plan.Oauth2Scopes.IsUnknown() {
		data.Oauth2Scopes = mistutils.ListOfStringTerraformToSdk(plan.Oauth2Scopes)
	} else {
		unset["-oauth2_scopes"] = ""
	}

	if !plan.Oauth2TokenUrl.IsNull() && !plan.Oauth2TokenUrl.IsUnknown() {
		data.Oauth2TokenUrl = plan.Oauth2TokenUrl.ValueStringPointer()
	} else {
		unset["-oauth2_token_url"] = ""
	}

	if !plan.Oauth2Username.IsNull() && !plan.Oauth2Username.IsUnknown() {
		data.Oauth2Username = plan.Oauth2Username.ValueStringPointer()
	} else {
		unset["-oauth2_username"] = ""
	}

	if !plan.Secret.IsNull() && !plan.Secret.IsUnknown() {
		data.Secret = models.NewOptional(models.ToPointer(plan.Secret.ValueString()))
	} else {
		unset["-secret"] = ""
	}

	if !plan.SingleEventPerMessage.IsNull() && !plan.SingleEventPerMessage.IsUnknown() {
		data.SingleEventPerMessage = plan.SingleEventPerMessage.ValueBoolPointer()
	} else {
		unset["-single_event_per_message"] = ""
	}

	if !plan.SplunkToken.IsNull() && !plan.SplunkToken.IsUnknown() {
		data.SplunkToken = models.NewOptional(models.ToPointer(plan.SplunkToken.ValueString()))
	} else {
		unset["-splunk_token"] = ""
	}

	if !plan.Topics.IsNull() && !plan.Topics.IsUnknown() {
		var items []string
		for _, v := range plan.Topics.Elements() {
			var sInterface interface{} = v
			s := sInterface.(basetypes.StringValue)
			items = append(items, s.ValueString())
		}
		data.Topics = items
	} else {
		unset["-topics"] = ""
	}

	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		data.Type = models.ToPointer(models.WebhookTypeEnum(plan.Type.ValueString()))
	} else {
		unset["-type"] = ""
	}

	if !plan.Url.IsNull() && !plan.Url.IsUnknown() {
		data.Url = plan.Url.ValueStringPointer()
	} else {
		unset["-url"] = ""
	}

	if !plan.VerifyCert.IsNull() && !plan.VerifyCert.IsUnknown() {
		data.VerifyCert = plan.VerifyCert.ValueBoolPointer()
	} else {
		unset["-verify_cert"] = ""
	}

	data.AdditionalProperties = unset
	return data, diags
}
