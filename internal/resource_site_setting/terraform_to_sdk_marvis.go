package resource_site_setting

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

		if !d.ApInsufficientCapacity.IsNull() && !d.ApInsufficientCapacity.IsUnknown() {
			data.ApInsufficientCapacity = models.ToPointer(d.ApInsufficientCapacity.ValueBool())
		}
		if !d.ApLoop.IsNull() && !d.ApLoop.IsUnknown() {
			data.ApLoop = models.ToPointer(d.ApLoop.ValueBool())
		}
		if !d.ApNonCompliant.IsNull() && !d.ApNonCompliant.IsUnknown() {
			data.ApNonCompliant = models.ToPointer(d.ApNonCompliant.ValueBool())
		}
		if !d.BouncePortForAbnormalPoeClient.IsNull() && !d.BouncePortForAbnormalPoeClient.IsUnknown() {
			data.BouncePortForAbnormalPoeClient = models.ToPointer(d.BouncePortForAbnormalPoeClient.ValueBool())
		}
		if !d.DisablePortWhenDdosProtocolViolation.IsNull() && !d.DisablePortWhenDdosProtocolViolation.IsUnknown() {
			data.DisablePortWhenDdosProtocolViolation = models.ToPointer(d.DisablePortWhenDdosProtocolViolation.ValueBool())
		}
		if !d.DisablePortWhenRogueDhcpServerDetected.IsNull() && !d.DisablePortWhenRogueDhcpServerDetected.IsUnknown() {
			data.DisablePortWhenRogueDhcpServerDetected = models.ToPointer(d.DisablePortWhenRogueDhcpServerDetected.ValueBool())
		}
		if !d.GatewayNonCompliant.IsNull() && !d.GatewayNonCompliant.IsUnknown() {
			data.GatewayNonCompliant = models.ToPointer(d.GatewayNonCompliant.ValueBool())
		}
		if !d.SwitchMisconfiguredPort.IsNull() && !d.SwitchMisconfiguredPort.IsUnknown() {
			data.SwitchMisconfiguredPort = models.ToPointer(d.SwitchMisconfiguredPort.ValueBool())
		}
		if !d.SwitchPortStuck.IsNull() && !d.SwitchPortStuck.IsUnknown() {
			data.SwitchPortStuck = models.ToPointer(d.SwitchPortStuck.ValueBool())
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
