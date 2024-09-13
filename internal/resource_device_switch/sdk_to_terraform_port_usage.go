package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func portUsageStormControlSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SwitchPortUsageStormControl) basetypes.ObjectValue {
	var no_broadcast basetypes.BoolValue = types.BoolValue(false)
	var no_multicast basetypes.BoolValue
	var no_registered_multicast basetypes.BoolValue
	var no_unknown_unicast basetypes.BoolValue
	var percentage basetypes.Int64Value

	if d.NoBroadcast != nil {
		no_broadcast = types.BoolValue(*d.NoBroadcast)
	}
	if d.NoMulticast != nil {
		no_multicast = types.BoolValue(*d.NoMulticast)
	}
	if d.NoRegisteredMulticast != nil {
		no_registered_multicast = types.BoolValue(*d.NoRegisteredMulticast)
	}
	if d.NoUnknownUnicast != nil {
		no_unknown_unicast = types.BoolValue(*d.NoUnknownUnicast)
	}
	if d.Percentage != nil {
		percentage = types.Int64Value(int64(*d.Percentage))
	}

	data_map_attr_type := StormControlValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"no_broadcast":            no_broadcast,
		"no_multicast":            no_multicast,
		"no_registered_multicast": no_registered_multicast,
		"no_unknown_unicast":      no_unknown_unicast,
		"percentage":              percentage,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func portUsageRulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SwitchPortUsageDynamicRule) basetypes.ListValue {
	var value_list []attr.Value
	for _, d := range l {
		var equals basetypes.StringValue
		var equals_any basetypes.ListValue = types.ListNull(types.StringType)
		var expression basetypes.StringValue
		var src basetypes.StringValue = types.StringValue(string(d.Src))
		var usage basetypes.StringValue

		if d.Equals != nil {
			equals = types.StringValue(*d.Equals)
		}
		if d.EqualsAny != nil {
			equals_any = mist_transform.ListOfStringSdkToTerraform(ctx, d.EqualsAny)
		}
		if d.Expression != nil {
			expression = types.StringValue(*d.Expression)
		}
		if d.Usage != nil {
			usage = types.StringValue(*d.Usage)
		}

		data_map_attr_type := RulesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"equals":     equals,
			"equals_any": equals_any,
			"expression": expression,
			"src":        src,
			"usage":      usage,
		}
		data, e := NewRulesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		value_list = append(value_list, data)
	}

	state_list_type := RulesValue{}.Type(ctx)
	state_list, e := types.ListValueFrom(ctx, state_list_type, value_list)
	diags.Append(e...)

	return state_list
}

func portUsagesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchPortUsage) basetypes.MapValue {
	state_value_map := make(map[string]attr.Value)
	for k, d := range m {
		var all_networks basetypes.BoolValue
		var allow_dhcpd basetypes.BoolValue
		var allow_multiple_supplicants basetypes.BoolValue
		var bypass_auth_when_server_down basetypes.BoolValue
		var bypass_auth_when_server_down_for_unkonwn_client basetypes.BoolValue
		var description basetypes.StringValue
		var disable_autoneg basetypes.BoolValue
		var disabled basetypes.BoolValue
		var duplex basetypes.StringValue
		var dynamic_vlan_networks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var enable_mac_auth basetypes.BoolValue
		var enable_qos basetypes.BoolValue
		var guest_network basetypes.StringValue
		var inter_switch_link basetypes.BoolValue
		var mac_auth_only basetypes.BoolValue
		var mac_auth_preferred basetypes.BoolValue
		var mac_auth_protocol basetypes.StringValue
		var mac_limit basetypes.Int64Value
		var mode basetypes.StringValue
		var mtu basetypes.Int64Value
		var networks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var persist_mac basetypes.BoolValue
		var poe_disabled basetypes.BoolValue
		var port_auth basetypes.StringValue
		var port_network basetypes.StringValue
		var reauth_interval basetypes.Int64Value
		var reset_default_when basetypes.StringValue
		var rules basetypes.ListValue = types.ListNull(RulesValue{}.Type(ctx))
		var server_fail_network basetypes.StringValue
		var server_reject_network basetypes.StringValue
		var speed basetypes.StringValue
		var storm_control basetypes.ObjectValue = types.ObjectNull(StormControlValue{}.AttributeTypes(ctx))
		var stp_edge basetypes.BoolValue
		var stp_no_root_port basetypes.BoolValue
		var stp_p2p basetypes.BoolValue
		var voip_network basetypes.StringValue

		if d.AllNetworks != nil {
			all_networks = types.BoolValue(*d.AllNetworks)
		}
		if d.AllowDhcpd != nil {
			allow_dhcpd = types.BoolValue(*d.AllowDhcpd)
		}
		if d.AllowMultipleSupplicants != nil {
			allow_multiple_supplicants = types.BoolValue(*d.AllowMultipleSupplicants)
		}
		if d.BypassAuthWhenServerDown != nil {
			bypass_auth_when_server_down = types.BoolValue(*d.BypassAuthWhenServerDown)
		}
		if d.BypassAuthWhenServerDownForUnkonwnClient != nil {
			bypass_auth_when_server_down_for_unkonwn_client = types.BoolValue(*d.BypassAuthWhenServerDownForUnkonwnClient)
		}
		if d.Description != nil {
			description = types.StringValue(*d.Description)
		}
		if d.DisableAutoneg != nil {
			disable_autoneg = types.BoolValue(*d.DisableAutoneg)
		}
		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.Duplex != nil {
			duplex = types.StringValue(string(*d.Duplex))
		}
		if d.DynamicVlanNetworks != nil {
			dynamic_vlan_networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.DynamicVlanNetworks)
		}
		if d.EnableMacAuth != nil {
			enable_mac_auth = types.BoolValue(*d.EnableMacAuth)
		}
		if d.EnableQos != nil {
			enable_qos = types.BoolValue(*d.EnableQos)
		}
		if d.GuestNetwork.Value() != nil {
			guest_network = types.StringValue(*d.GuestNetwork.Value())
		}
		if d.InterSwitchLink != nil {
			inter_switch_link = types.BoolValue(*d.InterSwitchLink)
		}
		if d.MacAuthOnly != nil {
			mac_auth_only = types.BoolValue(*d.MacAuthOnly)
		}
		if d.MacAuthPreferred != nil {
			mac_auth_preferred = types.BoolValue(*d.MacAuthPreferred)
		}
		if d.MacAuthProtocol != nil {
			mac_auth_protocol = types.StringValue(string(*d.MacAuthProtocol))
		}
		if d.MacLimit != nil {
			mac_limit = types.Int64Value(int64(*d.MacLimit))
		}
		if d.Mode != nil {
			mode = types.StringValue(string(*d.Mode))
		}
		if d.Mtu != nil {
			mtu = types.Int64Value(int64(*d.Mtu))
		}
		if d.Networks != nil {
			networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Networks)
		}
		if d.PersistMac != nil {
			persist_mac = types.BoolValue(*d.PersistMac)
		}
		if d.PoeDisabled != nil {
			poe_disabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.PortAuth.Value() != nil {
			port_auth = types.StringValue(string(*d.PortAuth.Value()))
		}
		if d.PortNetwork != nil {
			port_network = types.StringValue(*d.PortNetwork)
		}
		if d.ReauthInterval != nil {
			reauth_interval = types.Int64Value(int64(*d.ReauthInterval))
		}
		if d.ResetDefaultWhen != nil {
			reset_default_when = types.StringValue(string(*d.ResetDefaultWhen))
		}
		if d.Rules != nil {
			rules = portUsageRulesSdkToTerraform(ctx, diags, d.Rules)
		}
		if d.ServerFailNetwork.Value() != nil {
			server_fail_network = types.StringValue(*d.ServerFailNetwork.Value())
		}
		if d.ServerRejectNetwork.Value() != nil {
			server_reject_network = types.StringValue(*d.ServerRejectNetwork.Value())
		}
		if d.Speed != nil {
			speed = types.StringValue(*d.Speed)
		}
		if d.StormControl != nil {
			storm_control = portUsageStormControlSdkToTerraform(ctx, diags, *d.StormControl)
		}
		if d.StpEdge != nil {
			stp_edge = types.BoolValue(*d.StpEdge)
		}
		if d.StpNoRootPort != nil {
			stp_no_root_port = types.BoolValue(*d.StpNoRootPort)
		}
		if d.StpP2p != nil {
			stp_p2p = types.BoolValue(*d.StpP2p)
		}
		if d.VoipNetwork != nil {
			voip_network = types.StringValue(*d.VoipNetwork)
		}

		data_map_attr_type := PortUsagesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"all_networks":                                    all_networks,
			"allow_dhcpd":                                     allow_dhcpd,
			"allow_multiple_supplicants":                      allow_multiple_supplicants,
			"bypass_auth_when_server_down":                    bypass_auth_when_server_down,
			"bypass_auth_when_server_down_for_unkonwn_client": bypass_auth_when_server_down_for_unkonwn_client,
			"description":                                     description,
			"disable_autoneg":                                 disable_autoneg,
			"disabled":                                        disabled,
			"duplex":                                          duplex,
			"dynamic_vlan_networks":                           dynamic_vlan_networks,
			"enable_mac_auth":                                 enable_mac_auth,
			"enable_qos":                                      enable_qos,
			"guest_network":                                   guest_network,
			"inter_switch_link":                               inter_switch_link,
			"mac_auth_only":                                   mac_auth_only,
			"mac_auth_preferred":                              mac_auth_preferred,
			"mac_auth_protocol":                               mac_auth_protocol,
			"mac_limit":                                       mac_limit,
			"mode":                                            mode,
			"mtu":                                             mtu,
			"networks":                                        networks,
			"persist_mac":                                     persist_mac,
			"poe_disabled":                                    poe_disabled,
			"port_auth":                                       port_auth,
			"port_network":                                    port_network,
			"reauth_interval":                                 reauth_interval,
			"reset_default_when":                              reset_default_when,
			"rules":                                           rules,
			"server_fail_network":                             server_fail_network,
			"server_reject_network":                           server_reject_network,
			"speed":                                           speed,
			"storm_control":                                   storm_control,
			"stp_edge":                                        stp_edge,
			"stp_no_root_port":                                stp_no_root_port,
			"stp_p2p":                                         stp_p2p,
			"voip_network":                                    voip_network,
		}
		data, e := NewPortUsagesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}
	state_type := PortUsagesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
