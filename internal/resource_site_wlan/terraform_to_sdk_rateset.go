package resource_site_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func bandRatesetTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RatesetValue) models.WlanDatarates {
	data := models.WlanDatarates{}
	if !d.IsNull() && !d.IsUnknown() {

		if d.Ht.ValueStringPointer() != nil {
			data.Ht = models.NewOptional(d.Ht.ValueStringPointer())
		}
		if !d.Legacy.IsNull() && !d.Legacy.IsUnknown() {
			var legacy []models.WlanDataratesLegacyItemEnum
			for _, item := range d.Legacy.Elements() {
				var s_interface interface{} = item
				s := s_interface.(basetypes.StringValue)
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

func ratesetTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, m basetypes.MapValue) map[string]models.WlanDatarates {
	data := make(map[string]models.WlanDatarates)

	for k, v := range m.Elements() {

		var v_interface interface{} = v
		d := v_interface.(RatesetValue)

		if !d.IsNull() && !d.IsUnknown() {
			data[k] = bandRatesetTerraformToSdk(ctx, diags, d)
		}
	}

	return data
}
