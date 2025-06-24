package resource_org_deviceprofile_ap

import (
	"context"

	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicVlanTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ApPortConfigDynamicVlan {
	data := models.ApPortConfigDynamicVlan{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewDynamicVlanValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {

			if plan.DefaultVlanId.ValueInt64Pointer() != nil {
				data.DefaultVlanId = models.ToPointer(int(plan.DefaultVlanId.ValueInt64()))
			}
			if plan.Enabled.ValueBoolPointer() != nil {
				data.Enabled = plan.Enabled.ValueBoolPointer()
			}
			if plan.DynamicVlanType.ValueStringPointer() != nil {
				data.Type = plan.DynamicVlanType.ValueStringPointer()
			}

			if !plan.Vlans.IsNull() && !plan.Vlans.IsUnknown() {
				vlans := make(map[string]string)
				for key, value := range plan.Vlans.Elements() {
					var valueInterface interface{} = value
					planVlan := valueInterface.(basetypes.StringValue)
					vlans[key] = planVlan.ValueString()
				}
				data.Vlans = vlans
			}
		}
	}
	return &data
}

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.WlanMistNac {
	data := models.WlanMistNac{}

	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewMistNacValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.Enabled.ValueBoolPointer() != nil {
				data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
			}
		}
	}
	return &data
}

func radiusAcctServersTerraformToSdk(d basetypes.ListValue) []models.RadiusAcctServer {

	var data []models.RadiusAcctServer
	for _, planAttr := range d.Elements() {
		var srvPlanInterface interface{} = planAttr
		srvPlan := srvPlanInterface.(AcctServersValue)

		srvData := models.RadiusAcctServer{}
		srvData.Host = srvPlan.Host.ValueString()
		if srvPlan.Port.ValueStringPointer() != nil {
			srvData.Port = models.ToPointer(models.RadiusAcctPortContainer.FromString(srvPlan.Port.ValueString()))
		}
		srvData.Secret = srvPlan.Secret.ValueString()
		if srvPlan.KeywrapEnabled.ValueBoolPointer() != nil {
			srvData.KeywrapEnabled = models.ToPointer(srvPlan.KeywrapEnabled.ValueBool())
		}
		if srvPlan.KeywrapFormat.ValueStringPointer() != nil {
			srvData.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(srvPlan.KeywrapFormat.ValueString()))
		}
		if srvPlan.KeywrapKek.ValueStringPointer() != nil {
			srvData.KeywrapKek = models.ToPointer(srvPlan.KeywrapKek.ValueString())
		}
		if srvPlan.KeywrapMack.ValueStringPointer() != nil {
			srvData.KeywrapMack = models.ToPointer(srvPlan.KeywrapMack.ValueString())
		}
		data = append(data, srvData)
	}
	return data
}

func radiusAuthServersTerraformToSdk(d basetypes.ListValue) []models.RadiusAuthServer {

	var data []models.RadiusAuthServer
	for _, planAttr := range d.Elements() {
		var srvPlanInterface interface{} = planAttr
		srvPlan := srvPlanInterface.(AuthServersValue)

		srvData := models.RadiusAuthServer{}
		srvData.Host = srvPlan.Host.ValueString()
		if srvPlan.Port.ValueStringPointer() != nil {
			srvData.Port = models.ToPointer(models.RadiusAuthPortContainer.FromString(srvPlan.Port.ValueString()))
		}
		srvData.Secret = srvPlan.Secret.ValueString()
		if srvPlan.KeywrapEnabled.ValueBoolPointer() != nil {
			srvData.KeywrapEnabled = models.ToPointer(srvPlan.KeywrapEnabled.ValueBool())
		}
		if srvPlan.KeywrapFormat.ValueStringPointer() != nil {
			srvData.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(srvPlan.KeywrapFormat.ValueString()))
		}
		if srvPlan.KeywrapKek.ValueStringPointer() != nil {
			srvData.KeywrapKek = models.ToPointer(srvPlan.KeywrapKek.ValueString())
		}
		if srvPlan.KeywrapMack.ValueStringPointer() != nil {
			srvData.KeywrapMack = models.ToPointer(srvPlan.KeywrapMack.ValueString())
		}
		if srvPlan.RequireMessageAuthenticator.ValueBoolPointer() != nil {
			srvData.RequireMessageAuthenticator = models.ToPointer(srvPlan.RequireMessageAuthenticator.ValueBool())
		}
		data = append(data, srvData)
	}
	return data
}

func radiusConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RadiusConfig {
	data := models.RadiusConfig{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewRadiusConfigValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.AcctInterimInterval.ValueInt64Pointer() != nil {
				data.AcctInterimInterval = models.ToPointer(int(plan.AcctInterimInterval.ValueInt64()))
			}
			if !plan.AcctServers.IsNull() && !plan.AcctServers.IsUnknown() {
				data.AcctServers = radiusAcctServersTerraformToSdk(plan.AcctServers)
			}
			if !plan.AuthServers.IsNull() && !plan.AuthServers.IsUnknown() {
				data.AuthServers = radiusAuthServersTerraformToSdk(plan.AuthServers)
			}
			if plan.AuthServersRetries.ValueInt64Pointer() != nil {
				data.AuthServersRetries = models.ToPointer(int(plan.AuthServersRetries.ValueInt64()))
			}
			if plan.AuthServersTimeout.ValueInt64Pointer() != nil {
				data.AuthServersTimeout = models.ToPointer(int(plan.AuthServersTimeout.ValueInt64()))
			}
			if plan.CoaEnabled.ValueBoolPointer() != nil {
				data.CoaEnabled = plan.CoaEnabled.ValueBoolPointer()
			}
			if plan.CoaPort.ValueInt64Pointer() != nil {
				data.CoaPort = models.ToPointer(int(plan.CoaPort.ValueInt64()))
			}
			if plan.Network.ValueStringPointer() != nil {
				data.Network = models.ToPointer(plan.Network.ValueString())
			}
			if plan.SourceIp.ValueStringPointer() != nil {
				data.SourceIp = models.ToPointer(plan.SourceIp.ValueString())
			}
		}
	}
	return &data
}

func radsecServersTerraformToSdk(d basetypes.ListValue) []models.RadsecServer {
	var dataList []models.RadsecServer
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(ServersValue)
		data := models.RadsecServer{}
		data.Host = plan.Host.ValueStringPointer()
		data.Port = models.ToPointer(int(plan.Port.ValueInt64()))

		dataList = append(dataList, data)
	}
	return dataList
}

func radsecTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.Radsec {
	data := models.Radsec{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewRadsecValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.CoaEnabled.ValueBoolPointer() != nil {
				data.CoaEnabled = plan.CoaEnabled.ValueBoolPointer()
			}
			if plan.Enabled.ValueBoolPointer() != nil {
				data.Enabled = plan.Enabled.ValueBoolPointer()
			}
			if plan.IdleTimeout.ValueStringPointer() != nil {
				data.IdleTimeout = models.ToPointer(models.RadsecIdleTimeoutContainer.FromString(plan.IdleTimeout.ValueString()))
			}
			if !plan.MxclusterIds.IsNull() && !plan.MxclusterIds.IsUnknown() {
				data.MxclusterIds = mistutils.ListOfUuidTerraformToSdk(plan.MxclusterIds)
			}
			if !plan.ProxyHosts.IsNull() && !plan.ProxyHosts.IsUnknown() {
				data.ProxyHosts = mistutils.ListOfStringTerraformToSdk(plan.ProxyHosts)
			}
			if plan.ServerName.ValueStringPointer() != nil {
				data.ServerName = plan.ServerName.ValueStringPointer()
			}

			servers := radsecServersTerraformToSdk(plan.Servers)
			data.Servers = servers

			data.UseMxedge = plan.UseMxedge.ValueBoolPointer()
			data.UseSiteMxedge = plan.UseSiteMxedge.ValueBoolPointer()
		}
	}
	return &data
}

func portConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.ApPortConfig {
	data := make(map[string]models.ApPortConfig)
	for key, value := range d.Elements() {
		var valueInterface interface{} = value
		plan := valueInterface.(PortConfigValue)

		new := models.ApPortConfig{}
		if plan.Disabled.ValueBoolPointer() != nil {
			new.Disabled = plan.Disabled.ValueBoolPointer()
		}
		if !plan.DynamicVlan.IsNull() && !plan.DynamicVlan.IsUnknown() {
			new.DynamicVlan = dynamicVlanTerraformToSdk(ctx, diags, plan.DynamicVlan)
		}
		if plan.EnableMacAuth.ValueBoolPointer() != nil {
			new.EnableMacAuth = plan.EnableMacAuth.ValueBoolPointer()
		}
		if plan.Forwarding.ValueStringPointer() != nil {
			new.Forwarding = models.ToPointer(models.ApPortConfigForwardingEnum(plan.Forwarding.ValueString()))
		}
		if plan.MacAuthPreferred.ValueBoolPointer() != nil {
			new.MacAuthPreferred = plan.MacAuthPreferred.ValueBoolPointer()
		}
		if plan.MacAuthProtocol.ValueStringPointer() != nil {
			new.MacAuthProtocol = models.ToPointer(models.ApPortConfigMacAuthProtocolEnum(plan.MacAuthProtocol.ValueString()))
		}
		if !plan.MistNac.IsNull() && !plan.MistNac.IsUnknown() {
			new.MistNac = mistNacTerraformToSdk(ctx, diags, plan.MistNac)
		}
		if plan.MxTunnelId.ValueStringPointer() != nil {
			mxTunnelId, e := uuid.Parse(plan.MxTunnelId.ValueString())
			if e == nil {
				new.MxTunnelId = &mxTunnelId
			} else {
				diags.AddError("Bad value for mxtunnel_id", e.Error())
			}
		}
		if plan.MxtunnelName.ValueStringPointer() != nil {
			new.MxtunnelName = plan.MxtunnelName.ValueStringPointer()
		}
		if plan.PortAuth.ValueStringPointer() != nil {
			new.PortAuth = models.ToPointer(models.ApPortConfigPortAuthEnum(plan.PortAuth.ValueString()))
		}
		if plan.PortVlanId.ValueInt64Pointer() != nil {
			new.PortVlanId = models.ToPointer(int(plan.PortVlanId.ValueInt64()))
		}
		if !plan.RadiusConfig.IsNull() && !plan.RadiusConfig.IsUnknown() {
			new.RadiusConfig = radiusConfigTerraformToSdk(ctx, diags, plan.RadiusConfig)
		}
		if !plan.Radsec.IsNull() && !plan.Radsec.IsUnknown() {
			new.Radsec = radsecTerraformToSdk(ctx, diags, plan.Radsec)
		}
		if plan.VlanId.ValueInt64Pointer() != nil {
			new.VlanId = models.ToPointer(int(plan.VlanId.ValueInt64()))
		}
		if !plan.VlanIds.IsNull() && !plan.VlanIds.IsUnknown() {
			new.VlanIds = mistutils.ListOfIntTerraformToSdk(plan.VlanIds)
		}
		if plan.WxtunnelId.ValueStringPointer() != nil {
			wxTunnelId, e := uuid.Parse(plan.WxtunnelId.ValueString())
			if e == nil {
				new.WxtunnelId = &wxTunnelId
			} else {
				diags.AddError("Bad value for wxtunnel_id", e.Error())
			}
		}
		if plan.WxtunnelRemoteId.ValueStringPointer() != nil {
			new.WxtunnelRemoteId = plan.WxtunnelRemoteId.ValueStringPointer()
		}
		data[key] = new
	}
	return data
}
