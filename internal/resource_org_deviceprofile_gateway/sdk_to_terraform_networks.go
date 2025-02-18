package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	mistnetwork "github.com/Juniper/terraform-provider-mist/internal/resource_org_network"
)

func multicastNetworksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NetworkMulticast) MulticastValue {

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

	return data
}

func networksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m []models.Network) basetypes.ListValue {
	var dataList []NetworksValue

	for _, d := range m {
		var disallowMistServices = types.BoolValue(false)
		var gateway basetypes.StringValue
		var gateway6 basetypes.StringValue
		var internalAccess = types.ObjectNull(mistnetwork.InternalAccessValue{}.AttributeTypes(ctx))
		var internetAccess = types.ObjectNull(mistnetwork.InternetAccessValue{}.AttributeTypes(ctx))
		var isolation basetypes.BoolValue
		var name basetypes.StringValue
		var multicast = NewMulticastValueNull()
		var routedForNetworks = misttransform.ListOfStringSdkToTerraformEmpty()
		var subnet basetypes.StringValue
		var subnet6 basetypes.StringValue
		var tenants = types.MapNull(mistnetwork.TenantsValue{}.Type(ctx))
		var vlanId basetypes.StringValue
		var vpnAccess = types.MapNull(mistnetwork.VpnAccessValue{}.Type(ctx))

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
			internalAccess, _ = mistnetwork.InternalAccessSdkToTerraform(ctx, diags, *d.InternalAccess).ToObjectValue(ctx)
		}
		if d.InternetAccess != nil {
			internetAccess, _ = mistnetwork.InternetAccessSdkToTerraform(ctx, diags, *d.InternetAccess).ToObjectValue(ctx)
		}
		if d.Isolation != nil {
			isolation = types.BoolValue(*d.Isolation)
		}
		if d.Multicast != nil {
			multicast = multicastNetworksSdkToTerraform(ctx, diags, d.Multicast)
		}
		name = types.StringValue(d.Name)
		if d.RoutedForNetworks != nil {
			routedForNetworks = misttransform.ListOfStringSdkToTerraform(d.RoutedForNetworks)
		}
		if d.Subnet != nil {
			subnet = types.StringValue(*d.Subnet)
		}
		if d.Subnet6 != nil {
			subnet6 = types.StringValue(*d.Subnet6)
		}
		if d.Tenants != nil && len(d.Tenants) > 0 {
			tenants = mistnetwork.TenantSdkToTerraform(ctx, diags, d.Tenants)
		}
		if d.VlanId != nil {
			vlanId = types.StringValue(d.VlanId.String())
		}
		if d.VpnAccess != nil && len(d.VpnAccess) > 0 {
			vpnAccess = mistnetwork.VpnSdkToTerraform(ctx, diags, d.VpnAccess)
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
	datalistType := NetworksValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}
