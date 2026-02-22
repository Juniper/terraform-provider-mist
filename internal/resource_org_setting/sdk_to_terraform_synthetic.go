package resource_org_setting

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func syntheticTestWanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SynthetictestConfigWanSpeedtest) basetypes.ObjectValue {
	var enabled basetypes.BoolValue
	var timeOfDay basetypes.StringValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.TimeOfDay != nil {
		timeOfDay = types.StringValue(*d.TimeOfDay)
	}

	dataMapValue := map[string]attr.Value{
		"enabled":     enabled,
		"time_of_day": timeOfDay,
	}
	data, e := basetypes.NewObjectValue(WanSpeedtestValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
func syntheticTestCustomProbesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SynthetictestConfigCustomProbe) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var aggressiveness basetypes.StringValue
		var target basetypes.StringValue
		var threshold basetypes.Int64Value
		var probeType basetypes.StringValue

		if d.Aggressiveness != nil {
			aggressiveness = types.StringValue(string(*d.Aggressiveness))
		}
		if d.Target != nil {
			target = types.StringValue(*d.Target)
		}
		if d.Threshold != nil {
			threshold = types.Int64Value(int64(*d.Threshold))
		}
		if d.Type != nil {
			probeType = types.StringValue(string(*d.Type))
		}

		dataMapValue := map[string]attr.Value{
			"aggressiveness": aggressiveness,
			"target":         target,
			"threshold":      threshold,
			"type":           probeType,
		}
		data, e := NewCustomProbesValue(CustomProbesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMapType := CustomProbesValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func syntheticTestLanNetworksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SynthetictestConfigLanNetwork) basetypes.ListValue {
	var dataList []LanNetworksValue
	for _, d := range l {
		var networks = mistutils.ListOfStringSdkToTerraformEmpty()
		var probes = types.ListNull(basetypes.StringType{})

		if d.Networks != nil {
			networks = mistutils.ListOfStringSdkToTerraform(d.Networks)
		}
		if d.Probes != nil {
			probes = mistutils.ListOfStringSdkToTerraform(d.Probes)
		}

		dataMapValue := map[string]attr.Value{
			"networks": networks,
			"probes":   probes,
		}
		data, e := NewLanNetworksValue(LanNetworksValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, LanNetworksValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func syntheticTestVlansSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SynthetictestConfigVlan) basetypes.ListValue {
	var dataList []VlansValue
	for _, d := range l {
		var customTestUrls = mistutils.ListOfStringSdkToTerraformEmpty()
		var disabled basetypes.BoolValue
		var probes = types.ListNull(basetypes.StringType{})
		var vlanIds = types.ListNull(basetypes.StringType{})

		if d.CustomTestUrls != nil {
			customTestUrls = mistutils.ListOfStringSdkToTerraform(d.CustomTestUrls)
		}
		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.Probes != nil {
			probes = mistutils.ListOfStringSdkToTerraform(d.Probes)
		}
		if d.VlanIds != nil {
			var items []attr.Value
			for _, item := range d.VlanIds {
				items = append(items, mistutils.VlanAsString(item))
			}
			vlanIds, _ = types.ListValue(basetypes.StringType{}, items)
		}

		dataMapValue := map[string]attr.Value{
			"custom_test_urls": customTestUrls,
			"disabled":         disabled,
			"probes":           probes,
			"vlan_ids":         vlanIds,
		}
		data, e := NewVlansValue(VlansValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, VlansValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func syntheticTestSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SynthetictestConfig) SyntheticTestValue {

	var aggressiveness basetypes.StringValue
	var customProbes = types.MapNull(CustomProbesValue{}.Type(ctx))
	var disabled basetypes.BoolValue
	var lanNetworks = types.ListNull(LanNetworksValue{}.Type(ctx))
	var vlans = types.ListNull(VlansValue{}.Type(ctx))
	var wanSpeedtest = types.ObjectNull(WanSpeedtestValue{}.AttributeTypes(ctx))

	if d.Aggressiveness != nil {
		aggressiveness = types.StringValue(string(*d.Aggressiveness))
	}
	if d != nil && d.CustomProbes != nil {
		customProbes = syntheticTestCustomProbesSdkToTerraform(ctx, diags, d.CustomProbes)
	}
	if d != nil && d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d != nil && d.LanNetworks != nil {
		lanNetworks = syntheticTestLanNetworksSdkToTerraform(ctx, diags, d.LanNetworks)
	}
	if d != nil && d.Vlans != nil {
		vlans = syntheticTestVlansSdkToTerraform(ctx, diags, d.Vlans)
	}
	if d != nil && d.WanSpeedtest != nil {
		wanSpeedtest = syntheticTestWanSdkToTerraform(ctx, diags, d.WanSpeedtest)
	}

	dataMapValue := map[string]attr.Value{
		"aggressiveness": aggressiveness,
		"custom_probes":  customProbes,
		"disabled":       disabled,
		"lan_networks":   lanNetworks,
		"vlans":          vlans,
		"wan_speedtest":  wanSpeedtest,
	}
	data, e := NewSyntheticTestValue(SyntheticTestValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
