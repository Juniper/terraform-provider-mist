package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func analyticSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingAnalytic) AnalyticValue {
	var enabled basetypes.BoolValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := NewAnalyticValue(AnalyticValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
