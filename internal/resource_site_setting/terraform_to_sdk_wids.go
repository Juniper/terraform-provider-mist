package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func widsAuthFailureTerraformToSdk(ctx context.Context, o basetypes.ObjectValue) *models.SiteWidsRepeatedAuthFailures {
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

func widsTerraformToSdk(ctx context.Context, d WidsValue) *models.SiteWids {
	data := models.SiteWids{}

	if !d.IsNull() {
		repeatedAuthFailures := widsAuthFailureTerraformToSdk(ctx, d.RepeatedAuthFailures)
		data.RepeatedAuthFailures = repeatedAuthFailures
	}

	return &data
}
