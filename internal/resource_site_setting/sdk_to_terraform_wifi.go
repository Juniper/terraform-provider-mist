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
	var ciscoEnabled basetypes.BoolValue
	var disable11k basetypes.BoolValue
	var disableRadiosWhenPowerConstrained basetypes.BoolValue
	var enableArpSpoofCheck basetypes.BoolValue
	var enableSharedRadioScanning basetypes.BoolValue
	var enabled basetypes.BoolValue
	var locateConnected basetypes.BoolValue
	var locateUnconnected basetypes.BoolValue
	var meshAllowDfs basetypes.BoolValue
	var meshEnableCrm basetypes.BoolValue
	var meshEnabled basetypes.BoolValue
	var meshPsk basetypes.StringValue
	var meshSsid basetypes.StringValue
	var proxyArp basetypes.StringValue

	if d != nil && d.CiscoEnabled != nil {
		ciscoEnabled = types.BoolValue(*d.CiscoEnabled)
	}
	if d != nil && d.Disable11k != nil {
		disable11k = types.BoolValue(*d.Disable11k)
	}
	if d != nil && d.DisableRadiosWhenPowerConstrained != nil {
		disableRadiosWhenPowerConstrained = types.BoolValue(*d.DisableRadiosWhenPowerConstrained)
	}
	if d != nil && d.EnableArpSpoofCheck != nil {
		enableArpSpoofCheck = types.BoolValue(*d.EnableArpSpoofCheck)
	}
	if d != nil && d.EnableSharedRadioScanning != nil {
		enableSharedRadioScanning = types.BoolValue(*d.EnableSharedRadioScanning)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.LocateConnected != nil {
		locateConnected = types.BoolValue(*d.LocateConnected)
	}
	if d != nil && d.LocateUnconnected != nil {
		locateUnconnected = types.BoolValue(*d.LocateUnconnected)
	}
	if d != nil && d.MeshAllowDfs != nil {
		meshAllowDfs = types.BoolValue(*d.MeshAllowDfs)
	}
	if d != nil && d.MeshEnableCrm != nil {
		meshEnableCrm = types.BoolValue(*d.MeshEnableCrm)
	}
	if d != nil && d.MeshEnabled != nil {
		meshEnabled = types.BoolValue(*d.MeshEnabled)
	}
	if d != nil && d.MeshPsk.Value() != nil {
		meshPsk = types.StringValue(*d.MeshPsk.Value())
	}
	if d != nil && d.MeshSsid.Value() != nil {
		meshSsid = types.StringValue(*d.MeshSsid.Value())
	}
	if d != nil && d.ProxyArp.Value() != nil {
		proxyArp = types.StringValue(string(*d.ProxyArp.Value()))
	}

	dataMapValue := map[string]attr.Value{
		"cisco_enabled":                         ciscoEnabled,
		"disable_11k":                           disable11k,
		"disable_radios_when_power_constrained": disableRadiosWhenPowerConstrained,
		"enable_arp_spoof_check":                enableArpSpoofCheck,
		"enable_shared_radio_scanning":          enableSharedRadioScanning,
		"enabled":                               enabled,
		"locate_connected":                      locateConnected,
		"locate_unconnected":                    locateUnconnected,
		"mesh_allow_dfs":                        meshAllowDfs,
		"mesh_enable_crm":                       meshEnableCrm,
		"mesh_enabled":                          meshEnabled,
		"mesh_psk":                              meshPsk,
		"mesh_ssid":                             meshSsid,
		"proxy_arp":                             proxyArp,
	}
	data, e := NewWifiValue(WifiValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
