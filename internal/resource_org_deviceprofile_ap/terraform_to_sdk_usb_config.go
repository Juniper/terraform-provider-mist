package resource_org_deviceprofile_ap

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func usbConfigTerraformToSdk(d UsbConfigValue) *models.ApUsb {
	data := models.ApUsb{}

	if d.Cacert.ValueStringPointer() != nil {
		data.Cacert = models.NewOptional(d.Cacert.ValueStringPointer())
	}
	if d.Channel.ValueInt64Pointer() != nil {
		data.Channel = models.ToPointer(int(d.Channel.ValueInt64()))
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Host.ValueStringPointer() != nil {
		data.Host = d.Host.ValueStringPointer()
	}
	if d.Port.ValueInt64Pointer() != nil {
		data.Port = models.ToPointer(int(d.Port.ValueInt64()))
	}
	if d.UsbConfigType.ValueStringPointer() != nil {
		data.Type = models.ToPointer(models.ApUsbTypeEnum(d.UsbConfigType.ValueString()))
	}
	if d.VerifyCert.ValueBoolPointer() != nil {
		data.VerifyCert = d.VerifyCert.ValueBoolPointer()
	}
	if d.VlanId.ValueInt64Pointer() != nil {
		data.VlanId = models.ToPointer(int(d.VlanId.ValueInt64()))
	}

	return &data
}
