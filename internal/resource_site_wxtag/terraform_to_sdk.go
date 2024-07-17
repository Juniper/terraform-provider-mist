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

	data := models.WxlanTag{}

	if !plan.Mac.IsNull() && !plan.Mac.IsUnknown() {
		data.Mac = models.NewOptional(models.ToPointer(plan.Mac.ValueString()))
	}
	if !plan.Match.IsNull() && !plan.Match.IsUnknown() {
		data.Match = models.ToPointer(models.WxlanTagMatchEnum(plan.Match.ValueString()))
	}
	data.Name = plan.Name.ValueString()
	data.Op = models.ToPointer(models.WxlanTagOperationEnum(plan.Op.ValueString()))
	if !plan.ResourceMac.IsNull() && !plan.ResourceMac.IsUnknown() {
		data.ResourceMac = models.NewOptional(models.ToPointer(plan.ResourceMac.ValueString()))
	}
	if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
		data.Services = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Services)
	}
	if !plan.Specs.IsNull() && !plan.Specs.IsUnknown() {
		data.Specs = specs
	}
	if !plan.Subnet.IsNull() && !plan.Subnet.IsUnknown() {
		data.Subnet = models.ToPointer(plan.Subnet.ValueString())
	}
	data.Type = models.WxlanTagTypeEnum(plan.Type.ValueString())
	if !plan.Values.IsNull() && !plan.Values.IsUnknown() {
		data.Values = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Values)
	}
	if !plan.VlanId.IsNull() && !plan.VlanId.IsUnknown() {
		data.VlanId = models.ToPointer(int(plan.VlanId.ValueInt64()))
	}

	return &data, diags

}
