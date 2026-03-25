package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.Mxedge) (OrgMxedgeModel, diag.Diagnostics) {
	var state OrgMxedgeModel
	var diags diag.Diagnostics

	var id = types.StringNull()
	var mac = types.StringNull()
	var model = types.StringNull()
	var mxagentRegistered = types.BoolNull()
	var mxclusterId = types.StringNull()
	var mxedgeMgmt = NewMxedgeMgmtValueNull()
	var name = types.StringNull()
	var notes = types.StringNull()
	var ntpServers = types.ListNull(types.StringType)
	var oobIpConfig = NewOobIpConfigValueNull()
	var orgId = types.StringNull()
	var proxy = NewProxyValueNull()
	var registrationCode = types.StringNull()
	var services = types.ListNull(types.StringType)
	var siteId = types.StringNull()
	var tuntermDhcpdConfig = types.MapNull(TuntermDhcpdConfigValue{}.Type(ctx))
	var tuntermExtraRoutes = types.MapNull(TuntermExtraRoutesValue{}.Type(ctx))
	var tuntermIgmpSnoopingConfig = NewTuntermIgmpSnoopingConfigValueNull()
	var tuntermIpConfig = NewTuntermIpConfigValueNull()
	var tuntermMonitoring = types.ListNull(types.ListType{ElemType: types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"host":        types.StringType,
			"port":        types.Int64Type,
			"protocol":    types.StringType,
			"src_vlan_id": types.Int64Type,
			"timeout":     types.Int64Type,
		},
	}})
	var tuntermMulticastConfig = NewTuntermMulticastConfigValueNull()
	var tuntermOtherIpConfigs = types.MapNull(TuntermOtherIpConfigsValue{}.Type(ctx))
	var tuntermPortConfig = NewTuntermPortConfigValueNull()
	var tuntermRegistered = types.BoolNull()
	var tuntermSwitchConfig = types.MapNull(TuntermSwitchConfigValue{}.Type(ctx))
	var versions = NewVersionsValueNull()

	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.Mac != nil {
		mac = types.StringValue(*data.Mac)
	}
	if data.Magic != nil {
		registrationCode = types.StringValue(*data.Magic)
	}
	model = types.StringValue(data.Model)
	if data.MxagentRegistered != nil {
		mxagentRegistered = types.BoolValue(*data.MxagentRegistered)
	}
	if data.MxclusterId != nil && data.MxclusterId.String() != "00000000-0000-0000-0000-000000000000" {
		mxclusterId = types.StringValue(data.MxclusterId.String())
	}
	if data.MxedgeMgmt != nil {
		mxedgeMgmt = mxedgeMgmtSdkToTerraform(ctx, &diags, data.MxedgeMgmt)
	}
	name = types.StringValue(data.Name)
	if data.Note != nil {
		notes = types.StringValue(*data.Note)
	}
	if data.NtpServers != nil {
		ntpServers = mistutils.ListOfStringSdkToTerraform(data.NtpServers)
	}
	if data.OobIpConfig != nil {
		oobIpConfig = oobIpConfigSdkToTerraform(ctx, &diags, data.OobIpConfig)
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.Proxy != nil {
		proxy = proxySdkToTerraform(ctx, &diags, data.Proxy)
	}
	if data.Services != nil {
		services = mistutils.ListOfStringSdkToTerraform(data.Services)
	}
	if data.SiteId != nil && data.SiteId.String() != "00000000-0000-0000-0000-000000000000" {
		siteId = types.StringValue(data.SiteId.String())
	}
	if data.TuntermDhcpdConfig != nil {
		tuntermDhcpdConfig = tuntermDhcpdConfigSdkToTerraform(ctx, &diags, data.TuntermDhcpdConfig)
	}
	if data.TuntermExtraRoutes != nil {
		tuntermExtraRoutes = tuntermExtraRoutesSdkToTerraform(ctx, &diags, data.TuntermExtraRoutes)
	}
	if data.TuntermIgmpSnoopingConfig != nil {
		tuntermIgmpSnoopingConfig = tuntermIgmpSnoopingConfigSdkToTerraform(ctx, &diags, data.TuntermIgmpSnoopingConfig)
	}
	if data.TuntermIpConfig != nil {
		tuntermIpConfig = tuntermIpConfigSdkToTerraform(ctx, &diags, data.TuntermIpConfig)
	}
	if data.TuntermMonitoring != nil {
		tuntermMonitoring = tuntermMonitoringSdkToTerraform(ctx, &diags, data.TuntermMonitoring)
	}
	if data.TuntermMulticastConfig != nil {
		tuntermMulticastConfig = tuntermMulticastConfigSdkToTerraform(ctx, &diags, data.TuntermMulticastConfig)
	}
	if data.TuntermOtherIpConfigs != nil {
		tuntermOtherIpConfigs = tuntermOtherIpConfigsSdkToTerraform(ctx, &diags, data.TuntermOtherIpConfigs)
	}
	if data.TuntermPortConfig != nil {
		tuntermPortConfig = tuntermPortConfigSdkToTerraform(ctx, &diags, data.TuntermPortConfig)
	}
	if data.TuntermRegistered != nil {
		tuntermRegistered = types.BoolValue(*data.TuntermRegistered)
	}
	if data.TuntermSwitchConfig != nil {
		tuntermSwitchConfig = tuntermSwitchConfigSdkToTerraform(ctx, &diags, data.TuntermSwitchConfig)
	}
	if data.Versions != nil {
		versions = versionsSdkToTerraform(ctx, &diags, data.Versions)
	}

	state.Id = id
	state.Mac = mac
	state.Model = model
	state.MxagentRegistered = mxagentRegistered
	state.MxclusterId = mxclusterId
	state.MxedgeMgmt = mxedgeMgmt
	state.Name = name
	state.Notes = notes
	state.NtpServers = ntpServers
	state.OobIpConfig = oobIpConfig
	state.OrgId = orgId
	state.Proxy = proxy
	state.RegistrationCode = registrationCode
	state.Services = services
	state.SiteId = siteId
	state.TuntermDhcpdConfig = tuntermDhcpdConfig
	state.TuntermExtraRoutes = tuntermExtraRoutes
	state.TuntermIgmpSnoopingConfig = tuntermIgmpSnoopingConfig
	state.TuntermIpConfig = tuntermIpConfig
	state.TuntermMonitoring = tuntermMonitoring
	state.TuntermMulticastConfig = tuntermMulticastConfig
	state.TuntermOtherIpConfigs = tuntermOtherIpConfigs
	state.TuntermPortConfig = tuntermPortConfig
	state.TuntermRegistered = tuntermRegistered
	state.TuntermSwitchConfig = tuntermSwitchConfig
	state.Versions = versions

	return state, diags
}
