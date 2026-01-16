package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermIgmpSnoopingConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d TuntermIgmpSnoopingConfigValue) *models.MxedgeTuntermIgmpSnoopingConfig {
	data := models.MxedgeTuntermIgmpSnoopingConfig{}

	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if !d.Querier.IsNull() && !d.Querier.IsUnknown() {
		var querier QuerierValue
		d.Querier.As(ctx, &querier, basetypes.ObjectAsOptions{})
		data.Querier = querierTerraformToSdk(ctx, diags, querier)
	}

	if !d.VlanIds.IsNull() && !d.VlanIds.IsUnknown() {
		var vlanIds []int
		d.VlanIds.ElementsAs(ctx, &vlanIds, false)
		data.VlanIds = vlanIds
	}

	return &data
}

func querierTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d QuerierValue) *models.MxedgeTuntermIgmpSnoopingQuerier {
	data := models.MxedgeTuntermIgmpSnoopingQuerier{}

	if !d.MaxResponseTime.IsNull() && !d.MaxResponseTime.IsUnknown() {
		data.MaxResponseTime = models.ToPointer(int(d.MaxResponseTime.ValueInt64()))
	}

	if !d.Mtu.IsNull() && !d.Mtu.IsUnknown() {
		data.Mtu = models.ToPointer(int(d.Mtu.ValueInt64()))
	}

	if !d.QueryInterval.IsNull() && !d.QueryInterval.IsUnknown() {
		data.QueryInterval = models.ToPointer(int(d.QueryInterval.ValueInt64()))
	}

	if !d.Robustness.IsNull() && !d.Robustness.IsUnknown() {
		data.Robustness = models.ToPointer(int(d.Robustness.ValueInt64()))
	}

	if !d.Version.IsNull() && !d.Version.IsUnknown() {
		data.Version = models.ToPointer(int(d.Version.ValueInt64()))
	}

	return &data
}
