package resource_org_network

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TenantTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkTenant {
	data_map := make(map[string]models.NetworkTenant)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(TenantsValue)
		data := models.NetworkTenant{}
		data.Addresses = mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.Addresses)
		data_map[k] = data
	}
	return data_map
}
