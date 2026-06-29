package resource_org_deviceprofile_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func mqttConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApMqtt) MqttConfigValue {

	var brokerHost basetypes.StringValue
	var brokerPort basetypes.Int64Value
	var brokerProto basetypes.StringValue
	var enabled basetypes.BoolValue
	var format basetypes.StringValue
	var password basetypes.StringValue
	var username basetypes.StringValue

	if d.BrokerHost != nil {
		brokerHost = types.StringValue(*d.BrokerHost)
	}
	if d.BrokerPort != nil {
		brokerPort = types.Int64Value(int64(*d.BrokerPort))
	}
	if d.BrokerProto != nil {
		brokerProto = types.StringValue(string(*d.BrokerProto))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Format != nil {
		format = types.StringValue(string(*d.Format))
	}
	if d.Password != nil {
		password = types.StringValue(*d.Password)
	}
	if d.Username != nil {
		username = types.StringValue(*d.Username)
	}

	dataMapValue := map[string]attr.Value{
		"broker_host":  brokerHost,
		"broker_port":  brokerPort,
		"broker_proto": brokerProto,
		"enabled":      enabled,
		"format":       format,
		"password":     password,
		"username":     username,
	}
	data, e := NewMqttConfigValue(MqttConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
