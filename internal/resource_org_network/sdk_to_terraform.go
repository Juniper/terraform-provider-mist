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

func SdkToTerraform(ctx context.Context, data models.Network) (OrgNetworkModel, diag.Diagnostics) {
	var state OrgNetworkModel
	var diags diag.Diagnostics

	var disallowMistServices basetypes.BoolValue
	var multicast = NewMulticastValueNull()
	var gateway basetypes.StringValue
	var gateway6 basetypes.StringValue
	var id basetypes.StringValue
	var internalAccess = NewInternalAccessValueNull()
	var internetAccess = NewInternetAccessValueNull()
	var isolation basetypes.BoolValue
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var routedForNetworks = types.ListValueMust(types.StringType, []attr.Value{})
	var subnet basetypes.StringValue
	var subnet6 basetypes.StringValue
	var tenants = types.MapNull(TenantsValue{}.Type(ctx))
	var vlanId basetypes.StringValue
	var vpnAccess = types.MapNull(VpnAccessValue{}.Type(ctx))

	if data.DisallowMistServices != nil {
		disallowMistServices = types.BoolValue(*data.DisallowMistServices)
	}
	if data.Gateway != nil {
		gateway = types.StringValue(*data.Gateway)
	}
	if data.Gateway6 != nil {
		gateway6 = types.StringValue(*data.Gateway6)
	}
	if data.Multicast != nil {
		multicast = MulticastSdkToTerraform(ctx, &diags, *data.Multicast)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.InternalAccess != nil {
		internalAccess = InternalAccessSdkToTerraform(ctx, &diags, *data.InternalAccess)
	}
	if data.InternetAccess != nil {
		internetAccess = InternetAccessSdkToTerraform(ctx, &diags, *data.InternetAccess)
	}
	if data.Isolation != nil {
		isolation = types.BoolValue(*data.Isolation)
	}
	name = types.StringValue(data.Name)
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.RoutedForNetworks != nil {
		routedForNetworks = mistutils.ListOfStringSdkToTerraform(data.RoutedForNetworks)
	}
	if data.Subnet != nil {
		subnet = types.StringValue(*data.Subnet)
	}
	if data.Subnet6 != nil {
		subnet6 = types.StringValue(*data.Subnet6)
	}
	if data.Tenants != nil && len(data.Tenants) > 0 {
		tenants = TenantSdkToTerraform(ctx, &diags, data.Tenants)
	}
	if data.VlanId != nil {
		vlanId = mistutils.VlanAsString(*data.VlanId)
	}
	if data.VpnAccess != nil && len(data.VpnAccess) > 0 {
		vpnAccess = VpnSdkToTerraform(ctx, &diags, data.VpnAccess)
	}

	state.DisallowMistServices = disallowMistServices
	state.Gateway = gateway
	state.Gateway6 = gateway6
	state.Multicast = multicast
	state.Id = id
	state.InternalAccess = internalAccess
	state.InternetAccess = internetAccess
	state.Isolation = isolation
	state.Name = name
	state.OrgId = orgId
	state.RoutedForNetworks = routedForNetworks
	state.Subnet = subnet
	state.Subnet6 = subnet6
	state.Tenants = tenants
	state.VlanId = vlanId
	state.VpnAccess = vpnAccess

	return state, diags
}
