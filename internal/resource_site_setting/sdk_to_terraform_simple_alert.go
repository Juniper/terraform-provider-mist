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
	var clientCount basetypes.Int64Value
	var duration basetypes.Int64Value
	var incidentCount basetypes.Int64Value

	if d != nil && d.ClientCount != nil {
		clientCount = types.Int64Value(int64(*d.ClientCount))
	}
	if d != nil && d.Duration != nil {
		duration = types.Int64Value(int64(*d.Duration))
	}
	if d != nil && d.IncidentCount != nil {
		incidentCount = types.Int64Value(int64(*d.IncidentCount))
	}

	dataMapValue := map[string]attr.Value{
		"client_count":   clientCount,
		"duration":       duration,
		"incident_count": incidentCount,
	}
	data, e := basetypes.NewObjectValue(ArpFailureValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
func simpleAlertDnsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SimpleAlertDnsFailure) basetypes.ObjectValue {
	var clientCount basetypes.Int64Value
	var duration basetypes.Int64Value
	var incidentCount basetypes.Int64Value

	if d != nil && d.ClientCount != nil {
		clientCount = types.Int64Value(int64(*d.ClientCount))
	}
	if d != nil && d.Duration != nil {
		duration = types.Int64Value(int64(*d.Duration))
	}
	if d != nil && d.IncidentCount != nil {
		incidentCount = types.Int64Value(int64(*d.IncidentCount))
	}

	dataMapValue := map[string]attr.Value{
		"client_count":   clientCount,
		"duration":       duration,
		"incident_count": incidentCount,
	}
	data, e := basetypes.NewObjectValue(DnsFailureValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
func simpleAlertDhcpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SimpleAlertDhcpFailure) basetypes.ObjectValue {

	var clientCount basetypes.Int64Value
	var duration basetypes.Int64Value
	var incidentCount basetypes.Int64Value

	if d != nil && d.ClientCount != nil {
		clientCount = types.Int64Value(int64(*d.ClientCount))
	}
	if d != nil && d.Duration != nil {
		duration = types.Int64Value(int64(*d.Duration))
	}
	if d != nil && d.IncidentCount != nil {
		incidentCount = types.Int64Value(int64(*d.IncidentCount))
	}

	dataMapValue := map[string]attr.Value{
		"client_count":   clientCount,
		"duration":       duration,
		"incident_count": incidentCount,
	}
	data, e := basetypes.NewObjectValue(DhcpFailureValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func simpleAlertSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SimpleAlert) SimpleAlertValue {
	var arpFailure = types.ObjectNull(ArpFailureValue{}.AttributeTypes(ctx))
	var dhcpFailure = types.ObjectNull(DhcpFailureValue{}.AttributeTypes(ctx))
	var dnsFailure = types.ObjectNull(DnsFailureValue{}.AttributeTypes(ctx))

	if d != nil && d.ArpFailure != nil {
		arpFailure = simpleAlertArpSdkToTerraform(ctx, diags, d.ArpFailure)
	}
	if d != nil && d.DhcpFailure != nil {
		dhcpFailure = simpleAlertDhcpSdkToTerraform(ctx, diags, d.DhcpFailure)
	}
	if d != nil && d.DnsFailure != nil {
		dnsFailure = simpleAlertDnsSdkToTerraform(ctx, diags, d.DnsFailure)
	}

	dataMapValue := map[string]attr.Value{
		"arp_failure":  arpFailure,
		"dhcp_failure": dhcpFailure,
		"dns_failure":  dnsFailure,
	}
	data, e := NewSimpleAlertValue(SimpleAlertValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
