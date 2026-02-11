package resource_org_sso_role

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data models.SsoRoleOrg) (OrgSsoRoleModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	if data.Id == nil || data.OrgId == nil {
		diags.AddError("Error: SsoRoleOrg ID/OrgID is nil.", "The SSO Role ID/OrgID is nil.")
		return OrgSsoRoleModel{}, diags
	}

	state := OrgSsoRoleModel{
		Id:         types.StringValue(data.Id.String()),
		Name:       types.StringValue(data.Name),
		OrgId:      types.StringValue(data.OrgId.String()),
		Privileges: privilegesSdkToTerraform(ctx, &diags, data.Privileges),
	}

	return state, diags
}

func privilegesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.PrivilegeOrg) basetypes.ListValue {
	var privileges []PrivilegesValue
	for _, val := range data {
		role := types.StringValue(string(val.Role))
		scope := types.StringValue(string(val.Scope))

		var siteId basetypes.StringValue
		if val.SiteId != nil {
			siteId = types.StringValue(val.SiteId.String())
		}

		var sitegroupId basetypes.StringValue
		if val.SitegroupId != nil {
			sitegroupId = types.StringValue(val.SitegroupId.String())
		}

		var viewsArray []attr.Value
		if val.View != nil && val.Views == nil {
			viewsArray = append(viewsArray, types.StringValue(string(*val.View)))
		}

		if val.Views != nil {
			for _, view := range val.Views {
				viewsArray = append(viewsArray, types.StringValue(string(view)))
			}
		}

		views := mistutils.ListOfStringSdkToTerraformEmpty()
		if len(viewsArray) > 0 {
			var err diag.Diagnostics
			views, err = types.ListValueFrom(ctx, types.StringType, viewsArray)
			diags.Append(err...)
		}

		dataMap := map[string]attr.Value{
			"role":         role,
			"scope":        scope,
			"site_id":      siteId,
			"sitegroup_id": sitegroupId,
			"views":        views,
		}
		item, err := NewPrivilegesValue(PrivilegesValue{}.AttributeTypes(ctx), dataMap)
		diags.Append(err...)

		privileges = append(privileges, item)
	}

	result, err := types.ListValueFrom(ctx, PrivilegesValue{}.Type(ctx), privileges)
	diags.Append(err...)

	return result
}
