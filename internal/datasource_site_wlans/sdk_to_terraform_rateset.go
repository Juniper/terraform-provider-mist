package datasource_site_wlans

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func bandRatesetSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.WlanDatarates) basetypes.ObjectValue {
	var eht basetypes.StringValue
	var he basetypes.StringValue
	var ht basetypes.StringValue
	var legacy = misttransform.ListOfStringSdkToTerraformEmpty()
	var minRssi basetypes.Int64Value
	var template basetypes.StringValue
	var vht basetypes.StringValue

	if d.Eht.Value() != nil {
		eht = types.StringValue(*d.Eht.Value())
	}
	if d.He.Value() != nil {
		he = types.StringValue(*d.He.Value())
	}
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
		minRssi = types.Int64Value(int64(*d.MinRssi))
	}
	if d.Template.Value() != nil {
		template = types.StringValue(string(*d.Template.Value()))
	}
	if d.Vht.Value() != nil {
		vht = types.StringValue(*d.Vht.Value())
	}

	dataMapValue := map[string]attr.Value{
		"eht":      eht,
		"he":       he,
		"ht":       ht,
		"legacy":   legacy,
		"min_rssi": minRssi,
		"template": template,
		"vht":      vht,
	}
	data, e := basetypes.NewObjectValue(RatesetValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func ratesetSkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.WlanDatarates) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)

	for k, d := range m {
		stateValueMap[k] = bandRatesetSkToTerraform(ctx, diags, d)
	}

	stateResult, e := types.MapValueFrom(ctx, RatesetValue{}.Type(ctx), stateValueMap)
	diags.Append(e...)
	return stateResult
}
