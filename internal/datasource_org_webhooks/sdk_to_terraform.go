package datasource_org_webhooks

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.Webhook, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := webhookSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func webhookSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Webhook) OrgWebhooksValue {

	var createdTime types.Float64
	var enabled types.Bool
	var headers = types.MapNull(types.StringType)
	var id types.String
	var modifiedTime types.Float64
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
	var splunkToken types.String
	var topics types.List
	var wtype types.String
	var url types.String
	var verifyCert types.Bool

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
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
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
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
		oauth2Scopes = misttransform.ListOfStringSdkToTerraform(d.Oauth2Scopes)
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
		wtype = types.StringValue(string(*d.Type))
	}
	if d.Url != nil {
		url = types.StringValue(*d.Url)
	}
	if d.VerifyCert != nil {
		verifyCert = types.BoolValue(*d.VerifyCert)
	}

	dataMapValue := map[string]attr.Value{
		"created_time":         createdTime,
		"enabled":              enabled,
		"headers":              headers,
		"id":                   id,
		"modified_time":        modifiedTime,
		"name":                 name,
		"oauth2_client_id":     oauth2ClientId,
		"oauth2_client_secret": oauth2ClientSecret,
		"oauth2_grant_type":    oauth2GrantType,
		"oauth2_password":      oauth2Password,
		"oauth2_scopes":        oauth2Scopes,
		"oauth2_token_url":     oauth2TokenUrl,
		"oauth2_username":      oauth2Username,
		"org_id":               orgId,
		"secret":               secret,
		"splunk_token":         splunkToken,
		"topics":               topics,
		"type":                 wtype,
		"url":                  url,
		"verify_cert":          verifyCert,
	}
	state, e := NewOrgWebhooksValue(OrgWebhooksValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return state

}
