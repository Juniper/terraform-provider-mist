package provider

type SiteNetworktemplateModel struct {
	AclPolicies                     []AclPoliciesValue                           `hcl:"acl_policies"`
	AclTags                         map[string]AclTagsValue                      `hcl:"acl_tags"`
	AdditionalConfigCmds            []string                                     `hcl:"additional_config_cmds"`
	DhcpSnooping                    *DhcpSnoopingValue                           `hcl:"dhcp_snooping"`
	DisabledSystemDefinedPortUsages []string                                     `hcl:"disabled_system_defined_port_usages"`
	DnsServers                      []string                                     `hcl:"dns_servers"`
	DnsSuffix                       []string                                     `hcl:"dns_suffix"`
	ExtraRoutes                     map[string]ExtraRoutesValue                  `hcl:"extra_routes"`
	ExtraRoutes6                    map[string]ExtraRoutes6Value                 `hcl:"extra_routes6"`
	MistNac                         *MistNacValue                                `hcl:"mist_nac"`
	Networks                        map[string]NetworksValue                     `hcl:"networks"`
	NtpServers                      []string                                     `hcl:"ntp_servers"`
	OspfAreas                       map[string]SiteNetworktemplateOspfAreasValue `hcl:"ospf_areas"`
	PortMirroring                   map[string]PortMirroringValue                `hcl:"port_mirroring"`
	PortUsages                      map[string]OrgNetworktemplatePortUsagesValue `hcl:"port_usages"`
	RadiusConfig                    *RadiusConfigValue                           `hcl:"radius_config"`
	RemoteSyslog                    *RemoteSyslogValue                           `hcl:"remote_syslog"`
	RemoveExistingConfigs           *bool                                        `hcl:"remove_existing_configs"`
	SiteId                          string                                       `hcl:"site_id"`
	SnmpConfig                      *SnmpConfigValue                             `hcl:"snmp_config"`
	SwitchMatching                  *SwitchMatchingValue                         `hcl:"switch_matching"`
	SwitchMgmt                      *SwitchMgmtValue                             `hcl:"switch_mgmt"`
	VrfConfig                       *VrfConfigValue                              `hcl:"vrf_config"`
	VrfInstances                    map[string]VrfInstancesValue                 `hcl:"vrf_instances"`
}

type SiteNetworktemplateOspfAreasValue struct {
	IncludeLoopback *bool                        `cty:"include_loopback"`
	OspfNetworks    map[string]OspfNetworksValue `cty:"ospf_networks"`
	OspfAreasType   *string                      `cty:"type"`
}
