package datasource_const_fingerprints

import (
	"context"
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func SdkToTerraform(_ context.Context, d models.ConstFingerprintTypes) (data ConstFingerprintsModel, diags diag.Diagnostics) {
	data.Family = misttransform.ListOfStringSdkToTerraform(d.Family)
	data.Mfg = misttransform.ListOfStringSdkToTerraform(d.Mfg)
	data.Model = misttransform.ListOfStringSdkToTerraform(d.Model)
	data.OsType = misttransform.ListOfStringSdkToTerraform(d.OsType)

	return data, diags
}
