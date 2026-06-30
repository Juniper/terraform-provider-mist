package resource_org_mxcluster

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func mistNacedgeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacedgeValue) *models.MxclusterNacedge {
	data := models.MxclusterNacedge{}

	if !d.AuthTtl.IsNull() && !d.AuthTtl.IsUnknown() {
		data.AuthTtl = models.ToPointer(int(d.AuthTtl.ValueInt64()))
	}

	if !d.CachingSiteIds.IsNull() && !d.CachingSiteIds.IsUnknown() {
		strs := mistutils.ListOfStringTerraformToSdk(d.CachingSiteIds)
		var uuids []uuid.UUID
		for _, s := range strs {
			if id, e := uuid.Parse(s); e == nil {
				uuids = append(uuids, id)
			} else {
				diags.AddError("Invalid UUID in caching_site_ids", e.Error())
			}
		}
		data.CachingSiteIds = uuids
	}

	if !d.DefaultDot1xVlan.IsNull() && !d.DefaultDot1xVlan.IsUnknown() {
		data.DefaultDot1xVlan = d.DefaultDot1xVlan.ValueStringPointer()
	}

	if !d.DefaultVlan.IsNull() && !d.DefaultVlan.IsUnknown() {
		data.DefaultVlan = d.DefaultVlan.ValueStringPointer()
	}

	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if !d.NacEdgeHosts.IsNull() && !d.NacEdgeHosts.IsUnknown() {
		data.NacEdgeHosts = mistutils.ListOfStringTerraformToSdk(d.NacEdgeHosts)
	}

	return &data
}
