package resource_org_mxcluster

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.Mxcluster) (OrgMxclusterModel, diag.Diagnostics) {
	var state OrgMxclusterModel
	var diags diag.Diagnostics

	var id = types.StringNull()
	var mistDas = NewMistDasValueNull()
	var mistNac = NewMistNacValueNull()
	var mxedgeMgmt = NewMxedgeMgmtValueNull()
	var name = types.StringNull()
	var orgId = types.StringNull()
	var proxy = NewProxyValueNull()
	var radsec = NewRadsecValueNull()
	var radsecTls = NewRadsecTlsValueNull()
	var siteId = types.StringNull()
	var tuntermApSubnets = types.ListNull(types.StringType)
	var tuntermDhcpdConfig = types.MapNull(TuntermDhcpdConfigValue{}.Type(ctx))
	var tuntermExtraRoutes = types.MapNull(TuntermExtraRoutesValue{}.Type(ctx))
	var tuntermHosts = types.ListNull(types.StringType)
	var tuntermHostsOrder = types.ListNull(types.Int64Type)
	var tuntermHostsSelection = types.StringNull()
	var tuntermMonitoring = types.ListNull(types.ListType{ElemType: types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"host":        types.StringType,
			"port":        types.Int64Type,
			"protocol":    types.StringType,
			"src_vlan_id": types.Int64Type,
			"timeout":     types.Int64Type,
		},
	}})
	var tuntermMonitoringDisabled = types.BoolNull()

	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if !mistutils.IsSdkDataEmpty(data.MistDas) {
		mistDas = mistDasSdkToTerraform(ctx, &diags, data.MistDas)
	}
	if !mistutils.IsSdkDataEmpty(data.MistNac) {
		mistNac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	if !mistutils.IsSdkDataEmpty(data.MxedgeMgmt) {
		mxedgeMgmt = mxedgeMgmtSdkToTerraform(ctx, &diags, data.MxedgeMgmt)
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if !mistutils.IsSdkDataEmpty(data.Proxy) {
		proxy = proxySdkToTerraform(ctx, &diags, data.Proxy)
	}
	if !mistutils.IsSdkDataEmpty(data.Radsec) {
		radsec = radsecSdkToTerraform(ctx, &diags, data.Radsec)
	}
	if !mistutils.IsSdkDataEmpty(data.RadsecTls) {
		radsecTls = radsecTlsSdkToTerraform(ctx, &diags, data.RadsecTls)
	}
	if data.SiteId != nil && data.SiteId.String() != "00000000-0000-0000-0000-000000000000" {
		siteId = types.StringValue(data.SiteId.String())
	}
	if len(data.TuntermApSubnets) > 0 {
		tuntermApSubnets = mistutils.ListOfStringSdkToTerraform(data.TuntermApSubnets)
	}
	if data.TuntermDhcpdConfig != nil {
		tuntermDhcpdConfig = tuntermDhcpdConfigSdkToTerraform(ctx, &diags, data.TuntermDhcpdConfig)
	}
	if data.TuntermExtraRoutes != nil {
		tuntermExtraRoutes = tuntermExtraRoutesSdkToTerraform(ctx, &diags, data.TuntermExtraRoutes)
	}
	if len(data.TuntermHosts) > 0 {
		tuntermHosts = mistutils.ListOfStringSdkToTerraform(data.TuntermHosts)
	}
	if data.TuntermHostsOrder != nil {
		tuntermHostsOrder = listOfInt64SdkToTerraform(data.TuntermHostsOrder)
	}
	if data.TuntermHostsSelection != nil {
		tuntermHostsSelection = types.StringValue(string(*data.TuntermHostsSelection))
	}
	if data.TuntermMonitoring != nil {
		tuntermMonitoring = tuntermMonitoringSdkToTerraform(ctx, &diags, data.TuntermMonitoring)
	}
	if data.TuntermMonitoringDisabled != nil {
		tuntermMonitoringDisabled = types.BoolValue(*data.TuntermMonitoringDisabled)
	}

	state.Id = id
	state.MistDas = mistDas
	state.MistNac = mistNac
	state.MxedgeMgmt = mxedgeMgmt
	state.Name = name
	state.OrgId = orgId
	state.Proxy = proxy
	state.Radsec = radsec
	state.RadsecTls = radsecTls
	state.SiteId = siteId
	state.TuntermApSubnets = tuntermApSubnets
	state.TuntermDhcpdConfig = tuntermDhcpdConfig
	state.TuntermExtraRoutes = tuntermExtraRoutes
	state.TuntermHosts = tuntermHosts
	state.TuntermHostsOrder = tuntermHostsOrder
	state.TuntermHostsSelection = tuntermHostsSelection
	state.TuntermMonitoring = tuntermMonitoring
	state.TuntermMonitoringDisabled = tuntermMonitoringDisabled

	return state, diags
}

func listOfInt64SdkToTerraform(list []int) types.List {
	var data []attr.Value
	for _, item := range list {
		data = append(data, types.Int64Value(int64(item)))
	}
	return types.ListValueMust(types.Int64Type, data)
}
