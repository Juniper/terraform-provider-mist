package resource_org_nacrule

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func matchingPortTypesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.NacRuleMatchingPortTypeEnum) basetypes.ListValue {
	list_attr_types := types.StringType
	var list_attr_values []attr.Value
	for _, v := range d {
		v_string := types.StringValue(string(v))
		list_attr_values = append(list_attr_values, v_string)
	}

	r, e := types.ListValueFrom(ctx, list_attr_types, list_attr_values)
	diags.Append(e...)
	return r
}

func matchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NacRuleMatching) MatchingValue {

	var auth_type basetypes.StringValue
	var nactags basetypes.ListValue = types.ListNull(types.StringType)
	var port_types basetypes.ListValue = types.ListNull(types.StringType)
	var site_ids basetypes.ListValue = types.ListNull(types.StringType)
	var sitegroup_ids basetypes.ListValue = types.ListNull(types.StringType)
	var vendor basetypes.ListValue = types.ListNull(types.StringType)

	if d.AuthType != nil {
		auth_type = types.StringValue(string(*d.AuthType))
	}
	if d.Nactags != nil {
		nactags = mist_transform.ListOfStringSdkToTerraform(ctx, d.Nactags)
	}
	if d.PortTypes != nil {
		port_types = matchingPortTypesSdkToTerraform(ctx, diags, d.PortTypes)
	}
	if d.SiteIds != nil {
		site_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, d.SiteIds)
	}
	if d.SitegroupIds != nil {
		sitegroup_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, d.SitegroupIds)
	}
	if d.Vendor != nil {
		vendor = mist_transform.ListOfStringSdkToTerraform(ctx, d.Vendor)
	}

	data_map_attr_type := MatchingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"auth_type":     auth_type,
		"nactags":       nactags,
		"port_types":    port_types,
		"site_ids":      site_ids,
		"sitegroup_ids": sitegroup_ids,
		"vendor":        vendor,
	}
	data, e := NewMatchingValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func notMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NacRuleMatching) NotMatchingValue {

	var auth_type basetypes.StringValue
	var nactags basetypes.ListValue = types.ListNull(types.StringType)
	var port_types basetypes.ListValue = types.ListNull(types.StringType)
	var site_ids basetypes.ListValue = types.ListNull(types.StringType)
	var sitegroup_ids basetypes.ListValue = types.ListNull(types.StringType)
	var vendor basetypes.ListValue = types.ListNull(types.StringType)

	if d.AuthType != nil {
		auth_type = types.StringValue(string(*d.AuthType))
	}
	if d.Nactags != nil {
		nactags = mist_transform.ListOfStringSdkToTerraform(ctx, d.Nactags)
	}
	if d.PortTypes != nil {
		port_types = matchingPortTypesSdkToTerraform(ctx, diags, d.PortTypes)
	}
	if d.SiteIds != nil {
		site_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, d.SiteIds)
	}
	if d.SitegroupIds != nil {
		sitegroup_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, d.SitegroupIds)
	}
	if d.Vendor != nil {
		vendor = mist_transform.ListOfStringSdkToTerraform(ctx, d.Vendor)
	}

	data_map_attr_type := NotMatchingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"auth_type":     auth_type,
		"nactags":       nactags,
		"port_types":    port_types,
		"site_ids":      site_ids,
		"sitegroup_ids": sitegroup_ids,
		"vendor":        vendor,
	}
	data, e := NewNotMatchingValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
