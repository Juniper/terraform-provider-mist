package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ospfAreasNetworksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.OspfAreasNetwork {
	data_map := make(map[string]models.OspfAreasNetwork)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(OspfNetworksValue)

		data_item := models.OspfAreasNetwork{}
		if !item_obj.AuthKeys.IsNull() && !item_obj.AuthKeys.IsUnknown() {
			auth_keys_vm := make(map[string]string)
			for k, v := range item_obj.AuthKeys.Elements() {
				auth_keys_vm[k] = v.String()
			}
			data_item.AuthKeys = auth_keys_vm
		}
		if !item_obj.AuthPassword.IsNull() && !item_obj.AuthPassword.IsUnknown() {
			data_item.AuthPassword = item_obj.AuthPassword.ValueStringPointer()
		}
		if !item_obj.AuthType.IsNull() && !item_obj.AuthType.IsUnknown() {
			data_item.AuthType = (*models.OspfAreaNetworkAuthTypeEnum)(item_obj.AuthType.ValueStringPointer())
		}
		if !item_obj.BfdMinimumInterval.IsNull() && !item_obj.BfdMinimumInterval.IsUnknown() {
			data_item.BfdMinimumInterval = models.ToPointer(int(item_obj.BfdMinimumInterval.ValueInt64()))
		}
		if !item_obj.DeadInterval.IsNull() && !item_obj.DeadInterval.IsUnknown() {
			data_item.DeadInterval = models.ToPointer(int(item_obj.DeadInterval.ValueInt64()))
		}
		if !item_obj.ExportPolicy.IsNull() && !item_obj.ExportPolicy.IsUnknown() {
			data_item.ExportPolicy = item_obj.ExportPolicy.ValueStringPointer()
		}
		if !item_obj.HelloInterval.IsNull() && !item_obj.HelloInterval.IsUnknown() {
			data_item.HelloInterval = models.ToPointer(int(item_obj.HelloInterval.ValueInt64()))
		}
		if !item_obj.ImportPolicy.IsNull() && !item_obj.ImportPolicy.IsUnknown() {
			data_item.ImportPolicy = item_obj.ImportPolicy.ValueStringPointer()
		}
		if !item_obj.InterfaceType.IsNull() && !item_obj.InterfaceType.IsUnknown() {
			data_item.InterfaceType = (*models.OspfAreaNetworkInterfaceTypeEnum)(item_obj.InterfaceType.ValueStringPointer())
		}
		if !item_obj.Metric.IsNull() && !item_obj.Metric.IsUnknown() {
			data_item.Metric = models.NewOptional(models.ToPointer(int(item_obj.Metric.ValueInt64())))
		}
		if !item_obj.NoReadvertiseToOverlay.IsNull() && !item_obj.NoReadvertiseToOverlay.IsUnknown() {
			data_item.NoReadvertiseToOverlay = item_obj.NoReadvertiseToOverlay.ValueBoolPointer()
		}
		if !item_obj.Passive.IsNull() && !item_obj.Passive.IsUnknown() {
			data_item.Passive = item_obj.Passive.ValueBoolPointer()
		}

		data_map[item_name] = data_item
	}
	return data_map
}

func ospfAreasTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.OspfArea {
	data := make(map[string]models.OspfArea)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(OspfAreasValue)

		data_item := models.OspfArea{}
		if !item_obj.IncludeLoopback.IsNull() && !item_obj.IncludeLoopback.IsUnknown() {
			data_item.IncludeLoopback = item_obj.IncludeLoopback.ValueBoolPointer()
		}
		if !item_obj.OspfNetworks.IsNull() && !item_obj.OspfNetworks.IsUnknown() {
			data_item.Networks = ospfAreasNetworksTerraformToSdk(ctx, diags, item_obj.OspfNetworks)
		}
		if !item_obj.OspfAreasType.IsNull() && !item_obj.OspfAreasType.IsUnknown() {
			data_item.Type = (*models.OspfAreaTypeEnum)(item_obj.OspfAreasType.ValueStringPointer())
		}
		data[item_name] = data_item
	}
	return data
}
