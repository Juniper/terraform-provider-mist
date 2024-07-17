package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func jcloudSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingJcloud) JcloudValue {
	var org_apitoken basetypes.StringValue
	var org_apitoken_name basetypes.StringValue
	var org_id basetypes.StringValue

	if d.OrgApitoken != nil {
		org_apitoken = types.StringValue(*d.OrgApitoken)
	}
	if d.OrgApitokenName != nil {
		org_apitoken_name = types.StringValue(*d.OrgApitokenName)
	}
	if d.OrgId != nil {
		org_id = types.StringValue(*d.OrgId)
	}

	data_map_attr_type := JcloudValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"org_apitoken":      org_apitoken,
		"org_apitoken_name": org_apitoken_name,
		"org_id":            org_id,
	}
	data, e := NewJcloudValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
