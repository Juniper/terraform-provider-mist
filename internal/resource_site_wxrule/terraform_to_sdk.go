package resource_site_wxrule

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *SiteWxruleModel) (*models.WxlanRule, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.WxlanRule{}

	data.Action = (*models.WxlanRuleActionEnum)(plan.Action.ValueStringPointer())
	data.ApplyTags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ApplyTags)
	data.BlockedApps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.BlockedApps)
	data.DstAllowWxtags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstAllowWxtags)
	data.DstDenyWxtags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstDenyWxtags)
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.Order = int(plan.Order.ValueInt64())
	data.SrcWxtags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.SrcWxtags)
	if len(plan.TemplateId.ValueString()) > 0 {
		template_id, e := uuid.Parse(plan.TemplateId.ValueString())
		if e != nil {
			diags.AddError("Unable to convert IdpprofileId", e.Error())
		} else {
			data.TemplateId = models.ToPointer(template_id)
		}
	}
	return &data, diags

}
