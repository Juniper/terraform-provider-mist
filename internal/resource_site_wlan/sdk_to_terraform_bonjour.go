package resource_site_wlan

import (
	"context"
	"strings"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bonjourServicesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.WlanBonjourServiceProperties) basetypes.MapValue {

	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {

		var disableLocal basetypes.BoolValue
		var radiusGroups = mistutils.ListOfStringSdkToTerraformEmpty()
		var scope basetypes.StringValue

		if d.DisableLocal != nil {
			disableLocal = types.BoolValue(*d.DisableLocal)
		}
		if d.RadiusGroups != nil {
			radiusGroups = mistutils.ListOfStringSdkToTerraform(d.RadiusGroups)
		}
		if d.Scope != nil {
			scope = types.StringValue(string(*d.Scope))
		}

		dataMapValue := map[string]attr.Value{
			"disable_local": disableLocal,
			"radius_groups": radiusGroups,
			"scope":         scope,
		}
		data, e := NewServicesValue(ServicesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	r, e := types.MapValueFrom(ctx, ServicesValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return r
}

func bonjourSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanBonjour) BonjourValue {
	var additionalVlanIds = types.ListNull(types.StringType)
	var enabled basetypes.BoolValue
	var services = types.MapNull(ServicesValue{}.Type(ctx))

	if d != nil {
		var items []attr.Value
		for _, item := range strings.Split(d.AdditionalVlanIds, ",") {
			if item != "" {
				items = append(items, types.StringValue(item))
			}
		}
		list, _ := types.ListValue(basetypes.StringType{}, items)
		additionalVlanIds = list
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Services != nil && len(d.Services) > 0 {
		services = bonjourServicesSdkToTerraform(ctx, diags, d.Services)
	}

	dataMapValue := map[string]attr.Value{
		"additional_vlan_ids": additionalVlanIds,
		"enabled":             enabled,
		"services":            services,
	}
	data, e := NewBonjourValue(BonjourValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
