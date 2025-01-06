package resource_org_usermac

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgUsermacModel) (models.UserMac, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.UserMac{}
	unset := make(map[string]interface{})

	if !plan.Labels.IsNull() && !plan.Labels.IsUnknown() {
		data.Labels = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Labels)
	} else {
		unset["-labels"] = ""
	}

	if plan.Mac.ValueStringPointer() != nil {
		data.Mac = plan.Mac.ValueString()
	} else {
		unset["-mac"] = ""
	}

	if plan.Name.ValueStringPointer() != nil {
		data.Name = plan.Name.ValueStringPointer()
	} else {
		unset["-name"] = ""
	}

	if plan.Notes.ValueStringPointer() != nil {
		data.Notes = plan.Notes.ValueStringPointer()
	} else {
		unset["-notes"] = ""
	}

	if plan.RadiusGroup.ValueStringPointer() != nil {
		data.RadiusGroup = plan.RadiusGroup.ValueStringPointer()
	} else {
		unset["-radius_group"] = ""
	}

	if plan.Vlan.ValueStringPointer() != nil {
		data.Vlan = plan.Vlan.ValueStringPointer()
	} else {
		unset["-vlan"] = ""
	}

	data.AdditionalProperties = unset
	return data, diags
}
