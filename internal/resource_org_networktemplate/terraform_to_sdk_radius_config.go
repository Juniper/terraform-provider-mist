package resource_org_networktemplate

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
		srvData.Port = models.ToPointer(int(srvPlan.Port.ValueInt64()))
		srvData.Secret = srvPlan.Secret.ValueString()
		if srvPlan.KeywrapEnabled.ValueBoolPointer() != nil {
			srvData.KeywrapEnabled = models.ToPointer(srvPlan.KeywrapEnabled.ValueBool())
		}
		if srvPlan.KeywrapFormat.ValueStringPointer() != nil {
			srvData.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(srvPlan.KeywrapFormat.ValueString()))
		}
		if srvPlan.KeywrapKek.ValueStringPointer() != nil {
			srvData.KeywrapKek = models.ToPointer(srvPlan.KeywrapKek.ValueString())
		}
		if srvPlan.KeywrapMack.ValueStringPointer() != nil {
			srvData.KeywrapMack = models.ToPointer(srvPlan.KeywrapMack.ValueString())
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
		srvData.Port = models.ToPointer(int(srvPlan.Port.ValueInt64()))
		srvData.Secret = srvPlan.Secret.ValueString()
		if srvPlan.KeywrapEnabled.ValueBoolPointer() != nil {
			srvData.KeywrapEnabled = models.ToPointer(srvPlan.KeywrapEnabled.ValueBool())
		}
		if srvPlan.KeywrapFormat.ValueStringPointer() != nil {
			srvData.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(srvPlan.KeywrapFormat.ValueString()))
		}
		if srvPlan.KeywrapKek.ValueStringPointer() != nil {
			srvData.KeywrapKek = models.ToPointer(srvPlan.KeywrapKek.ValueString())
		}
		if srvPlan.KeywrapMack.ValueStringPointer() != nil {
			srvData.KeywrapMack = models.ToPointer(srvPlan.KeywrapMack.ValueString())
		}
		if srvPlan.RequireMessageAuthenticator.ValueBoolPointer() != nil {
			srvData.RequireMessageAuthenticator = models.ToPointer(srvPlan.RequireMessageAuthenticator.ValueBool())
		}
		data = append(data, srvData)
	}
	return data
}

func radiusConfigTerraformToSdk(d RadiusConfigValue) *models.SwitchRadiusConfig {

	data := models.SwitchRadiusConfig{}
	if d.AcctInterimInterval.ValueInt64Pointer() != nil {
		data.AcctInterimInterval = models.ToPointer(int(d.AcctInterimInterval.ValueInt64()))
	}
	if !d.AcctServers.IsNull() && !d.AcctServers.IsUnknown() {
		data.AcctServers = radiusAcctServersTerraformToSdk(d.AcctServers)
	}
	if d.AuthServersRetries.ValueInt64Pointer() != nil {
		data.AuthServersRetries = models.ToPointer(int(d.AuthServersRetries.ValueInt64()))
	}
	if d.AuthServersTimeout.ValueInt64Pointer() != nil {
		data.AuthServersTimeout = models.ToPointer(int(d.AuthServersTimeout.ValueInt64()))
	}
	if d.Network.ValueStringPointer() != nil {
		data.Network = models.ToPointer(d.Network.ValueString())
	}
	if d.SourceIp.ValueStringPointer() != nil {
		data.SourceIp = models.ToPointer(d.SourceIp.ValueString())
	}
	if !d.AuthServers.IsNull() && !d.AuthServers.IsUnknown() {
		data.AuthServers = radiusAuthServersTerraformToSdk(d.AuthServers)
	}

	return &data
}
