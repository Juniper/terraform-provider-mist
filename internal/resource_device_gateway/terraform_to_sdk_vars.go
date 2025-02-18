package resource_device_gateway

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func varsTerraformToSdk(m basetypes.MapValue) map[string]string {
	dataMap := make(map[string]string)
	for k, v := range m.Elements() {
		var vi interface{} = v
		vd := vi.(basetypes.StringValue)
		dataMap[k] = vd.ValueString()
	}
	return dataMap
}
