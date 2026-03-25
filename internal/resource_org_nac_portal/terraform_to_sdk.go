package resource_org_nac_portal

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgNacPortalModel) (models.NacPortal, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.NacPortal{}
	unset := make(map[string]interface{})

	if !plan.AccessType.IsNull() && !plan.AccessType.IsUnknown() {
		data.AccessType = models.ToPointer(models.NacPortalAccessTypeEnum(plan.AccessType.ValueString()))
	} else {
		unset["-access_type"] = ""
	}

	if !plan.AdditionalCacerts.IsNull() && !plan.AdditionalCacerts.IsUnknown() {
		data.AdditionalCacerts = mistutils.ListOfStringTerraformToSdk(plan.AdditionalCacerts)
	} else {
		unset["-additional_cacerts"] = ""
	}

	if !plan.AdditionalNacServerName.IsNull() && !plan.AdditionalNacServerName.IsUnknown() {
		data.AdditionalNacServerName = mistutils.ListOfStringTerraformToSdk(plan.AdditionalNacServerName)
	} else {
		unset["-additional_nac_server_name"] = ""
	}

	if !plan.CertExpireTime.IsNull() && !plan.CertExpireTime.IsUnknown() {
		data.CertExpireTime = models.ToPointer(int(plan.CertExpireTime.ValueInt64()))
	} else {
		unset["-cert_expire_time"] = ""
	}

	if !plan.EapType.IsNull() && !plan.EapType.IsUnknown() {
		data.EapType = models.ToPointer(models.NacPortalEapTypeEnum(plan.EapType.ValueString()))
	} else {
		unset["-eap_type"] = ""
	}

	if !plan.EnableTelemetry.IsNull() && !plan.EnableTelemetry.IsUnknown() {
		data.EnableTelemetry = plan.EnableTelemetry.ValueBoolPointer()
	} else {
		unset["-enable_telemetry"] = ""
	}

	if !plan.ExpiryNotificationTime.IsNull() && !plan.ExpiryNotificationTime.IsUnknown() {
		data.ExpiryNotificationTime = models.ToPointer(int(plan.ExpiryNotificationTime.ValueInt64()))
	} else {
		unset["-expiry_notification_time"] = ""
	}

	data.Name = plan.Name.ValueStringPointer()

	if !plan.NotifyExpiry.IsNull() && !plan.NotifyExpiry.IsUnknown() {
		data.NotifyExpiry = plan.NotifyExpiry.ValueBoolPointer()
	} else {
		unset["-notify_expiry"] = ""
	}

	if !plan.Portal.IsNull() && !plan.Portal.IsUnknown() {
		data.Portal = portalTerraformToSdk(plan.Portal)
	} else {
		unset["-portal"] = ""
	}

	if !plan.Ssid.IsNull() && !plan.Ssid.IsUnknown() {
		data.Ssid = plan.Ssid.ValueStringPointer()
	} else {
		unset["-ssid"] = ""
	}

	if !plan.Sso.IsNull() && !plan.Sso.IsUnknown() {
		data.Sso = ssoTerraformToSdk(ctx, &diags, plan.Sso)
	} else {
		unset["-sso"] = ""
	}

	if !plan.Tos.IsNull() && !plan.Tos.IsUnknown() {
		data.Tos = plan.Tos.ValueStringPointer()
	} else {
		unset["-tos"] = ""
	}

	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		data.Type = models.ToPointer(models.NacPortalTypeEnum(plan.Type.ValueString()))
	} else {
		unset["-type"] = ""
	}

	data.AdditionalProperties = unset
	return data, diags
}
