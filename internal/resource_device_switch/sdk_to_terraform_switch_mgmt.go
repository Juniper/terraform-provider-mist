package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func switchMgmtProtecCustomtReSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ProtectReCustom) basetypes.ListValue {
	var dataList []CustomValue

	for _, d := range l {

		var portRange basetypes.StringValue
		var protocol basetypes.StringValue
		var subnets = misttransform.ListOfStringSdkToTerraformEmpty()

		if d.PortRange != nil {
			portRange = types.StringValue(*d.PortRange)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(string(*d.Protocol))
		}
		if d.Subnets != nil {
			subnets = misttransform.ListOfStringSdkToTerraform(d.Subnets)
		}

		dataMapValue := map[string]attr.Value{
			"port_range": portRange,
			"protocol":   protocol,
			"subnets":    subnets,
		}
		data, e := NewCustomValue(CustomValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := CustomValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)

	return r
}

func switchMgmtProtectReSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ProtectRe) basetypes.ObjectValue {

	var allowedServices = types.ListNull(types.StringType)
	var custom = basetypes.NewListValueMust(CustomValue{}.Type(ctx), []attr.Value{})
	var enabled basetypes.BoolValue
	var trustedHosts = types.ListNull(types.StringType)

	if d.AllowedServices != nil {
		var items []attr.Value
		var itemsType attr.Type = basetypes.StringType{}
		for _, item := range d.AllowedServices {
			items = append(items, types.StringValue(string(item)))
		}
		list, _ := types.ListValue(itemsType, items)
		allowedServices = list
	}
	if d.Custom != nil {
		custom = switchMgmtProtecCustomtReSdkToTerraform(ctx, diags, d.Custom)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TrustedHosts != nil {
		trustedHosts = misttransform.ListOfStringSdkToTerraform(d.TrustedHosts)
	}

	dataMapValue := map[string]attr.Value{
		"allowed_services": allowedServices,
		"custom":           custom,
		"enabled":          enabled,
		"trusted_hosts":    trustedHosts,
	}
	data, e := NewProtectReValue(ProtectReValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func switchMgmtTacacsAcctSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.TacacsAcctServer) basetypes.ListValue {

	var acctValueList []attr.Value
	for _, d := range l {
		var host basetypes.StringValue
		var port basetypes.StringValue
		var secret basetypes.StringValue
		var timeout basetypes.Int64Value

		if d.Host != nil {
			host = types.StringValue(*d.Host)
		}
		if d.Port != nil {
			port = types.StringValue(*d.Port)
		}
		if d.Secret != nil {
			secret = types.StringValue(*d.Secret)
		}
		if d.Timeout != nil {
			timeout = types.Int64Value(int64(*d.Timeout))
		}

		dataMapValue := map[string]attr.Value{
			"host":    host,
			"port":    port,
			"secret":  secret,
			"timeout": timeout,
		}
		data, e := NewTacacctServersValue(TacacctServersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		acctValueList = append(acctValueList, data)
	}

	acctStateListType := TacacctServersValue{}.Type(ctx)
	acctStateList, e := types.ListValueFrom(ctx, acctStateListType, acctValueList)
	diags.Append(e...)

	return acctStateList
}

func switchMgmtTacacsAuthSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.TacacsAuthServer) basetypes.ListValue {

	var acctValueList []attr.Value
	for _, d := range l {

		var host basetypes.StringValue
		var port basetypes.StringValue
		var secret basetypes.StringValue
		var timeout basetypes.Int64Value

		if d.Host != nil {
			host = types.StringValue(*d.Host)
		}
		if d.Port != nil {
			port = types.StringValue(*d.Port)
		}
		if d.Secret != nil {
			secret = types.StringValue(*d.Secret)
		}
		if d.Timeout != nil {
			timeout = types.Int64Value(int64(*d.Timeout))
		}

		dataMapValue := map[string]attr.Value{
			"host":    host,
			"port":    port,
			"secret":  secret,
			"timeout": timeout,
		}
		data, e := NewTacplusServersValue(TacplusServersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		acctValueList = append(acctValueList, data)
	}

	acctStateListType := TacplusServersValue{}.Type(ctx)
	acctStateList, e := types.ListValueFrom(ctx, acctStateListType, acctValueList)
	diags.Append(e...)

	return acctStateList
}

func switchMgmtTacacsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Tacacs) basetypes.ObjectValue {

	var defaultRole basetypes.StringValue
	var enabled basetypes.BoolValue
	var network basetypes.StringValue
	var acctServers = types.ListNull(TacacctServersValue{}.Type(ctx))
	var tacplusServers = types.ListNull(TacplusServersValue{}.Type(ctx))

	if d != nil {
		if d.DefaultRole != nil {
			defaultRole = types.StringValue(string(*d.DefaultRole))
		}
		if d.Enabled != nil {
			enabled = types.BoolValue(*d.Enabled)
		}
		if d.Network != nil {
			network = types.StringValue(*d.Network)
		}
		acctServers = switchMgmtTacacsAcctSdkToTerraform(ctx, diags, d.AcctServers)
		tacplusServers = switchMgmtTacacsAuthSdkToTerraform(ctx, diags, d.TacplusServers)

	}

	dataMapValue := map[string]attr.Value{
		"default_role":    defaultRole,
		"enabled":         enabled,
		"network":         network,
		"acct_servers":    acctServers,
		"tacplus_servers": tacplusServers,
	}
	data, e := NewTacacsValue(TacacsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func switchLocalAccountUserSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ConfigSwitchLocalAccountsUser) basetypes.MapValue {
	dataMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var password basetypes.StringValue
		var role basetypes.StringValue

		if d.Password != nil {
			password = types.StringValue(*d.Password)
		}
		if d.Role != nil {
			role = types.StringValue(string(*d.Role))
		}

		itemMapAttrType := LocalAccountsValue{}.AttributeTypes(ctx)
		itemMapValue := map[string]attr.Value{
			"password": password,
			"role":     role,
		}
		data, e := NewLocalAccountsValue(itemMapAttrType, itemMapValue)
		diags.Append(e...)

		dataMapValue[k] = data
	}
	stateType := LocalAccountsValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, dataMapValue)
	diags.Append(e...)
	return stateResult
}

func switchMgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchMgmt) SwitchMgmtValue {

	var apAffinityThreshold basetypes.Int64Value
	var cliBanner basetypes.StringValue
	var cliIdleTimeout basetypes.Int64Value
	var configRevertTimer basetypes.Int64Value
	var dhcpOptionFqdn basetypes.BoolValue
	var disableOobDownAlarm basetypes.BoolValue
	var fipsEnabled basetypes.BoolValue
	var localAccounts = types.MapNull(LocalAccountsValue{}.Type(ctx))
	var mxedgeProxyHost basetypes.StringValue
	var mxedgeProxyPort basetypes.Int64Value
	var protectRe = types.ObjectNull(ProtectReValue{}.AttributeTypes(ctx))
	var rootPassword basetypes.StringValue
	var tacacs = types.ObjectNull(TacacsValue{}.AttributeTypes(ctx))
	var useMxedgeProxy basetypes.BoolValue

	if d != nil {
		if d.ApAffinityThreshold != nil {
			apAffinityThreshold = types.Int64Value(int64(*d.ApAffinityThreshold))
		}
		if d.CliBanner != nil {
			cliBanner = types.StringValue(*d.CliBanner)
		}
		if d.ConfigRevertTimer != nil {
			configRevertTimer = types.Int64Value(int64(*d.ConfigRevertTimer))
		}
		if d.CliIdleTimeout != nil {
			cliIdleTimeout = types.Int64Value(int64(*d.CliIdleTimeout))
		}
		if d.DhcpOptionFqdn != nil {
			dhcpOptionFqdn = types.BoolValue(*d.DhcpOptionFqdn)
		}
		if d.DisableOobDownAlarm != nil {
			disableOobDownAlarm = types.BoolValue(*d.DisableOobDownAlarm)
		}
		if d.FipsEnabled != nil {
			fipsEnabled = types.BoolValue(*d.FipsEnabled)
		}
		if d.LocalAccounts != nil {
			localAccounts = switchLocalAccountUserSdkToTerraform(ctx, diags, d.LocalAccounts)
		}
		if d.MxedgeProxyHost != nil {
			mxedgeProxyHost = types.StringValue(*d.MxedgeProxyHost)
		}
		if d.MxedgeProxyPort != nil {
			mxedgeProxyPort = types.Int64Value(int64(*d.MxedgeProxyPort))
		}
		if d.ProtectRe != nil {
			protectRe = switchMgmtProtectReSdkToTerraform(ctx, diags, d.ProtectRe)
		}
		if d.RootPassword != nil {
			rootPassword = types.StringValue(*d.RootPassword)
		}
		if d.Tacacs != nil {
			tacacs = switchMgmtTacacsSdkToTerraform(ctx, diags, d.Tacacs)
		}
		if d.UseMxedgeProxy != nil {
			useMxedgeProxy = types.BoolValue(*d.UseMxedgeProxy)
		}
	}

	dataMapValue := map[string]attr.Value{
		"ap_affinity_threshold":  apAffinityThreshold,
		"cli_banner":             cliBanner,
		"cli_idle_timeout":       cliIdleTimeout,
		"config_revert_timer":    configRevertTimer,
		"dhcp_option_fqdn":       dhcpOptionFqdn,
		"disable_oob_down_alarm": disableOobDownAlarm,
		"fips_enabled":           fipsEnabled,
		"local_accounts":         localAccounts,
		"mxedge_proxy_host":      mxedgeProxyHost,
		"mxedge_proxy_port":      mxedgeProxyPort,
		"protect_re":             protectRe,
		"root_password":          rootPassword,
		"tacacs":                 tacacs,
		"use_mxedge_proxy":       useMxedgeProxy,
	}
	data, e := NewSwitchMgmtValue(SwitchMgmtValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
