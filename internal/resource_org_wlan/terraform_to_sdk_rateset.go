package resource_org_wlan

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func bandRatesetTerraformToSdk(d RatesetValue) models.WlanDatarates {
	data := models.WlanDatarates{}
	if !d.IsNull() && !d.IsUnknown() {

		if d.Eht.ValueStringPointer() != nil {
			data.Eht = models.NewOptional(d.Eht.ValueStringPointer())
		}
		if d.He.ValueStringPointer() != nil {
			data.He = models.NewOptional(d.He.ValueStringPointer())
		}
		if d.Ht.ValueStringPointer() != nil {
			data.Ht = models.NewOptional(d.Ht.ValueStringPointer())
		}
		if !d.Legacy.IsNull() && !d.Legacy.IsUnknown() {
			var legacy []models.WlanDataratesLegacyItemEnum
			for _, item := range d.Legacy.Elements() {
				var sInterface interface{} = item
				s := sInterface.(basetypes.StringValue)
				legacy = append(legacy, models.WlanDataratesLegacyItemEnum(s.ValueString()))
			}
			data.Legacy = legacy
		}
		if d.MinRssi.ValueInt64Pointer() != nil {
			data.MinRssi = models.ToPointer(int(d.MinRssi.ValueInt64()))
		}
		if d.Template.ValueStringPointer() != nil {
			data.Template = models.NewOptional((*models.WlanDataratesTemplateEnum)(d.Template.ValueStringPointer()))
		}
		if d.Vht.ValueStringPointer() != nil {
			data.Vht = models.NewOptional(d.Vht.ValueStringPointer())
		}
	}
	return data
}

func ratesetTerraformToSdk(m basetypes.MapValue) map[string]models.WlanDatarates {
	data := make(map[string]models.WlanDatarates)

	for k, v := range m.Elements() {

		var vInterface interface{} = v
		d := vInterface.(RatesetValue)

		if !d.IsNull() && !d.IsUnknown() {
			data[k] = bandRatesetTerraformToSdk(d)
		}
	}

	return data
}
