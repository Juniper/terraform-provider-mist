package resource_org_wlantemplate

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data models.Template) (OrgWlantemplateModel, diag.Diagnostics) {
	var state OrgWlantemplateModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Name = types.StringValue(data.Name)

	if data.Applies != nil {
		state.Applies = appliesSdkToTerraform(ctx, &diags, *data.Applies)
	}

	state.DeviceprofileIds = mistutils.ListOfUuidSdkToTerraform(data.DeviceprofileIds)

	if data.Exceptions != nil {
		state.Exceptions = exceptionsSdkToTerraform(ctx, &diags, *data.Exceptions)
	}

	if data.FilterByDeviceprofile != nil {
		state.FilterByDeviceprofile = types.BoolValue(*data.FilterByDeviceprofile)
	}

	return state, diags
}
