package resource_device_switch

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
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
	mapItemType := PortConfigValue{}.Type(ctx)
	for k, d := range m {

		itemObj := NewLocalPortConfigValueUnknown()
		itemObj.Usage = types.StringValue(d.Usage)

		if d.AllNetworks != nil {
			itemObj.AllNetworks = types.BoolValue(*d.AllNetworks)
		}
		if d.AllowDhcpd != nil {
			itemObj.AllowDhcpd = types.BoolValue(*d.AllowDhcpd)
		}
		if d.AllowMultipleSupplicants != nil {
			itemObj.AllowMultipleSupplicants = types.BoolValue(*d.AllowMultipleSupplicants)
		}
		if d.BypassAuthWhenServerDown != nil {
			itemObj.BypassAuthWhenServerDown = types.BoolValue(*d.BypassAuthWhenServerDown)
		}
		if d.BypassAuthWhenServerDownForUnknownClient != nil {
			itemObj.BypassAuthWhenServerDownForUnknownClient = types.BoolValue(*d.BypassAuthWhenServerDownForUnknownClient)
		}
		if d.Description != nil {
			itemObj.Description = types.StringValue(*d.Description)
		}
		if d.DisableAutoneg != nil {
			itemObj.DisableAutoneg = types.BoolValue(*d.DisableAutoneg)
		}
		if d.Disabled != nil {
			itemObj.Disabled = types.BoolValue(*d.Disabled)
		}
		if d.Duplex != nil {
			itemObj.Duplex = types.StringValue(string(*d.Duplex))
		}
		if d.DynamicVlanNetworks != nil {
			itemObj.DynamicVlanNetworks = misttransform.ListOfStringSdkToTerraform(d.DynamicVlanNetworks)
		}
		if d.EnableMacAuth != nil {
			itemObj.EnableMacAuth = types.BoolValue(*d.EnableMacAuth)
		}
		if d.EnableQos != nil {
			itemObj.EnableQos = types.BoolValue(*d.EnableQos)
		}
		if d.GuestNetwork.Value() != nil {
			itemObj.GuestNetwork = types.StringValue(*d.GuestNetwork.Value())
		}
		if d.InterSwitchLink != nil {
			itemObj.InterSwitchLink = types.BoolValue(*d.InterSwitchLink)
		}
		if d.MacAuthOnly != nil {
			itemObj.MacAuthOnly = types.BoolValue(*d.MacAuthOnly)
		}
		if d.MacAuthPreferred != nil {
			itemObj.MacAuthPreferred = types.BoolValue(*d.MacAuthPreferred)
		}
		if d.MacAuthProtocol != nil {
			itemObj.MacAuthProtocol = types.StringValue(string(*d.MacAuthProtocol))
		}
		if d.MacLimit != nil {
			itemObj.MacLimit = types.Int64Value(int64(*d.MacLimit))
		}
		if d.Mode != nil {
			itemObj.Mode = types.StringValue(string(*d.Mode))
		}
		if d.Mtu != nil {
			itemObj.Mtu = types.Int64Value(int64(*d.Mtu))
		}
		if d.Networks != nil {
			itemObj.Networks = misttransform.ListOfStringSdkToTerraform(d.Networks)
		}
		if d.Note != nil {
			itemObj.Note = types.StringValue(*d.Note)
		}
		if d.PersistMac != nil {
			itemObj.PersistMac = types.BoolValue(*d.PersistMac)
		}
		if d.PoeDisabled != nil {
			itemObj.PoeDisabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.PortAuth.Value() != nil {
			itemObj.PortAuth = types.StringValue(string(*d.PortAuth.Value()))
		}
		if d.PortNetwork != nil {
			itemObj.PortNetwork = types.StringValue(*d.PortNetwork)
		}
		if d.ReauthInterval != nil {
			itemObj.ReauthInterval = types.Int64Value(int64(*d.ReauthInterval))
		}
		if d.ServerFailNetwork.Value() != nil {
			itemObj.ServerFailNetwork = types.StringValue(*d.ServerFailNetwork.Value())
		}
		if d.ServerRejectNetwork.Value() != nil {
			itemObj.ServerRejectNetwork = types.StringValue(*d.ServerRejectNetwork.Value())
		}
		if d.Speed != nil {
			itemObj.Speed = types.StringValue(string(*d.Speed))
		}
		if d.StormControl != nil {
			itemObj.StormControl = localPortConfigStormControlSdkToTerraform(ctx, diags, *d.StormControl)
		}
		if d.StpEdge != nil {
			itemObj.StpEdge = types.BoolValue(*d.StpEdge)
		}
		if d.StpNoRootPort != nil {
			itemObj.StpNoRootPort = types.BoolValue(*d.StpNoRootPort)
		}
		if d.StpP2p != nil {
			itemObj.StpP2p = types.BoolValue(*d.StpP2p)
		}
		if d.UseVstp != nil {
			itemObj.UseVstp = types.BoolValue(*d.UseVstp)
		}

		if d.VoipNetwork != nil {
			itemObj.VoipNetwork = types.StringValue(*d.VoipNetwork)
		}

		mapItemValue[k] = itemObj
	}
	r, e := types.MapValueFrom(ctx, mapItemType, mapItemValue)
	diags.Append(e...)
	return r
}
