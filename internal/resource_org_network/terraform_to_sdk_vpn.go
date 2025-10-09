package resource_org_network

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatVpnTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkVpnAccessDestinationNatProperty {
	dataMap := make(map[string]models.NetworkVpnAccessDestinationNatProperty)
	for k, v := range d.Elements() {
		// Extract attributes directly from the ObjectValue instead of casting to specific type
		if objVal, ok := v.(basetypes.ObjectValue); ok {
			data := models.NetworkVpnAccessDestinationNatProperty{}
			attrs := objVal.Attributes()

			if internalIp, exists := attrs["internal_ip"]; exists {
				if strVal, ok := internalIp.(basetypes.StringValue); ok {
					data.InternalIp = strVal.ValueStringPointer()
				}
			}
			if name, exists := attrs["name"]; exists {
				if strVal, ok := name.(basetypes.StringValue); ok {
					data.Name = strVal.ValueStringPointer()
				}
			}
			if port, exists := attrs["port"]; exists {
				if strVal, ok := port.(basetypes.StringValue); ok {
					data.Port = strVal.ValueStringPointer()
				}
			}
			dataMap[k] = data
		}
	}
	return dataMap
}

func staticNatVpnTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkVpnAccessStaticNatProperty {
	dataMap := make(map[string]models.NetworkVpnAccessStaticNatProperty)
	for k, v := range d.Elements() {
		// Extract attributes directly from the ObjectValue instead of casting to specific type
		if objVal, ok := v.(basetypes.ObjectValue); ok {
			data := models.NetworkVpnAccessStaticNatProperty{}
			attrs := objVal.Attributes()

			if internalIp, exists := attrs["internal_ip"]; exists {
				if strVal, ok := internalIp.(basetypes.StringValue); ok {
					data.InternalIp = strVal.ValueStringPointer()
				}
			}
			if name, exists := attrs["name"]; exists {
				if strVal, ok := name.(basetypes.StringValue); ok {
					data.Name = strVal.ValueStringPointer()
				}
			}
			dataMap[k] = data
		}
	}
	return dataMap
}

func vpnTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkVpnAccessConfig {
	dataMap := make(map[string]models.NetworkVpnAccessConfig)
	for k, v := range d.Elements() {
		// Extract attributes directly from the ObjectValue instead of casting to specific type
		if objVal, ok := v.(basetypes.ObjectValue); ok {
			data := models.NetworkVpnAccessConfig{}
			attrs := objVal.Attributes()

			if advertisedSubnet, exists := attrs["advertised_subnet"]; exists {
				if strVal, ok := advertisedSubnet.(basetypes.StringValue); ok && !strVal.IsNull() && !strVal.IsUnknown() {
					data.AdvertisedSubnet = strVal.ValueStringPointer()
				}
			}
			if allowPing, exists := attrs["allow_ping"]; exists {
				if boolVal, ok := allowPing.(basetypes.BoolValue); ok && !boolVal.IsNull() && !boolVal.IsUnknown() {
					data.AllowPing = boolVal.ValueBoolPointer()
				}
			}
			if destinationNat, exists := attrs["destination_nat"]; exists {
				if mapVal, ok := destinationNat.(basetypes.MapValue); ok && !mapVal.IsNull() && !mapVal.IsUnknown() {
					data.DestinationNat = destinationNatVpnTerraformToSdk(mapVal)
				}
			}
			if natPool, exists := attrs["nat_pool"]; exists {
				if strVal, ok := natPool.(basetypes.StringValue); ok && !strVal.IsNull() && !strVal.IsUnknown() {
					data.NatPool = strVal.ValueStringPointer()
				}
			}
			if noReadvertiseToLanBgp, exists := attrs["no_readvertise_to_lan_bgp"]; exists {
				if boolVal, ok := noReadvertiseToLanBgp.(basetypes.BoolValue); ok && !boolVal.IsNull() && !boolVal.IsUnknown() {
					data.NoReadvertiseToLanBgp = boolVal.ValueBoolPointer()
				}
			}
			if noReadvertiseToLanOspf, exists := attrs["no_readvertise_to_lan_ospf"]; exists {
				if boolVal, ok := noReadvertiseToLanOspf.(basetypes.BoolValue); ok && !boolVal.IsNull() && !boolVal.IsUnknown() {
					data.NoReadvertiseToLanOspf = boolVal.ValueBoolPointer()
				}
			}
			if noReadvertiseToOverlay, exists := attrs["no_readvertise_to_overlay"]; exists {
				if boolVal, ok := noReadvertiseToOverlay.(basetypes.BoolValue); ok && !boolVal.IsNull() && !boolVal.IsUnknown() {
					data.NoReadvertiseToOverlay = boolVal.ValueBoolPointer()
				}
			}
			if otherVrfs, exists := attrs["other_vrfs"]; exists {
				if listVal, ok := otherVrfs.(basetypes.ListValue); ok && !listVal.IsNull() && !listVal.IsUnknown() {
					data.OtherVrfs = mistutils.ListOfStringTerraformToSdk(listVal)
				}
			}
			if routed, exists := attrs["routed"]; exists {
				if boolVal, ok := routed.(basetypes.BoolValue); ok && !boolVal.IsNull() && !boolVal.IsUnknown() {
					data.Routed = boolVal.ValueBoolPointer()
				}
			}
			if sourceNat, exists := attrs["source_nat"]; exists {
				if objVal, ok := sourceNat.(basetypes.ObjectValue); ok && !objVal.IsNull() && !objVal.IsUnknown() {
					data.SourceNat = sourceNatTerraformToSdk(ctx, diags, objVal)
				}
			}
			if staticNat, exists := attrs["static_nat"]; exists {
				if mapVal, ok := staticNat.(basetypes.MapValue); ok && !mapVal.IsNull() && !mapVal.IsUnknown() {
					data.StaticNat = staticNatVpnTerraformToSdk(mapVal)
				}
			}
			if summarizedSubnet, exists := attrs["summarized_subnet"]; exists {
				if strVal, ok := summarizedSubnet.(basetypes.StringValue); ok && !strVal.IsNull() && !strVal.IsUnknown() {
					data.SummarizedSubnet = strVal.ValueStringPointer()
				}
			}
			if summarizedSubnetToLanBgp, exists := attrs["summarized_subnet_to_lan_bgp"]; exists {
				if strVal, ok := summarizedSubnetToLanBgp.(basetypes.StringValue); ok && !strVal.IsNull() && !strVal.IsUnknown() {
					data.SummarizedSubnetToLanBgp = strVal.ValueStringPointer()
				}
			}
			if summarizedSubnetToLanOspf, exists := attrs["summarized_subnet_to_lan_ospf"]; exists {
				if strVal, ok := summarizedSubnetToLanOspf.(basetypes.StringValue); ok && !strVal.IsNull() && !strVal.IsUnknown() {
					data.SummarizedSubnetToLanOspf = strVal.ValueStringPointer()
				}
			}

			dataMap[k] = data
		}
	}
	return dataMap
}
