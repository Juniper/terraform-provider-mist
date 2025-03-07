package datasource_org_networks

import (
	"context"
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.Network, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := networkSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func networkSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Network) OrgNetworksValue {

	var createdTime basetypes.Float64Value
	var id basetypes.StringValue
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var subnet basetypes.StringValue
	var subnet6 basetypes.StringValue
	var vlanId basetypes.StringValue

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	name = types.StringValue(d.Name)
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Subnet != nil {
		subnet = types.StringValue(*d.Subnet)
	}
	if d.Subnet6 != nil {
		subnet = types.StringValue(*d.Subnet6)
	}
	if d.VlanId != nil {
		vlanId = mistutils.VlanAsString(*d.VlanId)
	}

	dataMapValue := map[string]attr.Value{
		"created_time":  createdTime,
		"id":            id,
		"modified_time": modifiedTime,
		"name":          name,
		"org_id":        orgId,
		"subnet":        subnet,
		"subnet6":       subnet6,
		"vlan_id":       vlanId,
	}
	data, e := NewOrgNetworksValue(OrgNetworksValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
