package resource_org_mxcluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func mistDasTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistDasValue) *models.MxedgeDas {
	data := models.MxedgeDas{}

	if !d.CoaServers.IsNull() && !d.CoaServers.IsUnknown() {
		var coaServersList []models.MxedgeDasCoaServer
		for _, item := range d.CoaServers.Elements() {
			itemValue := item.(CoaServersValue)

			coaServer := models.MxedgeDasCoaServer{}

			if !itemValue.DisableEventTimestampCheck.IsNull() && !itemValue.DisableEventTimestampCheck.IsUnknown() {
				coaServer.DisableEventTimestampCheck = itemValue.DisableEventTimestampCheck.ValueBoolPointer()
			}

			if !itemValue.Enabled.IsNull() && !itemValue.Enabled.IsUnknown() {
				coaServer.Enabled = itemValue.Enabled.ValueBoolPointer()
			}

			if !itemValue.Host.IsNull() && !itemValue.Host.IsUnknown() {
				coaServer.Host = itemValue.Host.ValueStringPointer()
			}

			if !itemValue.Port.IsNull() && !itemValue.Port.IsUnknown() {
				coaServer.Port = models.ToPointer(int(itemValue.Port.ValueInt64()))
			}

			if !itemValue.RequireMessageAuthenticator.IsNull() && !itemValue.RequireMessageAuthenticator.IsUnknown() {
				coaServer.RequireMessageAuthenticator = itemValue.RequireMessageAuthenticator.ValueBoolPointer()
			}

			if !itemValue.Secret.IsNull() && !itemValue.Secret.IsUnknown() {
				coaServer.Secret = itemValue.Secret.ValueStringPointer()
			}

			coaServersList = append(coaServersList, coaServer)
		}
		data.CoaServers = coaServersList
	}

	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
