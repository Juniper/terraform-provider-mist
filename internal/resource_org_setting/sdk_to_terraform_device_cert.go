package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func deviceCertSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingDeviceCert) DeviceCertValue {

	var cert basetypes.StringValue
	var key basetypes.StringValue

	if d.Cert != nil {
		cert = types.StringValue(*d.Cert)
	}
	if d.Key != nil {
		key = types.StringValue(*d.Key)
	}

	dataMapValue := map[string]attr.Value{
		"cert": cert,
		"key":  key,
	}
	data, e := NewDeviceCertValue(DeviceCertValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
