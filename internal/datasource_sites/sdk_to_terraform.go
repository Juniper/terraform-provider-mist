package datasource_sites

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.Site, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := siteSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func siteSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Site) SitesValue {

	var address basetypes.StringValue
	var alarmtemplateId basetypes.StringValue
	var aptemplateId basetypes.StringValue
	var countryCode basetypes.StringValue
	var createdTime basetypes.Float64Value
	var gatewaytemplateId basetypes.StringValue
	var id basetypes.StringValue
	var latlng = types.ObjectNull(LatlngValue{}.AttributeTypes(ctx))
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var networktemplateId basetypes.StringValue
	var notes basetypes.StringValue
	var orgId basetypes.StringValue
	var rftemplateId basetypes.StringValue
	var secpolicyId basetypes.StringValue
	var sitegroupIds = types.ListNull(types.StringType)
	var sitetemplateId basetypes.StringValue
	var timezone basetypes.StringValue

	if d.Address != nil {
		address = types.StringValue(*d.Address)
	}
	if d.AlarmtemplateId.Value() != nil {
		alarmtemplateId = types.StringValue(d.AlarmtemplateId.Value().String())
	}
	if d.AptemplateId.Value() != nil {
		aptemplateId = types.StringValue(d.AptemplateId.Value().String())
	}
	if d.CountryCode != nil {
		countryCode = types.StringValue(*d.CountryCode)
	}
	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.GatewaytemplateId.Value() != nil {
		gatewaytemplateId = types.StringValue(d.GatewaytemplateId.Value().String())
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Latlng != nil {

		t := map[string]attr.Type{
			"lat": basetypes.Float64Type{},
			"lng": basetypes.Float64Type{},
		}
		v := map[string]attr.Value{
			"lat": types.Float64Value(d.Latlng.Lat),
			"lng": types.Float64Value(d.Latlng.Lng),
		}
		res, e := NewLatlngValue(t, v)
		diags.Append(e...)
		latlng, e = res.ToObjectValue(ctx)
		diags.Append(e...)
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	name = types.StringValue(d.Name)

	if d.NetworktemplateId.Value() != nil {
		networktemplateId = types.StringValue(d.NetworktemplateId.Value().String())
	}
	if d.Notes != nil {
		notes = types.StringValue(*d.Notes)
	}
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.RftemplateId.Value() != nil {
		rftemplateId = types.StringValue(d.RftemplateId.Value().String())
	}
	if d.SecpolicyId.Value() != nil {
		secpolicyId = types.StringValue(d.SecpolicyId.Value().String())
	}
	if d.SitegroupIds != nil {
		var items []attr.Value
		var itemsType attr.Type = basetypes.StringType{}
		for _, item := range d.SitegroupIds {
			items = append(items, types.StringValue(item.String()))
		}
		tmpSitegroupIds, e := types.ListValue(itemsType, items)
		diags.Append(e...)
		sitegroupIds = tmpSitegroupIds
	}
	if d.SitetemplateId.Value() != nil {
		sitetemplateId = types.StringValue(d.SitetemplateId.Value().String())
	}
	if d.Timezone != nil {
		timezone = types.StringValue(*d.Timezone)
	}

	dataMapValue := map[string]attr.Value{
		"address":            address,
		"alarmtemplate_id":   alarmtemplateId,
		"aptemplate_id":      aptemplateId,
		"country_code":       countryCode,
		"created_time":       createdTime,
		"gatewaytemplate_id": gatewaytemplateId,
		"id":                 id,
		"latlng":             latlng,
		"modified_time":      modifiedTime,
		"name":               name,
		"networktemplate_id": networktemplateId,
		"notes":              notes,
		"org_id":             orgId,
		"rftemplate_id":      rftemplateId,
		"secpolicy_id":       secpolicyId,
		"sitegroup_ids":      sitegroupIds,
		"sitetemplate_id":    sitetemplateId,
		"timezone":           timezone,
	}
	data, e := NewSitesValue(SitesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
