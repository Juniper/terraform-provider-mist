package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func pathPreferencePathsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.GatewayPathPreferencesPath {
	tflog.Debug(ctx, "pathPreferencePathsTerraformToSdk")
	var data_list []models.GatewayPathPreferencesPath
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PathsValue)
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
			data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks)
		}
		if !plan.TargetIps.IsNull() && !plan.TargetIps.IsUnknown() {
			data.TargetIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.TargetIps)
		}
		if plan.PathsType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.GatewayPathTypeEnum(plan.PathsType.ValueString()))
		}
		if plan.WanName.ValueStringPointer() != nil {
			data.WanName = models.ToPointer(plan.WanName.ValueString())
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func pathPreferencesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayPathPreferences {
	tflog.Debug(ctx, "pathPreferencesTerraformToSdk")
	data_map := make(map[string]models.GatewayPathPreferences)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PathPreferencesValue)

		data := models.GatewayPathPreferences{}
		paths := pathPreferencePathsTerraformToSdk(ctx, diags, plan.Paths)
		if !plan.Paths.IsNull() && !plan.Paths.IsUnknown() {
			data.Paths = paths
		}
		if plan.Strategy.ValueStringPointer() != nil {
			data.Strategy = models.ToPointer(models.GatewayPathStrategyEnum(plan.Strategy.ValueString()))
		}
		data_map[k] = data
	}
	return data_map
}
