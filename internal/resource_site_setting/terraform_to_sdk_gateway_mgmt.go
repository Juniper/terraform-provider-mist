package resource_site_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func gatewayMgmtProtectReCustomTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ProtectReCustom {
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
		if !item_obj.Subnets.IsNull() && !item_obj.Subnets.IsUnknown() {
			data_item.Subnets = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Subnets)
		}

		data = append(data, data_item)
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
		var item_interface interface{} = item
		item_obj := item_interface.(ProtectReValue)

		if !item_obj.AllowedServices.IsNull() && !item_obj.AllowedServices.IsUnknown() {
			var items []models.ProtectReAllowedServiceEnum
			for _, item := range item_obj.AllowedServices.Elements() {
				var iface interface{} = item
				val := iface.(basetypes.StringValue)
				items = append(items, models.ProtectReAllowedServiceEnum(val.ValueString()))
			}
			data.AllowedServices = items
		}
		if !item_obj.Custom.IsNull() && !item_obj.Custom.IsUnknown() {
			data.Custom = gatewayMgmtProtectReCustomTerraformToSdk(ctx, diags, item_obj.Custom)
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

func gatewayMgmtAppProbingCustomTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.AppProbingCustomApp {
	var data_list []models.AppProbingCustomApp
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(CustomAppsValue)
		data := models.AppProbingCustomApp{}

		data.Name = plan.Name.ValueStringPointer()
		data.Protocol = (*models.AppProbingCustomAppProtocolEnum)(plan.Protocol.ValueStringPointer())

		data.Hostnames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hostnames)
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

		data_list = append(data_list, data)
	}
	return data_list
}

func gatewayMgmtAppProbingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.AppProbing {
	data := models.AppProbing{}
	if !d.IsNull() || !d.IsUnknown() {
		v, e := NewAppProbingValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.Apps = mist_transform.ListOfStringTerraformToSdk(ctx, v.Apps)
			data.CustomApps = gatewayMgmtAppProbingCustomTerraformToSdk(ctx, diags, v.CustomApps)
			data.Enabled = v.Enabled.ValueBoolPointer()
		}
	}
	return &data
}

func gatewayMgmtAutoSignatureUpdateTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SiteSettingGatewayMgmtAutoSignatureUpdate {
	data := models.SiteSettingGatewayMgmtAutoSignatureUpdate{}
	if !d.IsNull() || !d.IsUnknown() {
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
		data.AdminSshkeys = mist_transform.ListOfStringTerraformToSdk(ctx, d.AdminSshkeys)
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

	if !d.ProbeHosts.IsNull() && !d.ProbeHosts.IsUnknown() {
		data.ProbeHosts = mist_transform.ListOfStringTerraformToSdk(ctx, d.ProbeHosts)
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
