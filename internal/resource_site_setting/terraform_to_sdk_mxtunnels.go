package resource_site_setting

import (
	"context"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func mxtunnelTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MxtunnelValue) *models.SiteMxtunnel {
	data := models.SiteMxtunnel{}

	if !d.ApSubnets.IsNull() && !d.ApSubnets.IsUnknown() {
		data.ApSubnets = mistutils.ListOfStringTerraformToSdk(d.ApSubnets)
	}

	if !d.AutoPreemption.IsNull() && !d.AutoPreemption.IsUnknown() {
		attrs := d.AutoPreemption.Attributes()
		ap := models.AutoPreemption{}
		if v, ok := attrs["day_of_week"].(basetypes.StringValue); ok && !v.IsNull() && !v.IsUnknown() && v.ValueString() != "" {
			ap.DayOfWeek = (*models.DayOfWeekEnum)(v.ValueStringPointer())
		}
		if v, ok := attrs["enabled"].(basetypes.BoolValue); ok && !v.IsNull() && !v.IsUnknown() {
			ap.Enabled = v.ValueBoolPointer()
		}
		if v, ok := attrs["time_of_day"].(basetypes.StringValue); ok && !v.IsNull() && !v.IsUnknown() {
			ap.TimeOfDay = v.ValueStringPointer()
		}
		data.AutoPreemption = &ap
	}

	if !d.Clusters.IsNull() && !d.Clusters.IsUnknown() {
		var clusters []models.SiteMxtunnelCluster
		for _, item := range d.Clusters.Elements() {
			itemValue := item.(ClustersValue)
			cluster := models.SiteMxtunnelCluster{}
			if !itemValue.Name.IsNull() && !itemValue.Name.IsUnknown() {
				cluster.Name = itemValue.Name.ValueStringPointer()
			}
			if !itemValue.TuntermHosts.IsNull() && !itemValue.TuntermHosts.IsUnknown() {
				cluster.TuntermHosts = mistutils.ListOfStringTerraformToSdk(itemValue.TuntermHosts)
			}
			clusters = append(clusters, cluster)
		}
		data.Clusters = clusters
	}

	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if !d.HelloInterval.IsNull() && !d.HelloInterval.IsUnknown() {
		data.HelloInterval = models.ToPointer(int(d.HelloInterval.ValueInt64()))
	}

	if !d.HelloRetries.IsNull() && !d.HelloRetries.IsUnknown() {
		data.HelloRetries = models.ToPointer(int(d.HelloRetries.ValueInt64()))
	}

	if !d.Hosts.IsNull() && !d.Hosts.IsUnknown() {
		data.Hosts = mistutils.ListOfStringTerraformToSdk(d.Hosts)
	}

	if !d.Mtu.IsNull() && !d.Mtu.IsUnknown() {
		data.Mtu = models.ToPointer(int(d.Mtu.ValueInt64()))
	}

	if !d.Protocol.IsNull() && !d.Protocol.IsUnknown() && d.Protocol.ValueString() != "" {
		p := models.MxtunnelProtocolEnum(d.Protocol.ValueString())
		data.Protocol = &p
	}

	if !d.VlanIds.IsNull() && !d.VlanIds.IsUnknown() {
		data.VlanIds = mistutils.ListOfIntTerraformToSdk(d.VlanIds)
	}

	if !d.Radsec.IsNull() && !d.Radsec.IsUnknown() {
		data.Radsec = mxtunnelsRadsecTerraformToSdk(ctx, diags, d.Radsec)
	}

	if !d.AdditionalMxtunnels.IsNull() && !d.AdditionalMxtunnels.IsUnknown() {
		additionals := make(map[string]models.SiteMxtunnelAdditionalMxtunnel)
		for k, item := range d.AdditionalMxtunnels.Elements() {
			itemValue := item.(AdditionalMxtunnelsValue)
			additional := mxtunnelsAdditionalTerraformToSdk(ctx, diags, itemValue)
			additionals[k] = additional
		}
		data.AdditionalMxtunnels = additionals
	}

	return &data
}

func mxtunnelsAdditionalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AdditionalMxtunnelsValue) models.SiteMxtunnelAdditionalMxtunnel {
	data := models.SiteMxtunnelAdditionalMxtunnel{}

	if !d.HelloInterval.IsNull() && !d.HelloInterval.IsUnknown() {
		data.HelloInterval = models.ToPointer(int(d.HelloInterval.ValueInt64()))
	}
	if !d.HelloRetries.IsNull() && !d.HelloRetries.IsUnknown() {
		data.HelloRetries = models.ToPointer(int(d.HelloRetries.ValueInt64()))
	}
	if !d.Protocol.IsNull() && !d.Protocol.IsUnknown() && d.Protocol.ValueString() != "" {
		data.Protocol = (*models.SiteMxtunnelProtocolEnum)(d.Protocol.ValueStringPointer())
	}
	if !d.VlanIds.IsNull() && !d.VlanIds.IsUnknown() {
		data.VlanIds = mistutils.ListOfIntTerraformToSdk(d.VlanIds)
	}
	if !d.TuntermClusters.IsNull() && !d.TuntermClusters.IsUnknown() {
		var clusters []models.SiteMxtunnelCluster
		for _, item := range d.TuntermClusters.Elements() {
			itemValue := item.(TuntermClustersValue)
			cluster := models.SiteMxtunnelCluster{}
			if !itemValue.Name.IsNull() && !itemValue.Name.IsUnknown() {
				cluster.Name = itemValue.Name.ValueStringPointer()
			}
			if !itemValue.TuntermHosts.IsNull() && !itemValue.TuntermHosts.IsUnknown() {
				cluster.TuntermHosts = mistutils.ListOfStringTerraformToSdk(itemValue.TuntermHosts)
			}
			clusters = append(clusters, cluster)
		}
		data.Clusters = clusters
	}

	return data
}

func mxtunnelsRadsecTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, obj types.Object) *models.SiteMxtunnelRadsec {
	data := models.SiteMxtunnelRadsec{}
	if obj.IsNull() || obj.IsUnknown() {
		return &data
	}

	attrs := obj.Attributes()
	if v, ok := attrs["enabled"].(basetypes.BoolValue); ok && !v.IsNull() && !v.IsUnknown() {
		data.Enabled = v.ValueBoolPointer()
	}
	if v, ok := attrs["use_mxedge"].(basetypes.BoolValue); ok && !v.IsNull() && !v.IsUnknown() {
		data.UseMxedge = v.ValueBoolPointer()
	}

	if acctList, ok := attrs["acct_servers"].(basetypes.ListValue); ok && !acctList.IsNull() && !acctList.IsUnknown() {
		var servers []models.RadiusAcctServer
		for _, item := range acctList.Elements() {
			itemValue := item.(AcctServersValue)
			server := models.RadiusAcctServer{Host: itemValue.Host.ValueString()}
			if !itemValue.Secret.IsNull() && !itemValue.Secret.IsUnknown() {
				server.Secret = itemValue.Secret.ValueString()
			}
			if !itemValue.Port.IsNull() && !itemValue.Port.IsUnknown() && itemValue.Port.ValueString() != "" {
				portStr := itemValue.Port.ValueString()
				if portInt, err := strconv.Atoi(portStr); err == nil {
					server.Port = models.ToPointer(models.RadiusAcctPortContainer.FromNumber(portInt))
				} else {
					server.Port = models.ToPointer(models.RadiusAcctPortContainer.FromString(portStr))
				}
			}
			servers = append(servers, server)
		}
		data.AcctServers = servers
	}

	if authList, ok := attrs["auth_servers"].(basetypes.ListValue); ok && !authList.IsNull() && !authList.IsUnknown() {
		var servers []models.RadiusAuthServer
		for _, item := range authList.Elements() {
			itemValue := item.(AuthServersValue)
			server := models.RadiusAuthServer{Host: itemValue.Host.ValueString()}
			if !itemValue.Secret.IsNull() && !itemValue.Secret.IsUnknown() {
				server.Secret = itemValue.Secret.ValueString()
			}
			if !itemValue.KeywrapEnabled.IsNull() && !itemValue.KeywrapEnabled.IsUnknown() {
				server.KeywrapEnabled = itemValue.KeywrapEnabled.ValueBoolPointer()
			}
			if !itemValue.KeywrapFormat.IsNull() && !itemValue.KeywrapFormat.IsUnknown() && itemValue.KeywrapFormat.ValueString() != "" {
				server.KeywrapFormat = (*models.RadiusKeywrapFormatEnum)(itemValue.KeywrapFormat.ValueStringPointer())
			}
			if !itemValue.KeywrapKek.IsNull() && !itemValue.KeywrapKek.IsUnknown() {
				server.KeywrapKek = itemValue.KeywrapKek.ValueStringPointer()
			}
			if !itemValue.KeywrapMack.IsNull() && !itemValue.KeywrapMack.IsUnknown() {
				server.KeywrapMack = itemValue.KeywrapMack.ValueStringPointer()
			}
			if !itemValue.RequireMessageAuthenticator.IsNull() && !itemValue.RequireMessageAuthenticator.IsUnknown() {
				server.RequireMessageAuthenticator = itemValue.RequireMessageAuthenticator.ValueBoolPointer()
			}
			if !itemValue.Port.IsNull() && !itemValue.Port.IsUnknown() && itemValue.Port.ValueString() != "" {
				portStr := itemValue.Port.ValueString()
				if portInt, err := strconv.Atoi(portStr); err == nil {
					server.Port = models.ToPointer(models.RadiusAuthPortContainer.FromNumber(portInt))
				} else {
					server.Port = models.ToPointer(models.RadiusAuthPortContainer.FromString(portStr))
				}
			}
			servers = append(servers, server)
		}
		data.AuthServers = servers
	}

	return &data
}
