package resource_device_gateway

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func extraRoutesTerraformToSdk(d basetypes.MapValue) map[string]models.GatewayExtraRoute {
	dataMap := make(map[string]models.GatewayExtraRoute)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(ExtraRoutesValue)

		data := models.GatewayExtraRoute{}
		if plan.Via.ValueStringPointer() != nil {
			data.Via = plan.Via.ValueStringPointer()
		}

		dataMap[k] = data
	}
	return dataMap
}

func extraRoutes6TerraformToSdk(d basetypes.MapValue) map[string]models.GatewayExtraRoute {
	dataMap := make(map[string]models.GatewayExtraRoute)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(ExtraRoutes6Value)

		data := models.GatewayExtraRoute{}
		if plan.Via.ValueStringPointer() != nil {
			data.Via = plan.Via.ValueStringPointer()
		}

		dataMap[k] = data
	}
	return dataMap
}
