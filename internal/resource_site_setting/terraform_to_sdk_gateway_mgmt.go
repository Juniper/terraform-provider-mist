package resource_site_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func gatewayMgmtAppProbingCustomTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.AppProbingCustomApp {
	var data_list []models.AppProbingCustomApp
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(CustomAppsValue)
		data := models.AppProbingCustomApp{}

		data.Address = plan.Address.ValueStringPointer()
		data.AppType = plan.AppType.ValueStringPointer()
		if !plan.Hostname.IsNull() && !plan.Hostname.IsUnknown() {
			data.Hostname = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hostname)
		}
		data.Name = plan.Name.ValueStringPointer()
		data.Network = plan.Network.ValueStringPointer()
		data.Protocol = (*models.AppProbingCustomAppProtocolEnum)(plan.Protocol.ValueStringPointer())
		data.Url = plan.Url.ValueStringPointer()
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
