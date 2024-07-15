package resource_org_nacrule

import (
	"context"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgNacruleModel) (models.NacRule, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.NacRule{}

	data.Action = *models.ToPointer(models.NacRuleActionEnum(plan.Action.ValueString()))
	data.ApplyTags = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ApplyTags)
	data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
	data.Matching = matchingTerraformToSdk(ctx, &diags, plan.Matching)
	data.Name = plan.Name.ValueString()
	data.NotMatching = notMatchingTerraformToSdk(ctx, &diags, plan.NotMatching)
	data.Order = models.ToPointer(int(plan.Order.ValueInt64()))

	return data, diags
}
