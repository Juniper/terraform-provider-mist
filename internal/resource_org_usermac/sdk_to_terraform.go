package resource_org_usermac

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, d *models.UserMac, orgId *uuid.UUID) (OrgUsermacModel, diag.Diagnostics) {
	var state OrgUsermacModel
	var diags diag.Diagnostics

	var id types.String
	var labels types.List = types.ListNull(types.StringType)
	var mac types.String
	var name types.String
	var notes types.String
	var org_id types.String
	var radius_group types.String
	var vlan types.String

	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Labels != nil {
		labels = mist_transform.ListOfStringSdkToTerraform(ctx, d.Labels)
	}

	mac = types.StringValue(d.Mac)

	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.Notes != nil {
		notes = types.StringValue(*d.Notes)
	}

	org_id = types.StringValue(orgId.String())

	if d.RadiusGroup != nil {
		radius_group = types.StringValue(*d.RadiusGroup)
	}
	if d.Vlan != nil {
		vlan = types.StringValue(*d.Vlan)
	}

	state.Id = id
	state.Labels = labels
	state.Mac = mac
	state.Name = name
	state.Notes = notes
	state.OrgId = org_id
	state.RadiusGroup = radius_group
	state.Vlan = vlan

	return state, diags
}
