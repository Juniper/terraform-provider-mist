package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func centrakSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApCentrak) CentrakValue {

	r_attr_type := CentrakValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(*d.Enabled),
	}
	r, e := NewCentrakValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
