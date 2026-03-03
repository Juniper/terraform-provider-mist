package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermOtherIpConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.MxedgeTuntermOtherIpConfig) basetypes.MapValue {

	state_value_map_type := TuntermOtherIpConfigsValue{}.Type(ctx)
	state_value_map := make(map[string]attr.Value)

	for k, v := range d {
		var ip types.String
		var netmask types.String

		// Required fields
		ip = types.StringValue(v.Ip)
		netmask = types.StringValue(v.Netmask)

		data_map_attr_type := TuntermOtherIpConfigsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"ip":      ip,
			"netmask": netmask,
		}
		data, e := NewTuntermOtherIpConfigsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}

	state_result, e := types.MapValueFrom(ctx, state_value_map_type, state_value_map)
	diags.Append(e...)
	return state_result
}
