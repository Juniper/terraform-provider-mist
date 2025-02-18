package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func skyAtpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingSkyatp) SkyatpValue {

	var enabled basetypes.BoolValue
	var sendIpMacMapping basetypes.BoolValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.SendIpMacMapping != nil {
		sendIpMacMapping = types.BoolValue(*d.SendIpMacMapping)
	}

	dataMapValue := map[string]attr.Value{
		"enabled":             enabled,
		"send_ip_mac_mapping": sendIpMacMapping,
	}
	data, e := NewSkyatpValue(SkyatpValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
