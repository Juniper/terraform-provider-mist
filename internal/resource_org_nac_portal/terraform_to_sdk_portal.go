package resource_org_nac_portal

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func portalTerraformToSdk(d PortalValue) *models.NacPortalGuestPortal {
	data := models.NacPortalGuestPortal{}

	if d.Auth.ValueStringPointer() != nil {
		data.Auth = models.ToPointer(models.NacPortalGuestPortalAuthEnum(d.Auth.ValueString()))
	}
	if d.Expire.ValueInt64Pointer() != nil {
		data.Expire = models.ToPointer(int(d.Expire.ValueInt64()))
	}
	if d.ExternalPortalUrl.ValueStringPointer() != nil {
		data.ExternalPortalUrl = d.ExternalPortalUrl.ValueStringPointer()
	}
	if d.ForceReconnect.ValueBoolPointer() != nil {
		data.ForceReconnect = d.ForceReconnect.ValueBoolPointer()
	}
	if d.Forward.ValueBoolPointer() != nil {
		data.Forward = d.Forward.ValueBoolPointer()
	}
	if d.ForwardUrl.ValueStringPointer() != nil {
		data.ForwardUrl = d.ForwardUrl.ValueStringPointer()
	}
	if d.MaxNumDevices.ValueInt64Pointer() != nil {
		data.MaxNumDevices = models.ToPointer(int(d.MaxNumDevices.ValueInt64()))
	}
	if d.Privacy.ValueBoolPointer() != nil {
		data.Privacy = d.Privacy.ValueBoolPointer()
	}

	return &data
}
