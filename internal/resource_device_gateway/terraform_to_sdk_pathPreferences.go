package resource_device_gateway

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func pathPreferencePathsTerraformToSdk(d basetypes.ListValue) []models.GatewayPathPreferencesPath {
	var dataList []models.GatewayPathPreferencesPath
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(PathsValue)
		data := models.GatewayPathPreferencesPath{}
		if plan.Cost.ValueInt64Pointer() != nil {
			data.Cost = models.ToPointer(int(plan.Cost.ValueInt64()))
		}
		if plan.Disabled.ValueBoolPointer() != nil {
			data.Disabled = models.ToPointer(plan.Disabled.ValueBool())
		}
		if plan.GatewayIp.ValueStringPointer() != nil {
			data.GatewayIp = models.ToPointer(plan.GatewayIp.ValueString())
		}
		if plan.InternetAccess.ValueBoolPointer() != nil {
			data.InternetAccess = models.ToPointer(plan.InternetAccess.ValueBool())
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = models.ToPointer(plan.Name.ValueString())
		}
		if !plan.Networks.IsNull() && !plan.Networks.IsUnknown() {
			data.Networks = misttransform.ListOfStringTerraformToSdk(plan.Networks)
		}
		if !plan.TargetIps.IsNull() && !plan.TargetIps.IsUnknown() {
			data.TargetIps = misttransform.ListOfStringTerraformToSdk(plan.TargetIps)
		}
		if plan.PathsType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.GatewayPathTypeEnum(plan.PathsType.ValueString()))
		}
		if plan.WanName.ValueStringPointer() != nil {
			data.WanName = models.ToPointer(plan.WanName.ValueString())
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func pathPreferencesTerraformToSdk(d basetypes.MapValue) map[string]models.GatewayPathPreferences {
	dataMap := make(map[string]models.GatewayPathPreferences)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(PathPreferencesValue)

		data := models.GatewayPathPreferences{}
		paths := pathPreferencePathsTerraformToSdk(plan.Paths)
		if !plan.Paths.IsNull() && !plan.Paths.IsUnknown() {
			data.Paths = paths
		}
		if plan.Strategy.ValueStringPointer() != nil {
			data.Strategy = models.ToPointer(models.GatewayPathStrategyEnum(plan.Strategy.ValueString()))
		}
		dataMap[k] = data
	}
	return dataMap
}
