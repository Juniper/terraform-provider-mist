package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func dhcpdConfigFixedBindingsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.DhcpdConfigFixedBinding {
	tflog.Debug(ctx, "dhcpdConfigFixedBindingsTerraformToSdk")
	data_map := make(map[string]models.DhcpdConfigFixedBinding)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(FixedBindingsValue)

		data := models.DhcpdConfigFixedBinding{}
		if plan.Ip.ValueStringPointer() != nil {
			data.Ip = models.ToPointer(plan.Ip.ValueString())
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = models.ToPointer(plan.Name.ValueString())
		}

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.DhcpdConfigOption {
	tflog.Debug(ctx, "dhcpdConfigOptionsTerraformToSdk")
	data_map := make(map[string]models.DhcpdConfigOption)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OptionsValue)

		data := models.DhcpdConfigOption{}
		if plan.OptionsType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.DhcpdConfigOptionTypeEnum(plan.OptionsType.ValueString()))
		}
		if plan.Value.ValueStringPointer() != nil {
			data.Value = models.ToPointer(plan.Value.ValueString())
		}

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigVendorOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.DhcpdConfigVendorOption {
	tflog.Debug(ctx, "dhcpdConfigVendorOptionsTerraformToSdk")
	data_map := make(map[string]models.DhcpdConfigVendorOption)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VendorEncapulatedValue)

		data := models.DhcpdConfigVendorOption{}
		if plan.VendorEncapulatedType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.DhcpdConfigVendorOptionTypeEnum(plan.VendorEncapulatedType.ValueString()))
		}
		if plan.Value.ValueStringPointer() != nil {
			data.Value = models.ToPointer(plan.Value.ValueString())
		}

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]interface{} {
	tflog.Debug(ctx, "dhcpdConfigConfigsTerraformToSdk")
	data_map := make(map[string]interface{})
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ConfigValue)

		fixed_bindings := dhcpdConfigFixedBindingsTerraformToSdk(ctx, diags, plan.FixedBindings)
		options := dhcpdConfigOptionsTerraformToSdk(ctx, diags, plan.Options)
		vendor_encapulated := dhcpdConfigVendorOptionsTerraformToSdk(ctx, diags, plan.VendorEncapulated)

		data := models.DhcpdConfigProperty{}
		if !plan.DnsServers.IsNull() && !plan.DnsServers.IsUnknown() {
			data.DnsServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)
		}
		if !plan.DnsSuffix.IsNull() && !plan.DnsSuffix.IsUnknown() {
			data.DnsSuffix = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
		}
		if !plan.FixedBindings.IsNull() && !plan.FixedBindings.IsUnknown() {
			data.FixedBindings = fixed_bindings
		}
		if plan.Gateway.ValueStringPointer() != nil {
			data.Gateway = models.ToPointer(plan.Gateway.ValueString())
		}
		if plan.IpEnd4.ValueStringPointer() != nil {
			data.IpEnd = models.ToPointer(plan.IpEnd4.ValueString())
		}
		if plan.IpEnd6.ValueStringPointer() != nil {
			data.IpEnd6 = models.ToPointer(plan.IpEnd6.ValueString())
		}
		if plan.IpStart4.ValueStringPointer() != nil {
			data.IpStart = models.ToPointer(plan.IpStart4.ValueString())
		}
		if plan.IpStart6.ValueStringPointer() != nil {
			data.IpStart6 = models.ToPointer(plan.IpStart6.ValueString())
		}
		if plan.LeaseTime.ValueInt64Pointer() != nil {
			data.LeaseTime = models.ToPointer(int(plan.LeaseTime.ValueInt64()))
		}
		if !plan.Options.IsNull() && !plan.Options.IsUnknown() {
			data.Options = options
		}
		if plan.ServerIdOverride.ValueBoolPointer() != nil {
			data.ServerIdOverride = models.ToPointer(plan.ServerIdOverride.ValueBool())
		}
		if !plan.Servers4.IsNull() && !plan.Servers4.IsUnknown() {
			data.Servers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Servers4)
		}
		if !plan.Servers6.IsNull() && !plan.Servers6.IsUnknown() {
			data.Servers6 = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Servers6)
		}
		if plan.Type4.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.DhcpdConfigTypeEnum(plan.Type4.ValueString()))
		}
		if plan.Type6.ValueStringPointer() != nil {
			data.Type6 = models.ToPointer(models.DhcpdConfigTypeEnum(plan.Type6.ValueString()))
		}
		if !plan.VendorEncapulated.IsNull() && !plan.VendorEncapulated.IsUnknown() {
			data.VendorEncapulated = vendor_encapulated
		}

		data_map[k] = data
	}
	return data_map
}

func dhcpdConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpdConfigValue) models.DhcpdConfig {
	tflog.Debug(ctx, "dhcpdConfigTerraformToSdk")

	data := models.DhcpdConfig{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	if !d.Config.IsNull() && !d.Config.IsUnknown() {
		data.AdditionalProperties = dhcpdConfigConfigsTerraformToSdk(ctx, diags, d.Config)
	}

	return data
}
