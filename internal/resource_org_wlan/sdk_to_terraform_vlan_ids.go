package resource_org_wlan

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vlanIdsSkToTerraform(diags *diag.Diagnostics, data []models.VlanIdWithVariable) basetypes.ListValue {

	var list []attr.Value
	for _, v := range data {
		list = append(list, types.StringValue(v.String()))
	}
	r, e := types.ListValue(basetypes.StringType{}, list)
	diags.Append(e...)

	return r
}
