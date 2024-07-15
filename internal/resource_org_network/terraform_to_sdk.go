package resource_org_network

import (
	"context"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgNetworkModel) (*models.Network, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Network{}
	unset := make(map[string]interface{})

	if plan.Name.ValueStringPointer() != nil {
		data.Name = plan.Name.ValueString()
	} else {
		unset["-ap_updown_threshold"] = ""
	}

	if plan.DisallowMistServices.ValueBoolPointer() != nil {
		data.DisallowMistServices = plan.DisallowMistServices.ValueBoolPointer()
	} else {
		unset["-disallow_mist_services"] = ""
	}

	if plan.Gateway.ValueStringPointer() != nil {
		data.Gateway = plan.Gateway.ValueStringPointer()
	} else {
		unset["-gateway"] = ""
	}

	if plan.Gateway6.ValueStringPointer() != nil {
		data.Gateway6 = plan.Gateway6.ValueStringPointer()
	} else {
		unset["-gateway6"] = ""
	}

	internal_access := InternalAccessTerraformToSdk(ctx, &diags, plan.InternalAccess)
	if !plan.InternalAccess.IsNull() && !plan.InternalAccess.IsUnknown() {
		data.InternalAccess = internal_access
	} else {
		unset["-internal_access"] = ""
	}

	internet_access := InternetAccessTerraformToSdk(ctx, &diags, plan.InternetAccess)
	if !plan.InternetAccess.IsNull() && !plan.InternetAccess.IsUnknown() {
		data.InternetAccess = internet_access
	} else {
		unset["-internet_access"] = ""
	}

	if plan.Isolation.ValueBoolPointer() != nil {
		data.Isolation = plan.Isolation.ValueBoolPointer()
	} else {
		unset["-isolation"] = ""
	}

	if !plan.RoutedForNetworks.IsNull() && !plan.RoutedForNetworks.IsUnknown() {
		data.RoutedForNetworks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RoutedForNetworks)
	} else {
		unset["-routed_for_networks"] = ""
	}

	if plan.Subnet.ValueStringPointer() != nil {
		data.Subnet = plan.Subnet.ValueStringPointer()
	} else {
		unset["-subnet"] = ""
	}

	if plan.Subnet6.ValueStringPointer() != nil {
		data.Subnet6 = plan.Subnet6.ValueStringPointer()
	} else {
		unset["-subnet6"] = ""
	}

	if !plan.Tenants.IsNull() && !plan.Tenants.IsUnknown() {
		data.Tenants = TenantTerraformToSdk(ctx, &diags, plan.Tenants)
	} else {
		unset["-tenants"] = ""
	}

	if plan.VlanId.ValueInt64Pointer() != nil {
		data.VlanId = models.ToPointer(int(plan.VlanId.ValueInt64()))
	} else {
		unset["-vlan_id"] = ""
	}

	if !plan.VpnAccess.IsNull() && !plan.VpnAccess.IsUnknown() {
		data.VpnAccess = VpnTerraformToSdk(ctx, &diags, plan.VpnAccess)
	} else {
		unset["-vpn_access"] = ""
	}

	data.AdditionalProperties = unset
	return &data, diags
}
