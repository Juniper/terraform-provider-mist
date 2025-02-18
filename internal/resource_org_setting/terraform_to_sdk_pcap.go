package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func pcapTerraformToSdk(d PcapValue) *models.OrgSettingPcap {
	data := models.OrgSettingPcap{}

	if d.Bucket.ValueStringPointer() != nil {
		data.Bucket = d.Bucket.ValueStringPointer()
	}

	if d.MaxPktLen.ValueInt64Pointer() != nil {
		data.MaxPktLen = models.ToPointer(int(d.MaxPktLen.ValueInt64()))
	}

	return &data
}
