package resource_org_wlan

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func bandRatesetSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.WlanDatarates) basetypes.ObjectValue {
	var ht basetypes.StringValue
	var legacy basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var min_rssi basetypes.Int64Value
	var template basetypes.StringValue
	var vht basetypes.StringValue

	if d.Ht.Value() != nil {
		ht = types.StringValue(*d.Ht.Value())
	}
	if len(d.Legacy) > 0 {
		var items []attr.Value
		for _, item := range d.Legacy {
			items = append(items, types.StringValue(string(item)))
		}
		list, e := types.ListValue(basetypes.StringType{}, items)
		diags.Append(e...)
		legacy = list
	}
	if d.MinRssi != nil {
		min_rssi = types.Int64Value(int64(*d.MinRssi))
	}
	if d.Template.Value() != nil {
		template = types.StringValue(string(*d.Template.Value()))
	}
	if d.Vht.Value() != nil {
		vht = types.StringValue(*d.Vht.Value())
	}

	data_map_value := map[string]attr.Value{
		"ht":       ht,
		"legacy":   legacy,
		"min_rssi": min_rssi,
		"template": template,
		"vht":      vht,
	}
	data, e := basetypes.NewObjectValue(RatesetValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}

func ratesetSkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.WlanDatarates) basetypes.MapValue {
	state_value_map := make(map[string]attr.Value)

	for k, d := range m {
		state_value_map[k] = bandRatesetSkToTerraform(ctx, diags, d)
	}

	state_result, e := types.MapValueFrom(ctx, RatesetValue{}.Type(ctx), state_value_map)
	diags.Append(e...)
	return state_result
}
