package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	misthours "github.com/Juniper/terraform-provider-mist/internal/commons/hours"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func engagementDwellTagNamesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteEngagementDwellTagNames) basetypes.ObjectValue {
	// Return null object if data is nil or has no values
	if d == nil || (d.Bounce == nil && d.Engaged == nil && d.Passerby == nil && d.Stationed == nil) {
		return types.ObjectNull(DwellTagNamesValue{}.AttributeTypes(ctx))
	}

	var bounce = types.StringNull()
	var engaged = types.StringNull()
	var passerby = types.StringNull()
	var stationed = types.StringNull()

	if d.Bounce != nil {
		bounce = types.StringValue(*d.Bounce)
	}
	if d.Engaged != nil {
		engaged = types.StringValue(*d.Engaged)
	}
	if d.Passerby != nil {
		passerby = types.StringValue(*d.Passerby)
	}
	if d.Stationed != nil {
		stationed = types.StringValue(*d.Stationed)
	}

	dataMapValue := map[string]attr.Value{
		"bounce":    bounce,
		"engaged":   engaged,
		"passerby":  passerby,
		"stationed": stationed,
	}
	data, e := basetypes.NewObjectValue(DwellTagNamesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func engagementDwellTagsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteEngagementDwellTags) basetypes.ObjectValue {
	// Return null object if data is nil or has no values
	if d == nil || (d.Bounce.Value() == nil && d.Engaged.Value() == nil && d.Passerby.Value() == nil && d.Stationed.Value() == nil) {
		return types.ObjectNull(DwellTagsValue{}.AttributeTypes(ctx))
	}

	var bounce = types.StringNull()
	var engaged = types.StringNull()
	var passerby = types.StringNull()
	var stationed = types.StringNull()

	if d.Bounce.Value() != nil {
		bounce = types.StringValue(*d.Bounce.Value())
	}
	if d.Engaged.Value() != nil {
		engaged = types.StringValue(*d.Engaged.Value())
	}
	if d.Passerby.Value() != nil {
		passerby = types.StringValue(*d.Passerby.Value())
	}
	if d.Stationed.Value() != nil {
		stationed = types.StringValue(*d.Stationed.Value())
	}

	dataMapValue := map[string]attr.Value{
		"bounce":    bounce,
		"engaged":   engaged,
		"passerby":  passerby,
		"stationed": stationed,
	}
	data, e := basetypes.NewObjectValue(DwellTagsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func engagementSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteEngagement) EngagementValue {

	var dwellTagNames = types.ObjectNull(DwellTagNamesValue{}.AttributeTypes(ctx))
	var dwellTags = types.ObjectNull(DwellTagsValue{}.AttributeTypes(ctx))
	var hours = types.ObjectNull(HoursValue{}.AttributeTypes(ctx))
	var maxDwell = types.Int64Null()
	var minDwell = types.Int64Null()

	if d != nil && d.DwellTagNames != nil {
		dwellTagNames = engagementDwellTagNamesSdkToTerraform(ctx, diags, d.DwellTagNames)
	}
	if d != nil && d.DwellTags != nil {
		dwellTags = engagementDwellTagsSdkToTerraform(ctx, diags, d.DwellTags)
	}
	if d != nil && d.Hours != nil {
		hours = misthours.HoursSdkToTerraform(diags, d.Hours)
	}
	if d != nil && d.MaxDwell != nil {
		maxDwell = types.Int64Value(int64(*d.MaxDwell))
	}
	if d != nil && d.MinDwell != nil {
		minDwell = types.Int64Value(int64(*d.MinDwell))
	}

	dataMapValue := map[string]attr.Value{
		"dwell_tag_names": dwellTagNames,
		"dwell_tags":      dwellTags,
		"hours":           hours,
		"max_dwell":       maxDwell,
		"min_dwell":       minDwell,
	}
	data, e := NewEngagementValue(EngagementValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
