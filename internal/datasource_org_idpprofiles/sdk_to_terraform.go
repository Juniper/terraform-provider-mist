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

	var base_profile types.String
	var created_time basetypes.NumberValue
	var id types.String
	var modified_time basetypes.NumberValue
	var name types.String
	var org_id types.String
	var overwrites types.List = types.ListNull(OverwritesValue{}.Type(ctx))

	if d.BaseProfile != nil {
		base_profile = types.StringValue(string(*d.BaseProfile))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	org_id = types.StringValue(d.OrgId.String())

	if d.Overwrites != nil {
		overwrites = overwritesSdkToTerraform(ctx, diags, d.Overwrites)
	}

	data_map_attr_type := OrgIdpprofilesValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"base_profile":  base_profile,
		"created_time":  created_time,
		"id":            id,
		"modified_time": modified_time,
		"name":          name,
		"org_id":        org_id,
		"overwrites":    overwrites,
	}
	data, e := NewOrgIdpprofilesValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
