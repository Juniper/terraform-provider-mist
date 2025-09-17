package mist_utils

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func WlanVlanAsString(vlanId models.WlanVlanIdWithVariable) basetypes.StringValue {
	if v, ok := vlanId.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := vlanId.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}
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

func WlanDynamicVlanDefaultVlanIdDeprecatedAsString(vlanId models.WlanDynamicVlanDefaultVlanIdDeprecated) basetypes.StringValue {
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
	} else if arrayIds, arrayOk := vlanIds.AsArrayOfVlanIdWithVariable7(); arrayOk {
		for _, id := range *arrayIds {
			items = append(items, VlanAsString(id))
		}
	}
	if len(items) > 0 {
		list, e := types.ListValue(basetypes.StringType{}, items)
		diags.Append(e...)
		return list
	} else {
		return types.ListNull(types.StringType)
	}
}

func WlanVlanIdsAsArrayOfString(diags *diag.Diagnostics, vlanIds *models.WlanVlanIds) basetypes.ListValue {
	var items []attr.Value
	if stringIds, stringOk := vlanIds.AsString(); stringOk {
		for _, id := range strings.Split(*stringIds, ",") {
			if id != "" {
				items = append(items, types.StringValue(id))
			}
		}
	} else if arrayIds, arrayOk := vlanIds.AsArrayOfVlanIdWithVariable4(); arrayOk {
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

func BgpLocalAsAsString(bgpAs *models.BgpLocalAs) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func RadiusCoaPortAsString(bgpAs *models.RadiusCoaPort) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func RadiusAcctPortAsString(bgpAs *models.RadiusAcctPort) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func RadiusAuthPortAsString(bgpAs *models.RadiusAuthPort) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func SyslogFilesAsString(bgpAs *models.RemoteSyslogArchiveFiles) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func SyslogPortAsString(bgpAs *models.RemoteSyslogServerPort) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func SwitchPortUsageMacLimitAsString(bgpAs *models.SwitchPortUsageMacLimit) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}
func SwitchPortOverwriteUsageMacLimitAsString(bgpAs *models.SwitchPortUsageMacLimitOverwrite) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func SwitchPortUsageMtuAsString(bgpAs *models.SwitchPortUsageMtu) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func SwitchMgmtMxedgeProxyPortsAsString(bgpAs *models.SwitchMgmtMxedgeProxyPort) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func GatewayPortConfigRethIdxAsString(bgpAs *models.GatewayPortConfigRethIdx) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func RadsecIdleTimeoutAsString(bgpAs *models.RadsecIdleTimeout) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func SponsorLinkValidityDurationAsString(bgpAs *models.SponsorLinkValidityDuration) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func GbpTagAsString(gbpTag models.NacTagGbpTag) basetypes.StringValue {
	if v, ok := gbpTag.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := gbpTag.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}

func WlanLimitAsString(bgpAs *models.WlanLimit) basetypes.StringValue {
	if v, ok := bgpAs.AsString(); ok {
		return types.StringValue(*v)
	} else if v, ok := bgpAs.AsNumber(); ok {
		return types.StringValue(fmt.Sprint(*v))
	} else {
		return types.StringNull()
	}
}
