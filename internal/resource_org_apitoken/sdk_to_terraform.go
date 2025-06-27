package resource_org_apitoken

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data models.OrgApitoken, previousState *OrgApitokenModel) (OrgApitokenModel, diag.Diagnostics) {
	var state OrgApitokenModel
	var diags diag.Diagnostics

	// Since Mist is not returning the complete API Token afterward, keeping the value
	// sent back when the API Token has been created
	if previousState != nil {
		state = *previousState
	} else {
		state.Key = types.StringValue(*data.Key)
	}

	if data.CreatedBy.Value() != nil {
		state.CreatedBy = types.StringValue(*data.CreatedBy.Value())
	}
	state.Id = types.StringValue(data.Id.String())

	state.Name = types.StringValue(data.Name)
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Privileges = privilegesSdkToTerraform(ctx, &diags, data.Privileges)
	if data.SrcIps != nil {
		state.SrcIps = mistutils.ListOfStringSdkToTerraform(data.SrcIps)
	} else {
		state.SrcIps = types.ListNull(types.StringType)
	}

	return state, diags
}

func privilegesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.PrivilegeOrg) basetypes.ListValue {

	var dataList []PrivilegesValue
	for _, v := range data {
		var role types.String
		var scope types.String
		var siteId types.String
		var sitegroupId types.String

		role = types.StringValue(string(v.Role))
		scope = types.StringValue(string(v.Scope))
		if v.SiteId != nil {
			siteId = types.StringValue(v.SiteId.String())
		}
		if v.SitegroupId != nil {
			sitegroupId = types.StringValue(v.SitegroupId.String())
		}

		dataMapValue := map[string]attr.Value{
			"role":         role,
			"scope":        scope,
			"site_id":      siteId,
			"sitegroup_id": sitegroupId,
		}
		priv, e := NewPrivilegesValue(PrivilegesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, priv)
	}

	r, e := types.ListValueFrom(ctx, PrivilegesValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
