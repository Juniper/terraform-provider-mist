package resource_org_networktemplate

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func bgpConfigNeighborsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchBgpConfigNeighbor) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var exportPolicy basetypes.StringValue
		var holdTime basetypes.Int64Value
		var importPolicy basetypes.StringValue
		var multihopTtl basetypes.Int64Value
		var neighborAs basetypes.StringValue

		if d.ExportPolicy != nil {
			exportPolicy = types.StringValue(*d.ExportPolicy)
		}
		if d.HoldTime != nil {
			holdTime = types.Int64Value(int64(*d.HoldTime))
		}
		if d.ImportPolicy != nil {
			importPolicy = types.StringValue(*d.ImportPolicy)
		}
		if d.MultihopTtl != nil {
			multihopTtl = types.Int64Value(int64(*d.MultihopTtl))
		}

		// NeighborAs is a union type - try string first, then number
		if strVal, ok := d.NeighborAs.AsString(); ok && strVal != nil {
			neighborAs = types.StringValue(*strVal)
		} else if numVal, ok := d.NeighborAs.AsNumber(); ok && numVal != nil {
			neighborAs = types.StringValue(fmt.Sprintf("%d", *numVal))
		}

		dataMapValue := map[string]attr.Value{
			"export_policy": exportPolicy,
			"hold_time":     holdTime,
			"import_policy": importPolicy,
			"multihop_ttl":  multihopTtl,
			"neighbor_as":   neighborAs,
		}
		data, e := NewNeighborsValue(NeighborsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMap, e := types.MapValueFrom(ctx, NeighborsValue{}.Type(ctx), stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func bgpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchBgpConfig) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {
		var authKey basetypes.StringValue
		var bfdMinimumInterval basetypes.Int64Value
		var exportPolicy basetypes.StringValue
		var holdTime basetypes.Int64Value
		var importPolicy basetypes.StringValue
		var localAs basetypes.StringValue
		var neighbors = types.MapNull(NeighborsValue{}.Type(ctx))
		var networks = types.ListNull(types.StringType)
		var typeBgp basetypes.StringValue

		if d.AuthKey != nil {
			authKey = types.StringValue(*d.AuthKey)
		}
		if d.BfdMinimumInterval != nil {
			bfdMinimumInterval = types.Int64Value(int64(*d.BfdMinimumInterval))
		}
		if d.ExportPolicy != nil {
			exportPolicy = types.StringValue(*d.ExportPolicy)
		}
		if d.HoldTime != nil {
			holdTime = types.Int64Value(int64(*d.HoldTime))
		}
		if d.ImportPolicy != nil {
			importPolicy = types.StringValue(*d.ImportPolicy)
		}
		// LocalAs is a union type - try string first, then number
		if strVal, ok := d.LocalAs.AsString(); ok && strVal != nil {
			localAs = types.StringValue(*strVal)
		} else if numVal, ok := d.LocalAs.AsNumber(); ok && numVal != nil {
			localAs = types.StringValue(fmt.Sprintf("%d", *numVal))
		}
		if len(d.Neighbors) > 0 {
			neighbors = bgpConfigNeighborsSdkToTerraform(ctx, diags, d.Neighbors)
		}
		if d.Networks != nil {
			networks = mistutils.ListOfStringSdkToTerraform(d.Networks)
		}
		typeBgp = types.StringValue(string(d.Type))

		dataMapValue := map[string]attr.Value{
			"auth_key":             authKey,
			"bfd_minimum_interval": bfdMinimumInterval,
			"export_policy":        exportPolicy,
			"hold_time":            holdTime,
			"import_policy":        importPolicy,
			"local_as":             localAs,
			"neighbors":            neighbors,
			"networks":             networks,
			"type":                 typeBgp,
		}
		data, e := NewBgpConfigValue(BgpConfigValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMap[k] = data
	}
	stateType := BgpConfigValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, stateValueMap)
	diags.Append(e...)
	return stateResult
}
