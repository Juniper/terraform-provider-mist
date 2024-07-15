package resource_org_wlan

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ciscoCwaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan CiscoCwaValue) *models.WlanCiscoCwa {

	data := models.WlanCiscoCwa{}
	data.AllowedHostnames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AllowedHostnames)
	data.AllowedSubnets = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AllowedSubnets)
	data.BlockedSubnets = mist_transform.ListOfStringTerraformToSdk(ctx, plan.BlockedSubnets)
	data.Enabled = plan.Enabled.ValueBoolPointer()

	return &data
}
