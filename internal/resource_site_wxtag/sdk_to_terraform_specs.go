package resource_site_wxtag

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func specsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []models.WxlanTagSpec) basetypes.ListValue {

	var dataList []SpecsValue
	for _, v := range data {
		var portRange = types.StringValue("0")
		var protocol = types.StringValue("any")
		var subnets = types.ListNull(types.StringType)

		if v.PortRange != nil {
			portRange = types.StringValue(*v.PortRange)
		}
		if v.Protocol != nil {
			protocol = types.StringValue(*v.Protocol)
		}
		if len(v.Subnets) > 0 {
			subnets = mistutils.ListOfStringSdkToTerraform(v.Subnets)
		}

		dataMapValue := map[string]attr.Value{
			"port_range": portRange,
			"protocol":   protocol,
			"subnets":    subnets,
		}
		data, e := NewSpecsValue(SpecsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, SpecsValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
