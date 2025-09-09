package provider

type OrgDeviceprofileGatewayModel struct {
	AdditionalConfigCmds  []string                                          `hcl:"additional_config_cmds"`
	BgpConfig             map[string]BgpConfigValue                         `hcl:"bgp_config"`
	DhcpdConfig           *DhcpdConfigValue                                 `hcl:"dhcpd_config"`
	DnsOverride           *bool                                             `hcl:"dns_override"`
	DnsServers            []string                                          `hcl:"dns_servers"`
	DnsSuffix             []string                                          `hcl:"dns_suffix"`
	ExtraRoutes           map[string]ExtraRoutesValue                       `hcl:"extra_routes"`
	ExtraRoutes6          map[string]ExtraRoutes6Value                      `hcl:"extra_routes6"`
	IdpProfiles           map[string]IdpProfilesValue                       `hcl:"idp_profiles"`
	IpConfigs             map[string]IpConfigsValue                         `hcl:"ip_configs"`
	Name                  string                                            `hcl:"name"`
	Networks              []OrgDeviceprofileGatewayNetworksValue            `hcl:"networks"`
	NtpOverride           *bool                                             `hcl:"ntp_override"`
	NtpServers            []string                                          `hcl:"ntp_servers"`
	OobIpConfig           *OobIpConfigValue                                 `hcl:"oob_ip_config"`
	OrgId                 string                                            `hcl:"org_id"`
	PathPreferences       map[string]PathPreferencesValue                   `hcl:"path_preferences"`
	PortConfig            map[string]OrgDeviceprofileGatewayPortConfigValue `hcl:"port_config"`
	RouterId              *string                                           `hcl:"router_id"`
	RoutingPolicies       map[string]RoutingPoliciesValue                   `hcl:"routing_policies"`
	ServicePolicies       []ServicePoliciesValue                            `hcl:"service_policies"`
	TunnelConfigs         map[string]TunnelConfigsValue                     `hcl:"tunnel_configs"`
	TunnelProviderOptions *TunnelProviderOptionsValue                       `hcl:"tunnel_provider_options"`
	VrfConfig             *VrfConfigValue                                   `hcl:"vrf_config"`
	VrfInstances          map[string]VrfInstancesValue                      `hcl:"vrf_instances"`
}

type OrgDeviceprofileGatewayNetworksValue struct {
	DisallowMistServices *bool                     `cty:"disallow_mist_services" hcl:"disallow_mist_services"`
	Gateway              *string                   `cty:"gateway" hcl:"gateway"`
	Gateway6             *string                   `cty:"gateway6" hcl:"gateway6"`
	InternalAccess       *InternalAccessValue      `cty:"internal_access" hcl:"internal_access"`
	InternetAccess       *InternetAccessValue      `cty:"internet_access" hcl:"internet_access"`
	Isolation            *bool                     `cty:"isolation" hcl:"isolation"`
	Multicast            *MulticastValue           `cty:"multicast" hcl:"multicast"`
	Name                 *string                   `cty:"name" hcl:"name"`
	RoutedForNetworks    []string                  `cty:"routed_for_networks" hcl:"routed_for_networks"`
	Subnet               string                    `cty:"subnet" hcl:"subnet"`
	Subnet6              *string                   `cty:"subnet6" hcl:"subnet6"`
	Tenants              map[string]TenantsValue   `cty:"tenants" hcl:"tenants"`
	VlanId               *string                   `cty:"vlan_id" hcl:"vlan_id"`
	VpnAccess            map[string]VpnAccessValue `cty:"vpn_access" hcl:"vpn_access"`
}

type OrgDeviceprofileGatewayPortConfigValue struct {
	AeDisableLacp       *bool                          `cty:"ae_disable_lacp" hcl:"ae_disable_lacp"`
	AeIdx               *string                        `cty:"ae_idx" hcl:"ae_idx"`
	AeLacpForceUp       *bool                          `cty:"ae_lacp_force_up" hcl:"ae_lacp_force_up"`
	Aggregated          *bool                          `cty:"aggregated" hcl:"aggregated"`
	Critical            *bool                          `cty:"critical" hcl:"critical"`
	Description         *string                        `cty:"description" hcl:"description"`
	DisableAutoneg      *bool                          `cty:"disable_autoneg" hcl:"disable_autoneg"`
	Disabled            *bool                          `cty:"disabled" hcl:"disabled"`
	DslType             *string                        `cty:"dsl_type" hcl:"dsl_type"`
	DslVci              *int64                         `cty:"dsl_vci" hcl:"dsl_vci"`
	DslVpi              *int64                         `cty:"dsl_vpi" hcl:"dsl_vpi"`
	Duplex              *string                        `cty:"duplex" hcl:"duplex"`
	LteApn              *string                        `cty:"lte_apn" hcl:"lte_apn"`
	LteAuth             *string                        `cty:"lte_auth" hcl:"lte_auth"`
	LteBackup           *bool                          `cty:"lte_backup" hcl:"lte_backup"`
	LtePassword         *string                        `cty:"lte_password" hcl:"lte_password"`
	LteUsername         *string                        `cty:"lte_username" hcl:"lte_username"`
	Mtu                 *int64                         `cty:"mtu" hcl:"mtu"`
	Name                *string                        `cty:"name" hcl:"name"`
	Networks            []string                       `cty:"networks" hcl:"networks"`
	OuterVlanId         *int64                         `cty:"outer_vlan_id" hcl:"outer_vlan_id"`
	PoeDisabled         *bool                          `cty:"poe_disabled" hcl:"poe_disabled"`
	PortIpConfig        *PortIpConfigValue             `cty:"ip_config" hcl:"ip_config"`
	PortNetwork         *string                        `cty:"port_network" hcl:"port_network"`
	PreserveDscp        *bool                          `cty:"preserve_dscp" hcl:"preserve_dscp"`
	Redundant           *bool                          `cty:"redundant" hcl:"redundant"`
	RedundantGroup      *int64                         `cty:"redundant_group" hcl:"redundant_group"`
	RethIdx             *string                        `cty:"reth_idx" hcl:"reth_idx"`
	RethNode            *string                        `cty:"reth_node" hcl:"reth_node"`
	RethNodes           []string                       `cty:"reth_nodes" hcl:"reth_nodes"`
	Speed               *string                        `cty:"speed" hcl:"speed"`
	SsrNoVirtualMac     *bool                          `cty:"ssr_no_virtual_mac" hcl:"ssr_no_virtual_mac"`
	SvrPortRange        *string                        `cty:"svr_port_range" hcl:"svr_port_range"`
	TrafficShaping      *TrafficShapingValue           `cty:"traffic_shaping" hcl:"traffic_shaping"`
	Usage               string                         `cty:"usage" hcl:"usage"`
	VlanId              *string                        `cty:"vlan_id" hcl:"vlan_id"`
	VpnPaths            map[string]VpnPathsValue       `cty:"vpn_paths" hcl:"vpn_paths"`
	WanArpPolicer       *string                        `cty:"wan_arp_policer" hcl:"wan_arp_policer"`
	WanDisableSpeedtest *bool                          `cty:"wan_disable_speedtest" hcl:"wan_disable_speedtest"`
	WanExtIp            *string                        `cty:"wan_ext_ip" hcl:"wan_ext_ip"`
	WanExtraRoutes      map[string]WanExtraRoutesValue `cty:"wan_extra_routes" hcl:"wan_extra_routes"`
	WanExtraRoutes6     map[string]WanExtraRoutesValue `cty:"wan_extra_routes6" hcl:"wan_extra_routes6"`
	WanNetworks         []string                       `cty:"wan_networks" hcl:"wan_networks"`
	WanProbeOverride    *WanProbeOverrideValue         `cty:"wan_probe_override" hcl:"wan_probe_override"`
	WanSourceNat        *WanSourceNatValue             `cty:"wan_source_nat" hcl:"wan_source_nat"`
	WanType             *string                        `cty:"wan_type" hcl:"wan_type"`
}
