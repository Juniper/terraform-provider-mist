package resource_org_apitoken

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/google/uuid"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, plan *OrgApitokenModel) (*models.OrgApitoken, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.OrgApitoken{}

	data.Name = plan.Name.ValueString()
	data.Privileges = privilegesTerraformToSdk(ctx, &diags, plan.Privileges)
	data.SrcIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.SrcIps)

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

		data_list = append(data_list, data)
	}

	return data_list
}
