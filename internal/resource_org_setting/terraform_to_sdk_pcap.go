package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func pcapTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d PcapValue) *models.OrgSettingPcap {
	data := models.OrgSettingPcap{}

	if d.Bucket.ValueStringPointer() != nil {
		data.Bucket = d.Bucket.ValueStringPointer()
	}

	if d.MaxPktLen.ValueInt64Pointer() != nil {
		data.MaxPktLen = models.ToPointer(int(d.MaxPktLen.ValueInt64()))
	}

	return &data
}
