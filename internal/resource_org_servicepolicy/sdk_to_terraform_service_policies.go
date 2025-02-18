package resource_org_servicepolicy

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appQoeToTerraform(d *models.ServicePolicyAppqoe) AppqoeValue {
	data := AppqoeValue{}
	if d != nil && d.Enabled != nil {
		data.Enabled = types.BoolValue(*d.Enabled)
	}
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
