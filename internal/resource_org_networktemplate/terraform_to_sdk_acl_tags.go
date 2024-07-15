package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func actTagSpecsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.AclTagSpec {
	var data []models.AclTagSpec
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_state := v_interface.(SpecsValue)
		v_data := models.AclTagSpec{}
		if v_state.PortRange.ValueStringPointer() != nil {
			v_data.PortRange = models.ToPointer(v_state.PortRange.ValueString())
		}
		if v_state.Protocol.ValueStringPointer() != nil {
			v_data.Protocol = models.ToPointer(v_state.Protocol.ValueString())
		}
		data = append(data, v_data)
	}
	return data
}

func actTagsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.AclTag {
	data := make(map[string]models.AclTag)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(AclTagsValue)

		data_item := models.AclTag{}
		data_item.Type = models.AclTagTypeEnum(item_obj.AclTagsType.ValueString())
		if item_obj.GbpTag.ValueInt64Pointer() != nil {
			data_item.GbpTag = models.ToPointer(int(item_obj.GbpTag.ValueInt64()))
		}
		data_item.Macs = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Macs)
		if item_obj.Network.ValueStringPointer() != nil {
			data_item.Network = models.ToPointer(item_obj.Network.ValueString())
		}
		if item_obj.RadiusGroup.ValueStringPointer() != nil {
			data_item.RadiusGroup = models.ToPointer(item_obj.RadiusGroup.ValueString())
		}
		if !item_obj.Specs.IsNull() && !item_obj.Specs.IsUnknown() {
			data_item.Specs = actTagSpecsTerraformToSdk(ctx, diags, item_obj.Specs)
		}
		if !item_obj.Subnets.IsNull() && !item_obj.Subnets.IsUnknown() {
			data_item.Subnets = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Subnets)
		}
		data[item_name] = data_item
	}
	return data
}
