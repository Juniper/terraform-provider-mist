package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func portUsageStormControlSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SwitchPortUsageStormControl) basetypes.ObjectValue {
	var noBroadcast = types.BoolValue(false)
	var noMulticast basetypes.BoolValue
	var noRegisteredMulticast basetypes.BoolValue
	var noUnknownUnicast basetypes.BoolValue
	var percentage basetypes.Int64Value

	if d.NoBroadcast != nil {
		noBroadcast = types.BoolValue(*d.NoBroadcast)
	}
	if d.NoMulticast != nil {
		noMulticast = types.BoolValue(*d.NoMulticast)
	}
	if d.NoRegisteredMulticast != nil {
		noRegisteredMulticast = types.BoolValue(*d.NoRegisteredMulticast)
	}
	if d.NoUnknownUnicast != nil {
		noUnknownUnicast = types.BoolValue(*d.NoUnknownUnicast)
	}
	if d.Percentage != nil {
		percentage = types.Int64Value(int64(*d.Percentage))
	}

	dataMapValue := map[string]attr.Value{
		"no_broadcast":            noBroadcast,
		"no_multicast":            noMulticast,
		"no_registered_multicast": noRegisteredMulticast,
		"no_unknown_unicast":      noUnknownUnicast,
		"percentage":              percentage,
	}
	data, e := basetypes.NewObjectValue(StormControlValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func portUsageRulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SwitchPortUsageDynamicRule) basetypes.ListValue {
	var valueList []attr.Value
	for _, d := range l {
		var equals basetypes.StringValue
		var equalsAny = types.ListNull(types.StringType)
		var expression basetypes.StringValue
		var src = types.StringValue(string(d.Src))
		var usage basetypes.StringValue

		if d.Equals != nil {
			equals = types.StringValue(*d.Equals)
		}
		if d.EqualsAny != nil {
			equalsAny = mistutils.ListOfStringSdkToTerraform(d.EqualsAny)
		}
		if d.Expression != nil {
			expression = types.StringValue(*d.Expression)
		}
		if d.Usage != nil {
			usage = types.StringValue(*d.Usage)
		}

		dataMapValue := map[string]attr.Value{
			"equals":     equals,
			"equals_any": equalsAny,
			"expression": expression,
			"src":        src,
			"usage":      usage,
		}
		data, e := NewRulesValue(RulesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		valueList = append(valueList, data)
	}

	stateListType := RulesValue{}.Type(ctx)
	stateList, e := types.ListValueFrom(ctx, stateListType, valueList)
	diags.Append(e...)

	return stateList
}

func portUsagesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchPortUsage) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {
		var allNetworks basetypes.BoolValue
		var allowDhcpd basetypes.BoolValue
		var allowMultipleSupplicants basetypes.BoolValue
		var bypassAuthWhenServerDown basetypes.BoolValue
		var bypassAuthWhenServerDownForUnknownClient basetypes.BoolValue
		var communityVlanId basetypes.Int64Value
		var description basetypes.StringValue
		var disableAutoneg basetypes.BoolValue
		var disabled basetypes.BoolValue
		var duplex basetypes.StringValue
		var dynamicVlanNetworks = types.ListNull(types.StringType)
		var enableMacAuth basetypes.BoolValue
		var enableQos basetypes.BoolValue
		var guestNetwork basetypes.StringValue
		var interIsolationNetwork basetypes.BoolValue
		var interSwitchLink basetypes.BoolValue
		var macAuthOnly basetypes.BoolValue
		var macAuthPreferred basetypes.BoolValue
		var macAuthProtocol basetypes.StringValue
		var macLimit basetypes.StringValue
		var mode basetypes.StringValue
		var mtu basetypes.StringValue
		var networks = types.ListNull(types.StringType)
		var persistMac basetypes.BoolValue
		var poeDisabled basetypes.BoolValue
		var portAuth basetypes.StringValue
		var portNetwork basetypes.StringValue
		var reauthInterval basetypes.StringValue
		var resetDefaultWhen basetypes.StringValue
		var rules = types.ListNull(RulesValue{}.Type(ctx))
		var serverFailNetwork basetypes.StringValue
		var serverRejectNetwork basetypes.StringValue
		var speed basetypes.StringValue
		var stormControl = types.ObjectNull(StormControlValue{}.AttributeTypes(ctx))
		var stpEdge basetypes.BoolValue
		var stpNoRootPort basetypes.BoolValue
		var stpP2p basetypes.BoolValue
		var useVstp basetypes.BoolValue
		var voipNetwork basetypes.StringValue

		if d.AllNetworks != nil {
			allNetworks = types.BoolValue(*d.AllNetworks)
		}
		if d.AllowDhcpd != nil {
			allowDhcpd = types.BoolValue(*d.AllowDhcpd)
		}
		if d.AllowMultipleSupplicants != nil {
			allowMultipleSupplicants = types.BoolValue(*d.AllowMultipleSupplicants)
		}
		if d.BypassAuthWhenServerDown != nil {
			bypassAuthWhenServerDown = types.BoolValue(*d.BypassAuthWhenServerDown)
		}
		if d.BypassAuthWhenServerDownForUnknownClient != nil {
			bypassAuthWhenServerDownForUnknownClient = types.BoolValue(*d.BypassAuthWhenServerDownForUnknownClient)
		}
		if d.CommunityVlanId != nil {
			communityVlanId = types.Int64Value(int64(*d.CommunityVlanId))
		}
		if d.Description != nil {
			description = types.StringValue(*d.Description)
		}
		if d.DisableAutoneg != nil {
			disableAutoneg = types.BoolValue(*d.DisableAutoneg)
		}
		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.Duplex != nil {
			duplex = types.StringValue(string(*d.Duplex))
		}
		if d.DynamicVlanNetworks != nil {
			dynamicVlanNetworks = mistutils.ListOfStringSdkToTerraform(d.DynamicVlanNetworks)
		}
		if d.EnableMacAuth != nil {
			enableMacAuth = types.BoolValue(*d.EnableMacAuth)
		}
		if d.EnableQos != nil {
			enableQos = types.BoolValue(*d.EnableQos)
		}
		if d.GuestNetwork.Value() != nil {
			guestNetwork = types.StringValue(*d.GuestNetwork.Value())
		}
		if d.InterIsolationNetworkLink != nil {
			interIsolationNetwork = types.BoolValue(*d.InterIsolationNetworkLink)
		}
		if d.InterSwitchLink != nil {
			interSwitchLink = types.BoolValue(*d.InterSwitchLink)
		}
		if d.MacAuthOnly != nil {
			macAuthOnly = types.BoolValue(*d.MacAuthOnly)
		}
		if d.MacAuthPreferred != nil {
			macAuthPreferred = types.BoolValue(*d.MacAuthPreferred)
		}
		if d.MacAuthProtocol != nil {
			macAuthProtocol = types.StringValue(string(*d.MacAuthProtocol))
		}
		if d.MacLimit != nil {
			macLimit = mistutils.SwitchPortUsageMacLimitAsString(d.MacLimit)
		}
		if d.Mode != nil {
			mode = types.StringValue(string(*d.Mode))
		}
		if d.Mtu != nil {
			mtu = mistutils.SwitchPortUsageMtuAsString(d.Mtu)
		}
		if d.Networks != nil {
			networks = mistutils.ListOfStringSdkToTerraform(d.Networks)
		}
		if d.PersistMac != nil {
			persistMac = types.BoolValue(*d.PersistMac)
		}
		if d.PoeDisabled != nil {
			poeDisabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.PortAuth.Value() != nil {
			portAuth = types.StringValue(string(*d.PortAuth.Value()))
		}
		if d.PortNetwork != nil {
			portNetwork = types.StringValue(*d.PortNetwork)
		}
		if d.ReauthInterval != nil {
			reauthInterval = mistutils.SwitchPortUsageReauthIntervalAsString(*d.ReauthInterval)
		}
		if d.ResetDefaultWhen != nil {
			resetDefaultWhen = types.StringValue(string(*d.ResetDefaultWhen))
		}
		if d.Rules != nil {
			rules = portUsageRulesSdkToTerraform(ctx, diags, d.Rules)
		}
		if d.ServerFailNetwork.Value() != nil {
			serverFailNetwork = types.StringValue(*d.ServerFailNetwork.Value())
		}
		if d.ServerRejectNetwork.Value() != nil {
			serverRejectNetwork = types.StringValue(*d.ServerRejectNetwork.Value())
		}
		if d.Speed != nil {
			speed = types.StringValue(string(*d.Speed))
		}
		if d.StormControl != nil {
			stormControl = portUsageStormControlSdkToTerraform(ctx, diags, *d.StormControl)
		}
		if d.StpEdge != nil {
			stpEdge = types.BoolValue(*d.StpEdge)
		}
		if d.StpNoRootPort != nil {
			stpNoRootPort = types.BoolValue(*d.StpNoRootPort)
		}
		if d.StpP2p != nil {
			stpP2p = types.BoolValue(*d.StpP2p)
		}
		if d.UseVstp != nil {
			useVstp = types.BoolValue(*d.UseVstp)
		}
		if d.VoipNetwork.Value() != nil {
			voipNetwork = types.StringValue(*d.VoipNetwork.Value())
		}

		dataMapValue := map[string]attr.Value{
			"all_networks":                                    allNetworks,
			"allow_dhcpd":                                     allowDhcpd,
			"allow_multiple_supplicants":                      allowMultipleSupplicants,
			"bypass_auth_when_server_down":                    bypassAuthWhenServerDown,
			"bypass_auth_when_server_down_for_unknown_client": bypassAuthWhenServerDownForUnknownClient,
			"community_vlan_id":                               communityVlanId,
			"description":                                     description,
			"disable_autoneg":                                 disableAutoneg,
			"disabled":                                        disabled,
			"duplex":                                          duplex,
			"dynamic_vlan_networks":                           dynamicVlanNetworks,
			"enable_mac_auth":                                 enableMacAuth,
			"enable_qos":                                      enableQos,
			"guest_network":                                   guestNetwork,
			"inter_isolation_network_link":                    interIsolationNetwork,
			"inter_switch_link":                               interSwitchLink,
			"mac_auth_only":                                   macAuthOnly,
			"mac_auth_preferred":                              macAuthPreferred,
			"mac_auth_protocol":                               macAuthProtocol,
			"mac_limit":                                       macLimit,
			"mode":                                            mode,
			"mtu":                                             mtu,
			"networks":                                        networks,
			"persist_mac":                                     persistMac,
			"poe_disabled":                                    poeDisabled,
			"port_auth":                                       portAuth,
			"port_network":                                    portNetwork,
			"reauth_interval":                                 reauthInterval,
			"reset_default_when":                              resetDefaultWhen,
			"rules":                                           rules,
			"server_fail_network":                             serverFailNetwork,
			"server_reject_network":                           serverRejectNetwork,
			"speed":                                           speed,
			"storm_control":                                   stormControl,
			"stp_edge":                                        stpEdge,
			"stp_no_root_port":                                stpNoRootPort,
			"stp_p2p":                                         stpP2p,
			"use_vstp":                                        useVstp,
			"voip_network":                                    voipNetwork,
		}
		data, e := NewPortUsagesValue(PortUsagesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMap[k] = data
	}
	stateType := PortUsagesValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, stateValueMap)
	diags.Append(e...)
	return stateResult
}
