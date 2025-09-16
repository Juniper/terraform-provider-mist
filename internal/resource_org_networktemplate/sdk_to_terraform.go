package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, data models.NetworkTemplate) (OrgNetworktemplateModel, diag.Diagnostics) {
	var state OrgNetworktemplateModel
	var diags diag.Diagnostics

	var aclPolicies = types.ListNull(AclPoliciesValue{}.Type(ctx))
	var aclTags = types.MapNull(AclTagsValue{}.Type(ctx))
	var additionalConfigCmds = types.ListNull(types.StringType)
	var dhcpSnooping = NewDhcpSnoopingValueNull()
	var dnsServers = types.ListValueMust(types.StringType, []attr.Value{})
	var dnsSuffix = types.ListValueMust(types.StringType, []attr.Value{})
	var extraRoutes = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extraRoutes6 = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var id types.String
	var mistNac = NewMistNacValueNull()
	var name types.String
	var networks = types.MapNull(NetworksValue{}.Type(ctx))
	var ntpServers = types.ListValueMust(types.StringType, []attr.Value{})
	var orgId types.String
	var ospfAreas = types.MapNull(OspfAreasValue{}.Type(ctx))
	var portMirroring = types.MapNull(PortMirroringValue{}.Type(ctx))
	var portUsages = types.MapNull(PortUsagesValue{}.Type(ctx))
	var radiusConfig = NewRadiusConfigValueNull()
	var remoteSyslog = NewRemoteSyslogValueNull()
	var removeExistingConfigs = types.BoolNull()
	var snmpConfig = NewSnmpConfigValueNull()
	var switchMatching = NewSwitchMatchingValueNull()
	var switchMgmt = NewSwitchMgmtValueNull()
	var vrfConfig = NewVrfConfigValueNull()
	var vrfInstances = types.MapNull(VrfInstancesValue{}.Type(ctx))

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
	if data.DnsServers != nil {
		dnsServers = mistutils.ListOfStringSdkToTerraform(data.DnsServers)
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
		id = types.StringValue(data.Id.String())
	}
	if data.MistNac != nil {
		mistNac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if len(data.Networks) > 0 {
		networks = NetworksSdkToTerraform(ctx, &diags, data.Networks)
	}
	if data.NtpServers != nil {
		ntpServers = mistutils.ListOfStringSdkToTerraform(data.NtpServers)
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
	if len(data.PortUsages) > 0 {
		portUsages = portUsagesSdkToTerraform(ctx, &diags, data.PortUsages)
	}
	if data.RadiusConfig != nil {
		radiusConfig = radiusConfigSdkToTerraform(ctx, &diags, data.RadiusConfig)
	}
	if data.RemoteSyslog != nil {
		remoteSyslog = remoteSyslogSdkToTerraform(ctx, &diags, data.RemoteSyslog)
	}
	if data.RemoveExistingConfigs != nil {
		removeExistingConfigs = types.BoolValue(*data.RemoveExistingConfigs)
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
	if len(data.VrfInstances) > 0 {
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
