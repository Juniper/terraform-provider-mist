package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func widsAuthFailureTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, o basetypes.ObjectValue) *models.SiteWidsRepeatedAuthFailures {
	tflog.Debug(ctx, "widsAuthFailureTerraformToSdk")
	data := models.SiteWidsRepeatedAuthFailures{}
	if o.IsNull() || o.IsUnknown() {
		return &data
	} else {
		d := NewRepeatedAuthFailuresValueMust(o.AttributeTypes(ctx), o.Attributes())
		data.Duration = models.ToPointer(int(d.Duration.ValueInt64()))
		data.Threshold = models.ToPointer(int(d.Threshold.ValueInt64()))
		return &data
	}
}

func widsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WidsValue) *models.SiteWids {
	tflog.Debug(ctx, "widsTerraformToSdk")
	data := models.SiteWids{}

	if !d.IsNull() {
		repeated_auth_failures := widsAuthFailureTerraformToSdk(ctx, diags, d.RepeatedAuthFailures)
		data.RepeatedAuthFailures = repeated_auth_failures
	}

	return &data
}
