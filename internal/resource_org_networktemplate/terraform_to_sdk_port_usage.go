package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func portUsageScTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.SwitchPortUsageStormControl {
	data := models.SwitchPortUsageStormControl{}
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
func portUsageRulesTerraformToSdk(d basetypes.ListValue) []models.SwitchPortUsageDynamicRule {

	var data []models.SwitchPortUsageDynamicRule
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(RulesValue)
		rule := models.SwitchPortUsageDynamicRule{}
		if vPlan.Equals.ValueStringPointer() != nil {
			rule.Equals = models.ToPointer(vPlan.Equals.ValueString())
		}
		if !vPlan.EqualsAny.IsNull() && !vPlan.EqualsAny.IsUnknown() {
			rule.EqualsAny = misttransform.ListOfStringTerraformToSdk(vPlan.EqualsAny)
		}
		if vPlan.Expression.ValueStringPointer() != nil {
			rule.Expression = models.ToPointer(vPlan.Expression.ValueString())
		}
		if vPlan.Usage.ValueStringPointer() != nil {
			rule.Usage = models.ToPointer(vPlan.Usage.ValueString())
		}
		if vPlan.Src.ValueStringPointer() != nil {
			rule.Src = models.SwitchPortUsageDynamicRuleSrcEnum(vPlan.Src.ValueString())
		}
		data = append(data, rule)
	}
	return data
}
func portUsageTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.SwitchPortUsage {
	data := make(map[string]models.SwitchPortUsage)
	for puName, puAttr := range d.Elements() {
		var puAttrInterface interface{} = puAttr
		puAttrValue := puAttrInterface.(PortUsagesValue)

		newPu := models.SwitchPortUsage{}
		if puAttrValue.AllNetworks.ValueBoolPointer() != nil {
			newPu.AllNetworks = models.ToPointer(puAttrValue.AllNetworks.ValueBool())
		}
		if puAttrValue.AllowDhcpd.ValueBoolPointer() != nil {
			newPu.AllowDhcpd = models.ToPointer(puAttrValue.AllowDhcpd.ValueBool())
		}
		if puAttrValue.AllowMultipleSupplicants.ValueBoolPointer() != nil {
			newPu.AllowMultipleSupplicants = models.ToPointer(puAttrValue.AllowMultipleSupplicants.ValueBool())
		}
		if puAttrValue.BypassAuthWhenServerDown.ValueBoolPointer() != nil {
			newPu.BypassAuthWhenServerDown = models.ToPointer(puAttrValue.BypassAuthWhenServerDown.ValueBool())
		}
		if puAttrValue.BypassAuthWhenServerDownForUnkownClient.ValueBoolPointer() != nil {
			newPu.BypassAuthWhenServerDownForUnkownClient = models.ToPointer(puAttrValue.BypassAuthWhenServerDownForUnkownClient.ValueBool())
		}
		if puAttrValue.Description.ValueStringPointer() != nil {
			newPu.Description = models.ToPointer(puAttrValue.Description.ValueString())
		}
		if puAttrValue.DisableAutoneg.ValueBoolPointer() != nil {
			newPu.DisableAutoneg = models.ToPointer(puAttrValue.DisableAutoneg.ValueBool())
		}
		if puAttrValue.Disabled.ValueBoolPointer() != nil {
			newPu.Disabled = models.ToPointer(puAttrValue.Disabled.ValueBool())
		}
		if puAttrValue.Duplex.ValueStringPointer() != nil {
			newPu.Duplex = models.ToPointer(models.SwitchPortUsageDuplexEnum(puAttrValue.Duplex.ValueString()))
		}
		if !puAttrValue.DynamicVlanNetworks.IsNull() && !puAttrValue.DynamicVlanNetworks.IsUnknown() {
			newPu.DynamicVlanNetworks = misttransform.ListOfStringTerraformToSdk(puAttrValue.DynamicVlanNetworks)
		}
		if puAttrValue.EnableMacAuth.ValueBoolPointer() != nil {
			newPu.EnableMacAuth = models.ToPointer(puAttrValue.EnableMacAuth.ValueBool())
		}
		if puAttrValue.EnableQos.ValueBoolPointer() != nil {
			newPu.EnableQos = models.ToPointer(puAttrValue.EnableQos.ValueBool())
		}
		if puAttrValue.GuestNetwork.ValueStringPointer() != nil {
			newPu.GuestNetwork = models.NewOptional(models.ToPointer(puAttrValue.GuestNetwork.ValueString()))
		}
		if puAttrValue.InterSwitchLink.ValueBoolPointer() != nil {
			newPu.InterSwitchLink = models.ToPointer(puAttrValue.InterSwitchLink.ValueBool())
		}
		if puAttrValue.MacAuthOnly.ValueBoolPointer() != nil {
			newPu.MacAuthOnly = models.ToPointer(puAttrValue.MacAuthOnly.ValueBool())
		}
		if puAttrValue.MacAuthPreferred.ValueBoolPointer() != nil {
			newPu.MacAuthPreferred = models.ToPointer(puAttrValue.MacAuthPreferred.ValueBool())
		}
		if puAttrValue.MacAuthProtocol.ValueStringPointer() != nil {
			newPu.MacAuthProtocol = models.ToPointer(models.SwitchPortUsageMacAuthProtocolEnum(puAttrValue.MacAuthProtocol.ValueString()))
		}
		if puAttrValue.MacLimit.ValueInt64Pointer() != nil {
			newPu.MacLimit = models.ToPointer(int(puAttrValue.MacLimit.ValueInt64()))
		}
		if puAttrValue.Mode.ValueStringPointer() != nil {
			newPu.Mode = models.ToPointer(models.SwitchPortUsageModeEnum(puAttrValue.Mode.ValueString()))
		}
		if puAttrValue.Mtu.ValueInt64Pointer() != nil {
			newPu.Mtu = models.ToPointer(int(puAttrValue.Mtu.ValueInt64()))
		}
		if !puAttrValue.Networks.IsNull() && !puAttrValue.Networks.IsUnknown() {
			newPu.Networks = misttransform.ListOfStringTerraformToSdk(puAttrValue.Networks)
		}
		if puAttrValue.PersistMac.ValueBoolPointer() != nil {
			newPu.PersistMac = models.ToPointer(puAttrValue.PersistMac.ValueBool())
		}
		if puAttrValue.PoeDisabled.ValueBoolPointer() != nil {
			newPu.PoeDisabled = models.ToPointer(puAttrValue.PoeDisabled.ValueBool())
		}
		if puAttrValue.PortAuth.ValueStringPointer() != nil {
			newPu.PortAuth = models.NewOptional(models.ToPointer(models.SwitchPortUsageDot1xEnum(puAttrValue.PortAuth.ValueString())))
		}
		if puAttrValue.PortNetwork.ValueStringPointer() != nil {
			newPu.PortNetwork = models.ToPointer(puAttrValue.PortNetwork.ValueString())
		}
		if puAttrValue.ReauthInterval.ValueInt64Pointer() != nil {
			newPu.ReauthInterval = models.ToPointer(int(puAttrValue.ReauthInterval.ValueInt64()))
		}
		if !puAttrValue.Rules.IsNull() && !puAttrValue.Rules.IsUnknown() {
			newPu.Rules = portUsageRulesTerraformToSdk(puAttrValue.Rules)
		}
		if puAttrValue.ResetDefaultWhen.ValueStringPointer() != nil {
			newPu.ResetDefaultWhen = models.ToPointer(models.SwitchPortUsageDynamicResetDefaultWhenEnum(puAttrValue.ResetDefaultWhen.ValueString()))
		}
		if puAttrValue.ServerFailNetwork.ValueStringPointer() != nil {
			newPu.ServerFailNetwork = models.NewOptional(models.ToPointer(puAttrValue.ServerFailNetwork.ValueString()))
		}
		if puAttrValue.ServerRejectNetwork.ValueStringPointer() != nil {
			newPu.ServerRejectNetwork = models.NewOptional(models.ToPointer(puAttrValue.ServerRejectNetwork.ValueString()))
		}
		if puAttrValue.Speed.ValueStringPointer() != nil {
			newPu.Speed = (*models.SwitchPortUsageSpeedEnum)(puAttrValue.Speed.ValueStringPointer())
		}
		if !puAttrValue.StormControl.IsNull() && !puAttrValue.StormControl.IsUnknown() {
			stormControl := portUsageScTerraformToSdk(ctx, diags, puAttrValue.StormControl)
			newPu.StormControl = models.ToPointer(stormControl)
		}
		if puAttrValue.StpEdge.ValueBoolPointer() != nil {
			newPu.StpEdge = models.ToPointer(puAttrValue.StpEdge.ValueBool())
		}
		if puAttrValue.StpNoRootPort.ValueBoolPointer() != nil {
			newPu.StpNoRootPort = puAttrValue.StpNoRootPort.ValueBoolPointer()
		}
		if puAttrValue.StpP2p.ValueBoolPointer() != nil {
			newPu.StpP2p = puAttrValue.StpP2p.ValueBoolPointer()
		}
		if puAttrValue.UiEvpntopoId.ValueStringPointer() != nil {
			uiEvpntopoId, e := uuid.Parse(puAttrValue.UiEvpntopoId.ValueString())
			if e == nil {
				newPu.UiEvpntopoId = &uiEvpntopoId
			} else {
				diags.AddError("Bad value for ui_evpntopo_id", e.Error())
			}
		}
		if puAttrValue.UseVstp.ValueBoolPointer() != nil {
			newPu.UseVstp = puAttrValue.UseVstp.ValueBoolPointer()
		}
		if puAttrValue.VoipNetwork.ValueStringPointer() != nil {
			newPu.VoipNetwork = models.ToPointer(puAttrValue.VoipNetwork.ValueString())
		}

		data[puName] = newPu
	}
	return data
}
