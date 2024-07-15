package datasource_org_gatewaytemplates

import (
	"context"
	"math/big"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.GatewayTemplate) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := gatewaytemplateSdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(OrgGatewaytemplatesValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func gatewaytemplateSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.GatewayTemplate) OrgGatewaytemplatesValue {

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

	data_map_attr_type := OrgGatewaytemplatesValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"created_time":  created_time,
		"id":            id,
		"modified_time": modified_time,
		"name":          name,
		"org_id":        org_id,
	}
	data, e := NewOrgGatewaytemplatesValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
