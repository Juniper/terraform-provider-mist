package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func dhcpdConfigVendorEncapsulatedSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.DhcpdConfigVendorOption) basetypes.MapValue {

	rMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var typeOption basetypes.StringValue
		var value basetypes.StringValue

		if d.Type != nil {
			typeOption = types.StringValue(string(*d.Type))
		}
		if d.Value != nil {
			value = types.StringValue(*d.Value)
		}

		dataMapValue := map[string]attr.Value{
			"type":  typeOption,
			"value": value,
		}
		data, e := NewVendorEncapsulatedValue(VendorEncapsulatedValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		rMapValue[k] = data
	}
	stateResultMap, e := types.MapValueFrom(ctx, VendorEncapsulatedValue{}.Type(ctx), rMapValue)
	diags.Append(e...)
	return stateResultMap
}

func dhcpdConfigOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.DhcpdConfigOption) basetypes.MapValue {

	rMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var typeOption basetypes.StringValue
		var value basetypes.StringValue

		if d.Type != nil {
			typeOption = types.StringValue(string(*d.Type))
		}
		if d.Value != nil {
			value = types.StringValue(*d.Value)
		}

		dataMapValue := map[string]attr.Value{
			"type":  typeOption,
			"value": value,
		}
		data, e := NewOptionsValue(OptionsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		rMapValue[k] = data
	}
	stateResultMap, e := types.MapValueFrom(ctx, OptionsValue{}.Type(ctx), rMapValue)
	diags.Append(e...)
	return stateResultMap
}

func dhcpdConfigFixedBindingsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.DhcpdConfigFixedBinding) basetypes.MapValue {
	rMap := make(map[string]attr.Value)
	for k, d := range m {
		var ip basetypes.StringValue
		var ip6 basetypes.StringValue
		var name basetypes.StringValue

		if d.Ip != nil {
			ip = types.StringValue(*d.Ip)
		}
		if d.Ip6 != nil {
			ip6 = types.StringValue(*d.Ip6)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}

		dataMapValue := map[string]attr.Value{
			"ip":   ip,
			"ip6":  ip6,
			"name": name,
		}
		data, e := NewFixedBindingsValue(FixedBindingsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		rMap[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, FixedBindingsValue{}.Type(ctx), rMap)
	diags.Append(e...)
	return stateResult
}

func dhcpdConfigConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.DhcpdConfigProperty) basetypes.MapValue {
	rMapValue := make(map[string]attr.Value)
	for k, d := range m {
		if k != "enabled" {
			var dnsServers = types.ListNull(types.StringType)
			var dnsSuffix = types.ListNull(types.StringType)
			var fixedBindings = types.MapNull(FixedBindingsValue{}.Type(ctx))
			var gateway basetypes.StringValue
			var ipEnd basetypes.StringValue
			var ip6End basetypes.StringValue
			var ipStart basetypes.StringValue
			var ip6Start basetypes.StringValue
			var leaseTime basetypes.Int64Value
			var options = types.MapNull(OptionsValue{}.Type(ctx))
			var serverIdOverride basetypes.BoolValue
			var servers = types.ListNull(types.StringType)
			var serversv6 = types.ListNull(types.StringType)
			var type4 basetypes.StringValue
			var type6 basetypes.StringValue
			var vendorEncapsulated = types.MapNull(VendorEncapsulatedValue{}.Type(ctx))

			if d.DnsServers != nil {
				dnsServers = mistutils.ListOfStringSdkToTerraform(d.DnsServers)
			}
			if d.DnsSuffix != nil {
				dnsSuffix = mistutils.ListOfStringSdkToTerraform(d.DnsSuffix)
			}
			if len(d.FixedBindings) > 0 {
				fixedBindings = dhcpdConfigFixedBindingsSdkToTerraform(ctx, diags, d.FixedBindings)
			}
			if d.Gateway != nil {
				gateway = types.StringValue(*d.Gateway)
			}
			if d.IpEnd != nil {
				ipEnd = types.StringValue(*d.IpEnd)
			}
			if d.Ip6End != nil {
				ip6End = types.StringValue(*d.Ip6End)
			}
			if d.IpStart != nil {
				ipStart = types.StringValue(*d.IpStart)
			}
			if d.Ip6Start != nil {
				ip6Start = types.StringValue(*d.Ip6Start)
			}
			if d.LeaseTime != nil {
				leaseTime = types.Int64Value(int64(*d.LeaseTime))
			}
			if len(d.Options) > 0 {
				options = dhcpdConfigOptionsSdkToTerraform(ctx, diags, d.Options)
			}
			if d.ServerIdOverride != nil {
				serverIdOverride = types.BoolValue(*d.ServerIdOverride)
			}
			if d.Servers != nil {
				servers = mistutils.ListOfStringSdkToTerraform(d.Servers)
			}
			if d.Serversv6 != nil {
				serversv6 = mistutils.ListOfStringSdkToTerraform(d.Serversv6)
			}
			if d.Type != nil {
				type4 = types.StringValue(string(*d.Type))
			}
			if d.Type6 != nil {
				type6 = types.StringValue(string(*d.Type6))
			}
			if len(d.VendorEncapsulated) > 0 {
				vendorEncapsulated = dhcpdConfigVendorEncapsulatedSdkToTerraform(ctx, diags, d.VendorEncapsulated)
			}

			dataMapValue := map[string]attr.Value{
				"dns_servers":         dnsServers,
				"dns_suffix":          dnsSuffix,
				"fixed_bindings":      fixedBindings,
				"gateway":             gateway,
				"ip_end":              ipEnd,
				"ip6_end":             ip6End,
				"ip_start":            ipStart,
				"ip6_start":           ip6Start,
				"lease_time":          leaseTime,
				"options":             options,
				"server_id_override":  serverIdOverride,
				"servers":             servers,
				"serversv6":           serversv6,
				"type":                type4,
				"type6":               type6,
				"vendor_encapsulated": vendorEncapsulated,
			}
			data, e := NewConfigValue(ConfigValue{}.AttributeTypes(ctx), dataMapValue)
			diags.Append(e...)

			rMapValue[k] = data
		}
	}
	stateType := ConfigValue{}.Type(ctx)
	r, e := types.MapValueFrom(ctx, stateType, rMapValue)
	diags.Append(e...)
	return r
}

func dhcpdConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.DhcpdConfig) DhcpdConfigValue {

	var config = types.MapNull(ConfigValue{}.Type(ctx))
	var enabled basetypes.BoolValue

	if len(d.AdditionalProperties) > 0 {
		config = dhcpdConfigConfigsSdkToTerraform(ctx, diags, d.AdditionalProperties)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"config":  config,
		"enabled": enabled,
	}
	data, e := NewDhcpdConfigValue(DhcpdConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
