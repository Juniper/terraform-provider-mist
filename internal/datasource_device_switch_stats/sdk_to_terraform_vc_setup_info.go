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

	var configType basetypes.StringValue
	var errMissingDevIdFpc basetypes.BoolValue

	if d.ConfigType != nil {
		configType = types.StringValue(*d.ConfigType)
	}
	if d.ErrMissingDevIdFpc != nil {
		errMissingDevIdFpc = types.BoolValue(*d.ErrMissingDevIdFpc)
	}

	dataMapValue := map[string]attr.Value{
		"config_type":            configType,
		"err_missing_dev_id_fpc": errMissingDevIdFpc,
	}
	data, e := basetypes.NewObjectValue(VcSetupInfoValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
