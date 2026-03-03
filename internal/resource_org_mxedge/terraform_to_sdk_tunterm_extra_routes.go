package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermExtraRoutesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.MxedgeTuntermExtraRoute {
	data := make(map[string]models.MxedgeTuntermExtraRoute)

	if d.IsNull() || d.IsUnknown() {
		return data
	}

	// Convert the map to a map of TuntermExtraRoutesValue
	var tfMap map[string]TuntermExtraRoutesValue
	d.ElementsAs(ctx, &tfMap, false)

	for key, value := range tfMap {
		route := models.MxedgeTuntermExtraRoute{}

		if !value.Via.IsNull() && !value.Via.IsUnknown() {
			route.Via = value.Via.ValueStringPointer()
		}

		data[key] = route
	}

	return data
}
