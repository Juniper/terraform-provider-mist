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
	if len(d.Family) > 0 {
		family = mistutils.ListOfStringSdkToTerraform(d.Family)
	}
	if len(d.Mfg) > 0 {
		mfg = mistutils.ListOfStringSdkToTerraform(d.Mfg)
	}
	if len(d.Model) > 0 {
		model = mistutils.ListOfStringSdkToTerraform(d.Model)
	}
	if len(d.Nactags) > 0 {
		nactags = mistutils.ListOfStringSdkToTerraform(d.Nactags)
	}
	if len(d.OsType) > 0 {
		osType = mistutils.ListOfStringSdkToTerraform(d.OsType)
	}
	if d.PortTypes != nil {
		portTypes = matchingPortTypesSdkToTerraform(ctx, diags, d.PortTypes)
	}
	if len(d.SiteIds) > 0 {
		siteIds = mistutils.ListOfUuidSdkToTerraform(d.SiteIds)
	}
	if len(d.SitegroupIds) > 0 {
		sitegroupIds = mistutils.ListOfUuidSdkToTerraform(d.SitegroupIds)
	}
	if len(d.Vendor) > 0 {
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
	if len(d.Family) > 0 {
		family = mistutils.ListOfStringSdkToTerraform(d.Family)
	}
	if len(d.Mfg) > 0 {
		mfg = mistutils.ListOfStringSdkToTerraform(d.Mfg)
	}
	if len(d.Model) > 0 {
		model = mistutils.ListOfStringSdkToTerraform(d.Model)
	}
	if len(d.Nactags) > 0 {
		nactags = mistutils.ListOfStringSdkToTerraform(d.Nactags)
	}
	if len(d.OsType) > 0 {
		osType = mistutils.ListOfStringSdkToTerraform(d.OsType)
	}
	if d.PortTypes != nil {
		portTypes = matchingPortTypesSdkToTerraform(ctx, diags, d.PortTypes)
	}
	if len(d.SiteIds) > 0 {
		siteIds = mistutils.ListOfUuidSdkToTerraform(d.SiteIds)
	}
	if len(d.SitegroupIds) > 0 {
		sitegroupIds = mistutils.ListOfUuidSdkToTerraform(d.SitegroupIds)
	}
	if len(d.Vendor) > 0 {
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
