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

	var cpApiId basetypes.StringValue
	var cpApiKey basetypes.StringValue
	var ecmApiId basetypes.StringValue
	var ecmApiKey basetypes.StringValue
	var enableLldp basetypes.BoolValue

	if d.CpApiId != nil {
		cpApiId = types.StringValue(*d.CpApiId)
	}
	if d.CpApiKey != nil {
		cpApiKey = types.StringValue(*d.CpApiKey)
	}
	if d.EcmApiId != nil {
		ecmApiId = types.StringValue(*d.EcmApiId)
	}
	if d.EcmApiKey != nil {
		ecmApiKey = types.StringValue(*d.EcmApiKey)
	}
	if d.EnableLldp != nil {
		enableLldp = types.BoolValue(*d.EnableLldp)
	}

	dataMapValue := map[string]attr.Value{
		"cp_api_id":   cpApiId,
		"cp_api_key":  cpApiKey,
		"ecm_api_id":  ecmApiId,
		"ecm_api_key": ecmApiKey,
		"enable_lldp": enableLldp,
	}
	data, e := NewCradlepointValue(CradlepointValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
