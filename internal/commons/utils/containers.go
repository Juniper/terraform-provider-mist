package mist_utils

import (
	"fmt"

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
