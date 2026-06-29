package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func marvisSelfDrivingDomainSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MarvisSelfDrivingDomain, newValue func(map[string]attr.Type, map[string]attr.Value) (basetypes.ObjectValuable, diag.Diagnostics), attrTypes map[string]attr.Type) basetypes.ObjectValue {
	var enabled = types.BoolNull()
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	data, e := newValue(attrTypes, map[string]attr.Value{
		"enabled": enabled,
	})
	diags.Append(e...)
	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func marvisSelfDrivingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MarvisSelfDriving) basetypes.ObjectValue {
	var wan = types.ObjectNull(WanValue{}.AttributeTypes(ctx))
	var wired = types.ObjectNull(WiredValue{}.AttributeTypes(ctx))
	var wireless = types.ObjectNull(WirelessValue{}.AttributeTypes(ctx))

	if d != nil {
		if d.Wan != nil {
			wan = marvisSelfDrivingDomainSdkToTerraform(ctx, diags, d.Wan, func(at map[string]attr.Type, av map[string]attr.Value) (basetypes.ObjectValuable, diag.Diagnostics) {
				return NewWanValue(at, av)
			}, WanValue{}.AttributeTypes(ctx))
		}
		if d.Wired != nil {
			wired = marvisSelfDrivingDomainSdkToTerraform(ctx, diags, d.Wired, func(at map[string]attr.Type, av map[string]attr.Value) (basetypes.ObjectValuable, diag.Diagnostics) {
				return NewWiredValue(at, av)
			}, WiredValue{}.AttributeTypes(ctx))
		}
		if d.Wireless != nil {
			wireless = marvisSelfDrivingDomainSdkToTerraform(ctx, diags, d.Wireless, func(at map[string]attr.Type, av map[string]attr.Value) (basetypes.ObjectValuable, diag.Diagnostics) {
				return NewWirelessValue(at, av)
			}, WirelessValue{}.AttributeTypes(ctx))
		}
	}

	data, e := NewSelfDrivingValue(SelfDrivingValue{}.AttributeTypes(ctx), map[string]attr.Value{
		"wan":      wan,
		"wired":    wired,
		"wireless": wireless,
	})
	diags.Append(e...)
	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func marvisSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingMarvis) MarvisValue {
	var disableProactiveMonitoring types.Bool
	var selfDriving = types.ObjectNull(SelfDrivingValue{}.AttributeTypes(ctx))

	if d.DisableProactiveMonitoring != nil {
		disableProactiveMonitoring = types.BoolValue(*d.DisableProactiveMonitoring)
	}
	if d.SelfDriving != nil {
		selfDriving = marvisSelfDrivingSdkToTerraform(ctx, diags, d.SelfDriving)
	}

	data, e := NewMarvisValue(MarvisValue{}.AttributeTypes(ctx), map[string]attr.Value{
		"disable_proactive_monitoring": disableProactiveMonitoring,
		"self_driving":                 selfDriving,
	})
	diags.Append(e...)
	return data
}
