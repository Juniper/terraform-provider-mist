package resource_device_switch

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dhcpdConfigFixedBindingsTerraformToSdk(d basetypes.MapValue) map[string]models.DhcpdConfigFixedBinding {
	dataMap := make(map[string]models.DhcpdConfigFixedBinding)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(FixedBindingsValue)

		data := models.DhcpdConfigFixedBinding{}
		if plan.Ip.ValueStringPointer() != nil {
			data.Ip = models.ToPointer(plan.Ip.ValueString())
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = models.ToPointer(plan.Name.ValueString())
		}

		dataMap[k] = data
	}
	return dataMap
}

func dhcpdConfigOptionsTerraformToSdk(d basetypes.MapValue) map[string]models.DhcpdConfigOption {
	dataMap := make(map[string]models.DhcpdConfigOption)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(OptionsValue)

		data := models.DhcpdConfigOption{}
		if plan.OptionsType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.DhcpdConfigOptionTypeEnum(plan.OptionsType.ValueString()))
		}
		if plan.Value.ValueStringPointer() != nil {
			data.Value = models.ToPointer(plan.Value.ValueString())
		}

		dataMap[k] = data
	}
	return dataMap
}

func dhcpdConfigVendorOptionsTerraformToSdk(d basetypes.MapValue) map[string]models.DhcpdConfigVendorOption {
	dataMap := make(map[string]models.DhcpdConfigVendorOption)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(VendorEncapsulatedValue)

		data := models.DhcpdConfigVendorOption{}
		if plan.VendorEncapsulatedType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.DhcpdConfigVendorOptionTypeEnum(plan.VendorEncapsulatedType.ValueString()))
		}
		if plan.Value.ValueStringPointer() != nil {
			data.Value = models.ToPointer(plan.Value.ValueString())
		}

		dataMap[k] = data
	}
	return dataMap
}

func dhcpdConfigConfigsTerraformToSdk(d basetypes.MapValue) map[string]models.SwitchDhcpdConfigProperty {
	dataMap := make(map[string]models.SwitchDhcpdConfigProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(ConfigValue)

		fixedBindings := dhcpdConfigFixedBindingsTerraformToSdk(plan.FixedBindings)
		options := dhcpdConfigOptionsTerraformToSdk(plan.Options)
		vendorEncapsulated := dhcpdConfigVendorOptionsTerraformToSdk(plan.VendorEncapsulated)

		data := models.SwitchDhcpdConfigProperty{}
		if !plan.DnsServers.IsNull() && !plan.DnsServers.IsUnknown() {
			data.DnsServers = misttransform.ListOfStringTerraformToSdk(plan.DnsServers)
		}
		if !plan.DnsSuffix.IsNull() && !plan.DnsSuffix.IsUnknown() {
			data.DnsSuffix = misttransform.ListOfStringTerraformToSdk(plan.DnsSuffix)
		}
		if !plan.FixedBindings.IsNull() && !plan.FixedBindings.IsUnknown() {
			data.FixedBindings = fixedBindings
		}
		if plan.Gateway.ValueStringPointer() != nil {
			data.Gateway = models.ToPointer(plan.Gateway.ValueString())
		}
		if plan.IpEnd.ValueStringPointer() != nil {
			data.IpEnd = models.ToPointer(plan.IpEnd.ValueString())
		}
		if plan.IpEnd6.ValueStringPointer() != nil {
			data.IpEnd6 = models.ToPointer(plan.IpEnd6.ValueString())
		}
		if plan.IpStart.ValueStringPointer() != nil {
			data.IpStart = models.ToPointer(plan.IpStart.ValueString())
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
		if !plan.Servers.IsNull() && !plan.Servers.IsUnknown() {
			data.Servers = misttransform.ListOfStringTerraformToSdk(plan.Servers)
		}
		if !plan.Servers6.IsNull() && !plan.Servers6.IsUnknown() {
			data.Servers6 = misttransform.ListOfStringTerraformToSdk(plan.Servers6)
		}
		if plan.ConfigType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.SwitchDhcpdConfigTypeEnum(plan.ConfigType.ValueString()))
		}
		if plan.Type6.ValueStringPointer() != nil {
			data.Type6 = models.ToPointer(models.SwitchDhcpdConfigTypeEnum(plan.Type6.ValueString()))
		}
		if !plan.VendorEncapsulated.IsNull() && !plan.VendorEncapsulated.IsUnknown() {
			data.VendorEncapsulated = vendorEncapsulated
		}

		dataMap[k] = data
	}
	return dataMap
}

func dhcpdConfigTerraformToSdk(d DhcpdConfigValue) *models.SwitchDhcpdConfig {

	data := models.SwitchDhcpdConfig{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	if !d.Config.IsNull() && !d.Config.IsUnknown() {
		data.AdditionalProperties = dhcpdConfigConfigsTerraformToSdk(d.Config)
	}

	return &data
}
