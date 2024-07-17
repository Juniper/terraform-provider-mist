package resource_org_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func installerTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d InstallerValue) *models.OrgSettingInstaller {
	data := models.OrgSettingInstaller{}

	if d.AllowAllDevices.ValueBoolPointer() != nil {
		data.AllowAllDevices = d.AllowAllDevices.ValueBoolPointer()
	}

	if d.AllowAllSites.ValueBoolPointer() != nil {
		data.AllowAllSites = d.AllowAllSites.ValueBoolPointer()
	}

	if !d.ExtraSiteIds.IsNull() && !d.ExtraSiteIds.IsUnknown() {
		data.ExtraSiteIds = mist_transform.ListOfUuidTerraformToSdk(ctx, d.ExtraSiteIds)
	}

	if !d.GracePeriod.IsNull() && !d.GracePeriod.IsUnknown() {
		data.GracePeriod = models.ToPointer(int(d.GracePeriod.ValueInt64()))
	}

	return &data
}
