package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func authSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanAuth) AuthValue {

	var anticlog_threshold basetypes.Int64Value
	var eap_reauth basetypes.BoolValue
	var enable_mac_auth basetypes.BoolValue
	var key_idx basetypes.Int64Value
	var keys basetypes.ListValue = types.ListNull(types.StringType)
	var multi_psk_only basetypes.BoolValue
	var owe basetypes.StringValue
	var pairwise basetypes.ListValue = types.ListNull(types.StringType)
	var private_wlan basetypes.BoolValue = types.BoolValue(false)
	var psk basetypes.StringValue
	var type_auth basetypes.StringValue
	var wep_as_secondary_auth basetypes.BoolValue

	if d != nil && d.AnticlogThreshold != nil {
		anticlog_threshold = types.Int64Value(int64(*d.AnticlogThreshold))
	}
	if d != nil && d.EapReauth != nil {
		eap_reauth = types.BoolValue(*d.EapReauth)
	}
	if d != nil && d.EnableMacAuth != nil {
		enable_mac_auth = types.BoolValue(*d.EnableMacAuth)
	}
	if d != nil && d.KeyIdx != nil {
		key_idx = types.Int64Value(int64(*d.KeyIdx))
	}

	var keys_list []attr.Value
	if d != nil && d.Keys != nil {
		for _, item := range d.Keys {
			value := item
			keys_list = append(keys_list, types.StringValue(value))
		}
	}
	keys = types.ListValueMust(basetypes.StringType{}, keys_list)

	if d != nil && d.MultiPskOnly != nil {
		multi_psk_only = types.BoolValue(*d.MultiPskOnly)
	}
	if d != nil && d.Owe != nil {
		owe = types.StringValue(string(*d.Owe))
	}
	var pairwise_list []attr.Value
	if d != nil && d.Pairwise != nil {
		for _, item := range d.Pairwise {
			value := string(item)
			pairwise_list = append(pairwise_list, types.StringValue(value))
		}
	}
	pairwise = types.ListValueMust(basetypes.StringType{}, pairwise_list)

	if d != nil && d.PrivateWlan != nil {
		private_wlan = types.BoolValue(*d.PrivateWlan)
	}
	if d != nil && d.Psk.Value() != nil {
		psk = types.StringValue(*d.Psk.Value())
	}
	if d != nil {
		type_auth = types.StringValue(string(d.Type))
	}
	if d != nil && d.WepAsSecondaryAuth != nil {
		wep_as_secondary_auth = types.BoolValue(*d.WepAsSecondaryAuth)
	}

	data_map_attr_type := AuthValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"anticlog_threshold":    anticlog_threshold,
		"eap_reauth":            eap_reauth,
		"enable_mac_auth":       enable_mac_auth,
		"key_idx":               key_idx,
		"keys":                  keys,
		"multi_psk_only":        multi_psk_only,
		"owe":                   owe,
		"pairwise":              pairwise,
		"private_wlan":          private_wlan,
		"psk":                   psk,
		"type":                  type_auth,
		"wep_as_secondary_auth": wep_as_secondary_auth,
	}
	data, e := NewAuthValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
