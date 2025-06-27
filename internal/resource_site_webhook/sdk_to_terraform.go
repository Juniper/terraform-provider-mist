package resource_site_webhook

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, d *models.Webhook) (SiteWebhookModel, diag.Diagnostics) {
	var state SiteWebhookModel
	var diags diag.Diagnostics

	var enabled types.Bool
	var headers = types.MapNull(types.StringType)
	var id types.String
	var name types.String
	var oauth2ClientId types.String
	var oauth2ClientSecret types.String
	var oauth2GrantType types.String
	var oauth2Password types.String
	var oauth2Scopes = types.ListNull(types.StringType)
	var oauth2TokenUrl types.String
	var oauth2Username types.String
	var orgId types.String
	var secret types.String
	var singleEventPerMessage types.Bool
	var siteId types.String
	var splunkToken types.String
	var topics types.List
	var wType types.String
	var url types.String
	var verifyCert types.Bool

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Headers.Value() != nil {
		tmp, e := types.MapValueFrom(ctx, types.StringType, *d.Headers.Value())
		if e != nil {
			diags.Append(e...)
		} else {
			headers = tmp
		}
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}

	if d.Name.Value() != nil {
		name = types.StringValue(*d.Name.Value())
	}
	if d.Oauth2ClientId != nil {
		oauth2ClientId = types.StringValue(*d.Oauth2ClientId)
	}
	if d.Oauth2ClientSecret != nil {
		oauth2ClientSecret = types.StringValue(*d.Oauth2ClientSecret)
	}
	if d.Oauth2GrantType != nil {
		oauth2GrantType = types.StringValue(string(*d.Oauth2GrantType))
	}
	if d.Oauth2Password != nil {
		oauth2Password = types.StringValue(*d.Oauth2Password)
	}
	if d.Oauth2Scopes != nil {
		oauth2Scopes = mistutils.ListOfStringSdkToTerraform(d.Oauth2Scopes)
	}
	if d.Oauth2TokenUrl != nil {
		oauth2TokenUrl = types.StringValue(*d.Oauth2TokenUrl)
	}
	if d.Oauth2Username != nil {
		oauth2Username = types.StringValue(*d.Oauth2Username)
	}
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Secret.Value() != nil {
		secret = types.StringValue(*d.Secret.Value())
	}
	if d.SingleEventPerMessage != nil {
		singleEventPerMessage = types.BoolValue(*d.SingleEventPerMessage)
	}
	if d.SiteId != nil {
		siteId = types.StringValue(d.SiteId.String())
	}
	if d.SplunkToken.Value() != nil {
		splunkToken = types.StringValue(*d.SplunkToken.Value())
	}
	if d.Topics != nil {
		var items []attr.Value
		var itemsType attr.Type = basetypes.StringType{}
		for _, item := range d.Topics {
			items = append(items, types.StringValue(item))
		}
		list, _ := types.ListValue(itemsType, items)
		topics = list
	}
	if d.Type != nil {
		wType = types.StringValue(string(*d.Type))
	}
	if d.Url != nil {
		url = types.StringValue(*d.Url)
	}
	if d.VerifyCert != nil {
		verifyCert = types.BoolValue(*d.VerifyCert)
	}

	state.Enabled = enabled
	state.Headers = headers
	state.Id = id
	state.Name = name
	state.Oauth2ClientId = oauth2ClientId
	state.Oauth2ClientSecret = oauth2ClientSecret
	state.Oauth2GrantType = oauth2GrantType
	state.Oauth2Password = oauth2Password
	state.Oauth2Scopes = oauth2Scopes
	state.Oauth2TokenUrl = oauth2TokenUrl
	state.Oauth2Username = oauth2Username
	state.OrgId = orgId
	state.Secret = secret
	state.SingleEventPerMessage = singleEventPerMessage
	state.SiteId = siteId
	state.SplunkToken = splunkToken
	state.OrgId = orgId
	state.Topics = topics
	state.Type = wType
	state.Url = url
	state.VerifyCert = verifyCert

	return state, diags

}
