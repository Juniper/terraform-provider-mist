package resource_device_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func DeleteTerraformToSdk(ctx context.Context) (models.MistDevice, diag.Diagnostics) {
	var diags diag.Diagnostics
	//var data models.SiteSetting
	data := models.DeviceAp{}

	tmp := DeviceApResourceSchema(ctx)
	unset := make(map[string]interface{})
	for k := range tmp.Attributes {
		unset["-"+k] = ""
	}

	data.Type = string(models.DeviceTypeEnum_AP)
	data.AdditionalProperties = unset
	mistDevice := models.MistDeviceContainer.FromDeviceAp(data)
	return mistDevice, diags
}
