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

func skyAtpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingSkyatp) SkyatpValue {
	tflog.Debug(ctx, "skyAtpSdkToTerraform")

	var enabled basetypes.BoolValue
	var send_ip_mac_mapping basetypes.BoolValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.SendIpMacMapping != nil {
		send_ip_mac_mapping = types.BoolValue(*d.SendIpMacMapping)
	}

	data_map_attr_type := SkyatpValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":             enabled,
		"send_ip_mac_mapping": send_ip_mac_mapping,
	}
	data, e := NewSkyatpValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
