package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func syntheticTestVlansTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, l basetypes.ListValue) []models.SynthetictestProperties {
	tflog.Debug(ctx, "syntheticTestVlansTerraformToSdk")
	var data_list []models.SynthetictestProperties
	for _, v := range l.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VlansValue)
		data := models.SynthetictestProperties{}
		data.CustomTestUrls = mist_transform.ListOfStringTerraformToSdk(ctx, plan.CustomTestUrls)
		data.Disabled = plan.Disabled.ValueBoolPointer()
		data.VlanIds = mist_transform.ListOfIntTerraformToSdk(ctx, plan.VlanIds)

		data_list = append(data_list, data)
	}
	return data_list
}

func syntheticTestTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SyntheticTestValue) *models.SynthetictestConfig {
	tflog.Debug(ctx, "syntheticTestTerraformToSdk")
	data := models.SynthetictestConfig{}

	data.Disabled = d.Disabled.ValueBoolPointer()

	vlans := syntheticTestVlansTerraformToSdk(ctx, diags, d.Vlans)
	data.Vlans = vlans

	return &data
}
