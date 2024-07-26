package resource_org_nacrule

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data models.NacRule) (OrgNacruleModel, diag.Diagnostics) {
	var state OrgNacruleModel
	var diags diag.Diagnostics

	var action types.String
	var apply_tags types.List = types.ListNull(types.StringType)
	var enabled types.Bool
	var id types.String
	var matching MatchingValue = NewMatchingValueNull()
	var name types.String
	var not_matching NotMatchingValue = NewNotMatchingValueNull()
	var order types.Int64
	var org_id types.String

	action = types.StringValue(string(data.Action))
	if data.ApplyTags != nil {
		apply_tags = mist_transform.ListOfStringSdkToTerraform(ctx, data.ApplyTags)
	}
	if data.Enabled != nil {
		enabled = types.BoolValue(*data.Enabled)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.Matching != nil {
		matching = matchingSdkToTerraform(ctx, &diags, data.Matching)
	}
	name = types.StringValue(data.Name)
	if data.NotMatching != nil {
		not_matching = notMatchingSdkToTerraform(ctx, &diags, data.Matching)
	}
	if data.Order != nil {
		order = types.Int64Value(int64(*data.Order))
	}
	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}

	state.Action = action
	state.ApplyTags = apply_tags
	state.Enabled = enabled
	state.Id = id
	state.Matching = matching
	state.Name = name
	state.NotMatching = not_matching
	state.Order = order
	state.OrgId = org_id

	return state, diags
}
