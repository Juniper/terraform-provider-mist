package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vpnOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VpnOptionsValue) *models.OrgSettingVpnOptions {
	data := models.OrgSettingVpnOptions{}

	if d.AsBase.ValueInt64Pointer() != nil {
		data.AsBase = models.ToPointer(int(d.AsBase.ValueInt64()))
	}

	if d.StSubnet.ValueStringPointer() != nil {
		data.StSubnet = d.StSubnet.ValueStringPointer()
	}

	return &data
}
