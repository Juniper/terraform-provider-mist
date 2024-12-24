package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func groupMulticastTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkMulticastGroup {
	data_map := make(map[string]models.NetworkMulticastGroup)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(GroupsValue)
		data := models.NetworkMulticastGroup{}
		data.RpIp = v_plan.RpIp.ValueStringPointer()
		data_map[k] = data
	}
	return data_map
}

func MulticastTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MulticastValue) *models.NetworkMulticast {
	data := models.NetworkMulticast{}

	if d.DisableIgmp.ValueBoolPointer() != nil {
		data.DisableIgmp = d.DisableIgmp.ValueBoolPointer()
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.Groups.IsNull() && !d.Groups.IsUnknown() {
		data.Groups = groupMulticastTerraformToSdk(ctx, diags, d.Groups)
	}

	return &data
}
