package datasource_org_wlans

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dnsServerRewriteSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.WlanDnsServerRewrite) basetypes.ObjectValue {
	if data == nil {
		return basetypes.NewObjectNull(DnsServerRewriteValue{}.AttributeTypes(ctx))
	}

	var enabled basetypes.BoolValue
	if data.Enabled != nil {
		enabled = types.BoolValue(*data.Enabled)
	}

	radiusGroupsValues := make(map[string]attr.Value)
	for k, v := range data.RadiusGroups {
		radiusGroupsValues[k] = types.StringValue(v)
	}
	radiusGroups := types.MapValueMust(types.StringType, radiusGroupsValues)

	dataMapValue := map[string]attr.Value{
		"enabled":       enabled,
		"radius_groups": radiusGroups,
	}
	result, err := basetypes.NewObjectValue(DnsServerRewriteValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(err...)

	return result
}
