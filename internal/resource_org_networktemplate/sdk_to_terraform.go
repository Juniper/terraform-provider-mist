package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, data models.NetworkTemplate) (OrgNetworktemplateModel, diag.Diagnostics) {
	var state OrgNetworktemplateModel
	var diags diag.Diagnostics

	var aclPolicies = types.ListNull(AclPoliciesValue{}.Type(ctx))
	var aclTags = types.MapNull(AclTagsValue{}.Type(ctx))
	var additionalConfigCmds = types.ListNull(types.StringType)
	var dhcpSnooping = NewDhcpSnoopingValueNull()
	var dnsServers = types.ListNull(types.StringType)
	var dnsSuffix = types.ListNull(types.StringType)
	var extraRoutes = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extraRoutes6 = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var id types.String
	var mistNac = NewMistNacValueNull()
	var name types.String
	var networks = types.MapNull(NetworksValue{}.Type(ctx))
	var ntpServers = types.ListNull(types.StringType)
	var orgId types.String
	var ospfAreas = types.MapNull(OspfAreasValue{}.Type(ctx))
	var portMirroring = types.MapNull(PortMirroringValue{}.Type(ctx))
	var portUsages = types.MapNull(PortUsagesValue{}.Type(ctx))
	var radiusConfig = NewRadiusConfigValueNull()
	var remoteSyslog = NewRemoteSyslogValueNull()
	var removeExistingConfigs = types.BoolValue(false)
	var snmpConfig = NewSnmpConfigValueNull()
	var switchMatching = NewSwitchMatchingValueNull()
	var switchMgmt = NewSwitchMgmtValueNull()
	var vrfConfig = NewVrfConfigValueNull()
	var vrfInstances = types.MapNull(VrfInstancesValue{}.Type(ctx))

	if data.AclPolicies != nil {
		aclPolicies = aclPoliciesSdkToTerraform(ctx, &diags, data.AclPolicies)
	}
	if data.AclTags != nil && len(data.AclTags) > 0 {
		aclTags = aclTagsSdkToTerraform(ctx, &diags, data.AclTags)
	}
	if data.AdditionalConfigCmds != nil {
		additionalConfigCmds = misttransform.ListOfStringSdkToTerraform(data.AdditionalConfigCmds)
	}
	if data.DhcpSnooping != nil {
		dhcpSnooping = dhcpSnoopingSdkToTerraform(ctx, &diags, data.DhcpSnooping)
	}
	if data.DnsServers != nil {
		dnsServers = misttransform.ListOfStringSdkToTerraform(data.DnsServers)
	}
	if data.DnsSuffix != nil {
		dnsSuffix = misttransform.ListOfStringSdkToTerraform(data.DnsSuffix)
	}
	if data.ExtraRoutes != nil && len(data.ExtraRoutes) > 0 {
		extraRoutes = extraRoutesSdkToTerraform(ctx, &diags, data.ExtraRoutes)
	}
	if data.ExtraRoutes6 != nil && len(data.ExtraRoutes6) > 0 {
		extraRoutes6 = extraRoutes6SdkToTerraform(ctx, &diags, data.ExtraRoutes6)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.MistNac != nil {
		mistNac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.Networks != nil && len(data.Networks) > 0 {
		networks = NetworksSdkToTerraform(ctx, &diags, data.Networks)
	}
	if data.NtpServers != nil {
		ntpServers = misttransform.ListOfStringSdkToTerraform(data.NtpServers)
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.OspfAreas != nil {
		ospfAreas = ospfAreasSdkToTerraform(ctx, &diags, data.OspfAreas)
	}
	if data.PortMirroring != nil {
		portMirroring = portMirroringSdkToTerraform(ctx, &diags, data.PortMirroring)
	}
	if data.PortUsages != nil && len(data.PortUsages) > 0 {
		portUsages = portUsagesSdkToTerraform(ctx, &diags, data.PortUsages)
	}
	if data.RadiusConfig != nil {
		radiusConfig = radiusConfigSdkToTerraform(ctx, &diags, data.RadiusConfig)
	}
	if data.RemoteSyslog != nil {
		remoteSyslog = remoteSyslogSdkToTerraform(ctx, &diags, data.RemoteSyslog)
	}
	if data.RemoveExistingConfigs != nil {
		state.RemoveExistingConfigs = types.BoolValue(*data.RemoveExistingConfigs)
	}
	if data.SnmpConfig != nil {
		snmpConfig = snmpConfigSdkToTerraform(ctx, &diags, data.SnmpConfig)
	}
	if data.SwitchMatching != nil {
		switchMatching = switchMatchingSdkToTerraform(ctx, &diags, data.SwitchMatching)
	}
	if data.SwitchMgmt != nil {
		switchMgmt = switchMgmtSdkToTerraform(ctx, &diags, data.SwitchMgmt)
	}
	if data.VrfConfig != nil {
		vrfConfig = vrfConfigSdkToTerraform(ctx, &diags, data.VrfConfig)
	}
	if data.VrfInstances != nil && len(data.VrfInstances) > 0 {
		vrfInstances = vrfInstancesSdkToTerraform(ctx, &diags, data.VrfInstances)
	}

	state.Id = id
	state.OrgId = orgId
	state.Name = name
	state.AclPolicies = aclPolicies
	state.AclTags = aclTags
	state.AdditionalConfigCmds = additionalConfigCmds
	state.DhcpSnooping = dhcpSnooping
	state.DnsServers = dnsServers
	state.DnsSuffix = dnsSuffix
	state.ExtraRoutes = extraRoutes
	state.ExtraRoutes6 = extraRoutes6
	state.MistNac = mistNac
	state.NtpServers = ntpServers
	state.Networks = networks
	state.OspfAreas = ospfAreas
	state.PortMirroring = portMirroring
	state.PortUsages = portUsages
	state.RadiusConfig = radiusConfig
	state.RemoteSyslog = remoteSyslog
	state.RemoveExistingConfigs = removeExistingConfigs
	state.SnmpConfig = snmpConfig
	state.SwitchMatching = switchMatching
	state.SwitchMgmt = switchMgmt
	state.VrfConfig = vrfConfig
	state.VrfInstances = vrfInstances

	return state, diags
}
