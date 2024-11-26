package resource_org_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func band24RatesetSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.WlanDatarates) basetypes.ObjectValue {
	var ht basetypes.StringValue
	var legacy basetypes.ListValue = types.ListNull(types.StringType)
	var min_rssi basetypes.Int64Value
	var template basetypes.StringValue
	var vht basetypes.StringValue

	if d.Ht.Value() != nil {
		ht = types.StringValue(*d.Ht.Value())
	}
	if d.Legacy != nil {
		var items []attr.Value
		var items_type attr.Type = basetypes.StringType{}
		for _, item := range d.Legacy {
			items = append(items, types.StringValue(string(item)))
		}
		list, _ := types.ListValue(items_type, items)
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
	data, e := basetypes.NewObjectValue(Band24Value{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}

func band5RatesetSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.WlanDatarates) basetypes.ObjectValue {
	var ht basetypes.StringValue
	var legacy basetypes.ListValue = types.ListNull(types.StringType)
	var min_rssi basetypes.Int64Value
	var template basetypes.StringValue
	var vht basetypes.StringValue

	if d.Ht.Value() != nil {
		ht = types.StringValue(*d.Ht.Value())
	}
	if d.Legacy != nil && len(d.Legacy) > 0 {
		var items []attr.Value
		var items_type attr.Type = basetypes.StringType{}
		for _, item := range d.Legacy {
			items = append(items, types.StringValue(string(item)))
		}
		list, _ := types.ListValue(items_type, items)
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
	data, e := basetypes.NewObjectValue(Band5Value{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}

func ratesetSkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.WlanDatarates) RatesetValue {

	var band_24 types.Object = types.ObjectNull(Band24Value{}.AttributeTypes(ctx))
	var band_5 types.Object = types.ObjectNull(Band5Value{}.AttributeTypes(ctx))

	if rateset, ok := m["24"]; ok {
		band_24 = band24RatesetSkToTerraform(ctx, diags, rateset)
	}

	if rateset, ok := m["5"]; ok {
		band_5 = band24RatesetSkToTerraform(ctx, diags, rateset)
	}

	data_map_value := map[string]attr.Value{
		"band_24": band_24,
		"band_5":  band_5,
	}
	data, e := NewRatesetValue(RatesetValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}
