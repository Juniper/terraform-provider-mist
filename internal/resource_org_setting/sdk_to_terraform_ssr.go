package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func ssrProxySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SsrProxy) basetypes.ObjectValue {

	var disabled basetypes.BoolValue
	var url basetypes.StringValue

	if d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d.Url != nil {
		url = types.StringValue(*d.Url)
	}

	dataMapValue := map[string]attr.Value{
		"disabled": disabled,
		"url":      url,
	}
	data, e := NewProxyValue(ProxyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o

}

func ssrAutoUpgradeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SettingSsrAutoUpgrade) basetypes.ObjectValue {

	var channel basetypes.StringValue
	var customVersions = types.MapNull(types.StringType)
	var enabled basetypes.BoolValue
	var version basetypes.StringValue

	if d.Channel != nil {
		channel = types.StringValue(string(*d.Channel))
	}
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
	if d.Version != nil {
		version = types.StringValue(*d.Version)
	}

	dataMapValue := map[string]attr.Value{
		"channel":         channel,
		"custom_versions": customVersions,
		"enabled":         enabled,
		"version":         version,
	}
	data, e := NewSsrAutoUpgradeValue(SsrAutoUpgradeValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o

}

func ssrSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SettingSsr) SsrValue {

	var conductorHosts = types.ListNull(types.StringType)
	var conductorToken basetypes.StringValue
	var disableStats basetypes.BoolValue
	var proxy = types.ObjectNull(ProxyValue{}.AttributeTypes(ctx))
	var ssrAutoUpgrade = types.ObjectNull(SsrAutoUpgradeValue{}.AttributeTypes(ctx))

	if d != nil && d.ConductorHosts != nil {
		conductorHosts = mistutils.ListOfStringSdkToTerraform(d.ConductorHosts)
	}
	if d != nil && d.ConductorToken != nil {
		conductorToken = types.StringValue(*d.ConductorToken)
	}
	if d != nil && d.DisableStats != nil {
		disableStats = types.BoolValue(*d.DisableStats)
	}
	if d != nil && d.Proxy != nil {
		proxy = ssrProxySdkToTerraform(ctx, diags, *d.Proxy)
	}
	if d != nil && d.AutoUpgrade != nil {
		ssrAutoUpgrade = ssrAutoUpgradeSdkToTerraform(ctx, diags, *d.AutoUpgrade)
	}

	dataMapValue := map[string]attr.Value{
		"conductor_hosts": conductorHosts,
		"conductor_token": conductorToken,
		"disable_stats":   disableStats,
		"proxy":           proxy,
		"auto_upgrade":    ssrAutoUpgrade,
	}
	data, e := NewSsrValue(SsrValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
