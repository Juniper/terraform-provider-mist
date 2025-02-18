package datasource_org_servicepolicies

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appQoeToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicyAppqoe) basetypes.ObjectValue {
	var enabled basetypes.BoolValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := basetypes.NewObjectValue(AppqoeValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func ewfSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.ServicePolicyEwfRule) basetypes.ListValue {
	var dataList []EwfValue
	for _, v := range d {
		var alertOnly basetypes.BoolValue
		var blockMessage basetypes.StringValue
		var enabled = types.BoolValue(false)
		var profile = types.StringValue("strict")

		if v.AlertOnly != nil {
			alertOnly = types.BoolValue(*v.AlertOnly)
		}
		if v.BlockMessage != nil {
			blockMessage = types.StringValue(*v.BlockMessage)
		}
		if v.Enabled != nil {
			enabled = types.BoolValue(*v.Enabled)
		}
		if v.Profile != nil {
			profile = types.StringValue(string(*v.Profile))
		}

		dataMapValue := map[string]attr.Value{
			"alert_only":    alertOnly,
			"block_message": blockMessage,
			"enabled":       enabled,
			"profile":       profile,
		}
		data, e := NewEwfValue(EwfValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	datalistType := EwfValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}
