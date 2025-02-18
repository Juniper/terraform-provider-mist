package resource_org_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func dhcpSnoopingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.DhcpSnooping) DhcpSnoopingValue {
	var allNetworks basetypes.BoolValue
	var enableArpSpoofCheck basetypes.BoolValue
	var enableIpSourceGuard basetypes.BoolValue
	var enabled basetypes.BoolValue
	var networks = types.ListNull(types.StringType)

	if d != nil && d.AllNetworks != nil {
		allNetworks = types.BoolValue(*d.AllNetworks)
	}
	if d != nil && d.EnableArpSpoofCheck != nil {
		enableArpSpoofCheck = types.BoolValue(*d.EnableArpSpoofCheck)
	}
	if d != nil && d.EnableIpSourceGuard != nil {
		enableIpSourceGuard = types.BoolValue(*d.EnableIpSourceGuard)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Networks != nil {
		networks = misttransform.ListOfStringSdkToTerraform(d.Networks)
	}

	dataMapValue := map[string]attr.Value{
		"all_networks":           allNetworks,
		"enable_arp_spoof_check": enableArpSpoofCheck,
		"enable_ip_source_guard": enableIpSourceGuard,
		"enabled":                enabled,
		"networks":               networks,
	}
	data, e := NewDhcpSnoopingValue(DhcpSnoopingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
