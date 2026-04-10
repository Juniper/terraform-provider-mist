package resource_org_setting

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
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

	if len(d.CustomVersions) > 0 {
		rMapValue := make(map[string]attr.Value)
		for k, v := range d.CustomVersions {
			rMapValue[k] = types.StringValue(v)
		}
		m, e := types.MapValueFrom(ctx, types.StringType, rMapValue)
		if !e.HasError() {
			customVersions = m
		} else {
			diags.Append(e...)
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

	if d != nil && !mistutils.IsSdkDataEmpty(d.AutoUpgrade) {
		autoUpgrade = switchAutoUpgradesSdkToTerraform(ctx, diags, d.AutoUpgrade)
	}

	dataMapValue := map[string]attr.Value{
		"auto_upgrade": autoUpgrade,
	}
	data, e := NewSwitchValue(SwitchValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
