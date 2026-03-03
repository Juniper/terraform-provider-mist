package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermOtherIpConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.MxedgeTuntermOtherIpConfig {
	data := make(map[string]models.MxedgeTuntermOtherIpConfig)

	if d.IsNull() || d.IsUnknown() {
		return data
	}

	// Convert the map to a map of TuntermOtherIpConfigsValue
	var tfMap map[string]TuntermOtherIpConfigsValue
	d.ElementsAs(ctx, &tfMap, false)

	for key, value := range tfMap {
		config := models.MxedgeTuntermOtherIpConfig{}

		// Required fields
		config.Ip = value.Ip.ValueString()
		config.Netmask = value.Netmask.ValueString()

		data[key] = config
	}

	return data
}
