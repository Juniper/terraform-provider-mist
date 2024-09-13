package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func tunnelConfigAutoProvNodeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.TunnelConfigsAutoProvisionNode, r_type map[string]attr.Type) basetypes.ObjectValue {

	var num_hosts basetypes.StringValue
	var wan_names basetypes.ListValue = types.ListNull(types.StringType)

	if d.NumHosts != nil {
		num_hosts = types.StringValue(*d.NumHosts)
	}
	if d.WanNames != nil {
		wan_names = mist_transform.ListOfStringSdkToTerraform(ctx, d.WanNames)
	}

	r_attr_value := map[string]attr.Value{
		"num_hosts": num_hosts,
		"wan_names": wan_names,
	}
	r, e := basetypes.NewObjectValue(r_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelConfigAutoProvSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.TunnelConfigsAutoProvision) basetypes.ObjectValue {

	var primary basetypes.ObjectValue = types.ObjectNull(PrimaryValue{}.AttributeTypes(ctx))
	var secondary basetypes.ObjectValue = types.ObjectNull(SecondaryValue{}.AttributeTypes(ctx))
	var enable basetypes.BoolValue
	var latlng basetypes.ObjectValue = types.ObjectNull(LatlngValue{}.AttributeTypes(ctx))
	var region basetypes.StringValue = types.StringValue("auto")

	if d.Primary != nil {
		primary = tunnelConfigAutoProvNodeSdkToTerraform(ctx, diags, *d.Primary, PrimaryValue{}.AttributeTypes(ctx))
	}
	if d.Secondary != nil {
		secondary = tunnelConfigAutoProvNodeSdkToTerraform(ctx, diags, *d.Secondary, SecondaryValue{}.AttributeTypes(ctx))
	}
	if d.Enable != nil {
		enable = types.BoolValue(*d.Enable)
	}
	if d.Latlng != nil {
		latlng_type := map[string]attr.Type{
			"lat": basetypes.Float64Type{},
			"lng": basetypes.Float64Type{},
		}
		latlng_value := map[string]attr.Value{
			"lat": types.Float64Value(d.Latlng.Lat),
			"lng": types.Float64Value(d.Latlng.Lng),
		}
		tmp, e := NewLatlngValue(latlng_type, latlng_value)
		diags.Append(e...)
		latlng, _ = tmp.ToObjectValue(ctx)

	}

	data_map_attr_type := AutoProvisionValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"primary":   primary,
		"secondary": secondary,
		"enable":    enable,
		"latlng":    latlng,
		"region":    region,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func tunnelConfigIkeProposalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.GatewayTemplateTunnelIkeProposal) basetypes.ListValue {
	var data_list = []IkeProposalsValue{}
	for _, d := range l {

		var auth_algo basetypes.StringValue
		var dh_group basetypes.StringValue = types.StringValue("14")
		var enc_algo basetypes.StringValue = types.StringValue("aes256")

		if d.AuthAlgo != nil {
			auth_algo = types.StringValue(string(*d.AuthAlgo))
		}
		if d.DhGroup != nil {
			dh_group = types.StringValue(string(*d.DhGroup))
		}
		if d.EncAlgo.Value() != nil {
			enc_algo = types.StringValue(string(*d.EncAlgo.Value()))
		}

		data_map_attr_type := IkeProposalsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"auth_algo": auth_algo,
			"dh_group":  dh_group,
			"enc_algo":  enc_algo,
		}
		data, e := NewIkeProposalsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := IkeProposalsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func tunnelConfigIpsecProposalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.GatewayTemplateTunnelIpsecProposal) basetypes.ListValue {
	var data_list = []IpsecProposalsValue{}
	for _, d := range l {
		var auth_algo basetypes.StringValue
		var dh_group basetypes.StringValue = types.StringValue("14")
		var enc_algo basetypes.StringValue = types.StringValue("aes256")

		if d.AuthAlgo != nil {
			auth_algo = types.StringValue(string(*d.AuthAlgo))
		}
		if d.DhGroup != nil {
			dh_group = types.StringValue(string(*d.DhGroup))
		}
		if d.EncAlgo.Value() != nil {
			enc_algo = types.StringValue(string(*d.EncAlgo.Value()))
		}

		data_map_attr_type := IpsecProposalsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"auth_algo": auth_algo,
			"dh_group":  dh_group,
			"enc_algo":  enc_algo,
		}
		data, e := NewIpsecProposalsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := IpsecProposalsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func tunnelConfigNodeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.GatewayTemplateTunnelNode, data_map_attr_type map[string]attr.Type) basetypes.ObjectValue {
	var hosts basetypes.ListValue = types.ListNull(types.StringType)
	var internal_ips basetypes.ListValue = types.ListNull(types.StringType)
	var probe_ips basetypes.ListValue = types.ListNull(types.StringType)
	var remote_ids basetypes.ListValue = types.ListNull(types.StringType)
	var wan_names basetypes.ListValue = types.ListNull(types.StringType)

	if d.Hosts != nil {
		hosts = mist_transform.ListOfStringSdkToTerraform(ctx, d.Hosts)
	}
	if d.InternalIps != nil {
		internal_ips = mist_transform.ListOfStringSdkToTerraform(ctx, d.InternalIps)
	}
	if d.ProbeIps != nil {
		probe_ips = mist_transform.ListOfStringSdkToTerraform(ctx, d.ProbeIps)
	}
	if d.RemoteIds != nil {
		remote_ids = mist_transform.ListOfStringSdkToTerraform(ctx, d.RemoteIds)
	}
	if d.WanNames != nil {
		wan_names = mist_transform.ListOfStringSdkToTerraform(ctx, d.WanNames)
	}

	data_map_value := map[string]attr.Value{
		"hosts":        hosts,
		"internal_ips": internal_ips,
		"probe_ips":    probe_ips,
		"remote_ids":   remote_ids,
		"wan_names":    wan_names,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func tunnelConfigProbeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.GatewayTemplateTunnelProbe) basetypes.ObjectValue {
	var interval basetypes.Int64Value
	var threshold basetypes.Int64Value
	var timeout basetypes.Int64Value
	var type_probe basetypes.StringValue = types.StringValue("icmp")

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
		type_probe = types.StringValue(string(*d.Type))
	}

	data_map_attr_type := ProbeValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"interval":  interval,
		"threshold": threshold,
		"timeout":   timeout,
		"type":      type_probe,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func tunnelConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.TunnelConfigs) basetypes.MapValue {

	state_value_map := make(map[string]attr.Value)
	for k, d := range m {
		var auto_provision basetypes.ObjectValue = types.ObjectNull(AutoProvisionValue{}.AttributeTypes(ctx))
		var ike_lifetime basetypes.Int64Value
		var ike_mode basetypes.StringValue = types.StringValue("main")
		var ike_proposals basetypes.ListValue = types.ListNull(IkeProposalsValue{}.Type(ctx))
		var ipsec_lifetime basetypes.Int64Value
		var ipsec_proposals basetypes.ListValue = types.ListNull(IpsecProposalsValue{}.Type(ctx))
		var local_id basetypes.StringValue
		var mode basetypes.StringValue = types.StringValue("active-standby")
		var primary basetypes.ObjectValue = types.ObjectNull(PrimaryValue{}.AttributeTypes(ctx))
		var probe basetypes.ObjectValue = types.ObjectNull(ProbeValue{}.AttributeTypes(ctx))
		var protocol basetypes.StringValue
		var provider basetypes.StringValue
		var psk basetypes.StringValue
		var secondary basetypes.ObjectValue = types.ObjectNull(SecondaryValue{}.AttributeTypes(ctx))
		var version basetypes.StringValue = types.StringValue("2")

		if d.AutoProvision != nil {
			auto_provision = tunnelConfigAutoProvSdkToTerraform(ctx, diags, *d.AutoProvision)
		}
		if d.IkeLifetime != nil {
			ike_lifetime = types.Int64Value(int64(*d.IkeLifetime))
		}
		if d.IkeMode != nil {
			ike_mode = types.StringValue(string(*d.IkeMode))
		}
		if d.IkeProposals != nil {
			ike_proposals = tunnelConfigIkeProposalSdkToTerraform(ctx, diags, d.IkeProposals)
		}
		if d.IpsecLifetime != nil {
			ipsec_lifetime = types.Int64Value(int64(*d.IpsecLifetime))
		}
		if d.IpsecProposals != nil {
			ipsec_proposals = tunnelConfigIpsecProposalSdkToTerraform(ctx, diags, d.IpsecProposals)
		}
		if d.LocalId != nil {
			local_id = types.StringValue(*d.LocalId)
		}
		if d.Mode != nil {
			mode = types.StringValue(string(*d.Mode))
		}
		if d.Primary != nil {
			primary = tunnelConfigNodeSdkToTerraform(ctx, diags, *d.Primary, PrimaryValue{}.AttributeTypes(ctx))
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
		if d.Secondary != nil {
			secondary = tunnelConfigNodeSdkToTerraform(ctx, diags, *d.Secondary, SecondaryValue{}.AttributeTypes(ctx))
		}
		if d.Version != nil {
			version = types.StringValue(string(*d.Version))
		}

		data_map_attr_type := TunnelConfigsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"auto_provision":  auto_provision,
			"ike_lifetime":    ike_lifetime,
			"ike_mode":        ike_mode,
			"ike_proposals":   ike_proposals,
			"ipsec_lifetime":  ipsec_lifetime,
			"ipsec_proposals": ipsec_proposals,
			"local_id":        local_id,
			"mode":            mode,
			"primary":         primary,
			"probe":           probe,
			"protocol":        protocol,
			"provider":        provider,
			"psk":             psk,
			"secondary":       secondary,
			"version":         version,
		}
		data, e := NewTunnelConfigsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}
	state_type := TunnelConfigsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
