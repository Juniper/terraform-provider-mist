package datasource_site_webhooks

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.Webhook) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := webhookSdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(SiteWebhooksValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func webhookSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.Webhook) SiteWebhooksValue {
	var state SiteWebhooksValue

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

	data_map_attr_type := SiteWebhooksValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":              enabled,
		"headers":              headers,
		"id":                   id,
		"name":                 name,
		"oauth2_client_id":     oauth2_client_id,
		"oauth2_client_secret": oauth2_client_secret,
		"oauth2_grant_type":    oauth2_grant_type,
		"oauth2_password":      oauth2_password,
		"oauth2_scopes":        oauth2_scopes,
		"oauth2_token_url":     oauth2_token_url,
		"oauth2_username":      oauth2_username,
		"org_id":               org_id,
		"secret":               secret,
		"site_id":              site_id,
		"splunk_token":         splunk_token,
		"topics":               topics,
		"type":                 wtype,
		"url":                  url,
		"verify_cert":          verify_cert,
	}
	state, e := NewSiteWebhooksValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return state

}
