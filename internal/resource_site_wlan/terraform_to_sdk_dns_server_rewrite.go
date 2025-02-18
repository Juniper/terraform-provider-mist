package resource_site_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dnsServerRewriteTerraformToSdk(plan DnsServerRewriteValue) *models.WlanDnsServerRewrite {

	radiusGroups := make(map[string]string)
	for k, v := range plan.RadiusGroups.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(basetypes.StringValue)
		radiusGroups[k] = vPlan.ValueString()
	}

	data := models.WlanDnsServerRewrite{}
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.RadiusGroups = radiusGroups

	return &data
}
