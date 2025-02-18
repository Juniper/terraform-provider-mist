package resource_device_gateway

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_network"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func mulitcastNetworksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.NetworkMulticast {
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
			var internalAccessInterface interface{} = plan.InternalAccess
			internalAccessTf := internalAccessInterface.(resource_org_network.InternalAccessValue)
			data.InternalAccess = resource_org_network.InternalAccessTerraformToSdk(internalAccessTf)
		}

		if !plan.InternetAccess.IsNull() && !plan.InternetAccess.IsUnknown() {
			var internetAccessInterface interface{} = plan.InternetAccess
			internetAccessTf := internetAccessInterface.(resource_org_network.InternetAccessValue)
			data.InternetAccess = resource_org_network.InternetAccessTerraformToSdk(internetAccessTf)
		}

		if plan.Isolation.ValueBoolPointer() != nil {
			data.Isolation = models.ToPointer(plan.Isolation.ValueBool())
		}
		if !plan.Multicast.IsNull() && !plan.Multicast.IsUnknown() {
			data.Multicast = mulitcastNetworksTerraformToSdk(ctx, diags, plan.Multicast)
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueString()
		}
		if !plan.RoutedForNetworks.IsNull() && !plan.RoutedForNetworks.IsUnknown() {
			data.RoutedForNetworks = misttransform.ListOfStringTerraformToSdk(plan.RoutedForNetworks)
		}
		if plan.Subnet.ValueStringPointer() != nil {
			data.Subnet = models.ToPointer(plan.Subnet.ValueString())
		}
		if plan.Subnet6.ValueStringPointer() != nil {
			data.Subnet6 = models.ToPointer(plan.Subnet.ValueString())
		}

		if !plan.Tenants.IsNull() && !plan.Tenants.IsUnknown() {
			data.Tenants = resource_org_network.TenantTerraformToSdk(plan.Tenants)
		}

		if plan.VlanId.ValueStringPointer() != nil {
			data.VlanId = models.ToPointer(models.VlanIdWithVariableContainer.FromString(plan.VlanId.ValueString()))
		}

		if !plan.VpnAccess.IsNull() && !plan.VpnAccess.IsUnknown() {
			data.VpnAccess = resource_org_network.VpnTerraformToSdk(ctx, diags, plan.VpnAccess)
		}

		dataList = append(dataList, data)
	}
	return dataList
}
