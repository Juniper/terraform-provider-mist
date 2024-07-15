package resource_site_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func switchMgmtProtecCustomtReSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ProtectReCustom) basetypes.ListValue {
	tflog.Debug(ctx, "switchMgmtProtecCustomtReSdkToTerraform")
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
	tflog.Debug(ctx, "switchMgmtProtectReSdkToTerraform")

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
	tflog.Debug(ctx, "switchMgmtTacacsAcctSdkToTerraform")

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
	tflog.Debug(ctx, "switchMgmtTacacsAuthSdkToTerraform")

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
	tflog.Debug(ctx, "switchMgmtTacacsSdkToTerraform")

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
func switchMgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchMgmt) SwitchMgmtValue {
	tflog.Debug(ctx, "switchMgmtSdkToTerraform")

	var config_revert basetypes.Int64Value
	var protect_re basetypes.ObjectValue = types.ObjectNull(ProtectReValue{}.AttributeTypes(ctx))
	var root_password basetypes.StringValue
	var tacacs basetypes.ObjectValue = types.ObjectNull(TacacsValue{}.AttributeTypes(ctx))

	if d != nil {
		if d.ConfigRevert != nil {
			config_revert = types.Int64Value(int64(*d.ConfigRevert))
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
	}

	data_map_attr_type := SwitchMgmtValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"config_revert": config_revert,
		"protect_re":    protect_re,
		"root_password": root_password,
		"tacacs":        tacacs,
	}
	data, e := NewSwitchMgmtValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
