package datasource_device_switch_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func vcSetupInfoSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsSwitchVcSetupInfo) basetypes.ObjectValue {

	var config_type basetypes.StringValue
	var err_missing_dev_id_fpc basetypes.BoolValue

	if d.ConfigType != nil {
		config_type = types.StringValue(*d.ConfigType)
	}
	if d.ErrMissingDevIdFpc != nil {
		err_missing_dev_id_fpc = types.BoolValue(*d.ErrMissingDevIdFpc)
	}

	data_map_attr_type := VcSetupInfoValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"config_type":            config_type,
		"err_missing_dev_id_fpc": err_missing_dev_id_fpc,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
