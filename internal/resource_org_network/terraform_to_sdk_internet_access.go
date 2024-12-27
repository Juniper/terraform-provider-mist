package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatInternetAccesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkInternetAccessDestinationNatProperty {
	data_map := make(map[string]models.NetworkInternetAccessDestinationNatProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(InternetAccessDestinationNatValue)
		data := models.NetworkInternetAccessDestinationNatProperty{}
		data.InternalIp = v_plan.InternalIp.ValueStringPointer()
		data.Name = v_plan.Name.ValueStringPointer()
		data.Port = v_plan.Port.ValueStringPointer()
		data.WanName = v_plan.WanName.ValueStringPointer()
		data_map[k] = data
	}
	return data_map
}

func staticNatInternetAccesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkInternetAccessStaticNatProperty {
	data_map := make(map[string]models.NetworkInternetAccessStaticNatProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(InternetAccessStaticNatValue)
		data := models.NetworkInternetAccessStaticNatProperty{}
		data.InternalIp = v_plan.InternalIp.ValueStringPointer()
		data.Name = v_plan.Name.ValueStringPointer()
		data.WanName = v_plan.WanName.ValueStringPointer()
		data_map[k] = data
	}
	return data_map
}

func InternetAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d InternetAccessValue) *models.NetworkInternetAccess {
	data := models.NetworkInternetAccess{}

	if !d.CreateSimpleServicePolicy.IsNull() && !d.CreateSimpleServicePolicy.IsUnknown() {
		data.CreateSimpleServicePolicy = d.CreateSimpleServicePolicy.ValueBoolPointer()
	}
	if !d.InternetAccessDestinationNat.IsNull() && !d.InternetAccessDestinationNat.IsUnknown() {
		data.DestinationNat = destinationNatInternetAccesTerraformToSdk(ctx, diags, d.InternetAccessDestinationNat)
	}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.InternetAccessStaticNat.IsNull() && !d.InternetAccessStaticNat.IsUnknown() {
		data.StaticNat = staticNatInternetAccesTerraformToSdk(ctx, diags, d.InternetAccessStaticNat)
	}
	if !d.Restricted.IsNull() && !d.Restricted.IsUnknown() {
		data.Restricted = d.Restricted.ValueBoolPointer()
	}

	return &data
}
