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
	var defaultPsk basetypes.StringValue
	var defaultVlanId basetypes.StringValue
	var enabled basetypes.BoolValue
	var forceLookup basetypes.BoolValue
	var source basetypes.StringValue

	if d != nil && d.DefaultPsk != nil {
		defaultPsk = types.StringValue(*d.DefaultPsk)
	}
	if d != nil && d.DefaultVlanId != nil {
		defaultVlanId = types.StringValue(d.DefaultVlanId.String())
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.ForceLookup != nil {
		forceLookup = types.BoolValue(*d.ForceLookup)
	}
	if d != nil && d.Source != nil {
		source = types.StringValue(string(*d.Source))
	}

	dataMapValue := map[string]attr.Value{
		"default_psk":     defaultPsk,
		"default_vlan_id": defaultVlanId,
		"enabled":         enabled,
		"force_lookup":    forceLookup,
		"source":          source,
	}
	data, e := basetypes.NewObjectValue(DynamicPskValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
