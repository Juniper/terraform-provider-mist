package resource_org_deviceprofile_gateway

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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
			data.Disabled = plan.Disabled.ValueBoolPointer()
		}
		if plan.ExportPolicy.ValueStringPointer() != nil {
			data.ExportPolicy = plan.ExportPolicy.ValueStringPointer()
		}
		if plan.HoldTime.ValueInt64Pointer() != nil {
			data.HoldTime = models.ToPointer(int(plan.HoldTime.ValueInt64()))
		}
		if plan.ImportPolicy.ValueStringPointer() != nil {
			data.ImportPolicy = plan.ImportPolicy.ValueStringPointer()
		}
		if plan.MultihopTtl.ValueInt64Pointer() != nil {
			data.MultihopTtl = models.ToPointer(int(plan.MultihopTtl.ValueInt64()))
		}
		if plan.NeighborAs.ValueStringPointer() != nil {
			data.NeighborAs = models.BgpAsContainer.FromString(plan.NeighborAs.ValueString())
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
			data.AuthKey = plan.AuthKey.ValueStringPointer()
		}
		if plan.BfdMinimumInterval.ValueInt64Pointer() != nil {
			data.BfdMinimumInterval = models.NewOptional(models.ToPointer(int(plan.BfdMinimumInterval.ValueInt64())))
		}
		if plan.BfdMultiplier.ValueInt64Pointer() != nil {
			data.BfdMultiplier = models.NewOptional(models.ToPointer(int(plan.BfdMultiplier.ValueInt64())))
		}
		if plan.DisableBfd.ValueBoolPointer() != nil {
			data.DisableBfd = plan.DisableBfd.ValueBoolPointer()
		}
		if plan.Export.ValueStringPointer() != nil {
			data.Export = plan.Export.ValueStringPointer()
		}
		if plan.ExportPolicy.ValueStringPointer() != nil {
			data.ExportPolicy = plan.ExportPolicy.ValueStringPointer()
		}
		if plan.ExtendedV4Nexthop.ValueBoolPointer() != nil {
			data.ExtendedV4Nexthop = plan.ExtendedV4Nexthop.ValueBoolPointer()
		}
		if plan.GracefulRestartTime.ValueInt64Pointer() != nil {
			data.GracefulRestartTime = models.ToPointer(int(plan.GracefulRestartTime.ValueInt64()))
		}
		if plan.HoldTime.ValueInt64Pointer() != nil {
			data.HoldTime = models.ToPointer(int(plan.HoldTime.ValueInt64()))
		}
		if plan.Import.ValueStringPointer() != nil {
			data.Import = plan.Import.ValueStringPointer()
		}
		if plan.ImportPolicy.ValueStringPointer() != nil {
			data.ImportPolicy = plan.ImportPolicy.ValueStringPointer()
		}
		if plan.LocalAs.ValueStringPointer() != nil {
			data.LocalAs = models.ToPointer(models.BgpLocalAsContainer.FromString(plan.LocalAs.ValueString()))
		}
		if plan.NeighborAs.ValueStringPointer() != nil {
			data.NeighborAs = models.ToPointer(models.BgpAsContainer.FromString(plan.NeighborAs.ValueString()))
		}
		if !plan.Neighbors.IsNull() && !plan.Neighbors.IsUnknown() {
			data.Neighbors = bgpConfigNeighborsTerraformToSdk(plan.Neighbors)
		}
		if !plan.Networks.IsNull() && !plan.Networks.IsUnknown() {
			data.Networks = mistutils.ListOfStringTerraformToSdk(plan.Networks)
		}
		if plan.NoPrivateAs.ValueBoolPointer() != nil {
			data.NoPrivateAs = plan.NoPrivateAs.ValueBoolPointer()
		}
		if plan.NoReadvertiseToOverlay.ValueBoolPointer() != nil {
			data.NoReadvertiseToOverlay = plan.NoReadvertiseToOverlay.ValueBoolPointer()
		}
		if plan.BgpConfigType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.BgpConfigTypeEnum(plan.BgpConfigType.ValueString()))
		}
		if plan.TunnelName.ValueStringPointer() != nil {
			data.TunnelName = plan.TunnelName.ValueStringPointer()
		}
		if plan.Via.ValueStringPointer() != nil {
			data.Via = models.BgpConfigViaEnum(plan.Via.ValueString())
		}
		if plan.VpnName.ValueStringPointer() != nil {
			data.VpnName = plan.VpnName.ValueStringPointer()
		}
		if plan.WanName.ValueStringPointer() != nil {
			data.WanName = plan.WanName.ValueStringPointer()
		}

		dataMap[k] = data
	}
	return dataMap
}
