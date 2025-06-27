package datasource_org_networks

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.Network, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := networkSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func networkSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Network) OrgNetworksValue {

	var createdTime basetypes.Float64Value
	var disallowMistServices basetypes.BoolValue
	var multicast = types.ObjectNull(MulticastValue{}.AttributeTypes(ctx))
	var gateway basetypes.StringValue
	var gateway6 basetypes.StringValue
	var id basetypes.StringValue
	var internalAccess = types.ObjectNull(InternalAccessValue{}.AttributeTypes(ctx))
	var internetAccess = types.ObjectNull(InternetAccessValue{}.AttributeTypes(ctx))
	var isolation basetypes.BoolValue
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var routedForNetworks = types.ListValueMust(types.StringType, []attr.Value{})
	var subnet basetypes.StringValue
	var subnet6 basetypes.StringValue
	var tenants = types.MapNull(TenantsValue{}.Type(ctx))
	var vlanId basetypes.StringValue
	var vpnAccess = types.MapNull(VpnAccessValue{}.Type(ctx))

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.DisallowMistServices != nil {
		disallowMistServices = types.BoolValue(*d.DisallowMistServices)
	}
	if d.Gateway != nil {
		gateway = types.StringValue(*d.Gateway)
	}
	if d.Gateway6 != nil {
		gateway6 = types.StringValue(*d.Gateway6)
	}
	if d.Multicast != nil {
		multicast = MulticastSdkToTerraform(ctx, diags, *d.Multicast)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.InternalAccess != nil {
		internalAccess = InternalAccessSdkToTerraform(ctx, diags, *d.InternalAccess)
	}
	if d.InternetAccess != nil {
		internetAccess = InternetAccessSdkToTerraform(ctx, diags, *d.InternetAccess)
	}
	if d.Isolation != nil {
		isolation = types.BoolValue(*d.Isolation)
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	name = types.StringValue(d.Name)
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.RoutedForNetworks != nil {
		routedForNetworks = mistutils.ListOfStringSdkToTerraform(d.RoutedForNetworks)
	}
	if d.Subnet != nil {
		subnet = types.StringValue(*d.Subnet)
	}
	if d.Subnet6 != nil {
		subnet = types.StringValue(*d.Subnet6)
	}
	if len(d.Tenants) > 0 {
		tenants = TenantSdkToTerraform(ctx, diags, d.Tenants)
	}
	if d.VlanId != nil {
		vlanId = mistutils.VlanAsString(*d.VlanId)
	}
	if len(d.VpnAccess) > 0 {
		vpnAccess = VpnSdkToTerraform(ctx, diags, d.VpnAccess)
	}

	dataMapValue := map[string]attr.Value{
		"created_time":           createdTime,
		"disallow_mist_services": disallowMistServices,
		"gateway":                gateway,
		"gateway6":               gateway6,
		"multicast":              multicast,
		"id":                     id,
		"internal_access":        internalAccess,
		"internet_access":        internetAccess,
		"isolation":              isolation,
		"modified_time":          modifiedTime,
		"name":                   name,
		"org_id":                 orgId,
		"routed_for_networks":    routedForNetworks,
		"subnet":                 subnet,
		"subnet6":                subnet6,
		"tenants":                tenants,
		"vlan_id":                vlanId,
		"vpn_access":             vpnAccess,
	}
	data, e := NewOrgNetworksValue(OrgNetworksValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
