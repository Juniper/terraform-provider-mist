package resource_org_wlantemplate

import (
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"golang.org/x/net/context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func exceptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.TemplateExceptions) ExceptionsValue {

	var site_ids basetypes.ListValue = mist_transform.ListOfUuidSdkToTerraformEmpty(ctx)
	var sitegroup_ids basetypes.ListValue = mist_transform.ListOfUuidSdkToTerraformEmpty(ctx)

	if d.SiteIds != nil {
		site_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, d.SiteIds)
	}
	if d.SitegroupIds != nil {
		sitegroup_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, d.SitegroupIds)
	}

	data_map_attr_type := ExceptionsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"site_ids":      site_ids,
		"sitegroup_ids": sitegroup_ids,
	}
	data, e := NewExceptionsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
