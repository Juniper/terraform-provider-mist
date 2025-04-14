package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func wifiTerraformToSdk(d WifiValue) *models.SiteWifi {
	data := models.SiteWifi{}

	if d.CiscoEnabled.ValueBoolPointer() != nil {
		data.CiscoEnabled = d.CiscoEnabled.ValueBoolPointer()
	}

	if d.Disable11k.ValueBoolPointer() != nil {
		data.Disable11k = d.Disable11k.ValueBoolPointer()
	}

	if d.DisableRadiosWhenPowerConstrained.ValueBoolPointer() != nil {
		data.DisableRadiosWhenPowerConstrained = d.DisableRadiosWhenPowerConstrained.ValueBoolPointer()
	}

	if d.EnableArpSpoofCheck.ValueBoolPointer() != nil {
		data.EnableArpSpoofCheck = d.EnableArpSpoofCheck.ValueBoolPointer()
	}

	if d.EnableSharedRadioScanning.ValueBoolPointer() != nil {
		data.EnableSharedRadioScanning = d.EnableSharedRadioScanning.ValueBoolPointer()
	}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if d.LocateConnected.ValueBoolPointer() != nil {
		data.LocateConnected = d.LocateConnected.ValueBoolPointer()
	}

	if d.LocateUnconnected.ValueBoolPointer() != nil {
		data.LocateUnconnected = d.LocateUnconnected.ValueBoolPointer()
	}

	if d.MeshAllowDfs.ValueBoolPointer() != nil {
		data.MeshAllowDfs = d.MeshAllowDfs.ValueBoolPointer()
	}

	if d.MeshEnableCrm.ValueBoolPointer() != nil {
		data.MeshEnableCrm = d.MeshEnableCrm.ValueBoolPointer()
	}

	if d.MeshEnabled.ValueBoolPointer() != nil {
		data.MeshEnabled = d.MeshEnabled.ValueBoolPointer()
	}

	if d.MeshPsk.ValueStringPointer() != nil {
		data.MeshPsk = models.NewOptional(d.MeshPsk.ValueStringPointer())
	}

	if d.MeshSsid.ValueStringPointer() != nil {
		data.MeshSsid = models.NewOptional(d.MeshSsid.ValueStringPointer())
	}

	if d.ProxyArp.ValueStringPointer() != nil {
		data.ProxyArp = models.NewOptional(models.ToPointer(models.SiteWifiProxyArpEnum(d.ProxyArp.ValueString())))
	}

	return &data
}
