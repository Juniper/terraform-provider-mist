package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func juniperSrxAutoUpgradeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.JuniperSrxAutoUpgrade) basetypes.ObjectValue {

	var customVersions = types.MapNull(types.StringType)
	var enabled basetypes.BoolValue
	var snapshot basetypes.BoolValue
	var version basetypes.StringValue

	if d.CustomVersions != nil {
		rMapValue := make(map[string]attr.Value)
		for k, v := range d.CustomVersions {
			rMapValue[k] = types.StringValue(v)
		}
		m, e := types.MapValueFrom(ctx, types.StringType, rMapValue)
		diags.Append(e...)
		customVersions = m
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Snapshot != nil {
		snapshot = types.BoolValue(*d.Snapshot)
	}
	if d.Version != nil {
		version = types.StringValue(*d.Version)
	}

	dataMapValue := map[string]attr.Value{
		"custom_versions": customVersions,
		"enabled":         enabled,
		"snapshot":        snapshot,
		"version":         version,
	}
	data, e := NewSrxAutoUpgradeValue(SrxAutoUpgradeValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o

}

func juniperSrxSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingJuniperSrx) JuniperSrxValue {
	var autoUpgrade = types.ObjectNull(SrxAutoUpgradeValue{}.AttributeTypes(ctx))

	if d.AutoUpgrade != nil {
		autoUpgrade = juniperSrxAutoUpgradeSdkToTerraform(ctx, diags, *d.AutoUpgrade)
	}

	dataMapValue := map[string]attr.Value{
		"auto_upgrade": autoUpgrade,
	}
	data, e := NewJuniperSrxValue(JuniperSrxValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
