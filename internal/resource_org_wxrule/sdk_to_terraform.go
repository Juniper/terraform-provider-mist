package resource_org_wxrule

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(data *models.WxlanRule) (OrgWxruleModel, diag.Diagnostics) {
	var state OrgWxruleModel
	var diags diag.Diagnostics

	var action types.String
	var applyTags = types.ListNull(types.StringType)
	var blockedApps = types.ListNull(types.StringType)
	var dstAllowWxtags = types.ListNull(types.StringType)
	var dstDenyWxtags = types.ListNull(types.StringType)
	var dstWxtags = types.ListNull(types.StringType)
	var enabled types.Bool
	var srcWxtags = types.ListNull(types.StringType)

	if data.Action != nil {
		action = types.StringValue(string(*data.Action))
	}
	if data.ApplyTags != nil && len(data.ApplyTags) > 0 {
		applyTags = mistutils.ListOfStringSdkToTerraform(data.ApplyTags)
	}
	if data.BlockedApps != nil && len(data.BlockedApps) > 0 {
		blockedApps = mistutils.ListOfStringSdkToTerraform(data.BlockedApps)
	}
	if data.DstAllowWxtags != nil {
		dstAllowWxtags = mistutils.ListOfStringSdkToTerraform(data.DstAllowWxtags)
	}
	if data.DstDenyWxtags != nil {
		dstDenyWxtags = mistutils.ListOfStringSdkToTerraform(data.DstDenyWxtags)
	}
	if data.DstWxtags != nil {
		dstWxtags = mistutils.ListOfStringSdkToTerraform(data.DstWxtags)
	}
	if data.Enabled != nil {
		enabled = types.BoolValue(*data.Enabled)
	}
	if data.SrcWxtags != nil {
		srcWxtags = mistutils.ListOfStringSdkToTerraform(data.SrcWxtags)
	}

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.TemplateId = types.StringValue(data.TemplateId.String())
	state.Order = types.Int64Value(int64(data.Order))

	state.Action = action
	state.ApplyTags = applyTags
	state.BlockedApps = blockedApps
	state.DstAllowWxtags = dstAllowWxtags
	state.DstDenyWxtags = dstDenyWxtags
	state.DstWxtags = dstWxtags
	state.Enabled = enabled
	state.SrcWxtags = srcWxtags

	return state, diags
}
