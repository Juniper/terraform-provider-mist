package resource_org_deviceprofile_gateway

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bgpConfigNeighborsTerraformToSdk(d basetypes.MapValue) map[string]models.BgpConfigNeighbors {
	dataMap := make(map[string]models.BgpConfigNeighbors)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(NeighborsValue)

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

		dataMap[k] = data
	}
	return dataMap
}

func bgpConfigTerraformToSdk(d basetypes.MapValue) map[string]models.BgpConfig {
	dataMap := make(map[string]models.BgpConfig)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(BgpConfigValue)

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
		if plan.LocalAs.ValueInt64Pointer() != nil {
			data.LocalAs = models.ToPointer(int(plan.LocalAs.ValueInt64()))
		}
		if plan.NeighborAs.ValueInt64Pointer() != nil {
			data.NeighborAs = models.ToPointer(int(plan.NeighborAs.ValueInt64()))
		}
		if !plan.Neighbors.IsNull() && !plan.Neighbors.IsUnknown() {
			data.Neighbors = bgpConfigNeighborsTerraformToSdk(plan.Neighbors)
		}
		if !plan.Networks.IsNull() && !plan.Networks.IsUnknown() {
			data.Networks = misttransform.ListOfStringTerraformToSdk(plan.Networks)
		}
		if plan.NoReadvertiseToOverlay.ValueBoolPointer() != nil {
			data.NoReadvertiseToOverlay = models.ToPointer(plan.NoReadvertiseToOverlay.ValueBool())
		}
		if plan.BgpConfigType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.BgpConfigTypeEnum(plan.BgpConfigType.ValueString()))
		}
		if plan.TunnelName.ValueStringPointer() != nil {
			data.TunnelName = plan.TunnelName.ValueStringPointer()
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

		dataMap[k] = data
	}
	return dataMap
}
