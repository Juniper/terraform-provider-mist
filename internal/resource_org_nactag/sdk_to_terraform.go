package resource_org_nactag

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.NacTag) (OrgNactagModel, diag.Diagnostics) {
	var state OrgNactagModel
	var diags diag.Diagnostics

	var allow_usermac_override types.Bool
	var egress_vlan_names types.List = types.ListNull(types.StringType)
	var gbp_tag types.Int64
	var id types.String
	var match types.String
	var match_all types.Bool
	var name types.String
	var org_id types.String
	var radius_attrs types.List = types.ListNull(types.StringType)
	var radius_group types.String
	var radius_vendor_attrs types.List = types.ListNull(types.StringType)
	var session_timeout types.Int64
	var type_nactag types.String
	var values types.List = types.ListNull(types.StringType)
	var vlan types.String

	if data.AllowUsermacOverride != nil {
		allow_usermac_override = types.BoolValue(*data.AllowUsermacOverride)
	}
	if data.EgressVlanNames != nil {
		egress_vlan_names = mist_transform.ListOfStringSdkToTerraform(ctx, data.EgressVlanNames)
	}
	if data.GbpTag != nil {
		gbp_tag = types.Int64Value(int64(*data.GbpTag))
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.Match != nil {
		match = types.StringValue(string(*data.Match))
	}
	if data.MatchAll != nil {
		match_all = types.BoolValue(*data.MatchAll)
	}

	name = types.StringValue(data.Name)

	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}
	if data.RadiusAttrs != nil {
		radius_attrs = mist_transform.ListOfStringSdkToTerraform(ctx, data.RadiusAttrs)
	}
	if data.RadiusGroup != nil {
		radius_group = types.StringValue(*data.RadiusGroup)
	}
	if data.RadiusVendorAttrs != nil {
		radius_vendor_attrs = mist_transform.ListOfStringSdkToTerraform(ctx, data.RadiusVendorAttrs)
	}
	if data.SessionTimeout != nil {
		session_timeout = types.Int64Value(int64(*data.SessionTimeout))
	}
	type_nactag = types.StringValue(string(data.Type))
	if data.Values != nil {
		values = mist_transform.ListOfStringSdkToTerraform(ctx, data.Values)
	}
	if data.Vlan != nil {
		vlan = types.StringValue(*data.Vlan)
	}

	state.AllowUsermacOverride = allow_usermac_override
	state.EgressVlanNames = egress_vlan_names
	state.GbpTag = gbp_tag
	state.Id = id
	state.Match = match
	state.MatchAll = match_all
	state.Name = name
	state.OrgId = org_id
	state.RadiusAttrs = radius_attrs
	state.RadiusGroup = radius_group
	state.RadiusVendorAttrs = radius_vendor_attrs
	state.SessionTimeout = session_timeout
	state.Type = type_nactag
	state.Values = values
	state.Vlan = vlan

	return state, diags
}
