package resource_site_setting

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vsInstanceTerraformToSdk(m basetypes.MapValue) map[string]models.VsInstanceProperty {
	dataMap := make(map[string]models.VsInstanceProperty)
	for k, v := range m.Elements() {

		var vi interface{} = v
		vd := vi.(VsInstanceValue)
		data := models.VsInstanceProperty{}
		data.Networks = misttransform.ListOfStringTerraformToSdk(vd.Networks)
		dataMap[k] = data
	}
	return dataMap
}
