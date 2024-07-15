package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func mistNacdSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanMistNac) MistNacValue {
	var enabled basetypes.BoolValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := MistNacValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := NewMistNacValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
