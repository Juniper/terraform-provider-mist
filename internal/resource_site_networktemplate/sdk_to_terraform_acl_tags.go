package resource_site_networktemplate

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func aclTagSpecsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.AclTagSpec) basetypes.ListValue {

	var dataList []SpecsValue

	for _, d := range l {

		var portRange basetypes.StringValue
		var protocol basetypes.StringValue

		if d.PortRange != nil {
			portRange = types.StringValue(*d.PortRange)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(*d.Protocol)
		}

		dataMapValue := map[string]attr.Value{
			"port_range": portRange,
			"protocol":   protocol,
		}
		data, e := NewSpecsValue(SpecsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, SpecsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
func aclTagsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.AclTag) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var gbpTag basetypes.Int64Value
		var macs = mistutils.ListOfStringSdkToTerraformEmpty()
		var network basetypes.StringValue
		var radiusGroup basetypes.StringValue
		var specs = types.ListNull(SpecsValue{}.Type(ctx))
		var subnets = types.ListNull(types.StringType)
		var tagType = types.StringValue(string(d.Type))

		if d.GbpTag != nil {
			gbpTag = types.Int64Value(int64(*d.GbpTag))
		}
		if d.Macs != nil {
			macs = mistutils.ListOfStringSdkToTerraform(d.Macs)
		}
		if d.Network != nil {
			network = types.StringValue(*d.Network)
		}
		if d.RadiusGroup != nil {
			radiusGroup = types.StringValue(*d.RadiusGroup)
		}
		if d.Specs != nil {
			specs = aclTagSpecsSdkToTerraform(ctx, diags, d.Specs)
		}
		if d.Subnets != nil {
			subnets = mistutils.ListOfStringSdkToTerraform(d.Subnets)
		}

		dataMapValue := map[string]attr.Value{
			"gbp_tag":      gbpTag,
			"macs":         macs,
			"network":      network,
			"radius_group": radiusGroup,
			"specs":        specs,
			"subnets":      subnets,
			"type":         tagType,
		}
		data, e := NewAclTagsValue(AclTagsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMap, e := types.MapValueFrom(ctx, AclTagsValue{}.Type(ctx), stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}
