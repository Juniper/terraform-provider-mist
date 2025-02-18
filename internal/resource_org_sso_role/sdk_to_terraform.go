package resource_org_sso_role

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
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
	var orgId types.String
	var privileges = types.ListNull(PrivilegesValue{}.Type(ctx))

	id = types.StringValue(data.Id.String())
	name = types.StringValue(data.Name)
	orgId = types.StringValue(data.OrgId.String())
	privileges = privilegesSdkToTerraform(ctx, &diags, data.Privileges)
	state.Id = id
	state.Name = name
	state.OrgId = orgId
	state.Privileges = privileges

	return state, diags
}

func privilegesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.PrivilegeOrg) basetypes.ListValue {

	var dataList []PrivilegesValue
	for _, v := range data {
		var role types.String
		var scope types.String
		var siteId types.String
		var sitegroupId types.String
		var views = misttransform.ListOfStringSdkToTerraformEmpty()

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
