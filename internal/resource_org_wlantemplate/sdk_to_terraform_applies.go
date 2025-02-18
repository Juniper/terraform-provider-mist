package resource_org_wlantemplate

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"golang.org/x/net/context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appliesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.TemplateApplies) AppliesValue {

	var orgId basetypes.StringValue
	var siteIds = misttransform.ListOfUuidSdkToTerraformEmpty()
	var sitegroupIds = misttransform.ListOfUuidSdkToTerraformEmpty()

	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.SiteIds != nil {
		siteIds = misttransform.ListOfUuidSdkToTerraform(d.SiteIds)
	}
	if d.SitegroupIds != nil {
		sitegroupIds = misttransform.ListOfUuidSdkToTerraform(d.SitegroupIds)
	}

	dataMapValue := map[string]attr.Value{
		"org_id":        orgId,
		"site_ids":      siteIds,
		"sitegroup_ids": sitegroupIds,
	}
	data, e := NewAppliesValue(AppliesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
