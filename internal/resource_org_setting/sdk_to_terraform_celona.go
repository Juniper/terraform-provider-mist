package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func celonaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingCelona) CelonaValue {

	var apiKey basetypes.StringValue
	var apiPrefix basetypes.StringValue

	if d.ApiKey != nil {
		apiKey = types.StringValue(*d.ApiKey)
	}
	if d.ApiPrefix != nil {
		apiPrefix = types.StringValue(*d.ApiPrefix)
	}

	dataMapValue := map[string]attr.Value{
		"api_key":    apiKey,
		"api_prefix": apiPrefix,
	}
	data, e := NewCelonaValue(CelonaValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
