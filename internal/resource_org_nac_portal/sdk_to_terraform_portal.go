package resource_org_nac_portal

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func portalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NacPortalGuestPortal) PortalValue {
	var auth basetypes.StringValue
	var expire basetypes.Int64Value
	var externalPortalUrl basetypes.StringValue
	var forceReconnect basetypes.BoolValue
	var forward basetypes.BoolValue
	var forwardUrl basetypes.StringValue
	var maxNumDevices basetypes.Int64Value
	var privacy basetypes.BoolValue

	if d != nil && d.Auth != nil {
		auth = types.StringValue(string(*d.Auth))
	}
	if d != nil && d.Expire != nil {
		expire = types.Int64Value(int64(*d.Expire))
	}
	if d != nil && d.ExternalPortalUrl != nil {
		externalPortalUrl = types.StringValue(*d.ExternalPortalUrl)
	}
	if d != nil && d.ForceReconnect != nil {
		forceReconnect = types.BoolValue(*d.ForceReconnect)
	}
	if d != nil && d.Forward != nil {
		forward = types.BoolValue(*d.Forward)
	}
	if d != nil && d.ForwardUrl != nil {
		forwardUrl = types.StringValue(*d.ForwardUrl)
	}
	if d != nil && d.MaxNumDevices != nil {
		maxNumDevices = types.Int64Value(int64(*d.MaxNumDevices))
	}
	if d != nil && d.Privacy != nil {
		privacy = types.BoolValue(*d.Privacy)
	}

	dataMapValue := map[string]attr.Value{
		"auth":                auth,
		"expire":              expire,
		"external_portal_url": externalPortalUrl,
		"force_reconnect":     forceReconnect,
		"forward":             forward,
		"forward_url":         forwardUrl,
		"max_num_devices":     maxNumDevices,
		"privacy":             privacy,
	}
	data, e := NewPortalValue(PortalValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
