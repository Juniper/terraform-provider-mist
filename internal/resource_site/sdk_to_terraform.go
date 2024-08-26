package resource_site

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data *models.Site) (SiteModel, diag.Diagnostics) {
	var state SiteModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Name = types.StringValue(data.Name)

	var address basetypes.StringValue
	var latlng LatlngValue
	var country_code basetypes.StringValue
	var timezone basetypes.StringValue
	var notes basetypes.StringValue
	var alarmtemplate_id basetypes.StringValue
	var aptemplate_id basetypes.StringValue
	var gatewaytemplate_id basetypes.StringValue
	var networktemplate_id basetypes.StringValue
	var rftemplate_id basetypes.StringValue
	var secpolicy_id basetypes.StringValue
	var sitetemplate_id basetypes.StringValue
	var sitegroup_ids basetypes.ListValue = types.ListNull(types.StringType)

	if data.Address != nil {
		address = types.StringValue(*data.Address)
	}
	if data.Latlng != nil {
		t := map[string]attr.Type{
			"lat": basetypes.Float64Type{},
			"lng": basetypes.Float64Type{},
		}
		v := map[string]attr.Value{
			"lat": types.Float64Value(data.Latlng.Lat),
			"lng": types.Float64Value(data.Latlng.Lng),
		}
		res, e := NewLatlngValue(t, v)
		diags.Append(e...)
		latlng = res
	}
	if data.CountryCode != nil {
		country_code = types.StringValue(*data.CountryCode)
	}
	if data.Timezone != nil {
		timezone = types.StringValue(*data.Timezone)
	}
	if data.Notes != nil {
		notes = types.StringValue(*data.Notes)
	}

	if data.AlarmtemplateId.Value() != nil && data.AlarmtemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		alarmtemplate_id = types.StringValue(data.AlarmtemplateId.Value().String())
	}
	if data.AptemplateId.Value() != nil && data.AptemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		aptemplate_id = types.StringValue(data.AptemplateId.Value().String())
	}
	if data.GatewaytemplateId.Value() != nil && data.GatewaytemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		gatewaytemplate_id = types.StringValue(data.GatewaytemplateId.Value().String())
	}
	if data.NetworktemplateId.Value() != nil && data.NetworktemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		networktemplate_id = types.StringValue(data.NetworktemplateId.Value().String())
	}
	if data.RftemplateId.Value() != nil && data.RftemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		rftemplate_id = types.StringValue(data.RftemplateId.Value().String())
	}
	if data.SecpolicyId.Value() != nil && data.SecpolicyId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		secpolicy_id = types.StringValue(data.SecpolicyId.Value().String())
	}
	if data.SitetemplateId.Value() != nil && data.SitetemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		sitetemplate_id = types.StringValue(data.SitetemplateId.Value().String())
	}
	if data.SitegroupIds != nil {
		var items []attr.Value
		var items_type attr.Type = basetypes.StringType{}
		for _, item := range data.SitegroupIds {
			items = append(items, types.StringValue(item.String()))
		}
		list, e := types.ListValue(items_type, items)
		if e != nil {
			diags.Append(e...)
		} else {
			sitegroup_ids = list
		}
	}

	state.Address = address
	state.Latlng = latlng
	state.CountryCode = country_code
	state.Timezone = timezone
	state.Notes = notes
	state.AlarmtemplateId = alarmtemplate_id
	state.AptemplateId = aptemplate_id
	state.GatewaytemplateId = gatewaytemplate_id
	state.NetworktemplateId = networktemplate_id
	state.RftemplateId = rftemplate_id
	state.SecpolicyId = secpolicy_id
	state.SitetemplateId = sitetemplate_id
	state.SitegroupIds = sitegroup_ids

	return state, diags
}
