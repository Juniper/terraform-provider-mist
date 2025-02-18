package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func proxyTerraformToSdk(d ProxyValue) *models.Proxy {
	data := models.Proxy{}

	data.Url = d.Url.ValueStringPointer()

	return &data
}
