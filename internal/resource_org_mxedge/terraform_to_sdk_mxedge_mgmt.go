package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func mxedgeMgmtTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MxedgeMgmtValue) *models.MxedgeMgmt {
	data := models.MxedgeMgmt{}

	if !d.ConfigAutoRevert.IsNull() && !d.ConfigAutoRevert.IsUnknown() {
		data.ConfigAutoRevert = d.ConfigAutoRevert.ValueBoolPointer()
	}

	if !d.FipsEnabled.IsNull() && !d.FipsEnabled.IsUnknown() {
		data.FipsEnabled = d.FipsEnabled.ValueBoolPointer()
	}

	if !d.MistPassword.IsNull() && !d.MistPassword.IsUnknown() {
		data.MistPassword = d.MistPassword.ValueStringPointer()
	}

	if !d.OobIpType.IsNull() && !d.OobIpType.IsUnknown() {
		data.OobIpType = (*models.MxedgeMgmtOobIpTypeEnum)(d.OobIpType.ValueStringPointer())
	}

	if !d.OobIpType6.IsNull() && !d.OobIpType6.IsUnknown() {
		data.OobIpType6 = (*models.MxedgeMgmtOobIpType6Enum)(d.OobIpType6.ValueStringPointer())
	}

	if !d.RootPassword.IsNull() && !d.RootPassword.IsUnknown() {
		data.RootPassword = d.RootPassword.ValueStringPointer()
	}

	return &data
}
