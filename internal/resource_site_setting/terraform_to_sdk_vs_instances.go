package resource_site_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vsInstanceTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, m basetypes.MapValue) map[string]models.VsInstanceProperty {
	data_map := make(map[string]models.VsInstanceProperty)
	for k, v := range m.Elements() {

		var vi interface{} = v
		vd := vi.(VsInstanceValue)
		data := models.VsInstanceProperty{}
		data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, vd.Networks)
		data_map[k] = data
	}
	return data_map
}
