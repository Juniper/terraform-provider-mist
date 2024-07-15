package resource_org_nacrule

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data models.NacRule) (OrgNacruleModel, diag.Diagnostics) {
	var state OrgNacruleModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Name = types.StringValue(data.Name)

	state.Action = types.StringValue(string(data.Action))
	state.ApplyTags = mist_transform.ListOfStringSdkToTerraform(ctx, data.ApplyTags)
	state.Enabled = types.BoolValue(*data.Enabled)
	state.Matching = matchingSdkToTerraform(ctx, &diags, data.Matching)
	state.NotMatching = notMatchingSdkToTerraform(ctx, &diags, data.Matching)
	state.Order = types.Int64Value(int64(*data.Order))

	return state, diags
}
