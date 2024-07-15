package resource_org_network

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data models.Network) (OrgNetworkModel, diag.Diagnostics) {
	var state OrgNetworkModel
	var diags diag.Diagnostics

	var disallow_mist_services basetypes.BoolValue
	var gateway basetypes.StringValue
	var gateway6 basetypes.StringValue
	var id basetypes.StringValue
	var internal_access InternalAccessValue = NewInternalAccessValueNull()
	var internet_access InternetAccessValue = NewInternetAccessValueNull()
	var isolation basetypes.BoolValue
	var name basetypes.StringValue
	var org_id basetypes.StringValue
	var routed_for_networks basetypes.ListValue = types.ListNull(types.StringType)
	var subnet basetypes.StringValue
	var subnet6 basetypes.StringValue
	var tenants basetypes.MapValue = types.MapNull(TenantsValue{}.Type(ctx))
	var vlan_id basetypes.Int64Value
	var vpn_access basetypes.MapValue = types.MapNull(VpnAccessValue{}.Type(ctx))

	if data.DisallowMistServices != nil {
		disallow_mist_services = types.BoolValue(*data.DisallowMistServices)
	}
	if data.Gateway != nil {
		gateway = types.StringValue(*data.Gateway)
	}
	if data.Gateway6 != nil {
		gateway6 = types.StringValue(*data.Gateway6)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.InternalAccess != nil {
		internal_access = InternalAccessSdkToTerraform(ctx, &diags, *data.InternalAccess)
	}
	if data.InternetAccess != nil {
		internet_access = InternetAccessSdkToTerraform(ctx, &diags, *data.InternetAccess)
	}
	if data.Isolation != nil {
		isolation = types.BoolValue(*data.Isolation)
	}
	name = types.StringValue(data.Name)
	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}
	if data.RoutedForNetworks != nil {
		routed_for_networks = mist_transform.ListOfStringSdkToTerraform(ctx, data.RoutedForNetworks)
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
		vlan_id = types.Int64Value(int64(*data.VlanId))
	}
	if data.VpnAccess != nil && len(data.VpnAccess) > 0 {
		vpn_access = VpnSdkToTerraform(ctx, &diags, data.VpnAccess)
	}

	state.DisallowMistServices = disallow_mist_services
	state.Gateway = gateway
	state.Gateway6 = gateway6
	state.Id = id
	state.InternalAccess = internal_access
	state.InternetAccess = internet_access
	state.Isolation = isolation
	state.Name = name
	state.OrgId = org_id
	state.RoutedForNetworks = routed_for_networks
	state.Subnet = subnet
	state.Subnet6 = subnet6
	state.Tenants = tenants
	state.VlanId = vlan_id
	state.VpnAccess = vpn_access

	return state, diags
}
