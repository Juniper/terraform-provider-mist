package resource_org_setting

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func mistNacIdpsTerraformToSdk(diags *diag.Diagnostics, d basetypes.ListValue) []models.OrgSettingMistNacIdp {
	var dataList []models.OrgSettingMistNacIdp
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(IdpsValue)
		data := models.OrgSettingMistNacIdp{}

		if !plan.ExcludeRealms.IsNull() && !plan.ExcludeRealms.IsUnknown() {
			data.ExcludeRealms = misttransform.ListOfStringTerraformToSdk(plan.ExcludeRealms)
		}

		if plan.Id.ValueStringPointer() != nil {
			idUuid, e := uuid.Parse(plan.Id.String())
			if e == nil {
				data.Id = models.ToPointer(idUuid)
			} else {
				diags.AddWarning("Unable to parse Nac Idp UUID", e.Error())
			}
		}

		if !plan.UserRealms.IsNull() && !plan.UserRealms.IsUnknown() {
			data.UserRealms = misttransform.ListOfStringTerraformToSdk(plan.UserRealms)
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func mistNacServerCertTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.OrgSettingMistNacServerCert {
	data := models.OrgSettingMistNacServerCert{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewServerCertValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.Cert.ValueStringPointer() != nil {
				data.Cert = plan.Cert.ValueStringPointer()
			}

			if plan.Key.ValueStringPointer() != nil {
				data.Key = plan.Key.ValueStringPointer()
			}

			if plan.Password.ValueStringPointer() != nil {
				data.Password = plan.Password.ValueStringPointer()
			}
		}
	}
	return &data
}

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) *models.OrgSettingMistNac {
	data := models.OrgSettingMistNac{}

	if !d.Cacerts.IsNull() && !d.Cacerts.IsUnknown() {
		data.Cacerts = misttransform.ListOfStringTerraformToSdk(d.Cacerts)
	}

	if d.DefaultIdpId.ValueStringPointer() != nil {
		data.DefaultIdpId = d.DefaultIdpId.ValueStringPointer()
	}

	if d.DisableRsaeAlgorithms.ValueBoolPointer() != nil {
		data.DisableRsaeAlgorithms = d.DisableRsaeAlgorithms.ValueBoolPointer()
	}

	if d.EapSslSecurityLevel.ValueInt64Pointer() != nil {
		data.EapSslSecurityLevel = models.ToPointer(int(d.EapSslSecurityLevel.ValueInt64()))
	}

	if d.EuOnly.ValueBoolPointer() != nil {
		data.EuOnly = d.EuOnly.ValueBoolPointer()
	}

	if !d.Idps.IsNull() && !d.Idps.IsUnknown() {
		data.Idps = mistNacIdpsTerraformToSdk(diags, d.Idps)
	}

	if d.IdpMachineCertLookupField.ValueStringPointer() != nil {
		data.IdpMachineCertLookupField = (*models.IdpMachineCertLookupFieldEnum)(d.IdpMachineCertLookupField.ValueStringPointer())
	}

	if d.IdpUserCertLookupField.ValueStringPointer() != nil {
		data.IdpUserCertLookupField = (*models.IdpUserCertLookupFieldEnum)(d.IdpUserCertLookupField.ValueStringPointer())
	}

	if !d.ServerCert.IsNull() && !d.ServerCert.IsUnknown() {
		data.ServerCert = mistNacServerCertTerraformToSdk(ctx, diags, d.ServerCert)
	}

	if d.UseIpVersion.ValueStringPointer() != nil {
		data.UseIpVersion = (*models.OrgSettingMistNacIpVersionEnum)(d.UseIpVersion.ValueStringPointer())
	}

	if d.UseSslPort.ValueBoolPointer() != nil {
		data.UseSslPort = d.UseSslPort.ValueBoolPointer()
	}

	return &data
}
