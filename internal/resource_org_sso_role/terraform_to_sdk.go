package resource_org_sso_role

import (
	"github.com/google/uuid"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(plan *OrgSsoRoleModel) (*models.SsoRoleOrg, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.SsoRoleOrg{}

	data.Name = plan.Name.ValueString()
	data.Privileges = privilegesTerraformToSdk(&diags, plan.Privileges)

	return &data, diags
}

func privilegesTerraformToSdk(diags *diag.Diagnostics, d basetypes.ListValue) []models.PrivilegeOrg {
	var dataList []models.PrivilegeOrg
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(PrivilegesValue)
		data := models.PrivilegeOrg{}

		data.Role = models.PrivilegeOrgRoleEnum(*plan.Role.ValueStringPointer())
		data.Scope = models.PrivilegeOrgScopeEnum(*plan.Scope.ValueStringPointer())
		if !plan.SiteId.IsNull() && !plan.SiteId.IsUnknown() {
			tmp, err := uuid.Parse(plan.SiteId.ValueString())
			if err != nil {
				diags.AddError(
					"Invalid \"site_id\" value for \"org_apitoken.privilege\" resource",
					"Could parse the UUID: "+err.Error(),
				)
			} else {
				data.SiteId = &tmp
			}
		}
		if !plan.SitegroupId.IsNull() && !plan.SitegroupId.IsUnknown() {
			tmp, err := uuid.Parse(plan.SitegroupId.ValueString())
			if err != nil {
				diags.AddError(
					"Invalid \"sitegroup_id\" value for \"org_apitoken.privilege\" resource",
					"Could parse the UUID: "+err.Error(),
				)
			} else {
				data.SitegroupId = &tmp
			}
		}

		if !plan.Views.IsNull() && !plan.Views.IsUnknown() {
			for _, v := range plan.Views.Elements() {
				var vi interface{} = v
				role := vi.(basetypes.StringValue)
				data.Views = append(data.Views, models.AdminPrivilegeViewEnum(role.ValueString()))
			}
		}

		dataList = append(dataList, data)
	}

	return dataList
}
