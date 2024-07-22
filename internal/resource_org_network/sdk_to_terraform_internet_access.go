package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func InternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkInternetAccess) InternetAccessValue {
	var create_simple_service_policy basetypes.BoolValue = types.BoolValue(false)
	var destination_nat basetypes.MapValue = types.MapNull(DestinationNatValue{}.Type(ctx))
	var enabled basetypes.BoolValue
	var restricted basetypes.BoolValue = types.BoolValue(false)
	var static_nac basetypes.MapValue = types.MapNull(StaticNatValue{}.Type(ctx))

	if d.CreateSimpleServicePolicy != nil {
		create_simple_service_policy = types.BoolValue(*d.CreateSimpleServicePolicy)
	}
	if d.DestinationNat != nil && len(d.DestinationNat) > 0 {
		destination_nat = destinationNatSdkToTerraform(ctx, diags, d.DestinationNat)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Restricted != nil {
		restricted = types.BoolValue(*d.Restricted)
	}
	if d.StaticNat != nil && len(d.StaticNat) > 0 {
		static_nac = staticNatSdkToTerraform(ctx, diags, d.StaticNat)
	}

	data_map_attr_type := InternetAccessValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"create_simple_service_policy": create_simple_service_policy,
		"destination_nat":              destination_nat,
		"enabled":                      enabled,
		"restricted":                   restricted,
		"static_nat":                   static_nac,
	}
	data, e := NewInternetAccessValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
