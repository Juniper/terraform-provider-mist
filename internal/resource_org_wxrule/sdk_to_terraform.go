package resource_org_wxrule

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.WxlanRule) (OrgWxruleModel, diag.Diagnostics) {
	var state OrgWxruleModel
	var diags diag.Diagnostics

	var action types.String
	var apply_tags types.List = types.ListNull(types.StringType)
	var blocked_apps types.List = types.ListNull(types.StringType)
	var dst_allow_wxtags types.List = types.ListNull(types.StringType)
	var dst_deny_wxtags types.List = types.ListNull(types.StringType)
	var dst_wxtags types.List = types.ListNull(types.StringType)
	var enabled types.Bool
	var src_wxtags types.List = types.ListNull(types.StringType)

	if data.Action != nil {
		action = types.StringValue(string(*data.Action))
	}
	if data.ApplyTags != nil && len(data.ApplyTags) > 0 {
		apply_tags = mist_transform.ListOfStringSdkToTerraform(ctx, data.ApplyTags)
	}
	if data.BlockedApps != nil && len(data.BlockedApps) > 0 {
		blocked_apps = mist_transform.ListOfStringSdkToTerraform(ctx, data.BlockedApps)
	}
	if data.DstAllowWxtags != nil && len(data.DstAllowWxtags) > 0 {
		dst_allow_wxtags = mist_transform.ListOfStringSdkToTerraform(ctx, data.DstAllowWxtags)
	}
	if data.DstDenyWxtags != nil && len(data.DstDenyWxtags) > 0 {
		dst_deny_wxtags = mist_transform.ListOfStringSdkToTerraform(ctx, data.DstDenyWxtags)
	}
	if data.DstWxtags != nil {
		dst_wxtags = mist_transform.ListOfStringSdkToTerraform(ctx, data.DstWxtags)
	}
	if data.Enabled != nil {
		enabled = types.BoolValue(*data.Enabled)
	}
	if data.SrcWxtags != nil && len(data.SrcWxtags) > 0 {
		src_wxtags = mist_transform.ListOfStringSdkToTerraform(ctx, data.SrcWxtags)
	}

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.TemplateId = types.StringValue(data.TemplateId.String())
	state.Order = types.Int64Value(int64(data.Order))

	state.Action = action
	state.ApplyTags = apply_tags
	state.BlockedApps = blocked_apps
	state.DstAllowWxtags = dst_allow_wxtags
	state.DstDenyWxtags = dst_deny_wxtags
	state.DstWxtags = dst_wxtags
	state.Enabled = enabled
	state.SrcWxtags = src_wxtags

	return state, diags
}
