package datasource_org_wxtags

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func specsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.WxlanTagSpec) basetypes.ListValue {

	var dataList []SpecsValue
	for _, v := range data {
		dataMapValue := map[string]attr.Value{
			"port_range": types.StringValue(*v.PortRange),
			"protocol":   types.StringValue(*v.Protocol),
			"subnets":    misttransform.ListOfStringSdkToTerraform(v.Subnets),
		}
		data, e := NewSpecsValue(SpecsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, SpecsValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
