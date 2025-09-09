package resource_site

import (
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(plan *SiteModel) (*models.Site, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data := models.Site{}
	data.Name = plan.Name.ValueString()

	if plan.Address.ValueStringPointer() != nil {
		data.Address = models.NewOptional(plan.Address.ValueStringPointer())
	} else {
		unset["-address"] = ""
	}

	if (!plan.Latlng.IsNull()) && !plan.Latlng.IsUnknown() {
		var dataLatlng models.LatLng
		dataLatlng.Lat = plan.Latlng.Lat.ValueFloat64()
		dataLatlng.Lng = plan.Latlng.Lng.ValueFloat64()
		data.Latlng = models.ToPointer(dataLatlng)
	} else {
		unset["-latlng"] = ""
	}

	if (!plan.CountryCode.IsNull()) && !plan.CountryCode.IsUnknown() {
		data.CountryCode = plan.CountryCode.ValueStringPointer()
	} else {
		unset["-country_code"] = ""
	}

	if (!plan.Timezone.IsNull()) && !plan.Timezone.IsUnknown() {
		data.Timezone = plan.Timezone.ValueStringPointer()
	} else {
		unset["-timezone"] = ""
	}

	if plan.Notes.ValueStringPointer() != nil {
		data.Notes = models.NewOptional(plan.Notes.ValueStringPointer())
	} else {
		unset["-notes"] = ""
	}

	data.AlarmtemplateId = mistutils.RequiredMistIdField(&diags, plan.AlarmtemplateId)
	data.AptemplateId = mistutils.RequiredMistIdField(&diags, plan.AptemplateId)
	data.GatewaytemplateId = mistutils.RequiredMistIdField(&diags, plan.GatewaytemplateId)
	data.NetworktemplateId = mistutils.RequiredMistIdField(&diags, plan.NetworktemplateId)
	data.RftemplateId = mistutils.RequiredMistIdField(&diags, plan.RftemplateId)
	data.SecpolicyId = mistutils.RequiredMistIdField(&diags, plan.SecpolicyId)
	data.SitetemplateId = mistutils.RequiredMistIdField(&diags, plan.SitetemplateId)

	var items []uuid.UUID
	for _, item := range plan.SitegroupIds.Elements() {
		var iface interface{} = item
		val := iface.(basetypes.StringValue)
		items = append(items, uuid.MustParse(val.ValueString()))
	}
	data.SitegroupIds = items

	data.AdditionalProperties = unset
	return &data, diags
}
