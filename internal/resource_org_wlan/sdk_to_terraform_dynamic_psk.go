package resource_org_wlan

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicPskSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanDynamicPsk) DynamicPskValue {
	var defaultPsk basetypes.StringValue
	var defaultVlanId basetypes.StringValue
	var enabled basetypes.BoolValue
	var forceLookup basetypes.BoolValue
	var source basetypes.StringValue

	if d != nil && d.DefaultPsk != nil {
		defaultPsk = types.StringValue(*d.DefaultPsk)
	}
	if d != nil && d.DefaultVlanId != nil {
		defaultVlanId = mistutils.VlanAsString(*d.DefaultVlanId)
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
	data, e := NewDynamicPskValue(DynamicPskValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
