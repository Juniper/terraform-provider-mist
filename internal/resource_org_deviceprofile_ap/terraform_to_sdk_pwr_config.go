package resource_org_deviceprofile_ap

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func pwrConfigTerraformToSdk(d PwrConfigValue) *models.ApPwrConfig {
	data := models.ApPwrConfig{}

	if d.Base.ValueInt64Pointer() != nil {
		data.Base = models.ToPointer(int(d.Base.ValueInt64()))
	}
	if d.PreferUsbOverWifi.ValueBoolPointer() != nil {
		data.PreferUsbOverWifi = d.PreferUsbOverWifi.ValueBoolPointer()
	}

	return &data
}
