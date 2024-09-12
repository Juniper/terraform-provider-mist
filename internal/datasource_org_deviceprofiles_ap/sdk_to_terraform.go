package datasource_org_deviceprofiles_ap

import (
	"context"
	"math/big"

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

	var created_time basetypes.NumberValue
	var id basetypes.StringValue
	var modified_time basetypes.NumberValue
	var name basetypes.StringValue
	var org_id basetypes.StringValue
	var deviceprofile_type basetypes.StringValue

	if d.CreatedTime != nil {
		created_time = types.NumberValue(big.NewFloat(*d.CreatedTime))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modified_time = types.NumberValue(big.NewFloat(*d.ModifiedTime))
	}
	name = types.StringValue(*d.Name)
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	deviceprofile_type = types.StringValue(string(d.Type))

	data_map_attr_type := OrgDeviceprofilesApValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"created_time":  created_time,
		"id":            id,
		"modified_time": modified_time,
		"name":          name,
		"type":          deviceprofile_type,
		"org_id":        org_id,
	}
	data, e := NewOrgDeviceprofilesApValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
