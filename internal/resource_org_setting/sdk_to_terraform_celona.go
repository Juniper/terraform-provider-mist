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

	var api_key basetypes.StringValue
	var api_prefix basetypes.StringValue

	if d.ApiKey != nil {
		api_key = types.StringValue(*d.ApiKey)
	}
	if d.ApiPrefix != nil {
		api_prefix = types.StringValue(*d.ApiPrefix)
	}

	data_map_attr_type := CelonaValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"api_key":    api_key,
		"api_prefix": api_prefix,
	}
	data, e := NewCelonaValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
