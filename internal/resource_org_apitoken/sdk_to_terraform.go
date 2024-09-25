package resource_org_apitoken

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data models.OrgApitoken, previous_state *OrgApitokenModel) (OrgApitokenModel, diag.Diagnostics) {
	var state OrgApitokenModel
	var diags diag.Diagnostics

	var created_by types.String
	var id types.String
	var key types.String
	var name types.String
	var org_id types.String
	var privileges types.List = types.ListNull(PrivilegesValue{}.Type(ctx))
	var src_ips types.List = types.ListNull(types.StringType)

	if data.CreatedBy.Value() != nil {
		created_by = types.StringValue(*data.CreatedBy.Value())
	}
	id = types.StringValue(data.Id.String())
	// Since Mist is not returning the complete API Token afterward, keeping the value
	// sent back when the API Token has been created
	if previous_state != nil && previous_state.Key.ValueStringPointer() != nil {
		key = previous_state.Key
	} else {
		key = types.StringValue(*data.Key)
	}
	name = types.StringValue(data.Name)
	org_id = types.StringValue(data.OrgId.String())
	privileges = privilegesSdkToTerraform(ctx, &diags, data.Privileges)
	if data.SrcIps != nil {
		src_ips = mist_transform.ListOfStringSdkToTerraform(ctx, data.SrcIps)
	}

	state.CreatedBy = created_by
	state.Id = id
	state.Key = key
	state.Name = name
	state.OrgId = org_id
	state.Privileges = privileges
	state.SrcIps = src_ips

	return state, diags
}

func privilegesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.PrivilegeOrg) basetypes.ListValue {

	var data_list = []PrivilegesValue{}
	for _, v := range data {
		var role types.String
		var scope types.String
		var site_id types.String
		var sitegroup_id types.String

		role = types.StringValue(string(v.Role))
		scope = types.StringValue(string(v.Scope))
		if v.SiteId != nil {
			site_id = types.StringValue(v.SiteId.String())
		}
		if v.SitegroupId != nil {
			sitegroup_id = types.StringValue(v.SitegroupId.String())
		}

		data_map_attr_type := PrivilegesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"role":         role,
			"scope":        scope,
			"site_id":      site_id,
			"sitegroup_id": sitegroup_id,
		}
		data, e := NewPrivilegesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}

	r, e := types.ListValueFrom(ctx, PrivilegesValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
