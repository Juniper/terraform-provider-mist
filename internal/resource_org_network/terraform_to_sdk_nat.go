package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func sourceNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.NetworkSourceNat {
	data := models.NetworkSourceNat{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewSourceNatValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.ExternalIp = plan.ExternalIp.ValueStringPointer()
		}
	}
	return &data
}
