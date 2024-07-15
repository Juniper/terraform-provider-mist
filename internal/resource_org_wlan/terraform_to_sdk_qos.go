package resource_org_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func qosTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d QosValue) *models.WlanQos {
	data := models.WlanQos{}
	data.Class = models.ToPointer(models.WlanQosClassEnum(string(d.Class.ValueString())))
	data.Overwrite = d.Overwrite.ValueBoolPointer()

	return &data
}
