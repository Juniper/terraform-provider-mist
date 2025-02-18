package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func airwatchSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanAirwatch) AirwatchValue {
	var apiKey basetypes.StringValue
	var consoleUrl basetypes.StringValue
	var enabled basetypes.BoolValue
	var password basetypes.StringValue
	var username basetypes.StringValue

	if d != nil && d.ApiKey != nil {
		apiKey = types.StringValue(*d.ApiKey)
	}
	if d != nil && d.ConsoleUrl != nil {
		consoleUrl = types.StringValue(*d.ConsoleUrl)
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

	dataMapValue := map[string]attr.Value{
		"api_key":     apiKey,
		"console_url": consoleUrl,
		"enabled":     enabled,
		"password":    password,
		"username":    username,
	}
	data, e := NewAirwatchValue(AirwatchValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
