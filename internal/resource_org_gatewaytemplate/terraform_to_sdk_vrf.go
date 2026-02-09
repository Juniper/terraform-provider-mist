package resource_org_gatewaytemplate

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vrfConfigTerraformToSdk(d VrfConfigValue) *models.VrfConfig {
	if d.IsNull() || d.Enabled.ValueBoolPointer() == nil {
		return &models.VrfConfig{}
	}

	return &models.VrfConfig{
		Enabled: d.Enabled.ValueBoolPointer(),
	}
}

func vrfInstancesTerraformToSdk(data basetypes.MapValue) map[string]models.GatewayVrfInstance {
	result := make(map[string]models.GatewayVrfInstance)
	for itemName, itemValue := range data.Elements() {
		itemObj := itemValue.(VrfInstancesValue)
		if itemObj.Networks.IsNull() || itemObj.Networks.IsUnknown() {
			continue
		}

		item := models.GatewayVrfInstance{
			Networks: mistutils.ListOfStringTerraformToSdk(itemObj.Networks),
		}
		result[itemName] = item
	}
	return result
}
