package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func bgpConfigNeighborsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.BgpConfigNeighbors) basetypes.MapValue {

	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var disabled basetypes.BoolValue = types.BoolValue(false)
		var export_policy basetypes.StringValue
		var hold_time basetypes.Int64Value = types.Int64Value(90)
		var import_policy basetypes.StringValue
		var multihop_ttl basetypes.Int64Value
		var neighbor_as basetypes.Int64Value

		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.ExportPolicy != nil {
			export_policy = types.StringValue(*d.ExportPolicy)
		}
		if d.HoldTime != nil {
			hold_time = types.Int64Value(int64(*d.HoldTime))
		}
		if d.ImportPolicy != nil {
			import_policy = types.StringValue(*d.ImportPolicy)
		}
		if d.MultihopTtl != nil {
			multihop_ttl = types.Int64Value(int64(*d.MultihopTtl))
		}
		if d.NeighborAs != nil {
			neighbor_as = types.Int64Value(int64(*d.NeighborAs))
		}

		data_map_attr_type := NeighborsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"disabled":      disabled,
			"export_policy": export_policy,
			"hold_time":     hold_time,
			"import_policy": import_policy,
			"multihop_ttl":  multihop_ttl,
			"neighbor_as":   neighbor_as,
		}
		data, e := NewNeighborsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data
	}
	state_result_map_type := NeighborsValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func bgpConfigCommunitiesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.BgpConfigCommunity) basetypes.ListValue {
	var data_list = []CommunitiesValue{}

	for _, d := range l {
		var id basetypes.StringValue
		var local_preference basetypes.Int64Value
		var vpn_name basetypes.StringValue

		if d.Id != nil {
			id = types.StringValue(*d.Id)
		}
		if d.LocalPreference != nil {
			local_preference = types.Int64Value(int64(*d.LocalPreference))
		}
		if d.VpnName != nil {
			vpn_name = types.StringValue(*d.VpnName)
		}

		data_map_attr_type := CommunitiesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"id":               id,
			"local_preference": local_preference,
			"vpn_name":         vpn_name,
		}
		data, e := NewCommunitiesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := CommunitiesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func bgpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.BgpConfig) basetypes.MapValue {
	state_value_map := make(map[string]attr.Value)
	for k, d := range m {
		var auth_key basetypes.StringValue
		var bfd_minimum_interval basetypes.Int64Value = types.Int64Value(350)
		var bfd_multiplier basetypes.Int64Value = types.Int64Value(3)
		var communities basetypes.ListValue = types.ListNull(CommunitiesValue{}.Type(ctx))
		var disable_bfd basetypes.BoolValue = types.BoolValue(false)
		var export basetypes.StringValue
		var export_policy basetypes.StringValue
		var extended_v4_nexthop basetypes.BoolValue
		var graceful_restart_time basetypes.Int64Value = types.Int64Value(0)
		var hold_time basetypes.Int64Value = types.Int64Value(90)
		var import_bgp basetypes.StringValue
		var import_policy basetypes.StringValue
		var local_as basetypes.Int64Value
		var neighbor_as basetypes.Int64Value
		var neighbors basetypes.MapValue = types.MapNull(NeighborsValue{}.Type(ctx))
		var networks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var no_readvertise_to_overlay basetypes.BoolValue = types.BoolValue(false)
		var type_bgp basetypes.StringValue
		var via basetypes.StringValue = types.StringValue("lan")
		var vpn_name basetypes.StringValue
		var wan_name basetypes.StringValue

		if d.AuthKey != nil {
			auth_key = types.StringValue(*d.AuthKey)
		}
		if d.BfdMinimumInterval.Value() != nil {
			bfd_minimum_interval = types.Int64Value(int64(*d.BfdMinimumInterval.Value()))
		}
		if d.BfdMultiplier.Value() != nil {
			bfd_multiplier = types.Int64Value(int64(*d.BfdMultiplier.Value()))
		}
		if d.Communities != nil {
			communities = bgpConfigCommunitiesSdkToTerraform(ctx, diags, d.Communities)
		}
		if d.DisableBfd != nil {
			disable_bfd = types.BoolValue(*d.DisableBfd)
		}
		if d.Export != nil {
			export = types.StringValue(*d.Export)
		}
		if d.ExportPolicy != nil {
			export_policy = types.StringValue(*d.ExportPolicy)
		}
		if d.ExtendedV4Nexthop != nil {
			extended_v4_nexthop = types.BoolValue(*d.ExtendedV4Nexthop)
		}
		if d.GracefulRestartTime != nil {
			graceful_restart_time = types.Int64Value(int64(*d.GracefulRestartTime))
		}
		if d.HoldTime != nil {
			hold_time = types.Int64Value(int64(*d.HoldTime))
		}
		if d.Import != nil {
			import_bgp = types.StringValue(*d.Import)
		}
		if d.ImportPolicy != nil {
			import_policy = types.StringValue(*d.ImportPolicy)
		}
		if d.LocalAs != nil {
			local_as = types.Int64Value(int64(*d.LocalAs))
		}
		if d.NeighborAs != nil {
			neighbor_as = types.Int64Value(int64(*d.NeighborAs))
		}
		if d.Neighbors != nil && len(d.Neighbors) > 0 {
			neighbors = bgpConfigNeighborsSdkToTerraform(ctx, diags, d.Neighbors)
		}
		if d.Networks != nil {
			networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Networks)
		}
		if d.NoReadvertiseToOverlay != nil {
			no_readvertise_to_overlay = types.BoolValue(*d.NoReadvertiseToOverlay)
		}
		if d.Type != nil {
			type_bgp = types.StringValue(string(*d.Type))
		}
		if d.Via != nil {
			via = types.StringValue(string(*d.Via))
		}
		if d.VpnName != nil {
			vpn_name = types.StringValue(*d.VpnName)
		}
		if d.WanName != nil {
			wan_name = types.StringValue(*d.WanName)
		}

		data_map_attr_type := BgpConfigValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"auth_key":                  auth_key,
			"bfd_minimum_interval":      bfd_minimum_interval,
			"bfd_multiplier":            bfd_multiplier,
			"communities":               communities,
			"disable_bfd":               disable_bfd,
			"export":                    export,
			"export_policy":             export_policy,
			"extended_v4_nexthop":       extended_v4_nexthop,
			"graceful_restart_time":     graceful_restart_time,
			"hold_time":                 hold_time,
			"import":                    import_bgp,
			"import_policy":             import_policy,
			"local_as":                  local_as,
			"neighbor_as":               neighbor_as,
			"neighbors":                 neighbors,
			"networks":                  networks,
			"no_readvertise_to_overlay": no_readvertise_to_overlay,
			"type":                      type_bgp,
			"via":                       via,
			"vpn_name":                  vpn_name,
			"wan_name":                  wan_name,
		}
		data, e := NewBgpConfigValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}
	state_type := BgpConfigValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
