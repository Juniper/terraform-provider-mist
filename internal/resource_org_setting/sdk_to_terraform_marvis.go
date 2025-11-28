package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func marvisAutoOperationsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MarvisAutoOperations) basetypes.ObjectValue {
	var apInsufficientCapacity basetypes.BoolValue
	var apLoop basetypes.BoolValue
	var apNonCompliant basetypes.BoolValue
	var bouncePortForAbnormalPoeClient basetypes.BoolValue
	var disablePortWhenDdosProtocolViolation basetypes.BoolValue
	var disablePortWhenRogueDhcpServerDetected basetypes.BoolValue
	var gatewayNonCompliant basetypes.BoolValue
	var switchMisconfiguredPort basetypes.BoolValue
	var switchPortStuck basetypes.BoolValue

	if d.ApInsufficientCapacity != nil {
		apInsufficientCapacity = types.BoolValue(*d.ApInsufficientCapacity)
	}
	if d.ApLoop != nil {
		apLoop = types.BoolValue(*d.ApLoop)
	}
	if d.ApNonCompliant != nil {
		apNonCompliant = types.BoolValue(*d.ApNonCompliant)
	}
	if d.BouncePortForAbnormalPoeClient != nil {
		bouncePortForAbnormalPoeClient = types.BoolValue(*d.BouncePortForAbnormalPoeClient)
	}
	if d.DisablePortWhenDdosProtocolViolation != nil {
		disablePortWhenDdosProtocolViolation = types.BoolValue(*d.DisablePortWhenDdosProtocolViolation)
	}
	if d.DisablePortWhenRogueDhcpServerDetected != nil {
		disablePortWhenRogueDhcpServerDetected = types.BoolValue(*d.DisablePortWhenRogueDhcpServerDetected)
	}
	if d.GatewayNonCompliant != nil {
		gatewayNonCompliant = types.BoolValue(*d.GatewayNonCompliant)
	}
	if d.SwitchMisconfiguredPort != nil {
		switchMisconfiguredPort = types.BoolValue(*d.SwitchMisconfiguredPort)
	}
	if d.SwitchPortStuck != nil {
		switchPortStuck = types.BoolValue(*d.SwitchPortStuck)
	}

	dataMapValue := map[string]attr.Value{
		"ap_insufficient_capacity":                     apInsufficientCapacity,
		"ap_loop":                                      apLoop,
		"ap_non_compliant":                             apNonCompliant,
		"bounce_port_for_abnormal_poe_client":          bouncePortForAbnormalPoeClient,
		"disable_port_when_ddos_protocol_violation":    disablePortWhenDdosProtocolViolation,
		"disable_port_when_rogue_dhcp_server_detected": disablePortWhenRogueDhcpServerDetected,
		"gateway_non_compliant":                        gatewayNonCompliant,
		"switch_misconfigured_port":                    switchMisconfiguredPort,
		"switch_port_stuck":                            switchPortStuck,
	}
	data, e := NewAutoOperationsValue(AutoOperationsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func marvisSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Marvis) MarvisValue {

	var autoOperations = types.ObjectNull(AutoOperationsValue{}.AttributeTypes(ctx))

	if d.AutoOperations != nil {
		autoOperations = marvisAutoOperationsSdkToTerraform(ctx, diags, d.AutoOperations)
	}

	dataMapValue := map[string]attr.Value{
		"auto_operations": autoOperations,
	}
	data, e := NewMarvisValue(MarvisValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
