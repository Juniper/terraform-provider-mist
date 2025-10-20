package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func vpnOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingVpnOptions) VpnOptionsValue {

	var asBase basetypes.Int64Value
	var enableIpv6 basetypes.BoolValue
	var stSubnet basetypes.StringValue

	if d.AsBase != nil {
		asBase = types.Int64Value(int64(*d.AsBase))
	}
	if d.EnableIpv6 != nil {
		enableIpv6 = types.BoolValue(*d.EnableIpv6)
	}
	if d.StSubnet != nil {
		stSubnet = types.StringValue(*d.StSubnet)
	}

	dataMapValue := map[string]attr.Value{
		"as_base":     asBase,
		"enable_ipv6": enableIpv6,
		"st_subnet":   stSubnet,
	}
	data, e := NewVpnOptionsValue(VpnOptionsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
