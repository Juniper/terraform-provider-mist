package resource_org_deviceprofile_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, data *models.DeviceprofileSwitch) (OrgDeviceprofileSwitchModel, diag.Diagnostics) {
	var state OrgDeviceprofileSwitchModel
	var diags diag.Diagnostics

	var aclPolicies = types.ListNull(AclPoliciesValue{}.Type(ctx))
	var aclTags = types.MapNull(AclTagsValue{}.Type(ctx))
	var additionalConfigCmds = types.ListNull(types.StringType)
	var dhcpSnooping = NewDhcpSnoopingValueNull()
	var dhcpdConfig = NewDhcpdConfigValueNull()
	var dnsServers = types.ListValueMust(types.StringType, []attr.Value{})
	var dnsSuffix = types.ListValueMust(types.StringType, []attr.Value{})
	var evpnConfig = NewEvpnConfigValueNull()
	var extraRoutes = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extraRoutes6 = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var id types.String
	var iotConfig = types.MapNull(IotConfigValue{}.Type(ctx))
	var ipConfig = NewIpConfigValueNull()
	var mistNac = NewMistNacValueNull()
	var name types.String
	var networks = types.MapNull(NetworksValue{}.Type(ctx))
	var ntpServers = types.ListValueMust(types.StringType, []attr.Value{})
	var oobIpConfig = NewOobIpConfigValueNull()
	var orgId types.String
	var ospfAreas = types.MapNull(OspfAreasValue{}.Type(ctx))
	var otherIpConfigs = types.MapNull(OtherIpConfigsValue{}.Type(ctx))
	var portConfig = types.MapNull(PortConfigValue{}.Type(ctx))
	var portMirroring = types.MapNull(PortMirroringValue{}.Type(ctx))
	var portUsages = types.MapNull(PortUsagesValue{}.Type(ctx))
	var radiusConfig = NewRadiusConfigValueNull()
	var remoteSyslog = NewRemoteSyslogValueNull()
	var routingPolicies = types.MapNull(RoutingPoliciesValue{}.Type(ctx))
	var siteId types.String
	var snmpConfig = NewSnmpConfigValueNull()
	var stpConfig = NewStpConfigValueNull()
	var switchMgmt = NewSwitchMgmtValueNull()
	var deviceType types.String
	var useRouterIdAsSourceIp types.Bool
	var vrfConfig = NewVrfConfigValueNull()
	var vrfInstances = types.MapNull(VrfInstancesValue{}.Type(ctx))
	var vrrpConfig = NewVrrpConfigValueNull()

	if data.AclPolicies != nil {
		aclPolicies = aclPoliciesSdkToTerraform(ctx, &diags, data.AclPolicies)
	}
	if len(data.AclTags) > 0 {
		aclTags = aclTagsSdkToTerraform(ctx, &diags, data.AclTags)
	}
	if data.AdditionalConfigCmds != nil {
		additionalConfigCmds = mistutils.ListOfStringSdkToTerraform(data.AdditionalConfigCmds)
	}
	if data.DhcpSnooping != nil {
		dhcpSnooping = dhcpSnoopingSdkToTerraform(ctx, &diags, data.DhcpSnooping)
	}
	if data.DhcpdConfig != nil {
		dhcpdConfig = dhcpdConfigSdkToTerraform(ctx, &diags, data.DhcpdConfig)
	}
	if data.DnsServers != nil {
		dnsServers = mistutils.ListOfStringSdkToTerraform(data.DnsServers)
	}
	if data.DnsSuffix != nil {
		dnsSuffix = mistutils.ListOfStringSdkToTerraform(data.DnsSuffix)
	}
	if data.EvpnConfig != nil {
		evpnConfig = evpnConfigSdkToTerraform(ctx, &diags, data.EvpnConfig)
	}
	if len(data.ExtraRoutes) > 0 {
		extraRoutes = extraRoutesSdkToTerraform(ctx, &diags, data.ExtraRoutes)
	}
	if len(data.ExtraRoutes6) > 0 {
		extraRoutes6 = extraRoutes6SdkToTerraform(ctx, &diags, data.ExtraRoutes6)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if len(data.IotConfig) > 0 {
		iotConfig = iotConfigSdkToTerraform(ctx, &diags, data.IotConfig)
	}
	if data.IpConfig != nil {
		ipConfig = ipConfigSdkToTerraform(ctx, &diags, data.IpConfig)
	}
	if data.MistNac != nil {
		mistNac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	if data.Name != "" {
		name = types.StringValue(data.Name)
	}
	if len(data.Networks) > 0 {
		networks = NetworksSdkToTerraform(ctx, &diags, data.Networks)
	}
	if data.NtpServers != nil {
		ntpServers = mistutils.ListOfStringSdkToTerraform(data.NtpServers)
	}
	if data.OobIpConfig != nil {
		oobIpConfig = oobIpConfigsSdkToTerraform(ctx, &diags, data.OobIpConfig)
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.OspfAreas != nil {
		ospfAreas = ospfAreasSdkToTerraform(ctx, &diags, data.OspfAreas)
	}
	if len(data.OtherIpConfigs) > 0 {
		otherIpConfigs = otherIpConfigsSdkToTerraform(ctx, &diags, data.OtherIpConfigs)
	}
	if len(data.PortConfig) > 0 {
		portConfig = portConfigSdkToTerraform(ctx, &diags, data.PortConfig)
	}
	if len(data.PortMirroring) > 0 {
		portMirroring = portMirroringSdkToTerraform(ctx, &diags, data.PortMirroring)
	}
	if len(data.PortUsages) > 0 {
		portUsages = portUsagesSdkToTerraform(ctx, &diags, data.PortUsages)
	}
	if data.RadiusConfig != nil {
		radiusConfig = radiusConfigSdkToTerraform(ctx, &diags, data.RadiusConfig)
	}
	if data.RemoteSyslog != nil {
		remoteSyslog = remoteSyslogSdkToTerraform(ctx, &diags, data.RemoteSyslog)
	}
	if data.RoutingPolicies != nil {
		routingPolicies = routingPoliciesSdkToTerraform(ctx, &diags, data.RoutingPolicies)
	}
	if data.SiteId != nil {
		siteId = types.StringValue(data.SiteId.String())
	}
	if data.SnmpConfig != nil {
		snmpConfig = snmpConfigSdkToTerraform(ctx, &diags, data.SnmpConfig)
	}
	if data.StpConfig != nil {
		stpConfig = stpConfigSdkToTerraform(ctx, &diags, *data.StpConfig)
	}
	if data.SwitchMgmt != nil {
		switchMgmt = switchMgmtSdkToTerraform(ctx, &diags, data.SwitchMgmt)
	}
	if data.UseRouterIdAsSourceIp != nil {
		useRouterIdAsSourceIp = types.BoolValue(*data.UseRouterIdAsSourceIp)
	}
	if data.VrfConfig != nil {
		vrfConfig = vrfConfigSdkToTerraform(ctx, &diags, data.VrfConfig)
	}
	if len(data.VrfInstances) > 0 {
		vrfInstances = vrfInstancesSdkToTerraform(ctx, &diags, data.VrfInstances)
	}
	if data.VrrpConfig != nil {
		vrrpConfig = vrrpConfigInstancesSdkToTerraform(ctx, &diags, data.VrrpConfig)
	}

	deviceType = types.StringValue(data.Type)

	state.AclPolicies = aclPolicies
	state.AclTags = aclTags
	state.AdditionalConfigCmds = additionalConfigCmds
	state.DhcpSnooping = dhcpSnooping
	state.DhcpdConfig = dhcpdConfig
	state.DnsServers = dnsServers
	state.DnsSuffix = dnsSuffix
	state.EvpnConfig = evpnConfig
	state.ExtraRoutes = extraRoutes
	state.ExtraRoutes6 = extraRoutes6
	state.Id = id
	state.IotConfig = iotConfig
	state.IpConfig = ipConfig
	state.MistNac = mistNac
	state.Name = name
	state.Networks = networks
	state.NtpServers = ntpServers
	state.OobIpConfig = oobIpConfig
	state.OrgId = orgId
	state.OspfAreas = ospfAreas
	state.OtherIpConfigs = otherIpConfigs
	state.PortConfig = portConfig
	state.PortMirroring = portMirroring
	state.PortUsages = portUsages
	state.RadiusConfig = radiusConfig
	state.RemoteSyslog = remoteSyslog
	state.RoutingPolicies = routingPolicies
	state.SiteId = siteId
	state.SnmpConfig = snmpConfig
	state.StpConfig = stpConfig
	state.SwitchMgmt = switchMgmt
	state.Type = deviceType
	state.UseRouterIdAsSourceIp = useRouterIdAsSourceIp
	state.VrfConfig = vrfConfig
	state.VrfInstances = vrfInstances
	state.VrrpConfig = vrrpConfig

	return state, diags
}
