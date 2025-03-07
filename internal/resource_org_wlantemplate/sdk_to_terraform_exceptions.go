package resource_org_wlantemplate

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"golang.org/x/net/context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func exceptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.TemplateExceptions) ExceptionsValue {

	var siteIds = mistutils.ListOfUuidSdkToTerraformEmpty()
	var sitegroupIds = mistutils.ListOfUuidSdkToTerraformEmpty()

	if d.SiteIds != nil {
		siteIds = mistutils.ListOfUuidSdkToTerraform(d.SiteIds)
	}
	if d.SitegroupIds != nil {
		sitegroupIds = mistutils.ListOfUuidSdkToTerraform(d.SitegroupIds)
	}

	dataMapValue := map[string]attr.Value{
		"site_ids":      siteIds,
		"sitegroup_ids": sitegroupIds,
	}
	data, e := NewExceptionsValue(ExceptionsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
