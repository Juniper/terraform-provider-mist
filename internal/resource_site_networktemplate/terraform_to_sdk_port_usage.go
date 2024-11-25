package resource_site_networktemplate

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func portUsageScTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SwitchPortUsageStormControl {
	data := models.SwitchPortUsageStormControl{}
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
func portUsageRulesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SwitchPortUsageDynamicRule {

	var data []models.SwitchPortUsageDynamicRule
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(RulesValue)
		rule := models.SwitchPortUsageDynamicRule{}
		if v_plan.Equals.ValueStringPointer() != nil {
			rule.Equals = models.ToPointer(v_plan.Equals.ValueString())
		}
		if !v_plan.EqualsAny.IsNull() && !v_plan.EqualsAny.IsUnknown() {
			rule.EqualsAny = mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.EqualsAny)
		}
		if v_plan.Expression.ValueStringPointer() != nil {
			rule.Expression = models.ToPointer(v_plan.Expression.ValueString())
		}
		if v_plan.Usage.ValueStringPointer() != nil {
			rule.Usage = models.ToPointer(v_plan.Usage.ValueString())
		}
		if v_plan.Src.ValueStringPointer() != nil {
			rule.Src = models.SwitchPortUsageDynamicRuleSrcEnum(v_plan.Src.ValueString())
		}
		data = append(data, rule)
	}
	return data
}
func portUsageTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.SwitchPortUsage {
	data := make(map[string]models.SwitchPortUsage)
	for pu_name, pu_attr := range d.Elements() {
		var pu_attr_interface interface{} = pu_attr
		pu_attr_value := pu_attr_interface.(PortUsagesValue)

		new_pu := models.SwitchPortUsage{}
		if pu_attr_value.AllNetworks.ValueBoolPointer() != nil {
			new_pu.AllNetworks = models.ToPointer(pu_attr_value.AllNetworks.ValueBool())
		}
		if pu_attr_value.AllowDhcpd.ValueBoolPointer() != nil {
			new_pu.AllowDhcpd = models.ToPointer(pu_attr_value.AllowDhcpd.ValueBool())
		}
		if pu_attr_value.AllowMultipleSupplicants.ValueBoolPointer() != nil {
			new_pu.AllowMultipleSupplicants = models.ToPointer(pu_attr_value.AllowMultipleSupplicants.ValueBool())
		}
		if pu_attr_value.BypassAuthWhenServerDown.ValueBoolPointer() != nil {
			new_pu.BypassAuthWhenServerDown = models.ToPointer(pu_attr_value.BypassAuthWhenServerDown.ValueBool())
		}
		if pu_attr_value.BypassAuthWhenServerDownForUnkonwnClient.ValueBoolPointer() != nil {
			new_pu.BypassAuthWhenServerDownForUnkonwnClient = models.ToPointer(pu_attr_value.BypassAuthWhenServerDownForUnkonwnClient.ValueBool())
		}
		if pu_attr_value.Description.ValueStringPointer() != nil {
			new_pu.Description = models.ToPointer(pu_attr_value.Description.ValueString())
		}
		if pu_attr_value.DisableAutoneg.ValueBoolPointer() != nil {
			new_pu.DisableAutoneg = models.ToPointer(pu_attr_value.DisableAutoneg.ValueBool())
		}
		if pu_attr_value.Disabled.ValueBoolPointer() != nil {
			new_pu.Disabled = models.ToPointer(pu_attr_value.Disabled.ValueBool())
		}
		if pu_attr_value.Duplex.ValueStringPointer() != nil {
			new_pu.Duplex = models.ToPointer(models.SwitchPortUsageDuplexEnum(pu_attr_value.Duplex.ValueString()))
		}
		if !pu_attr_value.DynamicVlanNetworks.IsNull() && !pu_attr_value.DynamicVlanNetworks.IsUnknown() {
			new_pu.DynamicVlanNetworks = mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.DynamicVlanNetworks)
		}
		if pu_attr_value.EnableMacAuth.ValueBoolPointer() != nil {
			new_pu.EnableMacAuth = models.ToPointer(pu_attr_value.EnableMacAuth.ValueBool())
		}
		if pu_attr_value.EnableQos.ValueBoolPointer() != nil {
			new_pu.EnableQos = models.ToPointer(pu_attr_value.EnableQos.ValueBool())
		}
		if pu_attr_value.GuestNetwork.ValueStringPointer() != nil {
			new_pu.GuestNetwork = models.NewOptional(models.ToPointer(pu_attr_value.GuestNetwork.ValueString()))
		}
		if pu_attr_value.InterSwitchLink.ValueBoolPointer() != nil {
			new_pu.InterSwitchLink = models.ToPointer(pu_attr_value.InterSwitchLink.ValueBool())
		}
		if pu_attr_value.MacAuthOnly.ValueBoolPointer() != nil {
			new_pu.MacAuthOnly = models.ToPointer(pu_attr_value.MacAuthOnly.ValueBool())
		}
		if pu_attr_value.MacAuthPreferred.ValueBoolPointer() != nil {
			new_pu.MacAuthPreferred = models.ToPointer(pu_attr_value.MacAuthPreferred.ValueBool())
		}
		if pu_attr_value.MacAuthProtocol.ValueStringPointer() != nil {
			new_pu.MacAuthProtocol = models.ToPointer(models.SwitchPortUsageMacAuthProtocolEnum(pu_attr_value.MacAuthProtocol.ValueString()))
		}
		if pu_attr_value.MacLimit.ValueInt64Pointer() != nil {
			new_pu.MacLimit = models.ToPointer(int(pu_attr_value.MacLimit.ValueInt64()))
		}
		if pu_attr_value.Mode.ValueStringPointer() != nil {
			new_pu.Mode = models.ToPointer(models.SwitchPortUsageModeEnum(pu_attr_value.Mode.ValueString()))
		}
		if pu_attr_value.Mtu.ValueInt64Pointer() != nil {
			new_pu.Mtu = models.ToPointer(int(pu_attr_value.Mtu.ValueInt64()))
		}
		if !pu_attr_value.Networks.IsNull() && !pu_attr_value.Networks.IsUnknown() {
			new_pu.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, pu_attr_value.Networks)
		}
		if pu_attr_value.PersistMac.ValueBoolPointer() != nil {
			new_pu.PersistMac = models.ToPointer(pu_attr_value.PersistMac.ValueBool())
		}
		if pu_attr_value.PoeDisabled.ValueBoolPointer() != nil {
			new_pu.PoeDisabled = models.ToPointer(pu_attr_value.PoeDisabled.ValueBool())
		}
		if pu_attr_value.PortAuth.ValueStringPointer() != nil {
			new_pu.PortAuth = models.NewOptional(models.ToPointer(models.SwitchPortUsageDot1xEnum(pu_attr_value.PortAuth.ValueString())))
		}
		if pu_attr_value.PortNetwork.ValueStringPointer() != nil {
			new_pu.PortNetwork = models.ToPointer(pu_attr_value.PortNetwork.ValueString())
		}
		if pu_attr_value.ReauthInterval.ValueInt64Pointer() != nil {
			new_pu.ReauthInterval = models.ToPointer(int(pu_attr_value.ReauthInterval.ValueInt64()))
		}
		if !pu_attr_value.Rules.IsNull() && !pu_attr_value.Rules.IsUnknown() {
			new_pu.Rules = portUsageRulesTerraformToSdk(ctx, diags, pu_attr_value.Rules)
		}
		if pu_attr_value.ResetDefaultWhen.ValueStringPointer() != nil {
			new_pu.ResetDefaultWhen = models.ToPointer(models.SwitchPortUsageDynamicResetDefaultWhenEnum(pu_attr_value.ResetDefaultWhen.ValueString()))
		}
		if pu_attr_value.ServerFailNetwork.ValueStringPointer() != nil {
			new_pu.ServerFailNetwork = models.NewOptional(models.ToPointer(pu_attr_value.ServerFailNetwork.ValueString()))
		}
		if pu_attr_value.ServerRejectNetwork.ValueStringPointer() != nil {
			new_pu.ServerRejectNetwork = models.NewOptional(models.ToPointer(pu_attr_value.ServerRejectNetwork.ValueString()))
		}
		if pu_attr_value.Speed.ValueStringPointer() != nil {
			new_pu.Speed = (*models.SwitchPortUsageSpeedEnum)(pu_attr_value.Speed.ValueStringPointer())
		}
		if !pu_attr_value.StormControl.IsNull() && !pu_attr_value.StormControl.IsUnknown() {
			storm_control := portUsageScTerraformToSdk(ctx, diags, pu_attr_value.StormControl)
			new_pu.StormControl = models.ToPointer(storm_control)
		}
		if pu_attr_value.StpEdge.ValueBoolPointer() != nil {
			new_pu.StpEdge = models.ToPointer(pu_attr_value.StpEdge.ValueBool())
		}
		if pu_attr_value.StpNoRootPort.ValueBoolPointer() != nil {
			new_pu.StpNoRootPort = pu_attr_value.StpNoRootPort.ValueBoolPointer()
		}
		if pu_attr_value.StpP2p.ValueBoolPointer() != nil {
			new_pu.StpP2p = pu_attr_value.StpP2p.ValueBoolPointer()
		}
		if pu_attr_value.UiEvpntopoId.ValueStringPointer() != nil {
			ui_evpntopo_id, e := uuid.Parse(pu_attr_value.UiEvpntopoId.ValueString())
			if e == nil {
				new_pu.UiEvpntopoId = &ui_evpntopo_id
			} else {
				diags.AddError("Bad value for ui_evpntopo_id", e.Error())
			}
		}
		if pu_attr_value.UseVstp.ValueBoolPointer() != nil {
			new_pu.UseVstp = pu_attr_value.UseVstp.ValueBoolPointer()
		}
		if pu_attr_value.VoipNetwork.ValueStringPointer() != nil {
			new_pu.VoipNetwork = models.ToPointer(pu_attr_value.VoipNetwork.ValueString())
		}

		data[pu_name] = new_pu
	}
	return data
}
