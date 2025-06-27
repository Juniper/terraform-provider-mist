package resource_org_wlantemplate

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"golang.org/x/net/context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appliesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.TemplateApplies) AppliesValue {

	var orgId basetypes.StringValue
	var siteIds = mistutils.ListOfUuidSdkToTerraformEmpty()
	var sitegroupIds = mistutils.ListOfUuidSdkToTerraformEmpty()

	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.SiteIds != nil {
		siteIds = mistutils.ListOfUuidSdkToTerraform(d.SiteIds)
	}
	if d.SitegroupIds != nil {
		sitegroupIds = mistutils.ListOfUuidSdkToTerraform(d.SitegroupIds)
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
