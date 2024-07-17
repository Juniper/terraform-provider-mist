package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func simpleAlertArpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SimpleAlertArpFailure) basetypes.ObjectValue {
	var client_count basetypes.Int64Value
	var duration basetypes.Int64Value
	var incident_count basetypes.Int64Value

	if d != nil && d.ClientCount != nil {
		client_count = types.Int64Value(int64(*d.ClientCount))
	}
	if d != nil && d.Duration != nil {
		duration = types.Int64Value(int64(*d.Duration))
	}
	if d != nil && d.IncidentCount != nil {
		incident_count = types.Int64Value(int64(*d.IncidentCount))
	}

	data_map_attr_type := ArpFailureValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"client_count":   client_count,
		"duration":       duration,
		"incident_count": incident_count,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
func simpleAlertDnsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SimpleAlertDnsFailure) basetypes.ObjectValue {
	var client_count basetypes.Int64Value
	var duration basetypes.Int64Value
	var incident_count basetypes.Int64Value

	if d != nil && d.ClientCount != nil {
		client_count = types.Int64Value(int64(*d.ClientCount))
	}
	if d != nil && d.Duration != nil {
		duration = types.Int64Value(int64(*d.Duration))
	}
	if d != nil && d.IncidentCount != nil {
		incident_count = types.Int64Value(int64(*d.IncidentCount))
	}

	data_map_attr_type := DnsFailureValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"client_count":   client_count,
		"duration":       duration,
		"incident_count": incident_count,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
func simpleAlertDhcpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SimpleAlertDhcpFailure) basetypes.ObjectValue {

	var client_count basetypes.Int64Value
	var duration basetypes.Int64Value
	var incident_count basetypes.Int64Value

	if d != nil && d.ClientCount != nil {
		client_count = types.Int64Value(int64(*d.ClientCount))
	}
	if d != nil && d.Duration != nil {
		duration = types.Int64Value(int64(*d.Duration))
	}
	if d != nil && d.IncidentCount != nil {
		incident_count = types.Int64Value(int64(*d.IncidentCount))
	}

	data_map_attr_type := DhcpFailureValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"client_count":   client_count,
		"duration":       duration,
		"incident_count": incident_count,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func simpleAlertSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SimpleAlert) SimpleAlertValue {
	var arp_failure basetypes.ObjectValue = types.ObjectNull(ArpFailureValue{}.AttributeTypes(ctx))
	var dhcp_failure basetypes.ObjectValue = types.ObjectNull(DhcpFailureValue{}.AttributeTypes(ctx))
	var dns_failure basetypes.ObjectValue = types.ObjectNull(DnsFailureValue{}.AttributeTypes(ctx))

	if d != nil && d.ArpFailure != nil {
		arp_failure = simpleAlertArpSdkToTerraform(ctx, diags, d.ArpFailure)
	}
	if d != nil && d.DhcpFailure != nil {
		dhcp_failure = simpleAlertDhcpSdkToTerraform(ctx, diags, d.DhcpFailure)
	}
	if d != nil && d.DnsFailure != nil {
		dns_failure = simpleAlertDnsSdkToTerraform(ctx, diags, d.DnsFailure)
	}

	data_map_attr_type := SimpleAlertValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"arp_failure":  arp_failure,
		"dhcp_failure": dhcp_failure,
		"dns_failure":  dns_failure,
	}
	data, e := NewSimpleAlertValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
