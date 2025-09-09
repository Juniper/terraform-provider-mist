package resource_device_gateway

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tunnelConfigsAutoProvisionPrimaryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.TunnelConfigAutoProvisionNode {
	data := models.TunnelConfigAutoProvisionNode{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan, e := NewAutoProvisionPrimaryValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.ProbeIps.IsNull() && !plan.ProbeIps.IsUnknown() {
				data.ProbeIps = mistutils.ListOfStringTerraformToSdk(plan.ProbeIps)
			}
			if !plan.WanNames.IsNull() && !plan.WanNames.IsUnknown() {
				data.WanNames = mistutils.ListOfStringTerraformToSdk(plan.WanNames)
			}
		}
		return &data
	}
}

func tunnelConfigsAutoProvisionSecondaryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.TunnelConfigAutoProvisionNode {
	data := models.TunnelConfigAutoProvisionNode{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan, e := NewAutoProvisionSecondaryValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.ProbeIps.IsNull() && !plan.ProbeIps.IsUnknown() {
				data.ProbeIps = mistutils.ListOfStringTerraformToSdk(plan.ProbeIps)
			}
			if !plan.WanNames.IsNull() && !plan.WanNames.IsUnknown() {
				data.WanNames = mistutils.ListOfStringTerraformToSdk(plan.WanNames)
			}
		}
		return &data
	}
}

func tunnelConfigsAutoProvisionTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelConfigAutoProvision {
	data := models.TunnelConfigAutoProvision{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan, e := NewAutoProvisionValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.AutoProvisionPrimary.IsNull() && !plan.AutoProvisionPrimary.IsUnknown() {
				data.Primary = tunnelConfigsAutoProvisionPrimaryTerraformToSdk(ctx, diags, plan.AutoProvisionPrimary)
			}

			if !plan.AutoProvisionSecondary.IsNull() && !plan.AutoProvisionSecondary.IsUnknown() {
				data.Secondary = tunnelConfigsAutoProvisionSecondaryTerraformToSdk(ctx, diags, plan.AutoProvisionSecondary)
			}
			if plan.Enabled.ValueBoolPointer() != nil {
				data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
			}

			if !plan.Latlng.IsNull() && !plan.Latlng.IsUnknown() {
				var planLatlngInterface interface{} = plan.Latlng
				planLatlng := planLatlngInterface.(LatlngValue)

				var latlng models.TunnelConfigAutoProvisionLatLng
				latlng.Lat = planLatlng.Lng.ValueFloat64()
				latlng.Lng = planLatlng.Lng.ValueFloat64()
				data.Latlng = models.ToPointer(latlng)
			}

			if plan.Provider.ValueStringPointer() != nil {
				data.Provider = models.TunnelConfigAutoProvisionProviderEnum(*plan.Provider.ValueStringPointer())
			}

			if plan.Region.ValueStringPointer() != nil {
				data.Region = plan.Region.ValueStringPointer()
			}

			if plan.ServiceConnection.ValueStringPointer() != nil {
				data.ServiceConnection = plan.ServiceConnection.ValueStringPointer()
			}

		}
		return data
	}
}

func gatewayTemplateTunnelIkeProposalTerraformToSdk(d basetypes.ListValue) []models.TunnelConfigIkeProposal {
	var dataList []models.TunnelConfigIkeProposal
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(IkeProposalsValue)
		data := models.TunnelConfigIkeProposal{}
		if plan.AuthAlgo.ValueStringPointer() != nil {
			data.AuthAlgo = models.ToPointer(models.TunnelConfigAuthAlgoEnum(plan.AuthAlgo.ValueString()))
		}
		if plan.DhGroup.ValueStringPointer() != nil {
			data.DhGroup = models.ToPointer(models.TunnelConfigIkeDhGroupEnum(plan.DhGroup.ValueString()))
		}
		if plan.EncAlgo.ValueStringPointer() != nil {
			data.EncAlgo = models.NewOptional(models.ToPointer(models.TunnelConfigEncAlgoEnum(plan.EncAlgo.ValueString())))
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func gatewayTemplateTunnelIpsecProposalTerraformToSdk(d basetypes.ListValue) []models.TunnelConfigIpsecProposal {
	var dataList []models.TunnelConfigIpsecProposal
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(IpsecProposalsValue)
		data := models.TunnelConfigIpsecProposal{}
		if plan.AuthAlgo.ValueStringPointer() != nil {
			data.AuthAlgo = models.ToPointer(models.TunnelConfigAuthAlgoEnum(plan.AuthAlgo.ValueString()))
		}
		if plan.DhGroup.ValueStringPointer() != nil {
			data.DhGroup = models.ToPointer(models.TunnelConfigDhGroupEnum(plan.DhGroup.ValueString()))
		}
		if plan.EncAlgo.ValueStringPointer() != nil {
			data.EncAlgo = models.NewOptional(models.ToPointer(models.TunnelConfigEncAlgoEnum(plan.EncAlgo.ValueString())))
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func gatewayTemplateTunnelProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelConfigProbe {
	data := models.TunnelConfigProbe{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan, e := NewProbeValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.Interval.ValueInt64Pointer() != nil {
				data.Interval = models.ToPointer(int(plan.Interval.ValueInt64()))
			}
			if plan.Threshold.ValueInt64Pointer() != nil {
				data.Threshold = models.ToPointer(int(plan.Threshold.ValueInt64()))
			}
			if plan.Timeout.ValueInt64Pointer() != nil {
				data.Timeout = models.ToPointer(int(plan.Timeout.ValueInt64()))
			}
			if plan.ProbeType.ValueStringPointer() != nil {
				data.Type = models.ToPointer(models.TunnelConfigProbeTypeEnum(plan.ProbeType.ValueString()))
			}
		}
		return data
	}
}

func gatewayTemplateTunnelPrimaryProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelConfigNode {
	data := models.TunnelConfigNode{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan, e := NewPrimaryValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.Hosts.IsNull() && !plan.Hosts.IsUnknown() {
				data.Hosts = mistutils.ListOfStringTerraformToSdk(plan.Hosts)
			}
			if !plan.InternalIps.IsNull() && !plan.InternalIps.IsUnknown() {
				data.InternalIps = mistutils.ListOfStringTerraformToSdk(plan.InternalIps)
			}
			if !plan.ProbeIps.IsNull() && !plan.ProbeIps.IsUnknown() {
				data.ProbeIps = mistutils.ListOfStringTerraformToSdk(plan.ProbeIps)
			}
			if !plan.RemoteIds.IsNull() && !plan.RemoteIds.IsUnknown() {
				data.RemoteIds = mistutils.ListOfStringTerraformToSdk(plan.RemoteIds)
			}
			if !plan.WanNames.IsNull() && !plan.WanNames.IsUnknown() {
				data.WanNames = mistutils.ListOfStringTerraformToSdk(plan.WanNames)
			}
		}
		return data
	}
}

func gatewayTemplateTunnelSecondaryProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelConfigNode {
	data := models.TunnelConfigNode{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan, e := NewSecondaryValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.Hosts.IsNull() && !plan.Hosts.IsUnknown() {
				data.Hosts = mistutils.ListOfStringTerraformToSdk(plan.Hosts)
			}
			if !plan.InternalIps.IsNull() && !plan.InternalIps.IsUnknown() {
				data.InternalIps = mistutils.ListOfStringTerraformToSdk(plan.InternalIps)
			}
			if !plan.ProbeIps.IsNull() && !plan.ProbeIps.IsUnknown() {
				data.ProbeIps = mistutils.ListOfStringTerraformToSdk(plan.ProbeIps)
			}
			if !plan.RemoteIds.IsNull() && !plan.RemoteIds.IsUnknown() {
				data.RemoteIds = mistutils.ListOfStringTerraformToSdk(plan.RemoteIds)
			}
			if !plan.WanNames.IsNull() && !plan.WanNames.IsUnknown() {
				data.WanNames = mistutils.ListOfStringTerraformToSdk(plan.WanNames)
			}
		}
		return data
	}
}

func tunnelConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.TunnelConfig {
	dataMap := make(map[string]models.TunnelConfig)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(TunnelConfigsValue)

		data := models.TunnelConfig{}

		if !plan.AutoProvision.IsNull() && !plan.AutoProvision.IsUnknown() {
			autoProvision := tunnelConfigsAutoProvisionTerraformToSdk(ctx, diags, plan.AutoProvision)
			data.AutoProvision = &autoProvision
		}

		if plan.IkeLifetime.ValueInt64Pointer() != nil {
			data.IkeLifetime = models.ToPointer(int(plan.IkeLifetime.ValueInt64()))
		}
		if plan.IkeMode.ValueStringPointer() != nil {
			data.IkeMode = models.ToPointer(models.TunnelConfigIkeModeEnum(plan.IkeMode.ValueString()))
		}

		if plan.IpsecLifetime.ValueInt64Pointer() != nil {
			data.IpsecLifetime = models.ToPointer(int(plan.IpsecLifetime.ValueInt64()))
		}

		if !plan.IkeProposals.IsNull() && !plan.IkeProposals.IsUnknown() {
			ikeProposals := gatewayTemplateTunnelIkeProposalTerraformToSdk(plan.IkeProposals)
			data.IkeProposals = ikeProposals
		}

		if plan.LocalId.ValueStringPointer() != nil {
			data.LocalId = models.ToPointer(plan.LocalId.ValueString())
		}

		if !plan.LocalSubnets.IsNull() && !plan.LocalSubnets.IsUnknown() {
			data.LocalSubnets = mistutils.ListOfStringTerraformToSdk(plan.LocalSubnets)
		}

		if plan.Mode.ValueStringPointer() != nil {
			data.Mode = (*models.TunnelConfigTunnelModeEnum)(plan.Mode.ValueStringPointer())
		}

		if !plan.Networks.IsNull() && !plan.Networks.IsUnknown() {
			data.Networks = mistutils.ListOfStringTerraformToSdk(plan.Networks)
		}

		if !plan.Primary.IsNull() && !plan.Primary.IsUnknown() {
			primary := gatewayTemplateTunnelPrimaryProbeTerraformToSdk(ctx, diags, plan.Primary)
			data.Primary = &primary
		}

		if !plan.IpsecProposals.IsNull() && !plan.IpsecProposals.IsUnknown() {
			ipsecProposals := gatewayTemplateTunnelIpsecProposalTerraformToSdk(plan.IpsecProposals)
			data.IpsecProposals = ipsecProposals
		}

		if !plan.Probe.IsNull() && !plan.Probe.IsUnknown() {
			probe := gatewayTemplateTunnelProbeTerraformToSdk(ctx, diags, plan.Probe)
			data.Probe = &probe
		}

		if plan.Protocol.ValueStringPointer() != nil {
			data.Protocol = models.ToPointer(models.TunnelConfigProtocolEnum(plan.Protocol.ValueString()))
		}
		if plan.Provider.ValueStringPointer() != nil {
			data.Provider = models.ToPointer(models.TunnelConfigProviderEnum(plan.Provider.ValueString()))
		}
		if plan.Psk.ValueStringPointer() != nil {
			data.Psk = models.ToPointer(plan.Psk.ValueString())
		}

		if !plan.RemoteSubnets.IsNull() && !plan.RemoteSubnets.IsUnknown() {
			data.RemoteSubnets = mistutils.ListOfStringTerraformToSdk(plan.RemoteSubnets)
		}

		if !plan.Secondary.IsNull() && !plan.Secondary.IsUnknown() {
			secondary := gatewayTemplateTunnelSecondaryProbeTerraformToSdk(ctx, diags, plan.Secondary)
			data.Secondary = &secondary
		}

		if plan.Version.ValueStringPointer() != nil {
			data.Version = models.ToPointer(models.TunnelConfigVersionEnum(plan.Version.ValueString()))
		}

		dataMap[k] = data
	}
	return dataMap
}
