package datasource_org_servicepolicies

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func idpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpConfig) basetypes.ObjectValue {

	var alert_only basetypes.BoolValue
	var enabled basetypes.BoolValue
	var idpprofile_id basetypes.StringValue
	var profile basetypes.StringValue

	if d.AlertOnly != nil {
		alert_only = types.BoolValue(*d.AlertOnly)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.IdpprofileId != nil {
		idpprofile_id = types.StringValue(d.IdpprofileId.String())
	}
	if d.Profile != nil {
		profile = types.StringValue(*d.Profile)
	}

	data_map_attr_type := IdpValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"alert_only":    alert_only,
		"enabled":       enabled,
		"idpprofile_id": idpprofile_id,
		"profile":       profile,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
