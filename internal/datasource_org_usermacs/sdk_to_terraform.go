package datasource_org_usermacs

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, l *[]models.UserMac, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := inventorySdkToTerraform(ctx, &diags, d)
		*elements = append(*elements, elem)
	}

	return diags
}

func inventorySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.UserMac) OrgUsermacsValue {
	var id types.String
	var labels = types.ListNull(types.StringType)
	var mac types.String
	var name types.String
	var notes types.String
	var radiusGroup types.String
	var vlan types.String

	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Labels != nil {
		labels = mistutils.ListOfStringSdkToTerraform(d.Labels)
	}

	mac = types.StringValue(d.Mac)

	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.Notes != nil {
		notes = types.StringValue(*d.Notes)
	}

	if d.RadiusGroup != nil {
		radiusGroup = types.StringValue(*d.RadiusGroup)
	}
	if d.Vlan != nil {
		vlan = types.StringValue(*d.Vlan)
	}

	dataMapValue := map[string]attr.Value{
		"id":           id,
		"labels":       labels,
		"mac":          mac,
		"name":         name,
		"notes":        notes,
		"radius_group": radiusGroup,
		"vlan":         vlan,
	}
	data, e := NewOrgUsermacsValue(OrgUsermacsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
