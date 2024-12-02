package resource_org_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func radiusAcctServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RadiusAcctServer {

	var data []models.RadiusAcctServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(AcctServersValue)
		srv_data := models.RadiusAcctServer{}
		srv_data.Host = srv_plan.Host.ValueString()
		if srv_plan.Port.ValueInt64Pointer() != nil {
			srv_data.Port = models.ToPointer(int(srv_plan.Port.ValueInt64()))
		}
		srv_data.Secret = srv_plan.Secret.ValueString()
		if srv_plan.KeywrapEnabled.ValueBoolPointer() != nil {
			srv_data.KeywrapEnabled = srv_plan.KeywrapEnabled.ValueBoolPointer()
		}
		if srv_plan.KeywrapFormat.ValueStringPointer() != nil {
			srv_data.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(string(srv_plan.KeywrapFormat.ValueString())))
		}
		if srv_plan.KeywrapKek.ValueStringPointer() != nil {
			srv_data.KeywrapKek = srv_plan.KeywrapKek.ValueStringPointer()
		}
		if srv_plan.KeywrapMack.ValueStringPointer() != nil {
			srv_data.KeywrapMack = srv_plan.KeywrapMack.ValueStringPointer()
		}
		data = append(data, srv_data)
	}
	return data
}

func radiusAuthServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RadiusAuthServer {

	var data []models.RadiusAuthServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(AuthServersValue)
		srv_data := models.RadiusAuthServer{}
		srv_data.Host = srv_plan.Host.ValueString()
		if srv_plan.Port.ValueInt64Pointer() != nil {
			srv_data.Port = models.ToPointer(int(srv_plan.Port.ValueInt64()))
		}
		srv_data.Secret = srv_plan.Secret.ValueString()
		if srv_plan.RequireMessageAuthenticator.ValueBoolPointer() != nil {
			srv_data.RequireMessageAuthenticator = srv_plan.RequireMessageAuthenticator.ValueBoolPointer()
		}
		if srv_plan.KeywrapEnabled.ValueBoolPointer() != nil {
			srv_data.KeywrapEnabled = srv_plan.KeywrapEnabled.ValueBoolPointer()
		}
		if srv_plan.KeywrapFormat.ValueStringPointer() != nil {
			srv_data.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(string(srv_plan.KeywrapFormat.ValueString())))
		}
		if srv_plan.KeywrapKek.ValueStringPointer() != nil {
			srv_data.KeywrapKek = srv_plan.KeywrapKek.ValueStringPointer()
		}
		if srv_plan.KeywrapMack.ValueStringPointer() != nil {
			srv_data.KeywrapMack = srv_plan.KeywrapMack.ValueStringPointer()
		}
		data = append(data, srv_data)
	}
	return data
}
