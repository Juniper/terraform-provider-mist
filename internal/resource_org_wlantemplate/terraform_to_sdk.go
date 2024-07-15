package resource_org_wlantemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgWlantemplateModel) (*models.Template, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.Template{}
	data.Name = plan.Name.ValueString()

	data.Applies = appliesTerraformToSdk(ctx, &diags, plan.Applies)

	data.DeviceprofileIds = mist_transform.ListOfUuidTerraformToSdk(ctx, plan.DeviceprofileIds)

	data.Exceptions = exceptionsTerraformToSdk(ctx, &diags, plan.Exceptions)

	data.FilterByDeviceprofile = models.ToPointer(plan.FilterByDeviceprofile.ValueBool())

	return &data, diags
}
