package resource_org_usermac

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(d *models.UserMac, mistOrgId *uuid.UUID) (OrgUsermacModel, diag.Diagnostics) {
	var state OrgUsermacModel
	var diags diag.Diagnostics

	var id types.String
	var labels = types.ListNull(types.StringType)
	var mac types.String
	var name types.String
	var notes types.String
	var orgId types.String
	var radiusGroup types.String
	var vlan types.String

	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Labels != nil {
		labels = misttransform.ListOfStringSdkToTerraform(d.Labels)
	}

	mac = types.StringValue(d.Mac)

	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.Notes != nil {
		notes = types.StringValue(*d.Notes)
	}

	orgId = types.StringValue(mistOrgId.String())

	if d.RadiusGroup != nil {
		radiusGroup = types.StringValue(*d.RadiusGroup)
	}
	if d.Vlan != nil {
		vlan = types.StringValue(*d.Vlan)
	}

	state.Id = id
	state.Labels = labels
	state.Mac = mac
	state.Name = name
	state.Notes = notes
	state.OrgId = orgId
	state.RadiusGroup = radiusGroup
	state.Vlan = vlan

	return state, diags
}
