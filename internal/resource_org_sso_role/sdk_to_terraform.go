package resource_org_sso_role

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data models.SsoRoleOrg) (OrgSsoRoleModel, diag.Diagnostics) {
	var state OrgSsoRoleModel
	var diags diag.Diagnostics

	var id types.String
	var name types.String
	var org_id types.String
	var privileges types.List = types.ListNull(PrivilegesValue{}.Type(ctx))

	id = types.StringValue(data.Id.String())
	name = types.StringValue(data.Name)
	org_id = types.StringValue(data.OrgId.String())
	privileges = privilegesSdkToTerraform(ctx, &diags, data.Privileges)
	state.Id = id
	state.Name = name
	state.OrgId = org_id
	state.Privileges = privileges

	return state, diags
}

func privilegesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.PrivilegeOrg) basetypes.ListValue {

	var data_list = []PrivilegesValue{}
	for _, v := range data {
		var role types.String
		var scope types.String
		var site_id types.String
		var sitegroup_id types.String
		var views types.List = types.ListNull(types.StringType)

		role = types.StringValue(string(v.Role))
		scope = types.StringValue(string(v.Scope))
		if v.SiteId != nil {
			site_id = types.StringValue(v.SiteId.String())
		}
		if v.SitegroupId != nil {
			sitegroup_id = types.StringValue(v.SitegroupId.String())
		}
		if v.Views != nil {
			var views_array []attr.Value
			for _, role := range v.Views {
				views_array = append(views_array, types.StringValue(string(role)))
			}
			tmp, e := types.ListValueFrom(ctx, types.StringType, views_array)
			if e != nil {
				diags.Append(e...)
			} else {
				views = tmp
			}
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
