package resource_site_wxrule

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *SiteWxruleModel) (*models.WxlanRule, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.WxlanRule{}

	if !plan.Action.IsNull() && !plan.Action.IsUnknown() {
		data.Action = (*models.WxlanRuleActionEnum)(plan.Action.ValueStringPointer())
	}
	if !plan.ApplyTags.IsNull() && !plan.ApplyTags.IsUnknown() {
		data.ApplyTags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ApplyTags)
	}
	if !plan.BlockedApps.IsNull() && !plan.BlockedApps.IsUnknown() {
		data.BlockedApps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.BlockedApps)
	}
	if !plan.DstAllowWxtags.IsNull() && !plan.DstAllowWxtags.IsUnknown() {
		data.DstAllowWxtags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstAllowWxtags)
	}
	if !plan.DstDenyWxtags.IsNull() && !plan.DstDenyWxtags.IsUnknown() {
		data.DstDenyWxtags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstDenyWxtags)
	}
	if !plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
		data.Enabled = plan.Enabled.ValueBoolPointer()
	}
	if !plan.Order.IsNull() && !plan.Order.IsUnknown() {
		data.Order = int(plan.Order.ValueInt64())
	}
	if !plan.SrcWxtags.IsNull() && !plan.SrcWxtags.IsUnknown() {
		data.SrcWxtags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.SrcWxtags)
	} else {
		data.SrcWxtags = make([]string, 0)
	}

	return &data, diags

}
