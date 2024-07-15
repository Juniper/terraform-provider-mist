package resource_org_wxtag

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func specsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.WxlanTagSpec) basetypes.ListValue {

	var data_list = []SpecsValue{}
	for _, v := range data {
		data_map_value := map[string]attr.Value{
			"port_range": types.StringValue(*v.PortRange),
			"protocol":   types.StringValue(*v.Protocol),
			"subnets":    mist_transform.ListOfStringSdkToTerraform(ctx, v.Subnets),
		}
		data, e := NewSpecsValue(SpecsValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, SpecsValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
