package resource_device_switch

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func localPortConfigScTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SwitchPortLocalUsageStormControl {
	data := models.SwitchPortLocalUsageStormControl{}
	if !d.IsNull() && !d.IsUnknown() {
		v_plan, e := NewStormControlValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if v_plan.NoBroadcast.ValueBoolPointer() != nil {
				data.NoBroadcast = models.ToPointer(v_plan.NoBroadcast.ValueBool())
			}
			if v_plan.NoMulticast.ValueBoolPointer() != nil {
				data.NoMulticast = models.ToPointer(v_plan.NoMulticast.ValueBool())
			}
			if v_plan.NoRegisteredMulticast.ValueBoolPointer() != nil {
				data.NoRegisteredMulticast = models.ToPointer(v_plan.NoRegisteredMulticast.ValueBool())
			}
			if v_plan.NoUnknownUnicast.ValueBoolPointer() != nil {
				data.NoUnknownUnicast = models.ToPointer(v_plan.NoUnknownUnicast.ValueBool())
			}
			if v_plan.Percentage.ValueInt64Pointer() != nil {
				data.Percentage = models.ToPointer(int(v_plan.Percentage.ValueInt64()))
			}
		}
	}
	return data
}

func LocalPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.JunosLocalPortConfig {

	data := make(map[string]models.JunosLocalPortConfig)
	for k, v := range d.Elements() {
		var plan_interface interface{} = v
		plan_obj := plan_interface.(LocalPortConfigValue)
		item_obj := models.JunosLocalPortConfig{}

		item_obj.Usage = plan_obj.Usage.ValueString()
		if plan_obj.AllNetworks.ValueBoolPointer() != nil {
			item_obj.AllNetworks = models.ToPointer(plan_obj.AllNetworks.ValueBool())
		}
		if plan_obj.AllowDhcpd.ValueBoolPointer() != nil {
			item_obj.AllowDhcpd = models.ToPointer(plan_obj.AllowDhcpd.ValueBool())
		}
		if plan_obj.AllowMultipleSupplicants.ValueBoolPointer() != nil {
			item_obj.AllowMultipleSupplicants = models.ToPointer(plan_obj.AllowMultipleSupplicants.ValueBool())
		}
		if plan_obj.BypassAuthWhenServerDown.ValueBoolPointer() != nil {
			item_obj.BypassAuthWhenServerDown = models.ToPointer(plan_obj.BypassAuthWhenServerDown.ValueBool())
		}
		if plan_obj.BypassAuthWhenServerDownForUnkonwnClient.ValueBoolPointer() != nil {
			item_obj.BypassAuthWhenServerDownForUnkonwnClient = models.ToPointer(plan_obj.BypassAuthWhenServerDownForUnkonwnClient.ValueBool())
		}
		if plan_obj.Description.ValueStringPointer() != nil {
			item_obj.Description = models.ToPointer(plan_obj.Description.ValueString())
		}
		if plan_obj.DisableAutoneg.ValueBoolPointer() != nil {
			item_obj.DisableAutoneg = models.ToPointer(plan_obj.DisableAutoneg.ValueBool())
		}
		if plan_obj.Disabled.ValueBoolPointer() != nil {
			item_obj.Disabled = models.ToPointer(plan_obj.Disabled.ValueBool())
		}
		if plan_obj.Duplex.ValueStringPointer() != nil {
			item_obj.Duplex = models.ToPointer(models.SwitchPortLocalUsageDuplexEnum(plan_obj.Duplex.ValueString()))
		}
		if !plan_obj.DynamicVlanNetworks.IsNull() && !plan_obj.DynamicVlanNetworks.IsUnknown() {
			item_obj.DynamicVlanNetworks = mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.DynamicVlanNetworks)
		}
		if plan_obj.EnableMacAuth.ValueBoolPointer() != nil {
			item_obj.EnableMacAuth = models.ToPointer(plan_obj.EnableMacAuth.ValueBool())
		}
		if plan_obj.EnableQos.ValueBoolPointer() != nil {
			item_obj.EnableQos = models.ToPointer(plan_obj.EnableQos.ValueBool())
		}
		if plan_obj.GuestNetwork.ValueStringPointer() != nil {
			item_obj.GuestNetwork = models.NewOptional(models.ToPointer(plan_obj.GuestNetwork.ValueString()))
		}
		if plan_obj.InterSwitchLink.ValueBoolPointer() != nil {
			item_obj.InterSwitchLink = models.ToPointer(plan_obj.InterSwitchLink.ValueBool())
		}
		if plan_obj.MacAuthOnly.ValueBoolPointer() != nil {
			item_obj.MacAuthOnly = models.ToPointer(plan_obj.MacAuthOnly.ValueBool())
		}
		if plan_obj.MacAuthPreferred.ValueBoolPointer() != nil {
			item_obj.MacAuthPreferred = models.ToPointer(plan_obj.MacAuthPreferred.ValueBool())
		}
		if plan_obj.MacAuthProtocol.ValueStringPointer() != nil {
			item_obj.MacAuthProtocol = models.ToPointer(models.SwitchPortLocalUsageMacAuthProtocolEnum(plan_obj.MacAuthProtocol.ValueString()))
		}
		if plan_obj.MacLimit.ValueInt64Pointer() != nil {
			item_obj.MacLimit = models.ToPointer(int(plan_obj.MacLimit.ValueInt64()))
		}
		if plan_obj.Mode.ValueStringPointer() != nil {
			item_obj.Mode = models.ToPointer(models.SwitchPortLocalUsageModeEnum(plan_obj.Mode.ValueString()))
		}
		if plan_obj.Mtu.ValueInt64Pointer() != nil {
			item_obj.Mtu = models.ToPointer(int(plan_obj.Mtu.ValueInt64()))
		}
		if !plan_obj.Networks.IsNull() && !plan_obj.Networks.IsUnknown() {
			item_obj.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, plan_obj.Networks)
		}
		if plan_obj.PersistMac.ValueBoolPointer() != nil {
			item_obj.PersistMac = models.ToPointer(plan_obj.PersistMac.ValueBool())
		}
		if plan_obj.PoeDisabled.ValueBoolPointer() != nil {
			item_obj.PoeDisabled = models.ToPointer(plan_obj.PoeDisabled.ValueBool())
		}
		if plan_obj.PortAuth.ValueStringPointer() != nil {
			item_obj.PortAuth = models.NewOptional(models.ToPointer(models.SwitchPortLocalUsageDot1xEnum(plan_obj.PortAuth.ValueString())))
		}
		if plan_obj.PortNetwork.ValueStringPointer() != nil {
			item_obj.PortNetwork = models.ToPointer(plan_obj.PortNetwork.ValueString())
		}
		if plan_obj.ReauthInterval.ValueInt64Pointer() != nil {
			item_obj.ReauthInterval = models.ToPointer(int(plan_obj.ReauthInterval.ValueInt64()))
		}
		if plan_obj.ServerFailNetwork.ValueStringPointer() != nil {
			item_obj.ServerFailNetwork = models.NewOptional(models.ToPointer(plan_obj.ServerFailNetwork.ValueString()))
		}
		if plan_obj.ServerRejectNetwork.ValueStringPointer() != nil {
			item_obj.ServerRejectNetwork = models.NewOptional(models.ToPointer(plan_obj.ServerRejectNetwork.ValueString()))
		}
		if plan_obj.Speed.ValueStringPointer() != nil {
			item_obj.Speed = (*models.JunosPortConfigSpeedEnum)(plan_obj.Speed.ValueStringPointer())
		}
		if !plan_obj.StormControl.IsNull() && !plan_obj.StormControl.IsUnknown() {
			storm_control := localPortConfigScTerraformToSdk(ctx, diags, plan_obj.StormControl)
			item_obj.StormControl = models.ToPointer(storm_control)
		}
		if plan_obj.StpEdge.ValueBoolPointer() != nil {
			item_obj.StpEdge = models.ToPointer(plan_obj.StpEdge.ValueBool())
		}
		if plan_obj.StpNoRootPort.ValueBoolPointer() != nil {
			item_obj.StpNoRootPort = plan_obj.StpNoRootPort.ValueBoolPointer()
		}
		if plan_obj.StpP2p.ValueBoolPointer() != nil {
			item_obj.StpP2p = plan_obj.StpP2p.ValueBoolPointer()
		}
		if plan_obj.UseVstp.ValueBoolPointer() != nil {
			item_obj.UseVstp = plan_obj.UseVstp.ValueBoolPointer()
		}
		if plan_obj.VoipNetwork.ValueStringPointer() != nil {
			item_obj.VoipNetwork = models.ToPointer(plan_obj.VoipNetwork.ValueString())
		}
		data[k] = item_obj
	}
	return data
}
