package resource_org_wxtag

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgWxtagModel) (*models.WxlanTag, diag.Diagnostics) {
	var diags diag.Diagnostics
	specs := specsTerraformToSdk(ctx, &diags, plan.Specs)

	data := models.WxlanTag{}

	if plan.Mac.ValueStringPointer() != nil {
		data.Mac = models.NewOptional(models.ToPointer(plan.Mac.ValueString()))
	}

	if plan.Match.ValueStringPointer() != nil {
		data.Match = models.ToPointer(models.WxlanTagMatchEnum(plan.Match.ValueString()))
	}

	data.Name = plan.Name.ValueString()

	data.Op = models.ToPointer(models.WxlanTagOperationEnum(plan.Op.ValueString()))

	if plan.ResourceMac.ValueStringPointer() != nil {
		data.ResourceMac = models.NewOptional(models.ToPointer(plan.ResourceMac.ValueString()))
	}

	if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
		data.Services = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Services)
	}

	if !plan.Specs.IsNull() && !plan.Specs.IsUnknown() {
		data.Specs = specs
	}

	if plan.Subnet.ValueStringPointer() != nil {
		data.Subnet = models.ToPointer(plan.Subnet.ValueString())
	}

	data.Type = models.WxlanTagTypeEnum(plan.Type.ValueString())

	if !plan.Values.IsNull() && !plan.Values.IsUnknown() {
		data.Values = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Values)
	}

	if plan.VlanId.ValueStringPointer() != nil {
		data.VlanId = models.ToPointer(models.WxlanTagVlanIdContainer.FromString(plan.VlanId.ValueString()))
	}

	return &data, diags

}
