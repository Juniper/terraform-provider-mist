package datasource_sites

import (
	"context"
	"math/big"

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
	var alarmtemplate_id basetypes.StringValue
	var aptemplate_id basetypes.StringValue
	var country_code basetypes.StringValue
	var created_time basetypes.NumberValue
	var gatewaytemplate_id basetypes.StringValue
	var id basetypes.StringValue
	var latlng basetypes.ObjectValue = types.ObjectNull(LatlngValue{}.AttributeTypes(ctx))
	var modified_time basetypes.NumberValue
	var name basetypes.StringValue
	var networktemplate_id basetypes.StringValue
	var notes basetypes.StringValue
	var org_id basetypes.StringValue
	var rftemplate_id basetypes.StringValue
	var secpolicy_id basetypes.StringValue
	var sitegroup_ids basetypes.ListValue = types.ListNull(types.StringType)
	var sitetemplate_id basetypes.StringValue
	var timezone basetypes.StringValue

	if d.Address != nil {
		address = types.StringValue(*d.Address)
	}
	if d.AlarmtemplateId.Value() != nil {
		alarmtemplate_id = types.StringValue(d.AlarmtemplateId.Value().String())
	}
	if d.AptemplateId.Value() != nil {
		aptemplate_id = types.StringValue(d.AptemplateId.Value().String())
	}
	if d.CountryCode != nil {
		country_code = types.StringValue(*d.CountryCode)
	}
	if d.CreatedTime != nil {
		created_time = types.NumberValue(big.NewFloat(*d.CreatedTime))
	}
	if d.GatewaytemplateId.Value() != nil {
		gatewaytemplate_id = types.StringValue(d.GatewaytemplateId.Value().String())
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
		modified_time = types.NumberValue(big.NewFloat(*d.ModifiedTime))
	}
	name = types.StringValue(d.Name)

	if d.NetworktemplateId.Value() != nil {
		networktemplate_id = types.StringValue(d.NetworktemplateId.Value().String())
	}
	if d.Notes != nil {
		notes = types.StringValue(*d.Notes)
	}
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.RftemplateId.Value() != nil {
		rftemplate_id = types.StringValue(d.RftemplateId.Value().String())
	}
	if d.SecpolicyId.Value() != nil {
		secpolicy_id = types.StringValue(d.SecpolicyId.Value().String())
	}
	if d.SitegroupIds != nil {
		var items []attr.Value
		var items_type attr.Type = basetypes.StringType{}
		for _, item := range d.SitegroupIds {
			items = append(items, types.StringValue(item.String()))
		}
		tmp_sitegroup_ids, e := types.ListValue(items_type, items)
		diags.Append(e...)
		sitegroup_ids = tmp_sitegroup_ids
	}
	if d.SitetemplateId.Value() != nil {
		sitetemplate_id = types.StringValue(d.SitetemplateId.Value().String())
	}
	if d.Timezone != nil {
		timezone = types.StringValue(*d.Timezone)
	}

	data_map_attr_type := SitesValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"address":            address,
		"alarmtemplate_id":   alarmtemplate_id,
		"aptemplate_id":      aptemplate_id,
		"country_code":       country_code,
		"created_time":       created_time,
		"gatewaytemplate_id": gatewaytemplate_id,
		"id":                 id,
		"latlng":             latlng,
		"modified_time":      modified_time,
		"name":               name,
		"networktemplate_id": networktemplate_id,
		"notes":              notes,
		"org_id":             org_id,
		"rftemplate_id":      rftemplate_id,
		"secpolicy_id":       secpolicy_id,
		"sitegroup_ids":      sitegroup_ids,
		"sitetemplate_id":    sitetemplate_id,
		"timezone":           timezone,
	}
	data, e := NewSitesValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
