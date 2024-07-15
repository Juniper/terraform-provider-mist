package resource_org_wxtag

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func specsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []models.WxlanTagSpec {
	var data_list []models.WxlanTagSpec
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		p := v_interface.(SpecsValue)

		data := models.WxlanTagSpec{
			PortRange: models.ToPointer(string(p.PortRange.ValueString())),
			Protocol:  models.ToPointer(string(p.Protocol.ValueString())),
			Subnets:   mist_transform.ListOfStringTerraformToSdk(ctx, p.Subnets),
		}

		data_list = append(data_list, data)
	}
	return data_list

}
