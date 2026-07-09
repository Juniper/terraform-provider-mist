package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func iotproxyVisionlineSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IotproxyVisionline) basetypes.ObjectValue {
	var accessId basetypes.StringValue
	var cacerts = types.ListNull(types.StringType)
	var enabled basetypes.BoolValue
	var host basetypes.StringValue
	var password basetypes.StringValue
	var port basetypes.Int64Value
	var username basetypes.StringValue

	if d.AccessId != nil {
		accessId = types.StringValue(*d.AccessId)
	}
	if d.Cacerts != nil {
		items := make([]attr.Value, len(d.Cacerts))
		for i, v := range d.Cacerts {
			items[i] = types.StringValue(v)
		}
		list, e := types.ListValue(types.StringType, items)
		diags.Append(e...)
		cacerts = list
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Host != nil {
		host = types.StringValue(*d.Host)
	}
	if d.Password != nil {
		password = types.StringValue(*d.Password)
	}
	if d.Port != nil {
		port = types.Int64Value(int64(*d.Port))
	}
	if d.Username != nil {
		username = types.StringValue(*d.Username)
	}

	dataMapValue := map[string]attr.Value{
		"access_id": accessId,
		"cacerts":   cacerts,
		"enabled":   enabled,
		"host":      host,
		"password":  password,
		"port":      port,
		"username":  username,
	}
	r, e := basetypes.NewObjectValue(VisionlineValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)
	return r
}

func iotproxySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Iotproxy) IotproxyValue {
	var enabled basetypes.BoolValue
	var visionline = types.ObjectNull(VisionlineValue{}.AttributeTypes(ctx))

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Visionline != nil {
		visionline = iotproxyVisionlineSdkToTerraform(ctx, diags, d.Visionline)
	}

	dataMapValue := map[string]attr.Value{
		"enabled":    enabled,
		"visionline": visionline,
	}
	data, e := NewIotproxyValue(IotproxyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
