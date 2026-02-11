package resource_org_gatewaytemplate

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vrfConfigTerraformToSdk(data VrfConfigValue) *models.VrfConfig {
	if data.IsNull() || data.Enabled.ValueBoolPointer() == nil {
		return &models.VrfConfig{}
	}

	return &models.VrfConfig{
		Enabled: data.Enabled.ValueBoolPointer(),
	}
}

func vrfInstancesTerraformToSdk(data basetypes.MapValue) map[string]models.GatewayVrfInstance {
	result := make(map[string]models.GatewayVrfInstance)
	for key, val := range data.Elements() {
		item := val.(VrfInstancesValue)
		if item.Networks.IsUnknown() {
			continue
		}

		if item.Networks.IsNull() {
			result[key] = models.GatewayVrfInstance{}
			continue
		}

		result[key] = models.GatewayVrfInstance{
			Networks: mistutils.ListOfStringTerraformToSdk(item.Networks),
		}
	}
	return result
}
