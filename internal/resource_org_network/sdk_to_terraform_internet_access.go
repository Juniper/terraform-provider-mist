package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatInternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkInternetAccessDestinationNatProperty) basetypes.MapValue {
	stateValueMapValue := make(map[string]attr.Value)
	for k, v := range d {
		var internalIp basetypes.StringValue
		var name basetypes.StringValue
		var port basetypes.StringValue
		var wanName basetypes.StringValue

		if v.InternalIp != nil {
			internalIp = types.StringValue(*v.InternalIp)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.Port != nil {
			port = types.StringValue(*v.Port)
		}
		if v.WanName != nil {
			wanName = types.StringValue(*v.WanName)
		}

		stateValueMapAttrValue := map[string]attr.Value{
			"internal_ip": internalIp,
			"name":        name,
			"port":        port,
			"wan_name":    wanName,
		}
		n, e := NewInternetAccessDestinationNatValue(InternetAccessDestinationNatValue{}.AttributeTypes(ctx), stateValueMapAttrValue)
		diags.Append(e...)

		stateValueMapValue[k] = n
	}
	stateResultMap, e := types.MapValueFrom(ctx, InternetAccessDestinationNatValue{}.Type(ctx), stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func staticNatInternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkInternetAccessStaticNatProperty) basetypes.MapValue {
	stateValueMapValue := make(map[string]attr.Value)
	for k, v := range d {
		var internalIp basetypes.StringValue
		var name basetypes.StringValue
		var wanName basetypes.StringValue

		if v.InternalIp != nil {
			internalIp = types.StringValue(*v.InternalIp)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.WanName != nil {
			wanName = types.StringValue(*v.WanName)
		}

		stateValueMapAttrValue := map[string]attr.Value{
			"internal_ip": internalIp,
			"name":        name,
			"wan_name":    wanName,
		}
		n, e := NewInternetAccessStaticNatValue(InternetAccessStaticNatValue{}.AttributeTypes(ctx), stateValueMapAttrValue)
		diags.Append(e...)

		stateValueMapValue[k] = n
	}
	stateResultMap, e := types.MapValueFrom(ctx, InternetAccessStaticNatValue{}.Type(ctx), stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func InternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkInternetAccess) InternetAccessValue {
	var createSimpleServicePolicy = types.BoolValue(false)
	var destinationNat = types.MapNull(InternetAccessDestinationNatValue{}.Type(ctx))
	var enabled basetypes.BoolValue
	var restricted = types.BoolValue(false)
	var staticNac = types.MapNull(InternetAccessStaticNatValue{}.Type(ctx))

	if d.CreateSimpleServicePolicy != nil {
		createSimpleServicePolicy = types.BoolValue(*d.CreateSimpleServicePolicy)
	}
	if d.DestinationNat != nil && len(d.DestinationNat) > 0 {
		destinationNat = destinationNatInternetAccessSdkToTerraform(ctx, diags, d.DestinationNat)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Restricted != nil {
		restricted = types.BoolValue(*d.Restricted)
	}
	if d.StaticNat != nil && len(d.StaticNat) > 0 {
		staticNac = staticNatInternetAccessSdkToTerraform(ctx, diags, d.StaticNat)
	}

	dataMapValue := map[string]attr.Value{
		"create_simple_service_policy": createSimpleServicePolicy,
		"destination_nat":              destinationNat,
		"enabled":                      enabled,
		"restricted":                   restricted,
		"static_nat":                   staticNac,
	}
	data, e := NewInternetAccessValue(InternetAccessValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
