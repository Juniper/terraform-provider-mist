package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func centrakSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApCentrak) CentrakValue {

	rAttrType := CentrakValue{}.AttributeTypes(ctx)
	rAttrValue := map[string]attr.Value{
		"enabled": types.BoolValue(*d.Enabled),
	}
	r, e := NewCentrakValue(rAttrType, rAttrValue)
	diags.Append(e...)
	return r
}
