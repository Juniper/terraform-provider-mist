package mist_utils

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func VlanAsString(vlanId models.VlanIdWithVariable) basetypes.StringValue {
	if v, ok := vlanId.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := vlanId.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func GatewayVlanAsString(vlanId models.GatewayPortVlanIdWithVariable) basetypes.StringValue {
	if v, ok := vlanId.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := vlanId.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func WlanDynamicVlanAsString(vlanId models.WlanDynamicVlanDefaultVlanId) basetypes.StringValue {
	if v, ok := vlanId.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := vlanId.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func PskVlanAsString(vlanId models.PskVlanId) basetypes.StringValue {
	if v, ok := vlanId.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := vlanId.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func SwitchPortUsageReauthIntervalAsString(vlanId models.SwitchPortUsageReauthInterval) basetypes.StringValue {
	if v, ok := vlanId.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := vlanId.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func DscpAsString(vlanId models.Dscp) basetypes.StringValue {
	if v, ok := vlanId.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := vlanId.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func WlanBonjourAdditionalVlanIdsAsArrayOfString(diags *diag.Diagnostics, vlanIds models.AdditionalVlanIds) basetypes.ListValue {
	var items []attr.Value
	if stringIds, stringOk := vlanIds.AsString(); stringOk {
		for _, id := range strings.Split(*stringIds, ",") {
			if id != "" {
				items = append(items, types.StringValue(id))
			}
		}
	} else if arrayIds, arrayOk := vlanIds.AsArrayOfVlanIdWithVariable8(); arrayOk {
		for _, id := range *arrayIds {
			items = append(items, VlanAsString(id))
		}
	}
	list, e := types.ListValue(basetypes.StringType{}, items)
	diags.Append(e...)
	return list
}

func WlanVlanIdsAsArrayOfString(diags *diag.Diagnostics, vlanIds *models.WlanVlanIds) basetypes.ListValue {
	var items []attr.Value
	if stringIds, stringOk := vlanIds.AsString(); stringOk {
		for _, id := range strings.Split(*stringIds, ",") {
			if id != "" {
				items = append(items, types.StringValue(id))
			}
		}
	} else if arrayIds, arrayOk := vlanIds.AsArrayOfVlanIdWithVariable5(); arrayOk {
		for _, id := range *arrayIds {
			items = append(items, VlanAsString(id))
		}
	}
	list, e := types.ListValue(basetypes.StringType{}, items)
	diags.Append(e...)
	return list
}

func BgpAsAsString(bgpAs *models.BgpAs) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}
