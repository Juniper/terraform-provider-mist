package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func coaServerTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []models.CoaServer {

	var data_list []models.CoaServer
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(CoaServersValue)
		data := models.CoaServer{}
		if v_plan.Ip.ValueStringPointer() != nil {
			data.Ip = v_plan.Ip.ValueString()
		}
		if v_plan.Secret.ValueStringPointer() != nil {
			data.Secret = v_plan.Secret.ValueString()
		}
		if v_plan.DisableEventTimestampCheck.ValueBoolPointer() != nil {
			data.DisableEventTimestampCheck = v_plan.DisableEventTimestampCheck.ValueBoolPointer()
		}
		if v_plan.Enabled.ValueBoolPointer() != nil {
			data.Enabled = v_plan.Enabled.ValueBoolPointer()
		}
		if v_plan.Port.ValueInt64Pointer() != nil {
			data.Port = models.ToPointer(int(v_plan.Port.ValueInt64()))
		}
		data_list = append(data_list, data)
	}

	return data_list
}
