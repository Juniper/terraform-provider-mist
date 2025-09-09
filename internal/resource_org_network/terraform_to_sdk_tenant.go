package resource_org_network

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TenantTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkTenant {
	dataMap := make(map[string]models.NetworkTenant)
	for k, v := range d.Elements() {
		// Extract attributes directly from the ObjectValue instead of casting to specific type
		if objVal, ok := v.(basetypes.ObjectValue); ok {
			data := models.NetworkTenant{}
			attrs := objVal.Attributes()

			if addresses, exists := attrs["addresses"]; exists {
				if listVal, ok := addresses.(basetypes.ListValue); ok {
					data.Addresses = mistutils.ListOfStringTerraformToSdk(listVal)
				}
			}
			dataMap[k] = data
		}
	}
	return dataMap
}
