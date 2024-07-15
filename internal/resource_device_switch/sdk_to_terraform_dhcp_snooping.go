package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func dhcpSnoopingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.DhcpSnooping) DhcpSnoopingValue {
	var all_networks basetypes.BoolValue
	var enable_arp_spoof_check basetypes.BoolValue
	var enable_ip_source_guard basetypes.BoolValue
	var enabled basetypes.BoolValue
	var networks basetypes.ListValue = types.ListNull(types.StringType)

	if d != nil && d.AllNetworks != nil {
		all_networks = types.BoolValue(*d.AllNetworks)
	}
	if d != nil && d.EnableArpSpoofCheck != nil {
		enable_arp_spoof_check = types.BoolValue(*d.EnableArpSpoofCheck)
	}
	if d != nil && d.EnableIpSourceGuard != nil {
		enable_ip_source_guard = types.BoolValue(*d.EnableIpSourceGuard)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Networks != nil {
		networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Networks)
	}

	data_map_attr_type := DhcpSnoopingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"all_networks":           all_networks,
		"enable_arp_spoof_check": enable_arp_spoof_check,
		"enable_ip_source_guard": enable_ip_source_guard,
		"enabled":                enabled,
		"networks":               networks,
	}
	data, e := NewDhcpSnoopingValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
