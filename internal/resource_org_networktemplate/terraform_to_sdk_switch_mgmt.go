package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func switchMgmtProtectReCustomTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ProtectReCustom {
	var data []models.ProtectReCustom
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(CustomValue)

		data_item := models.ProtectReCustom{}
		if item_obj.PortRange.ValueStringPointer() != nil {
			data_item.PortRange = models.ToPointer(item_obj.PortRange.ValueString())
		}
		if item_obj.Protocol.ValueStringPointer() != nil {
			data_item.Protocol = models.ToPointer(models.ProtectReCustomProtocolEnum(item_obj.Protocol.ValueString()))
		}
		if !item_obj.Subnet.IsNull() && !item_obj.Subnet.IsUnknown() {
			data_item.Subnet = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Subnet)
		}

		data = append(data, data_item)
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
		var item_interface interface{} = item
		item_obj := item_interface.(ProtectReValue)

		if !item_obj.AllowedServices.IsNull() && !item_obj.AllowedServices.IsUnknown() {
			data.AllowedServices = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.AllowedServices)
		}
		if !item_obj.Custom.IsNull() && !item_obj.Custom.IsUnknown() {
			data.Custom = switchMgmtProtectReCustomTerraformToSdk(ctx, diags, item_obj.Custom)
		}
		if item_obj.Enabled.ValueBoolPointer() != nil {
			data.Enabled = models.ToPointer(item_obj.Enabled.ValueBool())
		}
		if !item_obj.TrustedHosts.IsNull() && !item_obj.TrustedHosts.IsUnknown() {
			data.TrustedHosts = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.TrustedHosts)
		}
		return &data
	}
}

func TacacsAcctServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.TacacsAcctServer {
	var data []models.TacacsAcctServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(TacacctServersValue)

		srv_data := models.TacacsAcctServer{}
		if srv_plan.Host.ValueStringPointer() != nil {
			srv_data.Host = srv_plan.Host.ValueStringPointer()
		}
		if srv_plan.Port.ValueStringPointer() != nil {
			srv_data.Port = srv_plan.Port.ValueStringPointer()
		}
		if srv_plan.Secret.ValueStringPointer() != nil {
			srv_data.Secret = srv_plan.Secret.ValueStringPointer()
		}
		if srv_plan.Timeout.ValueInt64Pointer() != nil {
			srv_data.Timeout = models.ToPointer(int(srv_plan.Timeout.ValueInt64()))
		}
		data = append(data, srv_data)
	}
	return data
}

func TacacsAuthServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.TacacsAuthServer {
	var data []models.TacacsAuthServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(TacplusServersValue)

		srv_data := models.TacacsAuthServer{}
		if srv_plan.Host.ValueStringPointer() != nil {
			srv_data.Host = srv_plan.Host.ValueStringPointer()
		}
		if srv_plan.Port.ValueStringPointer() != nil {
			srv_data.Port = srv_plan.Port.ValueStringPointer()
		}
		if srv_plan.Secret.ValueStringPointer() != nil {
			srv_data.Secret = srv_plan.Secret.ValueStringPointer()
		}
		if srv_plan.Timeout.ValueInt64Pointer() != nil {
			srv_data.Timeout = models.ToPointer(int(srv_plan.Timeout.ValueInt64()))
		}
		data = append(data, srv_data)
	}
	return data
}

func switchMgmtTacacsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.Tacacs {
	data := models.Tacacs{}

	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		item, e := NewProtectReValue(TacacsValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var item_interface interface{} = item
		item_obj := item_interface.(TacacsValue)

		if item_obj.Enabled.ValueBoolPointer() != nil {
			data.Enabled = models.ToPointer(item_obj.Enabled.ValueBool())
		}
		if item_obj.Network.ValueStringPointer() != nil {
			data.Network = models.ToPointer(item_obj.Network.ValueString())
		}
		if !item_obj.TacacctServers.IsNull() && !item_obj.TacacctServers.IsUnknown() {
			data.AcctServers = TacacsAcctServersTerraformToSdk(ctx, diags, item_obj.TacacctServers)
		}
		if !item_obj.TacplusServers.IsNull() && !item_obj.TacplusServers.IsUnknown() {
			data.TacplusServers = TacacsAuthServersTerraformToSdk(ctx, diags, item_obj.TacplusServers)
		}
		if item_obj.DefaultRole.ValueStringPointer() != nil {
			data.DefaultRole = models.ToPointer(models.TacacsDefaultRoleEnum(item_obj.DefaultRole.ValueString()))
		}

		return &data
	}
}

func switchLocalAccountUsersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.ConfigSwitchLocalAccountsUser {
	data := make(map[string]models.ConfigSwitchLocalAccountsUser)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(LocalAccountsValue)

		data_item := models.ConfigSwitchLocalAccountsUser{}
		if item_obj.Password.ValueStringPointer() != nil {
			data_item.Password = item_obj.Password.ValueStringPointer()
		}
		if item_obj.Role.ValueStringPointer() != nil {
			data_item.Role = (*models.ConfigSwitchLocalAccountsUserRoleEnum)(item_obj.Role.ValueStringPointer())
		}

		data[item_name] = data_item
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
		if !d.LocalAccounts.IsNull() && !d.LocalAccounts.IsUnknown() {
			data.LocalAccounts = switchLocalAccountUsersTerraformToSdk(ctx, diags, d.LocalAccounts)
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
