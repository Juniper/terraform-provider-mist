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
	data.Mac = models.NewOptional(models.ToPointer(plan.Mac.ValueString()))
	data.Match = models.ToPointer(models.WxlanTagMatchEnum(plan.Match.ValueString()))
	data.Name = plan.Name.ValueString()
	data.Op = models.ToPointer(models.WxlanTagOperationEnum(plan.Op.ValueString()))
	data.ResourceMac = models.NewOptional(models.ToPointer(plan.ResourceMac.ValueString()))
	data.Services = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Services)
	data.Specs = specs
	data.Subnet = models.ToPointer(plan.Subnet.ValueString())
	data.Type = models.WxlanTagTypeEnum(plan.Type.ValueString())
	data.Values = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Values)
	data.VlanId = models.ToPointer(int(plan.VlanId.ValueInt64()))

	return &data, diags

}
