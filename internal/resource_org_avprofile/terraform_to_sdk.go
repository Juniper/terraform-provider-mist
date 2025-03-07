package resource_org_avprofile

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(plan *OrgAvprofileModel) (*models.Avprofile, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})
	data := models.Avprofile{}

	if plan.FallbackAction.ValueStringPointer() != nil {
		data.FallbackAction = (*models.AvprofileFallbackActionEnum)(plan.FallbackAction.ValueStringPointer())
	} else {
		unset["-fallback_action"] = ""
	}
	if plan.MaxFilesize.ValueInt64Pointer() != nil {
		data.MaxFilesize = models.ToPointer(int(plan.MaxFilesize.ValueInt64()))
	}

	if !plan.MimeWhitelist.IsNull() && !plan.MimeWhitelist.IsUnknown() {
		data.MimeWhitelist = mistutils.ListOfStringTerraformToSdk(plan.MimeWhitelist)
	} else {
		unset["-mime_whitelist"] = ""
	}

	data.Name = plan.Name.ValueString()

	var items []models.AvprofileProtocolsEnum
	for _, p := range plan.Protocols.Elements() {
		var sInterface interface{} = p
		s := sInterface.(basetypes.StringValue)
		items = append(items, (models.AvprofileProtocolsEnum)(s.ValueString()))
	}
	data.Protocols = items

	if !plan.UrlWhitelist.IsNull() && !plan.UrlWhitelist.IsUnknown() {
		data.UrlWhitelist = mistutils.ListOfStringTerraformToSdk(plan.UrlWhitelist)
	} else {
		unset["-url_whitelist"] = ""
	}

	return &data, diags
}
