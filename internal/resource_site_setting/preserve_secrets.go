package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// PreserveMxtunnelsRadsecSecrets replaces the radsec object in state with the
// one from prior (plan on Create/Update, prior state on Read).
func PreserveMxtunnelsRadsecSecrets(_ context.Context, _ *diag.Diagnostics, state *SiteSettingModel, prior *SiteSettingModel) {
	if prior == nil || state == nil {
		return
	}
	if prior.Mxtunnels.IsNull() || prior.Mxtunnels.IsUnknown() {
		return
	}
	if state.Mxtunnels.IsNull() || state.Mxtunnels.IsUnknown() {
		return
	}
	if prior.Mxtunnels.Radsec.IsNull() || prior.Mxtunnels.Radsec.IsUnknown() {
		return
	}

	mxt := state.Mxtunnels
	mxt.Radsec = prior.Mxtunnels.Radsec
	state.Mxtunnels = mxt
}
