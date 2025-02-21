package resource_org_servicepolicy

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func idpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpConfig) IdpValue {

	var alertOnly basetypes.BoolValue
	var enabled basetypes.BoolValue
	var idpprofileId basetypes.StringValue
	var profile basetypes.StringValue

	if d.AlertOnly != nil {
		alertOnly = types.BoolValue(*d.AlertOnly)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.IdpprofileId != nil {
		idpprofileId = types.StringValue(d.IdpprofileId.String())
	}
	if d.Profile != nil {
		profile = types.StringValue(*d.Profile)
	}

	dataMapValue := map[string]attr.Value{
		"alert_only":    alertOnly,
		"enabled":       enabled,
		"idpprofile_id": idpprofileId,
		"profile":       profile,
	}
	data, e := NewIdpValue(IdpValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
