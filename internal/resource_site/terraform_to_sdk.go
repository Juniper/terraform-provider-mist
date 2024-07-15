package resource_site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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

	alarmtemplate_id, e := uuid.Parse(plan.AlarmtemplateId.ValueString())
	if e == nil {
		data.AlarmtemplateId = models.NewOptional(&alarmtemplate_id)
	} else {
		unset["alarmtemplate_id"] = nil
	}

	aptemplate_id, e := uuid.Parse(plan.AptemplateId.ValueString())
	if e == nil {
		data.AptemplateId = models.NewOptional(&aptemplate_id)
	} else {
		unset["-aptemplate_id"] = ""
	}

	gatewaytemplate_id, e := uuid.Parse(plan.GatewaytemplateId.ValueString())
	if e == nil {
		data.GatewaytemplateId = models.NewOptional(&gatewaytemplate_id)
	} else {
		unset["gatewaytemplate_id"] = nil
	}

	networktemplate_id, e := uuid.Parse(plan.NetworktemplateId.ValueString())
	if e == nil {
		data.NetworktemplateId = models.NewOptional(&networktemplate_id)
	} else {
		unset["networktemplate_id"] = nil
	}

	rftemplate_id, e := uuid.Parse(plan.RftemplateId.ValueString())
	if e == nil {
		data.RftemplateId = models.NewOptional(&rftemplate_id)
	} else {
		unset["rftemplate_id"] = nil
	}

	secpolicy_id, e := uuid.Parse(plan.SecpolicyId.ValueString())
	if e == nil {
		data.SecpolicyId = models.NewOptional(&secpolicy_id)
	} else {
		unset["secpolicy_id"] = nil
	}

	sitetemplate_id, e := uuid.Parse(plan.SitetemplateId.ValueString())
	if e == nil {
		data.SitetemplateId = models.NewOptional(&sitetemplate_id)
	} else {
		tflog.Error(ctx, e.Error())
		unset["sitetemplate_id"] = nil
	}

	var items []uuid.UUID
	for _, item := range plan.SitegroupIds.Elements() {
		items = append(items, uuid.MustParse(item.String()))
	}
	data.SitegroupIds = items

	data.AdditionalProperties = unset
	return &data, diags
}
