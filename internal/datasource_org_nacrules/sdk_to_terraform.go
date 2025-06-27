package datasource_org_nacrules

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.NacRule, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := nacruleSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func nacruleSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NacRule) OrgNacrulesValue {

	var createdTime basetypes.Float64Value
	var enabled basetypes.BoolValue
	var id basetypes.StringValue
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var order basetypes.Int64Value
	var orgId basetypes.StringValue

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	if d.Order != nil {
		order = types.Int64Value(int64(*d.Order))
	}
	name = types.StringValue(d.Name)
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}

	dataMapValue := map[string]attr.Value{
		"created_time":  createdTime,
		"enabled":       enabled,
		"id":            id,
		"modified_time": modifiedTime,
		"name":          name,
		"order":         order,
		"org_id":        orgId,
	}
	data, e := NewOrgNacrulesValue(OrgNacrulesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
