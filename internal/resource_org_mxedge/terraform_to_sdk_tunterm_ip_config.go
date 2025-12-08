package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func tuntermIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d TuntermIpConfigValue) *models.MxedgeTuntermIpConfig {
	data := models.MxedgeTuntermIpConfig{}

	// Required fields
	data.Gateway = d.Gateway.ValueString()
	data.Ip = d.Ip.ValueString()
	data.Netmask = d.Netmask.ValueString()

	// Optional fields
	if !d.Gateway6.IsNull() && !d.Gateway6.IsUnknown() {
		data.Gateway6 = d.Gateway6.ValueStringPointer()
	}

	if !d.Ip6.IsNull() && !d.Ip6.IsUnknown() {
		data.Ip6 = d.Ip6.ValueStringPointer()
	}

	if !d.Netmask6.IsNull() && !d.Netmask6.IsUnknown() {
		data.Netmask6 = d.Netmask6.ValueStringPointer()
	}

	return &data
}
