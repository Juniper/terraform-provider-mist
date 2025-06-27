package datasource_org_servicepolicies

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func avSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAntivirus) basetypes.ObjectValue {

	var avprofileId basetypes.StringValue
	var enabled basetypes.BoolValue
	var profile basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.AvprofileId != nil {
		avprofileId = types.StringValue(d.AvprofileId.String())
	}
	if d.Profile != nil && *d.Profile != avprofileId.ValueString() {
		profile = types.StringValue(*d.Profile)
	}

	dataMapValue := map[string]attr.Value{
		"avprofile_id": avprofileId,
		"enabled":      enabled,
		"profile":      profile,
	}
	data, e := basetypes.NewObjectValue(AntivirusValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
