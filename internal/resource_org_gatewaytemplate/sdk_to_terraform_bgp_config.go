package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func bgpConfigNeighborsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.BgpConfigNeighbors) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var disabled = types.BoolValue(false)
		var exportPolicy basetypes.StringValue
		var holdTime basetypes.Int64Value
		var importPolicy basetypes.StringValue
		var multihopTtl basetypes.Int64Value
		var neighborAs basetypes.StringValue

		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
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
		if d.MultihopTtl != nil {
			multihopTtl = types.Int64Value(int64(*d.MultihopTtl))
		}

		neighborAs = mistutils.BgpAsAsString(&d.NeighborAs)

		dataMapValue := map[string]attr.Value{
			"disabled":      disabled,
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

func bgpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.BgpConfig) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {
		var authKey basetypes.StringValue
		var bfdMinimumInterval basetypes.Int64Value
		var bfdMultiplier basetypes.Int64Value
		var disableBfd basetypes.BoolValue
		var export basetypes.StringValue
		var exportPolicy basetypes.StringValue
		var extendedV4Nexthop basetypes.BoolValue
		var gracefulRestartTime basetypes.Int64Value
		var holdTime basetypes.Int64Value
		var importBgp basetypes.StringValue
		var importPolicy basetypes.StringValue
		var localAs basetypes.StringValue
		var neighborAs basetypes.StringValue
		var neighbors = types.MapNull(NeighborsValue{}.Type(ctx))
		var networks = types.ListNull(types.StringType)
		var noPrivateAs basetypes.BoolValue
		var noReadvertiseToOverlay basetypes.BoolValue
		var typeBgp basetypes.StringValue
		var tunnelName basetypes.StringValue
		var via basetypes.StringValue
		var vpnName basetypes.StringValue
		var wanName basetypes.StringValue

		if d.AuthKey != nil {
			authKey = types.StringValue(*d.AuthKey)
		}
		if d.BfdMinimumInterval.Value() != nil {
			bfdMinimumInterval = types.Int64Value(int64(*d.BfdMinimumInterval.Value()))
		}
		if d.BfdMultiplier.Value() != nil {
			bfdMultiplier = types.Int64Value(int64(*d.BfdMultiplier.Value()))
		}
		if d.DisableBfd != nil {
			disableBfd = types.BoolValue(*d.DisableBfd)
		}
		if d.Export != nil {
			export = types.StringValue(*d.Export)
		}
		if d.ExportPolicy != nil {
			exportPolicy = types.StringValue(*d.ExportPolicy)
		}
		if d.ExtendedV4Nexthop != nil {
			extendedV4Nexthop = types.BoolValue(*d.ExtendedV4Nexthop)
		}
		if d.GracefulRestartTime != nil {
			gracefulRestartTime = types.Int64Value(int64(*d.GracefulRestartTime))
		}
		if d.HoldTime != nil {
			holdTime = types.Int64Value(int64(*d.HoldTime))
		}
		if d.Import != nil {
			importBgp = types.StringValue(*d.Import)
		}
		if d.ImportPolicy != nil {
			importPolicy = types.StringValue(*d.ImportPolicy)
		}
		if d.LocalAs != nil {
			localAs = mistutils.BgpLocalAsAsString(d.LocalAs)
		}
		if d.NeighborAs != nil {
			neighborAs = mistutils.BgpAsAsString(d.NeighborAs)
		}
		if len(d.Neighbors) > 0 {
			neighbors = bgpConfigNeighborsSdkToTerraform(ctx, diags, d.Neighbors)
		}
		if d.Networks != nil {
			networks = mistutils.ListOfStringSdkToTerraform(d.Networks)
		}
		if d.NoPrivateAs != nil {
			noPrivateAs = types.BoolValue(*d.NoPrivateAs)
		}
		if d.NoReadvertiseToOverlay != nil {
			noReadvertiseToOverlay = types.BoolValue(*d.NoReadvertiseToOverlay)
		}
		if d.Type != nil {
			typeBgp = types.StringValue(string(*d.Type))
		}
		if d.TunnelName != nil {
			tunnelName = types.StringValue(*d.TunnelName)
		}

		via = types.StringValue(string(d.Via))

		if d.VpnName != nil {
			vpnName = types.StringValue(*d.VpnName)
		}
		if d.WanName != nil {
			wanName = types.StringValue(*d.WanName)
		}

		dataMapValue := map[string]attr.Value{
			"auth_key":                  authKey,
			"bfd_minimum_interval":      bfdMinimumInterval,
			"bfd_multiplier":            bfdMultiplier,
			"disable_bfd":               disableBfd,
			"export":                    export,
			"export_policy":             exportPolicy,
			"extended_v4_nexthop":       extendedV4Nexthop,
			"graceful_restart_time":     gracefulRestartTime,
			"hold_time":                 holdTime,
			"import":                    importBgp,
			"import_policy":             importPolicy,
			"local_as":                  localAs,
			"neighbor_as":               neighborAs,
			"neighbors":                 neighbors,
			"networks":                  networks,
			"no_private_as":             noPrivateAs,
			"no_readvertise_to_overlay": noReadvertiseToOverlay,
			"type":                      typeBgp,
			"tunnel_name":               tunnelName,
			"via":                       via,
			"vpn_name":                  vpnName,
			"wan_name":                  wanName,
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
