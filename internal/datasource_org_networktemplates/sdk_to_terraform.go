package datasource_org_networktemplates

import (
	"context"
	"math/big"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.NetworkTemplate, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := networktemplateSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func networktemplateSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NetworkTemplate) OrgNetworktemplatesValue {

	var createdTime basetypes.NumberValue
	var id basetypes.StringValue
	var modifiedTime basetypes.NumberValue
	var name basetypes.StringValue
	var orgId basetypes.StringValue

	if d.CreatedTime != nil {
		createdTime = types.NumberValue(big.NewFloat(*d.CreatedTime))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.NumberValue(big.NewFloat(*d.ModifiedTime))
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}

	dataMapValue := map[string]attr.Value{
		"created_time":  createdTime,
		"id":            id,
		"modified_time": modifiedTime,
		"name":          name,
		"org_id":        orgId,
	}
	data, e := NewOrgNetworktemplatesValue(OrgNetworktemplatesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
