package resource_org_apitoken

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data models.OrgApitoken, previousState *OrgApitokenModel) (OrgApitokenModel, diag.Diagnostics) {
	var state OrgApitokenModel
	var diags diag.Diagnostics

	var createdBy types.String
	var id types.String
	var key types.String
	var name types.String
	var orgId types.String
	var privileges = types.ListNull(PrivilegesValue{}.Type(ctx))
	var srcIps = types.ListNull(types.StringType)

	if data.CreatedBy.Value() != nil {
		createdBy = types.StringValue(*data.CreatedBy.Value())
	}
	id = types.StringValue(data.Id.String())
	// Since Mist is not returning the complete API Token afterward, keeping the value
	// sent back when the API Token has been created
	if previousState != nil && previousState.Key.ValueStringPointer() != nil {
		key = previousState.Key
	} else {
		key = types.StringValue(*data.Key)
	}
	name = types.StringValue(data.Name)
	orgId = types.StringValue(data.OrgId.String())
	privileges = privilegesSdkToTerraform(ctx, &diags, data.Privileges)
	if data.SrcIps != nil {
		srcIps = misttransform.ListOfStringSdkToTerraform(data.SrcIps)
	}

	state.CreatedBy = createdBy
	state.Id = id
	state.Key = key
	state.Name = name
	state.OrgId = orgId
	state.Privileges = privileges
	state.SrcIps = srcIps

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
		data, e := NewPrivilegesValue(PrivilegesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}

	r, e := types.ListValueFrom(ctx, PrivilegesValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
