package datasource_org_nactags

import (
	"context"
	"math/big"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.NacTag, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := nactagSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func nactagSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NacTag) OrgNactagsValue {

	var allow_usermac_override types.Bool
	var created_time basetypes.NumberValue
	var egress_vlan_names types.List = types.ListNull(types.StringType)
	var gbp_tag types.Int64
	var id types.String
	var match types.String
	var match_all types.Bool
	var modified_time basetypes.NumberValue
	var name types.String
	var org_id types.String
	var radius_attrs types.List = types.ListNull(types.StringType)
	var radius_group types.String
	var radius_vendor_attrs types.List = types.ListNull(types.StringType)
	var session_timeout types.Int64
	var type_nactag types.String
	var username_attr types.String
	var values types.List = types.ListNull(types.StringType)
	var vlan types.String

	if d.AllowUsermacOverride != nil {
		allow_usermac_override = types.BoolValue(*d.AllowUsermacOverride)
	}
	if d.CreatedTime != nil {
		created_time = types.NumberValue(big.NewFloat(*d.CreatedTime))
	}
	if d.EgressVlanNames != nil {
		egress_vlan_names = mist_transform.ListOfStringSdkToTerraform(ctx, d.EgressVlanNames)
	}
	if d.GbpTag != nil {
		gbp_tag = types.Int64Value(int64(*d.GbpTag))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Match != nil {
		match = types.StringValue(string(*d.Match))
	}
	if d.MatchAll != nil {
		match_all = types.BoolValue(*d.MatchAll)
	}
	if d.ModifiedTime != nil {
		modified_time = types.NumberValue(big.NewFloat(*d.ModifiedTime))
	}

	name = types.StringValue(d.Name)

	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.RadiusAttrs != nil {
		radius_attrs = mist_transform.ListOfStringSdkToTerraform(ctx, d.RadiusAttrs)
	}
	if d.RadiusGroup != nil {
		radius_group = types.StringValue(*d.RadiusGroup)
	}
	if d.RadiusVendorAttrs != nil {
		radius_vendor_attrs = mist_transform.ListOfStringSdkToTerraform(ctx, d.RadiusVendorAttrs)
	}
	if d.SessionTimeout != nil {
		session_timeout = types.Int64Value(int64(*d.SessionTimeout))
	}

	type_nactag = types.StringValue(string(d.Type))

	if d.UsernameAttr != nil {
		username_attr = types.StringValue(string(*d.UsernameAttr))
	}
	if d.Values != nil {
		values = mist_transform.ListOfStringSdkToTerraform(ctx, d.Values)
	}
	if d.Vlan != nil {
		vlan = types.StringValue(*d.Vlan)
	}

	data_map_attr_type := OrgNactagsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_usermac_override": allow_usermac_override,
		"created_time":           created_time,
		"egress_vlan_names":      egress_vlan_names,
		"gbp_tag":                gbp_tag,
		"id":                     id,
		"match":                  match,
		"match_all":              match_all,
		"modified_time":          modified_time,
		"name":                   name,
		"org_id":                 org_id,
		"radius_attrs":           radius_attrs,
		"radius_group":           radius_group,
		"radius_vendor_attrs":    radius_vendor_attrs,
		"session_timeout":        session_timeout,
		"type":                   type_nactag,
		"username_attr":          username_attr,
		"values":                 values,
		"vlan":                   vlan,
	}
	data, e := NewOrgNactagsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
