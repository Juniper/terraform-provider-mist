package resource_site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
)

func TerraformToSdk(ctx context.Context, plan *SiteModel) (*models.Site, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data := models.Site{}
	data.Name = plan.Name.ValueString()

	data.Address = models.ToPointer(plan.Address.ValueString())

	var data_latlng models.LatLng
	data_latlng.Lat = plan.Latlng.Lat.ValueFloat64()
	data_latlng.Lng = plan.Latlng.Lng.ValueFloat64()
	data.Latlng = models.ToPointer(data_latlng)

	data.CountryCode = plan.CountryCode.ValueStringPointer()

	data.Timezone = plan.Timezone.ValueStringPointer()

	data.Notes = plan.Notes.ValueStringPointer()

	if len(plan.AlarmtemplateId.ValueString()) > 0 {
		alarmtemplate_id, e := uuid.Parse(plan.AlarmtemplateId.ValueString())
		if e == nil {
			data.AlarmtemplateId = models.NewOptional(&alarmtemplate_id)
		} else {
			diags.AddError("Bad value for alarmtemplate_id", e.Error())
		}
	} else {
		unset["-alarmtemplate_id"] = ""
	}

	if len(plan.AptemplateId.ValueString()) > 0 {
		aptemplate_id, e := uuid.Parse(plan.AptemplateId.ValueString())
		if e == nil {
			data.AptemplateId = models.NewOptional(&aptemplate_id)
		} else {
			diags.AddError("Bad value for aptemplate_id", e.Error())
		}
	} else {
		unset["-aptemplate_id"] = ""
	}

	if len(plan.GatewaytemplateId.ValueString()) > 0 {
		gatewaytemplate_id, e := uuid.Parse(plan.GatewaytemplateId.ValueString())
		if e == nil {
			data.GatewaytemplateId = models.NewOptional(&gatewaytemplate_id)
		} else {
			diags.AddError("Bad value for gatewaytemplate_id", e.Error())
		}
	} else {
		unset["-gatewaytemplate_id"] = ""
	}

	if len(plan.NetworktemplateId.ValueString()) > 0 {
		networktemplate_id, e := uuid.Parse(plan.NetworktemplateId.ValueString())
		if e == nil {
			data.NetworktemplateId = models.NewOptional(&networktemplate_id)
		} else {
			diags.AddError("Bad value for networktemplate_id", e.Error())
		}
	} else {
		unset["-networktemplate_id"] = ""
	}

	if len(plan.RftemplateId.ValueString()) > 0 {
		rftemplate_id, e := uuid.Parse(plan.RftemplateId.ValueString())
		if e == nil {
			data.RftemplateId = models.NewOptional(&rftemplate_id)
		} else {
			diags.AddError("Bad value for rftemplate_id", e.Error())
		}
	} else {
		unset["rftemplate_id"] = nil
	}

	if len(plan.SecpolicyId.ValueString()) > 0 {
		secpolicy_id, e := uuid.Parse(plan.SecpolicyId.ValueString())
		if e == nil {
			data.SecpolicyId = models.NewOptional(&secpolicy_id)
		} else {
			diags.AddError("Bad value for secpolicy_id", e.Error())
		}
	} else {
		unset["-secpolicy_id"] = ""
	}

	if len(plan.SitetemplateId.ValueString()) > 0 {
		sitetemplate_id, e := uuid.Parse(plan.SitetemplateId.ValueString())
		if e == nil {
			data.SitetemplateId = models.NewOptional(&sitetemplate_id)
		} else {
			diags.AddError("Bad value for sitetemplate_id", e.Error())
		}
	} else {
		unset["-sitetemplate_id"] = ""
	}

	var items []uuid.UUID
	for _, item := range plan.SitegroupIds.Elements() {
		items = append(items, uuid.MustParse(item.String()))
	}
	data.SitegroupIds = items

	data.AdditionalProperties = unset
	return &data, diags
}
