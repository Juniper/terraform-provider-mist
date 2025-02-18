package resource_site_wlan

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func radiusAcctServersTerraformToSdk(d basetypes.ListValue) []models.RadiusAcctServer {

	var data []models.RadiusAcctServer
	for _, planAttr := range d.Elements() {
		var srvPlanInterface interface{} = planAttr
		srvPlan := srvPlanInterface.(AcctServersValue)
		srvData := models.RadiusAcctServer{}
		srvData.Host = srvPlan.Host.ValueString()
		if srvPlan.Port.ValueInt64Pointer() != nil {
			srvData.Port = models.ToPointer(int(srvPlan.Port.ValueInt64()))
		}
		srvData.Secret = srvPlan.Secret.ValueString()
		if srvPlan.KeywrapEnabled.ValueBoolPointer() != nil {
			srvData.KeywrapEnabled = srvPlan.KeywrapEnabled.ValueBoolPointer()
		}
		if srvPlan.KeywrapFormat.ValueStringPointer() != nil {
			srvData.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(srvPlan.KeywrapFormat.ValueString()))
		}
		if srvPlan.KeywrapKek.ValueStringPointer() != nil {
			srvData.KeywrapKek = srvPlan.KeywrapKek.ValueStringPointer()
		}
		if srvPlan.KeywrapMack.ValueStringPointer() != nil {
			srvData.KeywrapMack = srvPlan.KeywrapMack.ValueStringPointer()
		}
		data = append(data, srvData)
	}
	return data
}

func radiusAuthServersTerraformToSdk(d basetypes.ListValue) []models.RadiusAuthServer {

	var data []models.RadiusAuthServer
	for _, planAttr := range d.Elements() {
		var srvPlanInterface interface{} = planAttr
		srvPlan := srvPlanInterface.(AuthServersValue)
		srvData := models.RadiusAuthServer{}
		srvData.Host = srvPlan.Host.ValueString()
		if srvPlan.Port.ValueInt64Pointer() != nil {
			srvData.Port = models.ToPointer(int(srvPlan.Port.ValueInt64()))
		}
		srvData.Secret = srvPlan.Secret.ValueString()
		if srvPlan.RequireMessageAuthenticator.ValueBoolPointer() != nil {
			srvData.RequireMessageAuthenticator = srvPlan.RequireMessageAuthenticator.ValueBoolPointer()
		}
		if srvPlan.KeywrapEnabled.ValueBoolPointer() != nil {
			srvData.KeywrapEnabled = srvPlan.KeywrapEnabled.ValueBoolPointer()
		}
		if srvPlan.KeywrapFormat.ValueStringPointer() != nil {
			srvData.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(srvPlan.KeywrapFormat.ValueString()))
		}
		if srvPlan.KeywrapKek.ValueStringPointer() != nil {
			srvData.KeywrapKek = srvPlan.KeywrapKek.ValueStringPointer()
		}
		if srvPlan.KeywrapMack.ValueStringPointer() != nil {
			srvData.KeywrapMack = srvPlan.KeywrapMack.ValueStringPointer()
		}
		data = append(data, srvData)
	}
	return data
}
