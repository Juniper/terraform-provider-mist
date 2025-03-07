package resource_org_gatewaytemplate

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ipConfigsTerraformToSdk(d basetypes.MapValue) map[string]models.GatewayIpConfigProperty {
	dataMap := make(map[string]models.GatewayIpConfigProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(IpConfigsValue)

		data := models.GatewayIpConfigProperty{}
		if plan.Ip.ValueStringPointer() != nil {
			data.Ip = models.ToPointer(plan.Ip.ValueString())
		}
		if plan.Netmask.ValueStringPointer() != nil {
			data.Netmask = models.ToPointer(plan.Netmask.ValueString())
		}
		if !plan.SecondaryIps.IsNull() && !plan.SecondaryIps.IsUnknown() {
			data.SecondaryIps = mistutils.ListOfStringTerraformToSdk(plan.SecondaryIps)
		}
		if !plan.IpConfigsType.IsNull() && !plan.IpConfigsType.IsUnknown() {
			data.Type = models.ToPointer(models.IpTypeEnum(plan.IpConfigsType.ValueString()))
		}
		dataMap[k] = data
	}
	return dataMap
}
