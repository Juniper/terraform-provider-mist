package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ospfConfigAreasTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.OspfConfigArea {
	data_map := make(map[string]models.OspfConfigArea)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(AreasValue)
		data := models.OspfConfigArea{}
		if plan.NoSummary.ValueBoolPointer() != nil {
			data.NoSummary = plan.NoSummary.ValueBoolPointer()
		}
		data_map[k] = data
	}
	return data_map
}

func ospfConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OspfConfigValue) *models.OspfConfig {
	tflog.Debug(ctx, "ospfConfigTerraformToSdk")

	data := models.OspfConfig{}

	if !d.Areas.IsNull() && !d.Areas.IsUnknown() {
		data.Areas = ospfConfigAreasTerraformToSdk(ctx, diags, d.Areas)
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.ReferenceBandwidth.ValueStringPointer() != nil {
		data.ReferenceBandwidth = d.ReferenceBandwidth.ValueStringPointer()
	}
	return &data
}
