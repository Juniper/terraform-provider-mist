package resource_org_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func installerSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingInstaller) InstallerValue {
	var allow_all_devices basetypes.BoolValue
	var allow_all_sites basetypes.BoolValue
	var extra_site_ids basetypes.ListValue = mist_transform.ListOfUuidSdkToTerraformEmpty(ctx)
	var grace_period basetypes.Int64Value

	if d.AllowAllDevices != nil {
		allow_all_devices = types.BoolValue(*d.AllowAllDevices)
	}
	if d.AllowAllSites != nil {
		allow_all_sites = types.BoolValue(*d.AllowAllSites)
	}
	if d.ExtraSiteIds != nil {
		extra_site_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, d.ExtraSiteIds)
	}
	if d.GracePeriod != nil {
		grace_period = types.Int64Value(int64(*d.GracePeriod))
	}

	data_map_attr_type := InstallerValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_all_devices": allow_all_devices,
		"allow_all_sites":   allow_all_sites,
		"extra_site_ids":    extra_site_ids,
		"grace_period":      grace_period,
	}
	data, e := NewInstallerValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
