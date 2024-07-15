package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func ssrSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingSsr) SsrValue {
	tflog.Debug(ctx, "ssrSdkToTerraform")

	var conductor_hosts basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var disable_stats basetypes.BoolValue

	if d != nil && d.ConductorHosts != nil {
		conductor_hosts = mist_transform.ListOfStringSdkToTerraform(ctx, d.ConductorHosts)
	}
	if d != nil && d.DisableStats != nil {
		disable_stats = types.BoolValue(*d.DisableStats)
	}

	data_map_attr_type := SsrValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"conductor_hosts": conductor_hosts,
		"disable_stats":   disable_stats,
	}
	data, e := NewSsrValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
