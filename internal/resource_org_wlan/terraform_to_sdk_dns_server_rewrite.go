package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dnsServerRewriteTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan DnsServerRewriteValue) *models.WlanDnsServerRewrite {

	radius_groups := make(map[string]string)
	for k, v := range plan.RadiusGroups.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		radius_groups[k] = v_plan.ValueString()
	}

	data := models.WlanDnsServerRewrite{}
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.RadiusGroups = radius_groups

	return &data
}
