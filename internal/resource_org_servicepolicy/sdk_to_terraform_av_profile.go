package resource_org_servicepolicy

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func avSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAntivirus) AntivirusValue {

	var avprofileId basetypes.StringValue
	var enabled basetypes.BoolValue
	var profile basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.AvprofileId != nil {
		avprofileId = types.StringValue(d.AvprofileId.String())
	}
	if d.Profile != nil && *d.Profile != d.AvprofileId.String() {
		profile = types.StringValue(*d.Profile)
	}

	dataMapValue := map[string]attr.Value{
		"avpprofile_id": avprofileId,
		"enabled":       enabled,
		"profile":       profile,
	}
	data, e := NewAntivirusValue(AntivirusValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
