package resource_org_wlan

import (
	"context"
	"strings"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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
		if v_plan.DisableLocal.ValueBoolPointer() != nil {
			v_data.DisableLocal = v_plan.DisableLocal.ValueBoolPointer()
		}
		if !v_plan.RadiusGroups.IsNull() && !v_plan.RadiusGroups.IsUnknown() {
			v_data.RadiusGroups = mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.RadiusGroups)
		}
		if v_plan.Scope.ValueStringPointer() != nil {
			v_data.Scope = models.ToPointer(models.WlanBonjourServicePropertiesScopeEnum(string(v_plan.Scope.ValueString())))
		}
		data_map[k] = v_data
	}
	return data_map
}
func bonjourTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan BonjourValue) *models.WlanBonjour {

	data := models.WlanBonjour{}

	var tmp []string
	for _, v := range plan.AdditionalVlanIds.Elements() {
		var i interface{} = v
		s := i.(basetypes.StringValue)
		tmp = append(tmp, s.ValueString())
	}
	data.AdditionalVlanIds = strings.Join(tmp, ",")
	data.Services = bonjourServicesTerraformToSdk(ctx, diags, plan.Services)
	data.Enabled = plan.Enabled.ValueBoolPointer()

	return &data
}
