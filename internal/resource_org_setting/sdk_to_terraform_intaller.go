package resource_org_setting

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func installerSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingInstaller) InstallerValue {
	var allowAllDevices basetypes.BoolValue
	var allowAllSites basetypes.BoolValue
	var extraSiteIds = mistutils.ListOfUuidSdkToTerraformEmpty()
	var gracePeriod basetypes.Int64Value

	if d.AllowAllDevices != nil {
		allowAllDevices = types.BoolValue(*d.AllowAllDevices)
	}
	if d.AllowAllSites != nil {
		allowAllSites = types.BoolValue(*d.AllowAllSites)
	}
	if d.ExtraSiteIds != nil {
		extraSiteIds = mistutils.ListOfUuidSdkToTerraform(d.ExtraSiteIds)
	}
	if d.GracePeriod != nil {
		gracePeriod = types.Int64Value(int64(*d.GracePeriod))
	}

	dataMapValue := map[string]attr.Value{
		"allow_all_devices": allowAllDevices,
		"allow_all_sites":   allowAllSites,
		"extra_site_ids":    extraSiteIds,
		"grace_period":      gracePeriod,
	}
	data, e := NewInstallerValue(InstallerValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
