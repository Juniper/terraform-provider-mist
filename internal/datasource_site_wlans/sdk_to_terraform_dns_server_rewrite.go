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
	var radiusGroups = types.MapNull(types.StringType)

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	radiusGroupsValues := make(map[string]attr.Value)
	if d != nil && d.RadiusGroups != nil && len(d.RadiusGroups) > 0 {
		for k, v := range d.RadiusGroups {
			radiusGroupsValues[k] = types.StringValue(v)
		}
	}
	radiusGroups = types.MapValueMust(types.StringType, radiusGroupsValues)

	dataMapValue := map[string]attr.Value{
		"enabled":       enabled,
		"radius_groups": radiusGroups,
	}
	data, e := basetypes.NewObjectValue(DnsServerRewriteValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
