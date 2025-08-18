package resource_org_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func mistNacTerraformToSdk(d MistNacValue) *models.WlanMistNac {
	data := models.WlanMistNac{}

	if d.AcctInterimInterval.ValueInt64Pointer() != nil {
		data.AcctInterimInterval = models.ToPointer(int(d.AcctInterimInterval.ValueInt64()))
	}
	if d.AuthServersRetries.ValueInt64Pointer() != nil {
		data.AuthServersRetries = models.ToPointer(int(d.AuthServersRetries.ValueInt64()))
	}
	if d.AuthServersTimeout.ValueInt64Pointer() != nil {
		data.AuthServersTimeout = models.ToPointer(int(d.AuthServersTimeout.ValueInt64()))
	}
	if d.CoaEnabled.ValueBoolPointer() != nil {
		data.CoaEnabled = d.CoaEnabled.ValueBoolPointer()
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.FastDot1xTimers.ValueBoolPointer() != nil {
		data.FastDot1xTimers = models.ToPointer(d.FastDot1xTimers.ValueBool())
	}
	if d.Network.ValueStringPointer() != nil {
		data.Network = models.NewOptional(d.Network.ValueStringPointer())
	}
	if d.SourceIp.ValueStringPointer() != nil {
		data.SourceIp = models.NewOptional(d.SourceIp.ValueStringPointer())
	}

	return &data
}
