package resource_org_mxcluster

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, plan *OrgMxclusterModel) (*models.Mxcluster, diag.Diagnostics) {
	var data models.Mxcluster
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	if !plan.MistDas.IsNull() && !plan.MistDas.IsUnknown() {
		data.MistDas = mistDasTerraformToSdk(ctx, &diags, plan.MistDas)
	} else {
		unset["-mist_das"] = ""
	}

	if !plan.MistNac.IsNull() && !plan.MistNac.IsUnknown() {
		data.MistNac = mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
	} else {
		unset["-mist_nac"] = ""
	}

	if !plan.MxedgeMgmt.IsNull() && !plan.MxedgeMgmt.IsUnknown() {
		data.MxedgeMgmt = mxedgeMgmtTerraformToSdk(ctx, &diags, plan.MxedgeMgmt)
	} else {
		unset["-mxedge_mgmt"] = ""
	}

	data.Name = plan.Name.ValueStringPointer()

	// org_id is required, so we always parse it
	orgId, e := uuid.Parse(plan.OrgId.ValueString())
	if e == nil {
		data.OrgId = &orgId
	} else {
		diags.AddError("Invalid value for org_id", e.Error())
	}

	if !plan.Proxy.IsNull() && !plan.Proxy.IsUnknown() {
		data.Proxy = proxyTerraformToSdk(ctx, &diags, plan.Proxy)
	} else {
		unset["-proxy"] = ""
	}

	if !plan.Radsec.IsNull() && !plan.Radsec.IsUnknown() {
		data.Radsec = radsecTerraformToSdk(ctx, &diags, plan.Radsec)
	} else {
		unset["-radsec"] = ""
	}

	if !plan.RadsecTls.IsNull() && !plan.RadsecTls.IsUnknown() {
		data.RadsecTls = radsecTlsTerraformToSdk(ctx, &diags, plan.RadsecTls)
	} else {
		unset["-radsec_tls"] = ""
	}

	if len(plan.SiteId.ValueString()) > 0 {
		siteId, e := uuid.Parse(plan.SiteId.ValueString())
		if e == nil {
			data.SiteId = &siteId
		} else {
			diags.AddError("Bad value for site_id", e.Error())
		}
	} else {
		unset["-site_id"] = ""
	}

	if !plan.TuntermApSubnets.IsNull() && !plan.TuntermApSubnets.IsUnknown() {
		data.TuntermApSubnets = mistutils.ListOfStringTerraformToSdk(plan.TuntermApSubnets)
	} else {
		unset["-tunterm_ap_subnets"] = ""
	}

	if !plan.TuntermDhcpdConfig.IsNull() && !plan.TuntermDhcpdConfig.IsUnknown() {
		data.TuntermDhcpdConfig = tuntermDhcpdConfigTerraformToSdk(ctx, &diags, plan.TuntermDhcpdConfig)
	} else {
		unset["-tunterm_dhcpd_config"] = ""
	}

	if !plan.TuntermExtraRoutes.IsNull() && !plan.TuntermExtraRoutes.IsUnknown() {
		data.TuntermExtraRoutes = tuntermExtraRoutesTerraformToSdk(ctx, &diags, plan.TuntermExtraRoutes)
	} else {
		unset["-tunterm_extra_routes"] = ""
	}

	if !plan.TuntermHosts.IsNull() && !plan.TuntermHosts.IsUnknown() {
		data.TuntermHosts = mistutils.ListOfStringTerraformToSdk(plan.TuntermHosts)
	} else {
		unset["-tunterm_hosts"] = ""
	}

	if !plan.TuntermHostsOrder.IsNull() && !plan.TuntermHostsOrder.IsUnknown() {
		data.TuntermHostsOrder = listOfInt64TerraformToSdk(plan.TuntermHostsOrder)
	} else {
		unset["-tunterm_hosts_order"] = ""
	}

	if !plan.TuntermHostsSelection.IsNull() && !plan.TuntermHostsSelection.IsUnknown() {
		data.TuntermHostsSelection = (*models.MxclusterTuntermHostsSelectionEnum)(plan.TuntermHostsSelection.ValueStringPointer())
	} else {
		unset["-tunterm_hosts_selection"] = ""
	}

	if !plan.TuntermMonitoring.IsNull() && !plan.TuntermMonitoring.IsUnknown() {
		data.TuntermMonitoring = tuntermMonitoringTerraformToSdk(ctx, &diags, plan.TuntermMonitoring)
	} else {
		unset["-tunterm_monitoring"] = ""
	}

	if !plan.TuntermMonitoringDisabled.IsNull() && !plan.TuntermMonitoringDisabled.IsUnknown() {
		data.TuntermMonitoringDisabled = plan.TuntermMonitoringDisabled.ValueBoolPointer()
	} else {
		unset["-tunterm_monitoring_disabled"] = ""
	}

	data.AdditionalProperties = unset

	return &data, diags
}

func listOfInt64TerraformToSdk(list basetypes.ListValue) []int {
	var data []int
	for _, item := range list.Elements() {
		if v, ok := item.(basetypes.Int64Value); ok && !v.IsNull() && !v.IsUnknown() {
			data = append(data, int(v.ValueInt64()))
		}
	}
	return data
}
