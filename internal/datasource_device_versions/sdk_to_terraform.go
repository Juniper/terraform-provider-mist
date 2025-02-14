package datasource_device_versions

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.DeviceVersionItem) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := constWebhookSdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(DeviceVersionsValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func constWebhookSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.DeviceVersionItem) DeviceVersionsValue {
	var model basetypes.StringValue
	var tag basetypes.StringValue
	var version basetypes.StringValue

	model = types.StringValue(d.Model)

	if d.Tag != nil {
		tag = types.StringValue(*d.Tag)
	}

	version = types.StringValue(d.Version)

	o, e := NewDeviceVersionsValue(
		DeviceVersionsValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"model":   model,
			"tag":     tag,
			"version": version,
		},
	)
	diags.Append(e...)
	return o
}
