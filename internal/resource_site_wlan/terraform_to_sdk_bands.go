package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bandsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []models.Dot11BandEnum {

	var data_list []models.Dot11BandEnum
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		data := models.Dot11BandEnum(string(v_plan.ValueString()))
		data_list = append(data_list, data)
	}

	return data_list
}
