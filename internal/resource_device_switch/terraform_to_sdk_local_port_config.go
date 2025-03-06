package resource_device_switch

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func localPortConfigScTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SwitchPortLocalUsageStormControl {
	data := models.SwitchPortLocalUsageStormControl{}
	if !d.IsNull() && !d.IsUnknown() {
		vPlan, e := NewStormControlValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if vPlan.NoBroadcast.ValueBoolPointer() != nil {
				data.NoBroadcast = models.ToPointer(vPlan.NoBroadcast.ValueBool())
			}
			if vPlan.NoMulticast.ValueBoolPointer() != nil {
				data.NoMulticast = models.ToPointer(vPlan.NoMulticast.ValueBool())
			}
			if vPlan.NoRegisteredMulticast.ValueBoolPointer() != nil {
				data.NoRegisteredMulticast = models.ToPointer(vPlan.NoRegisteredMulticast.ValueBool())
			}
			if vPlan.NoUnknownUnicast.ValueBoolPointer() != nil {
				data.NoUnknownUnicast = models.ToPointer(vPlan.NoUnknownUnicast.ValueBool())
			}
			if vPlan.Percentage.ValueInt64Pointer() != nil {
				data.Percentage = models.ToPointer(int(vPlan.Percentage.ValueInt64()))
			}
		}
	}
	return data
}

func LocalPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.JunosLocalPortConfig {

	data := make(map[string]models.JunosLocalPortConfig)
	for k, v := range d.Elements() {
		var planInterface interface{} = v
		planObj := planInterface.(LocalPortConfigValue)
		itemObj := models.JunosLocalPortConfig{}

		itemObj.Usage = planObj.Usage.ValueString()
		if planObj.AllNetworks.ValueBoolPointer() != nil {
			itemObj.AllNetworks = models.ToPointer(planObj.AllNetworks.ValueBool())
		}
		if planObj.AllowDhcpd.ValueBoolPointer() != nil {
			itemObj.AllowDhcpd = models.ToPointer(planObj.AllowDhcpd.ValueBool())
		}
		if planObj.AllowMultipleSupplicants.ValueBoolPointer() != nil {
			itemObj.AllowMultipleSupplicants = models.ToPointer(planObj.AllowMultipleSupplicants.ValueBool())
		}
		if planObj.BypassAuthWhenServerDown.ValueBoolPointer() != nil {
			itemObj.BypassAuthWhenServerDown = models.ToPointer(planObj.BypassAuthWhenServerDown.ValueBool())
		}
		if planObj.BypassAuthWhenServerDownForUnknownClient.ValueBoolPointer() != nil {
			itemObj.BypassAuthWhenServerDownForUnknownClient = models.ToPointer(planObj.BypassAuthWhenServerDownForUnknownClient.ValueBool())
		}
		if planObj.Description.ValueStringPointer() != nil {
			itemObj.Description = models.ToPointer(planObj.Description.ValueString())
		}
		if planObj.DisableAutoneg.ValueBoolPointer() != nil {
			itemObj.DisableAutoneg = models.ToPointer(planObj.DisableAutoneg.ValueBool())
		}
		if planObj.Disabled.ValueBoolPointer() != nil {
			itemObj.Disabled = models.ToPointer(planObj.Disabled.ValueBool())
		}
		if planObj.Duplex.ValueStringPointer() != nil {
			itemObj.Duplex = models.ToPointer(models.SwitchPortLocalUsageDuplexEnum(planObj.Duplex.ValueString()))
		}
		if !planObj.DynamicVlanNetworks.IsNull() && !planObj.DynamicVlanNetworks.IsUnknown() {
			itemObj.DynamicVlanNetworks = misttransform.ListOfStringTerraformToSdk(planObj.DynamicVlanNetworks)
		}
		if planObj.EnableMacAuth.ValueBoolPointer() != nil {
			itemObj.EnableMacAuth = models.ToPointer(planObj.EnableMacAuth.ValueBool())
		}
		if planObj.EnableQos.ValueBoolPointer() != nil {
			itemObj.EnableQos = models.ToPointer(planObj.EnableQos.ValueBool())
		}
		if planObj.GuestNetwork.ValueStringPointer() != nil {
			itemObj.GuestNetwork = models.NewOptional(models.ToPointer(planObj.GuestNetwork.ValueString()))
		}
		if planObj.InterSwitchLink.ValueBoolPointer() != nil {
			itemObj.InterSwitchLink = models.ToPointer(planObj.InterSwitchLink.ValueBool())
		}
		if planObj.MacAuthOnly.ValueBoolPointer() != nil {
			itemObj.MacAuthOnly = models.ToPointer(planObj.MacAuthOnly.ValueBool())
		}
		if planObj.MacAuthPreferred.ValueBoolPointer() != nil {
			itemObj.MacAuthPreferred = models.ToPointer(planObj.MacAuthPreferred.ValueBool())
		}
		if planObj.MacAuthProtocol.ValueStringPointer() != nil {
			itemObj.MacAuthProtocol = models.ToPointer(models.SwitchPortLocalUsageMacAuthProtocolEnum(planObj.MacAuthProtocol.ValueString()))
		}
		if planObj.MacLimit.ValueInt64Pointer() != nil {
			itemObj.MacLimit = models.ToPointer(int(planObj.MacLimit.ValueInt64()))
		}
		if planObj.Mode.ValueStringPointer() != nil {
			itemObj.Mode = models.ToPointer(models.SwitchPortLocalUsageModeEnum(planObj.Mode.ValueString()))
		}
		if planObj.Mtu.ValueInt64Pointer() != nil {
			itemObj.Mtu = models.ToPointer(int(planObj.Mtu.ValueInt64()))
		}
		if !planObj.Networks.IsNull() && !planObj.Networks.IsUnknown() {
			itemObj.Networks = misttransform.ListOfStringTerraformToSdk(planObj.Networks)
		}
		if planObj.Note.ValueStringPointer() != nil {
			itemObj.Note = planObj.Note.ValueStringPointer()
		}
		if planObj.PersistMac.ValueBoolPointer() != nil {
			itemObj.PersistMac = models.ToPointer(planObj.PersistMac.ValueBool())
		}
		if planObj.PoeDisabled.ValueBoolPointer() != nil {
			itemObj.PoeDisabled = models.ToPointer(planObj.PoeDisabled.ValueBool())
		}
		if planObj.PortAuth.ValueStringPointer() != nil {
			itemObj.PortAuth = models.NewOptional(models.ToPointer(models.SwitchPortLocalUsageDot1xEnum(planObj.PortAuth.ValueString())))
		}
		if planObj.PortNetwork.ValueStringPointer() != nil {
			itemObj.PortNetwork = models.ToPointer(planObj.PortNetwork.ValueString())
		}
		if planObj.ReauthInterval.ValueStringPointer() != nil {
			itemObj.ReauthInterval = models.ToPointer(models.SwitchPortUsageReauthIntervalContainer.FromString(planObj.ReauthInterval.ValueString()))
		}
		if planObj.ServerFailNetwork.ValueStringPointer() != nil {
			itemObj.ServerFailNetwork = models.NewOptional(models.ToPointer(planObj.ServerFailNetwork.ValueString()))
		}
		if planObj.ServerRejectNetwork.ValueStringPointer() != nil {
			itemObj.ServerRejectNetwork = models.NewOptional(models.ToPointer(planObj.ServerRejectNetwork.ValueString()))
		}
		if planObj.Speed.ValueStringPointer() != nil {
			itemObj.Speed = (*models.JunosPortConfigSpeedEnum)(planObj.Speed.ValueStringPointer())
		}
		if !planObj.StormControl.IsNull() && !planObj.StormControl.IsUnknown() {
			stormControl := localPortConfigScTerraformToSdk(ctx, diags, planObj.StormControl)
			itemObj.StormControl = models.ToPointer(stormControl)
		}
		if planObj.StpEdge.ValueBoolPointer() != nil {
			itemObj.StpEdge = models.ToPointer(planObj.StpEdge.ValueBool())
		}
		if planObj.StpNoRootPort.ValueBoolPointer() != nil {
			itemObj.StpNoRootPort = planObj.StpNoRootPort.ValueBoolPointer()
		}
		if planObj.StpP2p.ValueBoolPointer() != nil {
			itemObj.StpP2p = planObj.StpP2p.ValueBoolPointer()
		}
		if planObj.UseVstp.ValueBoolPointer() != nil {
			itemObj.UseVstp = planObj.UseVstp.ValueBoolPointer()
		}
		if planObj.VoipNetwork.ValueStringPointer() != nil {
			itemObj.VoipNetwork = models.ToPointer(planObj.VoipNetwork.ValueString())
		}
		data[k] = itemObj
	}
	return data
}
