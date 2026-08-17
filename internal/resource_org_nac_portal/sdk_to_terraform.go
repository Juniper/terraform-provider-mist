package resource_org_nac_portal

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, d *models.NacPortal) (OrgNacPortalModel, diag.Diagnostics) {
	var state OrgNacPortalModel
	var diags diag.Diagnostics

	var accessType types.String
	var additionalCacerts = types.ListNull(types.StringType)
	var additionalNacServerName = types.ListNull(types.StringType)
	var certExpireTime types.Int64
	var eapType types.String
	var enableLocation types.Bool
	var enableTelemetry types.Bool
	var expiryNotificationTime types.Int64
	var id types.String
	var name types.String
	var notifyExpiry types.Bool
	var orgId types.String
	var portal = NewPortalValueNull()
	var ssid types.String
	var sso = NewSsoValueNull()
	var tos types.String
	var typeNacPortal types.String

	if d.AccessType != nil {
		accessType = types.StringValue(string(*d.AccessType))
	}
	if d.AdditionalCacerts != nil {
		additionalCacerts = mistutils.ListOfStringSdkToTerraform(d.AdditionalCacerts)
	}
	if d.AdditionalNacServerName != nil {
		additionalNacServerName = mistutils.ListOfStringSdkToTerraform(d.AdditionalNacServerName)
	}
	if d.CertExpireTime != nil {
		certExpireTime = types.Int64Value(int64(*d.CertExpireTime))
	}
	if d.EapType != nil {
		eapType = types.StringValue(string(*d.EapType))
	}
	if d.EnableLocation != nil {
		enableLocation = types.BoolValue(*d.EnableLocation)
	}
	if d.EnableTelemetry != nil {
		enableTelemetry = types.BoolValue(*d.EnableTelemetry)
	}
	if d.ExpiryNotificationTime != nil {
		expiryNotificationTime = types.Int64Value(int64(*d.ExpiryNotificationTime))
	}

	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}

	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.NotifyExpiry != nil {
		notifyExpiry = types.BoolValue(*d.NotifyExpiry)
	}

	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}

	if d.Portal != nil {
		portal = portalSdkToTerraform(ctx, &diags, d.Portal)
	}
	if d.Ssid != nil {
		ssid = types.StringValue(*d.Ssid)
	}
	if d.Sso != nil {
		sso = ssoSdkToTerraform(ctx, &diags, d.Sso)
	}
	if d.Tos != nil {
		tos = types.StringValue(*d.Tos)
	}
	if d.Type != nil {
		typeNacPortal = types.StringValue(string(*d.Type))
	}

	state.AccessType = accessType
	state.AdditionalCacerts = additionalCacerts
	state.AdditionalNacServerName = additionalNacServerName
	state.CertExpireTime = certExpireTime
	state.EapType = eapType
	state.EnableLocation = enableLocation
	state.EnableTelemetry = enableTelemetry
	state.ExpiryNotificationTime = expiryNotificationTime
	state.Id = id
	state.Name = name
	state.NotifyExpiry = notifyExpiry
	state.OrgId = orgId
	state.Portal = portal
	state.Ssid = ssid
	state.Sso = sso
	state.Tos = tos
	state.Type = typeNacPortal

	return state, diags
}
