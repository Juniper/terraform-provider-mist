package resource_device_gateway

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vrfConfigTerraformToSdk(d VrfConfigValue) *models.VrfConfig {
	data := models.VrfConfig{}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	return &data
}

func vrfInstancesTerraformToSdk(d basetypes.MapValue) map[string]models.GatewayVrfInstance {
	data := make(map[string]models.GatewayVrfInstance)
	for itemName, itemValue := range d.Elements() {
		itemObj := itemValue.(VrfInstancesValue)
		if itemObj.Networks.IsNull() || itemObj.Networks.IsUnknown() {
			continue
		}

		dataItem := models.GatewayVrfInstance{
			Networks: mistutils.ListOfStringTerraformToSdk(itemObj.Networks),
		}
		data[itemName] = dataItem
	}
	return data
}
