package resource_site_setting

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func rogueTerraformToSdk(d RogueValue) *models.SiteRogue {
	data := models.SiteRogue{}

	if !d.AllowedVlanIds.IsNull() && !d.AllowedVlanIds.IsUnknown() {
		data.AllowedVlanIds = mistutils.ListOfIntTerraformToSdk(d.AllowedVlanIds)
	}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.HoneypotEnabled.IsNull() && !d.HoneypotEnabled.IsUnknown() {
		data.HoneypotEnabled = d.HoneypotEnabled.ValueBoolPointer()
	}
	if !d.MinDuration.IsNull() && !d.MinDuration.IsUnknown() {
		data.MinDuration = models.ToPointer(int(d.MinDuration.ValueInt64()))
	}
	if !d.MinRogueDuration.IsNull() && !d.MinRogueDuration.IsUnknown() {
		data.MinRogueDuration = models.ToPointer(int(d.MinRogueDuration.ValueInt64()))
	}
	if !d.MinRssi.IsNull() && !d.MinRssi.IsUnknown() {
		data.MinRssi = models.ToPointer(int(d.MinRssi.ValueInt64()))
	}
	if !d.MinRogueRssi.IsNull() && !d.MinRogueRssi.IsUnknown() {
		data.MinRogueRssi = models.ToPointer(int(d.MinRogueRssi.ValueInt64()))
	}
	if !d.WhitelistedBssids.IsNull() && !d.WhitelistedBssids.IsUnknown() {
		data.WhitelistedBssids = mistutils.ListOfStringTerraformToSdk(d.WhitelistedBssids)
	}
	if !d.WhitelistedSsids.IsNull() && !d.WhitelistedSsids.IsUnknown() {
		data.WhitelistedSsids = mistutils.ListOfStringTerraformToSdk(d.WhitelistedSsids)
	}
	return &data
}
