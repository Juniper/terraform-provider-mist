package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func airwatchSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanAirwatch) AirwatchValue {
	var api_key basetypes.StringValue
	var console_url basetypes.StringValue
	var enabled basetypes.BoolValue
	var password basetypes.StringValue
	var username basetypes.StringValue

	if d != nil && d.ApiKey != nil {
		api_key = types.StringValue(*d.ApiKey)
	}
	if d != nil && d.ConsoleUrl != nil {
		console_url = types.StringValue(*d.ConsoleUrl)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Password != nil {
		password = types.StringValue(*d.Password)
	}
	if d != nil && d.Username != nil {
		username = types.StringValue(*d.Username)
	}

	data_map_attr_type := AirwatchValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"api_key":     api_key,
		"console_url": console_url,
		"enabled":     enabled,
		"password":    password,
		"username":    username,
	}
	data, e := NewAirwatchValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
