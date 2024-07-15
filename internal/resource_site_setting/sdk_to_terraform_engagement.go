package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mist_hours "terraform-provider-mist/internal/commons/hours"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func engagementDwellTagNamesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteEngagementDwellTagNames) basetypes.ObjectValue {
	tflog.Debug(ctx, "engagementDwellTagNamesSdkToTerraform")
	var bounce basetypes.StringValue
	var engaged basetypes.StringValue
	var passerby basetypes.StringValue
	var stationed basetypes.StringValue

	if d != nil && d.Bounce != nil {
		bounce = types.StringValue(*d.Bounce)
	}
	if d != nil && d.Engaged != nil {
		engaged = types.StringValue(*d.Engaged)
	}
	if d != nil && d.Passerby != nil {
		passerby = types.StringValue(*d.Passerby)
	}
	if d != nil && d.Stationed != nil {
		stationed = types.StringValue(*d.Stationed)
	}

	data_map_attr_type := DwellTagNamesValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"bounce":    bounce,
		"engaged":   engaged,
		"passerby":  passerby,
		"stationed": stationed,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func engagementDwellTagsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteEngagementDwellTags) basetypes.ObjectValue {
	tflog.Debug(ctx, "engagementDwellTagsSdkToTerraform")

	var bounce basetypes.StringValue
	var engaged basetypes.StringValue
	var passerby basetypes.StringValue
	var stationed basetypes.StringValue

	if d != nil && d.Bounce.Value() != nil {
		bounce = types.StringValue(*d.Bounce.Value())
	}
	if d != nil && d.Engaged.Value() != nil {
		engaged = types.StringValue(*d.Engaged.Value())
	}
	if d != nil && d.Passerby.Value() != nil {
		passerby = types.StringValue(*d.Passerby.Value())
	}
	if d != nil && d.Stationed.Value() != nil {
		stationed = types.StringValue(*d.Stationed.Value())
	}

	data_map_attr_type := DwellTagsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"bounce":    bounce,
		"engaged":   engaged,
		"passerby":  passerby,
		"stationed": stationed,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func engagementSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteEngagement) EngagementValue {
	tflog.Debug(ctx, "engagementSdkToTerraform")

	var dwell_tag_names basetypes.ObjectValue = types.ObjectNull(DwellTagNamesValue{}.AttributeTypes(ctx))
	var dwell_tags basetypes.ObjectValue = types.ObjectNull(DwellTagsValue{}.AttributeTypes(ctx))
	var hours basetypes.ObjectValue = types.ObjectNull(HoursValue{}.AttributeTypes(ctx))
	var max_dwell basetypes.Int64Value
	var min_dwell basetypes.Int64Value

	if d != nil && d.DwellTagNames != nil {
		dwell_tag_names = engagementDwellTagNamesSdkToTerraform(ctx, diags, d.DwellTagNames)
	}
	if d != nil && d.DwellTags != nil {
		dwell_tags = engagementDwellTagsSdkToTerraform(ctx, diags, d.DwellTags)
	}
	if d != nil && d.Hours != nil {
		hours = mist_hours.HoursSdkToTerraform(ctx, diags, d.Hours)
	}
	if d != nil && d.MaxDwell != nil {
		max_dwell = types.Int64Value(int64(*d.MaxDwell))
	}
	if d != nil && d.MinDwell != nil {
		min_dwell = types.Int64Value(int64(*d.MinDwell))
	}

	data_map_attr_type := EngagementValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"dwell_tag_names": dwell_tag_names,
		"dwell_tags":      dwell_tags,
		"hours":           hours,
		"max_dwell":       max_dwell,
		"min_dwell":       min_dwell,
	}
	data, e := NewEngagementValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
