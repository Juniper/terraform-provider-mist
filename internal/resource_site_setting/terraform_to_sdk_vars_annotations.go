package resource_site_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func varsAnnotationsTerraformToSdk(d basetypes.MapValue) map[string]models.VarsAnnotation {
	dataMap := make(map[string]models.VarsAnnotation)
	for k, v := range d.Elements() {
		var vi interface{} = v
		plan := vi.(VarsAnnotationsValue)
		data := models.VarsAnnotation{}
		if plan.Note.ValueStringPointer() != nil {
			data.Note = plan.Note.ValueStringPointer()
		}
		if plan.VarsAnnotationsType.ValueStringPointer() != nil {
			data.Type = plan.VarsAnnotationsType.ValueStringPointer()
		}
		dataMap[k] = data
	}
	return dataMap
}
