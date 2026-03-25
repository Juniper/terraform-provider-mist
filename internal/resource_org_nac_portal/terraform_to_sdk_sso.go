package resource_org_nac_portal

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ssoRoleMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.NacPortalSsoRoleMatching {
	var dataList []models.NacPortalSsoRoleMatching
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(SsoRoleMatchingValue)
		data := models.NacPortalSsoRoleMatching{}

		if plan.Assigned.ValueStringPointer() != nil {
			data.Assigned = plan.Assigned.ValueStringPointer()
		}
		if plan.Match.ValueStringPointer() != nil {
			data.Match = plan.Match.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func ssoTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SsoValue) *models.NacPortalSso {
	data := models.NacPortalSso{}

	if d.IdpCert.ValueStringPointer() != nil {
		data.IdpCert = d.IdpCert.ValueStringPointer()
	}
	if d.IdpSignAlgo.ValueStringPointer() != nil {
		data.IdpSignAlgo = models.ToPointer(models.NacPortalSsoIdpSignAlgoEnum(d.IdpSignAlgo.ValueString()))
	}
	if d.IdpSsoUrl.ValueStringPointer() != nil {
		data.IdpSsoUrl = d.IdpSsoUrl.ValueStringPointer()
	}
	if d.Issuer.ValueStringPointer() != nil {
		data.Issuer = d.Issuer.ValueStringPointer()
	}
	if d.NameidFormat.ValueStringPointer() != nil {
		data.NameidFormat = d.NameidFormat.ValueStringPointer()
	}
	if !d.SsoRoleMatching.IsNull() && !d.SsoRoleMatching.IsUnknown() {
		data.SsoRoleMatching = ssoRoleMatchingTerraformToSdk(ctx, diags, d.SsoRoleMatching)
	}
	if d.UseSsoRoleForCert.ValueBoolPointer() != nil {
		data.UseSsoRoleForCert = d.UseSsoRoleForCert.ValueBoolPointer()
	}

	return &data
}
