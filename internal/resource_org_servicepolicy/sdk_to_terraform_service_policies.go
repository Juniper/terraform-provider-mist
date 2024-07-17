package resource_org_servicepolicy

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appQoeToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAppqoe) AppqoeValue {
	data := AppqoeValue{}
	if d != nil && d.Enabled != nil {
		data.Enabled = types.BoolValue(*d.Enabled)
	}
	return data
}

func ewfSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.ServicePolicyEwfRule) basetypes.ListValue {
	var data_list = []EwfValue{}
	for _, v := range d {
		var alert_only basetypes.BoolValue
		var block_message basetypes.StringValue
		var enabled basetypes.BoolValue = types.BoolValue(false)
		var profile basetypes.StringValue = types.StringValue("strict")

		if v.AlertOnly != nil {
			alert_only = types.BoolValue(*v.AlertOnly)
		}
		if v.BlockMessage != nil {
			block_message = types.StringValue(*v.BlockMessage)
		}
		if v.Enabled != nil {
			enabled = types.BoolValue(*v.Enabled)
		}
		if v.Profile != nil {
			profile = types.StringValue(string(*v.Profile))
		}

		data_map_attr_type := EwfValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"alert_only":    alert_only,
			"block_message": block_message,
			"enabled":       enabled,
			"profile":       profile,
		}
		data, e := NewEwfValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := EwfValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
