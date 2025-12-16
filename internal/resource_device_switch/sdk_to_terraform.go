package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, data *models.DeviceSwitch) (DeviceSwitchModel, diag.Diagnostics) {
	var state DeviceSwitchModel
	var diags diag.Diagnostics

	var aclPolicies = types.ListNull(AclPoliciesValue{}.Type(ctx))
	var aclTags = types.MapNull(AclTagsValue{}.Type(ctx))
	var additionalConfigCmds = types.ListNull(types.StringType)
	var bgpConfig = types.MapNull(BgpConfigValue{}.Type(ctx))
	var dhcpSnooping = NewDhcpSnoopingValueNull()
	var dhcpdConfig = NewDhcpdConfigValueNull()
	var deviceId types.String
	var disableAutoConfig types.Bool
	var dnsServers = types.ListValueMust(types.StringType, []attr.Value{})
	var dnsSuffix = types.ListValueMust(types.StringType, []attr.Value{})
	var extraRoutes = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extraRoutes6 = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var image1Url = types.StringValue("not_present")
	var image2Url = types.StringValue("not_present")
	var image3Url = types.StringValue("not_present")
	var ipConfig = NewIpConfigValueNull()
	var localPortConfig = types.MapNull(LocalPortConfigValue{}.Type(ctx))
	var managed types.Bool
	var mapId types.String
	var mistNac = NewMistNacValueNull()
	var name types.String
	var notes types.String
	var networks = types.MapNull(NetworksValue{}.Type(ctx))
	var ntpServers = types.ListValueMust(types.StringType, []attr.Value{})
	var oobIpConfig = NewOobIpConfigValueNull()
	var ospfAreas = types.MapNull(OspfAreasValue{}.Type(ctx))
	var ospfConfig = NewOspfConfigValueNull()
	var otherIpConfigs = types.MapNull(OtherIpConfigsValue{}.Type(ctx))
	var orgId types.String
	var portConfig = types.MapNull(PortConfigValue{}.Type(ctx))
	var portConfigOverwrite = types.MapNull(PortConfigOverwriteValue{}.Type(ctx))
	var portMirroring = types.MapNull(PortMirroringValue{}.Type(ctx))
	var portUsages = types.MapNull(PortUsagesValue{}.Type(ctx))
	var radiusConfig = NewRadiusConfigValueNull()
	var remoteSyslog = NewRemoteSyslogValueNull()
	var role types.String
	var routerId types.String
	var siteId types.String
	var snmpConfig = NewSnmpConfigValueNull()
	var stpConfig = NewStpConfigValueNull()
	var switchMgmt = NewSwitchMgmtValueNull()
	var useRouterIdAsSourceIp types.Bool
	var vars = types.MapNull(types.StringType)
	var virtualChassis = NewVirtualChassisValueNull()
	var vrfConfig = NewVrfConfigValueNull()
	var vrfInstances = types.MapNull(VrfInstancesValue{}.Type(ctx))
	var vrrpConfig = NewVrrpConfigValueNull()
	var x types.Float64
	var y types.Float64

	var deviceType types.String
	var serial types.String
	var mac types.String
	var model types.String

	if data.AclPolicies != nil {
		aclPolicies = aclPoliciesSdkToTerraform(ctx, &diags, data.AclPolicies)
	}
	if len(data.AclTags) > 0 {
		aclTags = aclTagsSdkToTerraform(ctx, &diags, data.AclTags)
	}
	if data.AdditionalConfigCmds != nil {
		additionalConfigCmds = mistutils.ListOfStringSdkToTerraform(data.AdditionalConfigCmds)
	}
	if len(data.BgpConfig) > 0 {
		bgpConfig = bgpConfigSdkToTerraform(ctx, &diags, data.BgpConfig)
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
	if data.DisableAutoConfig != nil {
		disableAutoConfig = types.BoolValue(*data.DisableAutoConfig)
	}
	if data.DnsSuffix != nil {
		dnsSuffix = mistutils.ListOfStringSdkToTerraform(data.DnsSuffix)
	}
	if len(data.ExtraRoutes) > 0 {
		extraRoutes = extraRoutesSdkToTerraform(ctx, &diags, data.ExtraRoutes)
	}
	if len(data.ExtraRoutes6) > 0 {
		extraRoutes6 = extraRoutes6SdkToTerraform(ctx, &diags, data.ExtraRoutes6)
	}
	if data.Id != nil {
		deviceId = types.StringValue(data.Id.String())
	}
	if data.Image1Url.Value() != nil {
		image1Url = types.StringValue("present")
	}
	if data.Image2Url.Value() != nil {
		image2Url = types.StringValue("present")
	}
	if data.Image3Url.Value() != nil {
		image3Url = types.StringValue("present")
	}
	if data.IpConfig != nil {
		ipConfig = ipConfigSdkToTerraform(ctx, &diags, data.IpConfig)
	}
	if data.LocalPortConfig != nil {
		localPortConfig = localPortConfigSdkToTerraform(ctx, &diags, data.LocalPortConfig)
	}
	if data.Managed != nil {
		managed = types.BoolValue(*data.Managed)
	}
	if data.MapId != nil {
		mapId = types.StringValue(data.MapId.String())
	}
	if data.MistNac != nil {
		mistNac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.Notes != nil {
		notes = types.StringValue(*data.Notes)
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
	if data.OspfConfig != nil {
		ospfConfig = ospfConfigSdkToTerraform(ctx, &diags, data.OspfConfig)
	}
	if len(data.OtherIpConfigs) > 0 {
		otherIpConfigs = otherIpConfigsSdkToTerraform(ctx, &diags, data.OtherIpConfigs)
	}
	if len(data.PortConfig) > 0 {
		portConfig = portConfigSdkToTerraform(ctx, &diags, data.PortConfig)
	}
	if len(data.PortConfigOverwrite) > 0 {
		portConfigOverwrite = portConfigOverwriteSdkToTerraform(ctx, &diags, data.PortConfigOverwrite)
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
	if data.Role != nil {
		role = types.StringValue(*data.Role)
	}
	if data.RouterId != nil {
		routerId = types.StringValue(*data.RouterId)
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
	if data.VirtualChassis != nil {
		virtualChassis = virtualChassisSdkToTerraform(ctx, &diags, data.VirtualChassis)
	}
	if len(data.Vars) > 0 {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
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
	if data.X != nil {
		x = types.Float64Value(*data.X)
	}
	if data.Y != nil {
		y = types.Float64Value(*data.Y)
	}

	deviceType = types.StringValue(data.Type)

	if data.Serial != nil {
		serial = types.StringValue(*data.Serial)
	}

	if data.Mac != nil {
		mac = types.StringValue(*data.Mac)
	}

	if data.Model != nil {
		model = types.StringValue(*data.Model)
	}

	state.AclPolicies = aclPolicies
	state.AclTags = aclTags
	state.AdditionalConfigCmds = additionalConfigCmds
	state.BgpConfig = bgpConfig
	state.DeviceId = deviceId
	state.DhcpSnooping = dhcpSnooping
	state.DhcpdConfig = dhcpdConfig
	state.DisableAutoConfig = disableAutoConfig
	state.DnsServers = dnsServers
	state.DnsSuffix = dnsSuffix
	state.ExtraRoutes = extraRoutes
	state.ExtraRoutes6 = extraRoutes6
	state.Image1Url = image1Url
	state.Image2Url = image2Url
	state.Image3Url = image3Url
	state.IpConfig = ipConfig
	state.LocalPortConfig = localPortConfig
	state.Managed = managed
	state.MapId = mapId
	state.MistNac = mistNac
	state.Name = name
	state.Notes = notes
	state.NtpServers = ntpServers
	state.Networks = networks
	state.OobIpConfig = oobIpConfig
	state.OrgId = orgId
	state.OspfAreas = ospfAreas
	state.OspfConfig = ospfConfig
	state.OtherIpConfigs = otherIpConfigs
	state.PortConfig = portConfig
	state.PortConfigOverwrite = portConfigOverwrite
	state.PortMirroring = portMirroring
	state.PortUsages = portUsages
	state.RadiusConfig = radiusConfig
	state.RemoteSyslog = remoteSyslog
	state.Role = role
	state.RouterId = routerId
	state.SiteId = siteId
	state.SnmpConfig = snmpConfig
	state.StpConfig = stpConfig
	state.SwitchMgmt = switchMgmt
	state.UseRouterIdAsSourceIp = useRouterIdAsSourceIp
	state.Vars = vars
	state.VirtualChassis = virtualChassis
	state.VrfConfig = vrfConfig
	state.VrfInstances = vrfInstances
	state.VrrpConfig = vrrpConfig
	state.X = x
	state.Y = y
	state.Type = deviceType
	state.Serial = serial
	state.Mac = mac
	state.Model = model

	return state, diags
}
