package datasource_site_evpn_topologies

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.EvpnTopologyResponse, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := evpnTopologySdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func evpnTopologySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.EvpnTopologyResponse) SiteEvpnTopologiesValue {

	var createdTime basetypes.Float64Value
	var evpnOptions = basetypes.NewObjectNull(SiteEvpnTopologiesValue{}.AttributeTypes(ctx))
	var id basetypes.StringValue
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var siteId basetypes.StringValue
	var podNames = basetypes.NewMapNull(types.StringType)

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.EvpnOptions != nil {
		evpnOptions = evpnOptionsSdkToTerraform(ctx, diags, d.EvpnOptions)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.SiteId != nil {
		siteId = types.StringValue(d.SiteId.String())
	}
	if d.PodNames != nil {
		dataMap := make(map[string]string)
		for k, v := range d.PodNames {
			dataMap[k] = v
		}
		stateResult, e := types.MapValueFrom(ctx, types.StringType, dataMap)
		diags.Append(e...)
		podNames = stateResult
	}

	dataMapValue := map[string]attr.Value{
		"created_time":  createdTime,
		"evpn_options":  evpnOptions,
		"id":            id,
		"modified_time": modifiedTime,
		"name":          name,
		"site_id":       siteId,
		"pod_names":     podNames,
	}
	data, e := NewSiteEvpnTopologiesValue(SiteEvpnTopologiesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
