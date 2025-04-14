package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	misthours "github.com/Juniper/terraform-provider-mist/internal/commons/hours"
)

func engagementDwellTagNamesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SiteEngagementDwellTagNames {
	data := models.SiteEngagementDwellTagNames{}
	if !d.IsNull() && !d.IsUnknown() {
		v, e := NewDwellTagNamesValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.Bounce = v.Bounce.ValueStringPointer()
			data.Engaged = v.Engaged.ValueStringPointer()
			data.Passerby = v.Passerby.ValueStringPointer()
			data.Stationed = v.Stationed.ValueStringPointer()
		}
	}
	return &data
}

func engagementDwellTagsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SiteEngagementDwellTags {
	data := models.SiteEngagementDwellTags{}
	if !d.IsNull() && !d.IsUnknown() {
		v, e := NewDwellTagsValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.Bounce = models.NewOptional(v.Bounce.ValueStringPointer())
			data.Engaged = models.NewOptional(v.Engaged.ValueStringPointer())
			data.Passerby = models.NewOptional(v.Passerby.ValueStringPointer())
			data.Stationed = models.NewOptional(v.Stationed.ValueStringPointer())
		}
	}
	return &data
}

func engagementTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d EngagementValue) *models.SiteEngagement {
	data := models.SiteEngagement{}

	data.DwellTagNames = engagementDwellTagNamesTerraformToSdk(ctx, diags, d.DwellTagNames)

	data.DwellTags = engagementDwellTagsTerraformToSdk(ctx, diags, d.DwellTags)

	data.Hours = misthours.HoursTerraformToSdk(ctx, diags, d.Hours)

	if !d.MaxDwell.IsNull() && !d.MaxDwell.IsUnknown() {
		data.MaxDwell = models.ToPointer(int(d.MaxDwell.ValueInt64()))
	}
	if !d.MinDwell.IsNull() && !d.MinDwell.IsUnknown() {
		data.MinDwell = models.ToPointer(int(d.MinDwell.ValueInt64()))
	}

	return &data
}
