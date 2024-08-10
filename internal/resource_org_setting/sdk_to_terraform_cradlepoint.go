package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func cradlepointSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingCradlepoint) CradlepointValue {

	var cp_api_id basetypes.StringValue
	var cp_api_key basetypes.StringValue
	var ecm_api_id basetypes.StringValue
	var ecm_api_key basetypes.StringValue

	if d.CpApiId != nil {
		cp_api_id = types.StringValue(*d.CpApiId)
	}
	if d.CpApiKey != nil {
		cp_api_key = types.StringValue(*d.CpApiKey)
	}
	if d.EcmApiId != nil {
		ecm_api_id = types.StringValue(*d.EcmApiId)
	}
	if d.EcmApiKey != nil {
		ecm_api_key = types.StringValue(*d.EcmApiKey)
	}

	data_map_attr_type := CradlepointValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"cp_api_id":   cp_api_id,
		"cp_api_key":  cp_api_key,
		"ecm_api_id":  ecm_api_id,
		"ecm_api_key": ecm_api_key,
	}
	data, e := NewCradlepointValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
