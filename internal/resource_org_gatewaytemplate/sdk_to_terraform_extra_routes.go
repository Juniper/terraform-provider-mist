package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func extraRoutesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.GatewayExtraRoute) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var via basetypes.StringValue

		if d.Via != nil {
			via = types.StringValue(*d.Via)
		}

		dataMapValue := map[string]attr.Value{
			"via": via,
		}
		data, e := NewExtraRoutesValue(ExtraRoutesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMapType := ExtraRoutesValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func extraRoutes6SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.GatewayExtraRoute) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var via basetypes.StringValue

		if d.Via != nil {
			via = types.StringValue(*d.Via)
		}

		dataMapValue := map[string]attr.Value{
			"via": via,
		}
		data, e := NewExtraRoutesValue(ExtraRoutes6Value{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMapType := ExtraRoutes6Value{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}
