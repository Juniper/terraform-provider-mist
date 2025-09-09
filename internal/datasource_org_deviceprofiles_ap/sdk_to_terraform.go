package datasource_org_deviceprofiles_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.DeviceprofileAp, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := deviceprofileApSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func deviceprofileApSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.DeviceprofileAp) OrgDeviceprofilesApValue {

	var createdTime basetypes.Float64Value
	var id basetypes.StringValue
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var orgId basetypes.StringValue

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	if d.Name.Value() != nil {
		name = types.StringValue(*d.Name.Value())
	}
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}

	dataMapValue := map[string]attr.Value{
		"created_time":  createdTime,
		"id":            id,
		"modified_time": modifiedTime,
		"name":          name,
		"org_id":        orgId,
	}
	data, e := NewOrgDeviceprofilesApValue(OrgDeviceprofilesApValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
