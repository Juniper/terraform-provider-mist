package resource_org_deviceprofile_ap

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func uwbConfigTerraformToSdk(d UwbConfigValue) *models.ApUwbConfig {
	data := models.ApUwbConfig{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Host.ValueStringPointer() != nil {
		data.Host = d.Host.ValueStringPointer()
	}
	if d.Port.ValueInt64Pointer() != nil {
		data.Port = models.ToPointer(int(d.Port.ValueInt64()))
	}
	if d.Slot.ValueInt64Pointer() != nil {
		data.Slot = models.ToPointer(int(d.Slot.ValueInt64()))
	}
	if d.UwbConfigType.ValueStringPointer() != nil {
		data.Type = models.ToPointer(models.ApUwbConfigTypeEnum(d.UwbConfigType.ValueString()))
	}

	return &data
}
