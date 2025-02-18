package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func jcloudRaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingJcloudRa) JcloudRaValue {
	var orgApitoken basetypes.StringValue
	var orgApitokenName basetypes.StringValue
	var orgId basetypes.StringValue

	if d.OrgApitoken != nil {
		orgApitoken = types.StringValue(*d.OrgApitoken)
	}
	if d.OrgApitokenName != nil {
		orgApitokenName = types.StringValue(*d.OrgApitokenName)
	}
	if d.OrgId != nil {
		orgId = types.StringValue(*d.OrgId)
	}

	dataMapValue := map[string]attr.Value{
		"org_apitoken":      orgApitoken,
		"org_apitoken_name": orgApitokenName,
		"org_id":            orgId,
	}
	data, e := NewJcloudRaValue(JcloudRaValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
