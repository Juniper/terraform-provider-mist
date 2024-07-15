package resource_device_gateway

import (
	"context"
	"encoding/json"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func dhcpdConfigVendorEncapsulatedSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.DhcpdConfigVendorOption) basetypes.MapValue {
	tflog.Debug(ctx, "dhcpdConfigVendorEncapsulatedSdkToTerraform")

	r_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var type_option basetypes.StringValue
		var value basetypes.StringValue

		if d.Type != nil {
			type_option = types.StringValue(string(*d.Type))
		}
		if d.Value != nil {
			value = types.StringValue(*d.Value)
		}

		data_map_attr_type := VendorEncapulatedValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"type":  type_option,
			"value": value,
		}
		data, e := NewVendorEncapulatedValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		r_map_value[k] = data
	}
	state_result_map_type := VendorEncapulatedValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, r_map_value)
	diags.Append(e...)
	return state_result_map
}

func dhcpdConfigOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.DhcpdConfigOption) basetypes.MapValue {
	tflog.Debug(ctx, "dhcpdConfigOptionsSdkToTerraform")

	r_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var type_option basetypes.StringValue
		var value basetypes.StringValue

		if d.Type != nil {
			type_option = types.StringValue(string(*d.Type))
		}
		if d.Value != nil {
			value = types.StringValue(*d.Value)
		}

		data_map_attr_type := OptionsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"type":  type_option,
			"value": value,
		}
		data, e := NewOptionsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		r_map_value[k] = data
	}
	state_result_map_type := OptionsValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, r_map_value)
	diags.Append(e...)
	return state_result_map
}

func dhcpdConfigFixedBindingsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.DhcpdConfigFixedBinding) basetypes.MapValue {
	tflog.Debug(ctx, "dhcpdConfigFixedBindingsSdkToTerraform")
	r_map := make(map[string]attr.Value)
	for k, d := range m {
		var ip basetypes.StringValue
		var name basetypes.StringValue

		if d.Ip != nil {
			ip = types.StringValue(*d.Ip)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}

		data_map_attr_type := FixedBindingsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"ip":   ip,
			"name": name,
		}
		data, e := NewFixedBindingsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		r_map[k] = data
	}
	state_type := FixedBindingsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, r_map)
	diags.Append(e...)
	return state_result
}

func dhcpdConfigConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]interface{}) basetypes.MapValue {
	tflog.Debug(ctx, "dhcpdConfigConfigsSdkToTerraform")
	r_map_value := make(map[string]attr.Value)
	for k, d_interface := range m {
		if k != "enabled" {
			d_bytes, _ := json.Marshal(d_interface)
			d := models.DhcpdConfigProperty{}
			d.UnmarshalJSON(d_bytes)
			var dns_servers basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
			var dns_suffix basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
			var fixed_bindings basetypes.MapValue = types.MapNull(FixedBindingsValue{}.Type(ctx))
			var gateway basetypes.StringValue
			var ip_end basetypes.StringValue
			var ip_end6 basetypes.StringValue
			var ip_start basetypes.StringValue
			var ip_start6 basetypes.StringValue
			var lease_time basetypes.Int64Value = types.Int64Value(86400)
			var options basetypes.MapValue = types.MapNull(OptionsValue{}.Type(ctx))
			var server_id_override basetypes.BoolValue = types.BoolValue(false)
			var servers basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
			var servers6 basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
			var type4 basetypes.StringValue = types.StringValue("local")
			var type6 basetypes.StringValue = types.StringValue("none")
			var vendor_encapulated basetypes.MapValue = types.MapNull(VendorEncapulatedValue{}.Type(ctx))

			if d.DnsServers != nil {
				dns_servers = mist_transform.ListOfStringSdkToTerraform(ctx, d.DnsServers)
			}
			if d.DnsSuffix != nil {
				dns_suffix = mist_transform.ListOfStringSdkToTerraform(ctx, d.DnsSuffix)
			}
			if d.FixedBindings != nil && len(d.FixedBindings) > 0 {
				fixed_bindings = dhcpdConfigFixedBindingsSdkToTerraform(ctx, diags, d.FixedBindings)
			}
			if d.Gateway != nil {
				gateway = types.StringValue(*d.Gateway)
			}
			if d.IpEnd != nil {
				ip_end = types.StringValue(*d.IpEnd)
			}
			if d.IpEnd6 != nil {
				ip_end6 = types.StringValue(*d.IpEnd6)
			}
			if d.IpStart != nil {
				ip_start = types.StringValue(*d.IpStart)
			}
			if d.IpStart6 != nil {
				ip_start6 = types.StringValue(*d.IpStart6)
			}
			if d.LeaseTime != nil {
				lease_time = types.Int64Value(int64(*d.LeaseTime))
			}
			if d.Options != nil {
				options = dhcpdConfigOptionsSdkToTerraform(ctx, diags, d.Options)
			}
			if d.ServerIdOverride != nil {
				server_id_override = types.BoolValue(*d.ServerIdOverride)
			}
			if d.Servers != nil {
				servers = mist_transform.ListOfStringSdkToTerraform(ctx, d.Servers)
			}
			if d.Servers6 != nil {
				servers6 = mist_transform.ListOfStringSdkToTerraform(ctx, d.Servers6)
			}
			if d.Type != nil {
				type4 = types.StringValue(string(*d.Type))
			}
			if d.Type6 != nil {
				type6 = types.StringValue(string(*d.Type6))
			}
			if d.VendorEncapulated != nil {
				vendor_encapulated = dhcpdConfigVendorEncapsulatedSdkToTerraform(ctx, diags, d.VendorEncapulated)
			}

			data_map_attr_type := ConfigValue{}.AttributeTypes(ctx)
			data_map_value := map[string]attr.Value{
				"dns_servers":        dns_servers,
				"dns_suffix":         dns_suffix,
				"fixed_bindings":     fixed_bindings,
				"gateway":            gateway,
				"ip_end":             ip_end,
				"ip_end6":            ip_end6,
				"ip_start":           ip_start,
				"ip_start6":          ip_start6,
				"lease_time":         lease_time,
				"options":            options,
				"server_id_override": server_id_override,
				"servers":            servers,
				"servers6":           servers6,
				"type":               type4,
				"type6":              type6,
				"vendor_encapulated": vendor_encapulated,
			}
			data, e := NewConfigValue(data_map_attr_type, data_map_value)
			diags.Append(e...)

			r_map_value[k] = data
		}
	}
	state_type := ConfigValue{}.Type(ctx)
	r, e := types.MapValueFrom(ctx, state_type, r_map_value)
	diags.Append(e...)
	return r
}

func dhcpdConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.DhcpdConfig) DhcpdConfigValue {
	tflog.Debug(ctx, "dhcpdConfigSdkToTerraform")

	var config basetypes.MapValue = types.MapNull(ConfigValue{}.Type(ctx))
	var enabled basetypes.BoolValue = types.BoolValue(false)

	if len(d.AdditionalProperties) > 0 {
		config = dhcpdConfigConfigsSdkToTerraform(ctx, diags, d.AdditionalProperties)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := DhcpdConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"config":  config,
		"enabled": enabled,
	}
	data, e := NewDhcpdConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
