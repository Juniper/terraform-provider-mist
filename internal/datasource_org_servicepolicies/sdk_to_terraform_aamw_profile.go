package datasource_org_servicepolicies

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func aamwSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAamw) basetypes.ObjectValue {

	var aamwprofileId basetypes.StringValue
	var enabled basetypes.BoolValue
	var profile basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.AamwprofileId != nil {
		aamwprofileId = types.StringValue(d.AamwprofileId.String())
	}
	if d.Profile != nil {
		profile = types.StringValue(string(*d.Profile))
	}

	dataMapValue := map[string]attr.Value{
		"aamwprofile_id": aamwprofileId,
		"enabled":        enabled,
		"profile":        profile,
	}
	data, e := basetypes.NewObjectValue(AamwValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
