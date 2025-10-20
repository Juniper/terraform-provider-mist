package resource_org_network

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tenantTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkTenant {
	dataMap := make(map[string]models.NetworkTenant)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(TenantsValue)
		data := models.NetworkTenant{}
		data.Addresses = mistutils.ListOfStringTerraformToSdk(vPlan.Addresses)
		dataMap[k] = data
	}
	return dataMap
}
