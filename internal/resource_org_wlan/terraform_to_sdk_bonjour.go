package resource_org_wlan

import (
	"strings"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bonjourServicesTerraformToSdk(plan basetypes.MapValue) map[string]models.WlanBonjourServiceProperties {
	dataMap := make(map[string]models.WlanBonjourServiceProperties)
	for k, v := range plan.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(ServicesValue)
		vData := models.WlanBonjourServiceProperties{}
		if vPlan.DisableLocal.ValueBoolPointer() != nil {
			vData.DisableLocal = vPlan.DisableLocal.ValueBoolPointer()
		}
		if !vPlan.RadiusGroups.IsNull() && !vPlan.RadiusGroups.IsUnknown() {
			vData.RadiusGroups = mistutils.ListOfStringTerraformToSdk(vPlan.RadiusGroups)
		}
		if vPlan.Scope.ValueStringPointer() != nil {
			vData.Scope = models.ToPointer(models.WlanBonjourServicePropertiesScopeEnum(vPlan.Scope.ValueString()))
		}
		dataMap[k] = vData
	}
	return dataMap
}
func bonjourTerraformToSdk(plan BonjourValue) *models.WlanBonjour {

	data := models.WlanBonjour{}

	var tmp []string
	for _, v := range plan.AdditionalVlanIds.Elements() {
		var i interface{} = v
		s := i.(basetypes.StringValue)
		tmp = append(tmp, s.ValueString())
	}
	data.AdditionalVlanIds = strings.Join(tmp, ",")
	data.Services = bonjourServicesTerraformToSdk(plan.Services)
	data.Enabled = plan.Enabled.ValueBoolPointer()

	return &data
}
