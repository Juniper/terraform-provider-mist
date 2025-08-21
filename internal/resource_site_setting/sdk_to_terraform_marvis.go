package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func marvisAutoOperationsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MarvisAutoOperations) basetypes.ObjectValue {
	var bouncePortForAbnormalPoeClient basetypes.BoolValue
	var disablePortWhenDdosProtocolViolation basetypes.BoolValue
	var disablePortWhenRogueDhcpServerDetected basetypes.BoolValue

	if d.BouncePortForAbnormalPoeClient != nil {
		bouncePortForAbnormalPoeClient = types.BoolValue(*d.BouncePortForAbnormalPoeClient)
	}
	if d.DisablePortWhenDdosProtocolViolation != nil {
		disablePortWhenDdosProtocolViolation = types.BoolValue(*d.DisablePortWhenDdosProtocolViolation)
	}
	if d.DisablePortWhenRogueDhcpServerDetected != nil {
		disablePortWhenRogueDhcpServerDetected = types.BoolValue(*d.DisablePortWhenRogueDhcpServerDetected)
	}

	dataMapValue := map[string]attr.Value{
		"bounce_port_for_abnormal_poe_client":          bouncePortForAbnormalPoeClient,
		"disable_port_when_ddos_protocol_violation":    disablePortWhenDdosProtocolViolation,
		"disable_port_when_rogue_dhcp_server_detected": disablePortWhenRogueDhcpServerDetected,
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
