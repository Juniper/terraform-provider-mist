package resource_org_sso_role

import (
	"context"

	"github.com/google/uuid"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, plan *OrgSsoRoleModel) (*models.SsoRoleOrg, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.SsoRoleOrg{}

	data.Name = plan.Name.ValueString()
	data.Privileges = privilegesTerraformToSdk(ctx, &diags, plan.Privileges)

	return &data, diags
}

func privilegesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.PrivilegeOrg {
	var data_list []models.PrivilegeOrg
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PrivilegesValue)
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
			data.Views = models.ToPointer(models.AdminPrivilegeViewEnum(plan.Views.ValueString()))
		}

		data_list = append(data_list, data)
	}

	return data_list
}
