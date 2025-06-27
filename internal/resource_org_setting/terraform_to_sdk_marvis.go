package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func marvisAutoOperationTerraformToSdk(ctx context.Context, o basetypes.ObjectValue) *models.MarvisAutoOperations {
	data := models.MarvisAutoOperations{}
	if o.IsNull() || o.IsUnknown() {
		return &data
	} else {
		d := NewAutoOperationsValueMust(o.AttributeTypes(ctx), o.Attributes())

		if !d.BouncePortForAbnormalPoeClient.IsNull() && !d.BouncePortForAbnormalPoeClient.IsUnknown() {
			data.BouncePortForAbnormalPoeClient = models.ToPointer(d.BouncePortForAbnormalPoeClient.ValueBool())
		}
		if !d.DisablePortWhenDdosProtocolViolation.IsNull() && !d.DisablePortWhenDdosProtocolViolation.IsUnknown() {
			data.DisablePortWhenDdosProtocolViolation = models.ToPointer(d.DisablePortWhenDdosProtocolViolation.ValueBool())
		}
		if !d.DisablePortWhenRogueDhcpServerDetected.IsNull() && !d.DisablePortWhenRogueDhcpServerDetected.IsUnknown() {
			data.DisablePortWhenRogueDhcpServerDetected = models.ToPointer(d.DisablePortWhenRogueDhcpServerDetected.ValueBool())
		}

		return &data
	}
}

func marvisTerraformToSdk(ctx context.Context, d MarvisValue) *models.Marvis {
	data := models.Marvis{}
	if !d.AutoOperations.IsNull() && !d.AutoOperations.IsUnknown() {
		data.AutoOperations = marvisAutoOperationTerraformToSdk(ctx, d.AutoOperations)
	}
	return &data
}
