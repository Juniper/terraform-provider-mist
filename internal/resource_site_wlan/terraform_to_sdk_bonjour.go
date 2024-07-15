package resource_site_wlan

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bonjourServicesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.MapValue) map[string]models.WlanBonjourServiceProperties {
	data_map := make(map[string]models.WlanBonjourServiceProperties)
	for k, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(ServicesValue)
		v_data := models.WlanBonjourServiceProperties{}
		v_data.DisableLocal = v_plan.DisableLocal.ValueBoolPointer()
		v_data.RadiusGroups = mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.RadiusGroups)
		v_data.Scope = models.ToPointer(models.WlanBonjourServicePropertiesScopeEnum(string(v_plan.Scope.ValueString())))

		data_map[k] = v_data
	}
	return data_map
}
func bonjourTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan BonjourValue) *models.WlanBonjour {

	data := models.WlanBonjour{}
	data.AdditionalVlanIds = mist_transform.ListOfIntTerraformToSdk(ctx, plan.AdditionalVlanIds)
	data.Services = bonjourServicesTerraformToSdk(ctx, diags, plan.Services)
	data.Enabled = plan.Enabled.ValueBoolPointer()

	return &data
}
