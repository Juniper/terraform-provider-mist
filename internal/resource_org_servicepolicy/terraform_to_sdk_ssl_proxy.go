package resource_org_servicepolicy

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func sslProxyTerraformToSdk(_ *diag.Diagnostics, d SslProxyValue) *models.ServicePolicySslProxy {

	data := models.ServicePolicySslProxy{}
	if d.CiphersCategory.ValueStringPointer() != nil {
		data.CiphersCategory = (*models.SslProxyCiphersCategoryEnum)(d.CiphersCategory.ValueStringPointer())
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	return &data
}
