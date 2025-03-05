package resource_site_setting

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func gatewayMgmtProtectReCustomTerraformToSdk(d basetypes.ListValue) []models.ProtectReCustom {
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
			dataItem.Subnets = misttransform.ListOfStringTerraformToSdk(itemObj.Subnets)
		}

		data = append(data, dataItem)
	}
	return data
}
func gatewayMgmtProtectReTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ProtectRe {
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
			for _, item := range itemObj.AllowedServices.Elements() {
				var iface interface{} = item
				val := iface.(basetypes.StringValue)
				items = append(items, models.ProtectReAllowedServiceEnum(val.ValueString()))
			}
			data.AllowedServices = items
		}
		if !itemObj.Custom.IsNull() && !itemObj.Custom.IsUnknown() {
			data.Custom = gatewayMgmtProtectReCustomTerraformToSdk(itemObj.Custom)
		}
		if itemObj.Enabled.ValueBoolPointer() != nil {
			data.Enabled = models.ToPointer(itemObj.Enabled.ValueBool())
		}
		if !itemObj.TrustedHosts.IsNull() && !itemObj.TrustedHosts.IsUnknown() {
			data.TrustedHosts = misttransform.ListOfStringTerraformToSdk(itemObj.TrustedHosts)
		}
		return &data
	}
}

func gatewayMgmtAppProbingCustomTerraformToSdk(d basetypes.ListValue) []models.AppProbingCustomApp {
	var dataList []models.AppProbingCustomApp
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(CustomAppsValue)
		data := models.AppProbingCustomApp{}

		data.Name = plan.Name.ValueStringPointer()
		data.Protocol = (*models.AppProbingCustomAppProtocolEnum)(plan.Protocol.ValueStringPointer())

		data.Hostnames = misttransform.ListOfStringTerraformToSdk(plan.Hostnames)
		if len(data.Hostnames) > 0 {
			data.Key = &data.Hostnames[0]
		}

		if plan.PacketSize.ValueInt64Pointer() != nil {
			data.PacketSize = models.ToPointer(int(plan.PacketSize.ValueInt64()))
		}

		if plan.Protocol.ValueString() == "icmp" {
			data.Address = plan.Key.ValueStringPointer()
		}

		data.Network = plan.Network.ValueStringPointer()
		data.Vrf = plan.Vrf.ValueStringPointer()

		dataList = append(dataList, data)
	}
	return dataList
}

func gatewayMgmtAppProbingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.AppProbing {
	data := models.AppProbing{}
	if !d.IsNull() && !d.IsUnknown() {
		v, e := NewAppProbingValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.Apps = misttransform.ListOfStringTerraformToSdk(v.Apps)
			data.CustomApps = gatewayMgmtAppProbingCustomTerraformToSdk(v.CustomApps)
			data.Enabled = v.Enabled.ValueBoolPointer()
		}
	}
	return &data
}

func gatewayMgmtAutoSignatureUpdateTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SiteSettingGatewayMgmtAutoSignatureUpdate {
	data := models.SiteSettingGatewayMgmtAutoSignatureUpdate{}
	if !d.IsNull() && !d.IsUnknown() {
		v, e := NewAutoSignatureUpdateValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.DayOfWeek = (*models.DayOfWeekEnum)(v.DayOfWeek.ValueStringPointer())
			data.Enable = v.Enable.ValueBoolPointer()
			data.TimeOfDay = v.TimeOfDay.ValueStringPointer()
		}
	}
	return &data
}

func gatewayMgmtTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d GatewayMgmtValue) *models.SiteSettingGatewayMgmt {
	data := models.SiteSettingGatewayMgmt{}

	if !d.AdminSshkeys.IsNull() && !d.AdminSshkeys.IsUnknown() {
		data.AdminSshkeys = misttransform.ListOfStringTerraformToSdk(d.AdminSshkeys)
	}

	if !d.AppProbing.IsNull() && !d.AppProbing.IsUnknown() {
		data.AppProbing = gatewayMgmtAppProbingTerraformToSdk(ctx, diags, d.AppProbing)
	}

	if d.AppUsage.ValueBoolPointer() != nil {
		data.AppUsage = d.AppUsage.ValueBoolPointer()
	}

	if !d.AutoSignatureUpdate.IsNull() && !d.AutoSignatureUpdate.IsUnknown() {
		data.AutoSignatureUpdate = gatewayMgmtAutoSignatureUpdateTerraformToSdk(ctx, diags, d.AutoSignatureUpdate)
	}

	if d.ConfigRevertTimer.ValueInt64Pointer() != nil {
		data.ConfigRevertTimer = models.ToPointer(int(d.ConfigRevertTimer.ValueInt64()))
	}

	if d.DisableConsole.ValueBoolPointer() != nil {
		data.DisableConsole = d.DisableConsole.ValueBoolPointer()
	}

	if d.DisableOob.ValueBoolPointer() != nil {
		data.DisableOob = d.DisableOob.ValueBoolPointer()
	}

	if d.DisableUsb.ValueBoolPointer() != nil {
		data.DisableUsb = d.DisableUsb.ValueBoolPointer()
	}

	if d.FipsEnabled.ValueBoolPointer() != nil {
		data.FipsEnabled = d.FipsEnabled.ValueBoolPointer()
	}

	if !d.ProbeHosts.IsNull() && !d.ProbeHosts.IsUnknown() {
		data.ProbeHosts = misttransform.ListOfStringTerraformToSdk(d.ProbeHosts)
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
