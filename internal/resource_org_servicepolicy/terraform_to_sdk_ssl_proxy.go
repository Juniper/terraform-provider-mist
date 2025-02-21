package resource_org_servicepolicy

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func sslProxyTerraformToSdk(diags *diag.Diagnostics, d SslProxyValue) *models.ServicePolicySslProxy {

	data := models.ServicePolicySslProxy{}
	if d.CiphersCatagory.ValueStringPointer() != nil {
		data.CiphersCatagory = (*models.SslProxyCiphersCatagoryEnum)(d.CiphersCatagory.ValueStringPointer())
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	return &data
}
