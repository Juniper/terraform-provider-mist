package resource_device_gateway

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bgpConfigCommunitiesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.BgpConfigCommunity {

	var data_list []models.BgpConfigCommunity
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(CommunitiesValue)
		data := models.BgpConfigCommunity{}
		if plan.Id.ValueStringPointer() != nil {
			data.Id = plan.Id.ValueStringPointer()
		}
		if plan.LocalPreference.ValueInt64Pointer() != nil {
			data.LocalPreference = models.ToPointer(int(plan.LocalPreference.ValueInt64()))
		}
		if plan.VpnName.ValueStringPointer() != nil {
			data.VpnName = plan.VpnName.ValueStringPointer()
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func bgpConfigNeighborsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.BgpConfigNeighbors {
	data_map := make(map[string]models.BgpConfigNeighbors)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NeighborsValue)

		data := models.BgpConfigNeighbors{}
		if plan.Disabled.ValueBoolPointer() != nil {
			data.Disabled = models.ToPointer(plan.Disabled.ValueBool())
		}
		if plan.ExportPolicy.ValueStringPointer() != nil {
			data.ExportPolicy = models.ToPointer(plan.ExportPolicy.ValueString())
		}
		if plan.HoldTime.ValueInt64Pointer() != nil {
			data.HoldTime = models.ToPointer(int(plan.HoldTime.ValueInt64()))
		}
		if plan.ImportPolicy.ValueStringPointer() != nil {
			data.ImportPolicy = models.ToPointer(plan.ImportPolicy.ValueString())
		}
		if plan.MultihopTtl.ValueInt64Pointer() != nil {
			data.MultihopTtl = models.ToPointer(int(plan.MultihopTtl.ValueInt64()))
		}
		if plan.NeighborAs.ValueInt64Pointer() != nil {
			data.NeighborAs = models.ToPointer(int(plan.NeighborAs.ValueInt64()))
		}

		data_map[k] = data
	}
	return data_map
}

func bgpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.BgpConfig {
	data_map := make(map[string]models.BgpConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(BgpConfigValue)

		data := models.BgpConfig{}
		if plan.AuthKey.ValueStringPointer() != nil {
			data.AuthKey = models.ToPointer(plan.AuthKey.ValueString())
		}
		if plan.BfdMinimumInterval.ValueInt64Pointer() != nil {
			data.BfdMinimumInterval = models.NewOptional(models.ToPointer(int(plan.BfdMinimumInterval.ValueInt64())))
		}
		if plan.BfdMultiplier.ValueInt64Pointer() != nil {
			data.BfdMultiplier = models.NewOptional(models.ToPointer(int(plan.BfdMultiplier.ValueInt64())))
		}
		if !plan.Communities.IsNull() && !plan.Communities.IsUnknown() {
			data.Communities = bgpConfigCommunitiesTerraformToSdk(ctx, diags, plan.Communities)
		}
		if plan.DisableBfd.ValueBoolPointer() != nil {
			data.DisableBfd = models.ToPointer(plan.DisableBfd.ValueBool())
		}
		if plan.Export.ValueStringPointer() != nil {
			data.Export = models.ToPointer(plan.Export.ValueString())
		}
		if plan.ExportPolicy.ValueStringPointer() != nil {
			data.ExportPolicy = models.ToPointer(plan.ExportPolicy.ValueString())
		}
		if plan.ExtendedV4Nexthop.ValueBoolPointer() != nil {
			data.ExtendedV4Nexthop = models.ToPointer(plan.ExtendedV4Nexthop.ValueBool())
		}
		if plan.GracefulRestartTime.ValueInt64Pointer() != nil {
			data.GracefulRestartTime = models.ToPointer(int(plan.GracefulRestartTime.ValueInt64()))
		}
		if plan.HoldTime.ValueInt64Pointer() != nil {
			data.HoldTime = models.ToPointer(int(plan.HoldTime.ValueInt64()))
		}
		if plan.Import.ValueStringPointer() != nil {
			data.Import = models.ToPointer(plan.Import.ValueString())
		}
		if plan.ImportPolicy.ValueStringPointer() != nil {
			data.ImportPolicy = models.ToPointer(plan.ImportPolicy.ValueString())
		}
		if plan.NeighborAs.ValueInt64Pointer() != nil {
			data.NeighborAs = models.ToPointer(int(plan.NeighborAs.ValueInt64()))
		}
		if !plan.Neighbors.IsNull() && !plan.Neighbors.IsUnknown() {
			data.Neighbors = bgpConfigNeighborsTerraformToSdk(ctx, diags, plan.Neighbors)
		}
		if !plan.Networks.IsNull() && !plan.Networks.IsUnknown() {
			data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks)
		}
		if plan.NoReadvertiseToOverlay.ValueBoolPointer() != nil {
			data.NoReadvertiseToOverlay = models.ToPointer(plan.NoReadvertiseToOverlay.ValueBool())
		}
		if plan.TunnelName.ValueStringPointer() != nil {
			data.TunnelName = plan.TunnelName.ValueStringPointer()
		}
		if plan.BgpConfigType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.BgpConfigTypeEnum(plan.BgpConfigType.ValueString()))
		}
		if plan.Via.ValueStringPointer() != nil {
			data.Via = models.ToPointer(models.BgpConfigViaEnum(plan.Via.ValueString()))
		}
		if plan.VpnName.ValueStringPointer() != nil {
			data.VpnName = models.ToPointer(plan.VpnName.ValueString())
		}
		if plan.WanName.ValueStringPointer() != nil {
			data.WanName = models.ToPointer(plan.WanName.ValueString())
		}

		data_map[k] = data
	}
	return data_map
}
