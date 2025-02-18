package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func qosSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanQos) QosValue {

	var class basetypes.StringValue
	var overwrite basetypes.BoolValue

	if d != nil && d.Class != nil {
		class = types.StringValue(string(*d.Class))
	}
	if d != nil && d.Overwrite != nil {
		overwrite = types.BoolValue(*d.Overwrite)
	}

	dataMapValue := map[string]attr.Value{
		"class":     class,
		"overwrite": overwrite,
	}
	data, e := NewQosValue(QosValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
