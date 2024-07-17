package resource_org_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func mgmtTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MgmtValue) *models.OrgSettingMgmt {
	data := models.OrgSettingMgmt{}

	if !d.MxtunnelIds.IsNull() && !d.MxtunnelIds.IsUnknown() {
		data.MxtunnelIds = mist_transform.ListOfUuidTerraformToSdk(ctx, d.MxtunnelIds)
	}

	if d.UseMxtunnel.ValueBoolPointer() != nil {
		data.UseMxtunnel = d.UseMxtunnel.ValueBoolPointer()
	}

	if d.UseWxtunnel.ValueBoolPointer() != nil {
		data.UseWxtunnel = d.UseWxtunnel.ValueBoolPointer()
	}

	return &data
}
