package resource_org_network

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func VpnSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.NetworkVpnAccessConfig) basetypes.MapValue {

	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var advertised_subnet basetypes.StringValue
		var allow_ping basetypes.BoolValue
		var destination_nat basetypes.MapValue = types.MapNull(DestinationNatValue{}.Type(ctx))
		var nat_pool basetypes.StringValue
		var no_readvertise_to_lan_bgp basetypes.BoolValue = types.BoolValue(false)
		var no_readvertise_to_lan_ospf basetypes.BoolValue = types.BoolValue(false)
		var no_readvertise_to_overlay basetypes.BoolValue
		var other_vrfs basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var routed basetypes.BoolValue
		var source_nat basetypes.ObjectValue = types.ObjectNull(SourceNatValue{}.AttributeTypes(ctx))
		var static_nat basetypes.MapValue = types.MapNull(StaticNatValue{}.Type(ctx))
		var summarized_subnet basetypes.StringValue
		var summarized_subnet_to_lan_bgp basetypes.StringValue
		var summarized_subnet_to_lan_ospf basetypes.StringValue

		if d.AdvertisedSubnet != nil {
			advertised_subnet = types.StringValue(*d.AdvertisedSubnet)
		}
		if d.AllowPing != nil {
			allow_ping = types.BoolValue(*d.AllowPing)
		}
		if d.DestinationNat != nil && len(d.DestinationNat) > 0 {
			destination_nat = destinationNatSdkToTerraform(ctx, diags, d.DestinationNat)
		}
		if d.NatPool != nil {
			nat_pool = types.StringValue(*d.NatPool)
		}
		if d.NoReadvertiseToLanBgp != nil {
			no_readvertise_to_lan_bgp = types.BoolValue(*d.NoReadvertiseToLanBgp)
		}
		if d.NoReadvertiseToLanOspf != nil {
			no_readvertise_to_lan_ospf = types.BoolValue(*d.NoReadvertiseToLanOspf)
		}
		if d.NoReadvertiseToOverlay != nil {
			no_readvertise_to_overlay = types.BoolValue(*d.NoReadvertiseToOverlay)
		}
		if d.OtherVrfs != nil {
			other_vrfs = mist_transform.ListOfStringSdkToTerraform(ctx, d.OtherVrfs)
		}
		if d.Routed != nil {
			routed = types.BoolValue(*d.Routed)
		}
		if d.SourceNat != nil {
			source_nat = sourceNatSdkToTerraform(ctx, diags, d.SourceNat)
		}
		if d.StaticNat != nil {
			static_nat = staticNatSdkToTerraform(ctx, diags, d.StaticNat)
		}
		if d.SummarizedSubnet != nil {
			summarized_subnet = types.StringValue(*d.SummarizedSubnet)
		}
		if d.SummarizedSubnetToLanBgp != nil {
			summarized_subnet_to_lan_bgp = types.StringValue(*d.SummarizedSubnetToLanBgp)
		}
		if d.SummarizedSubnetToLanOspf != nil {
			summarized_subnet_to_lan_ospf = types.StringValue(*d.SummarizedSubnetToLanOspf)
		}

		data_map_attr_type := VpnAccessValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"advertised_subnet":             advertised_subnet,
			"allow_ping":                    allow_ping,
			"destination_nat":               destination_nat,
			"nat_pool":                      nat_pool,
			"no_readvertise_to_lan_bgp":     no_readvertise_to_lan_bgp,
			"no_readvertise_to_lan_ospf":    no_readvertise_to_lan_ospf,
			"no_readvertise_to_overlay":     no_readvertise_to_overlay,
			"other_vrfs":                    other_vrfs,
			"routed":                        routed,
			"source_nat":                    source_nat,
			"static_nat":                    static_nat,
			"summarized_subnet":             summarized_subnet,
			"summarized_subnet_to_lan_bgp":  summarized_subnet_to_lan_bgp,
			"summarized_subnet_to_lan_ospf": summarized_subnet_to_lan_ospf,
		}
		data, e := NewVpnAccessValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data
	}
	state_result_map_type := VpnAccessValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
