package resource_org_networktemplate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

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
func switchMatchingRulesIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SwitchMatchingRuleIpConfig {
	data := models.SwitchMatchingRuleIpConfig{}
	if !d.IsNull() && !d.IsUnknown() {
		item, e := NewIpConfigValue(IpConfigValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		if e == nil {
			if item.IpConfigType.ValueStringPointer() != nil {
				data.Type = (*models.IpTypeEnum)(item.IpConfigType.ValueStringPointer())
			}
			if item.Network.ValueStringPointer() != nil {
				data.Network = item.Network.ValueStringPointer()
			}
		}
	}
	return &data
}
func switchMatchingRulesOobIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SwitchMatchingRuleOobIpConfig {
	data := models.SwitchMatchingRuleOobIpConfig{}
	if !d.IsNull() && !d.IsUnknown() {
		item, e := NewOobIpConfigValue(OobIpConfigValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		if e == nil {
			if item.OobIpConfigType.ValueStringPointer() != nil {
				data.Type = (*models.IpTypeEnum)(item.OobIpConfigType.ValueStringPointer())
			}
			if item.UseMgmtVrf.ValueBoolPointer() != nil {
				data.UseMgmtVrf = item.UseMgmtVrf.ValueBoolPointer()
			}
			if item.UseMgmtVrfForHostOut.ValueBoolPointer() != nil {
				data.UseMgmtVrfForHostOut = item.UseMgmtVrfForHostOut.ValueBoolPointer()
			}
		}
	}
	return &data
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
		if plan_obj.Name.ValueStringPointer() != nil {
			item_obj.Name = models.ToPointer(plan_obj.Name.ValueString())
		}
		if !plan_obj.PortConfig.IsNull() && !plan_obj.PortConfig.IsUnknown() {
			item_obj.PortConfig = switchMatchingRulesPortConfigTerraformToSdk(ctx, diags, plan_obj.PortConfig)
		}
		if !plan_obj.PortMirroring.IsNull() && !plan_obj.PortMirroring.IsUnknown() {
			item_obj.PortMirroring = portMirroringTerraformToSdk(ctx, diags, plan_obj.PortMirroring)
		}
		if !plan_obj.IpConfig.IsNull() && !plan_obj.IpConfig.IsUnknown() {
			item_obj.IpConfig = switchMatchingRulesIpConfigTerraformToSdk(ctx, diags, plan_obj.IpConfig)
		}
		if !plan_obj.OobIpConfig.IsNull() && !plan_obj.OobIpConfig.IsUnknown() {
			item_obj.OobIpConfig = switchMatchingRulesOobIpConfigTerraformToSdk(ctx, diags, plan_obj.OobIpConfig)
		}

		match := map[string]interface{}{}
		if plan_obj.MatchType.ValueStringPointer() != nil && plan_obj.MatchType.ValueString() != "" {
			match_type := plan_obj.MatchType.ValueString()
			match[match_type] = plan_obj.MatchValue.ValueString()
		}

		if plan_obj.MatchModel.ValueStringPointer() != nil && plan_obj.MatchModel.ValueString() != "" {
			match_type := fmt.Sprintf(
				"match_model[0:%d]",
				len(plan_obj.MatchModel.ValueString()),
			)
			match[match_type] = plan_obj.MatchModel.ValueString()
		}

		if plan_obj.MatchName.ValueStringPointer() != nil && plan_obj.MatchName.ValueString() != "" {
			offset := 0
			if plan_obj.MatchNameOffset.ValueInt64Pointer() != nil {
				offset = int(plan_obj.MatchNameOffset.ValueInt64())
			}
			match_type := fmt.Sprintf(
				"match_name[%d:%d]",
				offset,
				offset+len(plan_obj.MatchName.ValueString()),
			)
			match[match_type] = plan_obj.MatchName.ValueString()
		}

		if plan_obj.MatchRole.ValueStringPointer() != nil {
			match["match_role"] = plan_obj.MatchRole.ValueString()
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
