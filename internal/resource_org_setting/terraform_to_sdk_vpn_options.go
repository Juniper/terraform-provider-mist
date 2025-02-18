package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vpnOptionsTerraformToSdk(d VpnOptionsValue) *models.OrgSettingVpnOptions {
	data := models.OrgSettingVpnOptions{}

	if d.AsBase.ValueInt64Pointer() != nil {
		data.AsBase = models.ToPointer(int(d.AsBase.ValueInt64()))
	}

	if d.StSubnet.ValueStringPointer() != nil {
		data.StSubnet = d.StSubnet.ValueStringPointer()
	}

	return &data
}
