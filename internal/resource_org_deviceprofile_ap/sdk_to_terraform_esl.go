package resource_org_deviceprofile_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func eslSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApEslConfig) EslConfigValue {
	var cacert basetypes.StringValue
	var channel basetypes.Int64Value
	var enabled basetypes.BoolValue
	var host basetypes.StringValue
	var port basetypes.Int64Value
	var typeEsl basetypes.StringValue
	var verifyCert basetypes.BoolValue
	var vlanId basetypes.Int64Value

	if d.Cacert != nil {
		cacert = types.StringValue(*d.Cacert)
	}
	if d.Channel != nil {
		channel = types.Int64Value(int64(*d.Channel))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Host != nil {
		host = types.StringValue(*d.Host)
	}
	if d.Port != nil {
		port = types.Int64Value(int64(*d.Port))
	}
	if d.Type != nil {
		typeEsl = types.StringValue(string(*d.Type))
	}
	if d.VerifyCert != nil {
		verifyCert = types.BoolValue(*d.VerifyCert)
	}
	if d.VlanId != nil {
		vlanId = types.Int64Value(int64(*d.VlanId))
	}

	dataMapValue := map[string]attr.Value{
		"cacert":      cacert,
		"channel":     channel,
		"enabled":     enabled,
		"host":        host,
		"port":        port,
		"type":        typeEsl,
		"verify_cert": verifyCert,
		"vlan_id":     vlanId,
	}
	data, e := NewEslConfigValue(EslConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
