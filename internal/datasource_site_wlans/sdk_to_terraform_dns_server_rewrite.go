package datasource_site_wlans

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dnsServerRewriteSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanDnsServerRewrite) basetypes.ObjectValue {

	var enabled basetypes.BoolValue
	var radius_groups basetypes.MapValue = types.MapNull(types.StringType)

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	radius_groups_values := make(map[string]attr.Value)
	if d != nil && d.RadiusGroups != nil && len(d.RadiusGroups) > 0 {
		for k, v := range d.RadiusGroups {
			radius_groups_values[k] = types.StringValue(v)
		}
	}
	radius_groups = types.MapValueMust(types.StringType, radius_groups_values)

	data_map_attr_type := DnsServerRewriteValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":       enabled,
		"radius_groups": radius_groups,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
