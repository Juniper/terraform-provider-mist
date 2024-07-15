package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	mist_network "github.com/Juniper/terraform-provider-mist/internal/resource_org_network"
)

func NetworksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m []models.Network) basetypes.ListValue {
	tflog.Debug(ctx, "NetworksSdkToTerraform")
	var data_list = []NetworksValue{}

	for _, d := range m {
		var disallow_mist_services basetypes.BoolValue = types.BoolValue(false)
		var gateway basetypes.StringValue
		var gateway6 basetypes.StringValue
		var internal_access basetypes.ObjectValue = types.ObjectNull(mist_network.InternalAccessValue{}.AttributeTypes(ctx))
		var internet_access basetypes.ObjectValue = types.ObjectNull(mist_network.InternetAccessValue{}.AttributeTypes(ctx))
		var isolation basetypes.BoolValue
		var name basetypes.StringValue
		var routed_for_networks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var subnet basetypes.StringValue
		var subnet6 basetypes.StringValue
		var tenants basetypes.MapValue = types.MapNull(mist_network.TenantsValue{}.Type(ctx))
		var vlan_id basetypes.Int64Value
		var vpn_access basetypes.MapValue = types.MapNull(mist_network.VpnAccessValue{}.Type(ctx))

		if d.DisallowMistServices != nil {
			disallow_mist_services = types.BoolValue(*d.DisallowMistServices)
		}
		if d.Gateway != nil {
			gateway = types.StringValue(*d.Gateway)
		}
		if d.Gateway6 != nil {
			gateway6 = types.StringValue(*d.Gateway6)
		}
		if d.InternalAccess != nil {
			internal_access, _ = mist_network.InternalAccessSdkToTerraform(ctx, diags, *d.InternalAccess).ToObjectValue(ctx)
		}
		if d.InternetAccess != nil {
			internet_access, _ = mist_network.InternetAccessSdkToTerraform(ctx, diags, *d.InternetAccess).ToObjectValue(ctx)
		}
		if d.Isolation != nil {
			isolation = types.BoolValue(*d.Isolation)
		}
		name = types.StringValue(d.Name)
		if d.RoutedForNetworks != nil {
			routed_for_networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.RoutedForNetworks)
		}
		if d.Subnet != nil {
			subnet = types.StringValue(*d.Subnet)
		}
		if d.Subnet6 != nil {
			subnet6 = types.StringValue(*d.Subnet6)
		}
		if d.Tenants != nil && len(d.Tenants) > 0 {
			tenants = mist_network.TenantSdkToTerraform(ctx, diags, d.Tenants)
		}
		if d.VlanId != nil {
			vlan_id = types.Int64Value(int64(*d.VlanId))
		}
		if d.VpnAccess != nil && len(d.VpnAccess) > 0 {
			vpn_access = mist_network.VpnSdkToTerraform(ctx, diags, d.VpnAccess)
		}

		data_map_attr_type := NetworksValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"disallow_mist_services": disallow_mist_services,
			"gateway":                gateway,
			"gateway6":               gateway6,
			"internal_access":        internal_access,
			"internet_access":        internet_access,
			"isolation":              isolation,
			"name":                   name,
			"routed_for_networks":    routed_for_networks,
			"subnet":                 subnet,
			"subnet6":                subnet6,
			"tenants":                tenants,
			"vlan_id":                vlan_id,
			"vpn_access":             vpn_access,
		}
		data, e := NewNetworksValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := NetworksValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
