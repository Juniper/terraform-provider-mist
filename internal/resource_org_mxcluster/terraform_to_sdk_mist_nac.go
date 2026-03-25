package resource_org_mxcluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) *models.MxclusterNac {
	data := models.MxclusterNac{}

	if !d.AcctServerPort.IsNull() && !d.AcctServerPort.IsUnknown() {
		data.AcctServerPort = models.ToPointer(int(d.AcctServerPort.ValueInt64()))
	}

	if !d.AuthServerPort.IsNull() && !d.AuthServerPort.IsUnknown() {
		data.AuthServerPort = models.ToPointer(int(d.AuthServerPort.ValueInt64()))
	}

	if !d.ClientIps.IsNull() && !d.ClientIps.IsUnknown() {
		clientIpsMap := make(map[string]models.MxclusterNacClientIp)
		var tfMap map[string]ClientIpsValue
		d.ClientIps.ElementsAs(ctx, &tfMap, false)

		for key := range tfMap {
			// ClientIpsValue has no fields based on the schema, it's an empty object
			clientIpsMap[key] = models.MxclusterNacClientIp{}
		}
		data.ClientIps = clientIpsMap
	}

	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if !d.Secret.IsNull() && !d.Secret.IsUnknown() {
		data.Secret = d.Secret.ValueStringPointer()
	}

	return &data
}
