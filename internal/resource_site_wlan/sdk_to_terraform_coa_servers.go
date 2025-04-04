package resource_site_wlan

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func coaServersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.CoaServer) basetypes.ListValue {

	var dataList []CoaServersValue
	for _, d := range l {
		var disableEventTimestampCheck basetypes.BoolValue
		var enabled basetypes.BoolValue
		var ip basetypes.StringValue
		var port basetypes.StringValue
		var secret basetypes.StringValue

		if d.DisableEventTimestampCheck != nil {
			disableEventTimestampCheck = types.BoolValue(*d.DisableEventTimestampCheck)
		}
		if d.Enabled != nil {
			enabled = types.BoolValue(*d.Enabled)
		}
		ip = types.StringValue(d.Ip)
		if d.Port != nil {
			port = mistutils.CoaPortAsString(d.Port)
		}
		secret = types.StringValue(d.Secret)

		dataMapValue := map[string]attr.Value{
			"disable_event_timestamp_check": disableEventTimestampCheck,
			"enabled":                       enabled,
			"ip":                            ip,
			"port":                          port,
			"secret":                        secret,
		}
		data, e := NewCoaServersValue(CoaServersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, CoaServersValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r

}
