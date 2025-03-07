package resource_site_setting

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func zoneOccupancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteZoneOccupancyAlert) ZoneOccupancyAlertValue {
	var emailNotifiers = mistutils.ListOfStringSdkToTerraformEmpty()
	var enabled basetypes.BoolValue
	var threshold basetypes.Int64Value

	if d.EmailNotifiers != nil {
		emailNotifiers = mistutils.ListOfStringSdkToTerraform(d.EmailNotifiers)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Threshold != nil {
		threshold = types.Int64Value(int64(*d.Threshold))
	}

	dataMapValue := map[string]attr.Value{
		"email_notifiers": emailNotifiers,
		"enabled":         enabled,
		"threshold":       threshold,
	}
	data, e := NewZoneOccupancyAlertValue(ZoneOccupancyAlertValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
