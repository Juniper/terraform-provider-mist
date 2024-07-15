package resource_org_nactag

import (
	"context"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgNactagModel) (models.NacTag, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.NacTag{}
	data.AllowUsermacOverride = plan.AllowUsermacOverride.ValueBoolPointer()
	data.EgressVlanNames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.EgressVlanNames)
	data.GbpTag = models.ToPointer(int(plan.GbpTag.ValueInt64()))
	data.Match = models.ToPointer(models.NacTagMatchEnum(plan.Match.ValueString()))
	data.MatchAll = plan.MatchAll.ValueBoolPointer()
	data.Name = plan.Name.ValueString()
	data.RadiusAttrs = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RadiusAttrs)
	data.RadiusGroup = plan.RadiusGroup.ValueStringPointer()
	data.RadiusVendorAttrs = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RadiusVendorAttrs)
	data.SessionTimeout = models.ToPointer(int(plan.SessionTimeout.ValueInt64()))
	data.Type = models.NacTagTypeEnum(plan.Type.ValueString())
	data.Values = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Values)
	data.Vlan = models.ToPointer(plan.Vlan.ValueString())

	return data, diags
}
