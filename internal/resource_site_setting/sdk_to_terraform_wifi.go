package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func wifiSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteWifi) WifiValue {
	var cisco_enabled basetypes.BoolValue
	var disable_11k basetypes.BoolValue
	var disable_radios_when_power_constrained basetypes.BoolValue
	var enable_arp_spoof_check basetypes.BoolValue
	var enable_shared_radio_scanning basetypes.BoolValue
	var enabled basetypes.BoolValue
	var locate_connected basetypes.BoolValue
	var locate_unconnected basetypes.BoolValue
	var mesh_allow_dfs basetypes.BoolValue
	var mesh_enable_crm basetypes.BoolValue
	var mesh_enabled basetypes.BoolValue
	var mesh_psk basetypes.StringValue
	var mesh_ssid basetypes.StringValue
	var proxy_arp basetypes.StringValue

	if d != nil && d.CiscoEnabled != nil {
		cisco_enabled = types.BoolValue(*d.CiscoEnabled)
	}
	if d != nil && d.Disable11k != nil {
		disable_11k = types.BoolValue(*d.Disable11k)
	}
	if d != nil && d.DisableRadiosWhenPowerConstrained != nil {
		disable_radios_when_power_constrained = types.BoolValue(*d.DisableRadiosWhenPowerConstrained)
	}
	if d != nil && d.EnableArpSpoofCheck != nil {
		enable_arp_spoof_check = types.BoolValue(*d.EnableArpSpoofCheck)
	}
	if d != nil && d.EnableSharedRadioScanning != nil {
		enable_shared_radio_scanning = types.BoolValue(*d.EnableSharedRadioScanning)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.LocateConnected != nil {
		locate_connected = types.BoolValue(*d.LocateConnected)
	}
	if d != nil && d.LocateUnconnected != nil {
		locate_unconnected = types.BoolValue(*d.LocateUnconnected)
	}
	if d != nil && d.MeshAllowDfs != nil {
		mesh_allow_dfs = types.BoolValue(*d.MeshAllowDfs)
	}
	if d != nil && d.MeshEnableCrm != nil {
		mesh_enable_crm = types.BoolValue(*d.MeshEnableCrm)
	}
	if d != nil && d.MeshEnabled != nil {
		mesh_enabled = types.BoolValue(*d.MeshEnabled)
	}
	if d != nil && d.MeshPsk.Value() != nil {
		mesh_psk = types.StringValue(*d.MeshPsk.Value())
	}
	if d != nil && d.MeshSsid.Value() != nil {
		mesh_ssid = types.StringValue(*d.MeshSsid.Value())
	}
	if d != nil && d.ProxyArp.Value() != nil {
		proxy_arp = types.StringValue(string(*d.ProxyArp.Value()))
	}

	data_map_attr_type := WifiValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"cisco_enabled":                         cisco_enabled,
		"disable_11k":                           disable_11k,
		"disable_radios_when_power_constrained": disable_radios_when_power_constrained,
		"enable_arp_spoof_check":                enable_arp_spoof_check,
		"enable_shared_radio_scanning":          enable_shared_radio_scanning,
		"enabled":                               enabled,
		"locate_connected":                      locate_connected,
		"locate_unconnected":                    locate_unconnected,
		"mesh_allow_dfs":                        mesh_allow_dfs,
		"mesh_enable_crm":                       mesh_enable_crm,
		"mesh_enabled":                          mesh_enabled,
		"mesh_psk":                              mesh_psk,
		"mesh_ssid":                             mesh_ssid,
		"proxy_arp":                             proxy_arp,
	}
	data, e := NewWifiValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
