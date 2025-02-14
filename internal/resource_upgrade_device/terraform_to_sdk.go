package resource_upgrade_device

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *UpgradeDeviceModel) (*models.DeviceUpgrade, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.DeviceUpgrade{}

	if plan.Reboot.ValueBoolPointer() != nil {
		data.Reboot = plan.Reboot.ValueBoolPointer()
	}

	if plan.RebootAt.ValueInt64Pointer() != nil {
		data.RebootAt = models.ToPointer(int(plan.RebootAt.ValueInt64()))
	}

	if plan.Snapshot.ValueBoolPointer() != nil {
		data.Snapshot = plan.Snapshot.ValueBoolPointer()
	}

	if plan.StartTime.ValueInt64Pointer() != nil {
		data.RebootAt = models.ToPointer(int(plan.RebootAt.ValueInt64()))
	}

	if plan.TargetVersion.ValueStringPointer() != nil {
		data.Version = *plan.TargetVersion.ValueStringPointer()
	}

	return &data, diags
}
