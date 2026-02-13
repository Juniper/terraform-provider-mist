package resource_org_wlantemplate

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data models.Template) (OrgWlantemplateModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	state := OrgWlantemplateModel{
		Id:               types.StringValue(data.Id.String()),
		OrgId:            types.StringValue(data.OrgId.String()),
		Name:             types.StringValue(data.Name),
		DeviceprofileIds: mistutils.ListOfUuidSdkToTerraform(data.DeviceprofileIds),
	}

	if data.Applies != nil {
		state.Applies = appliesSdkToTerraform(ctx, &diags, *data.Applies)
	}

	if data.Exceptions != nil {
		state.Exceptions = exceptionsSdkToTerraform(ctx, &diags, *data.Exceptions)
	}

	if data.FilterByDeviceprofile != nil {
		state.FilterByDeviceprofile = types.BoolValue(*data.FilterByDeviceprofile)
	}

	return state, diags
}

func exceptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data models.TemplateExceptions) ExceptionsValue {
	var siteIds = mistutils.ListOfUuidSdkToTerraformEmpty()
	if data.SiteIds != nil {
		siteIds = mistutils.ListOfUuidSdkToTerraform(data.SiteIds)
	}

	var sitegroupIds = mistutils.ListOfUuidSdkToTerraformEmpty()
	if data.SitegroupIds != nil {
		sitegroupIds = mistutils.ListOfUuidSdkToTerraform(data.SitegroupIds)
	}

	dataMapValue := map[string]attr.Value{
		"site_ids":      siteIds,
		"sitegroup_ids": sitegroupIds,
	}
	result, e := NewExceptionsValue(ExceptionsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return result
}

func appliesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data models.TemplateApplies) AppliesValue {
	var orgId basetypes.StringValue
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}

	var siteIds = mistutils.ListOfUuidSdkToTerraformEmpty()
	if data.SiteIds != nil {
		siteIds = mistutils.ListOfUuidSdkToTerraform(data.SiteIds)
	}

	var sitegroupIds = mistutils.ListOfUuidSdkToTerraformEmpty()
	if data.SitegroupIds != nil {
		sitegroupIds = mistutils.ListOfUuidSdkToTerraform(data.SitegroupIds)
	}

	dataMap := map[string]attr.Value{
		"org_id":        orgId,
		"site_ids":      siteIds,
		"sitegroup_ids": sitegroupIds,
	}
	result, err := NewAppliesValue(AppliesValue{}.AttributeTypes(ctx), dataMap)
	diags.Append(err...)

	return result
}
