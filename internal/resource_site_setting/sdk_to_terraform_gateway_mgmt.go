package resource_site_setting

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func gatewayMgmtProtectCustomReSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ProtectReCustom) basetypes.ListValue {
	var dataList []CustomValue

	for _, d := range l {

		var portRange basetypes.StringValue
		var protocol basetypes.StringValue
		var subnets = mistutils.ListOfStringSdkToTerraformEmpty()

		if d.PortRange != nil {
			portRange = types.StringValue(*d.PortRange)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(string(*d.Protocol))
		}
		if d.Subnets != nil {
			subnets = mistutils.ListOfStringSdkToTerraform(d.Subnets)
		}

		dataMapValue := map[string]attr.Value{
			"port_range": portRange,
			"protocol":   protocol,
			"subnets":    subnets,
		}
		data, e := NewCustomValue(CustomValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := CustomValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)

	return r
}
func gatewayMgmtProtectReSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ProtectRe) basetypes.ObjectValue {
	var allowedServices = types.ListNull(types.StringType)
	var custom = basetypes.NewListValueMust(CustomValue{}.Type(ctx), []attr.Value{})
	var enabled basetypes.BoolValue
	var trustedHosts = types.ListNull(types.StringType)

	if d.AllowedServices != nil {
		var items []attr.Value
		var itemsType attr.Type = basetypes.StringType{}
		for _, item := range d.AllowedServices {
			items = append(items, types.StringValue(string(item)))
		}
		list, _ := types.ListValue(itemsType, items)
		allowedServices = list
	}
	if d.Custom != nil {
		custom = gatewayMgmtProtectCustomReSdkToTerraform(ctx, diags, d.Custom)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TrustedHosts != nil {
		trustedHosts = mistutils.ListOfStringSdkToTerraform(d.TrustedHosts)
	}

	dataMapValue := map[string]attr.Value{
		"allowed_services": allowedServices,
		"custom":           custom,
		"enabled":          enabled,
		"trusted_hosts":    trustedHosts,
	}
	data, e := NewProtectReValue(ProtectReValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func gatewayMgmtAppProbingCustomSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.AppProbingCustomApp) basetypes.ListValue {
	var dataList []CustomAppsValue
	for _, d := range l {
		var address basetypes.StringValue
		var appType basetypes.StringValue
		var hostnames = mistutils.ListOfStringSdkToTerraformEmpty()
		var key basetypes.StringValue
		var name basetypes.StringValue
		var network basetypes.StringValue
		var packetSize basetypes.Int64Value
		var protocol basetypes.StringValue
		var url basetypes.StringValue
		var vrf basetypes.StringValue

		if d.Address != nil {
			address = types.StringValue(*d.Address)
		}
		if d.AppType != nil {
			appType = types.StringValue(*d.AppType)
		}
		if d.Hostnames != nil {
			hostnames = mistutils.ListOfStringSdkToTerraform(d.Hostnames)
		}
		if d.Key != nil {
			key = types.StringValue(*d.Key)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Network != nil {
			network = types.StringValue(*d.Network)
		}
		if d.PacketSize != nil {
			packetSize = types.Int64Value(int64(*d.PacketSize))
		}
		if d.Protocol != nil {
			protocol = types.StringValue(string(*d.Protocol))
		}
		if d.Url != nil {
			url = types.StringValue(*d.Url)
		}
		if d.Vrf != nil {
			vrf = types.StringValue(*d.Vrf)
		}

		dataMapValue := map[string]attr.Value{
			"address":     address,
			"app_type":    appType,
			"hostnames":   hostnames,
			"key":         key,
			"name":        name,
			"network":     network,
			"packet_size": packetSize,
			"protocol":    protocol,
			"url":         url,
			"vrf":         vrf,
		}
		data, e := NewCustomAppsValue(CustomAppsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := CustomAppsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}

func gatewayMgmtAppProbingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.AppProbing) basetypes.ObjectValue {
	var apps = types.ListNull(types.StringType)
	var customApps = types.ListNull(CustomAppsValue{}.Type(ctx))
	var enabled basetypes.BoolValue

	if d.Apps != nil {
		apps = mistutils.ListOfStringSdkToTerraform(d.Apps)
	}
	if d.CustomApps != nil {
		customApps = gatewayMgmtAppProbingCustomSdkToTerraform(ctx, diags, d.CustomApps)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"apps":        apps,
		"custom_apps": customApps,
		"enabled":     enabled,
	}
	data, e := basetypes.NewObjectValue(AppProbingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func gatewayMgmtAutoSignatureUpdateSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingGatewayMgmtAutoSignatureUpdate) basetypes.ObjectValue {
	var dayOfWeek basetypes.StringValue
	var enable basetypes.BoolValue
	var timeOfDay basetypes.StringValue

	if d.DayOfWeek != nil {
		dayOfWeek = types.StringValue(string(*d.DayOfWeek))
	}
	if d.Enable != nil {
		enable = types.BoolValue(*d.Enable)
	}
	if d.TimeOfDay != nil {
		timeOfDay = types.StringValue(*d.TimeOfDay)
	}

	dataMapValue := map[string]attr.Value{
		"day_of_week": dayOfWeek,
		"enable":      enable,
		"time_of_day": timeOfDay,
	}
	data, e := basetypes.NewObjectValue(AutoSignatureUpdateValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func gatewayMgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingGatewayMgmt) GatewayMgmtValue {
	var adminSshkeys = types.ListNull(types.StringType)
	var appProbing = types.ObjectNull(AppProbingValue{}.AttributeTypes(ctx))
	var appUsage basetypes.BoolValue
	var autoSignatureUpdate = types.ObjectNull(AutoSignatureUpdateValue{}.AttributeTypes(ctx))
	var configRevertTimer basetypes.Int64Value
	var disableConsole basetypes.BoolValue
	var disableOob basetypes.BoolValue
	var disableUsb basetypes.BoolValue
	var fipsEnabled basetypes.BoolValue
	var probeHosts = types.ListNull(types.StringType)
	var protectRe = types.ObjectNull(ProtectReValue{}.AttributeTypes(ctx))
	var rootPassword basetypes.StringValue
	var securityLogSourceAddress basetypes.StringValue
	var securityLogSourceInterface basetypes.StringValue

	if d.AdminSshkeys != nil {
		adminSshkeys = mistutils.ListOfStringSdkToTerraform(d.AdminSshkeys)
	}
	if d.AppProbing != nil {
		appProbing = gatewayMgmtAppProbingSdkToTerraform(ctx, diags, d.AppProbing)
	}
	if d.AppUsage != nil {
		appUsage = types.BoolValue(*d.AppUsage)
	}
	if d.AutoSignatureUpdate != nil {
		autoSignatureUpdate = gatewayMgmtAutoSignatureUpdateSdkToTerraform(ctx, diags, d.AutoSignatureUpdate)
	}
	if d.ConfigRevertTimer != nil {
		configRevertTimer = types.Int64Value(int64(*d.ConfigRevertTimer))
	}
	if d.DisableConsole != nil {
		disableConsole = types.BoolValue(*d.DisableConsole)
	}
	if d.DisableUsb != nil {
		disableUsb = types.BoolValue(*d.DisableUsb)
	}
	if d.FipsEnabled != nil {
		fipsEnabled = types.BoolValue(*d.FipsEnabled)
	}
	if d.DisableOob != nil {
		disableOob = types.BoolValue(*d.DisableOob)
	}
	if d.ProbeHosts != nil {
		probeHosts = mistutils.ListOfStringSdkToTerraform(d.ProbeHosts)
	}
	if d.ProtectRe != nil {
		protectRe = gatewayMgmtProtectReSdkToTerraform(ctx, diags, d.ProtectRe)
	}
	if d.RootPassword != nil {
		rootPassword = types.StringValue(*d.RootPassword)
	}
	if d.SecurityLogSourceAddress != nil {
		securityLogSourceAddress = types.StringValue(*d.SecurityLogSourceAddress)
	}
	if d.SecurityLogSourceInterface != nil {
		securityLogSourceInterface = types.StringValue(*d.SecurityLogSourceInterface)
	}

	dataMapValue := map[string]attr.Value{
		"admin_sshkeys":                 adminSshkeys,
		"app_probing":                   appProbing,
		"app_usage":                     appUsage,
		"auto_signature_update":         autoSignatureUpdate,
		"config_revert_timer":           configRevertTimer,
		"disable_console":               disableConsole,
		"disable_usb":                   disableUsb,
		"fips_enabled":                  fipsEnabled,
		"disable_oob":                   disableOob,
		"probe_hosts":                   probeHosts,
		"protect_re":                    protectRe,
		"root_password":                 rootPassword,
		"security_log_source_address":   securityLogSourceAddress,
		"security_log_source_interface": securityLogSourceInterface,
	}
	data, e := NewGatewayMgmtValue(GatewayMgmtValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
