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
		data.Ip = v_plan.Ip.ValueString()
		data.Secret = v_plan.Secret.ValueString()
		data.DisableEventTimestampCheck = v_plan.DisableEventTimestampCheck.ValueBoolPointer()
		data.Enabled = v_plan.Enabled.ValueBoolPointer()
		data.Port = models.ToPointer(int(v_plan.Port.ValueInt64()))

		data_list = append(data_list, data)
	}

	return data_list
}
