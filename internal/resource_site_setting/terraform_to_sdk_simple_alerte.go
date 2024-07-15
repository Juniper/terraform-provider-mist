package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func simpleAlertArpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, o basetypes.ObjectValue) *models.SimpleAlertArpFailure {
	tflog.Debug(ctx, "simpleAlertArpTerraformToSdk")
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

func simpleAlertDhcpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, o basetypes.ObjectValue) *models.SimpleAlertDhcpFailure {
	tflog.Debug(ctx, "simpleAlertDhcpTerraformToSdk")
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

func simpleAlertDnsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, o basetypes.ObjectValue) *models.SimpleAlertDnsFailure {
	tflog.Debug(ctx, "simpleAlertDnsTerraformToSdk")
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

func simpleAlertTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SimpleAlertValue) *models.SimpleAlert {
	tflog.Debug(ctx, "simpleAlertTerraformToSdk")
	data := models.SimpleAlert{}

	arp := simpleAlertArpTerraformToSdk(ctx, diags, d.ArpFailure)
	data.ArpFailure = arp

	dhcp := simpleAlertDhcpTerraformToSdk(ctx, diags, d.DhcpFailure)
	data.DhcpFailure = dhcp

	dns := simpleAlertDnsTerraformToSdk(ctx, diags, d.DnsFailure)
	data.DnsFailure = dns

	return &data
}
