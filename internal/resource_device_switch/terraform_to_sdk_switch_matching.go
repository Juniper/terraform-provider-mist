package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func portConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.JunosPortConfig {

	data := make(map[string]models.JunosPortConfig)
	for k, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(PortConfigValue)
		item_obj := models.JunosPortConfig{}
		item_obj.Usage = plan_obj.Usage.ValueString()
		if plan_obj.AeDisableLacp.ValueBoolPointer() != nil {
			item_obj.AeDisableLacp = models.ToPointer(plan_obj.AeDisableLacp.ValueBool())
		}
		if plan_obj.AeIdx.ValueInt64Pointer() != nil {
			item_obj.AeIdx = models.ToPointer(int(plan_obj.AeIdx.ValueInt64()))
		}
		if plan_obj.AeLacpSlow.ValueBoolPointer() != nil {
			item_obj.AeLacpSlow = models.ToPointer(plan_obj.AeLacpSlow.ValueBool())
		}
		if plan_obj.Aggregated.ValueBoolPointer() != nil {
			item_obj.Aggregated = models.ToPointer(plan_obj.Aggregated.ValueBool())
		}
		if plan_obj.Critical.ValueBoolPointer() != nil {
			item_obj.Critical = models.ToPointer(plan_obj.Critical.ValueBool())
		}
		if plan_obj.Description.ValueStringPointer() != nil {
			item_obj.Description = models.ToPointer(plan_obj.Description.ValueString())
		}
		if plan_obj.DisableAutoneg.ValueBoolPointer() != nil {
			item_obj.DisableAutoneg = models.ToPointer(plan_obj.DisableAutoneg.ValueBool())
		}
		if plan_obj.Duplex.ValueStringPointer() != nil {
			item_obj.Duplex = models.ToPointer(models.JunosPortConfigDuplexEnum(plan_obj.Duplex.ValueString()))
		}
		if plan_obj.DynamicUsage.ValueStringPointer() != nil {
			item_obj.DynamicUsage = models.NewOptional(models.ToPointer(plan_obj.DynamicUsage.ValueString()))
		}
		if plan_obj.Esilag.ValueBoolPointer() != nil {
			item_obj.Esilag = models.ToPointer(plan_obj.Esilag.ValueBool())
		}
		if plan_obj.Mtu.ValueInt64Pointer() != nil {
			item_obj.Mtu = models.ToPointer(int(plan_obj.Mtu.ValueInt64()))
		}
		if plan_obj.NoLocalOverwrite.ValueBoolPointer() != nil {
			item_obj.NoLocalOverwrite = models.ToPointer(plan_obj.NoLocalOverwrite.ValueBool())
		}
		if plan_obj.PoeDisabled.ValueBoolPointer() != nil {
			item_obj.PoeDisabled = models.ToPointer(plan_obj.PoeDisabled.ValueBool())
		}
		if plan_obj.Speed.ValueStringPointer() != nil {
			item_obj.Speed = models.ToPointer(models.JunosPortConfigSpeedEnum(plan_obj.Speed.ValueString()))
		}
		data[k] = item_obj
	}
	return data
}
