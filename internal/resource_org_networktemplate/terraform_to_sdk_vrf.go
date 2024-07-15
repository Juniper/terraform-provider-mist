package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vrfConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VrfConfigValue) *models.VrfConfig {
	data := models.VrfConfig{}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	return &data
}

func vrfInstanceExtraRouteTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.VrfExtraRoute {
	data := make(map[string]models.VrfExtraRoute)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(ExtraRoutesValue)

		data_item := models.VrfExtraRoute{}
		if item_obj.Via.ValueStringPointer() != nil {
			data_item.Via = models.ToPointer(item_obj.Via.ValueString())
		}
		data[item_name] = data_item
	}
	return data
}

func vrfInstancesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.VrfInstance {
	data := make(map[string]models.VrfInstance)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(VrfInstancesValue)

		data_item := models.VrfInstance{}
		if !item_obj.Networks.IsNull() && !item_obj.Networks.IsUnknown() {
			data_item.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Networks)
		}
		if !item_obj.VrfExtraRoutes.IsNull() && !item_obj.VrfExtraRoutes.IsUnknown() {
			data_item.ExtraRoutes = vrfInstanceExtraRouteTerraformToSdk(ctx, diags, item_obj.VrfExtraRoutes)
		}

		data[item_name] = data_item
	}
	return data
}
