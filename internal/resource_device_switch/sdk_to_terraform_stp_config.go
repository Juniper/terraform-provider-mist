package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func stpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SwitchStpConfig) StpConfigValue {

	var vstp_enabled basetypes.BoolValue

	if d.VstpEnabled != nil {
		vstp_enabled = types.BoolValue(*d.VstpEnabled)
	}

	data_map_attr_type := StpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"vstp_enabled": vstp_enabled,
	}
	data, e := NewStpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
