package resource_site_webhook

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, d *models.Webhook) (SiteWebhookModel, diag.Diagnostics) {
	var state SiteWebhookModel
	var diags diag.Diagnostics

	var enabled types.Bool
	var headers types.Map = types.MapNull(types.StringType)
	var id types.String
	var name types.String
	var oauth2_client_id types.String
	var oauth2_client_secret types.String
	var oauth2_grant_type types.String
	var oauth2_password types.String
	var oauth2_scopes types.List = types.ListNull(types.StringType)
	var oauth2_token_url types.String
	var oauth2_username types.String
	var org_id types.String
	var secret types.String
	var site_id types.String
	var splunk_token types.String
	var topics types.List
	var wtype types.String
	var url types.String
	var verify_cert types.Bool

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
		oauth2_client_id = types.StringValue(*d.Oauth2ClientId)
	}
	if d.Oauth2ClientSecret != nil {
		oauth2_client_secret = types.StringValue(*d.Oauth2ClientSecret)
	}
	if d.Oauth2GrantType != nil {
		oauth2_grant_type = types.StringValue(string(*d.Oauth2GrantType))
	}
	if d.Oauth2Password != nil {
		oauth2_password = types.StringValue(*d.Oauth2Password)
	}
	if d.Oauth2Scopes != nil {
		oauth2_scopes = mist_transform.ListOfStringSdkToTerraform(ctx, d.Oauth2Scopes)
	}
	if d.Oauth2TokenUrl != nil {
		oauth2_token_url = types.StringValue(*d.Oauth2TokenUrl)
	}
	if d.Oauth2Username != nil {
		oauth2_username = types.StringValue(*d.Oauth2Username)
	}
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.Secret.Value() != nil {
		secret = types.StringValue(*d.Secret.Value())
	}
	if d.SiteId != nil {
		site_id = types.StringValue(d.SiteId.String())
	}
	if d.SplunkToken.Value() != nil {
		splunk_token = types.StringValue(*d.SplunkToken.Value())
	}
	if d.Topics != nil {
		var items []attr.Value
		var items_type attr.Type = basetypes.StringType{}
		for _, item := range d.Topics {
			items = append(items, types.StringValue(string(item)))
		}
		list, _ := types.ListValue(items_type, items)
		topics = list
	}
	if d.Type != nil {
		wtype = types.StringValue(string(*d.Type))
	}
	if d.Url != nil {
		url = types.StringValue(*d.Url)
	}
	if d.VerifyCert != nil {
		verify_cert = types.BoolValue(*d.VerifyCert)
	}

	state.Enabled = enabled
	state.Headers = headers
	state.Id = id
	state.Name = name
	state.Oauth2ClientId = oauth2_client_id
	state.Oauth2ClientSecret = oauth2_client_secret
	state.Oauth2GrantType = oauth2_grant_type
	state.Oauth2Password = oauth2_password
	state.Oauth2Scopes = oauth2_scopes
	state.Oauth2TokenUrl = oauth2_token_url
	state.Oauth2Username = oauth2_username
	state.OrgId = org_id
	state.Secret = secret
	state.SiteId = site_id
	state.SplunkToken = splunk_token
	state.OrgId = org_id
	state.Topics = topics
	state.Type = wtype
	state.Url = url
	state.VerifyCert = verify_cert

	return state, diags

}
