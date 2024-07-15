package resource_site_wlan

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func hotspot20TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan Hotspot20Value) *models.WlanHotspot20 {

	var operators []models.WlanHotspot20OperatorsItemEnum
	for _, v := range plan.Operators.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		op := models.WlanHotspot20OperatorsItemEnum(string(v_plan.ValueString()))
		operators = append(operators, op)
	}

	data := models.WlanHotspot20{}
	data.DomainName = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DomainName)
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.NaiRealms = mist_transform.ListOfStringTerraformToSdk(ctx, plan.NaiRealms)
	data.Operators = operators
	data.Rcoi = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Rcoi)
	data.VenueName = plan.VenueName.ValueStringPointer()

	return &data
}
