package resource_org_nacrule

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data models.NacRule) (OrgNacruleModel, diag.Diagnostics) {
	var state OrgNacruleModel
	var diags diag.Diagnostics

	var action types.String
	var applyTags = types.ListValueMust(types.StringType, []attr.Value{})
	var enabled types.Bool
	var guestAuthState types.String
	var id types.String
	var matching = NewMatchingValueNull()
	var name types.String
	var notMatching = NewNotMatchingValueNull()
	var order types.Int64
	var orgId types.String

	action = types.StringValue(string(data.Action))
	if data.ApplyTags != nil {
		applyTags = mistutils.ListOfStringSdkToTerraform(data.ApplyTags)
	}
	if data.Enabled != nil {
		enabled = types.BoolValue(*data.Enabled)
	}
	if data.GuestAuthState != nil {
		guestAuthState = types.StringValue(string(*data.GuestAuthState))
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.Matching != nil {
		matching = matchingSdkToTerraform(ctx, &diags, data.Matching)
	}
	name = types.StringValue(data.Name)
	if data.NotMatching != nil {
		notMatching = notMatchingSdkToTerraform(ctx, &diags, data.NotMatching)
	}
	if data.Order != nil {
		order = types.Int64Value(int64(*data.Order))
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}

	state.Action = action
	state.ApplyTags = applyTags
	state.Enabled = enabled
	state.GuestAuthState = guestAuthState
	state.Id = id
	state.Matching = matching
	state.Name = name
	state.NotMatching = notMatching
	state.Order = order
	state.OrgId = orgId

	return state, diags
}
