package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgMxedgeModel) (*models.Mxedge, diag.Diagnostics) {
	var data models.Mxedge
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	if !plan.ForSite.IsNull() && !plan.ForSite.IsUnknown() {
		data.ForSite = plan.ForSite.ValueBoolPointer()
	} else {
		unset["-for_site"] = ""
	}

	if !plan.Magic.IsNull() && !plan.Magic.IsUnknown() {
		data.Magic = plan.Magic.ValueStringPointer()
	} else {
		unset["-magic"] = ""
	}

	data.Model = plan.Model.ValueString()

	if !plan.MxagentRegistered.IsNull() && !plan.MxagentRegistered.IsUnknown() {
		data.MxagentRegistered = plan.MxagentRegistered.ValueBoolPointer()
	} else {
		unset["-mxagent_registered"] = ""
	}

	if len(plan.MxclusterId.ValueString()) > 0 {
		mxclusterId, e := uuid.Parse(plan.MxclusterId.ValueString())
		if e == nil {
			data.MxclusterId = &mxclusterId
		} else {
			diags.AddError("Bad value for mxcluster_id", e.Error())
		}
	} else {
		unset["-mxcluster_id"] = ""
	}

	if !plan.MxedgeMgmt.IsNull() && !plan.MxedgeMgmt.IsUnknown() {
		data.MxedgeMgmt = mxedgeMgmtTerraformToSdk(ctx, &diags, plan.MxedgeMgmt)
	} else {
		unset["-mxedge_mgmt"] = ""
	}

	data.Name = plan.Name.ValueString()

	if !plan.Note.IsNull() && !plan.Note.IsUnknown() {
		data.Note = plan.Note.ValueStringPointer()
	} else {
		unset["-note"] = ""
	}

	if !plan.NtpServers.IsNull() && !plan.NtpServers.IsUnknown() {
		data.NtpServers = mistutils.ListOfStringTerraformToSdk(plan.NtpServers)
	} else {
		unset["-ntp_servers"] = ""
	}

	if !plan.OobIpConfig.IsNull() && !plan.OobIpConfig.IsUnknown() {
		data.OobIpConfig = oobIpConfigTerraformToSdk(ctx, &diags, plan.OobIpConfig)
	} else {
		unset["-oob_ip_config"] = ""
	}

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

	if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
		data.Services = mistutils.ListOfStringTerraformToSdk(plan.Services)
	} else {
		unset["-services"] = ""
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

	if !plan.TuntermIgmpSnoopingConfig.IsNull() && !plan.TuntermIgmpSnoopingConfig.IsUnknown() {
		data.TuntermIgmpSnoopingConfig = tuntermIgmpSnoopingConfigTerraformToSdk(ctx, &diags, plan.TuntermIgmpSnoopingConfig)
	} else {
		unset["-tunterm_igmp_snooping_config"] = ""
	}

	if !plan.TuntermIpConfig.IsNull() && !plan.TuntermIpConfig.IsUnknown() {
		data.TuntermIpConfig = tuntermIpConfigTerraformToSdk(ctx, &diags, plan.TuntermIpConfig)
	} else {
		unset["-tunterm_ip_config"] = ""
	}

	if !plan.TuntermMonitoring.IsNull() && !plan.TuntermMonitoring.IsUnknown() {
		data.TuntermMonitoring = tuntermMonitoringTerraformToSdk(ctx, &diags, plan.TuntermMonitoring)
	} else {
		unset["-tunterm_monitoring"] = ""
	}

	if !plan.TuntermMulticastConfig.IsNull() && !plan.TuntermMulticastConfig.IsUnknown() {
		data.TuntermMulticastConfig = tuntermMulticastConfigTerraformToSdk(ctx, &diags, plan.TuntermMulticastConfig)
	} else {
		unset["-tunterm_multicast_config"] = ""
	}

	if !plan.TuntermOtherIpConfigs.IsNull() && !plan.TuntermOtherIpConfigs.IsUnknown() {
		data.TuntermOtherIpConfigs = tuntermOtherIpConfigsTerraformToSdk(ctx, &diags, plan.TuntermOtherIpConfigs)
	} else {
		unset["-tunterm_other_ip_configs"] = ""
	}

	if !plan.TuntermPortConfig.IsNull() && !plan.TuntermPortConfig.IsUnknown() {
		data.TuntermPortConfig = tuntermPortConfigTerraformToSdk(ctx, &diags, plan.TuntermPortConfig)
	} else {
		unset["-tunterm_port_config"] = ""
	}

	if !plan.TuntermRegistered.IsNull() && !plan.TuntermRegistered.IsUnknown() {
		data.TuntermRegistered = plan.TuntermRegistered.ValueBoolPointer()
	} else {
		unset["-tunterm_registered"] = ""
	}

	if !plan.TuntermSwitchConfig.IsNull() && !plan.TuntermSwitchConfig.IsUnknown() {
		data.TuntermSwitchConfig = tuntermSwitchConfigTerraformToSdk(ctx, &diags, plan.TuntermSwitchConfig)
	} else {
		unset["-tunterm_switch_config"] = ""
	}

	if !plan.Versions.IsNull() && !plan.Versions.IsUnknown() {
		data.Versions = versionsTerraformToSdk(ctx, &diags, plan.Versions)
	} else {
		unset["-versions"] = ""
	}

	data.AdditionalProperties = unset

	return &data, diags
}
