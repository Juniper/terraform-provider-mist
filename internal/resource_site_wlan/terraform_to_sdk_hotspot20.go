package resource_site_wlan

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func hotspot20TerraformToSdk(plan Hotspot20Value) *models.WlanHotspot20 {

	var operators []models.WlanHotspot20OperatorsItemEnum
	for _, v := range plan.Operators.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(basetypes.StringValue)
		op := models.WlanHotspot20OperatorsItemEnum(vPlan.ValueString())
		operators = append(operators, op)
	}

	data := models.WlanHotspot20{}
	if !plan.DomainName.IsNull() && !plan.DomainName.IsUnknown() {
		data.DomainName = mistutils.ListOfStringTerraformToSdk(plan.DomainName)
	}
	if plan.Enabled.ValueBoolPointer() != nil {
		data.Enabled = plan.Enabled.ValueBoolPointer()
	}
	if !plan.NaiRealms.IsNull() && !plan.NaiRealms.IsUnknown() {
		data.NaiRealms = mistutils.ListOfStringTerraformToSdk(plan.NaiRealms)
	}
	data.Operators = operators
	if !plan.Rcoi.IsNull() && !plan.Rcoi.IsUnknown() {
		data.Rcoi = mistutils.ListOfStringTerraformToSdk(plan.Rcoi)
	}
	if plan.VenueName.ValueStringPointer() != nil {
		data.VenueName = plan.VenueName.ValueStringPointer()
	}

	return &data
}
