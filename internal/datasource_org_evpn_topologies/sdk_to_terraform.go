package datasource_org_evpn_topologies

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

func evpnTopologySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.EvpnTopologyResponse) EvpnTopologiesValue {

	var created_time basetypes.Float64Value
	var evpn_options basetypes.ObjectValue = basetypes.NewObjectNull(EvpnOptionsValue{}.AttributeTypes(ctx))
	var id basetypes.StringValue
	var modified_time basetypes.Float64Value
	var name basetypes.StringValue
	var org_id basetypes.StringValue
	var pod_names basetypes.MapValue = basetypes.NewMapNull(types.StringType)

	if d.CreatedTime != nil {
		created_time = types.Float64Value(float64(*d.CreatedTime))
	}
	if d.EvpnOptions != nil {
		evpn_options = evpnOptionsSdkToTerraform(ctx, diags, d.EvpnOptions)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modified_time = types.Float64Value(float64(*d.ModifiedTime))
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.PodNames != nil {
		data_map := make(map[string]string)
		for k, v := range d.PodNames {
			data_map[k] = v
		}
		state_result, e := types.MapValueFrom(ctx, types.StringType, data_map)
		diags.Append(e...)
		pod_names = state_result
	}

	data_map_value := map[string]attr.Value{
		"created_time":  created_time,
		"evpn_options":  evpn_options,
		"id":            id,
		"modified_time": modified_time,
		"name":          name,
		"org_id":        org_id,
		"pod_names":     pod_names,
	}
	data, e := NewEvpnTopologiesValue(EvpnTopologiesValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}
