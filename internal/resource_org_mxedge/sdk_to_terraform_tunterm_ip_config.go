package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func tuntermIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeTuntermIpConfig) TuntermIpConfigValue {

	var gateway types.String
	var gateway6 types.String
	var ip types.String
	var ip6 types.String
	var netmask types.String
	var netmask6 types.String

	// Required fields
	gateway = types.StringValue(d.Gateway)
	ip = types.StringValue(d.Ip)
	netmask = types.StringValue(d.Netmask)

	// Optional fields
	if d.Gateway6 != nil {
		gateway6 = types.StringValue(*d.Gateway6)
	}
	if d.Ip6 != nil {
		ip6 = types.StringValue(*d.Ip6)
	}
	if d.Netmask6 != nil {
		netmask6 = types.StringValue(*d.Netmask6)
	}

	data_map_attr_type := TuntermIpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"gateway":  gateway,
		"gateway6": gateway6,
		"ip":       ip,
		"ip6":      ip6,
		"netmask":  netmask,
		"netmask6": netmask6,
	}
	data, e := NewTuntermIpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
