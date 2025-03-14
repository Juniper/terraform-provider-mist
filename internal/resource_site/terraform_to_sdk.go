package resource_site

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
)

func TerraformToSdk(plan *SiteModel) (*models.Site, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data := models.Site{}
	data.Name = plan.Name.ValueString()

	if (!plan.Address.IsNull()) && !plan.Address.IsUnknown() {
		data.Address = models.ToPointer(plan.Address.ValueString())
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

	if (!plan.Notes.IsNull()) && !plan.Notes.IsUnknown() {
		data.Notes = plan.Notes.ValueStringPointer()
	} else {
		unset["-notes"] = ""
	}

	if !plan.AlarmtemplateId.IsNull() && !plan.AlarmtemplateId.IsUnknown() {
		alarmtemplateId, e := uuid.Parse(plan.AlarmtemplateId.ValueString())
		if e == nil {
			data.AlarmtemplateId = models.NewOptional(&alarmtemplateId)
		} else {
			diags.AddError("Bad value for alarmtemplate_id", e.Error())
		}
	} else {
		unset["-alarmtemplate_id"] = ""
	}

	if !plan.AptemplateId.IsNull() && !plan.AptemplateId.IsUnknown() {
		aptemplateId, e := uuid.Parse(plan.AptemplateId.ValueString())
		if e == nil {
			data.AptemplateId = models.NewOptional(&aptemplateId)
		} else {
			diags.AddError("Bad value for aptemplate_id", e.Error())
		}
	} else {
		unset["-aptemplate_id"] = ""
	}

	if !plan.GatewaytemplateId.IsNull() && !plan.GatewaytemplateId.IsUnknown() {
		gatewaytemplateId, e := uuid.Parse(plan.GatewaytemplateId.ValueString())
		if e == nil {
			data.GatewaytemplateId = models.NewOptional(&gatewaytemplateId)
		} else {
			diags.AddError("Bad value for gatewaytemplate_id", e.Error())
		}
	} else {
		unset["-gatewaytemplate_id"] = ""
	}

	if !plan.NetworktemplateId.IsNull() && !plan.NetworktemplateId.IsUnknown() {
		networktemplateId, e := uuid.Parse(plan.NetworktemplateId.ValueString())
		if e == nil {
			data.NetworktemplateId = models.NewOptional(&networktemplateId)
		} else {
			diags.AddError("Bad value for networktemplate_id", e.Error())
		}
	} else {
		unset["-networktemplate_id"] = ""
	}

	if !plan.RftemplateId.IsNull() && !plan.RftemplateId.IsUnknown() {
		rftemplateId, e := uuid.Parse(plan.RftemplateId.ValueString())
		if e == nil {
			data.RftemplateId = models.NewOptional(&rftemplateId)
		} else {
			diags.AddError("Bad value for rftemplate_id", e.Error())
		}
	} else {
		unset["-rftemplate_id"] = nil
	}

	if !plan.SecpolicyId.IsNull() && !plan.SecpolicyId.IsUnknown() {
		secpolicyId, e := uuid.Parse(plan.SecpolicyId.ValueString())
		if e == nil {
			data.SecpolicyId = models.NewOptional(&secpolicyId)
		} else {
			diags.AddError("Bad value for secpolicy_id", e.Error())
		}
	} else {
		unset["-secpolicy_id"] = ""
	}

	if !plan.SitetemplateId.IsNull() && !plan.SitetemplateId.IsUnknown() {
		sitetemplateId, e := uuid.Parse(plan.SitetemplateId.ValueString())
		if e == nil {
			data.SitetemplateId = models.NewOptional(&sitetemplateId)
		} else {
			diags.AddError("Bad value for sitetemplate_id", e.Error())
		}
	} else {
		unset["-sitetemplate_id"] = ""
	}

	if !plan.SitegroupIds.IsNull() && !plan.SitegroupIds.IsUnknown() {
		var items []uuid.UUID
		for _, item := range plan.SitegroupIds.Elements() {
			var iface interface{} = item
			val := iface.(basetypes.StringValue)
			items = append(items, uuid.MustParse(val.ValueString()))
		}
		data.SitegroupIds = items
	} else {
		unset["-sitegroup_ids"] = ""
	}

	data.AdditionalProperties = unset
	return &data, diags
}
