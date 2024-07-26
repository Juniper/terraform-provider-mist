package resource_org_nactag

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgNactagModel) (models.NacTag, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})
	data := models.NacTag{}

	if !plan.AllowUsermacOverride.IsNull() && !plan.AllowUsermacOverride.IsUnknown() {
		data.AllowUsermacOverride = plan.AllowUsermacOverride.ValueBoolPointer()
	} else {
		unset["-allow_usermac_override"] = ""
	}
	if !plan.EgressVlanNames.IsNull() && !plan.EgressVlanNames.IsUnknown() {
		data.EgressVlanNames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.EgressVlanNames)
	} else {
		unset["-egress_vlan_names"] = ""
	}
	if !plan.GbpTag.IsNull() && !plan.GbpTag.IsUnknown() {
		data.GbpTag = models.ToPointer(int(plan.GbpTag.ValueInt64()))
	} else {
		unset["-gbp_tag"] = ""
	}
	if !plan.Match.IsNull() && !plan.Match.IsUnknown() {
		data.Match = models.ToPointer(models.NacTagMatchEnum(plan.Match.ValueString()))
	} else {
		unset["-match"] = ""
	}
	if !plan.MatchAll.IsNull() && !plan.MatchAll.IsUnknown() {
		data.MatchAll = plan.MatchAll.ValueBoolPointer()
	} else {
		unset["-match_all"] = ""
	}
	data.Name = plan.Name.ValueString()
	if !plan.RadiusAttrs.IsNull() && !plan.RadiusAttrs.IsUnknown() {
		data.RadiusAttrs = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RadiusAttrs)
	} else {
		unset["-radius_attrs"] = ""
	}
	if !plan.RadiusGroup.IsNull() && !plan.RadiusGroup.IsUnknown() {
		data.RadiusGroup = plan.RadiusGroup.ValueStringPointer()
	} else {
		unset["-radius_group"] = ""
	}
	if !plan.RadiusVendorAttrs.IsNull() && !plan.RadiusVendorAttrs.IsUnknown() {
		data.RadiusVendorAttrs = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RadiusVendorAttrs)
	} else {
		unset["-radius_vendor_attrs"] = ""
	}
	if !plan.SessionTimeout.IsNull() && !plan.SessionTimeout.IsUnknown() {
		data.SessionTimeout = models.ToPointer(int(plan.SessionTimeout.ValueInt64()))
	} else {
		unset["-allow_usermac_override"] = ""
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		data.Type = models.NacTagTypeEnum(plan.Type.ValueString())
	} else {
		unset["-type"] = ""
	}
	if !plan.Values.IsNull() && !plan.Values.IsUnknown() {
		data.Values = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Values)
	} else {
		unset["-values"] = ""
	}
	if !plan.Vlan.IsNull() && !plan.Vlan.IsUnknown() {
		data.Vlan = models.ToPointer(plan.Vlan.ValueString())
	} else {
		unset["-vlan"] = ""
	}

	data.AdditionalProperties = unset
	return data, diags
}
