package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func extraRoutesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayExtraRoute {
	data_map := make(map[string]models.GatewayExtraRoute)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ExtraRoutesValue)

		data := models.GatewayExtraRoute{}
		if plan.Via.ValueStringPointer() != nil {
			data.Via = plan.Via.ValueStringPointer()
		}

		data_map[k] = data
	}
	return data_map
}

func extraRoutes6TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayExtraRoute {
	data_map := make(map[string]models.GatewayExtraRoute)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ExtraRoutes6Value)

		data := models.GatewayExtraRoute{}
		if plan.Via.ValueStringPointer() != nil {
			data.Via = plan.Via.ValueStringPointer()
		}

		data_map[k] = data
	}
	return data_map
}
