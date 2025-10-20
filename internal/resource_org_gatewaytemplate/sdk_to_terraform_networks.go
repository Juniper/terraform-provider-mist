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

func multicastNetworksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NetworkMulticast) basetypes.ObjectValue {

	var disableIgmp basetypes.BoolValue
	var enabled basetypes.BoolValue
	var groups = types.MapNull(GroupsValue{}.Type(ctx))

	if d != nil && d.DisableIgmp != nil {
		disableIgmp = types.BoolValue(*d.DisableIgmp)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Groups != nil {
		rMapValue := make(map[string]attr.Value)
		for k, v := range d.Groups {
			var rpIp types.String
			if v.RpIp != nil {
				rpIp = types.StringValue(*v.RpIp)
			}
			dataMapValue := map[string]attr.Value{
				"rp_ip": rpIp,
			}
			data, e := NewGroupsValue(GroupsValue{}.AttributeTypes(ctx), dataMapValue)
			diags.Append(e...)

			rMapValue[k] = data
		}
		r, e := types.MapValueFrom(ctx, GroupsValue{}.Type(ctx), rMapValue)
		if e != nil {
			diags.Append(e...)
		} else {
			groups = r
		}
	}

	dataMapValue := map[string]attr.Value{
		"disable_igmp": disableIgmp,
		"enabled":      enabled,
		"groups":       groups,
	}
	data, e := NewMulticastValue(MulticastValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func internalAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkInternalAccess) basetypes.ObjectValue {
	var enabled basetypes.BoolValue
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := NewInternalAccessValue(InternalAccessValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func destinationNatInternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkInternetAccessDestinationNatProperty) basetypes.MapValue {
	stateValueMapValue := make(map[string]attr.Value)
	for k, v := range d {
		var internalIp basetypes.StringValue
		var name basetypes.StringValue
		var port basetypes.StringValue
		var wanName basetypes.StringValue

		if v.InternalIp != nil {
			internalIp = types.StringValue(*v.InternalIp)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.Port != nil {
			port = types.StringValue(*v.Port)
		}
		if v.WanName != nil {
			wanName = types.StringValue(*v.WanName)
		}

		stateValueMapAttrValue := map[string]attr.Value{
			"internal_ip": internalIp,
			"name":        name,
			"port":        port,
			"wan_name":    wanName,
		}
		n, e := NewInternetAccessDestinationNatValue(InternetAccessDestinationNatValue{}.AttributeTypes(ctx), stateValueMapAttrValue)
		diags.Append(e...)

		stateValueMapValue[k] = n
	}
	stateResultMap, e := types.MapValueFrom(ctx, InternetAccessDestinationNatValue{}.Type(ctx), stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func staticNatInternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkInternetAccessStaticNatProperty) basetypes.MapValue {
	stateValueMapValue := make(map[string]attr.Value)
	for k, v := range d {
		var internalIp basetypes.StringValue
		var name basetypes.StringValue
		var wanName basetypes.StringValue

		if v.InternalIp != nil {
			internalIp = types.StringValue(*v.InternalIp)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.WanName != nil {
			wanName = types.StringValue(*v.WanName)
		}

		stateValueMapAttrValue := map[string]attr.Value{
			"internal_ip": internalIp,
			"name":        name,
			"wan_name":    wanName,
		}
		n, e := NewInternetAccessStaticNatValue(InternetAccessStaticNatValue{}.AttributeTypes(ctx), stateValueMapAttrValue)
		diags.Append(e...)

		stateValueMapValue[k] = n
	}
	stateResultMap, e := types.MapValueFrom(ctx, InternetAccessStaticNatValue{}.Type(ctx), stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func internetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkInternetAccess) basetypes.ObjectValue {
	var createSimpleServicePolicy = types.BoolValue(false)
	var destinationNat = types.MapNull(InternetAccessDestinationNatValue{}.Type(ctx))
	var enabled basetypes.BoolValue
	var restricted = types.BoolValue(false)
	var staticNat = types.MapNull(InternetAccessStaticNatValue{}.Type(ctx))

	if d.CreateSimpleServicePolicy != nil {
		createSimpleServicePolicy = types.BoolValue(*d.CreateSimpleServicePolicy)
	}
	if len(d.DestinationNat) > 0 {
		destinationNat = destinationNatInternetAccessSdkToTerraform(ctx, diags, d.DestinationNat)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Restricted != nil {
		restricted = types.BoolValue(*d.Restricted)
	}
	if len(d.StaticNat) > 0 {
		staticNat = staticNatInternetAccessSdkToTerraform(ctx, diags, d.StaticNat)
	}

	dataMapValue := map[string]attr.Value{
		"create_simple_service_policy": createSimpleServicePolicy,
		"destination_nat":              destinationNat,
		"enabled":                      enabled,
		"restricted":                   restricted,
		"static_nat":                   staticNat,
	}
	data, e := NewInternetAccessValue(InternetAccessValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)

	return o
}

func tenantSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkTenant) basetypes.MapValue {

	stateValueMapAttrType := TenantsValue{}.AttributeTypes(ctx)
	stateValueMapValue := make(map[string]attr.Value)
	for k, v := range d {
		stateValueMapAttrValue := map[string]attr.Value{
			"addresses": mistutils.ListOfStringSdkToTerraform(v.Addresses),
		}
		n, e := NewTenantsValue(stateValueMapAttrType, stateValueMapAttrValue)
		diags.Append(e...)
		stateValueMapValue[k] = n
	}
	stateResultMapType := TenantsValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

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

func sourceNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NetworkSourceNat) basetypes.ObjectValue {
	var externalIp basetypes.StringValue

	if d != nil && d.ExternalIp != nil {
		externalIp = types.StringValue(*d.ExternalIp)
	}

	rAttrType := SourceNatValue{}.AttributeTypes(ctx)
	rAttrValue := map[string]attr.Value{
		"external_ip": externalIp,
	}

	r, e := basetypes.NewObjectValue(rAttrType, rAttrValue)
	diags.Append(e...)
	return r
}

func vpnSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.NetworkVpnAccessConfig) basetypes.MapValue {

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
		if len(d.DestinationNat) > 0 {
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

func networksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m []models.Network) basetypes.ListValue {
	var dataList []NetworksValue

	for _, d := range m {
		var disallowMistServices = types.BoolValue(false)
		var gateway basetypes.StringValue
		var gateway6 basetypes.StringValue
		var internalAccess = types.ObjectNull(InternalAccessValue{}.AttributeTypes(ctx))
		var internetAccess = types.ObjectNull(InternetAccessValue{}.AttributeTypes(ctx))
		var isolation basetypes.BoolValue
		var multicast = types.ObjectNull(MulticastValue{}.AttributeTypes(ctx))
		var name basetypes.StringValue
		var routedForNetworks = types.ListNull(types.StringType)
		var subnet basetypes.StringValue
		var subnet6 basetypes.StringValue
		var tenants = types.MapNull(TenantsValue{}.Type(ctx))
		var vlanId basetypes.StringValue
		var vpnAccess = types.MapNull(VpnAccessValue{}.Type(ctx))

		if d.DisallowMistServices != nil {
			disallowMistServices = types.BoolValue(*d.DisallowMistServices)
		}
		if d.Gateway != nil {
			gateway = types.StringValue(*d.Gateway)
		}
		if d.Gateway6 != nil {
			gateway6 = types.StringValue(*d.Gateway6)
		}
		if d.InternalAccess != nil {
			internalAccess = internalAccessSdkToTerraform(ctx, diags, *d.InternalAccess)
		}
		if d.InternetAccess != nil {
			internetAccess = internetAccessSdkToTerraform(ctx, diags, *d.InternetAccess)
		}
		if d.Isolation != nil {
			isolation = types.BoolValue(*d.Isolation)
		}
		if d.Multicast != nil {
			multicast = multicastNetworksSdkToTerraform(ctx, diags, d.Multicast)
		}
		name = types.StringValue(d.Name)
		if d.RoutedForNetworks != nil {
			routedForNetworks = mistutils.ListOfStringSdkToTerraform(d.RoutedForNetworks)
		}
		if d.Subnet != nil {
			subnet = types.StringValue(*d.Subnet)
		}
		if d.Subnet6 != nil {
			subnet6 = types.StringValue(*d.Subnet6)
		}
		if len(d.Tenants) > 0 {
			tenants = tenantSdkToTerraform(ctx, diags, d.Tenants)
		}
		if d.VlanId != nil {
			vlanId = mistutils.VlanAsString(*d.VlanId)
		}
		if len(d.VpnAccess) > 0 {
			vpnAccess = vpnSdkToTerraform(ctx, diags, d.VpnAccess)
		}

		dataMapValue := map[string]attr.Value{
			"disallow_mist_services": disallowMistServices,
			"gateway":                gateway,
			"gateway6":               gateway6,
			"internal_access":        internalAccess,
			"internet_access":        internetAccess,
			"isolation":              isolation,
			"multicast":              multicast,
			"name":                   name,
			"routed_for_networks":    routedForNetworks,
			"subnet":                 subnet,
			"subnet6":                subnet6,
			"tenants":                tenants,
			"vlan_id":                vlanId,
			"vpn_access":             vpnAccess,
		}
		data, e := NewNetworksValue(NetworksValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, NetworksValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
