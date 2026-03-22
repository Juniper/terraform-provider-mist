package resource_org_mxcluster

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func radsecTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RadsecValue) *models.MxclusterRadsec {
	data := models.MxclusterRadsec{}

	if !d.AcctServers.IsNull() && !d.AcctServers.IsUnknown() {
		var acctServersList []models.MxclusterRadsecAcctServer
		for _, item := range d.AcctServers.Elements() {
			itemValue := item.(AcctServersValue)

			acctServer := models.MxclusterRadsecAcctServer{}

			if !itemValue.Host.IsNull() && !itemValue.Host.IsUnknown() {
				acctServer.Host = itemValue.Host.ValueStringPointer()
			}

			if !itemValue.Port.IsNull() && !itemValue.Port.IsUnknown() {
				acctServer.Port = models.ToPointer(int(itemValue.Port.ValueInt64()))
			}

			if !itemValue.Secret.IsNull() && !itemValue.Secret.IsUnknown() {
				acctServer.Secret = itemValue.Secret.ValueStringPointer()
			}

			if !itemValue.Ssids.IsNull() && !itemValue.Ssids.IsUnknown() {
				acctServer.Ssids = mistutils.ListOfStringTerraformToSdk(itemValue.Ssids)
			}

			acctServersList = append(acctServersList, acctServer)
		}
		data.AcctServers = acctServersList
	}

	if !d.AuthServers.IsNull() && !d.AuthServers.IsUnknown() {
		var authServersList []models.MxclusterRadsecAuthServer
		for _, item := range d.AuthServers.Elements() {
			itemValue := item.(AuthServersValue)

			authServer := models.MxclusterRadsecAuthServer{}

			if !itemValue.Host.IsNull() && !itemValue.Host.IsUnknown() {
				authServer.Host = itemValue.Host.ValueStringPointer()
			}

			if !itemValue.InbandStatusCheck.IsNull() && !itemValue.InbandStatusCheck.IsUnknown() {
				authServer.InbandStatusCheck = itemValue.InbandStatusCheck.ValueBoolPointer()
			}

			if !itemValue.InbandStatusInterval.IsNull() && !itemValue.InbandStatusInterval.IsUnknown() {
				authServer.InbandStatusInterval = models.ToPointer(int(itemValue.InbandStatusInterval.ValueInt64()))
			}

			if !itemValue.KeywrapEnabled.IsNull() && !itemValue.KeywrapEnabled.IsUnknown() {
				authServer.KeywrapEnabled = itemValue.KeywrapEnabled.ValueBoolPointer()
			}

			if !itemValue.KeywrapFormat.IsNull() && !itemValue.KeywrapFormat.IsUnknown() {
				authServer.KeywrapFormat = models.NewOptional(models.ToPointer(models.MxclusterRadAuthServerKeywrapFormatEnum(itemValue.KeywrapFormat.ValueString())))
			}

			if !itemValue.KeywrapKek.IsNull() && !itemValue.KeywrapKek.IsUnknown() {
				authServer.KeywrapKek = itemValue.KeywrapKek.ValueStringPointer()
			}

			if !itemValue.KeywrapMack.IsNull() && !itemValue.KeywrapMack.IsUnknown() {
				authServer.KeywrapMack = itemValue.KeywrapMack.ValueStringPointer()
			}

			if !itemValue.Port.IsNull() && !itemValue.Port.IsUnknown() {
				authServer.Port = models.ToPointer(int(itemValue.Port.ValueInt64()))
			}

			if !itemValue.Retry.IsNull() && !itemValue.Retry.IsUnknown() {
				authServer.Retry = models.ToPointer(int(itemValue.Retry.ValueInt64()))
			}

			if !itemValue.Secret.IsNull() && !itemValue.Secret.IsUnknown() {
				authServer.Secret = itemValue.Secret.ValueStringPointer()
			}

			if !itemValue.Ssids.IsNull() && !itemValue.Ssids.IsUnknown() {
				authServer.Ssids = mistutils.ListOfStringTerraformToSdk(itemValue.Ssids)
			}

			if !itemValue.Timeout.IsNull() && !itemValue.Timeout.IsUnknown() {
				authServer.Timeout = models.ToPointer(int(itemValue.Timeout.ValueInt64()))
			}

			authServersList = append(authServersList, authServer)
		}
		data.AuthServers = authServersList
	}

	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if !d.MatchSsid.IsNull() && !d.MatchSsid.IsUnknown() {
		data.MatchSsid = d.MatchSsid.ValueBoolPointer()
	}

	if !d.NasIpSource.IsNull() && !d.NasIpSource.IsUnknown() {
		data.NasIpSource = (*models.MxclusterRadsecNasIpSourceEnum)(d.NasIpSource.ValueStringPointer())
	}

	if !d.ProxyHosts.IsNull() && !d.ProxyHosts.IsUnknown() {
		data.ProxyHosts = mistutils.ListOfStringTerraformToSdk(d.ProxyHosts)
	}

	if !d.ServerSelection.IsNull() && !d.ServerSelection.IsUnknown() {
		data.ServerSelection = (*models.MxclusterRadsecServerSelectionEnum)(d.ServerSelection.ValueStringPointer())
	}

	if !d.SrcIpSource.IsNull() && !d.SrcIpSource.IsUnknown() {
		data.SrcIpSource = (*models.MxclusterRadsecSrcIpSourceEnum)(d.SrcIpSource.ValueStringPointer())
	}

	return &data
}
