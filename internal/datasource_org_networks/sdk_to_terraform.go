package datasource_org_networks

import (
	"context"
	"math/big"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.Network) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := networkSdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(OrgNetworksValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func networkSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.Network) OrgNetworksValue {

	var created_time basetypes.NumberValue
	var id basetypes.StringValue
	var modified_time basetypes.NumberValue
	var name basetypes.StringValue
	var org_id basetypes.StringValue

	if d.CreatedTime != nil {
		created_time = types.NumberValue(big.NewFloat(*d.CreatedTime))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modified_time = types.NumberValue(big.NewFloat(*d.ModifiedTime))
	}
	name = types.StringValue(d.Name)
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}

	data_map_attr_type := OrgNetworksValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"created_time":  created_time,
		"id":            id,
		"modified_time": modified_time,
		"name":          name,
		"org_id":        org_id,
	}
	data, e := NewOrgNetworksValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
