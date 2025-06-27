package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func ssrSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SettingSsr) SsrValue {

	var conductorHosts = mistutils.ListOfStringSdkToTerraformEmpty()
	var disableStats basetypes.BoolValue

	if d != nil && d.ConductorHosts != nil {
		conductorHosts = mistutils.ListOfStringSdkToTerraform(d.ConductorHosts)
	}
	if d != nil && d.DisableStats != nil {
		disableStats = types.BoolValue(*d.DisableStats)
	}

	dataMapValue := map[string]attr.Value{
		"conductor_hosts": conductorHosts,
		"disable_stats":   disableStats,
	}
	data, e := NewSsrValue(SsrValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
