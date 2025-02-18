package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func deviceCertTerraformToSdk(d DeviceCertValue) *models.OrgSettingDeviceCert {
	data := models.OrgSettingDeviceCert{}

	if d.Cert.ValueStringPointer() != nil {
		data.Cert = d.Cert.ValueStringPointer()
	}

	if d.Key.ValueStringPointer() != nil {
		data.Key = d.Key.ValueStringPointer()
	}

	return &data
}
