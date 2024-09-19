package resource_org_wxrule

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgWxruleModel) (*models.WxlanRule, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})
	data := models.WxlanRule{}

	if !plan.Action.IsNull() && !plan.Action.IsUnknown() {
		data.Action = (*models.WxlanRuleActionEnum)(plan.Action.ValueStringPointer())
	} else {
		unset["-action"] = ""
	}

	if !plan.ApplyTags.IsNull() && !plan.ApplyTags.IsUnknown() {
		data.ApplyTags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ApplyTags)
	} else {
		unset["-apply_tags"] = ""
	}

	if !plan.BlockedApps.IsNull() && !plan.BlockedApps.IsUnknown() {
		data.BlockedApps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.BlockedApps)
	} else {
		unset["-blocked_apps"] = ""
	}

	if !plan.DstAllowWxtags.IsNull() && !plan.DstAllowWxtags.IsUnknown() {
		data.DstAllowWxtags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstAllowWxtags)
	}
	if data.DstAllowWxtags == nil {
		data.DstAllowWxtags = make([]string, 0)
	}

	if !plan.DstDenyWxtags.IsNull() && !plan.DstDenyWxtags.IsUnknown() {
		data.DstDenyWxtags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstDenyWxtags)
	}
	if data.DstDenyWxtags == nil {
		data.DstDenyWxtags = make([]string, 0)
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

	if len(plan.TemplateId.ValueString()) > 0 {
		template_id, e := uuid.Parse(plan.TemplateId.ValueString())
		if e == nil {
			data.TemplateId = models.ToPointer(template_id)
		} else {
			diags.AddError("Bad value for template_id", e.Error())
		}
	}

	data.AdditionalProperties = unset
	return &data, diags
}
