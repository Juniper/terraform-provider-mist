package resource_org_network

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatVpnSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkVpnAccessDestinationNatProperty) basetypes.MapValue {
	stateValueMapValue := make(map[string]attr.Value)
	for k, v := range d {
		var internalIp basetypes.StringValue
		var name basetypes.StringValue
		var port basetypes.StringValue

		if v.InternalIp != nil {
			internalIp = types.StringValue(*v.InternalIp)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.Port != nil {
			port = types.StringValue(*v.Port)
		}

		stateValueMapAttrValue := map[string]attr.Value{
			"internal_ip": internalIp,
			"name":        name,
			"port":        port,
		}
		n, e := NewVpnAccessDestinationNatValue(VpnAccessDestinationNatValue{}.AttributeTypes(ctx), stateValueMapAttrValue)
		diags.Append(e...)

		stateValueMapValue[k] = n
	}
	stateResultMap, e := types.MapValueFrom(ctx, VpnAccessDestinationNatValue{}.Type(ctx), stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func staticNatVpnSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkVpnAccessStaticNatProperty) basetypes.MapValue {
	stateValueMapValue := make(map[string]attr.Value)
	for k, v := range d {
		var internalIp basetypes.StringValue
		var name basetypes.StringValue

		if v.InternalIp != nil {
			internalIp = types.StringValue(*v.InternalIp)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}

		stateValueMapAttrValue := map[string]attr.Value{
			"internal_ip": internalIp,
			"name":        name,
		}
		n, e := NewVpnAccessStaticNatValue(VpnAccessStaticNatValue{}.AttributeTypes(ctx), stateValueMapAttrValue)
		diags.Append(e...)

		stateValueMapValue[k] = n
	}
	stateResultMap, e := types.MapValueFrom(ctx, VpnAccessStaticNatValue{}.Type(ctx), stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func VpnSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.NetworkVpnAccessConfig) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var advertisedSubnet basetypes.StringValue
		var allowPing basetypes.BoolValue
		var destinationNat = types.MapNull(VpnAccessDestinationNatValue{}.Type(ctx))
		var natPool basetypes.StringValue
		var noReadvertiseToLanBgp = types.BoolValue(false)
		var noReadvertiseToLanOspf = types.BoolValue(false)
		var noReadvertiseToOverlay basetypes.BoolValue
		var otherVrfs = mistutils.ListOfStringSdkToTerraformEmpty()
		var routed basetypes.BoolValue
		var sourceNat = types.ObjectNull(SourceNatValue{}.AttributeTypes(ctx))
		var staticNat = types.MapNull(VpnAccessStaticNatValue{}.Type(ctx))
		var summarizedSubnet basetypes.StringValue
		var summarizedSubnetToLanBgp basetypes.StringValue
		var summarizedSubnetToLanOspf basetypes.StringValue

		if d.AdvertisedSubnet != nil {
			advertisedSubnet = types.StringValue(*d.AdvertisedSubnet)
		}
		if d.AllowPing != nil {
			allowPing = types.BoolValue(*d.AllowPing)
		}
		if d.DestinationNat != nil && len(d.DestinationNat) > 0 {
			destinationNat = destinationNatVpnSdkToTerraform(ctx, diags, d.DestinationNat)
		}
		if d.NatPool != nil {
			natPool = types.StringValue(*d.NatPool)
		}
		if d.NoReadvertiseToLanBgp != nil {
			noReadvertiseToLanBgp = types.BoolValue(*d.NoReadvertiseToLanBgp)
		}
		if d.NoReadvertiseToLanOspf != nil {
			noReadvertiseToLanOspf = types.BoolValue(*d.NoReadvertiseToLanOspf)
		}
		if d.NoReadvertiseToOverlay != nil {
			noReadvertiseToOverlay = types.BoolValue(*d.NoReadvertiseToOverlay)
		}
		if d.OtherVrfs != nil {
			otherVrfs = mistutils.ListOfStringSdkToTerraform(d.OtherVrfs)
		}
		if d.Routed != nil {
			routed = types.BoolValue(*d.Routed)
		}
		if d.SourceNat != nil {
			sourceNat = sourceNatSdkToTerraform(ctx, diags, d.SourceNat)
		}
		if d.StaticNat != nil {
			staticNat = staticNatVpnSdkToTerraform(ctx, diags, d.StaticNat)
		}
		if d.SummarizedSubnet != nil {
			summarizedSubnet = types.StringValue(*d.SummarizedSubnet)
		}
		if d.SummarizedSubnetToLanBgp != nil {
			summarizedSubnetToLanBgp = types.StringValue(*d.SummarizedSubnetToLanBgp)
		}
		if d.SummarizedSubnetToLanOspf != nil {
			summarizedSubnetToLanOspf = types.StringValue(*d.SummarizedSubnetToLanOspf)
		}

		dataMapValue := map[string]attr.Value{
			"advertised_subnet":             advertisedSubnet,
			"allow_ping":                    allowPing,
			"destination_nat":               destinationNat,
			"nat_pool":                      natPool,
			"no_readvertise_to_lan_bgp":     noReadvertiseToLanBgp,
			"no_readvertise_to_lan_ospf":    noReadvertiseToLanOspf,
			"no_readvertise_to_overlay":     noReadvertiseToOverlay,
			"other_vrfs":                    otherVrfs,
			"routed":                        routed,
			"source_nat":                    sourceNat,
			"static_nat":                    staticNat,
			"summarized_subnet":             summarizedSubnet,
			"summarized_subnet_to_lan_bgp":  summarizedSubnetToLanBgp,
			"summarized_subnet_to_lan_ospf": summarizedSubnetToLanOspf,
		}
		data, e := NewVpnAccessValue(VpnAccessValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMapType := VpnAccessValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}
