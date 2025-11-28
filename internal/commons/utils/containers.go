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

// StringOrNumberContainer is an interface for SDK types that can hold either a string or number value
type StringOrNumberContainer interface {
	AsString() (*string, bool)
	AsNumber() (*int, bool)
}

// ContainerAsString converts any SDK container type that implements StringOrNumberContainer to a Terraform StringValue
func ContainerAsString[T StringOrNumberContainer](container T) basetypes.StringValue {
	if v, ok := container.AsString(); ok && v != nil {
		return types.StringValue(*v)
	} else if v, ok := container.AsNumber(); ok && v != nil {
		return types.StringValue(fmt.Sprintf("%d", *v))
	} else {
		return types.StringNull()
	}
}

// Legacy function wrappers for backward compatibility
func WlanVlanAsString(vlanId models.WlanVlanIdWithVariable) basetypes.StringValue {
	return ContainerAsString(&vlanId)
}

func VlanAsString(vlanId models.VlanIdWithVariable) basetypes.StringValue {
	return ContainerAsString(&vlanId)
}

func GatewayVlanAsString(vlanId models.GatewayPortVlanIdWithVariable) basetypes.StringValue {
	return ContainerAsString(&vlanId)
}

func WlanDynamicVlanAsString(vlanId models.WlanDynamicVlanDefaultVlanId) basetypes.StringValue {
	return ContainerAsString(&vlanId)
}

func WlanDynamicVlanDefaultVlanIdDeprecatedAsString(vlanId models.WlanDynamicVlanDefaultVlanIdDeprecated) basetypes.StringValue {
	return ContainerAsString(&vlanId)
}

func PskVlanAsString(vlanId models.PskVlanId) basetypes.StringValue {
	return ContainerAsString(&vlanId)
}

func SwitchPortUsageReauthIntervalAsString(vlanId models.SwitchPortUsageReauthInterval) basetypes.StringValue {
	return ContainerAsString(&vlanId)
}

func DscpAsString(vlanId models.Dscp) basetypes.StringValue {
	return ContainerAsString(&vlanId)
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
	return ContainerAsString(bgpAs)
}

func BgpLocalAsAsString(bgpAs *models.BgpLocalAs) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func RadiusCoaPortAsString(bgpAs *models.RadiusCoaPort) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func RadiusAcctPortAsString(bgpAs *models.RadiusAcctPort) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func RadiusAuthPortAsString(bgpAs *models.RadiusAuthPort) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func SyslogFilesAsString(bgpAs *models.RemoteSyslogArchiveFiles) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func SyslogPortAsString(bgpAs *models.RemoteSyslogServerPort) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func SwitchPortUsageMacLimitAsString(bgpAs *models.SwitchPortUsageMacLimit) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func SwitchPortOverwriteUsageMacLimitAsString(bgpAs *models.SwitchPortUsageMacLimitOverwrite) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func SwitchPortUsageMtuAsString(bgpAs *models.SwitchPortUsageMtu) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func SwitchMgmtMxedgeProxyPortsAsString(bgpAs *models.SwitchMgmtMxedgeProxyPort) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func GatewayPortConfigRethIdxAsString(bgpAs *models.GatewayPortConfigRethIdx) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func RadsecIdleTimeoutAsString(bgpAs *models.RadsecIdleTimeout) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func SponsorLinkValidityDurationAsString(bgpAs *models.SponsorLinkValidityDuration) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func GbpTagAsString(gbpTag models.NacTagGbpTag) basetypes.StringValue {
	return ContainerAsString(&gbpTag)
}

func WlanLimitAsString(bgpAs *models.WlanLimit) basetypes.StringValue {
	return ContainerAsString(bgpAs)
}

func ServiceDscpAsString(container *models.ServiceDscp) basetypes.StringValue {
	return ContainerAsString(container)
}

func ServiceMaxJitterAsString(container *models.ServiceMaxJitter) basetypes.StringValue {
	return ContainerAsString(container)
}

func ServiceMaxLatencyAsString(container *models.ServiceMaxLatency) basetypes.StringValue {
	return ContainerAsString(container)
}

func ServiceMaxLossAsString(container *models.ServiceMaxLoss) basetypes.StringValue {
	return ContainerAsString(container)
}
