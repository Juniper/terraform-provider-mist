package datasource_org_deviceprofiles_ap

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.Deviceprofile) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		ap_js, e := d.MarshalJSON()
		if e != nil {
			diags.AddError("Unable to Marshal Deviceprofile AP", e.Error())
		} else {
			deviceprofile := models.DeviceprofileAp{}
			e := json.Unmarshal(ap_js, &deviceprofile)
			if e != nil {
				diags.AddError("Unable to unMarshal AP Stats", e.Error())
			}
			elem := deviceprofileApSdkToTerraform(ctx, &diags, &deviceprofile)
			elements = append(elements, elem)
		}
	}

	dataSet, err := types.SetValue(OrgDeviceprofilesApValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
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
	name = types.StringValue(*d.Name.Value())
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	deviceprofile_type = types.StringValue(string(*d.Type))

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
