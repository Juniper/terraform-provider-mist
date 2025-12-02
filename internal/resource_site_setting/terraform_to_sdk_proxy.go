package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func proxyTerraformToSdk(d ProxyValue) *models.Proxy {
	data := models.Proxy{}

	if !d.Disabled.IsNull() && !d.Disabled.IsUnknown() {
		data.Disabled = models.ToPointer(d.Disabled.ValueBool())
	}
	data.Url = d.Url.ValueStringPointer()

	return &data
}
