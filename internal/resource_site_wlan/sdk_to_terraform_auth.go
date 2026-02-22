package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func authSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.WlanAuth) AuthValue {
	if data == nil {
		return NewAuthValueNull()
	}

	var anticlogThreshold basetypes.Int64Value
	if data.AnticlogThreshold != nil {
		anticlogThreshold = types.Int64Value(int64(*data.AnticlogThreshold))
	}

	var eapReauth basetypes.BoolValue
	if data.EapReauth != nil {
		eapReauth = types.BoolValue(*data.EapReauth)
	}

	var enableMacAuth basetypes.BoolValue
	if data.EnableMacAuth != nil {
		enableMacAuth = types.BoolValue(*data.EnableMacAuth)
	}

	var keyIdx basetypes.Int64Value
	if data.KeyIdx != nil {
		keyIdx = types.Int64Value(int64(*data.KeyIdx))
	}

	var keysList []attr.Value
	if data.Keys != nil {
		for _, val := range data.Keys {
			keysList = append(keysList, types.StringValue(val))
		}
	}
	keys := types.ListValueMust(basetypes.StringType{}, keysList)

	var multiPskOnly basetypes.BoolValue
	if data.MultiPskOnly != nil {
		multiPskOnly = types.BoolValue(*data.MultiPskOnly)
	}

	var owe basetypes.StringValue
	if data.Owe != nil {
		owe = types.StringValue(string(*data.Owe))
	}

	var pairwise = types.ListNull(types.StringType)
	if data.Pairwise != nil {
		var pairwiseList []attr.Value
		for _, val := range data.Pairwise {
			pairwiseList = append(pairwiseList, types.StringValue(string(val)))
		}
		pairwise = types.ListValueMust(basetypes.StringType{}, pairwiseList)
	}

	var privateWlan = types.BoolValue(false)
	if data.PrivateWlan != nil {
		privateWlan = types.BoolValue(*data.PrivateWlan)
	}

	var psk = types.StringValue("")
	if data.Psk.Value() != nil {
		psk = types.StringValue(*data.Psk.Value())
	}

	typeAuth := types.StringValue(string(data.Type))

	var wepAsSecondaryAuth basetypes.BoolValue
	if data.WepAsSecondaryAuth != nil {
		wepAsSecondaryAuth = types.BoolValue(*data.WepAsSecondaryAuth)
	}

	dataMap := map[string]attr.Value{
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
	result, err := NewAuthValue(AuthValue{}.AttributeTypes(ctx), dataMap)
	diags.Append(err...)

	return result
}
