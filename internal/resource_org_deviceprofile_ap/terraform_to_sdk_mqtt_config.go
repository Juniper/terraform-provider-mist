package resource_org_deviceprofile_ap

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func mqttConfigTerraformToSdk(d MqttConfigValue) *models.ApMqtt {

	data := models.ApMqtt{}

	if d.BrokerHost.ValueStringPointer() != nil {
		data.BrokerHost = d.BrokerHost.ValueStringPointer()
	}
	if d.BrokerPort.ValueInt64Pointer() != nil {
		data.BrokerPort = models.ToPointer(int(d.BrokerPort.ValueInt64()))
	}
	if d.BrokerProto.ValueStringPointer() != nil {
		data.BrokerProto = models.ToPointer(models.ApMqttBrokerProtoEnum(d.BrokerProto.ValueString()))
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Format.ValueStringPointer() != nil {
		data.Format = models.ToPointer(models.ApMqttFormatEnum(d.Format.ValueString()))
	}
	if d.Password.ValueStringPointer() != nil {
		data.Password = d.Password.ValueStringPointer()
	}
	if d.Username.ValueStringPointer() != nil {
		data.Username = d.Username.ValueStringPointer()
	}

	return &data
}
