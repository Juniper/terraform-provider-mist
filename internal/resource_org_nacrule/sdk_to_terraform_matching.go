package resource_org_nacrule

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func matchingPortTypesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.NacRuleMatchingPortTypeEnum) basetypes.ListValue {
	listAttrTypes := types.StringType
	var listAttrValues []attr.Value
	for _, v := range d {
		vString := types.StringValue(string(v))
		listAttrValues = append(listAttrValues, vString)
	}

	r, e := types.ListValueFrom(ctx, listAttrTypes, listAttrValues)
	diags.Append(e...)
	return r
}

func matchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NacRuleMatching) MatchingValue {

	var authType basetypes.StringValue
	var family = types.ListNull(types.StringType)
	var mfg = types.ListNull(types.StringType)
	var model = types.ListNull(types.StringType)
	var nactags = types.ListNull(types.StringType)
	var osType = types.ListNull(types.StringType)
	var portTypes = types.ListNull(types.StringType)
	var siteIds = types.ListNull(types.StringType)
	var sitegroupIds = types.ListNull(types.StringType)
	var vendor = types.ListNull(types.StringType)

	if d.AuthType != nil {
		authType = types.StringValue(string(*d.AuthType))
	}
	if d.Family != nil {
		family = mistutils.ListOfStringSdkToTerraform(d.Family)
	}
	if d.Mfg != nil {
		mfg = mistutils.ListOfStringSdkToTerraform(d.Mfg)
	}
	if d.Model != nil {
		model = mistutils.ListOfStringSdkToTerraform(d.Model)
	}
	if d.Nactags != nil {
		nactags = mistutils.ListOfStringSdkToTerraform(d.Nactags)
	}
	if d.OsType != nil {
		osType = mistutils.ListOfStringSdkToTerraform(d.OsType)
	}
	if d.PortTypes != nil {
		portTypes = matchingPortTypesSdkToTerraform(ctx, diags, d.PortTypes)
	}
	if d.SiteIds != nil {
		siteIds = mistutils.ListOfUuidSdkToTerraform(d.SiteIds)
	}
	if d.SitegroupIds != nil {
		sitegroupIds = mistutils.ListOfUuidSdkToTerraform(d.SitegroupIds)
	}
	if d.Vendor != nil {
		vendor = mistutils.ListOfStringSdkToTerraform(d.Vendor)
	}

	dataMapValue := map[string]attr.Value{
		"auth_type":     authType,
		"family":        family,
		"mfg":           mfg,
		"model":         model,
		"nactags":       nactags,
		"os_type":       osType,
		"port_types":    portTypes,
		"site_ids":      siteIds,
		"sitegroup_ids": sitegroupIds,
		"vendor":        vendor,
	}
	data, e := NewMatchingValue(MatchingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func notMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NacRuleMatching) NotMatchingValue {

	var authType basetypes.StringValue
	var family = types.ListNull(types.StringType)
	var mfg = types.ListNull(types.StringType)
	var model = types.ListNull(types.StringType)
	var nactags = types.ListNull(types.StringType)
	var osType = types.ListNull(types.StringType)
	var portTypes = types.ListNull(types.StringType)
	var siteIds = types.ListNull(types.StringType)
	var sitegroupIds = types.ListNull(types.StringType)
	var vendor = types.ListNull(types.StringType)

	if d.AuthType != nil {
		authType = types.StringValue(string(*d.AuthType))
	}
	if d.Family != nil {
		family = mistutils.ListOfStringSdkToTerraform(d.Family)
	}
	if d.Mfg != nil {
		mfg = mistutils.ListOfStringSdkToTerraform(d.Mfg)
	}
	if d.Model != nil {
		model = mistutils.ListOfStringSdkToTerraform(d.Model)
	}
	if d.Nactags != nil {
		nactags = mistutils.ListOfStringSdkToTerraform(d.Nactags)
	}
	if d.OsType != nil {
		osType = mistutils.ListOfStringSdkToTerraform(d.OsType)
	}
	if d.PortTypes != nil {
		portTypes = matchingPortTypesSdkToTerraform(ctx, diags, d.PortTypes)
	}
	if d.SiteIds != nil {
		siteIds = mistutils.ListOfUuidSdkToTerraform(d.SiteIds)
	}
	if d.SitegroupIds != nil {
		sitegroupIds = mistutils.ListOfUuidSdkToTerraform(d.SitegroupIds)
	}
	if d.Vendor != nil {
		vendor = mistutils.ListOfStringSdkToTerraform(d.Vendor)
	}

	dataMapValue := map[string]attr.Value{
		"auth_type":     authType,
		"family":        family,
		"mfg":           mfg,
		"model":         model,
		"nactags":       nactags,
		"os_type":       osType,
		"port_types":    portTypes,
		"site_ids":      siteIds,
		"sitegroup_ids": sitegroupIds,
		"vendor":        vendor,
	}
	data, e := NewNotMatchingValue(NotMatchingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
