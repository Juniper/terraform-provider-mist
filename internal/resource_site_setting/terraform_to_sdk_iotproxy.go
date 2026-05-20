package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func iotproxyVisionlineTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VisionlineValue) *models.IotproxyVisionline {
	data := models.IotproxyVisionline{}

	if d.AccessId.ValueStringPointer() != nil {
		data.AccessId = d.AccessId.ValueStringPointer()
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Host.ValueStringPointer() != nil {
		data.Host = d.Host.ValueStringPointer()
	}
	if d.Password.ValueStringPointer() != nil {
		data.Password = d.Password.ValueStringPointer()
	}
	if d.Port.ValueInt64Pointer() != nil {
		data.Port = models.ToPointer(int(d.Port.ValueInt64()))
	}
	if d.Username.ValueStringPointer() != nil {
		data.Username = d.Username.ValueStringPointer()
	}

	return &data
}

func iotproxyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d IotproxyValue) *models.Iotproxy {
	data := models.Iotproxy{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.Visionline.IsNull() && !d.Visionline.IsUnknown() {
		visionlinePlan, e := NewVisionlineValue(d.Visionline.AttributeTypes(ctx), d.Visionline.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.Visionline = iotproxyVisionlineTerraformToSdk(ctx, diags, visionlinePlan)
		}
	}

	return &data
}
