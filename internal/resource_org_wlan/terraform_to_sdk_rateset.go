package resource_org_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func band24RatesetTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.WlanDatarates {
	data := models.WlanDatarates{}
	if !d.IsNull() && !d.IsUnknown() {

		item, e := NewBand24Value(Band24Value{}.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if item.Ht.ValueStringPointer() != nil {
				data.Ht = models.NewOptional(item.Ht.ValueStringPointer())
			}
			if !item.Legacy.IsNull() && !item.Legacy.IsUnknown() {
				var legacy []models.WlanDataratesLegacyItemEnum
				for _, item := range item.Legacy.Elements() {
					var s_interface interface{} = item
					s := s_interface.(basetypes.StringValue)
					legacy = append(legacy, models.WlanDataratesLegacyItemEnum(s.ValueString()))
				}
				data.Legacy = legacy
			}
			if item.MinRssi.ValueInt64Pointer() != nil {
				data.MinRssi = models.ToPointer(int(item.MinRssi.ValueInt64()))
			}
			if item.Template.ValueStringPointer() != nil {
				data.Template = models.NewOptional((*models.WlanDataratesTemplateEnum)(item.Template.ValueStringPointer()))
			}
			if item.Vht.ValueStringPointer() != nil {
				data.Vht = models.NewOptional(item.Vht.ValueStringPointer())
			}
		}
	}
	return data
}
func band5RatesetTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.WlanDatarates {
	data := models.WlanDatarates{}
	if !d.IsNull() && !d.IsUnknown() {

		item, e := NewBand5Value(Band5Value{}.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if item.Ht.ValueStringPointer() != nil {
				data.Ht = models.NewOptional(item.Ht.ValueStringPointer())
			}
			if !item.Legacy.IsNull() && !item.Legacy.IsUnknown() {
				var legacy []models.WlanDataratesLegacyItemEnum
				for _, item := range item.Legacy.Elements() {
					var s_interface interface{} = item
					s := s_interface.(basetypes.StringValue)
					legacy = append(legacy, models.WlanDataratesLegacyItemEnum(s.ValueString()))
				}
				data.Legacy = legacy
			}
			if item.MinRssi.ValueInt64Pointer() != nil {
				data.MinRssi = models.ToPointer(int(item.MinRssi.ValueInt64()))
			}
			if item.Template.ValueStringPointer() != nil {
				data.Template = models.NewOptional((*models.WlanDataratesTemplateEnum)(item.Template.ValueStringPointer()))
			}
			if item.Vht.ValueStringPointer() != nil {
				data.Vht = models.NewOptional(item.Vht.ValueStringPointer())
			}
		}
	}
	return data
}

func ratesetTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RatesetValue) map[string]models.WlanDatarates {
	data := make(map[string]models.WlanDatarates)

	if !d.Band24.IsNull() && !d.Band24.IsUnknown() {
		data["24"] = band24RatesetTerraformToSdk(ctx, diags, d.Band24)
	}
	if !d.Band5.IsNull() && !d.Band5.IsUnknown() {
		data["5"] = band5RatesetTerraformToSdk(ctx, diags, d.Band5)
	}

	return data
}
