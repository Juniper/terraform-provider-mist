package resource_device_switch

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func localPortConfigStormControlSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SwitchPortLocalUsageStormControl) basetypes.ObjectValue {
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
func localPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.JunosLocalPortConfig) basetypes.MapValue {
	map_item_value := make(map[string]attr.Value)
	map_item_type := PortConfigValue{}.Type(ctx)
	for k, d := range m {

		item_obj := NewLocalPortConfigValueUnknown()
		item_obj.Usage = types.StringValue(d.Usage)

		if d.AllNetworks != nil {
			item_obj.AllNetworks = types.BoolValue(*d.AllNetworks)
		}
		if d.AllowDhcpd != nil {
			item_obj.AllowDhcpd = types.BoolValue(*d.AllowDhcpd)
		}
		if d.AllowMultipleSupplicants != nil {
			item_obj.AllowMultipleSupplicants = types.BoolValue(*d.AllowMultipleSupplicants)
		}
		if d.BypassAuthWhenServerDown != nil {
			item_obj.BypassAuthWhenServerDown = types.BoolValue(*d.BypassAuthWhenServerDown)
		}
		if d.BypassAuthWhenServerDownForUnkonwnClient != nil {
			item_obj.BypassAuthWhenServerDownForUnkonwnClient = types.BoolValue(*d.BypassAuthWhenServerDownForUnkonwnClient)
		}
		if d.Description != nil {
			item_obj.Description = types.StringValue(*d.Description)
		}
		if d.DisableAutoneg != nil {
			item_obj.DisableAutoneg = types.BoolValue(*d.DisableAutoneg)
		}
		if d.Disabled != nil {
			item_obj.Disabled = types.BoolValue(*d.Disabled)
		}
		if d.Duplex != nil {
			item_obj.Duplex = types.StringValue(string(*d.Duplex))
		}
		if d.DynamicVlanNetworks != nil {
			item_obj.DynamicVlanNetworks = mist_transform.ListOfStringSdkToTerraform(ctx, d.DynamicVlanNetworks)
		}
		if d.EnableMacAuth != nil {
			item_obj.EnableMacAuth = types.BoolValue(*d.EnableMacAuth)
		}
		if d.EnableQos != nil {
			item_obj.EnableQos = types.BoolValue(*d.EnableQos)
		}
		if d.GuestNetwork.Value() != nil {
			item_obj.GuestNetwork = types.StringValue(*d.GuestNetwork.Value())
		}
		if d.InterSwitchLink != nil {
			item_obj.InterSwitchLink = types.BoolValue(*d.InterSwitchLink)
		}
		if d.MacAuthOnly != nil {
			item_obj.MacAuthOnly = types.BoolValue(*d.MacAuthOnly)
		}
		if d.MacAuthPreferred != nil {
			item_obj.MacAuthPreferred = types.BoolValue(*d.MacAuthPreferred)
		}
		if d.MacAuthProtocol != nil {
			item_obj.MacAuthProtocol = types.StringValue(string(*d.MacAuthProtocol))
		}
		if d.MacLimit != nil {
			item_obj.MacLimit = types.Int64Value(int64(*d.MacLimit))
		}
		if d.Mode != nil {
			item_obj.Mode = types.StringValue(string(*d.Mode))
		}
		if d.Mtu != nil {
			item_obj.Mtu = types.Int64Value(int64(*d.Mtu))
		}
		if d.Networks != nil {
			item_obj.Networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Networks)
		}
		if d.PersistMac != nil {
			item_obj.PersistMac = types.BoolValue(*d.PersistMac)
		}
		if d.PoeDisabled != nil {
			item_obj.PoeDisabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.PortAuth.Value() != nil {
			item_obj.PortAuth = types.StringValue(string(*d.PortAuth.Value()))
		}
		if d.PortNetwork != nil {
			item_obj.PortNetwork = types.StringValue(*d.PortNetwork)
		}
		if d.ReauthInterval != nil {
			item_obj.ReauthInterval = types.Int64Value(int64(*d.ReauthInterval))
		}
		if d.ServerFailNetwork.Value() != nil {
			item_obj.ServerFailNetwork = types.StringValue(*d.ServerFailNetwork.Value())
		}
		if d.ServerRejectNetwork.Value() != nil {
			item_obj.ServerRejectNetwork = types.StringValue(*d.ServerRejectNetwork.Value())
		}
		if d.Speed != nil {
			item_obj.Speed = types.StringValue(string(*d.Speed))
		}
		if d.StormControl != nil {
			item_obj.StormControl = localPortConfigStormControlSdkToTerraform(ctx, diags, *d.StormControl)
		}
		if d.StpEdge != nil {
			item_obj.StpEdge = types.BoolValue(*d.StpEdge)
		}
		if d.StpNoRootPort != nil {
			item_obj.StpNoRootPort = types.BoolValue(*d.StpNoRootPort)
		}
		if d.StpP2p != nil {
			item_obj.StpP2p = types.BoolValue(*d.StpP2p)
		}
		if d.UseVstp != nil {
			item_obj.UseVstp = types.BoolValue(*d.UseVstp)
		}

		if d.VoipNetwork != nil {
			item_obj.VoipNetwork = types.StringValue(*d.VoipNetwork)
		}

		map_item_value[k] = item_obj
	}
	r, e := types.MapValueFrom(ctx, map_item_type, map_item_value)
	diags.Append(e...)
	return r
}
