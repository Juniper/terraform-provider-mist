package resource_org_nac_portal_template

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func getLogo(diags *diag.Diagnostics, filepath string) string {
	var filestring string
	file, err := os.Open(filepath)
	if err != nil {
		diags.AddError(
			"Invalid \"logo\" value for \"mist_org_nac_portal_template\" resource",
			fmt.Sprintf("Unable to open the file \"%s\": %s", filepath, err.Error()),
		)
		return filestring
	}

	defer file.Close()
	fileData, err := io.ReadAll(file)
	if err != nil {
		diags.AddError(
			"Invalid \"logo\" value for \"mist_org_nac_portal_template\" resource",
			fmt.Sprintf("Unable to read the file \"%s\": %s", filepath, err.Error()),
		)
		return filestring
	}

	contentType := http.DetectContentType(fileData)
	imgBase64Str := base64.StdEncoding.EncodeToString(fileData)
	filestring = fmt.Sprintf("data:%s;base64,%s", contentType, imgBase64Str)
	return filestring
}

func TerraformToSdk(plan *OrgNacPortalTemplateModel) (models.NacPortalTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.NacPortalTemplate{}

	if plan.Alignment.ValueStringPointer() != nil {
		alignment := models.PortalTemplateAlignmentEnum(plan.Alignment.ValueString())
		data.Alignment = &alignment
	}
	if plan.Color.ValueStringPointer() != nil {
		data.Color = plan.Color.ValueStringPointer()
	}
	if plan.Logo.ValueStringPointer() != nil {
		logo := getLogo(&diags, plan.Logo.ValueString())
		if logo != "" {
			data.Logo = &logo
		}
	}
	if plan.PoweredBy.ValueBoolPointer() != nil {
		data.PoweredBy = plan.PoweredBy.ValueBoolPointer()
	}

	return data, diags
}
