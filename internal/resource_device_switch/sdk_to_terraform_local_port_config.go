package resource_device_switch

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func localPortConfigStormControlSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SwitchPortLocalUsageStormControl) basetypes.ObjectValue {
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
func localPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.JunosLocalPortConfig) basetypes.MapValue {
	mapItemValue := make(map[string]attr.Value)

	for k, d := range m {

		var usage = types.StringValue(d.Usage)
		var allNetworks basetypes.BoolValue
		var allowDhcpd basetypes.BoolValue
		var allowMultipleSupplicants basetypes.BoolValue
		var bypassAuthWhenServerDown basetypes.BoolValue
		var bypassAuthWhenServerDownForUnknownClient basetypes.BoolValue
		var description basetypes.StringValue
		var disableAutoneg basetypes.BoolValue
		var disabled basetypes.BoolValue
		var duplex basetypes.StringValue
		var dynamicVlanNetworks = types.ListNull(types.StringType)
		var enableMacAuth basetypes.BoolValue
		var enableQos basetypes.BoolValue
		var guestNetwork basetypes.StringValue
		var interSwitchLink basetypes.BoolValue
		var macAuthOnly basetypes.BoolValue
		var macAuthPreferred basetypes.BoolValue
		var macAuthProtocol basetypes.StringValue
		var macLimit basetypes.Int64Value
		var mode basetypes.StringValue
		var mtu basetypes.Int64Value
		var networks = types.ListNull(types.StringType)
		var note basetypes.StringValue
		var persistMac basetypes.BoolValue
		var poeDisabled basetypes.BoolValue
		var portAuth basetypes.StringValue
		var portNetwork basetypes.StringValue
		var reauthInterval basetypes.StringValue
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
			macLimit = types.Int64Value(int64(*d.MacLimit))
		}
		if d.Mode != nil {
			mode = types.StringValue(string(*d.Mode))
		}
		if d.Mtu != nil {
			mtu = types.Int64Value(int64(*d.Mtu))
		}
		if d.Networks != nil {
			networks = mistutils.ListOfStringSdkToTerraform(d.Networks)
		}
		if d.Note != nil {
			note = types.StringValue(*d.Note)
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
			stormControl = localPortConfigStormControlSdkToTerraform(ctx, diags, *d.StormControl)
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

		if d.VoipNetwork != nil {
			voipNetwork = types.StringValue(*d.VoipNetwork)
		}

		dataMapValue := map[string]attr.Value{
			"note":                         note,
			"all_networks":                 allNetworks,
			"allow_dhcpd":                  allowDhcpd,
			"allow_multiple_supplicants":   allowMultipleSupplicants,
			"bypass_auth_when_server_down": bypassAuthWhenServerDown,
			"bypass_auth_when_server_down_for_unknown_client": bypassAuthWhenServerDownForUnknownClient,
			"description":           description,
			"disable_autoneg":       disableAutoneg,
			"disabled":              disabled,
			"duplex":                duplex,
			"dynamic_vlan_networks": dynamicVlanNetworks,
			"enable_mac_auth":       enableMacAuth,
			"enable_qos":            enableQos,
			"guest_network":         guestNetwork,
			"inter_switch_link":     interSwitchLink,
			"mac_auth_only":         macAuthOnly,
			"mac_auth_preferred":    macAuthPreferred,
			"mac_auth_protocol":     macAuthProtocol,
			"mac_limit":             macLimit,
			"mode":                  mode,
			"mtu":                   mtu,
			"networks":              networks,
			"persist_mac":           persistMac,
			"poe_disabled":          poeDisabled,
			"port_auth":             portAuth,
			"port_network":          portNetwork,
			"reauth_interval":       reauthInterval,
			"server_fail_network":   serverFailNetwork,
			"server_reject_network": serverRejectNetwork,
			"speed":                 speed,
			"storm_control":         stormControl,
			"stp_edge":              stpEdge,
			"stp_no_root_port":      stpNoRootPort,
			"stp_p2p":               stpP2p,
			"usage":                 usage,
			"use_vstp":              useVstp,
			"voip_network":          voipNetwork,
		}
		data, e := NewLocalPortConfigValue(LocalPortConfigValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapItemValue[k] = data
	}
	r, e := types.MapValueFrom(ctx, LocalPortConfigValue{}.Type(ctx), mapItemValue)
	diags.Append(e...)
	return r
}
