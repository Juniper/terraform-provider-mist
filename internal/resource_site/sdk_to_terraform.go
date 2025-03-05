package resource_site

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(data *models.Site) (SiteModel, diag.Diagnostics) {
	var state SiteModel
	var diags diag.Diagnostics

	unset := make(map[string]interface{})

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Name = types.StringValue(data.Name)

	var address basetypes.StringValue
	var latlng LatlngValue
	var countryCode basetypes.StringValue
	var timezone basetypes.StringValue
	var notes basetypes.StringValue
	var alarmtemplateId basetypes.StringValue
	var aptemplateId basetypes.StringValue
	var gatewaytemplateId basetypes.StringValue
	var networktemplateId basetypes.StringValue
	var rftemplateId basetypes.StringValue
	var secpolicyId basetypes.StringValue
	var sitetemplateId basetypes.StringValue
	var sitegroupIds = types.ListNull(types.StringType)
	var tzOffset basetypes.Int64Value

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
		if e != nil {
			diags.Append(e...)
		}
		latlng = res
	}
	if data.CountryCode != nil {
		countryCode = types.StringValue(*data.CountryCode)
	}
	if data.Timezone != nil {
		timezone = types.StringValue(*data.Timezone)
	}
	if data.Notes != nil {
		notes = types.StringValue(*data.Notes)
	}

	if data.AlarmtemplateId.Value() != nil && data.AlarmtemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		alarmtemplateId = types.StringValue(data.AlarmtemplateId.Value().String())
	}
	if data.AptemplateId.Value() != nil && data.AptemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		aptemplateId = types.StringValue(data.AptemplateId.Value().String())
	}
	if data.GatewaytemplateId.Value() != nil && data.GatewaytemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		gatewaytemplateId = types.StringValue(data.GatewaytemplateId.Value().String())
	}
	if data.NetworktemplateId.Value() != nil && data.NetworktemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		networktemplateId = types.StringValue(data.NetworktemplateId.Value().String())
	}
	if data.RftemplateId.Value() != nil && data.RftemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		rftemplateId = types.StringValue(data.RftemplateId.Value().String())
	}
	if data.SecpolicyId.Value() != nil && data.SecpolicyId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		secpolicyId = types.StringValue(data.SecpolicyId.Value().String())
	}
	if data.SitetemplateId.Value() != nil && data.SitetemplateId.Value().String() != "00000000-0000-0000-0000-000000000000" {
		sitetemplateId = types.StringValue(data.SitetemplateId.Value().String())
	}
	if data.SitegroupIds != nil {
		var items []attr.Value
		var itemsType attr.Type = basetypes.StringType{}
		for _, item := range data.SitegroupIds {
			items = append(items, types.StringValue(item.String()))
		}
		list, e := types.ListValue(itemsType, items)
		if e != nil {
			diags.Append(e...)
		} else {
			sitegroupIds = list
		}
	}
	if data.Tzoffset != nil {
		tzOffset = types.Int64Value(int64(*data.Tzoffset))
	} else {
		unset["-tzoffset"] = ""
	}

	state.Address = address
	state.Latlng = latlng
	state.CountryCode = countryCode
	state.Timezone = timezone
	state.Notes = notes
	state.AlarmtemplateId = alarmtemplateId
	state.AptemplateId = aptemplateId
	state.GatewaytemplateId = gatewaytemplateId
	state.NetworktemplateId = networktemplateId
	state.RftemplateId = rftemplateId
	state.SecpolicyId = secpolicyId
	state.SitetemplateId = sitetemplateId
	state.SitegroupIds = sitegroupIds
	state.Tzoffset = tzOffset

	return state, diags
}
