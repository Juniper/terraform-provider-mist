package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func switchMgmtProtectReCustomTerraformToSdk(d basetypes.ListValue) []models.ProtectReCustom {
	var data []models.ProtectReCustom
	for _, item := range d.Elements() {
		var itemInterface interface{} = item
		itemObj := itemInterface.(CustomValue)

		dataItem := models.ProtectReCustom{}
		if itemObj.PortRange.ValueStringPointer() != nil {
			dataItem.PortRange = models.ToPointer(itemObj.PortRange.ValueString())
		}
		if itemObj.Protocol.ValueStringPointer() != nil {
			dataItem.Protocol = models.ToPointer(models.ProtectReCustomProtocolEnum(itemObj.Protocol.ValueString()))
		}
		if !itemObj.Subnets.IsNull() && !itemObj.Subnets.IsUnknown() {
			dataItem.Subnets = mistutils.ListOfStringTerraformToSdk(itemObj.Subnets)
		}

		data = append(data, dataItem)
	}
	return data
}

func switchMgmtProtectReTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ProtectRe {
	data := models.ProtectRe{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		item, e := NewProtectReValue(ProtectReValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var itemInterface interface{} = item
		itemObj := itemInterface.(ProtectReValue)

		if !itemObj.AllowedServices.IsNull() && !itemObj.AllowedServices.IsUnknown() {
			var items []models.ProtectReAllowedServiceEnum
			for _, i := range itemObj.AllowedServices.Elements() {
				var iface interface{} = i
				val := iface.(basetypes.StringValue)
				items = append(items, models.ProtectReAllowedServiceEnum(val.ValueString()))
			}
			data.AllowedServices = items
		}
		if !itemObj.Custom.IsNull() && !itemObj.Custom.IsUnknown() {
			data.Custom = switchMgmtProtectReCustomTerraformToSdk(itemObj.Custom)
		}
		if itemObj.Enabled.ValueBoolPointer() != nil {
			data.Enabled = models.ToPointer(itemObj.Enabled.ValueBool())
		}
		if !itemObj.TrustedHosts.IsNull() && !itemObj.TrustedHosts.IsUnknown() {
			data.TrustedHosts = mistutils.ListOfStringTerraformToSdk(itemObj.TrustedHosts)
		}
		return &data
	}
}

func tacacsAcctServersTerraformToSdk(d basetypes.ListValue) []models.TacacsAcctServer {

	var data []models.TacacsAcctServer
	for _, planAttr := range d.Elements() {
		var srvPlanInterface interface{} = planAttr
		srvPlan := srvPlanInterface.(TacacctServersValue)

		srvData := models.TacacsAcctServer{}
		if srvPlan.Host.ValueStringPointer() != nil {
			srvData.Host = srvPlan.Host.ValueStringPointer()
		}
		if srvPlan.Port.ValueStringPointer() != nil {
			srvData.Port = srvPlan.Port.ValueStringPointer()
		}
		if srvPlan.Secret.ValueStringPointer() != nil {
			srvData.Secret = srvPlan.Secret.ValueStringPointer()
		}
		if srvPlan.Timeout.ValueInt64Pointer() != nil {
			srvData.Timeout = models.ToPointer(int(srvPlan.Timeout.ValueInt64()))
		}
		data = append(data, srvData)
	}
	return data
}
func tacacsAuthServersTerraformToSdk(d basetypes.ListValue) []models.TacacsAuthServer {

	var data []models.TacacsAuthServer
	for _, planAttr := range d.Elements() {
		var srvPlanInterface interface{} = planAttr
		srvPlan := srvPlanInterface.(TacplusServersValue)

		srvData := models.TacacsAuthServer{}
		if srvPlan.Host.ValueStringPointer() != nil {
			srvData.Host = srvPlan.Host.ValueStringPointer()
		}
		if srvPlan.Port.ValueStringPointer() != nil {
			srvData.Port = srvPlan.Port.ValueStringPointer()
		}
		if srvPlan.Secret.ValueStringPointer() != nil {
			srvData.Secret = srvPlan.Secret.ValueStringPointer()
		}
		if srvPlan.Timeout.ValueInt64Pointer() != nil {
			srvData.Timeout = models.ToPointer(int(srvPlan.Timeout.ValueInt64()))
		}
		data = append(data, srvData)
	}
	return data
}

func switchMgmtTacacsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.Tacacs {

	data := models.Tacacs{}

	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		item, e := NewTacacsValue(d.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)

		if e != nil {
			diags.Append(e...)
		} else {
			if item.Enabled.ValueBoolPointer() != nil {
				data.Enabled = models.ToPointer(item.Enabled.ValueBool())
			}
			if item.Network.ValueStringPointer() != nil {
				data.Network = models.ToPointer(item.Network.ValueString())
			}
			if !item.TacacctServers.IsNull() && !item.TacacctServers.IsUnknown() {
				data.AcctServers = tacacsAcctServersTerraformToSdk(item.TacacctServers)
			}
			if !item.TacplusServers.IsNull() && !item.TacplusServers.IsUnknown() {
				data.TacplusServers = tacacsAuthServersTerraformToSdk(item.TacplusServers)
			}
			if item.DefaultRole.ValueStringPointer() != nil {
				data.DefaultRole = models.ToPointer(models.TacacsDefaultRoleEnum(item.DefaultRole.ValueString()))
			}
		}
		return &data
	}
}

func switchLocalAccountUsersTerraformToSdk(d basetypes.MapValue) map[string]models.ConfigSwitchLocalAccountsUser {
	data := make(map[string]models.ConfigSwitchLocalAccountsUser)
	for itemName, itemValue := range d.Elements() {
		var itemInterface interface{} = itemValue
		itemObj := itemInterface.(LocalAccountsValue)

		dataItem := models.ConfigSwitchLocalAccountsUser{}
		if itemObj.Password.ValueStringPointer() != nil {
			dataItem.Password = itemObj.Password.ValueStringPointer()
		}
		if itemObj.Role.ValueStringPointer() != nil {
			dataItem.Role = (*models.ConfigSwitchLocalAccountsUserRoleEnum)(itemObj.Role.ValueStringPointer())
		}

		data[itemName] = dataItem
	}
	return data
}

func switchMgmtTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SwitchMgmtValue) *models.SwitchMgmt {

	data := models.SwitchMgmt{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {

		if d.ApAffinityThreshold.ValueInt64Pointer() != nil {
			data.ApAffinityThreshold = models.ToPointer(int(d.ApAffinityThreshold.ValueInt64()))
		}
		if d.CliBanner.ValueStringPointer() != nil {
			data.CliBanner = d.CliBanner.ValueStringPointer()
		}
		if d.CliIdleTimeout.ValueInt64Pointer() != nil {
			data.CliIdleTimeout = models.ToPointer(int(d.CliIdleTimeout.ValueInt64()))
		}
		if d.ConfigRevertTimer.ValueInt64Pointer() != nil {
			data.ConfigRevertTimer = models.ToPointer(int(d.ConfigRevertTimer.ValueInt64()))
		}
		if d.DhcpOptionFqdn.ValueBoolPointer() != nil {
			data.DhcpOptionFqdn = d.DhcpOptionFqdn.ValueBoolPointer()
		}
		if d.DisableOobDownAlarm.ValueBoolPointer() != nil {
			data.DisableOobDownAlarm = d.DisableOobDownAlarm.ValueBoolPointer()
		}
		if d.FipsEnabled.ValueBoolPointer() != nil {
			data.FipsEnabled = d.FipsEnabled.ValueBoolPointer()
		}
		if !d.LocalAccounts.IsNull() && !d.LocalAccounts.IsUnknown() {
			data.LocalAccounts = switchLocalAccountUsersTerraformToSdk(d.LocalAccounts)
		}
		if d.MxedgeProxyHost.ValueStringPointer() != nil {
			data.MxedgeProxyHost = d.MxedgeProxyHost.ValueStringPointer()
		}
		if d.MxedgeProxyPort.ValueInt64Pointer() != nil {
			data.MxedgeProxyPort = models.ToPointer(int(d.MxedgeProxyPort.ValueInt64()))
		}
		if !d.ProtectRe.IsNull() && !d.ProtectRe.IsUnknown() {
			data.ProtectRe = switchMgmtProtectReTerraformToSdk(ctx, diags, d.ProtectRe)
		}
		if d.RootPassword.ValueStringPointer() != nil {
			data.RootPassword = models.ToPointer(d.RootPassword.ValueString())
		}
		if !d.Tacacs.IsNull() && !d.Tacacs.IsUnknown() {
			data.Tacacs = switchMgmtTacacsTerraformToSdk(ctx, diags, d.Tacacs)
		}
		if d.UseMxedgeProxy.ValueBoolPointer() != nil {
			data.UseMxedgeProxy = d.UseMxedgeProxy.ValueBoolPointer()
		}

		return &data
	}

}
