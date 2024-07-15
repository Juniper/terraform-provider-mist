package resource_org_nacrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func matchingPortTypesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.NacRuleMatchingPortTypeEnum {

	var data []models.NacRuleMatchingPortTypeEnum
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(basetypes.StringValue)
		data_item := models.NacRuleMatchingPortTypeEnum(plan.ValueString())
		data = append(data, data_item)
	}
	return data
}

func matchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MatchingValue) *models.NacRuleMatching {

	data := models.NacRuleMatching{}

	data.AuthType = models.ToPointer(models.NacRuleMatchingAuthTypeEnum(d.AuthType.ValueString()))
	data.Nactags = mist_transform.ListOfStringTerraformToSdk(ctx, d.Nactags)
	data.PortTypes = matchingPortTypesTerraformToSdk(ctx, diags, d.PortTypes)
	data.SitegroupIds = mist_transform.ListOfUuidTerraformToSdk(ctx, d.SitegroupIds)
	data.Vendor = mist_transform.ListOfStringTerraformToSdk(ctx, d.Vendor)

	return &data
}

func notMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d NotMatchingValue) *models.NacRuleMatching {

	data := models.NacRuleMatching{}

	data.AuthType = models.ToPointer(models.NacRuleMatchingAuthTypeEnum(d.AuthType.ValueString()))
	data.Nactags = mist_transform.ListOfStringTerraformToSdk(ctx, d.Nactags)
	data.PortTypes = matchingPortTypesTerraformToSdk(ctx, diags, d.PortTypes)
	data.SitegroupIds = mist_transform.ListOfUuidTerraformToSdk(ctx, d.SitegroupIds)
	data.Vendor = mist_transform.ListOfStringTerraformToSdk(ctx, d.Vendor)

	return &data
}
