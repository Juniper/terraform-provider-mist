package resource_device_gateway

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func tunnelConfigsAutoProvisionPrimaryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.TunnelConfigsAutoProvisionNode {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionPrimaryTerraformToSdk")
	data := models.TunnelConfigsAutoProvisionNode{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewAutoProvisionPrimaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.NumHosts.ValueStringPointer() != nil {
			data.NumHosts = models.ToPointer(plan.NumHosts.ValueString())
		}
		if !plan.WanNames.IsNull() && !plan.WanNames.IsUnknown() {
			data.WanNames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames)
		}

		return &data
	}
}

func tunnelConfigsAutoProvisionSecondaryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.TunnelConfigsAutoProvisionNode {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionSecondaryTerraformToSdk")
	data := models.TunnelConfigsAutoProvisionNode{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewAutoProvisionSecondaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.NumHosts.ValueStringPointer() != nil {
			data.NumHosts = models.ToPointer(plan.NumHosts.ValueString())
		}
		if !plan.WanNames.IsNull() && !plan.WanNames.IsUnknown() {
			data.WanNames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames)
		}

		return &data
	}
}

func tunnelConfigsAutoProvisionTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelConfigsAutoProvision {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionTerraformToSdk")
	data := models.TunnelConfigsAutoProvision{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAutoProvisionValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.Enable.ValueBoolPointer() != nil {
			data.Enable = models.ToPointer(plan.Enable.ValueBool())
		}

		var plan_latlng_interface interface{} = plan.Latlng
		plan_latlng := plan_latlng_interface.(LatlngValue)

		var latlng models.LatLng
		latlng.Lat = plan_latlng.Lng.ValueFloat64()
		latlng.Lng = plan_latlng.Lng.ValueFloat64()
		if !plan.Latlng.IsNull() && !plan.Latlng.IsUnknown() {
			data.Latlng = models.ToPointer(latlng)
		}

		if !plan.AutoProvisionPrimary.IsNull() && !plan.AutoProvisionPrimary.IsUnknown() {
			data.Primary = tunnelConfigsAutoProvisionPrimaryTerraformToSdk(ctx, diags, plan.AutoProvisionPrimary)
		}

		if plan.Region.ValueStringPointer() != nil {
			data.Region = models.ToPointer(models.TunnelConfigsAutoProvisionRegionEnum(plan.Region.ValueString()))
		}

		if !plan.AutoProvisionSecondary.IsNull() && !plan.AutoProvisionSecondary.IsUnknown() {
			data.Secondary = tunnelConfigsAutoProvisionSecondaryTerraformToSdk(ctx, diags, plan.AutoProvisionSecondary)
		}

		return data
	}
}

func gatewayTemplateTunnelIkeProposalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.GatewayTemplateTunnelIkeProposal {
	tflog.Debug(ctx, "gatewayTemplateTunnelIkeProposalTerraformToSdk")
	var data_list []models.GatewayTemplateTunnelIkeProposal
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IkeProposalsValue)
		data := models.GatewayTemplateTunnelIkeProposal{}
		if plan.AuthAlgo.ValueStringPointer() != nil {
			data.AuthAlgo = models.ToPointer(models.TunnelConfigsAuthAlgoEnum(plan.AuthAlgo.ValueString()))
		}
		if plan.DhGroup.ValueStringPointer() != nil {
			data.DhGroup = models.ToPointer(models.GatewayTemplateTunnelIkeDhGroupEnum(plan.DhGroup.ValueString()))
		}
		if plan.EncAlgo.ValueStringPointer() != nil {
			data.EncAlgo = models.NewOptional(models.ToPointer(models.TunnelConfigsEncAlgoEnum(plan.EncAlgo.ValueString())))
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func gatewayTemplateTunnelIpsecProposalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.GatewayTemplateTunnelIpsecProposal {
	tflog.Debug(ctx, "gatewayTemplateTunnelIpsecProposalTerraformToSdk")
	var data_list []models.GatewayTemplateTunnelIpsecProposal
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IpsecProposalsValue)
		data := models.GatewayTemplateTunnelIpsecProposal{}
		if plan.AuthAlgo.ValueStringPointer() != nil {
			data.AuthAlgo = models.ToPointer(models.TunnelConfigsAuthAlgoEnum(plan.AuthAlgo.ValueString()))
		}
		if plan.DhGroup.ValueStringPointer() != nil {
			data.DhGroup = models.ToPointer(models.TunnelConfigsDhGroupEnum(plan.DhGroup.ValueString()))
		}
		if plan.EncAlgo.ValueStringPointer() != nil {
			data.EncAlgo = models.NewOptional(models.ToPointer(models.TunnelConfigsEncAlgoEnum(plan.EncAlgo.ValueString())))
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func gatewayTemplateTunnelProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.GatewayTemplateTunnelProbe {
	tflog.Debug(ctx, "gatewayTemplateTunnelProbeTerraformToSdk")
	data := models.GatewayTemplateTunnelProbe{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewProbeValueMust(d.AttributeTypes(ctx), d.Attributes())
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
			data.Type = models.ToPointer(models.GatewayTemplateProbeTypeEnum(plan.ProbeType.ValueString()))
		}
		return data
	}
}

func gatewayTemplateTunnelPrimaryProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.GatewayTemplateTunnelNode {
	tflog.Debug(ctx, "gatewayTemplateTunnelPrimaryProbeTerraformToSdk")
	data := models.GatewayTemplateTunnelNode{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewPrimaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		if !plan.Hosts.IsNull() && !plan.Hosts.IsUnknown() {
			data.Hosts = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hosts)
		}
		if !plan.InternalIps.IsNull() && !plan.InternalIps.IsUnknown() {
			data.InternalIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.InternalIps)
		}
		if !plan.ProbeIps.IsNull() && !plan.ProbeIps.IsUnknown() {
			data.ProbeIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ProbeIps)
		}
		if !plan.RemoteIds.IsNull() && !plan.RemoteIds.IsUnknown() {
			data.RemoteIds = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RemoteIds)
		}
		if !plan.WanNames.IsNull() && !plan.WanNames.IsUnknown() {
			data.WanNames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames)
		}
		return data
	}
}

func gatewayTemplateTunnelSecondaryProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.GatewayTemplateTunnelNode {
	tflog.Debug(ctx, "gatewayTemplateTunnelSecondaryProbeTerraformToSdk")
	data := models.GatewayTemplateTunnelNode{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewSecondaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		if !plan.Hosts.IsNull() && !plan.Hosts.IsUnknown() {
			data.Hosts = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hosts)
		}
		if !plan.InternalIps.IsNull() && !plan.InternalIps.IsUnknown() {
			data.InternalIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.InternalIps)
		}
		if !plan.ProbeIps.IsNull() && !plan.ProbeIps.IsUnknown() {
			data.ProbeIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ProbeIps)
		}
		if !plan.RemoteIds.IsNull() && !plan.RemoteIds.IsUnknown() {
			data.RemoteIds = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RemoteIds)
		}
		if !plan.WanNames.IsNull() && !plan.WanNames.IsUnknown() {
			data.WanNames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames)
		}
		return data
	}
}

func tunnelConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.TunnelConfigs {
	tflog.Debug(ctx, "tunnelConfigsTerraformToSdk")
	data_map := make(map[string]models.TunnelConfigs)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TunnelConfigsValue)

		data := models.TunnelConfigs{}

		auto_provision := tunnelConfigsAutoProvisionTerraformToSdk(ctx, diags, plan.AutoProvision)
		if !plan.AutoProvision.IsNull() && !plan.AutoProvision.IsUnknown() {
			data.AutoProvision = &auto_provision
		}

		if plan.IkeLifetime.ValueInt64Pointer() != nil {
			data.IkeLifetime = models.ToPointer(int(plan.IkeLifetime.ValueInt64()))
		}
		if plan.IkeMode.ValueStringPointer() != nil {
			data.IkeMode = models.ToPointer(models.GatewayTemplateTunnelIkeModeEnum(plan.IkeMode.ValueString()))
		}

		ike_proposals := gatewayTemplateTunnelIkeProposalTerraformToSdk(ctx, diags, plan.IkeProposals)
		if !plan.IkeProposals.IsNull() && !plan.IkeProposals.IsUnknown() {
			data.IkeProposals = ike_proposals
		}

		if plan.IpsecLifetime.ValueInt64Pointer() != nil {
			data.IpsecLifetime = models.ToPointer(int(plan.IpsecLifetime.ValueInt64()))
		}

		primary := gatewayTemplateTunnelPrimaryProbeTerraformToSdk(ctx, diags, plan.Primary)
		if !plan.Primary.IsNull() && !plan.Primary.IsUnknown() {
			data.Primary = &primary
		}

		ipsec_proposals := gatewayTemplateTunnelIpsecProposalTerraformToSdk(ctx, diags, plan.IpsecProposals)
		if !plan.IpsecProposals.IsNull() && !plan.IpsecProposals.IsUnknown() {
			data.IpsecProposals = ipsec_proposals
		}

		if plan.LocalId.ValueStringPointer() != nil {
			data.LocalId = models.ToPointer(plan.LocalId.ValueString())
		}

		probe := gatewayTemplateTunnelProbeTerraformToSdk(ctx, diags, plan.Probe)
		if !plan.Probe.IsNull() && !plan.Probe.IsUnknown() {
			data.Probe = &probe
		}

		if plan.Protocol.ValueStringPointer() != nil {
			data.Protocol = models.ToPointer(models.GatewayTemplateTunnelProtocolEnum(plan.Protocol.ValueString()))
		}
		if plan.Provider.ValueStringPointer() != nil {
			data.Provider = models.ToPointer(models.TunnelProviderOptionsNameEnum(plan.Provider.ValueString()))
		}
		if plan.Psk.ValueStringPointer() != nil {
			data.Psk = models.ToPointer(plan.Psk.ValueString())
		}

		secondary := gatewayTemplateTunnelSecondaryProbeTerraformToSdk(ctx, diags, plan.Secondary)
		if !plan.Secondary.IsNull() && !plan.Secondary.IsUnknown() {
			data.Secondary = &secondary
		}

		if plan.Version.ValueStringPointer() != nil {
			data.Version = models.ToPointer(models.GatewayTemplateTunnelVersionEnum(plan.Version.ValueString()))
		}

		data_map[k] = data
	}
	return data_map
}
