package resource_org_nacrule

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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
	var nactags = misttransform.ListOfStringSdkToTerraformEmpty()
	var portTypes = misttransform.ListOfStringSdkToTerraformEmpty()
	var siteIds = misttransform.ListOfStringSdkToTerraformEmpty()
	var sitegroupIds = misttransform.ListOfStringSdkToTerraformEmpty()
	var vendor = misttransform.ListOfStringSdkToTerraformEmpty()

	if d.AuthType != nil {
		authType = types.StringValue(string(*d.AuthType))
	}
	if d.Nactags != nil {
		nactags = misttransform.ListOfStringSdkToTerraform(d.Nactags)
	}
	if d.PortTypes != nil {
		portTypes = matchingPortTypesSdkToTerraform(ctx, diags, d.PortTypes)
	}
	if d.SiteIds != nil {
		siteIds = misttransform.ListOfUuidSdkToTerraform(d.SiteIds)
	}
	if d.SitegroupIds != nil {
		sitegroupIds = misttransform.ListOfUuidSdkToTerraform(d.SitegroupIds)
	}
	if d.Vendor != nil {
		vendor = misttransform.ListOfStringSdkToTerraform(d.Vendor)
	}

	dataMapValue := map[string]attr.Value{
		"auth_type":     authType,
		"nactags":       nactags,
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
	var nactags = misttransform.ListOfStringSdkToTerraformEmpty()
	var portTypes = misttransform.ListOfStringSdkToTerraformEmpty()
	var siteIds = misttransform.ListOfStringSdkToTerraformEmpty()
	var sitegroupIds = misttransform.ListOfStringSdkToTerraformEmpty()
	var vendor = misttransform.ListOfStringSdkToTerraformEmpty()

	if d.AuthType != nil {
		authType = types.StringValue(string(*d.AuthType))
	}
	if d.Nactags != nil {
		nactags = misttransform.ListOfStringSdkToTerraform(d.Nactags)
	}
	if d.PortTypes != nil {
		portTypes = matchingPortTypesSdkToTerraform(ctx, diags, d.PortTypes)
	}
	if d.SiteIds != nil {
		siteIds = misttransform.ListOfUuidSdkToTerraform(d.SiteIds)
	}
	if d.SitegroupIds != nil {
		sitegroupIds = misttransform.ListOfUuidSdkToTerraform(d.SitegroupIds)
	}
	if d.Vendor != nil {
		vendor = misttransform.ListOfStringSdkToTerraform(d.Vendor)
	}

	dataMapValue := map[string]attr.Value{
		"auth_type":     authType,
		"nactags":       nactags,
		"port_types":    portTypes,
		"site_ids":      siteIds,
		"sitegroup_ids": sitegroupIds,
		"vendor":        vendor,
	}
	data, e := NewNotMatchingValue(NotMatchingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
