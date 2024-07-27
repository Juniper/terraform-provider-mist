package resource_site_wxtag

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *SiteWxtagModel) (*models.WxlanTag, diag.Diagnostics) {
	var diags diag.Diagnostics
	specs := specsTerraformToSdk(ctx, &diags, plan.Specs)
	unset := make(map[string]interface{})
	data := models.WxlanTag{}

	if !plan.Mac.IsNull() && !plan.Mac.IsUnknown() {
		data.Mac = models.NewOptional(models.ToPointer(plan.Mac.ValueString()))
	} else {
		unset["-mac"] = ""
	}

	if !plan.Match.IsNull() && !plan.Match.IsUnknown() {
		data.Match = models.ToPointer(models.WxlanTagMatchEnum(plan.Match.ValueString()))
	} else {
		unset["-match"] = ""
	}

	data.Name = plan.Name.ValueString()

	if !plan.Op.IsNull() && !plan.Op.IsUnknown() {
		data.Op = models.ToPointer(models.WxlanTagOperationEnum(plan.Op.ValueString()))
	} else {
		unset["-op"] = ""
	}

	if !plan.Specs.IsNull() && !plan.Specs.IsUnknown() {
		data.Specs = specs
	} else {
		unset["-specs"] = ""
	}

	data.Type = models.WxlanTagTypeEnum(plan.Type.ValueString())

	if !plan.Values.IsNull() && !plan.Values.IsUnknown() {
		data.Values = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Values)
	} else {
		unset["-values"] = ""
	}

	if !plan.VlanId.IsNull() && !plan.VlanId.IsUnknown() {
		data.VlanId = models.ToPointer(models.WxlanTagVlanIdContainer.FromString(plan.VlanId.ValueString()))
	} else {
		unset["-vlan_id"] = ""
	}

	data.AdditionalProperties = unset
	return &data, diags
}
