package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appQosAppsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.MapValue) map[string]models.WlanAppQosAppsProperties {
	data_map := make(map[string]models.WlanAppQosAppsProperties)
	for k, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(AppsValue)
		data := models.WlanAppQosAppsProperties{}
		data.Dscp = models.ToPointer(int(v_plan.Dscp.ValueInt64()))
		data.DstSubnet = v_plan.DstSubnet.ValueStringPointer()
		data.SrcSubnet = v_plan.SrcSubnet.ValueStringPointer()
		data_map[k] = data
	}
	return data_map
}
func appQosOthersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []models.WlanAppQosOthersItem {
	var data_list []models.WlanAppQosOthersItem
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(OthersValue)
		data := models.WlanAppQosOthersItem{}
		data.Dscp = models.ToPointer(int(v_plan.Dscp.ValueInt64()))
		data.DstSubnet = v_plan.DstSubnet.ValueStringPointer()
		data.PortRanges = v_plan.PortRanges.ValueStringPointer()
		data.Protocol = v_plan.Protocol.ValueStringPointer()
		data.SrcSubnet = v_plan.SrcSubnet.ValueStringPointer()
		data_list = append(data_list, data)
	}
	return data_list
}

func appQosTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AppQosValue) *models.WlanAppQos {

	data := models.WlanAppQos{}

	apps := appQosAppsTerraformToSdk(ctx, diags, plan.Apps)
	data.Apps = apps

	data.Enabled = plan.Enabled.ValueBoolPointer()

	others := appQosOthersTerraformToSdk(ctx, diags, plan.Others)
	data.Others = others

	return &data
}
