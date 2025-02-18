package resource_org_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func qosTerraformToSdk(d QosValue) *models.WlanQos {
	data := models.WlanQos{}
	data.Class = models.ToPointer(models.WlanQosClassEnum(d.Class.ValueString()))
	data.Overwrite = d.Overwrite.ValueBoolPointer()

	return &data
}
