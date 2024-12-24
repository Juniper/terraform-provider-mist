package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkDestinationNatProperty {
	data_map := make(map[string]models.NetworkDestinationNatProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(DestinationNatValue)
		data := models.NetworkDestinationNatProperty{}
		data.InternalIp = v_plan.InternalIp.ValueStringPointer()
		data.Name = v_plan.Name.ValueStringPointer()
		data.Port = models.ToPointer(int(v_plan.Port.ValueInt64()))
		data.WanName = v_plan.WanName.ValueStringPointer()
		data_map[k] = data
	}
	return data_map
}

func staticNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkStaticNatProperty {
	data_map := make(map[string]models.NetworkStaticNatProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(StaticNatValue)
		data := models.NetworkStaticNatProperty{}
		data.InternalIp = v_plan.InternalIp.ValueStringPointer()
		data.Name = v_plan.Name.ValueStringPointer()
		data.WanName = v_plan.WanName.ValueStringPointer()
		data_map[k] = data
	}
	return data_map
}

func sourceNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.NetworkSourceNat {
	data := models.NetworkSourceNat{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewSourceNatValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.ExternalIp = plan.ExternalIp.ValueStringPointer()
		}
	}
	return &data
}
