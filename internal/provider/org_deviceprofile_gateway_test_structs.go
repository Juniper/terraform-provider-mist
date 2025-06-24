package provider

type OrgDeviceprofileGatewayModel struct {
	AdditionalConfigCmds  []string                        `hcl:"additional_config_cmds"`
	BgpConfig             map[string]BgpConfigValue       `hcl:"bgp_config"`
	DhcpdConfig           *DhcpdConfigValue               `hcl:"dhcpd_config"`
	DnsOverride           *bool                           `hcl:"dns_override"`
	DnsServers            []string                        `hcl:"dns_servers"`
	DnsSuffix             []string                        `hcl:"dns_suffix"`
	ExtraRoutes           map[string]ExtraRoutesValue     `hcl:"extra_routes"`
	ExtraRoutes6          map[string]ExtraRoutes6Value    `hcl:"extra_routes6"`
	IdpProfiles           map[string]IdpProfilesValue     `hcl:"idp_profiles"`
	IpConfigs             map[string]IpConfigsValue       `hcl:"ip_configs"`
	Name                  string                          `hcl:"name"`
	Networks              []NetworksValue                 `hcl:"networks"`
	NtpOverride           *bool                           `hcl:"ntp_override"`
	NtpServers            []string                        `hcl:"ntp_servers"`
	OobIpConfig           *OobIpConfigValue               `hcl:"oob_ip_config"`
	OrgId                 string                          `hcl:"org_id"`
	PathPreferences       map[string]PathPreferencesValue `hcl:"path_preferences"`
	PortConfig            map[string]PortConfigValue      `hcl:"port_config"`
	RouterId              *string                         `hcl:"router_id"`
	RoutingPolicies       map[string]RoutingPoliciesValue `hcl:"routing_policies"`
	ServicePolicies       []ServicePoliciesValue          `hcl:"service_policies"`
	TunnelConfigs         map[string]TunnelConfigsValue   `hcl:"tunnel_configs"`
	TunnelProviderOptions *TunnelProviderOptionsValue     `hcl:"tunnel_provider_options"`
	VrfConfig             *VrfConfigValue                 `hcl:"vrf_config"`
	VrfInstances          map[string]VrfInstancesValue    `hcl:"vrf_instances"`
}
