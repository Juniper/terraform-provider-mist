package datasource_org_nactags

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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

	var allowUsermacOverride types.Bool
	var createdTime basetypes.Float64Value
	var egressVlanNames = types.ListNull(types.StringType)
	var gbpTag types.String
	var id types.String
	var match types.String
	var matchAll types.Bool
	var modifiedTime basetypes.Float64Value
	var name types.String
	var orgId types.String
	var radiusAttrs = types.ListNull(types.StringType)
	var radiusGroup types.String
	var radiusVendorAttrs = types.ListNull(types.StringType)
	var sessionTimeout types.Int64
	var typeNactag types.String
	var usernameAttr types.String
	var values = types.ListNull(types.StringType)
	var vlan types.String

	if d.AllowUsermacOverride != nil {
		allowUsermacOverride = types.BoolValue(*d.AllowUsermacOverride)
	}
	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.EgressVlanNames != nil {
		egressVlanNames = misttransform.ListOfStringSdkToTerraform(d.EgressVlanNames)
	}
	if d.GbpTag != nil {
		gbpTag = types.StringValue(d.GbpTag.String())
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Match != nil {
		match = types.StringValue(string(*d.Match))
	}
	if d.MatchAll != nil {
		matchAll = types.BoolValue(*d.MatchAll)
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}

	name = types.StringValue(d.Name)

	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.RadiusAttrs != nil {
		radiusAttrs = misttransform.ListOfStringSdkToTerraform(d.RadiusAttrs)
	}
	if d.RadiusGroup != nil {
		radiusGroup = types.StringValue(*d.RadiusGroup)
	}
	if d.RadiusVendorAttrs != nil {
		radiusVendorAttrs = misttransform.ListOfStringSdkToTerraform(d.RadiusVendorAttrs)
	}
	if d.SessionTimeout != nil {
		sessionTimeout = types.Int64Value(int64(*d.SessionTimeout))
	}

	typeNactag = types.StringValue(string(d.Type))

	if d.UsernameAttr != nil {
		usernameAttr = types.StringValue(string(*d.UsernameAttr))
	}
	if d.Values != nil {
		values = misttransform.ListOfStringSdkToTerraform(d.Values)
	}
	if d.Vlan != nil {
		vlan = types.StringValue(*d.Vlan)
	}

	dataMapValue := map[string]attr.Value{
		"allow_usermac_override": allowUsermacOverride,
		"created_time":           createdTime,
		"egress_vlan_names":      egressVlanNames,
		"gbp_tag":                gbpTag,
		"id":                     id,
		"match":                  match,
		"match_all":              matchAll,
		"modified_time":          modifiedTime,
		"name":                   name,
		"org_id":                 orgId,
		"radius_attrs":           radiusAttrs,
		"radius_group":           radiusGroup,
		"radius_vendor_attrs":    radiusVendorAttrs,
		"session_timeout":        sessionTimeout,
		"type":                   typeNactag,
		"username_attr":          usernameAttr,
		"values":                 values,
		"vlan":                   vlan,
	}
	data, e := NewOrgNactagsValue(OrgNactagsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
