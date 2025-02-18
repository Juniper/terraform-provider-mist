package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func simpleAlertArpTerraformToSdk(ctx context.Context, o basetypes.ObjectValue) *models.SimpleAlertArpFailure {
	data := models.SimpleAlertArpFailure{}
	if o.IsNull() || o.IsUnknown() {
		return &data
	} else {
		d := NewArpFailureValueMust(o.AttributeTypes(ctx), o.Attributes())

		data.ClientCount = models.ToPointer(int(d.ClientCount.ValueInt64()))
		data.Duration = models.ToPointer(int(d.Duration.ValueInt64()))
		data.IncidentCount = models.ToPointer(int(d.IncidentCount.ValueInt64()))

		return &data
	}
}

func simpleAlertDhcpTerraformToSdk(ctx context.Context, o basetypes.ObjectValue) *models.SimpleAlertDhcpFailure {
	data := models.SimpleAlertDhcpFailure{}
	if o.IsNull() || o.IsUnknown() {
		return &data
	} else {
		d := NewDhcpFailureValueMust(o.AttributeTypes(ctx), o.Attributes())

		data.ClientCount = models.ToPointer(int(d.ClientCount.ValueInt64()))
		data.Duration = models.ToPointer(int(d.Duration.ValueInt64()))
		data.IncidentCount = models.ToPointer(int(d.IncidentCount.ValueInt64()))

		return &data
	}
}

func simpleAlertDnsTerraformToSdk(ctx context.Context, o basetypes.ObjectValue) *models.SimpleAlertDnsFailure {
	data := models.SimpleAlertDnsFailure{}
	if o.IsNull() || o.IsUnknown() {
		return &data
	} else {
		d := NewDnsFailureValueMust(o.AttributeTypes(ctx), o.Attributes())

		data.ClientCount = models.ToPointer(int(d.ClientCount.ValueInt64()))
		data.Duration = models.ToPointer(int(d.Duration.ValueInt64()))
		data.IncidentCount = models.ToPointer(int(d.IncidentCount.ValueInt64()))

		return &data
	}
}

func simpleAlertTerraformToSdk(ctx context.Context, d SimpleAlertValue) *models.SimpleAlert {
	data := models.SimpleAlert{}

	arp := simpleAlertArpTerraformToSdk(ctx, d.ArpFailure)
	data.ArpFailure = arp

	dhcp := simpleAlertDhcpTerraformToSdk(ctx, d.DhcpFailure)
	data.DhcpFailure = dhcp

	dns := simpleAlertDnsTerraformToSdk(ctx, d.DnsFailure)
	data.DnsFailure = dns

	return &data
}
