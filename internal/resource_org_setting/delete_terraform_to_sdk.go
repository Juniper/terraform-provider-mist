package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func DeleteTerraformToSdk(ctx context.Context) (*models.OrgSetting, diag.Diagnostics) {
	var diags diag.Diagnostics
	//var data models.SiteSetting
	data := models.OrgSetting{}

	tmp := OrgSettingResourceSchema(ctx)
	unset := make(map[string]interface{})
	for k := range tmp.Attributes {
		unset["-"+k] = ""
	}
	data.AdditionalProperties = unset
	return &data, diags
}
