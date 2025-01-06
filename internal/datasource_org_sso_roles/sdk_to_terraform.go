package datasource_org_sso_roles

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.SsoRoleOrg, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := ssoRoleSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func ssoRoleSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SsoRoleOrg) OrgSsoRolesValue {
	var created_time basetypes.Float64Value
	var id basetypes.StringValue
	var modified_time basetypes.Float64Value
	var name basetypes.StringValue
	var org_id basetypes.StringValue
	var privileges basetypes.ListValue = types.ListNull(PrivilegesValue{}.Type(ctx))

	if d.CreatedTime != nil {
		created_time = types.Float64Value(*d.CreatedTime)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modified_time = types.Float64Value(*d.ModifiedTime)
	}

	name = types.StringValue(d.Name)

	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.Privileges != nil {
		privileges = privilegesSdkToTerraform(ctx, diags, d.Privileges)
	}

	data_map_attr_type := OrgSsoRolesValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"created_time":  created_time,
		"id":            id,
		"modified_time": modified_time,
		"name":          name,
		"org_id":        org_id,
		"privileges":    privileges,
	}
	data, e := NewOrgSsoRolesValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func privilegesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.PrivilegeOrg) basetypes.ListValue {

	var data_list = []PrivilegesValue{}
	for _, v := range data {
		var role types.String
		var scope types.String
		var site_id types.String
		var sitegroup_id types.String
		var views types.String

		role = types.StringValue(string(v.Role))
		scope = types.StringValue(string(v.Scope))
		if v.SiteId != nil {
			site_id = types.StringValue(v.SiteId.String())
		}
		if v.SitegroupId != nil {
			sitegroup_id = types.StringValue(v.SitegroupId.String())
		}
		if v.Views != nil {
			views = types.StringValue(string(*v.Views))
		}

		data_map_attr_type := PrivilegesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"role":         role,
			"scope":        scope,
			"site_id":      site_id,
			"sitegroup_id": sitegroup_id,
			"views":        views,
		}
		data, e := NewPrivilegesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}

	r, e := types.ListValueFrom(ctx, PrivilegesValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
