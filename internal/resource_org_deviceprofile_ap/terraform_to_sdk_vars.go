package resource_org_deviceprofile_ap

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func varsTerraformToSdk(m basetypes.MapValue) map[string]string {
	dataMap := make(map[string]string)
	for k, v := range m.Elements() {
		dataMap[k] = v.String()
	}
	return dataMap
}
