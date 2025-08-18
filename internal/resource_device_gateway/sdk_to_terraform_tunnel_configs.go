package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func tunnelConfigAutoProvNodeSdkToTerraform(diags *diag.Diagnostics, d models.TunnelConfigAutoProvisionNode, rType map[string]attr.Type) basetypes.ObjectValue {

	var probeIps = types.ListNull(types.StringType)
	var wanNames = types.ListNull(types.StringType)

	if d.ProbeIps != nil {
		probeIps = mistutils.ListOfStringSdkToTerraform(d.ProbeIps)
	}
	if d.WanNames != nil {
		wanNames = mistutils.ListOfStringSdkToTerraform(d.WanNames)
	}

	rAttrValue := map[string]attr.Value{
		"probe_ips": probeIps,
		"wan_names": wanNames,
	}
	r, e := basetypes.NewObjectValue(rType, rAttrValue)
	diags.Append(e...)
	return r
}

func tunnelConfigAutoProvSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.TunnelConfigAutoProvision) basetypes.ObjectValue {

	var primary = types.ObjectNull(AutoProvisionPrimaryValue{}.AttributeTypes(ctx))
	var secondary = types.ObjectNull(AutoProvisionSecondaryValue{}.AttributeTypes(ctx))
	var enabled basetypes.BoolValue
	var latlng = types.ObjectNull(LatlngValue{}.AttributeTypes(ctx))
	var provider basetypes.StringValue
	var region basetypes.StringValue
	var serviceConnection basetypes.StringValue

	if d.Primary != nil {
		primary = tunnelConfigAutoProvNodeSdkToTerraform(diags, *d.Primary, AutoProvisionPrimaryValue{}.AttributeTypes(ctx))
	}
	if d.Secondary != nil {
		secondary = tunnelConfigAutoProvNodeSdkToTerraform(diags, *d.Secondary, AutoProvisionSecondaryValue{}.AttributeTypes(ctx))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Latlng != nil {
		latlngValue := map[string]attr.Value{
			"lat": types.Float64Value(d.Latlng.Lat),
			"lng": types.Float64Value(d.Latlng.Lng),
		}
		tmp, e := NewLatlngValue(LatlngValue{}.AttributeTypes(ctx), latlngValue)
		diags.Append(e...)
		latlng, _ = tmp.ToObjectValue(ctx)
	}

	provider = types.StringValue(string(d.Provider))

	if d.Region != nil {
		region = types.StringValue(*d.Region)
	}

	if d.ServiceConnection != nil {
		serviceConnection = types.StringValue(*d.ServiceConnection)
	}

	dataMapValue := map[string]attr.Value{
		"primary":            primary,
		"secondary":          secondary,
		"enabled":            enabled,
		"latlng":             latlng,
		"provider":           provider,
		"region":             region,
		"service_connection": serviceConnection,
	}
	data, e := basetypes.NewObjectValue(AutoProvisionValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func tunnelConfigIkeProposalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.TunnelConfigIkeProposal) basetypes.ListValue {
	var dataList []IkeProposalsValue
	for _, d := range l {

		var authAlgo basetypes.StringValue
		var dhGroup basetypes.StringValue
		var encAlgo basetypes.StringValue

		if d.AuthAlgo != nil {
			authAlgo = types.StringValue(string(*d.AuthAlgo))
		}
		if d.DhGroup != nil {
			dhGroup = types.StringValue(string(*d.DhGroup))
		}
		if d.EncAlgo.Value() != nil {
			encAlgo = types.StringValue(string(*d.EncAlgo.Value()))
		}

		dataMapValue := map[string]attr.Value{
			"auth_algo": authAlgo,
			"dh_group":  dhGroup,
			"enc_algo":  encAlgo,
		}
		data, e := NewIkeProposalsValue(IkeProposalsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}

	r, e := types.ListValueFrom(ctx, IkeProposalsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func tunnelConfigIpsecProposalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.TunnelConfigIpsecProposal) basetypes.ListValue {
	var dataList []IpsecProposalsValue
	for _, d := range l {
		var authAlgo basetypes.StringValue
		var dhGroup basetypes.StringValue
		var encAlgo basetypes.StringValue

		if d.AuthAlgo != nil {
			authAlgo = types.StringValue(string(*d.AuthAlgo))
		}
		if d.DhGroup != nil {
			dhGroup = types.StringValue(string(*d.DhGroup))
		}
		if d.EncAlgo.Value() != nil {
			encAlgo = types.StringValue(string(*d.EncAlgo.Value()))
		}

		dataMapValue := map[string]attr.Value{
			"auth_algo": authAlgo,
			"dh_group":  dhGroup,
			"enc_algo":  encAlgo,
		}
		data, e := NewIpsecProposalsValue(IpsecProposalsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}

	r, e := types.ListValueFrom(ctx, IpsecProposalsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func tunnelConfigNodeSdkToTerraform(diags *diag.Diagnostics, d models.TunnelConfigNode, dataMapAttrType map[string]attr.Type) basetypes.ObjectValue {
	var hosts = types.ListNull(types.StringType)
	var internalIps = types.ListNull(types.StringType)
	var probeIps = types.ListNull(types.StringType)
	var remoteIds = types.ListNull(types.StringType)
	var wanNames = types.ListNull(types.StringType)

	if d.Hosts != nil {
		hosts = mistutils.ListOfStringSdkToTerraform(d.Hosts)
	}
	if d.InternalIps != nil {
		internalIps = mistutils.ListOfStringSdkToTerraform(d.InternalIps)
	}
	if d.ProbeIps != nil {
		probeIps = mistutils.ListOfStringSdkToTerraform(d.ProbeIps)
	}
	if d.RemoteIds != nil {
		remoteIds = mistutils.ListOfStringSdkToTerraform(d.RemoteIds)
	}
	if d.WanNames != nil {
		wanNames = mistutils.ListOfStringSdkToTerraform(d.WanNames)
	}

	dataMapValue := map[string]attr.Value{
		"hosts":        hosts,
		"internal_ips": internalIps,
		"probe_ips":    probeIps,
		"remote_ids":   remoteIds,
		"wan_names":    wanNames,
	}
	data, e := basetypes.NewObjectValue(dataMapAttrType, dataMapValue)
	diags.Append(e...)

	return data
}

func tunnelConfigProbeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.TunnelConfigProbe) basetypes.ObjectValue {
	var interval basetypes.Int64Value
	var threshold basetypes.Int64Value
	var timeout basetypes.Int64Value
	var typeProbe = types.StringValue("icmp")

	if d.Interval != nil {
		interval = types.Int64Value(int64(*d.Interval))
	}
	if d.Threshold != nil {
		threshold = types.Int64Value(int64(*d.Threshold))
	}
	if d.Timeout != nil {
		timeout = types.Int64Value(int64(*d.Timeout))
	}
	if d.Type != nil {
		typeProbe = types.StringValue(string(*d.Type))
	}

	dataMapValue := map[string]attr.Value{
		"interval":  interval,
		"threshold": threshold,
		"timeout":   timeout,
		"type":      typeProbe,
	}
	data, e := basetypes.NewObjectValue(ProbeValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func tunnelConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.TunnelConfig) basetypes.MapValue {

	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {
		var autoProvision = types.ObjectNull(AutoProvisionValue{}.AttributeTypes(ctx))
		var ikeLifetime basetypes.Int64Value
		var ikeMode basetypes.StringValue
		var ikeProposals = types.ListNull(IkeProposalsValue{}.Type(ctx))
		var ipsecLifetime basetypes.Int64Value
		var ipsecProposals = types.ListNull(IpsecProposalsValue{}.Type(ctx))
		var localId basetypes.StringValue
		var localSubnets = types.ListNull(types.StringType)
		var mode basetypes.StringValue
		var networks = mistutils.ListOfStringSdkToTerraformEmpty()
		var primary = types.ObjectNull(PrimaryValue{}.AttributeTypes(ctx))
		var probe = types.ObjectNull(ProbeValue{}.AttributeTypes(ctx))
		var protocol basetypes.StringValue
		var provider basetypes.StringValue
		var psk basetypes.StringValue
		var remoteSubnets = types.ListNull(types.StringType)
		var secondary = types.ObjectNull(SecondaryValue{}.AttributeTypes(ctx))
		var version basetypes.StringValue

		if d.AutoProvision != nil {
			autoProvision = tunnelConfigAutoProvSdkToTerraform(ctx, diags, *d.AutoProvision)
		}
		if d.IkeLifetime != nil {
			ikeLifetime = types.Int64Value(int64(*d.IkeLifetime))
		}
		if d.IkeMode != nil {
			ikeMode = types.StringValue(string(*d.IkeMode))
		}
		if d.IkeProposals != nil {
			ikeProposals = tunnelConfigIkeProposalSdkToTerraform(ctx, diags, d.IkeProposals)
		}
		if d.IpsecLifetime != nil {
			ipsecLifetime = types.Int64Value(int64(*d.IpsecLifetime))
		}
		if d.IpsecProposals != nil {
			ipsecProposals = tunnelConfigIpsecProposalSdkToTerraform(ctx, diags, d.IpsecProposals)
		}
		if d.LocalId != nil {
			localId = types.StringValue(*d.LocalId)
		}
		if d.LocalSubnets != nil {
			localSubnets = mistutils.ListOfStringSdkToTerraform(d.LocalSubnets)
		}
		if d.Mode != nil {
			mode = types.StringValue(string(*d.Mode))
		}
		if d.Networks != nil {
			networks = mistutils.ListOfStringSdkToTerraform(d.Networks)
		}
		if d.Primary != nil {
			primary = tunnelConfigNodeSdkToTerraform(diags, *d.Primary, PrimaryValue{}.AttributeTypes(ctx))
		}
		if d.Probe != nil {
			probe = tunnelConfigProbeSdkToTerraform(ctx, diags, *d.Probe)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(string(*d.Protocol))
		}
		if d.Provider != nil {
			provider = types.StringValue(string(*d.Provider))
		}
		if d.Psk != nil {
			psk = types.StringValue(*d.Psk)
		}
		if d.RemoteSubnets != nil {
			remoteSubnets = mistutils.ListOfStringSdkToTerraform(d.RemoteSubnets)
		}
		if d.Secondary != nil {
			secondary = tunnelConfigNodeSdkToTerraform(diags, *d.Secondary, SecondaryValue{}.AttributeTypes(ctx))
		}
		if d.Version != nil {
			version = types.StringValue(string(*d.Version))
		}

		dataMapValue := map[string]attr.Value{
			"auto_provision":  autoProvision,
			"ike_lifetime":    ikeLifetime,
			"ike_mode":        ikeMode,
			"ike_proposals":   ikeProposals,
			"ipsec_lifetime":  ipsecLifetime,
			"ipsec_proposals": ipsecProposals,
			"local_id":        localId,
			"local_subnets":   localSubnets,
			"mode":            mode,
			"networks":        networks,
			"primary":         primary,
			"probe":           probe,
			"protocol":        protocol,
			"provider":        provider,
			"psk":             psk,
			"remote_subnets":  remoteSubnets,
			"secondary":       secondary,
			"version":         version,
		}
		data, e := NewTunnelConfigsValue(TunnelConfigsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMap[k] = data
	}

	stateResult, e := types.MapValueFrom(ctx, TunnelConfigsValue{}.Type(ctx), stateValueMap)
	diags.Append(e...)
	return stateResult
}
