package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func tuntermPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d TuntermPortConfigValue) *models.TuntermPortConfig {
	data := models.TuntermPortConfig{}

	if !d.DownstreamPorts.IsNull() && !d.DownstreamPorts.IsUnknown() {
		data.DownstreamPorts = mistutils.ListOfStringTerraformToSdk(d.DownstreamPorts)
	}

	if !d.SeparateUpstreamDownstream.IsNull() && !d.SeparateUpstreamDownstream.IsUnknown() {
		data.SeparateUpstreamDownstream = d.SeparateUpstreamDownstream.ValueBoolPointer()
	}

	if !d.UpstreamPortVlanId.IsNull() && !d.UpstreamPortVlanId.IsUnknown() {
		data.UpstreamPortVlanId = models.ToPointer(int(d.UpstreamPortVlanId.ValueInt64()))
	}

	if !d.UpstreamPorts.IsNull() && !d.UpstreamPorts.IsUnknown() {
		data.UpstreamPorts = mistutils.ListOfStringTerraformToSdk(d.UpstreamPorts)
	}

	return &data
}
