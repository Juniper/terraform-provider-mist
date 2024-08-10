package resource_device_gateway

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_network"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func networksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.Network {
	var data_list []models.Network
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NetworksValue)
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
			var internal_access_interface interface{} = plan.InternalAccess
			internal_access_tf := internal_access_interface.(resource_org_network.InternalAccessValue)
			data.InternalAccess = resource_org_network.InternalAccessTerraformToSdk(ctx, diags, internal_access_tf)
		}

		if !plan.InternetAccess.IsNull() && !plan.InternetAccess.IsUnknown() {
			var internet_access_interface interface{} = plan.InternetAccess
			internet_access_tf := internet_access_interface.(resource_org_network.InternetAccessValue)
			data.InternetAccess = resource_org_network.InternetAccessTerraformToSdk(ctx, diags, internet_access_tf)
		}

		if plan.Isolation.ValueBoolPointer() != nil {
			data.Isolation = models.ToPointer(plan.Isolation.ValueBool())
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueString()
		}
		if !plan.RoutedForNetworks.IsNull() && !plan.RoutedForNetworks.IsUnknown() {
			data.RoutedForNetworks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RoutedForNetworks)
		}
		if plan.Subnet.ValueStringPointer() != nil {
			data.Subnet = models.ToPointer(plan.Subnet.ValueString())
		}
		if plan.Subnet6.ValueStringPointer() != nil {
			data.Subnet6 = models.ToPointer(plan.Subnet.ValueString())
		}

		if !plan.Tenants.IsNull() && !plan.Tenants.IsUnknown() {
			data.Tenants = resource_org_network.TenantTerraformToSdk(ctx, diags, plan.Tenants)
		}

		if plan.VlanId.ValueStringPointer() != nil {
			data.VlanId = models.ToPointer(models.VlanIdWithVariableContainer.FromString(plan.VlanId.ValueString()))
		}

		if !plan.VpnAccess.IsNull() && !plan.VpnAccess.IsUnknown() {
			data.VpnAccess = resource_org_network.VpnTerraformToSdk(ctx, diags, plan.VpnAccess)
		}

		data_list = append(data_list, data)
	}
	return data_list
}
