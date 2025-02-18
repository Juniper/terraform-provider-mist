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
	var createdTime basetypes.Float64Value
	var id basetypes.StringValue
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var privileges = types.ListNull(PrivilegesValue{}.Type(ctx))

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}

	name = types.StringValue(d.Name)

	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Privileges != nil {
		privileges = privilegesSdkToTerraform(ctx, diags, d.Privileges)
	}

	dataMapValue := map[string]attr.Value{
		"created_time":  createdTime,
		"id":            id,
		"modified_time": modifiedTime,
		"name":          name,
		"org_id":        orgId,
		"privileges":    privileges,
	}
	data, e := NewOrgSsoRolesValue(OrgSsoRolesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func privilegesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.PrivilegeOrg) basetypes.ListValue {

	var dataList []PrivilegesValue
	for _, v := range data {
		var role types.String
		var scope types.String
		var siteId types.String
		var sitegroupId types.String
		var views = types.ListNull(types.StringType)

		role = types.StringValue(string(v.Role))
		scope = types.StringValue(string(v.Scope))
		if v.SiteId != nil {
			siteId = types.StringValue(v.SiteId.String())
		}
		if v.SitegroupId != nil {
			sitegroupId = types.StringValue(v.SitegroupId.String())
		}
		if v.Views != nil {
			var viewsArray []attr.Value
			for _, role := range v.Views {
				viewsArray = append(viewsArray, types.StringValue(string(role)))
			}
			tmp, e := types.ListValueFrom(ctx, types.StringType, viewsArray)
			if e != nil {
				diags.Append(e...)
			} else {
				views = tmp
			}
		}

		dataMapValue := map[string]attr.Value{
			"role":         role,
			"scope":        scope,
			"site_id":      siteId,
			"sitegroup_id": sitegroupId,
			"views":        views,
		}
		data, e := NewPrivilegesValue(PrivilegesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}

	r, e := types.ListValueFrom(ctx, PrivilegesValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
