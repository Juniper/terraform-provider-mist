package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func cloudsharkSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingCloudshark) CloudsharkValue {

	var apitoken basetypes.StringValue
	var url basetypes.StringValue

	if d.Apitoken != nil {
		apitoken = types.StringValue(*d.Apitoken)
	}
	if d.Url != nil {
		url = types.StringValue(*d.Url)
	}

	dataMapValue := map[string]attr.Value{
		"apitoken": apitoken,
		"url":      url,
	}
	data, e := NewCloudsharkValue(CloudsharkValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
