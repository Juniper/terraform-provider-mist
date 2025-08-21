package mist_utils

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func RequiredMistIdField(diags *diag.Diagnostics, value basetypes.StringValue) models.Optional[uuid.UUID] {
	var mistId models.Optional[uuid.UUID]
	mistId.ShouldSetValue(true)
	if !value.IsNull() && !value.IsUnknown() {
		uuid, e := uuid.Parse(value.ValueString())
		if e == nil {
			mistId.SetValue(models.ToPointer(uuid))
		} else {
			diags.AddError("Bad value for alarmtemplate_id", e.Error())
		}
	}
	return mistId
}
