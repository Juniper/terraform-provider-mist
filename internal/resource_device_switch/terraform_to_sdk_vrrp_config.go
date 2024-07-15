package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func vrrpGroupsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.VrrpConfigGroup {
	data_map := make(map[string]models.VrrpConfigGroup)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(GroupsValue)
		data := models.VrrpConfigGroup{}

		if plan.Priority.ValueInt64Pointer() != nil {
			data.Priority = models.ToPointer(int(plan.Priority.ValueInt64()))
		}

		data_map[k] = data
	}
	return data_map
}

func vrrpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VrrpConfigValue) *models.VrrpConfig {
	tflog.Debug(ctx, "vrrpTerraformToSdk")

	data := models.VrrpConfig{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.Groups.IsNull() && !d.Groups.IsUnknown() {
		data.Groups = vrrpGroupsTerraformToSdk(ctx, diags, d.Groups)
	}

	return &data
}
