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

	var anticlogThreshold basetypes.Int64Value
	var eapReauth basetypes.BoolValue
	var enableMacAuth basetypes.BoolValue
	var keyIdx basetypes.Int64Value
	var keys = types.ListNull(types.StringType)
	var multiPskOnly basetypes.BoolValue
	var owe basetypes.StringValue
	var pairwise = types.ListNull(types.StringType)
	var privateWlan = types.BoolValue(false)
	var psk = types.StringValue("")
	var typeAuth basetypes.StringValue
	var wepAsSecondaryAuth basetypes.BoolValue

	if d != nil && d.AnticlogThreshold != nil {
		anticlogThreshold = types.Int64Value(int64(*d.AnticlogThreshold))
	}
	if d != nil && d.EapReauth != nil {
		eapReauth = types.BoolValue(*d.EapReauth)
	}
	if d != nil && d.EnableMacAuth != nil {
		enableMacAuth = types.BoolValue(*d.EnableMacAuth)
	}
	if d != nil && d.KeyIdx != nil {
		keyIdx = types.Int64Value(int64(*d.KeyIdx))
	}

	var keysList []attr.Value
	if d != nil && d.Keys != nil {
		for _, item := range d.Keys {
			value := item
			keysList = append(keysList, types.StringValue(value))
		}
	}
	keys = types.ListValueMust(basetypes.StringType{}, keysList)

	if d != nil && d.MultiPskOnly != nil {
		multiPskOnly = types.BoolValue(*d.MultiPskOnly)
	}
	if d != nil && d.Owe != nil {
		owe = types.StringValue(string(*d.Owe))
	}
	var pairwiseList []attr.Value
	if d != nil && d.Pairwise != nil {
		for _, item := range d.Pairwise {
			value := string(item)
			pairwiseList = append(pairwiseList, types.StringValue(value))
		}
	}
	pairwise = types.ListValueMust(basetypes.StringType{}, pairwiseList)

	if d != nil && d.PrivateWlan != nil {
		privateWlan = types.BoolValue(*d.PrivateWlan)
	}
	if d != nil && d.Psk.Value() != nil {
		psk = types.StringValue(*d.Psk.Value())
	}
	if d != nil {
		typeAuth = types.StringValue(string(d.Type))
	}
	if d != nil && d.WepAsSecondaryAuth != nil {
		wepAsSecondaryAuth = types.BoolValue(*d.WepAsSecondaryAuth)
	}

	dataMapValue := map[string]attr.Value{
		"anticlog_threshold":    anticlogThreshold,
		"eap_reauth":            eapReauth,
		"enable_mac_auth":       enableMacAuth,
		"key_idx":               keyIdx,
		"keys":                  keys,
		"multi_psk_only":        multiPskOnly,
		"owe":                   owe,
		"pairwise":              pairwise,
		"private_wlan":          privateWlan,
		"psk":                   psk,
		"type":                  typeAuth,
		"wep_as_secondary_auth": wepAsSecondaryAuth,
	}
	data, e := NewAuthValue(AuthValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
