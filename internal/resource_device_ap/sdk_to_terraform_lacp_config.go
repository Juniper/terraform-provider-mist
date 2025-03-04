package resource_device_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func lacpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.DeviceApLacpConfig) LacpConfigValue {

	var enabled basetypes.BoolValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := NewLacpConfigValue(LacpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
