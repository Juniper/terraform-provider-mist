package datasource_const_fingerprints

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func SdkToTerraform(_ context.Context, d models.ConstFingerprintTypes) (data ConstFingerprintsModel, diags diag.Diagnostics) {
	data.Family = mistutils.ListOfStringSdkToTerraform(d.Family)
	data.Mfg = mistutils.ListOfStringSdkToTerraform(d.Mfg)
	data.Model = mistutils.ListOfStringSdkToTerraform(d.Model)
	data.Os = mistutils.ListOfStringSdkToTerraform(d.Os)

	return data, diags
}
