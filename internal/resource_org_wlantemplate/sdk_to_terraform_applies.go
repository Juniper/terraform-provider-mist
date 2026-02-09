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
