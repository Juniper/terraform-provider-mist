package resource_device_gateway

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func multicastNetworksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.NetworkMulticast {
	data := models.NetworkMulticast{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewMulticastValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.DisableIgmp.ValueBoolPointer() != nil {
				data.DisableIgmp = plan.DisableIgmp.ValueBoolPointer()
			}
			if plan.Enabled.ValueBoolPointer() != nil {
				data.Enabled = plan.Enabled.ValueBoolPointer()
			}
			if !plan.Groups.IsNull() && !plan.Groups.IsUnknown() {
				groupMap := make(map[string]models.NetworkMulticastGroup)
				for k, v := range plan.Groups.Elements() {
					var vInterface interface{} = v
					p := vInterface.(GroupsValue)
					g := models.NetworkMulticastGroup{}
					if p.RpIp.ValueStringPointer() != nil {
						g.RpIp = p.RpIp.ValueStringPointer()
					}
					groupMap[k] = g
				}
				data.Groups = groupMap
			}
		}
	}
	return &data
}

func internalAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.NetworkInternalAccess {
	data := models.NetworkInternalAccess{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewInternalAccessValue(InternalAccessValue{}.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
			return &data
		}
		data.Enabled = plan.Enabled.ValueBoolPointer()
	}
	return &data
}

func destinationNatInternetAccessTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkInternetAccessDestinationNatProperty {
	dataMap := make(map[string]models.NetworkInternetAccessDestinationNatProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(InternetAccessDestinationNatValue)
		data := models.NetworkInternetAccessDestinationNatProperty{}
		data.InternalIp = vPlan.InternalIp.ValueStringPointer()
		data.Name = vPlan.Name.ValueStringPointer()
		data.Port = vPlan.Port.ValueStringPointer()
		data.WanName = vPlan.WanName.ValueStringPointer()
		dataMap[k] = data
	}
	return dataMap
}

func staticNatInternetAccessTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkInternetAccessStaticNatProperty {
	dataMap := make(map[string]models.NetworkInternetAccessStaticNatProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(InternetAccessStaticNatValue)
		data := models.NetworkInternetAccessStaticNatProperty{}
		data.InternalIp = vPlan.InternalIp.ValueStringPointer()
		data.Name = vPlan.Name.ValueStringPointer()
		data.WanName = vPlan.WanName.ValueStringPointer()
		dataMap[k] = data
	}
	return dataMap
}

func internetAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.NetworkInternetAccess {
	data := models.NetworkInternetAccess{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewInternetAccessValue(InternetAccessValue{}.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
			return &data
		}
		if !plan.CreateSimpleServicePolicy.IsNull() && !plan.CreateSimpleServicePolicy.IsUnknown() {
			data.CreateSimpleServicePolicy = plan.CreateSimpleServicePolicy.ValueBoolPointer()
		}
		if !plan.InternetAccessDestinationNat.IsNull() && !plan.InternetAccessDestinationNat.IsUnknown() {
			data.DestinationNat = destinationNatInternetAccessTerraformToSdk(plan.InternetAccessDestinationNat)
		}
		if !plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
			data.Enabled = plan.Enabled.ValueBoolPointer()
		}
		if !plan.InternetAccessStaticNat.IsNull() && !plan.InternetAccessStaticNat.IsUnknown() {
			data.StaticNat = staticNatInternetAccessTerraformToSdk(plan.InternetAccessStaticNat)
		}
		if !plan.Restricted.IsNull() && !plan.Restricted.IsUnknown() {
			data.Restricted = plan.Restricted.ValueBoolPointer()
		}
	}
	return &data
}

func destinationNatVpnTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkVpnAccessDestinationNatProperty {
	dataMap := make(map[string]models.NetworkVpnAccessDestinationNatProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(VpnAccessDestinationNatValue)
		data := models.NetworkVpnAccessDestinationNatProperty{}
		data.InternalIp = vPlan.InternalIp.ValueStringPointer()
		data.Name = vPlan.Name.ValueStringPointer()
		data.Port = vPlan.Port.ValueStringPointer()
		dataMap[k] = data
	}
	return dataMap
}

func staticNatVpnTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkVpnAccessStaticNatProperty {
	dataMap := make(map[string]models.NetworkVpnAccessStaticNatProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(VpnAccessStaticNatValue)
		data := models.NetworkVpnAccessStaticNatProperty{}
		data.InternalIp = vPlan.InternalIp.ValueStringPointer()
		data.Name = vPlan.Name.ValueStringPointer()
		dataMap[k] = data
	}
	return dataMap
}

func sourceNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.NetworkSourceNat {
	data := models.NetworkSourceNat{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewSourceNatValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.ExternalIp = plan.ExternalIp.ValueStringPointer()
		}
	}
	return &data
}

func vpnTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkVpnAccessConfig {
	dataMap := make(map[string]models.NetworkVpnAccessConfig)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(VpnAccessValue)

		data := models.NetworkVpnAccessConfig{}
		if plan.AdvertisedSubnet.ValueStringPointer() != nil {
			data.AdvertisedSubnet = plan.AdvertisedSubnet.ValueStringPointer()
		}
		if plan.AllowPing.ValueBoolPointer() != nil {
			data.AllowPing = plan.AllowPing.ValueBoolPointer()
		}
		if !plan.VpnAccessDestinationNat.IsNull() && !plan.VpnAccessDestinationNat.IsUnknown() {
			data.DestinationNat = destinationNatVpnTerraformToSdk(plan.VpnAccessDestinationNat)
		}
		if plan.NatPool.ValueStringPointer() != nil {
			data.NatPool = plan.NatPool.ValueStringPointer()
		}
		if plan.NoReadvertiseToLanBgp.ValueBoolPointer() != nil {
			data.NoReadvertiseToLanBgp = plan.NoReadvertiseToLanBgp.ValueBoolPointer()
		}
		if plan.NoReadvertiseToLanOspf.ValueBoolPointer() != nil {
			data.NoReadvertiseToLanOspf = plan.NoReadvertiseToLanOspf.ValueBoolPointer()
		}
		if plan.NoReadvertiseToOverlay.ValueBoolPointer() != nil {
			data.NoReadvertiseToOverlay = plan.NoReadvertiseToOverlay.ValueBoolPointer()
		}
		if !plan.OtherVrfs.IsNull() && !plan.OtherVrfs.IsUnknown() {
			data.OtherVrfs = mistutils.ListOfStringTerraformToSdk(plan.OtherVrfs)
		}
		if plan.Routed.ValueBoolPointer() != nil {
			data.Routed = plan.Routed.ValueBoolPointer()
		}
		if !plan.SourceNat.IsNull() && !plan.SourceNat.IsUnknown() {
			data.SourceNat = sourceNatTerraformToSdk(ctx, diags, plan.SourceNat)
		}
		if !plan.VpnAccessStaticNat.IsNull() && !plan.VpnAccessStaticNat.IsUnknown() {
			data.StaticNat = staticNatVpnTerraformToSdk(plan.VpnAccessStaticNat)
		}
		if plan.SummarizedSubnet.ValueStringPointer() != nil {
			data.SummarizedSubnet = plan.SummarizedSubnet.ValueStringPointer()
		}
		if plan.SummarizedSubnetToLanBgp.ValueStringPointer() != nil {
			data.SummarizedSubnetToLanBgp = plan.SummarizedSubnetToLanBgp.ValueStringPointer()
		}
		if plan.SummarizedSubnetToLanOspf.ValueStringPointer() != nil {
			data.SummarizedSubnetToLanOspf = plan.SummarizedSubnetToLanOspf.ValueStringPointer()
		}

		dataMap[k] = data
	}
	return dataMap
}

func tenantTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkTenant {
	dataMap := make(map[string]models.NetworkTenant)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(TenantsValue)
		data := models.NetworkTenant{}
		data.Addresses = mistutils.ListOfStringTerraformToSdk(vPlan.Addresses)
		dataMap[k] = data
	}
	return dataMap
}

func networksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Network {
	var dataList []models.Network
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(NetworksValue)
		data := models.Network{}

		if plan.DisallowMistServices.ValueBoolPointer() != nil {
			data.DisallowMistServices = models.ToPointer(plan.DisallowMistServices.ValueBool())
		}
		if plan.Gateway.ValueStringPointer() != nil {
			data.Gateway = models.ToPointer(plan.Gateway.ValueString())
		}
		if plan.Gateway6.ValueStringPointer() != nil {
			data.Gateway6 = models.ToPointer(plan.Gateway6.ValueString())
		}

		if !plan.InternalAccess.IsNull() && !plan.InternalAccess.IsUnknown() {
			data.InternalAccess = internalAccessTerraformToSdk(ctx, diags, plan.InternalAccess)
		}

		if !plan.InternetAccess.IsNull() && !plan.InternetAccess.IsUnknown() {
			data.InternetAccess = internetAccessTerraformToSdk(ctx, diags, plan.InternetAccess)
		}

		if plan.Isolation.ValueBoolPointer() != nil {
			data.Isolation = models.ToPointer(plan.Isolation.ValueBool())
		}
		if !plan.Multicast.IsNull() && !plan.Multicast.IsUnknown() {
			data.Multicast = multicastNetworksTerraformToSdk(ctx, diags, plan.Multicast)
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueString()
		}
		if !plan.RoutedForNetworks.IsNull() && !plan.RoutedForNetworks.IsUnknown() {
			data.RoutedForNetworks = mistutils.ListOfStringTerraformToSdk(plan.RoutedForNetworks)
		}
		if plan.Subnet.ValueStringPointer() != nil {
			data.Subnet = models.ToPointer(plan.Subnet.ValueString())
		}
		if plan.Subnet6.ValueStringPointer() != nil {
			data.Subnet6 = models.ToPointer(plan.Subnet6.ValueString())
		}

		if !plan.Tenants.IsNull() && !plan.Tenants.IsUnknown() {
			data.Tenants = tenantTerraformToSdk(plan.Tenants)
		}

		if plan.VlanId.ValueStringPointer() != nil {
			data.VlanId = models.ToPointer(models.VlanIdWithVariableContainer.FromString(plan.VlanId.ValueString()))
		}

		if !plan.VpnAccess.IsNull() && !plan.VpnAccess.IsUnknown() {
			data.VpnAccess = vpnTerraformToSdk(ctx, diags, plan.VpnAccess)
		}

		dataList = append(dataList, data)
	}
	return dataList
}
