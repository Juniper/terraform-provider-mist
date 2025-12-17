package resource_device_switch

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bgpConfigNeighborsTerraformToSdk(d basetypes.MapValue) map[string]models.SwitchBgpConfigNeighbor {
	dataMap := make(map[string]models.SwitchBgpConfigNeighbor)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(NeighborsValue)

		data := models.SwitchBgpConfigNeighbor{}
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
			data.NeighborAs = models.SwitchBgpConfigNeighborNeighborAsContainer.FromString(plan.NeighborAs.ValueString())
		}

		dataMap[k] = data
	}
	return dataMap
}

func bgpConfigTerraformToSdk(d basetypes.MapValue) map[string]models.SwitchBgpConfig {
	dataMap := make(map[string]models.SwitchBgpConfig)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(BgpConfigValue)

		data := models.SwitchBgpConfig{}
		if plan.AuthKey.ValueStringPointer() != nil {
			data.AuthKey = plan.AuthKey.ValueStringPointer()
		}
		if plan.BfdMinimumInterval.ValueInt64Pointer() != nil {
			data.BfdMinimumInterval = models.ToPointer(int(plan.BfdMinimumInterval.ValueInt64()))
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
		if plan.LocalAs.ValueStringPointer() != nil {
			data.LocalAs = models.BgpAsContainer.FromString(plan.LocalAs.ValueString())
		}
		if !plan.Neighbors.IsNull() && !plan.Neighbors.IsUnknown() {
			data.Neighbors = bgpConfigNeighborsTerraformToSdk(plan.Neighbors)
		}
		if !plan.Networks.IsNull() && !plan.Networks.IsUnknown() {
			data.Networks = mistutils.ListOfStringTerraformToSdk(plan.Networks)
		}
		if plan.BgpConfigType.ValueStringPointer() != nil {
			data.Type = models.SwitchBgpConfigTypeEnum(plan.BgpConfigType.ValueString())
		}

		dataMap[k] = data
	}
	return dataMap
}
