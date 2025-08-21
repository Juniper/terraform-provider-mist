package resource_org_setting

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func syntheticTestWanTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SynthetictestConfigWanSpeedtest {
	data := models.SynthetictestConfigWanSpeedtest{}
	if !d.IsNull() && !d.IsUnknown() {
		vd, e := NewWanSpeedtestValue(WanSpeedtestValue{}.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.Enabled = vd.Enabled.ValueBoolPointer()
			data.TimeOfDay = vd.TimeOfDay.ValueStringPointer()
		}
	}
	return &data
}
func syntheticTestCustomProbesTerraformToSdk(m basetypes.MapValue) map[string]models.SynthetictestConfigCustomProbe {
	var dataMap = make(map[string]models.SynthetictestConfigCustomProbe)
	for n, v := range m.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(CustomProbesValue)
		data := models.SynthetictestConfigCustomProbe{}

		if plan.Aggressiveness.ValueStringPointer() != nil {
			data.Aggressiveness = (*models.SynthetictestConfigAggressivenessEnum)(plan.Aggressiveness.ValueStringPointer())
		}
		if plan.Host.ValueStringPointer() != nil {
			data.Host = plan.Host.ValueStringPointer()
		}
		if plan.Port.ValueInt64Pointer() != nil {
			data.Port = models.ToPointer(int(plan.Port.ValueInt64()))
		}
		if plan.Threshold.ValueInt64Pointer() != nil {
			data.Threshold = models.ToPointer(int(plan.Threshold.ValueInt64()))
		}
		if plan.CustomProbesType.ValueStringPointer() != nil {
			data.Type = (*models.SynthetictestConfigCustomProbeTypeEnum)(plan.CustomProbesType.ValueStringPointer())
		}
		if plan.Url.ValueStringPointer() != nil {
			data.Url = plan.Url.ValueStringPointer()
		}

		dataMap[n] = data
	}
	return dataMap
}

func syntheticTestLanNetworksTerraformToSdk(d basetypes.ListValue) []models.SynthetictestConfigLanNetwork {
	var dataList []models.SynthetictestConfigLanNetwork
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(LanNetworksValue)
		data := models.SynthetictestConfigLanNetwork{}

		if !plan.Networks.IsNull() && !plan.Networks.IsUnknown() {
			data.Networks = mistutils.ListOfStringTerraformToSdk(plan.Networks)
		}
		if !plan.Probes.IsNull() && !plan.Probes.IsUnknown() {
			data.Probes = mistutils.ListOfStringTerraformToSdk(plan.Probes)
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func syntheticTestVlansTerraformToSdk(d basetypes.ListValue) []models.SynthetictestConfigVlan {
	var dataList []models.SynthetictestConfigVlan
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(VlansValue)
		data := models.SynthetictestConfigVlan{}

		if !plan.CustomTestUrls.IsNull() && !plan.CustomTestUrls.IsUnknown() {
			data.CustomTestUrls = mistutils.ListOfStringTerraformToSdk(plan.CustomTestUrls)
		}

		if plan.Disabled.ValueBoolPointer() != nil {
			data.Disabled = plan.Disabled.ValueBoolPointer()
		}

		if !plan.Probes.IsNull() && !plan.Probes.IsUnknown() {
			data.Probes = mistutils.ListOfStringTerraformToSdk(plan.Probes)
		}

		if !plan.VlanIds.IsNull() && !plan.VlanIds.IsUnknown() {
			var items []models.VlanIdWithVariable
			for _, item := range plan.VlanIds.Elements() {
				var itemInterface interface{} = item
				i := itemInterface.(basetypes.StringValue)
				v := models.VlanIdWithVariableContainer.FromString(i.ValueString())
				items = append(items, v)
			}
			data.VlanIds = items
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func syntheticTestTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SyntheticTestValue) *models.SynthetictestConfig {
	data := models.SynthetictestConfig{}

	if d.Aggressiveness.ValueStringPointer() != nil {
		data.Aggressiveness = (*models.SynthetictestConfigAggressivenessEnum)(d.Aggressiveness.ValueStringPointer())
	}

	if !d.CustomProbes.IsNull() && !d.CustomProbes.IsUnknown() {
		data.CustomProbes = syntheticTestCustomProbesTerraformToSdk(d.CustomProbes)
	}

	if !d.LanNetworks.IsNull() && !d.LanNetworks.IsUnknown() {
		data.LanNetworks = syntheticTestLanNetworksTerraformToSdk(d.LanNetworks)
	}

	if d.Disabled.ValueBoolPointer() != nil {
		data.Disabled = d.Disabled.ValueBoolPointer()
	}

	if !d.Vlans.IsNull() && !d.Vlans.IsUnknown() {
		data.Vlans = syntheticTestVlansTerraformToSdk(d.Vlans)
	}

	if !d.WanSpeedtest.IsNull() && !d.WanSpeedtest.IsUnknown() {
		data.WanSpeedtest = syntheticTestWanTerraformToSdk(ctx, diags, d.WanSpeedtest)
	}

	return &data
}
