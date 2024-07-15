package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ipConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayIpConfigProperty {
	tflog.Debug(ctx, "ipConfigsTerraformToSdk")
	data_map := make(map[string]models.GatewayIpConfigProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IpConfigsValue)

		data := models.GatewayIpConfigProperty{}
		if plan.Ip.ValueStringPointer() != nil {
			data.Ip = models.ToPointer(plan.Ip.ValueString())
		}
		if plan.Netmask.ValueStringPointer() != nil {
			data.Netmask = models.ToPointer(plan.Netmask.ValueString())
		}
		if !plan.SecondaryIps.IsNull() && !plan.SecondaryIps.IsUnknown() {
			data.SecondaryIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.SecondaryIps)
		}
		if !plan.IpConfigsType.IsNull() && !plan.IpConfigsType.IsUnknown() {
			data.Type = models.ToPointer(models.IpTypeEnum(plan.IpConfigsType.ValueString()))
		}
		data_map[k] = data
	}
	return data_map
}
