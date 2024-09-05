package resource_site_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func gatewayMgmtAppProbingCustomSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.AppProbingCustomApp) basetypes.ListValue {
	var data_list = []CustomAppsValue{}
	for _, d := range l {
		var address basetypes.StringValue
		var app_type basetypes.StringValue
		var hostname basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var name basetypes.StringValue
		var network basetypes.StringValue
		var protocol basetypes.StringValue
		var url basetypes.StringValue
		var vrf basetypes.StringValue

		if d.Address != nil {
			address = types.StringValue(*d.Address)
		}
		if d.AppType != nil {
			app_type = types.StringValue(*d.AppType)
		}
		if d.Hostname != nil {
			hostname = mist_transform.ListOfStringSdkToTerraform(ctx, d.Hostname)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Network != nil {
			network = types.StringValue(*d.Network)
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

		data_map_attr_type := CustomAppsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"address":  address,
			"app_type": app_type,
			"hostname": hostname,
			"name":     name,
			"network":  network,
			"protocol": protocol,
			"url":      url,
			"vrf":      vrf,
		}
		data, e := NewCustomAppsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := VlansValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func gatewayMgmtAppProbingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.AppProbing) basetypes.ObjectValue {
	var apps basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var custom_apps basetypes.ListValue
	var enabled basetypes.BoolValue

	if d.Apps != nil {
		apps = mist_transform.ListOfStringSdkToTerraform(ctx, d.Apps)
	}
	if d.CustomApps != nil {
		custom_apps = gatewayMgmtAppProbingCustomSdkToTerraform(ctx, diags, d.CustomApps)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := AppProbingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"apps":        apps,
		"custom_apps": custom_apps,
		"enabled":     enabled,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func gatewayMgmtAutoSignatureUpdateSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingGatewayMgmtAutoSignatureUpdate) basetypes.ObjectValue {
	var day_of_week basetypes.StringValue
	var enable basetypes.BoolValue
	var time_of_day basetypes.StringValue

	if d.DayOfWeek != nil {
		day_of_week = types.StringValue(string(*d.DayOfWeek))
	}
	if d.Enable != nil {
		enable = types.BoolValue(*d.Enable)
	}
	if d.TimeOfDay != nil {
		time_of_day = types.StringValue(*d.TimeOfDay)
	}

	data_map_attr_type := AutoSignatureUpdateValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"day_of_week": day_of_week,
		"enable":      enable,
		"time_of_day": time_of_day,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
func gatewayMgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingGatewayMgmt) GatewayMgmtValue {
	var admin_sshkeys basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var app_probing basetypes.ObjectValue = types.ObjectNull(AppProbingValue{}.AttributeTypes(ctx))
	var app_usage basetypes.BoolValue
	var auto_signature_update basetypes.ObjectValue = types.ObjectNull(AutoSignatureUpdateValue{}.AttributeTypes(ctx))
	var config_revert_timer basetypes.Int64Value
	var disable_console basetypes.BoolValue
	var disable_oob basetypes.BoolValue
	var probe_hosts basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var root_password basetypes.StringValue
	var security_log_source_address basetypes.StringValue
	var security_log_source_interface basetypes.StringValue

	if d.AdminSshkeys != nil {
		admin_sshkeys = mist_transform.ListOfStringSdkToTerraform(ctx, d.AdminSshkeys)
	}
	if d.AppProbing != nil {
		app_probing = gatewayMgmtAppProbingSdkToTerraform(ctx, diags, d.AppProbing)
	}
	if d.AppUsage != nil {
		app_usage = types.BoolValue(*d.AppUsage)
	}
	if d.AutoSignatureUpdate != nil {
		auto_signature_update = gatewayMgmtAutoSignatureUpdateSdkToTerraform(ctx, diags, d.AutoSignatureUpdate)
	}
	if d.ConfigRevertTimer != nil {
		config_revert_timer = types.Int64Value(int64(*d.ConfigRevertTimer))
	}
	if d.DisableConsole != nil {
		disable_console = types.BoolValue(*d.DisableConsole)
	}
	if d.DisableOob != nil {
		disable_oob = types.BoolValue(*d.DisableOob)
	}
	if d.ProbeHosts != nil {
		probe_hosts = mist_transform.ListOfStringSdkToTerraform(ctx, d.ProbeHosts)
	}
	if d.RootPassword != nil {
		root_password = types.StringValue(*d.RootPassword)
	}
	if d.SecurityLogSourceAddress != nil {
		security_log_source_address = types.StringValue(*d.SecurityLogSourceAddress)
	}
	if d.SecurityLogSourceInterface != nil {
		security_log_source_interface = types.StringValue(*d.SecurityLogSourceInterface)
	}

	data_map_attr_type := GatewayMgmtValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"admin_sshkeys":                 admin_sshkeys,
		"app_probing":                   app_probing,
		"app_usage":                     app_usage,
		"auto_signature_update":         auto_signature_update,
		"config_revert_timer":           config_revert_timer,
		"disable_console":               disable_console,
		"disable_oob":                   disable_oob,
		"probe_hosts":                   probe_hosts,
		"root_password":                 root_password,
		"security_log_source_address":   security_log_source_address,
		"security_log_source_interface": security_log_source_interface,
	}
	data, e := NewGatewayMgmtValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
