package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func portChannelTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d PortChannelizationValue) *models.PortChannelization {
	data := models.PortChannelization{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	for k, v := range d.Config.Elements() {
		data.AdditionalProperties[k] = v
	}

	return &data
}
