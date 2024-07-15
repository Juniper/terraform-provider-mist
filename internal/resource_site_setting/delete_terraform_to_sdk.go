package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func DeleteTerraformToSdk(ctx context.Context) (*models.SiteSetting, diag.Diagnostics) {
	var diags diag.Diagnostics
	//var data models.SiteSetting
	data := models.SiteSetting{}

	tmp := SiteSettingResourceSchema(ctx)
	unset := make(map[string]interface{})
	for k := range tmp.Attributes {
		unset["-"+k] = ""
	}
	data.AdditionalProperties = unset
	return &data, diags
}
