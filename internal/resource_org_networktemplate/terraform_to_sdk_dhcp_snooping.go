package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func dhcpSnoopingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpSnoopingValue) *models.DhcpSnooping {
	data := models.DhcpSnooping{}
	if d.AllNetworks.ValueBoolPointer() != nil {
		data.AllNetworks = models.ToPointer(d.AllNetworks.ValueBool())
	}
	if d.EnableArpSpoofCheck.ValueBoolPointer() != nil {
		data.EnableArpSpoofCheck = models.ToPointer(d.EnableArpSpoofCheck.ValueBool())
	}
	if d.EnableIpSourceGuard.ValueBoolPointer() != nil {
		data.EnableIpSourceGuard = models.ToPointer(d.EnableIpSourceGuard.ValueBool())
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	if !d.Networks.IsNull() && !d.Networks.IsUnknown() {
		data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, d.Networks)
	}
	return &data
}
