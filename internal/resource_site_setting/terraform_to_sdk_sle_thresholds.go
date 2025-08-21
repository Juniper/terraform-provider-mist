package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func sleThresholdsTerraformToSdk(d SleThresholdsValue) *models.SleThresholds {
	data := models.SleThresholds{}

	if d.Capacity.ValueInt64Pointer() != nil {
		data.Capacity = models.ToPointer(int(d.Capacity.ValueInt64()))
	}

	if d.Coverage.ValueInt64Pointer() != nil {
		data.Coverage = models.ToPointer(int(d.Coverage.ValueInt64()))
	}

	if d.Throughput.ValueInt64Pointer() != nil {
		data.Throughput = models.ToPointer(int(d.Throughput.ValueInt64()))
	}

	if d.Timetoconnect.ValueInt64Pointer() != nil {
		data.TimeToConnect = models.ToPointer(int(d.Timetoconnect.ValueInt64()))
	}

	return &data
}
