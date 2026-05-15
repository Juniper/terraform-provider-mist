package resource_device_gateway

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func gatewayMgmtCustomAppsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.AppProbingCustomApp {
	var dataList []models.AppProbingCustomApp
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(CustomAppsValue)
		data := models.AppProbingCustomApp{}

		if plan.Address.ValueStringPointer() != nil {
			data.Address = plan.Address.ValueStringPointer()
		}
		if plan.AppType.ValueStringPointer() != nil {
			data.AppType = plan.AppType.ValueStringPointer()
		}
		if !plan.Hostnames.IsNull() && !plan.Hostnames.IsUnknown() {
			data.Hostnames = mistutils.ListOfStringTerraformToSdk(plan.Hostnames)
		}
		if plan.Key.ValueStringPointer() != nil {
			data.Key = plan.Key.ValueStringPointer()
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}
		if plan.Network.ValueStringPointer() != nil {
			data.Network = plan.Network.ValueStringPointer()
		}
		if plan.PacketSize.ValueInt64Pointer() != nil {
			data.PacketSize = models.ToPointer(int(plan.PacketSize.ValueInt64()))
		}
		if plan.Protocol.ValueStringPointer() != nil {
			data.Protocol = (*models.AppProbingCustomAppProtocolEnum)(plan.Protocol.ValueStringPointer())
		}
		if plan.Url.ValueStringPointer() != nil {
			data.Url = plan.Url.ValueStringPointer()
		}
		if plan.Vrf.ValueStringPointer() != nil {
			data.Vrf = plan.Vrf.ValueStringPointer()
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func gatewayMgmtAppProbingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.AppProbing {
	data := models.AppProbing{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	}
	plan, e := NewAppProbingValue(d.AttributeTypes(ctx), d.Attributes())
	if e != nil {
		diags.Append(e...)
		return &data
	}
	if !plan.Apps.IsNull() && !plan.Apps.IsUnknown() {
		data.Apps = mistutils.ListOfStringTerraformToSdk(plan.Apps)
	}
	if !plan.CustomApps.IsNull() && !plan.CustomApps.IsUnknown() {
		data.CustomApps = gatewayMgmtCustomAppsTerraformToSdk(ctx, diags, plan.CustomApps)
	}
	if plan.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
	}
	return &data
}

func gatewayMgmtAutoSignatureUpdateTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.GatewayMgmtAutoSignatureUpdate {
	data := models.GatewayMgmtAutoSignatureUpdate{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	}
	plan, e := NewAutoSignatureUpdateValue(d.AttributeTypes(ctx), d.Attributes())
	if e != nil {
		diags.Append(e...)
		return &data
	}
	if plan.DayOfWeek.ValueStringPointer() != nil {
		data.DayOfWeek = (*models.DayOfWeekEnum)(plan.DayOfWeek.ValueStringPointer())
	}
	if plan.Enable.ValueBoolPointer() != nil {
		data.Enable = models.ToPointer(plan.Enable.ValueBool())
	}
	if plan.TimeOfDay.ValueStringPointer() != nil {
		data.TimeOfDay = plan.TimeOfDay.ValueStringPointer()
	}
	return &data
}

func gatewayMgmtProtectReCustomTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ProtectReCustom {
	var dataList []models.ProtectReCustom
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(CustomValue)
		data := models.ProtectReCustom{}

		if plan.PortRange.ValueStringPointer() != nil {
			data.PortRange = plan.PortRange.ValueStringPointer()
		}
		if plan.Protocol.ValueStringPointer() != nil {
			data.Protocol = (*models.ProtectReCustomProtocolEnum)(plan.Protocol.ValueStringPointer())
		}
		if !plan.Subnets.IsNull() && !plan.Subnets.IsUnknown() {
			data.Subnets = mistutils.ListOfStringTerraformToSdk(plan.Subnets)
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func gatewayMgmtProtectReTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ProtectRe {
	data := models.ProtectRe{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	}
	plan, e := NewProtectReValue(d.AttributeTypes(ctx), d.Attributes())
	if e != nil {
		diags.Append(e...)
		return &data
	}
	if !plan.AllowedServices.IsNull() && !plan.AllowedServices.IsUnknown() {
		for _, item := range plan.AllowedServices.Elements() {
			var itemInterface interface{} = item
			itemStr := itemInterface.(basetypes.StringValue)
			data.AllowedServices = append(data.AllowedServices, models.ProtectReAllowedServiceEnum(itemStr.ValueString()))
		}
	}
	if !plan.Custom.IsNull() && !plan.Custom.IsUnknown() {
		data.Custom = gatewayMgmtProtectReCustomTerraformToSdk(ctx, diags, plan.Custom)
	}
	if plan.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
	}
	if plan.HitCount.ValueBoolPointer() != nil {
		data.HitCount = models.ToPointer(plan.HitCount.ValueBool())
	}
	if !plan.TrustedHosts.IsNull() && !plan.TrustedHosts.IsUnknown() {
		data.TrustedHosts = mistutils.ListOfStringTerraformToSdk(plan.TrustedHosts)
	}
	return &data
}

func gatewayMgmtTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d GatewayMgmtValue) *models.GatewayMgmt {
	data := models.GatewayMgmt{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	}

	if !d.AdminSshkeys.IsNull() && !d.AdminSshkeys.IsUnknown() {
		data.AdminSshkeys = mistutils.ListOfStringTerraformToSdk(d.AdminSshkeys)
	}
	if !d.AppProbing.IsNull() && !d.AppProbing.IsUnknown() {
		data.AppProbing = gatewayMgmtAppProbingTerraformToSdk(ctx, diags, d.AppProbing)
	}
	if d.AppUsage.ValueBoolPointer() != nil {
		data.AppUsage = models.ToPointer(d.AppUsage.ValueBool())
	}
	if !d.AutoSignatureUpdate.IsNull() && !d.AutoSignatureUpdate.IsUnknown() {
		data.AutoSignatureUpdate = gatewayMgmtAutoSignatureUpdateTerraformToSdk(ctx, diags, d.AutoSignatureUpdate)
	}
	if d.ConfigRevertTimer.ValueInt64Pointer() != nil {
		data.ConfigRevertTimer = models.ToPointer(int(d.ConfigRevertTimer.ValueInt64()))
	}
	if d.DisableConsole.ValueBoolPointer() != nil {
		data.DisableConsole = models.ToPointer(d.DisableConsole.ValueBool())
	}
	if d.DisableOob.ValueBoolPointer() != nil {
		data.DisableOob = models.ToPointer(d.DisableOob.ValueBool())
	}
	if d.DisableUsb.ValueBoolPointer() != nil {
		data.DisableUsb = models.ToPointer(d.DisableUsb.ValueBool())
	}
	if d.FipsEnabled.ValueBoolPointer() != nil {
		data.FipsEnabled = models.ToPointer(d.FipsEnabled.ValueBool())
	}
	if !d.ProbeHosts.IsNull() && !d.ProbeHosts.IsUnknown() {
		data.ProbeHosts = mistutils.ListOfStringTerraformToSdk(d.ProbeHosts)
	}
	if !d.ProbeHostsv6.IsNull() && !d.ProbeHostsv6.IsUnknown() {
		data.ProbeHostsv6 = mistutils.ListOfStringTerraformToSdk(d.ProbeHostsv6)
	}
	if !d.ProtectRe.IsNull() && !d.ProtectRe.IsUnknown() {
		data.ProtectRe = gatewayMgmtProtectReTerraformToSdk(ctx, diags, d.ProtectRe)
	}
	if d.RootPassword.ValueStringPointer() != nil {
		data.RootPassword = d.RootPassword.ValueStringPointer()
	}
	if d.SecurityLogSourceAddress.ValueStringPointer() != nil {
		data.SecurityLogSourceAddress = d.SecurityLogSourceAddress.ValueStringPointer()
	}
	if d.SecurityLogSourceInterface.ValueStringPointer() != nil {
		data.SecurityLogSourceInterface = d.SecurityLogSourceInterface.ValueStringPointer()
	}
	return &data
}
