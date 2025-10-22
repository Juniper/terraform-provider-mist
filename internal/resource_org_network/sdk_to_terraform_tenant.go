package resource_org_network

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tenantSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkTenant) basetypes.MapValue {

	stateValueMapAttrType := TenantsValue{}.AttributeTypes(ctx)
	stateValueMapValue := make(map[string]attr.Value)
	for k, v := range d {
		stateValueMapAttrValue := map[string]attr.Value{
			"addresses": mistutils.ListOfStringSdkToTerraform(v.Addresses),
		}
		n, e := NewTenantsValue(stateValueMapAttrType, stateValueMapAttrValue)
		diags.Append(e...)
		stateValueMapValue[k] = n
	}
	stateResultMapType := TenantsValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}
