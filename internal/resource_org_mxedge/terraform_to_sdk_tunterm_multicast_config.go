package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermMulticastConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d TuntermMulticastConfigValue) *models.MxedgeTuntermMulticastConfig {
	data := models.MxedgeTuntermMulticastConfig{}

	if !d.Mdns.IsNull() && !d.Mdns.IsUnknown() {
		var mdns MdnsValue
		d.Mdns.As(ctx, &mdns, basetypes.ObjectAsOptions{})
		data.Mdns = mdnsTerraformToSdk(ctx, diags, mdns)
	}

	if !d.Ssdp.IsNull() && !d.Ssdp.IsUnknown() {
		var ssdp SsdpValue
		d.Ssdp.As(ctx, &ssdp, basetypes.ObjectAsOptions{})
		data.Ssdp = ssdpTerraformToSdk(ctx, diags, ssdp)
	}

	return &data
}

func mdnsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MdnsValue) *models.MxedgeTuntermMulticastMdns {
	data := models.MxedgeTuntermMulticastMdns{}

	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if !d.VlanIds.IsNull() && !d.VlanIds.IsUnknown() {
		data.VlanIds = mistutils.ListOfStringTerraformToSdk(d.VlanIds)
	}

	return &data
}

func ssdpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SsdpValue) *models.MxedgeTuntermMulticastSsdp {
	data := models.MxedgeTuntermMulticastSsdp{}

	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if !d.VlanIds.IsNull() && !d.VlanIds.IsUnknown() {
		data.VlanIds = mistutils.ListOfStringTerraformToSdk(d.VlanIds)
	}

	return &data
}
