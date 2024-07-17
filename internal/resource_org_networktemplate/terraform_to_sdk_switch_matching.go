package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func switchMatchingRulesPortMirroringTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.SwitchPortMirroringProperty {

	data := make(map[string]models.SwitchPortMirroringProperty)
	for k, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(PortMirroringValue)
		item_obj := models.SwitchPortMirroringProperty{}
		if !plan_obj.InputNetworksIngress.IsNull() && !plan_obj.InputNetworksIngress.IsUnknown() {
			item_obj.InputNetworksIngress = mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputNetworksIngress)
		}
		if !plan_obj.InputPortIdsEgress.IsNull() && !plan_obj.InputPortIdsEgress.IsUnknown() {
			item_obj.InputPortIdsEgress = mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputPortIdsEgress)
		}
		if !plan_obj.InputPortIdsIngress.IsNull() && !plan_obj.InputPortIdsIngress.IsUnknown() {
			item_obj.InputPortIdsIngress = mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.InputPortIdsIngress)
		}
		if plan_obj.OutputPortId.ValueStringPointer() != nil {
			item_obj.OutputPortId = models.ToPointer(plan_obj.OutputPortId.ValueString())
		}
		data[k] = item_obj
	}
	return data
}
func switchMatchingRulesPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.JunosPortConfig {

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
func switchMatchingRulesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SwitchMatchingRule {

	var data []models.SwitchMatchingRule
	for _, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(MatchingRulesValue)
		item_obj := models.SwitchMatchingRule{}

		if !plan_obj.AdditionalConfigCmds.IsNull() && !plan_obj.AdditionalConfigCmds.IsUnknown() {
			item_obj.AdditionalConfigCmds = mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.AdditionalConfigCmds)
		}
		if plan_obj.MatchRole.ValueStringPointer() != nil {
			item_obj.MatchRole = models.ToPointer(plan_obj.MatchRole.ValueString())
		}
		if plan_obj.Name.ValueStringPointer() != nil {
			item_obj.Name = models.ToPointer(plan_obj.Name.ValueString())
		}
		if !plan_obj.PortConfig.IsNull() && !plan_obj.PortConfig.IsUnknown() {
			item_obj.PortConfig = switchMatchingRulesPortConfigTerraformToSdk(ctx, diags, plan_obj.PortConfig)
		}
		if !plan_obj.PortMirroring.IsNull() && !plan_obj.PortMirroring.IsUnknown() {
			item_obj.PortMirroring = switchMatchingRulesPortMirroringTerraformToSdk(ctx, diags, plan_obj.PortMirroring)
		}

		match := map[string]interface{}{}
		if plan_obj.MatchType.ValueStringPointer() != nil {
			match_type := plan_obj.MatchType.ValueString()
			match_value := plan_obj.MatchValue.ValueString()
			match[match_type] = match_value
		}
		item_obj.AdditionalProperties = match

		data = append(data, item_obj)
	}
	return data
}

func switchMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SwitchMatchingValue) *models.SwitchMatching {

	data := models.SwitchMatching{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		if d.Enable.ValueBoolPointer() != nil {
			data.Enable = models.ToPointer(d.Enable.ValueBool())
		}
		if !d.MatchingRules.IsNull() && !d.MatchingRules.IsUnknown() {
			data.Rules = switchMatchingRulesTerraformToSdk(ctx, diags, d.MatchingRules)
		}

		return &data
	}

}
