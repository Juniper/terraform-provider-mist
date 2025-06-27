package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func switchAutoUpgradesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchAutoUpgrade) basetypes.ObjectValue {
	var customVersions = types.MapNull(types.StringType)
	var enabled basetypes.BoolValue
	var snapshot basetypes.BoolValue

	if d.CustomVersions != nil {
		var items []attr.Value
		var itemsType attr.Type = basetypes.StringType{}
		for k, v := range d.CustomVersions {
			items = append(items, types.StringValue(k+":"+v))
		}
		tmp, e := types.MapValueFrom(ctx, itemsType, items)
		if e != nil {
			diags.Append(e...)
		} else {
			customVersions = tmp
		}
	}

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	if d.Snapshot != nil {
		snapshot = types.BoolValue(*d.Snapshot)
	}

	dataMapValue := map[string]attr.Value{
		"custom_versions": customVersions,
		"enabled":         enabled,
		"snapshot":        snapshot,
	}
	data, e := NewAutoUpgradeValue(AutoUpgradeValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func switchSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingSwitch) SwitchValue {

	var autoUpgrade = types.ObjectNull(AutoUpgradeValue{}.AttributeTypes(ctx))

	if d != nil && d.AutoUpgrade != nil {
		autoUpgrade = switchAutoUpgradesSdkToTerraform(ctx, diags, d.AutoUpgrade)
	}

	dataMapValue := map[string]attr.Value{
		"auto_upgrade": autoUpgrade,
	}
	data, e := NewSwitchValue(SwitchValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
