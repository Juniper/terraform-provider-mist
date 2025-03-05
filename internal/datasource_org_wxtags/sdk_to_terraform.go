package datasource_org_wxtags

import (
	"context"
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.WxlanTag, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := wxtagSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func wxtagSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WxlanTag) OrgWxtagsValue {

	var createdTime basetypes.Float64Value
	var id basetypes.StringValue
	var mac types.String
	var match types.String
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var op types.String
	var specs = types.ListNull(SpecsValue{}.Type(ctx))
	var values = types.ListNull(types.StringType)
	var vlanId types.String

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Mac.Value() != nil {
		mac = types.StringValue(*d.Mac.Value())
	}
	if d.Match != nil {
		match = types.StringValue(string(*d.Match))
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	name = types.StringValue(d.Name)
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Op != nil {
		op = types.StringValue(string(*d.Op))
	}
	if d.Specs != nil {
		specs = specsSdkToTerraform(ctx, diags, d.Specs)
	}
	if d.Values != nil {
		values = misttransform.ListOfStringSdkToTerraform(d.Values)
	}
	if d.VlanId != nil {
		vlanId = types.StringValue(d.VlanId.String())
	}

	dataMapValue := map[string]attr.Value{
		"created_time":  createdTime,
		"id":            id,
		"mac":           mac,
		"match":         match,
		"modified_time": modifiedTime,
		"name":          name,
		"org_id":        orgId,
		"op":            op,
		"specs":         specs,
		"values":        values,
		"vlan_id":       vlanId,
	}
	data, e := NewOrgWxtagsValue(OrgWxtagsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
