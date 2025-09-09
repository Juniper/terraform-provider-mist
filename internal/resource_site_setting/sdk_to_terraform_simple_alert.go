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
	if d == nil {
		return types.ObjectNull(ArpFailureValue{}.AttributeTypes(ctx))
	}

	// If all fields are nil, return null object
	if d.ClientCount == nil && d.Duration == nil && d.IncidentCount == nil {
		return types.ObjectNull(ArpFailureValue{}.AttributeTypes(ctx))
	}

	var clientCount = types.Int64Null()
	var duration = types.Int64Null()
	var incidentCount = types.Int64Null()

	if d.ClientCount != nil {
		clientCount = types.Int64Value(int64(*d.ClientCount))
	}
	if d.Duration != nil {
		duration = types.Int64Value(int64(*d.Duration))
	}
	if d.IncidentCount != nil {
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
	if d == nil {
		return types.ObjectNull(DnsFailureValue{}.AttributeTypes(ctx))
	}

	// If all fields are nil, return null object
	if d.ClientCount == nil && d.Duration == nil && d.IncidentCount == nil {
		return types.ObjectNull(DnsFailureValue{}.AttributeTypes(ctx))
	}

	var clientCount = types.Int64Null()
	var duration = types.Int64Null()
	var incidentCount = types.Int64Null()

	if d.ClientCount != nil {
		clientCount = types.Int64Value(int64(*d.ClientCount))
	}
	if d.Duration != nil {
		duration = types.Int64Value(int64(*d.Duration))
	}
	if d.IncidentCount != nil {
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
	if d == nil {
		return types.ObjectNull(DhcpFailureValue{}.AttributeTypes(ctx))
	}

	// If all fields are nil, return null object
	if d.ClientCount == nil && d.Duration == nil && d.IncidentCount == nil {
		return types.ObjectNull(DhcpFailureValue{}.AttributeTypes(ctx))
	}

	var clientCount = types.Int64Null()
	var duration = types.Int64Null()
	var incidentCount = types.Int64Null()

	if d.ClientCount != nil {
		clientCount = types.Int64Value(int64(*d.ClientCount))
	}
	if d.Duration != nil {
		duration = types.Int64Value(int64(*d.Duration))
	}
	if d.IncidentCount != nil {
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
	if d == nil {
		return NewSimpleAlertValueNull()
	}

	var arpFailure = types.ObjectNull(ArpFailureValue{}.AttributeTypes(ctx))
	var dhcpFailure = types.ObjectNull(DhcpFailureValue{}.AttributeTypes(ctx))
	var dnsFailure = types.ObjectNull(DnsFailureValue{}.AttributeTypes(ctx))

	if d.ArpFailure != nil {
		arpFailure = simpleAlertArpSdkToTerraform(ctx, diags, d.ArpFailure)
	}
	if d.DhcpFailure != nil {
		dhcpFailure = simpleAlertDhcpSdkToTerraform(ctx, diags, d.DhcpFailure)
	}
	if d.DnsFailure != nil {
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
