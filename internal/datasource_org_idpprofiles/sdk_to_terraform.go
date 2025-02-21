package datasource_org_idpprofiles

import (
	"context"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.IdpProfile, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := servicepolicieSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func servicepolicieSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpProfile) OrgIdpprofilesValue {

	var baseProfile types.String
	var createdTime basetypes.Float64Value
	var id types.String
	var modifiedTime basetypes.Float64Value
	var name types.String
	var orgId types.String
	var overwrites = types.ListNull(OverwritesValue{}.Type(ctx))

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.BaseProfile != nil {
		baseProfile = types.StringValue(string(*d.BaseProfile))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	orgId = types.StringValue(d.OrgId.String())

	if d.Overwrites != nil {
		overwrites = overwritesSdkToTerraform(ctx, diags, d.Overwrites)
	}

	dataMapValue := map[string]attr.Value{
		"base_profile":  baseProfile,
		"created_time":  createdTime,
		"id":            id,
		"modified_time": modifiedTime,
		"name":          name,
		"org_id":        orgId,
		"overwrites":    overwrites,
	}
	data, e := NewOrgIdpprofilesValue(OrgIdpprofilesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
