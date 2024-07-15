package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func srxAppSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingSrxApp) SrxAppValue {
	tflog.Debug(ctx, "srxAppSdkToTerraform")

	var enabled basetypes.BoolValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := SrxAppValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := NewSrxAppValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
