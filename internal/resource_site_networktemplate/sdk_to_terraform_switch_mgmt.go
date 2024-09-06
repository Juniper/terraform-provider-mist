package resource_site_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func switchMgmtProtecCustomtReSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ProtectReCustom) basetypes.ListValue {
	var data_list = []CustomValue{}

	for _, d := range l {

		var port_range basetypes.StringValue
		var protocol basetypes.StringValue
		var subnet basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

		if d.PortRange != nil {
			port_range = types.StringValue(*d.PortRange)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(string(*d.Protocol))
		}
		if d.Subnet != nil {
			subnet = mist_transform.ListOfStringSdkToTerraform(ctx, d.Subnet)
		}

		data_map_attr_type := CustomValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"port_range": port_range,
			"protocol":   protocol,
			"subnet":     subnet,
		}
		data, e := NewCustomValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := CustomValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)

	return r
}
func switchMgmtProtectReSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ProtectRe) basetypes.ObjectValue {
	var allowed_services basetypes.ListValue = types.ListNull(types.StringType)
	var custom basetypes.ListValue = types.ListNull(CustomValue{}.Type(ctx))
	var enabled basetypes.BoolValue
	var trusted_hosts basetypes.ListValue = types.ListNull(types.StringType)

	if d.AllowedServices != nil {
		allowed_services = mist_transform.ListOfStringSdkToTerraform(ctx, d.AllowedServices)
	}
	if d.Custom != nil {
		custom = switchMgmtProtecCustomtReSdkToTerraform(ctx, diags, d.Custom)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TrustedHosts != nil {
		trusted_hosts = mist_transform.ListOfStringSdkToTerraform(ctx, d.TrustedHosts)
	}

	data_map_attr_type := ProtectReValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allowed_services": allowed_services,
		"custom":           custom,
		"enabled":          enabled,
		"trusted_hosts":    trusted_hosts,
	}
	data, e := NewProtectReValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func switchMgmtTacacsAcctSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.TacacsAcctServer) basetypes.ListValue {
	var acct_value_list []attr.Value
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

		data_map_attr_type := TacacctServersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"host":    host,
			"port":    port,
			"secret":  secret,
			"timeout": timeout,
		}
		data, e := NewTacacctServersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		acct_value_list = append(acct_value_list, data)
	}

	acct_state_list_type := TacacctServersValue{}.Type(ctx)
	acct_state_list, e := types.ListValueFrom(ctx, acct_state_list_type, acct_value_list)
	diags.Append(e...)

	return acct_state_list
}

func switchMgmtTacacsAuthSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.TacacsAuthServer) basetypes.ListValue {
	var acct_value_list []attr.Value
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

		data_map_attr_type := TacplusServersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"host":    host,
			"port":    port,
			"secret":  secret,
			"timeout": timeout,
		}
		data, e := NewTacplusServersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		acct_value_list = append(acct_value_list, data)
	}

	acct_state_list_type := TacplusServersValue{}.Type(ctx)
	acct_state_list, e := types.ListValueFrom(ctx, acct_state_list_type, acct_value_list)
	diags.Append(e...)

	return acct_state_list
}

func switchMgmtTacacsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Tacacs) basetypes.ObjectValue {
	var default_role basetypes.StringValue
	var enabled basetypes.BoolValue
	var network basetypes.StringValue
	var acct_servers basetypes.ListValue = types.ListNull(TacacctServersValue{}.Type(ctx))
	var tacplus_servers basetypes.ListValue = types.ListNull(TacplusServersValue{}.Type(ctx))

	if d != nil {
		if d.DefaultRole != nil {
			default_role = types.StringValue(string(*d.DefaultRole))
		}
		if d.Enabled != nil {
			enabled = types.BoolValue(*d.Enabled)
		}
		if d.Network != nil {
			network = types.StringValue(*d.Network)
		}
		acct_servers = switchMgmtTacacsAcctSdkToTerraform(ctx, diags, d.AcctServers)
		tacplus_servers = switchMgmtTacacsAuthSdkToTerraform(ctx, diags, d.TacplusServers)

	}

	data_map_attr_type := TacacsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"default_role":    default_role,
		"enabled":         enabled,
		"network":         network,
		"acct_servers":    acct_servers,
		"tacplus_servers": tacplus_servers,
	}
	data, e := NewTacacsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func switchLocalAccountUserSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ConfigSwitchLocalAccountsUser) basetypes.MapValue {
	data_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var password basetypes.StringValue
		var role basetypes.StringValue

		if d.Password != nil {
			password = types.StringValue(*d.Password)
		}
		if d.Role != nil {
			role = types.StringValue(string(*d.Role))
		}

		data_map_attr_type := LocalAccountsValue{}.AttributeTypes(ctx)
		item_map_value := map[string]attr.Value{
			"password": password,
			"role":     role,
		}
		data, e := NewLocalAccountsValue(data_map_attr_type, item_map_value)
		diags.Append(e...)

		data_map_value[k] = data
	}
	state_type := LocalAccountsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, data_map_value)
	diags.Append(e...)
	return state_result
}

func switchMgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchMgmt) SwitchMgmtValue {

	var ap_affinity_threshold basetypes.Int64Value
	var cli_banner basetypes.StringValue
	var cli_idle_timeout basetypes.Int64Value
	var config_revert_timer basetypes.Int64Value
	var dhcp_option_fqdn basetypes.BoolValue
	var local_accounts basetypes.MapValue = types.MapNull(LocalAccountsValue{}.Type(ctx))
	var mxedge_proxy_host basetypes.StringValue
	var mxedge_proxy_port basetypes.Int64Value
	var protect_re basetypes.ObjectValue = types.ObjectNull(ProtectReValue{}.AttributeTypes(ctx))
	var root_password basetypes.StringValue
	var tacacs basetypes.ObjectValue = types.ObjectNull(TacacsValue{}.AttributeTypes(ctx))
	var use_mxedge_proxy basetypes.BoolValue

	if d != nil {
		if d.ApAffinityThreshold != nil {
			ap_affinity_threshold = types.Int64Value(int64(*d.ApAffinityThreshold))
		}
		if d.CliBanner != nil {
			cli_banner = types.StringValue(*d.CliBanner)
		}
		if d.ConfigRevertTimer != nil {
			config_revert_timer = types.Int64Value(int64(*d.ConfigRevertTimer))
		}
		if d.CliIdleTimeout != nil {
			cli_idle_timeout = types.Int64Value(int64(*d.CliIdleTimeout))
		}
		if d.DhcpOptionFqdn != nil {
			dhcp_option_fqdn = types.BoolValue(*d.DhcpOptionFqdn)
		}
		if d.LocalAccounts != nil {
			local_accounts = switchLocalAccountUserSdkToTerraform(ctx, diags, d.LocalAccounts)
		}
		if d.MxedgeProxyHost != nil {
			mxedge_proxy_host = types.StringValue(*d.MxedgeProxyHost)
		}
		if d.MxedgeProxyPort != nil {
			mxedge_proxy_port = types.Int64Value(int64(*d.MxedgeProxyPort))
		}
		if d.ProtectRe != nil {
			protect_re = switchMgmtProtectReSdkToTerraform(ctx, diags, d.ProtectRe)
		}
		if d.RootPassword != nil {
			root_password = types.StringValue(*d.RootPassword)
		}
		if d.Tacacs != nil {
			tacacs = switchMgmtTacacsSdkToTerraform(ctx, diags, d.Tacacs)
		}
		if d.UseMxedgeProxy != nil {
			use_mxedge_proxy = types.BoolValue(*d.UseMxedgeProxy)
		}
	}

	data_map_attr_type := SwitchMgmtValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"ap_affinity_threshold": ap_affinity_threshold,
		"cli_banner":            cli_banner,
		"cli_idle_timeout":      cli_idle_timeout,
		"config_revert_timer":   config_revert_timer,
		"dhcp_option_fqdn":      dhcp_option_fqdn,
		"local_accounts":        local_accounts,
		"mxedge_proxy_host":     mxedge_proxy_host,
		"mxedge_proxy_port":     mxedge_proxy_port,
		"protect_re":            protect_re,
		"root_password":         root_password,
		"tacacs":                tacacs,
		"use_mxedge_proxy":      use_mxedge_proxy,
	}
	data, e := NewSwitchMgmtValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
