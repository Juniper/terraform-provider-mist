package resource_org_wlan

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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
	if !plan.DomainName.IsNull() && !plan.DomainName.IsUnknown() {
		data.DomainName = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DomainName)
	}
	if plan.Enabled.ValueBoolPointer() != nil {
		data.Enabled = plan.Enabled.ValueBoolPointer()
	}
	if !plan.NaiRealms.IsNull() && !plan.NaiRealms.IsUnknown() {
		data.NaiRealms = mist_transform.ListOfStringTerraformToSdk(ctx, plan.NaiRealms)
	}
	data.Operators = operators
	if !plan.Rcoi.IsNull() && !plan.Rcoi.IsUnknown() {
		data.Rcoi = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Rcoi)
	}
	if plan.VenueName.ValueStringPointer() != nil {
		data.VenueName = plan.VenueName.ValueStringPointer()
	}

	return &data
}
