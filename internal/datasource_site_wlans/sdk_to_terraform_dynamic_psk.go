package datasource_site_wlans

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicPskSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanDynamicPsk) basetypes.ObjectValue {
	var default_psk basetypes.StringValue
	var default_vlan_id basetypes.StringValue
	var enabled basetypes.BoolValue
	var force_lookup basetypes.BoolValue
	var source basetypes.StringValue

	if d != nil && d.DefaultPsk != nil {
		default_psk = types.StringValue(*d.DefaultPsk)
	}
	if d != nil && d.DefaultVlanId != nil {
		default_vlan_id = types.StringValue(d.DefaultVlanId.String())
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.ForceLookup != nil {
		force_lookup = types.BoolValue(*d.ForceLookup)
	}
	if d != nil && d.Source != nil {
		source = types.StringValue(string(*d.Source))
	}

	data_map_attr_type := DynamicPskValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"default_psk":     default_psk,
		"default_vlan_id": default_vlan_id,
		"enabled":         enabled,
		"force_lookup":    force_lookup,
		"source":          source,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
