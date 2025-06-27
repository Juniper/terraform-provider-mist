package resource_site_networktemplate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func switchMatchingRulesPortConfigTerraformToSdk(d basetypes.MapValue) map[string]models.JunosPortConfig {

	data := make(map[string]models.JunosPortConfig)
	for k, v := range d.Elements() {
		var planInterface interface{} = v
		planObj := planInterface.(PortConfigValue)
		itemObj := models.JunosPortConfig{}
		itemObj.Usage = planObj.Usage.ValueString()
		if planObj.AeDisableLacp.ValueBoolPointer() != nil {
			itemObj.AeDisableLacp = models.ToPointer(planObj.AeDisableLacp.ValueBool())
		}
		if planObj.AeIdx.ValueInt64Pointer() != nil {
			itemObj.AeIdx = models.ToPointer(int(planObj.AeIdx.ValueInt64()))
		}
		if planObj.AeLacpSlow.ValueBoolPointer() != nil {
			itemObj.AeLacpSlow = models.ToPointer(planObj.AeLacpSlow.ValueBool())
		}
		if planObj.Aggregated.ValueBoolPointer() != nil {
			itemObj.Aggregated = models.ToPointer(planObj.Aggregated.ValueBool())
		}
		if planObj.Critical.ValueBoolPointer() != nil {
			itemObj.Critical = models.ToPointer(planObj.Critical.ValueBool())
		}
		if planObj.Description.ValueStringPointer() != nil {
			itemObj.Description = models.ToPointer(planObj.Description.ValueString())
		}
		if planObj.DisableAutoneg.ValueBoolPointer() != nil {
			itemObj.DisableAutoneg = models.ToPointer(planObj.DisableAutoneg.ValueBool())
		}
		if planObj.Duplex.ValueStringPointer() != nil {
			itemObj.Duplex = models.ToPointer(models.JunosPortConfigDuplexEnum(planObj.Duplex.ValueString()))
		}
		if planObj.DynamicUsage.ValueStringPointer() != nil {
			itemObj.DynamicUsage = models.NewOptional(models.ToPointer(planObj.DynamicUsage.ValueString()))
		}
		if planObj.Esilag.ValueBoolPointer() != nil {
			itemObj.Esilag = models.ToPointer(planObj.Esilag.ValueBool())
		}
		if planObj.Mtu.ValueInt64Pointer() != nil {
			itemObj.Mtu = models.ToPointer(int(planObj.Mtu.ValueInt64()))
		}
		if planObj.NoLocalOverwrite.ValueBoolPointer() != nil {
			itemObj.NoLocalOverwrite = models.ToPointer(planObj.NoLocalOverwrite.ValueBool())
		}
		if planObj.PoeDisabled.ValueBoolPointer() != nil {
			itemObj.PoeDisabled = models.ToPointer(planObj.PoeDisabled.ValueBool())
		}
		if planObj.Speed.ValueStringPointer() != nil {
			itemObj.Speed = models.ToPointer(models.JunosPortConfigSpeedEnum(planObj.Speed.ValueString()))
		}
		data[k] = itemObj
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
		var planInterface interface{} = v
		planObj := planInterface.(MatchingRulesValue)
		itemObj := models.SwitchMatchingRule{}

		if !planObj.AdditionalConfigCmds.IsNull() && !planObj.AdditionalConfigCmds.IsUnknown() {
			itemObj.AdditionalConfigCmds = mistutils.ListOfStringTerraformToSdk(planObj.AdditionalConfigCmds)
		}
		if planObj.Name.ValueStringPointer() != nil {
			itemObj.Name = models.ToPointer(planObj.Name.ValueString())
		}
		if !planObj.PortConfig.IsNull() && !planObj.PortConfig.IsUnknown() {
			itemObj.PortConfig = switchMatchingRulesPortConfigTerraformToSdk(planObj.PortConfig)
		}
		if !planObj.PortMirroring.IsNull() && !planObj.PortMirroring.IsUnknown() {
			itemObj.PortMirroring = portMirroringTerraformToSdk(planObj.PortMirroring)
		}
		if !planObj.IpConfig.IsNull() && !planObj.IpConfig.IsUnknown() {
			itemObj.IpConfig = switchMatchingRulesIpConfigTerraformToSdk(ctx, diags, planObj.IpConfig)
		}
		if !planObj.OobIpConfig.IsNull() && !planObj.OobIpConfig.IsUnknown() {
			itemObj.OobIpConfig = switchMatchingRulesOobIpConfigTerraformToSdk(ctx, diags, planObj.OobIpConfig)
		}

		match := make(map[string]string)
		if planObj.MatchType.ValueStringPointer() != nil && planObj.MatchType.ValueString() != "" {
			matchType := planObj.MatchType.ValueString()
			match[matchType] = planObj.MatchValue.ValueString()
		}

		if planObj.MatchModel.ValueStringPointer() != nil && planObj.MatchModel.ValueString() != "" {
			matchType := fmt.Sprintf(
				"match_model[0:%d]",
				len(planObj.MatchModel.ValueString()),
			)
			match[matchType] = planObj.MatchModel.ValueString()
		}

		if planObj.MatchName.ValueStringPointer() != nil && planObj.MatchName.ValueString() != "" {
			offset := 0
			if planObj.MatchNameOffset.ValueInt64Pointer() != nil {
				offset = int(planObj.MatchNameOffset.ValueInt64())
			}
			matchType := fmt.Sprintf(
				"match_name[%d:%d]",
				offset,
				offset+len(planObj.MatchName.ValueString()),
			)
			match[matchType] = planObj.MatchName.ValueString()
		}

		if planObj.MatchRole.ValueStringPointer() != nil {
			match["match_role"] = planObj.MatchRole.ValueString()
		}

		itemObj.AdditionalProperties = match

		data = append(data, itemObj)
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
