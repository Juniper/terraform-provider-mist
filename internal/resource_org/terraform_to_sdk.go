package resource_org

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgModel) (*models.Org, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Org{}
	unset := make(map[string]interface{})

	if plan.AlarmtemplateId.IsNull() || plan.AlarmtemplateId.IsUnknown() {
		unset["-alarmtemplate_id"] = ""
	} else {
		data.AlarmtemplateId = models.NewOptional(models.ToPointer(uuid.MustParse(plan.AlarmtemplateId.ValueString())))
	}

	if plan.AllowMist.IsNull() || plan.AllowMist.IsUnknown() {
		unset["-allow_mist"] = ""
	} else {
		data.AllowMist = plan.AllowMist.ValueBoolPointer()
	}

	data.Name = plan.Name.ValueString()

	if plan.SessionExpiry.IsNull() || plan.SessionExpiry.IsUnknown() {
		unset["-session_expiry"] = ""
	} else {
		data.SessionExpiry = models.ToPointer(int(plan.SessionExpiry.ValueInt64()))
	}

	return &data, diags
}
